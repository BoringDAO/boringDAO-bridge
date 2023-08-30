package monitor

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"go.uber.org/atomic"
)

var _ IMonitor = (*Monitor)(nil)

type Monitor struct {
	index         map[uint64]uint64
	bridgeAbi     abi.ABI
	crossOutC     chan *Coco
	crossInC      chan *Coco
	finishedC     chan *Coco
	bridgeWrapper *BridgeWrapper
	storage       storage.Storage
	config        *repo.BridgeConfig
	token         map[string]string
	chainIDs      []uint64
	cocoNum       atomic.Int64

	logger logrus.FieldLogger
	ctx    context.Context
	cancel context.CancelFunc
}

func New(repoRoot string, config *repo.BridgeConfig, token map[string]string, chainIDs []uint64, logger logrus.FieldLogger) (IMonitor, error) {
	storagePath := repo.GetStoragePath(repoRoot, fmt.Sprintf("bridge_%d", config.ChainID))
	storage, err := leveldb.New(storagePath)
	if err != nil {
		return nil, err
	}

	bw, err := NewBridgeWrapper(config, logger)
	if err != nil {
		return nil, err
	}

	borAbi, err := abi.JSON(bytes.NewReader([]byte(NBridgeABI)))
	if err != nil {
		return nil, fmt.Errorf("abi unmarshal: %s", err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Monitor{
		config:        config,
		token:         token,
		storage:       storage,
		bridgeWrapper: bw,
		bridgeAbi:     borAbi,
		crossOutC:     make(chan *Coco, 1024),
		crossInC:      make(chan *Coco, 1024),
		finishedC:     make(chan *Coco, 1024),
		chainIDs:      chainIDs,
		index:         make(map[uint64]uint64),
		logger:        logger,
		ctx:           ctx,
		cancel:        cancel,
	}, nil
}

func (m *Monitor) Start() error {
	m.loadIndexFromStorage()

	go m.listenCrossOutEvent()
	go m.handleCrossInCocoC()

	return nil
}

func (m *Monitor) Stop() error {
	m.cancel()
	m.bridgeWrapper.Close()
	return nil
}

func (m *Monitor) ListenCrossOutC() chan *Coco {
	return m.crossOutC
}

func (m *Monitor) HandleCrossIn(coco *Coco) {
	m.crossInC <- coco
}

func (m *Monitor) ListenFinishedCocoC() chan *Coco {
	return m.finishedC
}

func (m *Monitor) HandleFinishedCoco(coco *Coco) {
	m.persistIndex(coco.ToChainId.Uint64(), coco.Index)
	m.putTxID(coco.TxId, coco)
	m.cocoNum.Dec()
}

func (m *Monitor) listenCrossOutEvent() {
	ticker := time.NewTicker(100 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			num := m.bridgeWrapper.BlockNumber(context.TODO())
			end := num - m.config.MinConfirms
			for _, chainId := range m.chainIDs {
				start, ok := m.index[chainId]
				if !ok {
					m.logger.Warnf("ignore native chain:[%d]", chainId)
					continue
				}
				chainBigInt := new(big.Int).SetUint64(chainId)
				index := m.bridgeWrapper.Index(chainBigInt)
				for i := start + 1; i <= index.Uint64(); {
					m.logger.Infof("listenEvent chainId:[%d], index:[%d]", chainId, i)
					indexHeight := m.bridgeWrapper.IndexHeight(chainBigInt, new(big.Int).SetUint64(i)).Uint64()
					if indexHeight > end || indexHeight == 0 {
						m.logger.Warnf("IndexHeight[%d][%d]: %s", chainId, index.Uint64(), "indexHeight > end || indexHeight == 0")
						num := m.bridgeWrapper.BlockNumber(context.TODO())
						end = num - m.config.MinConfirms
						time.Sleep(50 * time.Second)
						continue
					}
					hasEvent := false
					for !hasEvent {
						filter := m.bridgeWrapper.FilterCrossOut(&bind.FilterOpts{Start: indexHeight, End: &indexHeight, Context: m.ctx})
						for filter.Next() {
							if filter.Event.ToChainId.Uint64() != chainId {
								continue
							}
							hasEvent = true
							m.handleCross(filter.Event, i)
							m.index[chainId] = i
							i++
						}
					}
				}
				m.logger.Infof("listenEvent chainId:[%d], index from [%d] to [%d]", chainId, start, index.Uint64())
			}

		case <-m.ctx.Done():
			m.logger.Info("listenCrossOutEvent done")
			return
		}
	}
}

func (m *Monitor) handleCrossInCocoC() {
	for {
		select {
		case coco := <-m.crossInC:
			if err := m.CrossIn(coco.OriginToken, coco.OriginChainId, coco.FromChainId, coco.ToChainId, coco.From, coco.To, coco.Amount, coco.TxId); err != nil {
				m.logger.WithFields(logrus.Fields{
					"OriginToken":   coco.OriginToken.String(),
					"OriginChainId": coco.OriginChainId.String(),
					"FromChainId":   coco.FromChainId.String(),
					"ToChainId":     coco.ToChainId.String(),
					"From":          coco.From.String(),
					"To":            coco.To.String(),
					"Amount":        coco.Amount.String(),
					"TxId":          coco.TxId,
					"error":         err.Error(),
				}).Panic("CrossIn failed")
			}
			// time.Sleep(2 * time.Second)

			m.finishedC <- coco

		case <-m.ctx.Done():
			m.logger.Info("handleCrossInEvent done")
			return
		}
	}
}

func (m *Monitor) handleCross(crossOut *NBridgeCrossOut, index uint64) bool {
	if !strings.EqualFold(crossOut.Raw.Address.String(), m.config.BridgeContract) {
		m.logger.Warnf("ignore log with contract address: %s", crossOut.Raw.Address.String())
		return false
	}

	if !m.checkSupportedToken(crossOut.OriginToken.String(), crossOut.OriginChainId.String()) {
		m.logger.Warnf("ignore log with unsupported original token: %s", crossOut.OriginToken.String())
		return false
	}

	coco := &Coco{
		Index:         index,
		OriginToken:   crossOut.OriginToken,
		OriginChainId: crossOut.OriginChainId,
		FromChainId:   crossOut.FromChainId,
		ToChainId:     crossOut.ToChainId,
		From:          crossOut.From,
		To:            crossOut.To,
		Amount:        crossOut.Amount,
		TxId:          crossOut.Raw.TxHash.String(),
		BlockHeight:   crossOut.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"OriginToken":   coco.OriginToken.String(),
		"OriginChainId": coco.OriginChainId.String(),
		"FromChainId":   coco.FromChainId.String(),
		"ToChainId":     coco.ToChainId.String(),
		"From":          coco.From.String(),
		"To":            coco.To.String(),
		"Amount":        coco.Amount.String(),
		"Index":         coco.Index,
		"TxId":          crossOut.Raw.TxHash.String(),
		"block_height":  crossOut.Raw.BlockNumber,
		"removed":       crossOut.Raw.Removed,
		"hasTx":         m.HasTx(crossOut.Raw.TxHash.String()),
	}).Info("CrossOut")

	if m.HasTx(crossOut.Raw.TxHash.String()) {
		return false
	}

	if crossOut.Raw.Removed {
		return false
	}

	m.crossOutC <- coco

	return true
}

