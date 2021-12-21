package monitor

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const Deposited = 0
const CrossOuted = 1
const CrossInFailed = 2
const Withdrawed = 3

type Coco struct {
	Typ         int            `json:"typ"`
	TxId        string         `json:"tx_id"`
	BlockHeight uint64         `json:"block_height"`
	Index       uint           `json:"index"`
	FromChainId *big.Int       `json:"from_chain_id"`
	FromToken   common.Address `json:"from_token"`
	From        common.Address `json:"from"`
	ToChainId   *big.Int       `json:"to_chain_id"`
	ToToken     common.Address `json:"to_token"`
	To          common.Address `json:"to"`
	Amount      *big.Int       `json:"amount"`
}

type Mnt interface {
	Start() error

	CrossIn(fromToken, toToken common.Address, from, to common.Address, fromChainID, toChainID, amount *big.Int, txid string) error

	HandleCocoC() chan *Coco

	HasTx(txId string, coco *Coco) bool

	PutTxID(txId string, coco *Coco)

	Name() string

	MntLock()

	MntUnlock()
}
