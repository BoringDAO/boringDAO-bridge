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

	"github.com/boringdao/bridge/pkg/bridge"

	repo2 "github.com/boringdao/bridge/pkg/repo"

	"github.com/boringdao/bridge/internal/monitor/contracts/edge"

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
	cocoC       chan *bridge.Coco
	logger      logrus.FieldLogger
	storage     storage.Storage
	config      *repo2.EdgeConfig
	minConfirms uint64
	gasFeeRate  float64
	edgeAddr    common.Address
	address     common.Address
	mut         sync.Mutex
	ctx         context.Context
	cancel      context.CancelFunc
	chainIds    []uint64
}

func New(repoRoot string, config *repo2.EdgeConfig, ds []uint64, logger logrus.FieldLogger) (bridge.Mnt, error) {
	storagePath := repo2.GetStoragePath(repoRoot, fmt.Sprintf("%s_%d", config.Name, config.ChainID))
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
	if config.GasFeeRate < 1.2 {
		gasFeeRate = 1.2
	} else {
		gasFeeRate = config.GasFeeRate
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Monitor{
		index:       make(map[uint64]uint64),
		config:      config,
		storage:     ethStorage,
		address:     address,
		gasFeeRate:  gasFeeRate,
		wrapper:     wrapper,
		chainIds:    ds,
		edgeAddr:    common.HexToAddress(config.EdgeContract),
		minConfirms: uint64(minConfirms),
		cocoC:       make(chan *bridge.Coco),
		logger:      logger,
		ctx:         ctx,
		cancel:      cancel,
	}, nil
}

func (m *Monitor) Start() error {
	m.loadIndexFromStorage()
	go m.listenEvent()
	return nil
}

func (m *Monitor) Stop() error {
	m.cancel()
	m.wrapper.Close()
	return nil
}

func (m *Monitor) listenEvent() {
	ticker := time.NewTicker(20 * time.Second)
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
			for _, chainId := range m.chainIds {
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
						filter := m.wrapper.FilterDeposited(&bind.FilterOpts{Start: indexHeight, End: &indexHeight, Context: m.ctx})
						for filter.Next() {
							if chainId != m.config.ChainID {
								continue
							}
							m.handleDeposited(filter.Event)
							hasEvent = true
						}
						if !hasEvent {
							filter := m.wrapper.FilterCrossOuted(&bind.FilterOpts{Start: indexHeight, End: &indexHeight, Context: m.ctx})
							for filter.Next() {
								if filter.Event.P.ToChainId.Uint64() != chainId {
									continue
								}
								m.handleCrossOuted(filter.Event)
								hasEvent = true
							}
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
			}
		case <-m.ctx.Done():
			m.logger.Info("CrossLockLockIterator done")
			return
		}
	}
}

func (m *Monitor) HandleCocoC() chan *bridge.Coco {
	return m.cocoC
}

func (m *Monitor) handleDeposited(lock *edge.TwoWayEdgeDeposited) {
	if !strings.EqualFold(lock.Raw.Address.String(), m.config.EdgeContract) {
		return
	}

	if m.storage.Has(TxKey(lock.Raw.TxHash.String(), bridge.Deposited, lock.Raw.Index)) {
		return
	}
	coco := &bridge.Coco{
		Typ:         bridge.Deposited,
		From:        lock.From.String(),
		Amount:      lock.Amount,
		FromChainId: lock.FromChainId,
		FromToken:   lock.FromToken.String(),
		Index:       lock.Raw.Index,
		TxId:        lock.Raw.TxHash.String(),
		BlockHeight: lock.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"from":          coco.From,
		"to":            coco.To,
		"from_chain_id": coco.FromChainId.String(),
		"to_chain_id":   coco.ToChainId.String(),
		"from_token":    coco.FromToken,
		"to_token":      coco.ToToken,
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
	for {
		if m.storage.Has(TxKey(coco.TxId, coco.Typ, coco.Index)) {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (m *Monitor) handleCrossOuted(crossBurn *edge.TwoWayEdgeCrossOuted) {
	if !strings.EqualFold(crossBurn.Raw.Address.String(), m.config.EdgeContract) {
		return
	}

	if m.storage.Has(TxKey(crossBurn.Raw.TxHash.String(), bridge.CrossOuted, crossBurn.Raw.Index)) {
		return
	}
	coco := &bridge.Coco{
		Typ:         bridge.CrossOuted,
		From:        crossBurn.P.From.String(),
		To:          crossBurn.P.To.String(),
		FromToken:   crossBurn.P.FromToken.String(),
		FromChainId: crossBurn.P.FromChainId,
		ToChainId:   crossBurn.P.ToChainId,
		Amount:      crossBurn.P.Amount,
		Index:       crossBurn.Raw.Index,
		TxId:        crossBurn.Raw.TxHash.String(),
		BlockHeight: crossBurn.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"from":          coco.From,
		"to":            coco.To,
		"from_chain_id": coco.FromChainId.String(),
		"to_chain_id":   coco.ToChainId.String(),
		"from_token":    coco.FromToken,
		"to_token":      coco.ToToken,
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
	for {
		if m.storage.Has(TxKey(coco.TxId, coco.Typ, coco.Index)) {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (m *Monitor) CrossIn(fromToken, toToken string, from, to string, fromChainID, toChainID, amount *big.Int, txId string) error {
	unlocked := m.wrapper.TxHandled(txId)
	if unlocked {
		m.logger.Infof("find TxHandled txId:%s", txId)
		return nil
	}

	m.logger.WithFields(logrus.Fields{
		"tx_id":       txId,
		"from_token":  fromToken,
		"to_token":    toToken,
		"from":        from,
		"to":          to,
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
			price = decimal.NewFromBigInt(m.wrapper.session.TransactOpts.GasPrice, 0).Add(decimal.NewFromFloat(1)).BigInt()
		}
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(m.gasFeeRate))
		if gasPrice.GreaterThan(decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(5))) {
			gasPrice = decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(5))
		}
		if m.wrapper.session.TransactOpts.GasPrice == nil ||
			gasPrice.BigInt().Cmp(m.wrapper.session.TransactOpts.GasPrice) == 1 {
			m.wrapper.session.TransactOpts.GasPrice = gasPrice.BigInt()

			var hash common.Hash
			transaction, hash = m.wrapper.CrossIn(common.HexToAddress(fromToken), common.HexToAddress(toToken), common.HexToAddress(from), common.HexToAddress(to), fromChainID, toChainID, amount, txId)
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

func (m *Monitor) HasTx(txId string, coco *bridge.Coco) bool {
	return m.storage.Has(TxKey(txId, coco.Typ, coco.Index))
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

func TxKey(hash string, typ int, idx uint) []byte {
	return []byte(fmt.Sprintf("tx-%d-%s-%d", typ, hash, idx))
}

func (m *Monitor) PutTxID(txId string, coco *bridge.Coco) {
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

func (m *Monitor) persistIndex(chainId, index uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, index)
	m.storage.Put(indexKey(chainId), buf)
	m.logger.Infof("handled cross out events ChainId:[%d] Index:[%d]", chainId, index)
}