func (m *Monitor) checkSupportedToken(token, chainID string) bool {
	for tokenAddr, originChainId := range m.token {
		if strings.EqualFold(token, strings.TrimSpace(tokenAddr)) && strings.EqualFold(chainID, strings.TrimSpace(originChainId)) {
			return true
		}
	}

	return false
}

func (m *Monitor) confirmEvent(event types.Log) bool {
	for {
		num := m.bridgeWrapper.BlockNumber(context.TODO())
		isConfirmed := num-event.BlockNumber >= m.config.MinConfirms
		if !isConfirmed {
			time.Sleep(15 * time.Second)
			continue
		}
		log, err := m.GetLockLog(event.TxHash.String())
		if err != nil {
			m.logger.WithFields(logrus.Fields{
				"err":        err,
				"now_height": num,
			}).Error("confirmEvent")
			return false
		}
		return log.BlockHeight == event.BlockNumber
	}
}

func (m *Monitor) CrossIn(originToken common.Address, originChainId *big.Int, fromChainId *big.Int, toChainId *big.Int,
	from common.Address, to common.Address, amount *big.Int, txId string) error {
	minted := m.bridgeWrapper.TxHandled(txId)
	if minted {
		m.logger.Infof("find TxMinted txId:%s", txId)
		return nil
	}

	m.logger.WithFields(logrus.Fields{
		"from": from.String(),
		"to":   to.String(),
		"txId": txId,
	}).Info("will crossIn")

	var (
		transaction *types.Transaction
		receipt     *types.Receipt
		err         error
		hash        common.Hash
		hashes      []common.Hash
	)

	m.bridgeWrapper.session.TransactOpts.Nonce = nil
	m.bridgeWrapper.session.TransactOpts.GasPrice = nil

	for {
		price := m.bridgeWrapper.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		if m.bridgeWrapper.session.TransactOpts.GasPrice == nil ||
			gasPrice.BigInt().Cmp(m.bridgeWrapper.session.TransactOpts.GasPrice) == 1 {
			m.bridgeWrapper.session.TransactOpts.GasPrice = gasPrice.BigInt()

			transaction, hash = m.bridgeWrapper.CrossIn(originToken, originChainId, fromChainId, toChainId, from, to, amount, txId)
			m.bridgeWrapper.session.TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
			hashes = append(hashes, hash)

			m.logger.Infof("send CrossIn tx %s with gasPrice %s and nonce %d",
				hash.String(), gasPrice.String(), transaction.Nonce())
		}

		receipt, err = m.bridgeWrapper.TransactionReceiptsLimitedRetry(context.TODO(), hashes)
		if err == nil {
			break
		}
	}

	if receipt.Status == 1 {
		m.logger.WithField("tx_hash", hash.String()).Info("CrossIn success")
	} else {
		return fmt.Errorf("CrossIn fail:%s", hash.String())
	}

	return nil
}

