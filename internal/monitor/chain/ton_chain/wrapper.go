package ton

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

const DEPOSIT uint64 = 0x0f20e64b
const CROSSOUT uint64 = 0xc75ea505

const SEND_MESSAGE uint64 = 0xf1381e5b
const NEW_ORDER uint64 = 0xf718510f

type Wrapper struct {
	client       *liteclient.ConnectionPool
	api          ton.APIClientWrapped
	wallet       *wallet.Wallet
	logger       logrus.FieldLogger
	contractAddr *address.Address
	ctx          context.Context
	cancel       context.CancelFunc
}

type EventDate struct {
	MultisigAddress string
	IsInit          *big.Int
	EventIndex      *big.Int
	Event           *Event
}

type Event struct {
	Opcode              uint64
	EventAddress        string
	QueryId             *big.Int
	JettonWalletAddress *address.Address
	FromUser            *address.Address
	JettonAmount        *big.Int
	ToChainId           *big.Int
	ToAddress           string
	EventCount          *big.Int
}

func NewWrapper(logger logrus.FieldLogger, config *repo.EdgeTonConfig) (*Wrapper, error) {
	client := liteclient.NewConnectionPool()

	err := client.AddConnectionsFromConfigUrl(context.Background(), config.Addrs[0])
	if err != nil {
		panic(err)
	}
	api := ton.NewAPIClient(client).WithTimeout(30 * time.Second).WithRetry(10)

	wl, err := wallet.FromSeed(api, strings.Split(config.MNEMONIC, " "), wallet.V4R2)
	if err != nil {
		return nil, fmt.Errorf("failed to create ton  wallet: %w", err)
	}
	logger.Info("ton wallet: ", wl.Address())

	ctx, cancel := context.WithCancel(context.Background())
	return &Wrapper{
		client:       client,
		api:          api,
		wallet:       wl,
		logger:       logger,
		contractAddr: address.MustParseAddr(config.EdgeContract),
		ctx:          ctx,
		cancel:       cancel,
	}, nil
}

func (w *Wrapper) Close() error {
	w.client.Stop()
	return nil
}

func (w *Wrapper) HeaderByNumber(ctx context.Context, number uint32) tlb.BlockHeader {
	data, err := w.api.GetBlockData(
		ctx,
		&ton.BlockIDExt{Workchain: 0, SeqNo: number},
	)
	if err != nil {
		logrus.Printf("HeaderByNumber err: %s", err.Error())
		return tlb.BlockHeader{}
	}
	return data.BlockInfo
}

func (w *Wrapper) CrossIn(ctx context.Context, jettonToken, fromAddr, toAddr string, amount *big.Int, txId string) error {
	var err error
	msgBase64 := ""
	retry.Retry(func(attempt uint) error {
		msgBase64, err = getCrossInBoc(fromAddr, toAddr, amount.String(), jettonToken, w.contractAddr.String(), txId)
		if err != nil {
			logrus.Printf("getCrossInBoc err: %s", err.Error())
		}
		return err
	}, strategy.Wait(3*time.Second))
	data, err := base64.StdEncoding.DecodeString(msgBase64)
	if err != nil {
		return err
	}
	boc, err := cell.FromBOC(data)
	if err != nil {
		return err
	}
	tx, block, err := w.wallet.SendWaitTransaction(w.ctx, &wallet.Message{
		Mode: 1,
		InternalMessage: &tlb.InternalMessage{
			IHRDisabled: true,
			Bounce:      true,
			DstAddr:     w.contractAddr,
			Body:        boc,
			Amount:      tlb.MustFromTON("0.05"),
		},
	})
	if err != nil {
		return err
	}
	w.logger.Infof("Ton CrossIn tx: %s, block: %d", hex.EncodeToString(tx.Hash), block.SeqNo)

	return nil
}

func (w *Wrapper) GetEventAddress(ctx context.Context, eventIndex uint) (*address.Address, error) {
	block, err := w.api.CurrentMasterchainInfo(ctx)
	if err != nil {
		return nil, err
	}

	res, err := w.api.WaitForBlock(block.SeqNo).RunGetMethod(ctx, block, w.contractAddr, "get_event_address", eventIndex)
	if err != nil {
		return nil, err
	}
	slice, err := res.Slice(0)
	if err != nil {
		return nil, err
	}

	return slice.LoadAddr()
}

