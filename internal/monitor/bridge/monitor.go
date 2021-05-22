package bridge

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
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

type Monitor interface {
	Start() error

	CrossMint(ethToken common.Address, txId string, addrFromEth common.Address, recipient common.Address, amount *big.Int) error

	ListenCrossBurnEvent()

	HandleCocoC() chan *Coco

	HasTx(txId string) bool

	PutTxID(txId string, coco *Coco)
}
