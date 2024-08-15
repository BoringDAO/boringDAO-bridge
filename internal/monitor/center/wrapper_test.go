package center_chain

import (
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"testing"

	"github.com/boringdao/bridge/internal/repo"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestEvent(t *testing.T) {
	wrapper, err := NewWrapper(&repo.CenterConfig{
		Name:           "amoy",
		Addrs:          []string{"https://rpc.ankr.com/polygon_amoy"},
		ChainID:        80002,
		MinConfirms:    1,
		PrivKey:        generatePrivateKey(),
		GasLimit:       300000,
		GasFeeRate:     1.2,
		Index:          map[uint64]uint64{},
		CenterContract: "0xBA6E3cfAfd0bF84e5d4E93560a7eac6c8c59201f",
	}, &logrus.Logger{})
	require.NoError(t, err)
	chainId := big.NewInt(11155111)
	idx := wrapper.Index(chainId)
	for i := 0; i <= int(idx.Uint64()); i++ {
		height := wrapper.IndexHeight(chainId, big.NewInt(int64(i)))
		fmt.Printf("idx: %d, height: %d\n", i, height)
	}
}
func generatePrivateKey() string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	return hex.EncodeToString(privateKeyBytes)
}
