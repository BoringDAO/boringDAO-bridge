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

	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

var _ IMonitor = (*Monitor)(nil)

type Monitor struct {
	bHeight       uint64
	bridgeAbi     abi.ABI
	crossOutC     chan *Coco
	crossInC      chan *Coco
	finishedC     chan *Coco
	bridgeWrapper *BridgeWrapper
	storage       storage.Storage
	config        *repo.BridgeConfig
	nft           *repo.Nft
	chainIDs      []uint64

	logger logrus.FieldLogger
	ctx    context.Context
	cancel context.CancelFunc
}

func New(repoRoot string, config *repo.BridgeConfig, nft *repo.Nft, chainIDs []uint64, logger logrus.FieldLogger) (*Monitor, error) {
	storagePath := repo.GetStoragePath(repoRoot, fmt.Sprintf("bridge_%d", config.ChainID))
	storage, err := leveldb.New(storagePath)
	if err != nil {
		return nil, err
	}

	bw, err := NewBridgeWrapper(config, logger)
	if err != nil {
		return nil, err
	}

	borAbi, err := abi.JSON(bytes.NewReader([]byte(BridgeABI)))
	if err != nil {
		return nil, fmt.Errorf("abi unmarshal: %s", err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Monitor{
		config:        config,
		nft:           nft,
		storage:       storage,
		bridgeWrapper: bw,
		bridgeAbi:     borAbi,
		crossOutC:     make(chan *Coco, 128),
		crossInC:      make(chan *Coco, 128),
		finishedC:     make(chan *Coco, 128),
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
	m.persistBHeight(coco.TargetChainID.Uint64(), coco.BlockHeight)
	m.putTxID(coco.TxId, coco)
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
				m.handleCross(filter.Event, true)
			}
			m.logger.WithFields(logrus.Fields{"start": start, "end": end}).Infof("FilterCrossOut end")

			start = end + 1
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
			originTokenInfo := BridgetokenInfo{
				Token:   coco.OriginToken0,
				ChainId: coco.OriginChainId,
				TokenId: coco.OriginTokenId,
			}

			if err := m.CrossIn(originTokenInfo, coco.OriginTokenUri, coco.From, coco.To, coco.TxId); err != nil {
				m.logger.WithFields(logrus.Fields{
					"From":           coco.From.String(),
					"To":             coco.To.String(),
					"OriginChainId":  coco.OriginChainId.String(),
					"OriginToken0":   coco.OriginToken0.String(),
					"OriginTokenId":  coco.OriginTokenId.String(),
					"OriginTokenUri": coco.OriginTokenUri,
					"TxId":           coco.TxId,
					"error":          err.Error(),
				}).Panic("CrossIn failed")
			}

			m.finishedC <- coco

		case <-m.ctx.Done():
			m.logger.Info("handleCrossInEvent done")
			return
		}
	}
}

func (m *Monitor) handleCross(crossOut *BridgeCrossOut, isHistory bool) {
	if !strings.EqualFold(crossOut.Raw.Address.String(), m.config.BridgeContract) {
		m.logger.Warnf("ignore log with contract address: %s", crossOut.Raw.Address.String())
		return
	}

	if !m.checkSupportedToken(crossOut.OriginToken0.String()) {
		m.logger.Warnf("ignore log with unsupported original token: %s", crossOut.OriginToken0.String())
		return
	}

	if m.HasTx(crossOut.Raw.TxHash.String()) {
		return
	}

	coco := &Coco{
		IsHistory:      isHistory,
		OriginToken0:   crossOut.OriginToken0,
		OriginChainId:  crossOut.OriginChainId,
		OriginTokenId:  crossOut.OriginTokenId,
		OriginTokenUri: crossOut.OriginTokenUri,
		TargetChainID:  crossOut.TargetChainID,
		From:           crossOut.From,
		To:             crossOut.To,
		TxId:           crossOut.Raw.TxHash.String(),
		SourceChainID:  m.config.ChainID,
		BlockHeight:    crossOut.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"OriginToken0":   coco.OriginToken0.String(),
		"OriginChainId":  coco.OriginChainId.String(),
		"OriginTokenId":  coco.OriginTokenId.String(),
		"OriginTokenUri": coco.OriginTokenUri,
		"TargetChainID":  coco.TargetChainID.String(),
		"From":           coco.From.String(),
		"To":             coco.To.String(),
		"TxId":           crossOut.Raw.TxHash.String(),
		"SourceChainID":  coco.SourceChainID,
		"block_height":   crossOut.Raw.BlockNumber,
		"removed":        crossOut.Raw.Removed,
	}).Info("CrossOut")

	if crossOut.Raw.Removed {
		return
	}

	if !m.confirmEvent(crossOut.Raw) {
		m.logger.WithFields(logrus.Fields{
			"txId":          crossOut.Raw.TxHash.String(),
			"SourceChainID": coco.SourceChainID,
			"TargetChainID": coco.TargetChainID.String(),
			"block_height":  crossOut.Raw.BlockNumber,
		}).Info("CrossOut has not confirmed")
		return
	}

	m.logger.WithField("tx", crossOut.Raw.TxHash.String()).Info("confirmEvent")

	m.crossOutC <- coco
}

func (m *Monitor) checkSupportedToken(token string) bool {
	for _, tokenAddr := range m.nft.Tokens {
		if strings.EqualFold(token, tokenAddr) {
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

func (m *Monitor) CrossIn(originTokenInfo BridgetokenInfo, originTokenUri string, from common.Address, to common.Address, txId string) error {
	minted := m.bridgeWrapper.TxMinted(txId)
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

			transaction, hash = m.bridgeWrapper.CrossIn(originTokenInfo, originTokenUri, from, to, txId)
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
				var lock BridgeCrossOut
				if err := m.bridgeAbi.UnpackIntoInterface(&lock, "CrossOut", log.Data); err != nil {
					m.logger.Error(err)
					continue
				}
				return &Coco{
					OriginToken0:   lock.OriginToken0,
					OriginChainId:  lock.OriginChainId,
					OriginTokenId:  lock.OriginTokenId,
					OriginTokenUri: lock.OriginTokenUri,
					From:           lock.From,
					To:             lock.To,
					TargetChainID:  lock.TargetChainID,
					TxId:           log.TxHash.String(),
					BlockHeight:    receipt.BlockNumber.Uint64(),
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
	m.logger.WithFields(logrus.Fields{
		"chainID": chainID,
		"height":  height,
	}).Info("Persist Bridge Block Height")
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
