package chain

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/boringdao/bridge/internal/monitor"
	mnt "github.com/boringdao/bridge/internal/monitor/contracts"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/hexutil"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Monitor struct {
	lHeight     uint64
	cHeight     uint64
	rHeight     uint64
	twoWayAbi   abi.ABI
	wrapper     *Wrapper
	cocoC       chan *monitor.Coco
	logger      logrus.FieldLogger
	storage     storage.Storage
	config      *repo.BridgeConfig
	minConfirms uint64
	twoWayAddr  common.Address
	address     common.Address

	ctx    context.Context
	cancel context.CancelFunc
}

func New(repoRoot string, config *repo.BridgeConfig, logger logrus.FieldLogger) (monitor.Mnt, error) {
	storagePath := repo.GetStoragePath(repoRoot, fmt.Sprintf("%s_%d", config.Name, config.ChainID))
	ethStorage, err := leveldb.New(storagePath)
	if err != nil {
		return nil, err
	}

	privKey, err := crypto.ToECDSA(hexutil.Decode(config.PrivKey))
	if err != nil {
		return nil, err
	}

	address := crypto.PubkeyToAddress(privKey.PublicKey)

	minConfirms := 0
	if config.MinConfirms > 0 {
		minConfirms = int(config.MinConfirms)
	}

	twoWayAbi, err := abi.JSON(bytes.NewReader([]byte(mnt.TwoWayABI)))
	if err != nil {
		return nil, fmt.Errorf("abi unmarshal: %s", err.Error())
	}

	wrapper, err := NewWrapper(config, logger)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Monitor{
		config:      config,
		storage:     ethStorage,
		address:     address,
		wrapper:     wrapper,
		twoWayAbi:   twoWayAbi,
		twoWayAddr:  common.HexToAddress(config.TwoWayContract),
		minConfirms: uint64(minConfirms),
		cocoC:       make(chan *monitor.Coco),
		logger:      logger,
		ctx:         ctx,
		cancel:      cancel,
	}, nil
}

func (m *Monitor) Start() error {
	m.loadHeightFromStorage()
	go m.listenLockEvent()
	go m.listenCrossBurnEvent()
	go m.listenRollback()
	return nil
}

func (m *Monitor) Stop() error {
	m.cancel()
	m.wrapper.Close()
	return nil
}

func (m *Monitor) listenLockEvent() {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	start := m.lHeight

	for {
		select {
		case <-ticker.C:
			num, err := m.fetchBlockNum()
			if err != nil {
				continue
			}
			end := num - m.minConfirms
			if num < m.minConfirms || end < start {
				continue
			}
			if end >= start+300 {
				end = start + 300
			}
			filter := m.wrapper.FilterLock(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				m.handleLock(filter.Event)
			}

			m.logger.WithFields(logrus.Fields{"start": start, "end": end, "current": num}).Infof("CrossLockLockIterator")
			start = end + 1
		case <-m.ctx.Done():
			m.logger.Info("CrossLockLockIterator done")
			return
		}
	}
}

func (m *Monitor) listenCrossBurnEvent() {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	start := m.cHeight

	for {
		select {
		case <-ticker.C:
			num, err := m.fetchBlockNum()
			if err != nil {
				continue
			}
			end := num - m.minConfirms
			if num < m.minConfirms || end < start {
				continue
			}
			if end >= start+300 {
				end = start + 300
			}
			filter := m.wrapper.FilterCrossBurn(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				m.handleCrossBurn(filter.Event)
			}

			m.logger.WithFields(logrus.Fields{"start": start, "end": end, "current": num}).Infof("CrossBurn")
			start = end + 1
		case <-m.ctx.Done():
			m.logger.Info("CrossBurn done")
			return
		}
	}
}

