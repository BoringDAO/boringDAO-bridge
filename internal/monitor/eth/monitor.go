package eth

import (
	"bytes"
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	mnt "github.com/boringdao/bridge/internal/monitor"

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

const Lock = 0
const CrossBurn = 1

type Coco struct {
	Typ         int            `json:"typ"`
	IsHistory   bool           `json:"isHistory"`
	Sender      common.Address `json:"sender"`
	Recipient   common.Address `json:"recipient"`
	Amount      *big.Int       `json:"amount"`
	SrcToken    common.Address `json:"src_token"`
	DstToken    common.Address `json:"dst_token"`
	TxId        string         `json:"tx_id"`
	BlockHeight uint64         `json:"block_height"`
}

type Monitor struct {
	lHeight      uint64
	cHeight      uint64
	pegProxyAbi  abi.ABI
	ethWrapper   *EthWrapper
	cocoC        chan *Coco
	logger       logrus.FieldLogger
	storage      storage.Storage
	config       *repo.Config
	minConfirms  uint64
	pegProxyAddr common.Address
	address      common.Address

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

	pegProxyAbi, err := abi.JSON(bytes.NewReader([]byte(mnt.PegProxyABI)))
	if err != nil {
		return nil, fmt.Errorf("abi unmarshal: %s", err.Error())
	}

	ethWrapper, err := NewEthWrapper(&config.Eth, logger)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Monitor{
		config:       config,
		storage:      ethStorage,
		address:      address,
		ethWrapper:   ethWrapper,
		pegProxyAbi:  pegProxyAbi,
		pegProxyAddr: common.HexToAddress(config.Eth.PegBridgeContract),
		minConfirms:  uint64(minConfirms),
		cocoC:        make(chan *Coco),
		logger:       logger,
		ctx:          ctx,
		cancel:       cancel,
	}, nil
}

func (m *Monitor) Start() error {
	m.loadHeightFromStorage()
	go m.listenLockEvent()
	go m.listenCrossBurnEvent()
	return nil
}

func (m *Monitor) Stop() error {
	m.cancel()
	m.ethWrapper.Close()
	return nil
}

func (m *Monitor) listenLockEvent() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	start := m.lHeight

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
			filter := m.ethWrapper.FilterLock(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				m.handleLock(filter.Event, true)
			}

			m.logger.WithFields(logrus.Fields{"start": start, "end": end, "current": num}).Infof("CrossLockLockIterator")
			start = end + 1
		case <-m.ctx.Done():
			m.logger.Info("CrossLockLockIterator done")
			return
		}
	}
}

func (m *Monitor) listenCrossBurnEvent() {
	ticker := time.NewTicker(60 * time.Second)
	defer ticker.Stop()

	start := m.cHeight

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
			filter := m.ethWrapper.FilterCrossBurn(&bind.FilterOpts{Start: start, End: &end, Context: m.ctx})
			for filter.Next() {
				m.handleCrossBurn(filter.Event, true)
			}

			m.logger.WithFields(logrus.Fields{"start": start, "end": end, "current": num}).Infof("CrossBurn")
			start = end + 1
		case <-m.ctx.Done():
			m.logger.Info("CrossBurn done")
			return
		}
	}
}

func (m *Monitor) HandleCocoC() chan *Coco {
	return m.cocoC
}

func (m *Monitor) handleLock(lock *mnt.PegProxyLock, isHistory bool) {
	if !strings.EqualFold(lock.Raw.Address.String(), m.config.Eth.PegBridgeContract) {
		return
	}

	if m.storage.Has(TxKey(lock.Raw.TxHash.String())) {
		return
	}
	coco := &Coco{
		Typ:         Lock,
		IsHistory:   isHistory,
		SrcToken:    lock.SrcToken,
		DstToken:    lock.DestToken,
		Sender:      lock.From,
		Recipient:   lock.To,
		Amount:      lock.Amount,
		TxId:        lock.Raw.TxHash.String(),
		BlockHeight: lock.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"src_token":    coco.SrcToken.String(),
		"dst_token":    coco.DstToken.String(),
		"sender":       coco.Sender.String(),
		"recipient":    coco.Recipient.String(),
		"amount":       coco.Amount.String(),
		"txId":         lock.Raw.TxHash.String(),
		"block_height": lock.Raw.BlockNumber,
		"removed":      lock.Raw.Removed,
	}).Info("PegProxyLock")

	if lock.Raw.Removed {
		return
	}

	if !m.confirmEvent(lock.Raw, Lock) {
		m.logger.WithFields(logrus.Fields{
			"txId":         lock.Raw.TxHash.String(),
			"block_height": lock.Raw.BlockNumber,
		}).Info("PegProxyLock has not confirmed")
		return
	}

	m.logger.WithField("tx", lock.Raw.TxHash.String()).Info("confirmEvent")
	m.cocoC <- coco
	m.persistLBlockHeight(lock.Raw.TxHash.String(), lock.Raw.BlockNumber)
}

