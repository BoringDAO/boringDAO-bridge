package chain

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/boringdao/bridge/internal/monitor/contracts/edge"

	"github.com/boringdao/bridge/internal/monitor"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/hexutil"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Monitor struct {
	depositedHeight     uint64
	crossOutedHeight    uint64
	crossInFailedHeight uint64
	wrapper             *Wrapper
	cocoC               chan *monitor.Coco
	logger              logrus.FieldLogger
	storage             storage.Storage
	config              *repo.EdgeConfig
	minConfirms         uint64
	edgeAddr            common.Address
	address             common.Address
	mut                 sync.Mutex
	ctx                 context.Context
	cancel              context.CancelFunc
}

func New(repoRoot string, config *repo.EdgeConfig, logger logrus.FieldLogger) (monitor.Mnt, error) {
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
		edgeAddr:    common.HexToAddress(config.EdgeContract),
		minConfirms: uint64(minConfirms),
		cocoC:       make(chan *monitor.Coco),
		logger:      logger,
		ctx:         ctx,
		cancel:      cancel,
	}, nil
}

func (m *Monitor) Start() error {
	m.loadHeightFromStorage()
	go m.listenDepositedEvent()
	go m.listenCrossOutedEvent()
	go m.listenCrossInFailedEvent()
	return nil
}

func (m *Monitor) Stop() error {
	m.cancel()
	m.wrapper.Close()
	return nil
}

func (m *Monitor) listenDepositedEvent() {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	start := m.depositedHeight

	for {
		select {
		case <-ticker.C:
			hasEvent := false
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
			filter := m.wrapper.FilterDeposited(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				m.handleDeposited(filter.Event)
				hasEvent = true
			}

			m.logger.WithFields(logrus.Fields{"start": start, "end": end, "current": num}).Infof("FilterDeposited")
			start = end + 1

			if !hasEvent {
				m.persistDepositedHeight(end)
			}
		case <-m.ctx.Done():
			m.logger.Info("CrossLockLockIterator done")
			return
		}
	}
}

func (m *Monitor) listenCrossOutedEvent() {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	start := m.crossOutedHeight

	for {
		select {
		case <-ticker.C:
			hasEvent := false
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
			filter := m.wrapper.FilterCrossOuted(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				m.handleCrossOuted(filter.Event)
				hasEvent = true
			}

			m.logger.WithFields(logrus.Fields{"start": start, "end": end, "current": num}).Infof("FilterCrossOuted")
			start = end + 1

			if !hasEvent {
				m.persistCrossOutedHeight(end)
			}
		case <-m.ctx.Done():
			m.logger.Info("CrossBurn done")
			return
		}
	}
}

func (m *Monitor) listenCrossInFailedEvent() {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	start := m.crossInFailedHeight

	for {
		select {
		case <-ticker.C:
			hasEvent := false
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
			filter := m.wrapper.FilterCrossInFailed(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				m.handleCrossInFailed(filter.Event)
				hasEvent = true
			}

			m.logger.WithFields(logrus.Fields{"start": start, "end": end, "current": num}).Infof("FilterCrossInFailed")
			start = end + 1

			if !hasEvent {
				m.persistCrossInFailedHeight(end)
			}
		case <-m.ctx.Done():
			m.logger.Info("RollbackIterator done")
			return
		}
	}
}

func (m *Monitor) HandleCocoC() chan *monitor.Coco {
	return m.cocoC
}

func (m *Monitor) handleCrossInFailed(rollback *edge.TwoWayEdgeCrossInFailed) {
	if !strings.EqualFold(rollback.Raw.Address.String(), m.config.EdgeContract) {
		return
	}

	if m.storage.Has(TxKey(rollback.Raw.TxHash.String(), monitor.CrossInFailed, rollback.Raw.Index)) {
		return
	}
	coco := &monitor.Coco{
		Typ:         monitor.CrossInFailed,
		From:        rollback.P.From,
		To:          rollback.P.To,
		FromToken:   rollback.P.FromToken,
		ToToken:     rollback.P.ToToken,
		FromChainId: rollback.P.FromChainId,
		ToChainId:   rollback.P.ToChainId,
		Amount:      rollback.P.Amount,
		Index:       rollback.Raw.Index,
		TxId:        rollback.Raw.TxHash.String(),
		BlockHeight: rollback.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"from":          coco.From.String(),
		"to":            coco.To.String(),
		"from_chain_id": coco.FromChainId.String(),
		"to_chain_id":   coco.ToChainId.String(),
		"from_token":    coco.FromToken.String(),
		"to_token":      coco.ToToken.String(),
		"amount":        coco.Amount.String(),
		"index":         coco.Index,
		"txId":          rollback.Raw.TxHash.String(),
		"block_height":  rollback.Raw.BlockNumber,
		"removed":       rollback.Raw.Removed,
	}).Info("CrossInFailed")

	if rollback.Raw.Removed {
		return
	}

	m.cocoC <- coco
	m.persistCrossInFailedBlockHeight(rollback.Raw.TxHash.String(), rollback.Raw.BlockNumber, coco)
}

