package filter

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/boringdao/bridge/internal/monitor/chain"

	"github.com/boringdao/bridge/internal/monitor"
	"github.com/boringdao/bridge/internal/monitor/contracts/edge"
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
	depositedHeight  uint64
	crossOutedHeight uint64
	wrapper          *chain.Wrapper
	cocoC            chan *monitor.Coco
	logger           logrus.FieldLogger
	storage          storage.Storage
	config           *repo.EdgeConfig
	minConfirms      uint64
	gasFeeRate       float64
	edgeAddr         common.Address
	address          common.Address
	mut              sync.Mutex
	rangeHeight      uint64
	ctx              context.Context
	cancel           context.CancelFunc
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

	wrapper, err := chain.NewWrapper(config, logger)
	if err != nil {
		return nil, err
	}
	var gasFeeRate float64
	if config.GasFeeRate < 1.2 {
		gasFeeRate = 1.2
	} else {
		gasFeeRate = config.GasFeeRate
	}
	rangeHeight := uint64(0)
	if config.RangeHeight == 0 {
		rangeHeight = 300
	} else {
		rangeHeight = config.RangeHeight
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Monitor{
		config:      config,
		storage:     ethStorage,
		address:     address,
		gasFeeRate:  gasFeeRate,
		wrapper:     wrapper,
		rangeHeight: rangeHeight,
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
			if end >= start+m.rangeHeight {
				end = start + m.rangeHeight
			}
			filter := m.wrapper.FilterDeposited(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				m.handleDeposited(filter.Event)
				hasEvent = true
			}

			time.Sleep(3 * time.Second)
			// check again
			if !hasEvent {
				filter := m.wrapper.FilterDeposited(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
				for filter.Next() {
					m.handleDeposited(filter.Event)
					hasEvent = true
				}
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

			time.Sleep(3 * time.Second)
			// check again
			if !hasEvent {
				filter := m.wrapper.FilterCrossOuted(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
				for filter.Next() {
					m.handleCrossOuted(filter.Event)
					hasEvent = true
				}
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

func (m *Monitor) HandleCocoC() chan *monitor.Coco {
	return m.cocoC
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

	m.wrapper.Session().TransactOpts.Nonce = nil
	m.wrapper.Session().TransactOpts.GasPrice = nil

	for {
		price := m.wrapper.SuggestGasPrice(context.TODO())
		if m.wrapper.Session().TransactOpts.GasPrice != nil && price.Cmp(m.wrapper.Session().TransactOpts.GasPrice) != 1 {
			price = decimal.NewFromBigInt(m.wrapper.Session().TransactOpts.GasPrice, 0).Add(decimal.NewFromFloat(1)).BigInt()
		}
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(m.gasFeeRate))
		if gasPrice.GreaterThan(decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(5))) {
			gasPrice = decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(5))
		}
		if m.wrapper.Session().TransactOpts.GasPrice == nil ||
			gasPrice.BigInt().Cmp(m.wrapper.Session().TransactOpts.GasPrice) == 1 {
			m.wrapper.Session().TransactOpts.GasPrice = gasPrice.BigInt()

			var hash common.Hash
			transaction, hash = m.wrapper.CrossIn(fromToken, toToken, from, to, fromChainID, toChainID, amount, txId)
			if transaction != nil {
				m.wrapper.Session().TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
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
		if unlocked := m.wrapper.TxHandled(txId); unlocked {
			m.logger.Infof("find TxHandled txId:%s", txId)
			return nil
		}
		return fmt.Errorf("CrossIn fail:%s", receipt.TxHash.String())
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

	if m.config.DepositedHeight != 0 {
		m.depositedHeight = m.config.DepositedHeight
	}

	if m.config.CrossOutedHeight != 0 {
		m.crossOutedHeight = m.config.CrossOutedHeight
	}

	m.logger.WithFields(logrus.Fields{
		"depositedHeight":  m.depositedHeight,
		"crossOutedHeight": m.crossOutedHeight,
		"address":          m.address.String(),
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
			m.logger.WithFields(logrus.Fields{
				"height": m.depositedHeight,
			}).Info("Persist Deposited Height")
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
