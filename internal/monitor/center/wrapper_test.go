package center_chain

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/stretchr/testify/require"

	"github.com/boringdao/bridge/internal/monitor/contracts/center"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func TestFilter(t *testing.T) {
	rpcClient, err := rpc.DialContext(context.Background(), "https://rpc-mainnet.maticvigil.com")
	require.Nil(t, err)
	etherCli := ethclient.NewClient(rpcClient)
	twoWay, err := center.NewTwoWayCenter(common.HexToAddress("0x6dc551088afaf828b95a9c9ad590aefe797c8e87"), etherCli)
	require.Nil(t, err)

	end := uint64(23711026)
	filter, err := twoWay.FilterForwardCrossOuted(&bind.FilterOpts{
		Start:   23710963,
		End:     &end,
		Context: context.Background(),
	})
	require.Nil(t, err)
	for filter.Next() {
		_ = filter.Event
	}
}