func (m *Monitor) handleDeposited(lock *edge.TwoWayEdgeDeposited) {
	if !strings.EqualFold(lock.Raw.Address.String(), m.config.EdgeContract) {
		return
	}

	if m.storage.Has(TxKey(lock.Raw.TxHash.String(), monitor.Deposited, lock.Raw.Index)) {
		return
	}
	coco := &monitor.Coco{
		Typ:         monitor.Deposited,
		From:        lock.From,
		Amount:      lock.Amount,
		FromChainId: lock.FromChainId,
		FromToken:   lock.FromToken,
		Index:       lock.Raw.Index,
		TxId:        lock.Raw.TxHash.String(),
		BlockHeight: lock.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"from":          coco.From.String(),
		"to":            coco.To.String(),
		"from_chain_id": coco.FromChainId.String(),
		"to_chain_id":   coco.ToChainId.String(),
		"from_token":    coco.FromToken.String(),
		"to_token":      coco.ToToken.String(),
		"amount":        coco.Amount.String(),
		"index":         coco.Index,
		"txId":          lock.Raw.TxHash.String(),
		"block_height":  lock.Raw.BlockNumber,
		"removed":       lock.Raw.Removed,
	}).Info("Deposited")

	if lock.Raw.Removed {
		return
	}

	m.cocoC <- coco
	m.persistDepositedBlockHeight(lock.Raw.TxHash.String(), lock.Raw.BlockNumber, coco)
}

func (m *Monitor) handleCrossOuted(crossBurn *edge.TwoWayEdgeCrossOuted) {
	if !strings.EqualFold(crossBurn.Raw.Address.String(), m.config.EdgeContract) {
		return
	}

	if m.storage.Has(TxKey(crossBurn.Raw.TxHash.String(), monitor.CrossOuted, crossBurn.Raw.Index)) {
		return
	}
	coco := &monitor.Coco{
		Typ:         monitor.CrossOuted,
		From:        crossBurn.P.From,
		To:          crossBurn.P.To,
		FromToken:   crossBurn.P.FromToken,
		FromChainId: crossBurn.P.FromChainId,
		ToChainId:   crossBurn.P.ToChainId,
		Amount:      crossBurn.P.Amount,
		Index:       crossBurn.Raw.Index,
		TxId:        crossBurn.Raw.TxHash.String(),
		BlockHeight: crossBurn.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"from":          coco.From.String(),
		"to":            coco.To.String(),
		"from_chain_id": coco.FromChainId.String(),
		"to_chain_id":   coco.ToChainId.String(),
		"from_token":    coco.FromToken.String(),
		"to_token":      coco.ToToken.String(),
		"amount":        coco.Amount.String(),
		"index":         coco.Index,
		"txId":          crossBurn.Raw.TxHash.String(),
		"block_height":  crossBurn.Raw.BlockNumber,
		"removed":       crossBurn.Raw.Removed,
	}).Info("CrossOuted")

	if crossBurn.Raw.Removed {
		return
	}

	m.cocoC <- coco
	m.persistCrossOutedBlockHeight(crossBurn.Raw.TxHash.String(), crossBurn.Raw.BlockNumber, coco)
}

