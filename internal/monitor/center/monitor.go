package center_chain

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/boringdao/bridge/internal/monitor"
	"github.com/boringdao/bridge/internal/monitor/contracts/center"
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
	crossOutedHeight uint64
	withdrawedHeight uint64
	wrapper          *Wrapper
	cocoC            chan *monitor.Coco
	logger           logrus.FieldLogger
	storage          storage.Storage
	config           *repo.CenterConfig
	minConfirms      uint64
	centerAddr       common.Address
	address          common.Address
	mut              sync.Mutex
	ctx              context.Context
	cancel           context.CancelFunc
}

func New(repoRoot string, config *repo.CenterConfig, logger logrus.FieldLogger) (*Monitor, error) {
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
		centerAddr:  common.HexToAddress(config.CenterContract),
		minConfirms: uint64(minConfirms),
		cocoC:       make(chan *monitor.Coco),
		logger:      logger,
		ctx:         ctx,
		cancel:      cancel,
	}, nil
}

func (m *Monitor) Start() error {
	m.loadHeightFromStorage()
	go m.listenWithdrawedEvent()
	go m.listenCenterCrossOutedEvent()
	return nil
}

func (m *Monitor) Stop() error {
	m.cancel()
	m.wrapper.Close()
	return nil
}

func (m *Monitor) listenWithdrawedEvent() {
	ticker := time.NewTicker(20 * time.Second)
	defer ticker.Stop()

	start := m.withdrawedHeight

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
			filter := m.wrapper.FilterWithdrawed(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				m.handleWithdrawed(filter.Event)
				hasEvent = true
			}

			m.logger.WithFields(logrus.Fields{"start": start, "end": end, "current": num}).Infof("FilterWithdrawed")
			start = end + 1

			if !hasEvent {
				m.persistWithdrawedHeight(end)
			}
		case <-m.ctx.Done():
			m.logger.Info("FilterWithdrawed done")
			return
		}
	}
}

func (m *Monitor) listenCenterCrossOutedEvent() {
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
			filter := m.wrapper.FilterCenterCrossOuted(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				m.handleCenterCrossOuted(filter.Event)
				hasEvent = true
			}

			m.logger.WithFields(logrus.Fields{"start": start, "end": end, "current": num}).Infof("FilterCenterCrossOuted")
			start = end + 1

			if !hasEvent {
				m.persistCrossOutedHeight(end)
			}
		case <-m.ctx.Done():
			m.logger.Info("FilterCenterCrossOuted done")
			return
		}
	}
}

func (m *Monitor) HandleCocoC() chan *monitor.Coco {
	return m.cocoC
}

func (m *Monitor) handleWithdrawed(withdrawed *center.TwoWayCenterWithdrawed) {
	if !strings.EqualFold(withdrawed.Raw.Address.String(), m.config.CenterContract) {
		return
	}

	if m.storage.Has(TxKey(withdrawed.Raw.TxHash.String(), monitor.Withdrawed, withdrawed.Raw.Index)) {
		return
	}
	coco := &monitor.Coco{
		Typ:         monitor.Withdrawed,
		From:        withdrawed.P.From,
		To:          withdrawed.P.To,
		FromToken:   withdrawed.P.FromToken,
		ToToken:     withdrawed.P.ToToken,
		FromChainId: withdrawed.P.FromChainId,
		ToChainId:   withdrawed.P.ToChainId,
		Amount:      withdrawed.P.Amount,
		Index:       withdrawed.Raw.Index,
		TxId:        withdrawed.Raw.TxHash.String(),
		BlockHeight: withdrawed.Raw.BlockNumber,
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
		"txId":          withdrawed.Raw.TxHash.String(),
		"block_height":  withdrawed.Raw.BlockNumber,
		"removed":       withdrawed.Raw.Removed,
	}).Info("Withdrawed")

	if withdrawed.Raw.Removed {
		return
	}

	m.cocoC <- coco
	m.persistWithdrawedBlockHeight(withdrawed.Raw.TxHash.String(), withdrawed.Raw.BlockNumber, coco)
}