func (m *Monitor) handleCrossBurn(crossBurn *mnt.PegProxyCrossBurn, isHistory bool) {
	if !strings.EqualFold(crossBurn.Raw.Address.String(), m.config.Eth.PegBridgeContract) {
		return
	}

	if m.storage.Has(TxKey(crossBurn.Raw.TxHash.String())) {
		return
	}
	coco := &Coco{
		Typ:         CrossBurn,
		IsHistory:   isHistory,
		SrcToken:    crossBurn.SrcToken,
		DstToken:    crossBurn.DestToken,
		Sender:      crossBurn.From,
		Recipient:   crossBurn.To,
		Amount:      crossBurn.Amount,
		TxId:        crossBurn.Raw.TxHash.String(),
		BlockHeight: crossBurn.Raw.BlockNumber,
	}

	m.logger.WithFields(logrus.Fields{
		"src_token":    coco.SrcToken.String(),
		"dst_token":    coco.DstToken.String(),
		"sender":       coco.Sender.String(),
		"recipient":    coco.Recipient.String(),
		"amount":       coco.Amount.String(),
		"txId":         crossBurn.Raw.TxHash.String(),
		"block_height": crossBurn.Raw.BlockNumber,
		"removed":      crossBurn.Raw.Removed,
	}).Info("PegProxyCrossBurn")

	if crossBurn.Raw.Removed {
		return
	}

	if !m.confirmEvent(crossBurn.Raw, CrossBurn) {
		m.logger.WithFields(logrus.Fields{
			"txId":         crossBurn.Raw.TxHash.String(),
			"block_height": crossBurn.Raw.BlockNumber,
		}).Info("PegProxyCrossBurn has not confirmed")
		return
	}

	m.logger.WithField("tx", crossBurn.Raw.TxHash.String()).Info("confirmEvent")
	m.cocoC <- coco
	m.persistCBlockHeight(crossBurn.Raw.TxHash.String(), crossBurn.Raw.BlockNumber)
}