func (m *Monitor) CrossIn(fromToken, toToken common.Address, from, to common.Address, fromChainID, toChainID, amount *big.Int, txId string) error {
	unlocked := m.wrapper.TxHandled(txId)
	if unlocked {
		m.logger.Infof("find TxHandled txId:%s", txId)
		return nil
	}

	m.logger.WithFields(logrus.Fields{
		"tx_id":       txId,
		"from_token":  fromToken.String(),
		"to_token":    toToken.String(),
		"from":        from.String(),
		"to":          to.String(),
		"fromChainId": fromChainID.String(),
		"toChainId":   toChainID.String(),
		"amount":      amount.String(),
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
			transaction, hash = m.wrapper.CrossIn(fromToken, toToken, from, to, fromChainID, toChainID, amount, txId)
			if transaction != nil {
				m.wrapper.session.TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
				hashes = append(hashes, hash)

				m.logger.Infof("send CrossIn tx %s with gasPrice %s and nonce %d",
					hash, gasPrice.String(), transaction.Nonce())
			}
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
	b := m.storage.Get(depositedHeightKey())
	if b == nil {
		m.depositedHeight = header.Number.Uint64() - m.minConfirms
		m.persistDepositedHeight(m.depositedHeight)
	} else {
		m.depositedHeight = binary.LittleEndian.Uint64(b)

	}

	// load block height
	c := m.storage.Get(crossOutedHeightKey())
	if c == nil {
		m.crossOutedHeight = header.Number.Uint64() - m.minConfirms
		m.persistCrossOutedHeight(m.crossOutedHeight)
	} else {
		m.crossOutedHeight = binary.LittleEndian.Uint64(c)

	}

	// load block height
	r := m.storage.Get(crossInFailedHeightKey())
	if r == nil {
		m.crossInFailedHeight = header.Number.Uint64() - m.minConfirms
		m.persistCrossInFailedHeight(m.crossInFailedHeight)
	} else {
		m.crossInFailedHeight = binary.LittleEndian.Uint64(r)

	}

	if m.config.DepositedHeight != 0 {
		m.depositedHeight = m.config.DepositedHeight
	}

	if m.config.CrossOutedHeight != 0 {
		m.crossOutedHeight = m.config.CrossOutedHeight
	}

	if m.config.CrossInFailedHeight != 0 {
		m.crossInFailedHeight = m.config.CrossInFailedHeight
	}

	m.logger.WithFields(logrus.Fields{
		"depositedHeight":     m.depositedHeight,
		"crossOutedHeight":    m.crossOutedHeight,
		"crossInFailedHeight": m.crossInFailedHeight,
		"address":             m.address.String(),
	}).Info("Subscribe")
}

func depositedHeightKey() []byte {
	return []byte(fmt.Sprintf("depositedHeight"))
}

func crossOutedHeightKey() []byte {
	return []byte(fmt.Sprintf("crossOutedHeight"))
}

func crossInFailedHeightKey() []byte {
	return []byte(fmt.Sprintf("crossInFailedHeight"))
}

func (m *Monitor) persistDepositedBlockHeight(txId string, height uint64, coco *monitor.Coco) {
	m.persistDepositedHeight(height)
	for {
		if m.storage.Has(TxKey(txId, coco.Typ, coco.Index)) {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (m *Monitor) persistCrossOutedBlockHeight(txId string, height uint64, coco *monitor.Coco) {
	m.persistCrossOutedHeight(height)
	for {
		if m.storage.Has(TxKey(txId, coco.Typ, coco.Index)) {
			m.logger.WithFields(logrus.Fields{
				"height": m.crossOutedHeight,
			}).Info("Persist Cross Outed Height")
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (m *Monitor) persistCrossInFailedBlockHeight(txId string, height uint64, coco *monitor.Coco) {
	m.persistCrossInFailedHeight(height)
	for {
		if m.storage.Has(TxKey(txId, coco.Typ, coco.Index)) {
			m.logger.WithFields(logrus.Fields{
				"height": m.crossInFailedHeight,
			}).Info("Persist Cross In Failed Height")
			return
		}
		time.Sleep(500 * time.Millisecond)
	}

}

func (m *Monitor) HasTx(txId string, coco *monitor.Coco) bool {
	return m.storage.Has(TxKey(txId, coco.Typ, coco.Index))
}

func (m *Monitor) persistDepositedHeight(height uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, height)
	m.storage.Put(depositedHeightKey(), buf)
	m.depositedHeight = height
}

func (m *Monitor) persistCrossOutedHeight(height uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, height)
	m.storage.Put(crossOutedHeightKey(), buf)
	m.crossOutedHeight = height
}

func (m *Monitor) persistCrossInFailedHeight(height uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, height)
	m.storage.Put(crossInFailedHeightKey(), buf)
	m.crossInFailedHeight = height
}

func TxKey(hash string, typ int, idx uint) []byte {
	return []byte(fmt.Sprintf("tx-%d-%s-%d", typ, hash, idx))
}

func (m *Monitor) PutTxID(txId string, coco *monitor.Coco) {
	data, err := json.Marshal(&coco)
	if err != nil {
		m.logger.Error(err)
	}
	m.storage.Put(TxKey(txId, coco.Typ, coco.Index), data)
}

func (m *Monitor) MntLock() {
	m.mut.Lock()
}
func (m *Monitor) MntUnlock() {
	m.mut.Unlock()
}
