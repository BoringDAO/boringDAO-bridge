package bsc

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

	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/hexutil"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type Coco struct {
	IsHistory   bool           `json:"isHistory"`
	Coin        string         `json:"coin"`
	Sender      common.Address `json:"sender"`
	Recipient   common.Address `json:"recipient"`
	Amount      *big.Int       `json:"amount"`
	TxId        string         `json:"tx_id"`
	BlockHeight uint64         `json:"block_height"`
	EthToken    common.Address `json:"eth_token"`
	BscToken    common.Address `json:"bsc_token"`
}

type Monitor struct {
	bHeight     uint64
	borBscAbi   abi.ABI
	bscWrapper  *BscWrapper
	topic       common.Hash
	cocoC       chan *Coco
	logger      logrus.FieldLogger
	storage     storage.Storage
	config      *repo.Config
	minConfirms uint64
	mux         sync.Mutex
	borBscAddr  common.Address
	address     common.Address

	ctx    context.Context
	cancel context.CancelFunc
}

func New(config *repo.Config, logger logrus.FieldLogger) (*Monitor, error) {
	storagePath := repo.GetStoragePath(config.RepoRoot, "bsc")
	ethStorage, err := leveldb.New(storagePath)
	if err != nil {
		return nil, err
	}

	privKey, err := crypto.ToECDSA(hexutil.Decode(config.Bsc.PrivKey))
	if err != nil {
		return nil, err
	}

	address := crypto.PubkeyToAddress(privKey.PublicKey)

	auth := bind.NewKeyedTransactor(privKey)
	if config.Bsc.GasLimit < 800000 {
		auth.GasLimit = 1500000
	} else {
		auth.GasLimit = config.Bsc.GasLimit
	}

	minConfirms := 0
	if config.Bsc.MinConfirms > 0 {
		minConfirms = int(config.Bsc.MinConfirms)
	}

	borAbi, err := abi.JSON(bytes.NewReader([]byte(BorBSCABI)))
	if err != nil {
		return nil, fmt.Errorf("abi unmarshal: %s", err.Error())
	}

	bscWrapper, err := NewBscWrapper(&config.Bsc, logger)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Monitor{
		config:      config,
		storage:     ethStorage,
		bscWrapper:  bscWrapper,
		address:     address,
		borBscAbi:   borAbi,
		borBscAddr:  common.HexToAddress(config.Bsc.BorBscContract),
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
	m.bscWrapper.Close()
	return nil
}

func (m *Monitor) listenLockEvent() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	start := m.bHeight

	for {
		select {
		case <-ticker.C:
			num, err := m.fetchBlockNum()
			if err != nil {
				continue
			}

			end := num - m.minConfirms
			if num < m.minConfirms || end < start {
				continue
			}
			if end >= start+2000 {
				end = start + 2000
			}
			filter := m.bscWrapper.FilterCrossBurn(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				m.handleCross(filter.Event, true)
			}
			m.logger.WithFields(logrus.Fields{"start": start, "end": end, "current": num}).Infof("BridgeCrossBurnIterator")

			start = end + 1
		case <-m.ctx.Done():
			m.logger.Info("BridgeCrossBurnIterator done")
			return
		}
	}

}

func (m *Monitor) handleCross(lock *BorBSCCrossBurn, isHistory bool) {
	if !strings.EqualFold(lock.Raw.Address.String(), m.config.Bsc.BorBscContract) {
		return
	}

	if m.storage.Has(TxKey(lock.Raw.TxHash.String())) {
		return
	}

	coco := &Coco{
		IsHistory:   isHistory,
		Sender:      lock.From,
		Recipient:   lock.To,
		Amount:      lock.Amount,
		EthToken:    lock.EthToken,
		BscToken:    lock.BscToken,
		TxId:        lock.Raw.TxHash.String(),
		BlockHeight: lock.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"Sender":       coco.Sender.String(),
		"Recipient":    coco.Recipient.String(),
		"EthToken":     coco.EthToken.String(),
		"BscToken":     coco.BscToken.String(),
		"amount":       coco.Amount.String(),
		"txId":         lock.Raw.TxHash.String(),
		"block_height": lock.Raw.BlockNumber,
		"removed":      lock.Raw.Removed,
	}).Info("CrossBurn")

	if lock.Raw.Removed {
		return
	}

	if !m.confirmEvent(lock.Raw) {
		m.logger.WithFields(logrus.Fields{
			"txId":         lock.Raw.TxHash.String(),
			"block_height": lock.Raw.BlockNumber,
		}).Info("CrossBurn has not confirmed")
		return
	}

	m.logger.WithField("tx", lock.Raw.TxHash.String()).Info("confirmEvent")

	m.cocoC <- coco
	m.persistBBlockHeight(lock.Raw.TxHash.String(), lock.Raw.BlockNumber)
}

