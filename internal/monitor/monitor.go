package monitor

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	lru "github.com/hashicorp/golang-lru"
	"github.com/sirupsen/logrus"
)

var _ IMonitor = (*Monitor)(nil)

type Monitor struct {
	bHeight     uint64
	usdtWrapper *UsdtWrapper
	storage     storage.Storage
	config      *repo.StatisticsConfig
	cache       *lru.Cache

	logger logrus.FieldLogger
	ctx    context.Context
	cancel context.CancelFunc
}

type TransferEvent struct {
	Transfer  *UsdtTransfer
	Timestamp uint64
}

func New(repoRoot string, config *repo.StatisticsConfig, logger logrus.FieldLogger) (*Monitor, error) {
	storagePath := repo.GetStoragePath(repoRoot, fmt.Sprintf("statistics_%d", config.ChainID))
	storage, err := leveldb.New(storagePath)
	if err != nil {
		return nil, err
	}

	bw, err := NewUsdtWrapper(config, logger)
	if err != nil {
		return nil, err
	}

	cache, err := lru.New(1000)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Monitor{
		config:      config,
		storage:     storage,
		usdtWrapper: bw,
		cache:       cache,
		logger:      logger,
		ctx:         ctx,
		cancel:      cancel,
	}, nil
}

func (m *Monitor) Start() error {
	m.loadHeightFromStorage()

	go m.listenTransferEvent()

	return nil
}

func (m *Monitor) Stop() error {
	m.cancel()
	m.usdtWrapper.Close()
	return nil
}

func (m *Monitor) listenTransferEvent() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	var once sync.Once

	step := uint64(1000)
	start := m.bHeight
	swapPair := common.HexToAddress(m.config.SwapPairContract)

	for {
		select {
		case <-ticker.C:
			num := m.usdtWrapper.BlockNumber(context.TODO())
			end := num - m.config.MinConfirms
			if end < start {
				continue
			}

			if end-start > step {
				end = start + step
			} else {
				once.Do(func() {
					ticker.Reset(30 * time.Second)
				})
			}

			iteratorAdd, iteratorSub := m.usdtWrapper.FilterTransfer(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx}, swapPair)
			for iteratorAdd.Next() {
				m.handleAddLiquidityEvent(iteratorAdd.Event)
			}
			for iteratorSub.Next() {
				m.handleRemoveLiquidityEvent(iteratorSub.Event)
			}
			m.logger.WithFields(logrus.Fields{"start": start, "end": end}).Infof("FilterTransfer end")

			start = end + 1

			m.persistBHeight(end)

		case <-m.ctx.Done():
			m.logger.Info("listenAddLiquidityEvent done")
			return
		}
	}
}

func (m *Monitor) handleAddLiquidityEvent(transfer *UsdtTransfer) bool {
	if !strings.EqualFold(transfer.Raw.Address.String(), m.config.UsdtContract) {
		m.logger.Warnf("ignore log with unexpected contract address: %s", transfer.Raw.Address.String())
		return false
	}

	if m.hasAddLiquidityTx(transfer.Raw.TxHash.String(), transfer.From.String()) {
		m.logger.Warnf("add liquidity tx %s has been processed", transfer.Raw.TxHash.String())
		return false
	}

	if transfer.To.String() != m.config.SwapPairContract {
		m.logger.Warnf("get unexpected add liquidity event from %s to %s", transfer.From.String(), transfer.To.String())
		return false
	}

	transferEvent := &TransferEvent{
		Transfer:  transfer,
		Timestamp: m.getBlockTimeStamp(transfer.Raw.BlockNumber),
	}

	m.persistEvent(transferEvent, true)

	m.logger.Infof("user %s add liquidity %s at block %d timestamp %d",
		transfer.From.String(), transfer.Value.String(), transfer.Raw.BlockNumber, transferEvent.Timestamp)

	return true
}