func (m *Monitor) listenRollback() {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	start := m.rHeight

	for {
		select {
		case <-ticker.C:
			num, err := m.fetchBlockNum()
			if err != nil {
				continue
			}
			end := num - m.minConfirms
			if num < m.minConfirms || end < start {
				continue
			}
			if end >= start+300 {
				end = start + 300
			}
			filter := m.wrapper.FilterRollback(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				m.handleRollback(filter.Event)
			}

			m.logger.WithFields(logrus.Fields{"start": start, "end": end, "current": num}).Infof("RollbackIterator")
			start = end + 1
		case <-m.ctx.Done():
			m.logger.Info("RollbackIterator done")
			return
		}
	}
}

func (m *Monitor) HandleCocoC() chan *monitor.Coco {
	return m.cocoC
}

func (m *Monitor) handleRollback(rollback *mnt.TwoWayRollback) {
	if !strings.EqualFold(rollback.Raw.Address.String(), m.config.TwoWayContract) {
		return
	}

	if m.storage.Has(TxKey(rollback.Raw.TxHash.String(), monitor.Rollback, rollback.Raw.Index)) {
		return
	}
	coco := &monitor.Coco{
		Typ:         monitor.Rollback,
		Token0:      rollback.Token0,
		Token1:      rollback.Token1,
		ChainID0:    rollback.ChainID0,
		ChainID1:    rollback.ChainID1,
		From:        rollback.From,
		To:          rollback.To,
		Amount:      rollback.Amount,
		Index:       rollback.Raw.Index,
		TxId:        rollback.Raw.TxHash.String(),
		BlockHeight: rollback.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"token0":       coco.Token0.String(),
		"token1":       coco.Token1.String(),
		"from":         coco.From.String(),
		"to":           coco.To.String(),
		"chain0":       coco.ChainID0.String(),
		"chain1":       coco.ChainID1.String(),
		"amount":       coco.Amount.String(),
		"index":        coco.Index,
		"txId":         rollback.Raw.TxHash.String(),
		"block_height": rollback.Raw.BlockNumber,
		"removed":      rollback.Raw.Removed,
	}).Info("TwoWayRollback")

	if rollback.Raw.Removed {
		return
	}

	if !m.confirmEvent(rollback.Raw, monitor.Rollback) {
		m.logger.WithFields(logrus.Fields{
			"txId":         rollback.Raw.TxHash.String(),
			"block_height": rollback.Raw.BlockNumber,
		}).Info("TwoWay has not confirmed")
		return
	}

	m.logger.WithField("tx", rollback.Raw.TxHash.String()).Info("confirmEvent")
	m.cocoC <- coco
	m.persistRBlockHeight(rollback.Raw.TxHash.String(), rollback.Raw.BlockNumber, coco)
}

func (m *Monitor) handleLock(lock *mnt.TwoWayLock) {
	if !strings.EqualFold(lock.Raw.Address.String(), m.config.TwoWayContract) {
		return
	}

	if m.storage.Has(TxKey(lock.Raw.TxHash.String(), monitor.Lock, lock.Raw.Index)) {
		return
	}
	coco := &monitor.Coco{
		Typ:         monitor.Lock,
		Token0:      lock.Token0,
		Token1:      lock.Token1,
		ChainID0:    lock.ChainID0,
		ChainID1:    lock.ChainID1,
		From:        lock.From,
		To:          lock.To,
		Amount:      lock.Amount,
		Index:       lock.Raw.Index,
		TxId:        lock.Raw.TxHash.String(),
		BlockHeight: lock.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"token0":       coco.Token0.String(),
		"token1":       coco.Token1.String(),
		"from":         coco.From.String(),
		"to":           coco.To.String(),
		"chain0":       coco.ChainID0.String(),
		"chain1":       coco.ChainID1.String(),
		"amount":       coco.Amount.String(),
		"index":        coco.Index,
		"txId":         lock.Raw.TxHash.String(),
		"block_height": lock.Raw.BlockNumber,
		"removed":      lock.Raw.Removed,
	}).Info("TwoWayLock")

	if lock.Raw.Removed {
		return
	}

	if !m.confirmEvent(lock.Raw, monitor.Lock) {
		m.logger.WithFields(logrus.Fields{
			"txId":         lock.Raw.TxHash.String(),
			"block_height": lock.Raw.BlockNumber,
		}).Info("TwoWayLock has not confirmed")
		return
	}

	m.logger.WithField("tx", lock.Raw.TxHash.String()).Info("confirmEvent")
	m.cocoC <- coco
	m.persistLBlockHeight(lock.Raw.TxHash.String(), lock.Raw.BlockNumber, coco)
}

