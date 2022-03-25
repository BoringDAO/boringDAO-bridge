package monitor

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Coco struct {
	FromToken   common.Address `json:"from_token"`
	ToToken     common.Address `json:"to_token"`
	FromChainId *big.Int       `json:"from_chain_id"`
	ToChainId   *big.Int       `json:"to_chain_id"`
	From        common.Address `json:"from"`
	To          common.Address `json:"to"`
	Amount      *big.Int       `json:"amount"`
	Index       uint64         `json:"index"`
	TxId        string         `json:"tx_id"`
	BlockHeight uint64         `json:"block_height"`
}

type IMonitor interface {
	Start() error

	Stop() error

	ListenCrossOutC() chan *Coco

	HandleCrossIn(coco *Coco)

	ListenFinishedCocoC() chan *Coco

	HandleFinishedCoco(coco *Coco)
}