func (m *Monitor) confirmEvent(event types.Log, typ int) bool {
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
		var log *Coco
		switch typ {
		case Lock:
			log, err = m.GetLockLog(event.TxHash.String())
		case CrossBurn:
			log, err = m.GetCrossBurnLog(event.TxHash.String())
		}
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

func (m *Monitor) Unlock(txId string, token common.Address, from common.Address, recipient common.Address, amount *big.Int) error {
	unlocked := m.ethWrapper.TxUnlocked(txId)
	if unlocked {
		m.logger.Infof("find txUnlocked Chain %d txId:%s", txId)
		return nil
	}

	m.logger.WithFields(logrus.Fields{
		"tx_id":     txId,
		"token":     token.String(),
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

	m.ethWrapper.session.TransactOpts.Nonce = nil
	m.ethWrapper.session.TransactOpts.GasPrice = nil

	for {
		price := m.ethWrapper.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		if m.ethWrapper.session.TransactOpts.GasPrice == nil ||
			gasPrice.BigInt().Cmp(m.ethWrapper.session.TransactOpts.GasPrice) == 1 {
			m.ethWrapper.session.TransactOpts.GasPrice = gasPrice.BigInt()

			transaction, _ = m.ethWrapper.Unlock(token, from, recipient, amount, txId)
			m.ethWrapper.session.TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
			hashes = append(hashes, transaction.Hash())

			m.logger.Infof("send UnlockBor tx %s with gasPrice %s and nonce %d",
				transaction.Hash().String(), gasPrice.String(), transaction.Nonce())
		}
		receipt, err = m.ethWrapper.TransactionReceiptsLimitedRetry(context.TODO(), hashes)
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

func (m *Monitor) CrossIn(txId string, token common.Address, from common.Address, recipient common.Address, amount *big.Int) error {
	unlocked := m.ethWrapper.TxMinted(txId)
	if unlocked {
		m.logger.Infof("find TxMinted txId:%s", txId)
		return nil
	}

	m.logger.WithFields(logrus.Fields{
		"tx_id":     txId,
		"token":     token.String(),
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

	m.ethWrapper.session.TransactOpts.Nonce = nil
	m.ethWrapper.session.TransactOpts.GasPrice = nil

	for {
		price := m.ethWrapper.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		if m.ethWrapper.session.TransactOpts.GasPrice == nil ||
			gasPrice.BigInt().Cmp(m.ethWrapper.session.TransactOpts.GasPrice) == 1 {
			m.ethWrapper.session.TransactOpts.GasPrice = gasPrice.BigInt()

			transaction, _ = m.ethWrapper.CrossIn(token, from, recipient, amount, txId)
			m.ethWrapper.session.TransactOpts.Nonce = big.NewInt(int64(transaction.Nonce()))
			hashes = append(hashes, transaction.Hash())

			m.logger.Infof("send UnlockBor tx %s with gasPrice %s and nonce %d",
				transaction.Hash().String(), gasPrice.String(), transaction.Nonce())
		}
		receipt, err = m.ethWrapper.TransactionReceiptsLimitedRetry(context.TODO(), hashes)
		if err == nil {
			break
		}
	}

	if receipt.Status == 1 {
		m.logger.WithField("tx_hash", receipt.TxHash.String()).Info("CrossIn success")
	} else {
		return fmt.Errorf("unlock fail:%s", receipt.TxHash.String())
	}
	return nil
}

func (m *Monitor) GetLockLog(txId string) (*Coco, error) {
	receipt := m.ethWrapper.TransactionReceipt(context.TODO(), common.HexToHash(txId))
	for _, log := range receipt.Logs {
		if !strings.EqualFold(log.Address.String(), m.config.Eth.PegBridgeContract) {
			continue
		}

		if log.Removed {
			continue
		}
		lock, err := m.ethWrapper.pegProxy.ParseLock(*log)
		if err != nil {
			return nil, err
		}
		return &Coco{
			SrcToken:    lock.SrcToken,
			DstToken:    lock.DestToken,
			Sender:      lock.From,
			Recipient:   lock.To,
			Amount:      lock.Amount,
			TxId:        log.TxHash.String(),
			BlockHeight: receipt.BlockNumber.Uint64(),
		}, nil
	}
	return nil, fmt.Errorf("not found Lock log in tx:%s", txId)
}

func (m *Monitor) GetCrossBurnLog(txId string) (*Coco, error) {
	receipt := m.ethWrapper.TransactionReceipt(context.TODO(), common.HexToHash(txId))
	for _, log := range receipt.Logs {
		if !strings.EqualFold(log.Address.String(), m.config.Eth.PegBridgeContract) {
			continue
		}

		if log.Removed {
			continue
		}
		crossBurn, err := m.ethWrapper.pegProxy.ParseCrossBurn(*log)
		if err != nil {
			return nil, err
		}
		return &Coco{
			SrcToken:    crossBurn.SrcToken,
			DstToken:    crossBurn.DestToken,
			Sender:      crossBurn.From,
			Recipient:   crossBurn.To,
			Amount:      crossBurn.Amount,
			TxId:        log.TxHash.String(),
			BlockHeight: receipt.BlockNumber.Uint64(),
		}, nil
	}
	return nil, fmt.Errorf("not found crossBurn log in tx:%s", txId)
}

func (m *Monitor) fetchBlockNum() (uint64, error) {
	header := m.ethWrapper.HeaderByNumber(context.TODO(), nil)
	return header.Number.Uint64(), nil
}

func (m *Monitor) loadHeightFromStorage() {
	var header *types.Header
	header = m.ethWrapper.HeaderByNumber(context.TODO(), nil)

	// load block height
	b := m.storage.Get(lHeightKey())
	if b == nil {
		m.lHeight = header.Number.Uint64() - m.minConfirms
		m.persistLHeight(m.lHeight)
	} else {
		m.lHeight = binary.LittleEndian.Uint64(b)

	}

	// load block height
	c := m.storage.Get(cHeightKey())
	if c == nil {
		m.cHeight = header.Number.Uint64() - m.minConfirms
		m.persistCHeight(m.cHeight)
	} else {
		m.cHeight = binary.LittleEndian.Uint64(c)

	}

	if m.config.Eth.LockHeight != 0 {
		m.lHeight = m.config.Eth.LockHeight
	}

	if m.config.Eth.CrossBurnHeight != 0 {
		m.cHeight = m.config.Eth.CrossBurnHeight
	}

	m.logger.WithFields(logrus.Fields{
		"lock_height":  m.lHeight,
		"cross_height": m.cHeight,
	}).Info("Subscribe")
}

func lHeightKey() []byte {
	return []byte(fmt.Sprintf("lHeight"))
}

func cHeightKey() []byte {
	return []byte(fmt.Sprintf("cHeight"))
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

func (m *Monitor) persistCBlockHeight(txId string, height uint64) {
	m.persistCHeight(height)
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

func (m *Monitor) persistCHeight(height uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, height)
	m.storage.Put(cHeightKey(), buf)
	m.cHeight = height
	m.logger.WithFields(logrus.Fields{
		"height": m.cHeight,
	}).Info("Persist CrossBurn Block Height")
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