func (m *Monitor) handleCrossBurn(crossBurn *mnt.TwoWayCrossBurn) {
	if !strings.EqualFold(crossBurn.Raw.Address.String(), m.config.TwoWayContract) {
		return
	}

	if m.storage.Has(TxKey(crossBurn.Raw.TxHash.String(), monitor.CrossBurn, crossBurn.Raw.Index)) {
		return
	}
	coco := &monitor.Coco{
		Typ:         monitor.CrossBurn,
		Token0:      crossBurn.Token0,
		Token1:      crossBurn.Token1,
		ChainID0:    crossBurn.ChainID0,
		ChainID1:    crossBurn.ChainID1,
		From:        crossBurn.From,
		To:          crossBurn.To,
		Amount:      crossBurn.Amount,
		Index:       crossBurn.Raw.Index,
		TxId:        crossBurn.Raw.TxHash.String(),
		BlockHeight: crossBurn.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"token0":       coco.Token0.String(),
		"token1":       coco.Token1.String(),
		"from":         coco.From.String(),
		"to":           coco.To.String(),
		"chain0":       coco.ChainID0.String(),
		"chain1":       coco.ChainID1.String(),
		"amount":       coco.Amount.String(),
		"index":        coco.Index,
		"txId":         crossBurn.Raw.TxHash.String(),
		"block_height": crossBurn.Raw.BlockNumber,
		"removed":      crossBurn.Raw.Removed,
	}).Info("TwoWayCrossBurn")

	if crossBurn.Raw.Removed {
		return
	}

	if !m.confirmEvent(crossBurn.Raw, monitor.CrossBurn) {
		m.logger.WithFields(logrus.Fields{
			"txId":         crossBurn.Raw.TxHash.String(),
			"block_height": crossBurn.Raw.BlockNumber,
		}).Info("TwoWayCrossBurn has not confirmed")
		return
	}

	m.logger.WithField("tx", crossBurn.Raw.TxHash.String()).Info("confirmEvent")
	m.cocoC <- coco
	m.persistCBlockHeight(crossBurn.Raw.TxHash.String(), crossBurn.Raw.BlockNumber, coco)
}

func (m *Monitor) confirmEvent(event types.Log, typ int) bool {
	for {
		num, err := m.fetchBlockNum()
		if err != nil {
			time.Sleep(15 * time.Second)
			continue
		}
		isConfirmed := num-event.BlockNumber >= m.minConfirms
		if !isConfirmed {
			time.Sleep(15 * time.Second)
			continue
		}
		var log *monitor.Coco
		switch typ {
		case monitor.Lock:
			log, err = m.GetLockLog(event.TxHash.String())
		case monitor.CrossBurn:
			log, err = m.GetCrossBurnLog(event.TxHash.String())
		case monitor.Rollback:
			log, err = m.GetRollback(event.TxHash.String())
		}
		if err != nil {
			m.logger.WithFields(logrus.Fields{
				"err":        err,
				"now_height": num,
			}).Error("confirmEvent")
			continue
		}
		return log.BlockHeight == event.BlockNumber
	}
}

