package ton

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/boringdao/bridge/internal/repo"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
	"github.com/xssnick/tonutils-go/ton/wallet"
	"github.com/xssnick/tonutils-go/tvm/cell"
)

const configUrl string = "https://ton.org/global.config.json"

const multisigContractAddress string = "EQAoWYTgcJM7SCzmJvn0UjicfgZRvfVKa7WQSo3dSUb7X7Rn"
const eventIndex uint = 8

func TestWrapper_GetEventData(t *testing.T) {

	w, _ := NewWrapper(&logrus.Logger{}, &repo.EdgeTonConfig{
		Addrs:        []string{configUrl},
		EdgeContract: multisigContractAddress,
		MNEMONIC:     strings.Join(wallet.NewSeed(), " "),
	})
	address, _ := w.GetEventAddress(context.Background(), eventIndex)
	fmt.Println("address: ", address)

	e, err := w.GetEventData(context.Background(), address)
	require.NoError(t, err)
	data, _ := json.Marshal(e)
	fmt.Println("e: ", string(data))
}

func TestWrapper_CrossIn(t *testing.T) {
	words := ""
	w, _ := NewWrapper(&logrus.Logger{}, &repo.EdgeTonConfig{
		Addrs:        []string{configUrl},
		EdgeContract: multisigContractAddress,
		MNEMONIC:     words,
	})
	txId := "12345678"
	orderId := crypto.Keccak256Hash([]byte(txId)).Hex()
	fmt.Println("orderId: ", orderId)
	// err := w.CrossIn(context.Background(), "EQCxE6mUtQJKFnGfaROTKOt1lZbDiiX1kCixRv7Nw2Id_sDs", "0xd75d38de8C8227097Df34A439C1a553F6E31372e", "UQA8ue90JugDvs0wvLR35tuiraNFV7-_Uxczu7Ku-xK1tjgi", big.NewInt(1), orderId)
	// require.NoError(t, err)

	exist, err := w.isOrderHandled(context.Background(), txId)
	require.NoError(t, err)
	fmt.Println("exist: ", exist)

}

// te6cckEBAQEAWQAArg+KfqVd3sSIJfkh4RAYAHlz3uhN0Ad9mmF5aO/Nt0VbRoqvf36mLmd3ZV32JWttACgqbij9YuuM55MpGaWYlN9Sc/EImulJyyKf+47yrLtQAgIAAAAAMZTEEnU=
// te6cckEBAgEAXAABpA+KfqUAAAAAAAAAABAYAHlz3uhN0Ad9mmF5aO/Nt0VbRoqvf36mLmd3ZV32JWttACgqbij9YuuM55MpGaWYlN9Sc/EImulJyyKf+47yrLtQAgMBAAoAAAAAMbIWb4I=

// te6cckEBBQEAnwABAdEBAQrxOB5bAQIBaCIAdNAttakBuwz0+ruabnD48YqznfsXDJtMnUtil23jfFwgF9eEAAAAAAAAAAAAAAAAAAEDAaQPin6lAAAAAAAAAAAQGAB5c97oTdAHfZpheWjvzbdFW0aKr39+pi5nd2Vd9iVrbQAoKm4o/WLrjOeTKRmlmJTfUnPxCJrpScsin/uO8qy7UAIDBAAKAAAAADFhq00y
// te6cckEBBAEAnAABAcABAQvR4nA8tgMCAWYiAHTQLbWpAbsM9Pq7mm5w+PGKs537FwybTJ1LYpdt43xcHMS0AAAAAAAAAAAAAAAAAAEDAK4Pin6lXd7EiCX5IeEQGAB5c97oTdAHfZpheWjvzbdFW0aKr39+pi5nd2Vd9iVrbQAoKm4o/WLrjOeTKRmlmJTfUnPxCJrpScsin/uO8qy7UAICAAAAADHbFfh5

func TestCell(t *testing.T) {
	cellBase64 := "te6cckEBBAEATAABGAAAAAIAAAAAAAAAAAEBGaAAAAAAAAAAAJzEtAQCAUOAGiWY6NS1Ble1dz0Eq/najuEgguKYaHuH4JkM39HRQQ0wAwAMMS5qc29uiB4bhQ=="
	data, err := base64.StdEncoding.DecodeString(cellBase64)
	require.NoError(t, err)
	boc, err := cell.FromBOC(data)
	require.NoError(t, err)
	fmt.Print(boc.ToBOC())
}
