package monitor

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const Lock = 0
const CrossBurn = 1
const Rollback = 2

type Coco struct {
	Typ         int            `json:"typ"`
	From        common.Address `json:"from"`
	To          common.Address `json:"to"`
	Amount      *big.Int       `json:"amount"`
	TxId        string         `json:"tx_id"`
	BlockHeight uint64         `json:"block_height"`
	Index       uint           `json:"index"`
	Token0      common.Address `json:"token0"`
	Token1      common.Address `json:"token1"`
	ChainID0    *big.Int       `json:"chain_id_0"`
	ChainID1    *big.Int       `json:"chain_id_1"`
}

type Mnt interface {
	Start() error

	CrossIn(txId string, token common.Address, from common.Address, recipient common.Address, chainID, amount *big.Int) error

	Unlock(txId string, token common.Address, from common.Address, recipient common.Address, chainID, amount *big.Int) error

	Rollback(txId string, token common.Address, from common.Address, recipient common.Address, chainID, amount *big.Int) error

	HandleCocoC() chan *Coco

	HasTx(txId string, coco *Coco) bool

	PutTxID(txId string, coco *Coco)

	Name() string

	MntLock()

	MntUnlock()
}