func (m *Monitor) Unlock(txId string, token common.Address, from common.Address, recipient common.Address, chainID, amount *big.Int) error {
	unlocked := m.wrapper.TxUnlocked(txId)
	if unlocked {
		m.logger.Infof("find txUnlocked Chain %d txId:%s", chainID.Uint64(), txId)
		return nil
	}

	m.logger.WithFields(logrus.Fields{
		"tx_id":     txId,
		"token":     token.String(),
		"sender":    from.String(),
		"recipient": recipient.String(),
		"amount":    amount.String(),
	}).Info("will unlock")

	var (
		transaction *types.Transaction
		receipt     *types.Receipt
		err         error
		hashes      []common.Hash
	)

	m.wrapper.session.TransactOpts.Nonce = nil
	m.wrapper.session.TransactOpts.GasPrice = nil

	for {
		price := m.wrapper.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		if m.wrapper.session.TransactOpts.GasPrice == nil ||
			gasPrice.BigInt().Cmp(m.wrapper.session.TransactOpts.GasPrice) == 1 {
			m.wrapper.session.TransactOpts.GasPrice = gasPrice.BigInt()
			var hash common.Hash
			transaction, hash = m.wrapper.Unlock(token, from, recipient, chainID, amount, txId)
			m.wrapper.session.TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
			hashes = append(hashes, hash)

			m.logger.Infof("send UnlockBor tx %s with gasPrice %s and nonce %d",
				hash, gasPrice.String(), transaction.Nonce())
		}
		receipt, err = m.wrapper.TransactionReceiptsLimitedRetry(context.TODO(), hashes)
		if err == nil {
			break
		}
	}

	if receipt.Status == 1 {
		m.logger.WithField("tx_hash", receipt.TxHash.String()).Info("unlock success")
	} else {
		return fmt.Errorf("unlock fail:%s", receipt.TxHash.String())
	}
	return nil
}

func (m *Monitor) Rollback(txId string, token common.Address, from common.Address, recipient common.Address, chainID, amount *big.Int) error {
	unlocked := m.wrapper.TxUnlocked(txId)
	if unlocked {
		m.logger.Infof("find bridge.Rollback Chain %d txId:%s", chainID.Uint64(), txId)
		return nil
	}

	m.logger.WithFields(logrus.Fields{
		"tx_id":     txId,
		"token":     token.String(),
		"sender":    from.String(),
		"chainId":   chainID.String(),
		"recipient": recipient.String(),
		"amount":    amount.String(),
	}).Info("will rollback")

	var (
		transaction *types.Transaction
		receipt     *types.Receipt
		err         error
		hashes      []common.Hash
	)

	m.wrapper.session.TransactOpts.Nonce = nil
	m.wrapper.session.TransactOpts.GasPrice = nil

	for {
		price := m.wrapper.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		if m.wrapper.session.TransactOpts.GasPrice == nil ||
			gasPrice.BigInt().Cmp(m.wrapper.session.TransactOpts.GasPrice) == 1 {
			m.wrapper.session.TransactOpts.GasPrice = gasPrice.BigInt()
			var hash common.Hash
			transaction, hash = m.wrapper.Rollback(token, from, chainID, amount, txId)
			m.wrapper.session.TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
			hashes = append(hashes, hash)

			m.logger.Infof("send bridge.Rollback tx %s with gasPrice %s and nonce %d",
				hash, gasPrice.String(), transaction.Nonce())
		}
		receipt, err = m.wrapper.TransactionReceiptsLimitedRetry(context.TODO(), hashes)
		if err == nil {
			break
		}
	}

	if receipt.Status == 1 {
		m.logger.WithField("tx_hash", receipt.TxHash.String()).Info("rollback success")
	} else {
		return fmt.Errorf("rollback fail:%s", receipt.TxHash.String())
	}
	return nil
}

