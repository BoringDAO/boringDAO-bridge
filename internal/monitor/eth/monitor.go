package eth

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/hexutil"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
)

type Coco struct {
	IsHistory   bool           `json:"isHistory"`
	Sender      common.Address `json:"sender"`
	Recipient   common.Address `json:"recipient"`
	Amount      *big.Int       `json:"amount"`
	EthToken    common.Address `json:"eth_token"`
	BscToken    common.Address `json:"bsc_token"`
	ChainID     *big.Int       `json:"chain_id"`
	TxId        string         `json:"tx_id"`
	BlockHeight uint64         `json:"block_height"`
}

type Monitor struct {
	lHeight     uint64
	lockAbi     abi.ABI
	topic       common.Hash
	cocoC       chan *Coco
	lockWrapper *LockWrapper
	logger      logrus.FieldLogger
	storage     storage.Storage
	config      *repo.Config
	minConfirms uint64
	mux         sync.Mutex
	lockAddr    common.Address
	address     common.Address
	BorAddr     common.Address

	ctx    context.Context
	cancel context.CancelFunc
}

func New(config *repo.Config, logger logrus.FieldLogger) (*Monitor, error) {
	storagePath := repo.GetStoragePath(config.RepoRoot, "eth")
	ethStorage, err := leveldb.New(storagePath)
	if err != nil {
		return nil, err
	}

	privKey, err := crypto.ToECDSA(hexutil.Decode(config.Eth.PrivKey))
	if err != nil {
		return nil, err
	}

	address := crypto.PubkeyToAddress(privKey.PublicKey)

	minConfirms := 0
	if config.Eth.MinConfirms > 0 {
		minConfirms = int(config.Eth.MinConfirms)
	}

	lw, err := NewLockWrapper(&config.Eth, logger)
	if err != nil {
		return nil, err
	}

	lockBorAbi, err := abi.JSON(bytes.NewReader([]byte(CrossLockABI)))
	if err != nil {
		return nil, fmt.Errorf("abi unmarshal: %s", err.Error())
	}
	ctx, cancel := context.WithCancel(context.Background())

	return &Monitor{
		config:      config,
		storage:     ethStorage,
		lockWrapper: lw,
		address:     address,
		lockAbi:     lockBorAbi,
		lockAddr:    common.HexToAddress(config.Eth.CrossLockContract),
		minConfirms: uint64(minConfirms),
		cocoC:       make(chan *Coco),
		logger:      logger,
		ctx:         ctx,
		cancel:      cancel,
	}, nil
}

func (m *Monitor) Start() error {
	m.loadHeightFromStorage()
	go m.listenLockEvent()
	return nil
}

func (m *Monitor) Stop() error {
	m.cancel()
	m.lockWrapper.Close()
	return nil
}

func (m *Monitor) listenLockEvent() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	start := m.lHeight

	for {
		select {
		case <-ticker.C:
			num := m.fetchBlockNum()
			end := num - m.minConfirms
			if num < m.minConfirms || end < start {
				continue
			}

			if end-start > 2000 {
				end = start + 2000
			}

			m.logger.WithFields(logrus.Fields{"start": start, "end": end}).Infof("CrossLockLockIterator start")

			filter := m.lockWrapper.FilterLock(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				m.handleLock(filter.Event, true)
			}

			m.logger.WithFields(logrus.Fields{"start": start, "end": end}).Infof("CrossLockLockIterator end")

			start = end + 1
		case <-m.ctx.Done():
			m.logger.Info("CrossLockLockIterator done")
			return
		}
	}
}

func (m *Monitor) HandleCocoC() chan *Coco {
	return m.cocoC
}

func (m *Monitor) handleLock(lock *CrossLockLock, isHistory bool) {
	if !strings.EqualFold(lock.Raw.Address.String(), m.config.Eth.CrossLockContract) {
		m.logger.Debugf("ignore eth log with contract address: %s", lock.Raw.Address.String())
		return
	}

	token1, ok := m.config.Token[strings.ToLower(lock.Token0.String())]
	if !ok || !strings.EqualFold(token1, lock.Token1.String()) {
		m.logger.Debugf("ignore eth log with token address: %s, %s", lock.Token0.String(), lock.Token1.String())
		return
	}

	if m.storage.Has(TxKey(lock.Raw.TxHash.String())) {
		return
	}
	coco := &Coco{
		IsHistory:   isHistory,
		EthToken:    lock.Token0,
		BscToken:    lock.Token1,
		ChainID:     lock.ChainID,
		Sender:      lock.Locker,
		Recipient:   lock.To,
		Amount:      lock.Amount,
		TxId:        lock.Raw.TxHash.String(),
		BlockHeight: lock.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"eth_token":    coco.EthToken.String(),
		"bsc_token":    coco.BscToken.String(),
		"sender":       coco.Sender.String(),
		"recipient":    coco.Recipient.String(),
		"amount":       coco.Amount.String(),
		"txId":         lock.Raw.TxHash.String(),
		"block_height": lock.Raw.BlockNumber,
		"removed":      lock.Raw.Removed,
	}).Info("LockBorLock")

	if lock.Raw.Removed {
		return
	}

	if !m.confirmEvent(lock.Raw) {
		m.logger.WithFields(logrus.Fields{
			"txId":         lock.Raw.TxHash.String(),
			"block_height": lock.Raw.BlockNumber,
		}).Info("LockBorLock has not confirmed")
		return
	}

	m.logger.WithField("tx", lock.Raw.TxHash.String()).Info("confirmEvent")
	m.cocoC <- coco
	m.persistLBlockHeight(lock.Raw.TxHash.String(), lock.Raw.BlockNumber)
}