func (m *Monitor) confirmEvent(event types.Log) bool {
	for {
		num, err := m.fetchBlockNum()
		if err != nil {
			time.Sleep(15 * time.Second)
			continue
		}
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
			return false
		}
		return log.BlockHeight == event.BlockNumber
	}
}
func (m *Monitor) HandleCocoC() chan *Coco {
	return m.cocoC
}

func (m *Monitor) CrossMint(txId string, addrFromEth common.Address, recipient common.Address, amount *big.Int) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	unlocked := m.bscWrapper.TxMinted(txId)
	if unlocked {
		m.logger.Infof("find TxMinted eth txId:%s", txId)
		return nil
	}

	m.logger.WithFields(logrus.Fields{
		"recipient": recipient.String(),
		"amount":    amount.String(),
	}).Info("will crossMint")

	var (
		transaction *types.Transaction
		receipt     *types.Receipt
		err         error
		hash        common.Hash
		hashes      []common.Hash
	)

	m.bscWrapper.session.TransactOpts.Nonce = nil
	m.bscWrapper.session.TransactOpts.GasPrice = nil

	for {
		price := m.bscWrapper.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		if m.bscWrapper.session.TransactOpts.GasPrice == nil ||
			gasPrice.BigInt().Cmp(m.bscWrapper.session.TransactOpts.GasPrice) == 1 {
			m.bscWrapper.session.TransactOpts.GasPrice = gasPrice.BigInt()

			transaction, hash = m.bscWrapper.CrossMint(addrFromEth, recipient, amount, txId)
			m.bscWrapper.session.TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
			hashes = append(hashes, hash)

			m.logger.Infof("send CrossMint tx %s with gasPrice %s and nonce %d",
				hash.String(), gasPrice.String(), transaction.Nonce())
		}

		receipt, err = m.bscWrapper.TransactionReceiptsLimitedRetry(context.TODO(), hashes)
		if err == nil {
			break
		}
	}

	if receipt.Status == 1 {
		m.logger.WithField("tx_hash", hash.String()).Info("crossMint success")
	} else {
		return fmt.Errorf("crossMint fail:%s", hash.String())
	}

	return nil
}

func (m *Monitor) GetLockLog(txId string) (*Coco, error) {
	receipt := m.bscWrapper.TransactionReceipt(context.TODO(), common.HexToHash(txId))
	for _, log := range receipt.Logs {
		if !strings.EqualFold(log.Address.String(), m.config.Bsc.BorBscContract) {
			continue
		}

		if log.Removed {
			continue
		}
		for _, topic := range log.Topics {
			if strings.EqualFold(topic.String(), m.borBscAbi.Events["CrossBurn"].ID.String()) {
				var lock BorBSCCrossBurn
				if err := m.borBscAbi.UnpackIntoInterface(&lock, "CrossBurn", log.Data); err != nil {
					m.logger.Error(err)
					continue
				}
				return &Coco{
					Sender:      lock.From,
					Recipient:   lock.To,
					Amount:      lock.Amount,
					EthToken:    lock.EthToken,
					BscToken:    lock.BscToken,
					TxId:        log.TxHash.String(),
					BlockHeight: receipt.BlockNumber.Uint64(),
				}, nil
			}
		}
	}
	return nil, fmt.Errorf("not found BorBSCCrossBurn log in tx:%s", txId)
}

func (m *Monitor) fetchBlockNum() (uint64, error) {
	header := m.bscWrapper.HeaderByNumber(context.TODO(), nil)
	return header.Number.Uint64(), nil
}

func (m *Monitor) loadHeightFromStorage() {
	header := m.bscWrapper.HeaderByNumber(context.TODO(), nil)

	// load block height
	b := m.storage.Get(bHeightKey())
	if b == nil {
		m.bHeight = header.Number.Uint64() - m.minConfirms
		m.persistBHeight(m.bHeight)
	} else {
		m.bHeight = binary.LittleEndian.Uint64(b)
	}

	if m.config.Bsc.Height != 0 {
		m.bHeight = m.config.Bsc.Height
	}

	m.logger.WithFields(logrus.Fields{
		"bsc_height": m.bHeight,
	}).Info("Subscribe")
}

func bHeightKey() []byte {
	return []byte(fmt.Sprintf("bHeight"))
}

func (m *Monitor) persistBBlockHeight(txId string, height uint64) {
	m.persistBHeight(height)
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

func (m *Monitor) persistBHeight(height uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, height)
	m.storage.Put(bHeightKey(), buf)
	m.bHeight = height
	m.logger.WithFields(logrus.Fields{
		"height": m.bHeight,
	}).Info("Persist Bsc Block Height")
}

func TxKey(hash string) []byte {
	return []byte(fmt.Sprintf("bsc-tx-%s", hash))
}
func (m *Monitor) PutTxID(txId string, coco *Coco) {
	data, err := json.Marshal(&coco)
	if err != nil {
		m.logger.Error(err)
	}
	m.storage.Put(TxKey(txId), data)
}