func (m *Monitor) CrossIn(txId string, token common.Address, from common.Address, recipient common.Address, chainID, amount *big.Int) error {
	unlocked := m.wrapper.TxMinted(txId)
	if unlocked {
		m.logger.Infof("find TxMinted txId:%s", txId)
		return nil
	}

	m.logger.WithFields(logrus.Fields{
		"tx_id":     txId,
		"token":     token.String(),
		"sender":    from.String(),
		"recipient": recipient.String(),
		"amount":    amount.String(),
	}).Info("will crossIn")

	var (
		transaction *types.Transaction
		receipt     *types.Receipt
		err         error
		hashes      []common.Hash
	)

	m.wrapper.session.TransactOpts.Nonce = nil
	m.wrapper.session.TransactOpts.GasPrice = nil

	for {
		price := m.wrapper.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		if m.wrapper.session.TransactOpts.GasPrice == nil ||
			gasPrice.BigInt().Cmp(m.wrapper.session.TransactOpts.GasPrice) == 1 {
			m.wrapper.session.TransactOpts.GasPrice = gasPrice.BigInt()

			var hash common.Hash
			transaction, hash = m.wrapper.CrossIn(token, from, recipient, chainID, amount, txId)
			m.wrapper.session.TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
			hashes = append(hashes, hash)

			m.logger.Infof("send CrossIn tx %s with gasPrice %s and nonce %d",
				hash, gasPrice.String(), transaction.Nonce())
		}
		receipt, err = m.wrapper.TransactionReceiptsLimitedRetry(context.TODO(), hashes)
		if err == nil {
			break
		}
	}

	if receipt.Status == 1 {
		m.logger.WithField("tx_hash", receipt.TxHash.String()).Info("CrossIn success")
	} else {
		return fmt.Errorf("unlock fail:%s", receipt.TxHash.String())
	}
	return nil
}

func (m *Monitor) GetLockLog(txId string) (*monitor.Coco, error) {
	receipt := m.wrapper.TransactionReceipt(context.TODO(), common.HexToHash(txId))
	for _, log := range receipt.Logs {
		if !strings.EqualFold(log.Address.String(), m.config.TwoWayContract) {
			continue
		}

		if log.Removed {
			continue
		}
		lock, err := m.wrapper.twoWay.ParseLock(*log)
		if err != nil {
			continue
		}
		return &monitor.Coco{
			Token0:      lock.Token0,
			Token1:      lock.Token1,
			ChainID0:    lock.ChainID0,
			ChainID1:    lock.ChainID1,
			From:        lock.From,
			To:          lock.To,
			Amount:      lock.Amount,
			TxId:        log.TxHash.String(),
			BlockHeight: receipt.BlockNumber.Uint64(),
		}, nil
	}
	return nil, fmt.Errorf("not found Lock log in tx:%s", txId)
}

func (m *Monitor) GetRollback(txId string) (*monitor.Coco, error) {
	receipt := m.wrapper.TransactionReceipt(context.TODO(), common.HexToHash(txId))
	for _, log := range receipt.Logs {
		if !strings.EqualFold(log.Address.String(), m.config.TwoWayContract) {
			continue
		}

		if log.Removed {
			continue
		}
		rollback, err := m.wrapper.twoWay.ParseRollback(*log)
		if err != nil {
			continue
		}
		return &monitor.Coco{
			Token0:      rollback.Token0,
			Token1:      rollback.Token1,
			ChainID0:    rollback.ChainID0,
			ChainID1:    rollback.ChainID1,
			From:        rollback.From,
			To:          rollback.To,
			Amount:      rollback.Amount,
			TxId:        log.TxHash.String(),
			BlockHeight: receipt.BlockNumber.Uint64(),
		}, nil
	}
	return nil, fmt.Errorf("not found rollback log in tx:%s", txId)
}

func (m *Monitor) GetCrossBurnLog(txId string) (*monitor.Coco, error) {
	receipt := m.wrapper.TransactionReceipt(context.TODO(), common.HexToHash(txId))
	for _, log := range receipt.Logs {
		if !strings.EqualFold(log.Address.String(), m.config.TwoWayContract) {
			continue
		}

		if log.Removed {
			continue
		}
		crossBurn, err := m.wrapper.twoWay.ParseCrossBurn(*log)
		if err != nil {
			continue
		}
		return &monitor.Coco{
			Token0:      crossBurn.Token0,
			Token1:      crossBurn.Token1,
			ChainID0:    crossBurn.ChainID0,
			ChainID1:    crossBurn.ChainID1,
			From:        crossBurn.From,
			To:          crossBurn.To,
			Amount:      crossBurn.Amount,
			TxId:        log.TxHash.String(),
			BlockHeight: receipt.BlockNumber.Uint64(),
		}, nil
	}
	return nil, fmt.Errorf("not found crossBurn log in tx:%s", txId)
}

func (m *Monitor) Name() string {
	return m.config.Name
}