func (m *Monitor) handleCenterCrossOuted(outed *center.TwoWayCenterCenterCrossOuted) {
	if !strings.EqualFold(outed.Raw.Address.String(), m.config.CenterContract) {
		return
	}

	if m.storage.Has(TxKey(outed.Raw.TxHash.String(), monitor.CrossOuted, outed.Raw.Index)) {
		return
	}
	coco := &monitor.Coco{
		Typ:         monitor.CrossOuted,
		From:        outed.P.From,
		To:          outed.P.To,
		FromToken:   outed.P.FromToken,
		ToToken:     outed.P.ToToken,
		FromChainId: outed.P.FromChainId,
		ToChainId:   outed.P.ToChainId,
		Amount:      outed.P.Amount,
		Index:       outed.Raw.Index,
		TxId:        outed.Raw.TxHash.String(),
		BlockHeight: outed.Raw.BlockNumber,
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
		"txId":          outed.Raw.TxHash.String(),
		"block_height":  outed.Raw.BlockNumber,
		"removed":       outed.Raw.Removed,
	}).Info("CrossOuted")

	if outed.Raw.Removed {
		return
	}

	m.cocoC <- coco
	m.persistCrossOutedBlockHeight(outed.Raw.TxHash.String(), outed.Raw.BlockNumber, coco)
}

func (m *Monitor) CrossIn(fromToken common.Address, from, to common.Address, fromChainID, toChainID, amount *big.Int, txId string) error {
	unlocked := m.wrapper.TxHandled(txId)
	if unlocked {
		m.logger.Infof("find TxHandled txId:%s", txId)
		return nil
	}

	m.logger.WithFields(logrus.Fields{
		"tx_id":       txId,
		"from_token":  fromToken.String(),
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
			transaction, hash = m.wrapper.CrossIn(fromToken, from, to, fromChainID, toChainID, amount, txId)
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
		return fmt.Errorf("crossin fail:%s", receipt.TxHash.String())
	}
	return nil
}

func (m *Monitor) Issue(fromToken, toToken common.Address, from, to common.Address, fromChainID, toChainID, amount *big.Int, txId string) error {
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
	}).Info("will issue")

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
			transaction, hash = m.wrapper.Issue(fromToken, toToken, from, to, fromChainID, toChainID, amount, txId)
			if transaction != nil {
				m.wrapper.session.TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
				hashes = append(hashes, hash)

				m.logger.Infof("send Issue tx %s with gasPrice %s and nonce %d",
					hash, gasPrice.String(), transaction.Nonce())
			}
		}
		receipt, err = m.wrapper.TransactionReceiptsLimitedRetry(context.TODO(), hashes)
		if err == nil {
			break
		}
	}

	if receipt.Status == 1 {
		m.logger.WithField("tx_hash", receipt.TxHash.String()).Info("Issue success")
	} else {
		return fmt.Errorf("issue fail:%s", receipt.TxHash.String())
	}
	return nil
}

func (m *Monitor) RollbackCrossIn(fromToken, toToken common.Address, from, to common.Address, fromChainID, toChainID, amount *big.Int, txId string) error {
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
	}).Info("will rollback crossIn")

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
			transaction, hash = m.wrapper.RollbackCrossIn(fromToken, toToken, from, to, fromChainID, toChainID, amount, txId)
			if transaction != nil {
				m.wrapper.session.TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
				hashes = append(hashes, hash)

				m.logger.Infof("send RollbackCrossIn tx %s with gasPrice %s and nonce %d",
					hash, gasPrice.String(), transaction.Nonce())
			}
		}
		receipt, err = m.wrapper.TransactionReceiptsLimitedRetry(context.TODO(), hashes)
		if err == nil {
			break
		}
	}

	if receipt.Status == 1 {
		m.logger.WithField("tx_hash", receipt.TxHash.String()).Info("RollbackCrossIn success")
	} else {
		return fmt.Errorf("RollbackCrossIn fail:%s", receipt.TxHash.String())
	}
	return nil
}

