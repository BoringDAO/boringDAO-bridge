package chain

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/stretchr/testify/require"

	"github.com/boringdao/bridge/internal/monitor/contracts/edge"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func Test_TxHandled(t *testing.T) {
	rpcClient, err := rpc.DialContext(context.Background(), "https://andromeda.metis.io/?owner=1088")
	require.Nil(t, err)
	etherCli := ethclient.NewClient(rpcClient)

	twoWay, err := edge.NewTwoWayEdge(common.HexToAddress("0x57E05346A576A3f75972aC0E7c9d9f46765B346E"), etherCli)
	require.Nil(t, err)
	session := &edge.TwoWayEdgeSession{
		Contract: twoWay,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
	}
	unlocked, err := session.TxHandled("0xeb43e69e77f7076825adf73e0d340aa720aabd0033de9041548c1eb52d139750#ForwardCrossOuted")
	require.Nil(t, err)
	require.True(t, unlocked)
}
