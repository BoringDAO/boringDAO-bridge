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
	index       map[uint64]uint64
	wrapper     *Wrapper
	cocoC       chan *monitor.Coco
	logger      logrus.FieldLogger
	storage     storage.Storage
	config      *repo.CenterConfig
	chainIds    []uint64
	minConfirms uint64
	centerAddr  common.Address
	address     common.Address
	mut         sync.Mutex
	ctx         context.Context
	cancel      context.CancelFunc
	gasFeeRate  float64
}

func New(repoRoot string, config *repo.CenterConfig, ds []uint64, logger logrus.FieldLogger) (*Monitor, error) {
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

	var gasFeeRate float64
	if config.GasFeeRate < 1.5 {
		gasFeeRate = 1.5
	} else {
		gasFeeRate = config.GasFeeRate
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Monitor{
		index:       make(map[uint64]uint64),
		config:      config,
		storage:     ethStorage,
		gasFeeRate:  gasFeeRate,
		address:     address,
		wrapper:     wrapper,
		chainIds:    ds,
		centerAddr:  common.HexToAddress(config.CenterContract),
		minConfirms: uint64(minConfirms),
		cocoC:       make(chan *monitor.Coco),
		logger:      logger,
		ctx:         ctx,
		cancel:      cancel,
	}, nil
}

func (m *Monitor) Start() error {
	m.loadIndexFromStorage()
	for _, chainId := range m.chainIds {
		go m.listenEvent(chainId)
	}

	return nil
}

func (m *Monitor) Stop() error {
	m.cancel()
	m.wrapper.Close()
	return nil
}

func (m *Monitor) listenEvent(chainId uint64) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			num, err := m.fetchBlockNum()
			if err != nil {
				continue
			}
			end := num - m.config.MinConfirms
			hasEvent := false

			start, ok := m.index[chainId]
			if !ok {
				m.logger.Warnf("ignore native chain:[%d]", chainId)
				continue
			}
			chainBigInt := new(big.Int).SetUint64(chainId)
			index := m.wrapper.Index(chainBigInt)
			for i := start + 1; i <= index.Uint64(); i++ {
				indexHeight := m.wrapper.IndexHeight(chainBigInt, new(big.Int).SetUint64(i)).Uint64()
				if indexHeight > end || indexHeight == 0 {
					break
				}
				for !hasEvent {
					j := i
					withdrawedFilter := m.wrapper.FilterWithdrawed(&bind.FilterOpts{Start: indexHeight, End: &indexHeight, Context: m.ctx})
					for withdrawedFilter.Next() {
						if withdrawedFilter.Event.P.ToChainId.Uint64() != chainId {
							continue
						}
						m.handleWithdrawed(withdrawedFilter.Event)
						hasEvent = true
					}
					forwardFilter := m.wrapper.FilterForwardCrossOuted(&bind.FilterOpts{Start: indexHeight, End: &indexHeight, Context: m.ctx})
					for forwardFilter.Next() {
						if forwardFilter.Event.P.ToChainId.Uint64() != chainId {
							continue
						}
						m.handleForwardCrossOuted(forwardFilter.Event)
						hasEvent = true
					}
					crossOutedfilter := m.wrapper.FilterCenterCrossOuted(&bind.FilterOpts{Start: indexHeight, End: &indexHeight, Context: m.ctx})
					for crossOutedfilter.Next() {
						if crossOutedfilter.Event.P.ToChainId.Uint64() != chainId {
							continue
						}
						m.handleCenterCrossOuted(crossOutedfilter.Event)
						hasEvent = true

					}
					if hasEvent {
						m.persistIndex(chainId, j)
						m.index[chainId] = j
						j++
					}
					if j > i+1 {
						i = j
					}
				}

			}
			m.logger.Infof("listenEvent chainId:[%d], index from [%d] to [%d]", chainId, start, index.Uint64())
		case <-m.ctx.Done():
			m.logger.Info("FilterWithdrawed done")
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

	for {
		if m.storage.Has(TxKey(coco.TxId, coco.Typ, coco.Index)) {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (m *Monitor) handleCenterCrossOuted(outed *center.TwoWayCenterCrossOuted) {
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

	for {
		if m.storage.Has(TxKey(coco.TxId, coco.Typ, coco.Index)) {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (m *Monitor) handleForwardCrossOuted(outed *center.TwoWayCenterForwardCrossOuted) {
	if !strings.EqualFold(outed.Raw.Address.String(), m.config.CenterContract) {
		return
	}

	if m.storage.Has(TxKey(outed.Raw.TxHash.String(), monitor.ForwardCrossOuted, outed.Raw.Index)) {
		return
	}
	coco := &monitor.Coco{
		Typ:         monitor.ForwardCrossOuted,
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
	}).Info("ForwardCrossOuted")

	if outed.Raw.Removed {
		return
	}

	m.cocoC <- coco
	for {
		if m.storage.Has(TxKey(coco.TxId, coco.Typ, coco.Index)) {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
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
		if m.wrapper.session.TransactOpts.GasPrice != nil && price.Cmp(m.wrapper.session.TransactOpts.GasPrice) != 1 {
			price = decimal.NewFromBigInt(m.wrapper.session.TransactOpts.GasPrice, 0).Mul(decimal.NewFromFloat(1.1)).BigInt()
		}
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(m.gasFeeRate))
		if gasPrice.GreaterThan(decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(5))) {
			gasPrice = decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(5))
		}
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
		if unlocked := m.wrapper.TxHandled(txId); unlocked {
			m.logger.Infof("find TxHandled txId:%s", txId)
			return nil
		}
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
		if m.wrapper.session.TransactOpts.GasPrice != nil && price.Cmp(m.wrapper.session.TransactOpts.GasPrice) != 1 {
			price = decimal.NewFromBigInt(m.wrapper.session.TransactOpts.GasPrice, 0).Mul(decimal.NewFromFloat(1.1)).BigInt()
		}
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(m.gasFeeRate))
		if gasPrice.GreaterThan(decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(5))) {
			gasPrice = decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(5))
		}
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
		if unlocked := m.wrapper.TxHandled(txId); unlocked {
			m.logger.Infof("find TxHandled txId:%s", txId)
			return nil
		}
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
		if m.wrapper.session.TransactOpts.GasPrice != nil && price.Cmp(m.wrapper.session.TransactOpts.GasPrice) != 1 {
			price = decimal.NewFromBigInt(m.wrapper.session.TransactOpts.GasPrice, 0).Mul(decimal.NewFromFloat(1.1)).BigInt()
		}
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(m.gasFeeRate))
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
		if unlocked := m.wrapper.TxHandled(txId); unlocked {
			m.logger.Infof("find TxHandled txId:%s", txId)
			return nil
		}
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
		if m.wrapper.session.TransactOpts.GasPrice != nil && price.Cmp(m.wrapper.session.TransactOpts.GasPrice) != 1 {
			price = decimal.NewFromBigInt(m.wrapper.session.TransactOpts.GasPrice, 0).Mul(decimal.NewFromFloat(1.1)).BigInt()
		}
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(m.gasFeeRate))
		if gasPrice.GreaterThan(decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(5))) {
			gasPrice = decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(5))
		}
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
		if unlocked := m.wrapper.TxHandled(txId); unlocked {
			m.logger.Infof("find TxHandled txId:%s", txId)
			return nil
		}
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