func (m *Monitor) handleRemoveLiquidityEvent(transfer *UsdtTransfer) bool {
	if !strings.EqualFold(transfer.Raw.Address.String(), m.config.UsdtContract) {
		m.logger.Warnf("ignore log with unexpected contract address: %s", transfer.Raw.Address.String())
		return false
	}

	if m.hasRemoveLiquidityTx(transfer.Raw.TxHash.String(), transfer.To.String()) {
		m.logger.Warnf("remove liquidity tx %s has been processed", transfer.Raw.TxHash.String())
		return false
	}

	if transfer.From.String() != m.config.SwapPairContract {
		m.logger.Warnf("get unexpected remove liquidity event from %s to %s", transfer.From.String(), transfer.To.String())
		return false
	}

	transferEvent := &TransferEvent{
		Transfer:  transfer,
		Timestamp: m.getBlockTimeStamp(transfer.Raw.BlockNumber),
	}

	m.persistEvent(transferEvent, false)

	m.logger.Infof("user %s remove liquidity %s at block %d timestamp %d",
		transfer.To.String(), transfer.Value.String(), transfer.Raw.BlockNumber, transferEvent.Timestamp)

	return true
}

func (m *Monitor) loadHeightFromStorage() {
	height := m.usdtWrapper.BlockNumber(context.TODO())
	m.bHeight = height - m.config.MinConfirms

	b := m.storage.Get(bHeightKey(m.config.ChainID))
	if b != nil {
		height := binary.LittleEndian.Uint64(b)
		if height < m.bHeight {
			m.bHeight = height
		}
	}

	if m.config.Height != 0 {
		m.bHeight = m.config.Height
	}

	m.persistBHeight(m.bHeight)

	m.logger.WithFields(logrus.Fields{
		"height":  m.bHeight,
		"chainID": m.config.ChainID,
	}).Info("Subscribe")
}

func (m *Monitor) persistEvent(event *TransferEvent, add bool) {
	batch := m.storage.NewBatch()
	data, err := json.Marshal(event)
	if err != nil {
		m.logger.Panicf("marshal transfer event %v: %v", event, err)
	}
	if add {
		batch.Put(addLiquidityTxKey(event.Transfer.From.String(), event.Transfer.Raw.TxHash.String()), data)
	} else {
		batch.Put(removeLiquidityTxKey(event.Transfer.To.String(), event.Transfer.Raw.TxHash.String()), data)
	}
	batch.Commit()
}

func bHeightKey(chainID uint64) []byte {
	return []byte(fmt.Sprintf("bHeight-%d", chainID))
}

func (m *Monitor) hasAddLiquidityTx(txId, user string) bool {
	return m.storage.Has(addLiquidityTxKey(user, txId))
}

func (m *Monitor) hasRemoveLiquidityTx(txId, user string) bool {
	return m.storage.Has(removeLiquidityTxKey(user, txId))
}

func (m *Monitor) persistBHeight(height uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, height)
	m.storage.Put(bHeightKey(m.config.ChainID), buf)
	//m.logger.WithFields(logrus.Fields{
	//	"chainID": m.config.ChainID,
	//	"height":  height,
	//}).Info("Persist Bridge Block Height")
}

func (m *Monitor) getBlockTimeStamp(block uint64) uint64 {
	val, ok := m.cache.Get(block)
	if ok {
		return val.(uint64)
	}

	header := m.usdtWrapper.HeaderByNumber(context.TODO(), big.NewInt(int64(block)))
	m.cache.Add(block, header.Time)

	return header.Time
}

func addLiquidityTxKey(user, hash string) []byte {
	return []byte(fmt.Sprintf("tx-add-%s-%s", hash, user))
}

func removeLiquidityTxKey(user, hash string) []byte {
	return []byte(fmt.Sprintf("tx-remove-%s-%s", hash, user))
}

func (m *Monitor) addUserAmount(user string, amount *big.Int) *big.Int {
	balance := big.NewInt(0)
	val := m.storage.Get(userKey(user))
	if len(val) != 0 {
		balance.SetBytes(val)
	}

	balance = balance.Add(balance, amount)

	return balance
}

func (m *Monitor) subUserAmount(user string, amount *big.Int) *big.Int {
	balance := big.NewInt(0)
	val := m.storage.Get(userKey(user))
	if len(val) != 0 {
		balance = balance.SetBytes(val)
	}

	balance = balance.Sub(balance, amount)

	m.storage.Put(userKey(user), balance.Bytes())

	return balance
}

func userKey(addr string) []byte {
	return []byte(fmt.Sprintf("user-%s", addr))
}