func (m *Monitor) confirmEvent(event types.Log) bool {
	for {
		num := m.fetchBlockNum()
		isConfirmed := num-event.BlockNumber >= m.minConfirms
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
			continue
		}
		return log.BlockHeight == event.BlockNumber
	}
}

func (m *Monitor) UnlockBor(txId string, token common.Address, chainId *big.Int, from common.Address, recipient common.Address, amount *big.Int) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	unlocked := m.lockWrapper.TxUnlocked(txId)
	if unlocked {
		m.logger.Infof("find txUnlocked bsc txId:%s", txId)
		return nil
	}

	m.logger.WithFields(logrus.Fields{
		"tx_id":     txId,
		"token":     token.String(),
		"chainID":   chainId.String(),
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

	m.lockWrapper.session.TransactOpts.Nonce = nil
	m.lockWrapper.session.TransactOpts.GasPrice = nil

	for {
		price := m.lockWrapper.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		if m.lockWrapper.session.TransactOpts.GasPrice == nil ||
			gasPrice.BigInt().Cmp(m.lockWrapper.session.TransactOpts.GasPrice) == 1 {
			m.lockWrapper.session.TransactOpts.GasPrice = gasPrice.BigInt()

			transaction = m.lockWrapper.Unlock(token, chainId, from, recipient, amount, txId)
			m.lockWrapper.session.TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
			hashes = append(hashes, transaction.Hash())

			m.logger.Infof("send UnlockBor tx %s with gasPrice %s and nonce %d",
				transaction.Hash().String(), gasPrice.String(), transaction.Nonce())
		}
		receipt, err = m.lockWrapper.TransactionReceiptsLimitedRetry(context.TODO(), hashes)
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

func (m *Monitor) GetLockLog(txId string) (*Coco, error) {
	receipt := m.lockWrapper.TransactionReceipt(context.TODO(), common.HexToHash(txId))
	for _, log := range receipt.Logs {
		if !strings.EqualFold(log.Address.String(), m.config.Eth.CrossLockContract) {
			continue
		}

		if log.Removed {
			continue
		}
		for _, topic := range log.Topics {
			if strings.EqualFold(topic.String(), m.lockAbi.Events["Lock"].ID.String()) {
				var lock CrossLockLock
				if err := m.lockAbi.UnpackIntoInterface(&lock, "Lock", log.Data); err != nil {
					m.logger.Error(err)
					continue
				}
				return &Coco{
					EthToken:    lock.Token0,
					BscToken:    lock.Token1,
					ChainID:     lock.ChainID,
					Sender:      lock.Locker,
					Recipient:   lock.To,
					Amount:      lock.Amount,
					TxId:        log.TxHash.String(),
					BlockHeight: receipt.BlockNumber.Uint64(),
				}, nil
			}
		}
	}
	return nil, fmt.Errorf("not found Lock log in tx:%s", txId)
}

func (m *Monitor) fetchBlockNum() uint64 {
	header := m.lockWrapper.HeaderByNumber(context.TODO(), nil)
	return header.Number.Uint64()
}

func (m *Monitor) loadHeightFromStorage() {
	height := m.fetchBlockNum()
	// load block height
	b := m.storage.Get(lHeightKey())
	if b == nil {
		m.lHeight = height - m.minConfirms
		if m.config.Eth.Height != 0 && m.config.Eth.Height < m.lHeight {
			m.lHeight = m.config.Eth.Height
		}

		m.persistLHeight(m.lHeight)
	} else {
		m.lHeight = binary.LittleEndian.Uint64(b)
		if m.config.Eth.Height != 0 {
			m.lHeight = m.config.Eth.Height
		}
	}

	m.logger.WithFields(logrus.Fields{
		"Lock_height": m.lHeight,
	}).Info("Subscribe")
}

func lHeightKey() []byte {
	return []byte(fmt.Sprintf("lHeight"))
}

func (m *Monitor) persistLBlockHeight(txId string, height uint64) {
	m.persistLHeight(height)
	for {
		if m.storage.Has(TxKey(txId)) {
			return
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (m *Monitor) HasTx(txId string) bool {
	return m.storage.Has(TxKey(txId))
}

func (m *Monitor) persistLHeight(height uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, height)
	m.storage.Put(lHeightKey(), buf)
	m.logger.WithFields(logrus.Fields{
		"height": height,
	}).Info("Persist Lock Block Height")
}

func TxKey(hash string) []byte {
	return []byte(fmt.Sprintf("eth-tx-%s", hash))
}

func (m *Monitor) PutTxID(txId string, coco *Coco) {
	data, err := json.Marshal(&coco)
	if err != nil {
		m.logger.Error(err)
	}
	m.storage.Put(TxKey(txId), data)
}
