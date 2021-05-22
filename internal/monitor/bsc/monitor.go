package bsc

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
	Coin        string         `json:"coin"`
	Sender      common.Address `json:"sender"`
	Recipient   common.Address `json:"recipient"`
	Amount      *big.Int       `json:"amount"`
	TxId        string         `json:"tx_id"`
	BlockHeight uint64         `json:"block_height"`
	EthToken    common.Address `json:"eth_token"`
	BscToken    common.Address `json:"bsc_token"`
	ChainID     *big.Int       `json:"chain_id"`
}

type Monitor struct {
	bHeight       uint64
	bridgeAbi     abi.ABI
	topic         common.Hash
	cocoC         chan *Coco
	bridgeWrapper *BridgeWrapper
	logger        logrus.FieldLogger
	storage       storage.Storage
	config        *repo.Config
	minConfirms   uint64
	mux           sync.Mutex
	bridgeAddr    common.Address
	address       common.Address

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

	minConfirms := 0
	if config.Bsc.MinConfirms > 0 {
		minConfirms = int(config.Bsc.MinConfirms)
	}

	bw, err := NewBridgeWrapper(&config.Bsc, logger)
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
		storage:       ethStorage,
		bridgeWrapper: bw,
		address:       address,
		bridgeAbi:     borAbi,
		bridgeAddr:    common.HexToAddress(config.Bsc.BridgeContract),
		minConfirms:   uint64(minConfirms),
		cocoC:         make(chan *Coco),
		logger:        logger,
		ctx:           ctx,
		cancel:        cancel,
	}, nil
}

func (m *Monitor) Start() error {
	m.loadHeightFromStorage()
	go m.listenLockEvent()
	return nil
}

func (m *Monitor) Stop() error {
	m.cancel()
	m.bridgeWrapper.Close()
	return nil
}

func (m *Monitor) listenLockEvent() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	start := m.bHeight

	for {
		select {
		case <-ticker.C:
			num := m.fetchBlockNum()
			end := num - m.minConfirms
			if end < start {
				continue
			}

			if end-start > 2000 {
				end = start + 2000
			}

			filter := m.bridgeWrapper.FilterCrossBurn(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				m.handleCross(filter.Event, true)
			}
			m.logger.WithFields(logrus.Fields{"start": start, "end": end}).Infof("BridgeCrossBurnIterator")

			start = end + 1
		case <-m.ctx.Done():
			m.logger.Info("BridgeCrossBurnIterator done")
			return
		}
	}

}

func (m *Monitor) handleCross(lock *BridgeCrossBurn, isHistory bool) {
	if !strings.EqualFold(lock.Raw.Address.String(), m.config.Bsc.BridgeContract) {
		m.logger.Debugf("ignore bsc log with contract address: %s", lock.Raw.Address.String())
		return
	}

	token1, ok := m.config.Token[strings.ToLower(lock.Token0.String())]
	if !ok || !strings.EqualFold(token1, lock.Token1.String()) {
		m.logger.Debugf("ignore bsc log with token address: %s, %s", lock.Token0.String(), lock.Token1.String())
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
		EthToken:    lock.Token0,
		BscToken:    lock.Token1,
		ChainID:     lock.ChainID,
		TxId:        lock.Raw.TxHash.String(),
		BlockHeight: lock.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"Sender":       coco.Sender.String(),
		"Recipient":    coco.Recipient.String(),
		"EthToken":     coco.EthToken.String(),
		"BscToken":     coco.BscToken.String(),
		"ChainID":      coco.ChainID.String(),
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
			return false
		}
		return log.BlockHeight == event.BlockNumber
	}
}
func (m *Monitor) HandleCocoC() chan *Coco {
	return m.cocoC
}

func (m *Monitor) CrossMint(ethToken common.Address, txId string, addrFromEth common.Address, recipient common.Address, amount *big.Int) error {
	m.mux.Lock()
	defer m.mux.Unlock()

	unlocked := m.bridgeWrapper.TxMinted(txId)
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

			transaction = m.bridgeWrapper.CrossMint(ethToken, addrFromEth, recipient, amount, txId)
			m.bridgeWrapper.session.TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
			hashes = append(hashes, transaction.Hash())

			m.logger.Infof("send CrossMint tx %s with gasPrice %s and nonce %d",
				transaction.Hash().String(), gasPrice.String(), transaction.Nonce())
		}

		receipt, err = m.bridgeWrapper.TransactionReceiptsLimitedRetry(context.TODO(), hashes)
		if err == nil {
			break
		}
	}

	if receipt.Status == 1 {
		m.logger.WithField("tx_hash", transaction.Hash().String()).Info("crossMint success")
	} else {
		return fmt.Errorf("crossMint fail:%s", transaction.Hash().String())
	}

	return nil
}

func (m *Monitor) GetLockLog(txId string) (*Coco, error) {
	receipt := m.bridgeWrapper.TransactionReceipt(context.TODO(), common.HexToHash(txId))
	for _, log := range receipt.Logs {
		if !strings.EqualFold(log.Address.String(), m.config.Bsc.BridgeContract) {
			continue
		}

		if log.Removed {
			continue
		}
		for _, topic := range log.Topics {
			if strings.EqualFold(topic.String(), m.bridgeAbi.Events["CrossBurn"].ID.String()) {
				var lock BridgeCrossBurn
				if err := m.bridgeAbi.UnpackIntoInterface(&lock, "CrossBurn", log.Data); err != nil {
					m.logger.Error(err)
					continue
				}
				return &Coco{
					Sender:      lock.From,
					Recipient:   lock.To,
					Amount:      lock.Amount,
					EthToken:    lock.Token0,
					BscToken:    lock.Token1,
					ChainID:     lock.ChainID,
					TxId:        log.TxHash.String(),
					BlockHeight: receipt.BlockNumber.Uint64(),
				}, nil
			}
		}
	}
	return nil, fmt.Errorf("not found BridgeCrossBurn log in tx:%s", txId)
}

func (m *Monitor) fetchBlockNum() uint64 {
	header := m.bridgeWrapper.HeaderByNumber(context.TODO(), nil)
	return header.Number.Uint64()
}

func (m *Monitor) loadHeightFromStorage() {
	height := m.fetchBlockNum()

	// load block height
	b := m.storage.Get(bHeightKey())
	if b == nil {
		m.bHeight = height - m.minConfirms
		if m.config.Bsc.Height != 0 && m.config.Bsc.Height < m.bHeight {
			m.bHeight = m.config.Bsc.Height
		}
		m.persistBHeight(m.bHeight)
	} else {
		m.bHeight = binary.LittleEndian.Uint64(b)
		if m.config.Bsc.Height != 0 {
			m.bHeight = m.config.Bsc.Height
		}
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
	m.logger.WithFields(logrus.Fields{
		"height": height,
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