func (m *Monitor) HasTx(txId string, coco *monitor.Coco) bool {
	return m.storage.Has(TxKey(txId, coco.Typ, coco.Index))
}

func (m *Monitor) persistIndex(chainId, index uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, index)
	m.storage.Put(indexKey(chainId), buf)
	m.logger.Infof("handled cross out events ChainId:[%d] Index:[%d]", chainId, index)
}

func TxKey(hash string, typ int, idx uint) []byte {
	return []byte(fmt.Sprintf("tx-%d-%s-%d", typ, hash, idx))
}

func indexKey(chainId uint64) []byte {
	return []byte(fmt.Sprintf("index-%d", chainId))
}

func (m *Monitor) loadIndexFromStorage() {
	for _, chainId := range m.chainIds {
		buf := m.storage.Get(indexKey(chainId))
		if buf != nil {
			m.index[chainId] = binary.LittleEndian.Uint64(buf)
		} else {
			m.index[chainId] = 0
		}
	}
	if m.config.Index != nil && len(m.config.Index) != 0 {
		for chainId, index := range m.config.Index {
			m.index[chainId] = index
		}
	}
	m.logger.WithFields(logrus.Fields{
		"index":   m.index,
		"chainID": m.config.ChainID,
	}).Info("Subscribe")
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