func (m *Monitor) fetchBlockNum() (uint64, error) {
	header := m.wrapper.HeaderByNumber(context.TODO(), nil)
	return header.Number.Uint64(), nil
}

func (m *Monitor) loadHeightFromStorage() {
	var header *types.Header
	header = m.wrapper.HeaderByNumber(context.TODO(), nil)

	// load block height
	b := m.storage.Get(lHeightKey())
	if b == nil {
		m.lHeight = header.Number.Uint64() - m.minConfirms
		m.persistLHeight(m.lHeight)
	} else {
		m.lHeight = binary.LittleEndian.Uint64(b)

	}

	// load block height
	c := m.storage.Get(cHeightKey())
	if c == nil {
		m.cHeight = header.Number.Uint64() - m.minConfirms
		m.persistCHeight(m.cHeight)
	} else {
		m.cHeight = binary.LittleEndian.Uint64(c)

	}

	// load block height
	r := m.storage.Get(rHeightKey())
	if r == nil {
		m.rHeight = header.Number.Uint64() - m.minConfirms
		m.persistRHeight(m.rHeight)
	} else {
		m.rHeight = binary.LittleEndian.Uint64(r)

	}

	if m.config.LockHeight != 0 {
		m.lHeight = m.config.LockHeight
	}

	if m.config.CrossBurnHeight != 0 {
		m.cHeight = m.config.CrossBurnHeight
	}

	if m.config.RollbackHeight != 0 {
		m.rHeight = m.config.RollbackHeight
	}

	m.logger.WithFields(logrus.Fields{
		"lock_height":     m.lHeight,
		"cross_height":    m.cHeight,
		"rollback_height": m.rHeight,
	}).Info("Subscribe")
}

func lHeightKey() []byte {
	return []byte(fmt.Sprintf("lHeight"))
}

func cHeightKey() []byte {
	return []byte(fmt.Sprintf("cHeight"))
}

func rHeightKey() []byte {
	return []byte(fmt.Sprintf("rHeight"))
}

func (m *Monitor) persistLBlockHeight(txId string, height uint64, coco *monitor.Coco) {
	m.persistLHeight(height)
	for {
		if m.storage.Has(TxKey(txId, coco.Typ, coco.Index)) {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (m *Monitor) persistCBlockHeight(txId string, height uint64, coco *monitor.Coco) {
	m.persistCHeight(height)
	for {
		if m.storage.Has(TxKey(txId, coco.Typ, coco.Index)) {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (m *Monitor) persistRBlockHeight(txId string, height uint64, coco *monitor.Coco) {
	m.persistRHeight(height)
	for {
		if m.storage.Has(TxKey(txId, coco.Typ, coco.Index)) {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (m *Monitor) HasTx(txId string, coco *monitor.Coco) bool {
	return m.storage.Has(TxKey(txId, coco.Typ, coco.Index))
}

func (m *Monitor) persistLHeight(height uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, height)
	m.storage.Put(lHeightKey(), buf)
	m.lHeight = height
	m.logger.WithFields(logrus.Fields{
		"height": m.lHeight,
	}).Info("Persist Lock Block Height")
}

func (m *Monitor) persistRHeight(height uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, height)
	m.storage.Put(rHeightKey(), buf)
	m.rHeight = height
	m.logger.WithFields(logrus.Fields{
		"height": m.rHeight,
	}).Info("Persist Rollback Block Height")
}

func (m *Monitor) persistCHeight(height uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, height)
	m.storage.Put(cHeightKey(), buf)
	m.cHeight = height
	m.logger.WithFields(logrus.Fields{
		"height": m.cHeight,
	}).Info("Persist CrossBurn Block Height")
}

func TxKey(hash string, typ int, idx uint) []byte {
	return []byte(fmt.Sprintf("tx-%d-%s", typ, hash))
}

func (m *Monitor) PutTxID(txId string, coco *monitor.Coco) {
	data, err := json.Marshal(&coco)
	if err != nil {
		m.logger.Error(err)
	}
	m.storage.Put(TxKey(txId, coco.Typ, coco.Index), data)
}