func (w *Wrapper) isOrderHandled(ctx context.Context, txId string) (bool, error) {
	block, err := w.api.CurrentMasterchainInfo(ctx)
	if err != nil {
		return false, err
	}

	orderId := crypto.Keccak256Hash([]byte(txId)).Big()
	res, err := w.api.WaitForBlock(block.SeqNo).RunGetMethod(ctx, block, w.contractAddr, "get_order_address", orderId)
	if err != nil {
		return false, err
	}
	slice, err := res.Slice(0)
	if err != nil {
		return false, err
	}

	orderAddr, err := slice.LoadAddr()
	if err != nil {
		return false, err
	}

	_, err = w.api.WaitForBlock(block.SeqNo).RunGetMethod(ctx, block, orderAddr, "get_order_data")
	if err != nil {
		if strings.Contains(err.Error(), "contract is not initialized") {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (w *Wrapper) GetOrderData(ctx context.Context, orderAddr *address.Address) error {
	block, err := w.api.CurrentMasterchainInfo(ctx)
	if err != nil {
		return err
	}

	_, err = w.api.WaitForBlock(block.SeqNo).RunGetMethod(ctx, block, orderAddr, "get_order_data")
	if err != nil {
		if strings.Contains(err.Error(), "contract is not initialized") {
			return nil
		}
		return err
	}
	return nil
}

func (w *Wrapper) GetEventData(ctx context.Context, eventAddr *address.Address) (*EventDate, error) {
	block, err := w.api.CurrentMasterchainInfo(ctx)
	if err != nil {
		return nil, err
	}

	getEventData, err := w.api.WaitForBlock(block.SeqNo).RunGetMethod(ctx, block, eventAddr, "get_event_data")
	if err != nil {
		if strings.Contains(err.Error(), "contract is not initialized") {
			return nil, nil
		}
		return nil, err
	}

	return buildEventData(eventAddr, getEventData), nil
}

func buildEventData(eventAddr *address.Address, eventData *ton.ExecutionResult) *EventDate {
	multisigAddress := eventData.MustSlice(0)
	multisigAddr := multisigAddress.MustLoadAddr()
	multisigAddressString := multisigAddr.String()

	isInit := eventData.MustInt(1)
	if isInit.Cmp(big.NewInt(0)) == 0 {
		return nil
	}
	eventIndex := eventData.MustInt(2)

	event := eventData.MustCell(3)
	eventBody := event.BeginParse()

	opcode := eventBody.MustLoadUInt(32)

	e := &Event{}
	e.Opcode = opcode
	e.EventAddress = eventAddr.String()
	if opcode == DEPOSIT {
		queryId := eventBody.MustLoadBigUInt(64)
		bodyRef := eventBody.MustLoadRef()
		jettonWalletAddress := bodyRef.MustLoadBigUInt(256)
		fromUser := bodyRef.MustLoadAddr()
		jettonAmount := bodyRef.MustLoadBigUInt(256)
		extraRef := eventBody.MustLoadRef()
		toAddress := extraRef.MustLoadBinarySnake()
		eventCount := eventBody.MustLoadBigUInt(256)

		e.QueryId = queryId
		e.ToAddress = fmt.Sprintf("%x", toAddress)
		e.JettonWalletAddress = address.MustParseRawAddr("0:" + padStart(jettonWalletAddress.Text(16), 64, "0"))
		e.FromUser = fromUser
		e.JettonAmount = jettonAmount
		e.EventCount = eventCount
	} else if opcode == CROSSOUT {
		queryId := eventBody.MustLoadBigUInt(64)
		bodyRef := eventBody.MustLoadRef()
		jettonWalletAddress := bodyRef.MustLoadBigUInt(256)
		fromUser := bodyRef.MustLoadAddr()
		jettonAmount := bodyRef.MustLoadBigUInt(256)
		extraRef := eventBody.MustLoadRef()
		toChainId := extraRef.MustLoadBigUInt(256)
		toAddress := extraRef.MustLoadBinarySnake()
		eventCount := eventBody.MustLoadBigUInt(256)

		e.QueryId = queryId
		e.JettonWalletAddress = address.MustParseRawAddr("0:" + padStart(jettonWalletAddress.Text(16), 64, "0"))
		e.FromUser = fromUser
		e.JettonAmount = jettonAmount
		e.EventCount = eventCount
		e.ToChainId = toChainId
		e.ToAddress = fmt.Sprintf("%x", toAddress)
	} else {
		log.Fatalln("unknown opcode")
	}

	return &EventDate{
		MultisigAddress: multisigAddressString,
		IsInit:          isInit,
		EventIndex:      eventIndex,
		Event:           e,
	}
}

func getCrossInBoc(fromAddr, toAddr, amount, jettonWalletAddress, multisigAddress, orderId string) (string, error) {
	url := fmt.Sprintf("http://localhost:3000/getCrossInBoc?fromAddr=%s&toAddr=%s&amount=%s&jettonWalletAddress=%s&multisigAddress=%s&orderId=%s",
		fromAddr, toAddr, amount, jettonWalletAddress, multisigAddress, orderId)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("accept", "*/*")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func padStart(str string, length int, padChar string) string {
	if len(str) >= length {
		return str
	}
	padding := strings.Repeat(padChar, length-len(str))
	return padding + str
}
