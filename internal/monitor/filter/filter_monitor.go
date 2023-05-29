package filter

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
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"go.uber.org/atomic"
)

var _ monitor.IMonitor = (*Monitor)(nil)

type Monitor struct {
	bHeight       uint64
	bridgeAbi     abi.ABI
	crossOutC     chan *monitor.Coco
	crossInC      chan *monitor.Coco
	finishedC     chan *monitor.Coco
	bridgeWrapper *monitor.BridgeWrapper
	storage       storage.Storage
	config        *repo.BridgeConfig
	token         map[string]string
	chainIDs      []uint64
	cocoNum       atomic.Int64

	logger logrus.FieldLogger
	ctx    context.Context
	cancel context.CancelFunc
}

func New(repoRoot string, config *repo.BridgeConfig, token map[string]string, chainIDs []uint64, logger logrus.FieldLogger) (monitor.IMonitor, error) {
	storagePath := repo.GetStoragePath(repoRoot, fmt.Sprintf("bridge_%d", config.ChainID))
	storage, err := leveldb.New(storagePath)
	if err != nil {
		return nil, err
	}

	bw, err := monitor.NewBridgeWrapper(config, logger)
	if err != nil {
		return nil, err
	}

	borAbi, err := abi.JSON(bytes.NewReader([]byte(monitor.NBridgeABI)))
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
		crossOutC:     make(chan *monitor.Coco, 128),
		crossInC:      make(chan *monitor.Coco, 128),
		finishedC:     make(chan *monitor.Coco, 128),
		chainIDs:      chainIDs,
		logger:        logger,
		ctx:           ctx,
		cancel:        cancel,
	}, nil
}

func (m *Monitor) Start() error {
	m.loadHeightFromStorage()

	go m.listenCrossOutEvent()
	go m.handleCrossInCocoC()

	return nil
}

func (m *Monitor) Stop() error {
	m.cancel()
	m.bridgeWrapper.Close()
	return nil
}

func (m *Monitor) ListenCrossOutC() chan *monitor.Coco {
	return m.crossOutC
}

func (m *Monitor) HandleCrossIn(coco *monitor.Coco) {
	m.crossInC <- coco
}

func (m *Monitor) ListenFinishedCocoC() chan *monitor.Coco {
	return m.finishedC
}

func (m *Monitor) HandleFinishedCoco(coco *monitor.Coco) {
	m.persistBHeight(coco.ToChainId.Uint64(), coco.BlockHeight)
	m.putTxID(coco.TxId, coco)
	m.cocoNum.Dec()
}

func (m *Monitor) listenCrossOutEvent() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	start := m.bHeight

	for {
		select {
		case <-ticker.C:
			num := m.bridgeWrapper.BlockNumber(context.TODO())
			end := num - m.config.MinConfirms
			if end < start {
				continue
			}

			if end-start > 300 {
				end = start + 300
			}

			filter := m.bridgeWrapper.FilterCrossOut(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				res := m.handleCross(filter.Event, true)
				if res {
					m.cocoNum.Inc()
				}
			}
			// filter again
			if m.cocoNum.Load() == 0 {
				filter := m.bridgeWrapper.FilterCrossOut(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
				for filter.Next() {
					res := m.handleCross(filter.Event, true)
					if res {
						m.cocoNum.Inc()
					}
				}
			}
			m.logger.Infof("finished filtering cross out event between block [%d, %d]", start, end)

			start = end + 1

			if m.cocoNum.Load() == 0 {
				for _, chainID := range m.chainIDs {
					if chainID == m.config.ChainID {
						continue
					}
					m.persistBHeight(chainID, end)
				}
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

			m.finishedC <- coco

		case <-m.ctx.Done():
			m.logger.Info("handleCrossInEvent done")
			return
		}
	}
}

func (m *Monitor) handleCross(crossOut *monitor.NBridgeCrossOut, isHistory bool) bool {
	if !strings.EqualFold(crossOut.Raw.Address.String(), m.config.BridgeContract) {
		m.logger.Warnf("ignore log with contract address: %s", crossOut.Raw.Address.String())
		return false
	}

	if !m.checkSupportedToken(crossOut.OriginToken.String(), crossOut.OriginChainId.String()) {
		m.logger.Warnf("ignore log with unsupported original token: %s", crossOut.OriginToken.String())
		return false
	}

	if m.HasTx(crossOut.Raw.TxHash.String()) {
		return false
	}

	coco := &monitor.Coco{
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
		"TxId":          crossOut.Raw.TxHash.String(),
		"block_height":  crossOut.Raw.BlockNumber,
		"removed":       crossOut.Raw.Removed,
	}).Info("CrossOut")

	if crossOut.Raw.Removed {
		return false
	}

	if !m.confirmEvent(crossOut.Raw) {
		m.logger.WithFields(logrus.Fields{
			"txId":         crossOut.Raw.TxHash.String(),
			"FromChainId":  coco.FromChainId.String(),
			"ToChainId":    coco.ToChainId.String(),
			"block_height": crossOut.Raw.BlockNumber,
		}).Info("CrossOut has not confirmed")
		return false
	}

	m.logger.WithField("tx", crossOut.Raw.TxHash.String()).Info("confirmEvent")

	m.crossOutC <- coco

	return true
}

