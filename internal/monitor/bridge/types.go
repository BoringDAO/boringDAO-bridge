package bridge

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Coco struct {
	IsHistory   bool           `json:"isHistory"`
	Coin        string         `json:"coin"`
	Sender      common.Address `json:"sender"`
	Recipient   common.Address `json:"recipient"`
	Amount      *big.Int       `json:"amount"`
	TxId        string         `json:"tx_id"`
	BlockHeight uint64         `json:"block_height"`
	Token0      common.Address `json:"token0"`
	Token1      common.Address `json:"token1"`
	ChainID     *big.Int       `json:"chain_id"`
}

type IMonitor interface {
	Start() error

	CrossMint(ethToken common.Address, txId string, addrFromEth common.Address, recipient common.Address, amount *big.Int) error

	ListenCrossBurnEvent()

	HandleCocoC() chan *Coco

	HasTx(txId string) bool

	PutTxID(txId string, coco *Coco)
}