func (m *Monitor) GetLockLog(txId string) (*Coco, error) {
	receipt := m.bridgeWrapper.TransactionReceipt(context.TODO(), common.HexToHash(txId))
	for _, log := range receipt.Logs {
		if !strings.EqualFold(log.Address.String(), m.config.BridgeContract) {
			continue
		}

		if log.Removed {
			continue
		}
		for _, topic := range log.Topics {
			if strings.EqualFold(topic.String(), m.bridgeAbi.Events["CrossOut"].ID.String()) {
				var lock NBridgeCrossOut
				if err := m.bridgeAbi.UnpackIntoInterface(&lock, "CrossOut", log.Data); err != nil {
					m.logger.Error(err)
					continue
				}
				return &Coco{
					OriginToken:   lock.OriginToken,
					OriginChainId: lock.OriginChainId,
					FromChainId:   lock.FromChainId,
					ToChainId:     lock.ToChainId,
					From:          lock.From,
					To:            lock.To,
					Amount:        lock.Amount,
					TxId:          log.TxHash.String(),
					BlockHeight:   receipt.BlockNumber.Uint64(),
				}, nil
			}
		}
	}
	return nil, fmt.Errorf("not found CrossOut log in tx:%s", txId)
}

func (m *Monitor) loadIndexFromStorage() {

	for _, chainId := range m.chainIDs {
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

func indexKey(chainId uint64) []byte {
	return []byte(fmt.Sprintf("index-%d", chainId))
}

func (m *Monitor) HasTx(txId string) bool {
	return m.storage.Has(TxKey(txId, m.config.ChainID))
}

func (m *Monitor) persistIndex(chainId, index uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, index)
	m.storage.Put(indexKey(chainId), buf)
	m.logger.Infof("handled cross out events ChainId:[%d] Index:[%d]", chainId, index)
}

func TxKey(hash string, chainID uint64) []byte {
	return []byte(fmt.Sprintf("bridge-tx-%d-%s", chainID, hash))
}
func (m *Monitor) putTxID(txId string, coco *Coco) {
	data, err := json.Marshal(&coco)
	if err != nil {
		m.logger.Error(err)
	}
	m.storage.Put(TxKey(txId, m.config.ChainID), data)
}