func (m *Monitor) checkSupportedToken(token, chainID string) bool {
	for tokenAddr, originChainId := range m.token {
		if strings.EqualFold(token, tokenAddr) && strings.EqualFold(chainID, originChainId) {
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

	m.bridgeWrapper.Session().TransactOpts.Nonce = nil
	m.bridgeWrapper.Session().TransactOpts.GasPrice = nil

	for {
		price := m.bridgeWrapper.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		if m.bridgeWrapper.Session().TransactOpts.GasPrice == nil ||
			gasPrice.BigInt().Cmp(m.bridgeWrapper.Session().TransactOpts.GasPrice) == 1 {
			m.bridgeWrapper.Session().TransactOpts.GasPrice = gasPrice.BigInt()

			transaction, hash = m.bridgeWrapper.CrossIn(originToken, originChainId, fromChainId, toChainId, from, to, amount, txId)
			m.bridgeWrapper.Session().TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
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

func (m *Monitor) GetLockLog(txId string) (*monitor.Coco, error) {
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
				var lock monitor.NBridgeCrossOut
				if err := m.bridgeAbi.UnpackIntoInterface(&lock, "CrossOut", log.Data); err != nil {
					m.logger.Error(err)
					continue
				}
				return &monitor.Coco{
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

func (m *Monitor) loadHeightFromStorage() {
	height := m.bridgeWrapper.BlockNumber(context.TODO())
	m.bHeight = height - m.config.MinConfirms

	for _, chainID := range m.chainIDs {
		if chainID == m.config.ChainID {
			continue
		}
		b := m.storage.Get(bHeightKey(chainID))
		if b != nil {
			height := binary.LittleEndian.Uint64(b)
			if height < m.bHeight {
				m.bHeight = height
			}
		}
	}

	if m.config.Height != 0 {
		m.bHeight = m.config.Height
	}

	for _, chainID := range m.chainIDs {
		if chainID == m.config.ChainID {
			continue
		}
		m.persistBHeight(chainID, m.bHeight)
	}

	m.logger.WithFields(logrus.Fields{
		"height":  m.bHeight,
		"chainID": m.config.ChainID,
	}).Info("Subscribe")
}

func bHeightKey(chainID uint64) []byte {
	return []byte(fmt.Sprintf("bHeight-%d", chainID))
}

func (m *Monitor) HasTx(txId string) bool {
	return m.storage.Has(TxKey(txId, m.config.ChainID))
}

func (m *Monitor) persistBHeight(chainID uint64, height uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, height)
	m.storage.Put(bHeightKey(chainID), buf)
	m.logger.Infof("handled cross out events before block %d on chain %d", height, chainID)
}

func TxKey(hash string, chainID uint64) []byte {
	return []byte(fmt.Sprintf("bridge-tx-%d-%s", chainID, hash))
}
func (m *Monitor) putTxID(txId string, coco *monitor.Coco) {
	data, err := json.Marshal(&coco)
	if err != nil {
		m.logger.Error(err)
	}
	m.storage.Put(TxKey(txId, m.config.ChainID), data)
}