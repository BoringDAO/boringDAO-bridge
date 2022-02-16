package bridge

import (
	"math/big"
)

const Deposited = 0
const CrossOuted = 1
const ForwardCrossOuted = 2
const Withdrawed = 3

type Coco struct {
	Typ         int      `json:"typ"`
	TxId        string   `json:"tx_id"`
	BlockHeight uint64   `json:"block_height"`
	Index       uint     `json:"index"`
	FromChainId *big.Int `json:"from_chain_id"`
	FromToken   string   `json:"from_token"`
	From        string   `json:"from"`
	ToChainId   *big.Int `json:"to_chain_id"`
	ToToken     string   `json:"to_token"`
	To          string   `json:"to"`
	Amount      *big.Int `json:"amount"`
}

type Mnt interface {
	Start() error

	CrossIn(fromToken, toToken string, from, to string, fromChainID, toChainID, amount *big.Int, txid string) error

	HandleCocoC() chan *Coco

	HasTx(txId string, coco *Coco) bool

	PutTxID(txId string, coco *Coco)

	Name() string

	MntLock()

	MntUnlock()
}
