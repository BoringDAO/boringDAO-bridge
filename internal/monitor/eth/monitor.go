package eth

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/hexutil"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Coco struct {
	IsHistory   bool           `json:"isHistory"`
	Sender      common.Address `json:"sender"`
	Recipient   common.Address `json:"recipient"`
	Amount      *big.Int       `json:"amount"`
	EthToken    common.Address `json:"eth_token"`
	BscToken    common.Address `json:"bsc_token"`
	TxId        string         `json:"tx_id"`
	BlockHeight uint64         `json:"block_height"`
}

type Monitor struct {
	lHeight     uint64
	lockAbi     abi.ABI
	ethClient   *ethclient.Client
	topic       common.Hash
	cocoC       chan *Coco
	lock        *CrossLock
	session     *CrossLockSession
	logger      logrus.FieldLogger
	storage     storage.Storage
	config      *repo.Config
	minConfirms uint64
	mux         sync.Mutex
	lockAddr    common.Address
	address     common.Address

	BorAddr common.Address

	ctx    context.Context
	cancel context.CancelFunc
}

func New(config *repo.Config, logger logrus.FieldLogger) (*Monitor, error) {
	storagePath := repo.GetStoragePath(config.RepoRoot, "eth")
	ethStorage, err := leveldb.New(storagePath)
	if err != nil {
		return nil, err
	}

	etherCli, err := ethclient.Dial(config.Eth.Addr)
	if err != nil {
		return nil, fmt.Errorf("dial ethereum node: %w", err)
	}

	privKey, err := crypto.ToECDSA(hexutil.Decode(config.Eth.PrivKey))
	if err != nil {
		return nil, err
	}

	address := crypto.PubkeyToAddress(privKey.PublicKey)

	auth := bind.NewKeyedTransactor(privKey)
	if config.Eth.GasLimit < 800000 {
		auth.GasLimit = 1500000
	} else {
		auth.GasLimit = config.Eth.GasLimit
	}

	minConfirms := 0
	if config.Eth.MinConfirms > 0 {
		minConfirms = int(config.Eth.MinConfirms)
	}

	broker, err := NewCrossLock(common.HexToAddress(config.Eth.LockContract), etherCli)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate a lock contract: %w", err)
	}
	session := &CrossLockSession{
		Contract: broker,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
		TransactOpts: *auth,
	}

	lockBorAbi, err := abi.JSON(bytes.NewReader([]byte(CrossLockABI)))
	if err != nil {
		return nil, fmt.Errorf("abi unmarshal: %s", err.Error())
	}
	ctx, cancel := context.WithCancel(context.Background())

	return &Monitor{
		config:      config,
		storage:     ethStorage,
		session:     session,
		lock:        broker,
		ethClient:   etherCli,
		address:     address,
		lockAbi:     lockBorAbi,
		lockAddr:    common.HexToAddress(config.Eth.LockContract),
		BorAddr:     common.HexToAddress(config.Eth.BorContract),
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
	m.ethClient.Close()
	return nil
}

func (m *Monitor) listenLockEvent() {
	var filter *CrossLockLockIterator
	var err error
	err = retry.Retry(func(attempt uint) error {
		filter, err = m.lock.FilterLock(&bind.FilterOpts{Start: m.lHeight, End: nil, Context: m.ctx})
		if err != nil {
			m.logger.Errorf("FilterLock error:%w", err)
		}
		return err
	}, strategy.Wait(3*time.Second))
	if err != nil {
		m.logger.Error(err)
	}
	for filter.Next() {
		m.handleLock(filter.Event, true)
		if filter.done {
			break
		}
	}
	sink := make(chan *CrossLockLock, 0)
	event.Resubscribe(100*time.Millisecond, func(ctx context.Context) (event.Subscription, error) {
		sub, err := m.lock.WatchLock(&bind.WatchOpts{
			Start:   &m.lHeight,
			Context: ctx,
		}, sink)
		if err != nil {
			m.logger.Error(err)
			return nil, err
		}
		return sub, nil
	})
	for be := range sink {
		m.handleLock(be, false)
	}
}

func (m *Monitor) HandleCocoC() chan *Coco {
	return m.cocoC
}

func (m *Monitor) handleLock(lock *CrossLockLock, isHistory bool) {
	if !strings.EqualFold(lock.Raw.Address.String(), m.config.Eth.LockContract) {
		return
	}
	if m.storage.Has(TxKey(lock.Raw.TxHash.String())) {
		return
	}
	coco := &Coco{
		IsHistory:   isHistory,
		EthToken:    lock.EthToken,
		BscToken:    lock.BscToken,
		Sender:      lock.Locker,
		Recipient:   lock.Recipient,
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

func (m *Monitor) UnlockBor(txId string, token common.Address, from common.Address, recipient common.Address, amount *big.Int) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	price, err := m.ethClient.SuggestGasPrice(context.TODO())
	if err != nil {
		return err
	}
	gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
	m.session.TransactOpts.GasPrice = gasPrice.BigInt()

	m.logger.WithFields(logrus.Fields{
		"token":     token.String(),
		"recipient": recipient.String(),
		"amount":    amount.String(),
	}).Info("will unlock")

	transaction, err := m.session.Unlock(token, from, recipient, amount, txId)
	if err != nil {
		return fmt.Errorf("unlock error:%v", err)
	}
	var receipt *types.Receipt
	err = retry.Retry(func(attempt uint) error {
		receipt, err = m.ethClient.TransactionReceipt(context.TODO(), transaction.Hash())
		if err != nil {
			return err
		}
		if receipt == nil {
			return errors.Errorf("%s not found receipt", transaction.Hash().String())
		}
		return nil
	}, strategy.Wait(10*time.Second))
	if err != nil {
		m.logger.Error(err)
	}

	if receipt.Status == 1 {
		m.logger.WithField("tx_hash", transaction.Hash().String()).Info("unlock success")
	} else {
		m.logger.Errorf("unlock fail:%s", transaction.Hash().String())
	}
	return nil
}

func (m *Monitor) GetLockLog(txId string) (*Coco, error) {
	receipt, err := m.ethClient.TransactionReceipt(context.TODO(), common.HexToHash(txId))
	if err != nil {
		return nil, err
	}
	for _, log := range receipt.Logs {
		if !strings.EqualFold(log.Address.String(), m.config.Eth.LockContract) {
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
					EthToken:    lock.EthToken,
					BscToken:    lock.BscToken,
					Sender:      lock.Locker,
					Recipient:   lock.Recipient,
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
	header, err := m.ethClient.HeaderByNumber(context.TODO(), nil)
	if err != nil {
		m.logger.Error(err)
		return 0
	}
	return header.Number.Uint64()
}

func (m *Monitor) loadHeightFromStorage() {
	var header *types.Header
	var err error
	if err = retry.Retry(func(attempt uint) error {
		header, err = m.ethClient.HeaderByNumber(context.TODO(), nil)
		if err != nil {
			m.logger.Errorf("HeaderByNumber error:%v", err)
		}
		return err
	}, strategy.Wait(3*time.Second)); err != nil {
		m.logger.Error(err)
	}

	// load block height
	b := m.storage.Get(lHeightKey())
	if b == nil {
		m.lHeight = header.Number.Uint64() - m.minConfirms
		m.persistLHeight(m.lHeight)
	} else {
		m.lHeight = binary.LittleEndian.Uint64(b)
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
	m.lHeight = height
	m.logger.WithFields(logrus.Fields{
		"height": m.lHeight,
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
