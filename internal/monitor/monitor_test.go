package monitor

import (
	"context"
	"fmt"
	"testing"

	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/log"
	"github.com/stretchr/testify/assert"
)

func TestBridgeWrapper_HeaderByNumber(t *testing.T) {
	config, err := repo.UnmarshalConfig("../../../config")
	assert.Nil(t, err)

	for _, bConfig := range config.Bridges {
		bConfig.PrivKey = "0x7de69ec0b3e74a7d16f13d68288beb7a44fad919c27b7b50f294394ec818d364"

		bw, err := NewBridgeWrapper(bConfig, log.NewWithModule("harmony"))
		assert.Nil(t, err)

		number, err := bw.ethClient.BlockNumber(context.TODO())
		assert.Nil(t, err)

		fmt.Println(number)
	}
}
