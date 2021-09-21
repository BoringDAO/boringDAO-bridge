package monitor

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Coco struct {
	IsHistory      bool           `json:"isHistory"`
	OriginToken0   common.Address `json:"origin_token0"`
	OriginChainId  *big.Int       `json:"origin_chain_id"`
	OriginTokenId  *big.Int       `json:"origin_token_id"`
	OriginTokenUri string         `json:"origin_token_uri"`
	TargetChainID  *big.Int       `json:"target_chain_id"`
	From           common.Address `json:"from"`
	To             common.Address `json:"to"`
	TxId           string         `json:"tx_id"`
	SourceChainID  uint64         `json:"source_chain_id"`
	BlockHeight    uint64         `json:"block_height"`
}

type IMonitor interface {
	Start() error

	Stop() error

	ListenCrossOutC() chan *Coco

	HandleCrossIn(coco *Coco)

	ListenFinishedCocoC() chan *Coco

	HandleFinishedCoco(coco *Coco)
}
