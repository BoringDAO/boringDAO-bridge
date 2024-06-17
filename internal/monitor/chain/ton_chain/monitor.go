package ton

import (
	"context"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/boringdao/bridge/internal/monitor"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/hexutil"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
)

type Monitor struct {
	index   uint64
	wrapper *Wrapper
	cocoC   chan *monitor.Coco
	logger  logrus.FieldLogger
	storage storage.Storage
	config  *repo.EdgeTonConfig
	tokens  map[string]string
	mut     sync.Mutex
	ctx     context.Context
	cancel  context.CancelFunc
}

func New(repoRoot string, config *repo.EdgeTonConfig, ds []uint64, logger logrus.FieldLogger) (monitor.Mnt, error) {
	storagePath := repo.GetStoragePath(repoRoot, fmt.Sprintf("%s_%d", config.Name, config.ChainID))
	storage, err := leveldb.New(storagePath)
	if err != nil {
		return nil, err
	}

	wrapper, err := NewWrapper(logger, config)
	if err != nil {
		return nil, err
	}
	logger.Infof("Ton Tokens: %v", config.Tokens)

	tokens := make(map[string]string)
	for jettonAddr, evmAddr := range config.Tokens {
		tokens[jettonAddr] = evmAddr
		tokens[evmAddr] = jettonAddr
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Monitor{
		config:  config,
		storage: storage,
		wrapper: wrapper,
		cocoC:   make(chan *monitor.Coco),
		tokens:  tokens,
		logger:  logger,
		ctx:     ctx,
		cancel:  cancel,
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
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			index, err := m.wrapper.GetEventAddress(m.ctx, uint(m.index))
			if err != nil {
				m.logger.Errorf("Ton GetEventAddress failed: %s", err)
				continue
			}
			event, err := m.wrapper.GetEventData(m.ctx, index)
			if err != nil {
				m.logger.Errorf("Ton GetEventData failed: %s", err)
				continue
			}
			if event == nil {
				continue
			}
			typ := monitor.CrossOuted
			if event.Event.Opcode == DEPOSIT {
				typ = monitor.CrossOuted
			}

			if !common.IsHexAddress(event.Event.ToAddress) {
				m.logger.Errorf("Ton GetEventData failed: %s[%d], %s", "invalid address", m.index, event.Event.ToAddress)
				m.index++
				m.persistIndex(m.index)
				continue
			}
			toToken, ok := m.tokens[event.Event.JettonWalletAddress.String()]
			if !ok {
				m.logger.Errorf("Ton GetEventData failed: %s[%d], %s", "invalid token", m.index, event.Event.JettonWalletAddress.String())
				m.index++
				m.persistIndex(m.index)
				continue
			}
			m.cocoC <- &monitor.Coco{
				Typ:         typ,
				TxId:        event.Event.EventAddress,
				BlockHeight: 0,
				Index:       0,
				FromChainId: big.NewInt(int64(m.config.ChainID)),
				FromToken:   common.BytesToAddress(event.Event.JettonWalletAddress.Data()),
				From:        event.Event.FromUser.Data(),
				ToChainId:   event.Event.ToChainId,
				ToToken:     common.HexToAddress(toToken),
				To:          common.HexToAddress(event.Event.ToAddress).Bytes(),
				Amount:      event.Event.JettonAmount,
			}
			m.index++
			m.persistIndex(m.index)
		case <-m.ctx.Done():
			m.logger.Info("CrossLockLockIterator done")
			return
		}
	}
}

func (m *Monitor) HandleCocoC() chan *monitor.Coco {
	return m.cocoC
}

func (m *Monitor) CrossIn(fromToken, toToken common.Address, from, to []byte, fromChainID, toChainID, amount *big.Int, txId string) error {
	orderId := crypto.Keccak256Hash([]byte(txId)).Hex()
	exists, err := m.wrapper.isOrderHandled(m.ctx, txId)
	if err != nil {
		return err
	}
	if exists {
		m.logger.Warnf("order %s already handled", txId)
		return nil
	}

	m.logger.WithFields(logrus.Fields{
		"tx_id":       txId,
		"from_token":  fromToken.String(),
		"to_token":    toToken.String(),
		"from":        hexutil.Encode(from),
		"to":          hexutil.Encode(to),
		"fromChainId": fromChainID.String(),
		"toChainId":   toChainID.String(),
		"amount":      amount.String(),
	}).Info("will crossIn")
	jettonToken, ok := m.tokens[toToken.Hex()]
	if !ok {
		m.logger.Errorf("Ton CrossIn failed: %s, %s", "invalid token", toToken.Hex())
		return fmt.Errorf("invalid token %s", toToken.Hex())
	}
	return m.wrapper.CrossIn(m.ctx, jettonToken, common.BytesToAddress(from).Hex(), string(to), amount, orderId)
}

func (m *Monitor) Name() string {
	return m.config.Name
}

func (m *Monitor) HasTx(txId string, coco *monitor.Coco) bool {
	return m.storage.Has(TxKey(txId, coco.Typ, coco.Index))
}

func indexKey() []byte {
	return []byte("index")
}

func (m *Monitor) loadIndexFromStorage() {
	buf := m.storage.Get(indexKey())
	if buf != nil {
		m.index = binary.LittleEndian.Uint64(buf)
	} else {
		m.index = 0
	}
	if m.config.Index != 0 {
		m.index = m.config.Index
	}

	m.logger.WithFields(logrus.Fields{
		"index":   m.index,
		"chainID": m.config.ChainID,
	}).Info("Subscribe")
}

func TxKey(hash string, typ int, idx uint) []byte {
	return []byte(fmt.Sprintf("tx-%d-%s-%d", typ, hash, idx))
}

func (m *Monitor) PutTxID(txId string, coco *monitor.Coco) {
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

func (m *Monitor) persistIndex(index uint64) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf, index)
	m.storage.Put(indexKey(), buf)
	m.logger.Infof("handled cross out events  Index:[%d]", index)
}
