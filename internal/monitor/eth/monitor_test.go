package eth

import (
	"context"
	"fmt"
	"testing"

	"github.com/boringdao/bridge/internal/loggers"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestNewLockWrapper(t *testing.T) {
	config, err := repo.UnmarshalConfig("../../../config")
	assert.Nil(t, err)

	config.Eth.PrivKey = "0x7de69ec0b3e74a7d16f13d68288beb7a44fad919c27b7b50f294394ec818d364"

	lw, err := NewLockWrapper(&config.Eth, log.NewWithModule(loggers.ETH))
	assert.Nil(t, err)
	assert.NotNil(t, lw)
}

func TestLockWrapper_HeaderByNumber(t *testing.T) {
	config, err := repo.UnmarshalConfig("../../../config")
	assert.Nil(t, err)

	config.Eth.PrivKey = "0x7de69ec0b3e74a7d16f13d68288beb7a44fad919c27b7b50f294394ec818d364"

	config.Eth.Addrs[0] = "https://opsten.infura.io/v3/3e8a2d2286964b62bc6b1d67c379badd"
	lw, err := NewLockWrapper(&config.Eth, log.NewWithModule(loggers.ETH))
	assert.Nil(t, err)

	header := lw.HeaderByNumber(context.TODO(), nil)
	assert.NotNil(t, header)

	fmt.Println(header.Number.String())
}

func newWrapper(t *testing.T) *LockWrapper {
	config, err := repo.UnmarshalConfig("../../../config")
	assert.Nil(t, err)

	config.Eth.PrivKey = "0x7de69ec0b3e74a7d16f13d68288beb7a44fad919c27b7b50f294394ec818d364"

	config.Eth.Addrs[0] = "https://ropsten.infura.io/v3/3e8a2d2286964b62bc6b1d67c379badd"
	lw, err := NewLockWrapper(&config.Eth, log.NewWithModule(loggers.ETH))
	assert.Nil(t, err)

	return lw
}

func TestLockWrapper_FilterLock(t *testing.T) {
	lw := newWrapper(t)

	start := uint64(9868259)
	end := start + 100
	iterator := lw.FilterLock(&bind.FilterOpts{Start: start, End: &end, Context: context.TODO()})
	assert.NotNil(t, iterator)

	fmt.Println(iterator.Next())
}

func TestLockWrapper_SuggestGasPrice(t *testing.T) {
	lw := newWrapper(t)

	result := lw.SuggestGasPrice(context.TODO())
	assert.NotNil(t, result)

	fmt.Println(result.String())

	gasPrice := decimal.NewFromBigInt(result, 0).Mul(decimal.NewFromFloat(1.2))
	fmt.Println(gasPrice.String())
}