func (m *Monitor) ForwardCrossOut(fromToken common.Address, from, to common.Address, fromChainID, toChainID, amount *big.Int, txId string) error {
	unlocked := m.wrapper.TxHandled(txId)
	if unlocked {
		m.logger.Infof("find TxHandled txId:%s", txId)
		return nil
	}

	m.logger.WithFields(logrus.Fields{
		"tx_id":       txId,
		"from_token":  fromToken.String(),
		"from":        from.String(),
		"to":          to.String(),
		"fromChainId": fromChainID.String(),
		"toChainId":   toChainID.String(),
		"amount":      amount.String(),
	}).Info("will forwardCrossOut")

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
			transaction, hash = m.wrapper.ForwardCrossOut(fromToken, from, to, fromChainID, toChainID, amount, txId)
			if transaction != nil {
				m.wrapper.session.TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
				hashes = append(hashes, hash)

				m.logger.Infof("send forwardCrossOut tx %s with gasPrice %s and nonce %d",
					hash, gasPrice.String(), transaction.Nonce())
			}
		}
		receipt, err = m.wrapper.TransactionReceiptsLimitedRetry(context.TODO(), hashes)
		if err == nil {
			break
		}
	}

	if receipt.Status == 1 {
		m.logger.WithField("tx_hash", receipt.TxHash.String()).Info("ForwardCrossOut success")
	} else {
		return fmt.Errorf("forwardCrossOut fail:%s", receipt.TxHash.String())
	}
	return nil
}

func (m *Monitor) Name() string {
	return m.config.Name
}

func (m *Monitor) ChainId() uint64 {
	return m.config.ChainID
}

func (m *Monitor) fetchBlockNum() (uint64, error) {
	header := m.wrapper.HeaderByNumber(context.TODO(), nil)
	return header.Number.Uint64(), nil
}

func (m *Monitor) loadHeightFromStorage() {
	var header *types.Header
	header = m.wrapper.HeaderByNumber(context.TODO(), nil)

	// load block height
	c := m.storage.Get(crossOutedHeightKey())
	if c == nil {
		m.crossOutedHeight = header.Number.Uint64() - m.minConfirms
		m.persistCrossOutedHeight(m.crossOutedHeight)
	} else {
		m.crossOutedHeight = binary.LittleEndian.Uint64(c)

	}

	// load block height
	r := m.storage.Get(WithdrawedHeightKey())
	if r == nil {
		m.withdrawedHeight = header.Number.Uint64() - m.minConfirms
		m.persistWithdrawedHeight(m.withdrawedHeight)
	} else {
		m.withdrawedHeight = binary.LittleEndian.Uint64(r)

	}

	if m.config.CrossOutedHeight != 0 {
		m.crossOutedHeight = m.config.CrossOutedHeight
	}

	if m.config.WithdrawedHeight != 0 {
		m.withdrawedHeight = m.config.WithdrawedHeight
	}

	m.logger.WithFields(logrus.Fields{
		"crossOutedHeight": m.crossOutedHeight,
		"WithdrawedHeight": m.withdrawedHeight,
		"address":          m.address.String(),
	}).Info("Subscribe")
}

func crossOutedHeightKey() []byte {
	return []byte(fmt.Sprintf("crossOutedHeight"))
}

func WithdrawedHeightKey() []byte {
	return []byte(fmt.Sprintf("withdrawedHeight"))
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

func (m *Monitor) persistWithdrawedBlockHeight(txId string, height uint64, coco *monitor.Coco) {
	m.persistWithdrawedHeight(height)
	for {
		if m.storage.Has(TxKey(txId, coco.Typ, coco.Index)) {
			m.logger.WithFields(logrus.Fields{
				"height": m.withdrawedHeight,
			}).Info("Persist Withdrawed Height")
			return
		}
		time.Sleep(500 * time.Millisecond)
	}

}

func (m *Monitor) HasTx(txId string, coco *monitor.Coco) bool {
	return m.storage.Has(TxKey(txId, coco.Typ, coco.Index))
}

func (m *Monitor) persistCrossOutedHeight(height uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, height)
	m.storage.Put(crossOutedHeightKey(), buf)
	m.crossOutedHeight = height
}

func (m *Monitor) persistWithdrawedHeight(height uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, height)
	m.storage.Put(WithdrawedHeightKey(), buf)
	m.withdrawedHeight = height
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
