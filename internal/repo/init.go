package repo

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pelletier/go-toml"

	"github.com/fatih/color"

	"github.com/boringdao/bridge/pkg/kit/fileutil"
)

const (
	packPath = "../../config"
)

var supportedChain = []string{"eth", "bsc", "okex", "polygon", "avalanche", "harmony", "heco", "fantom", "xdai"}

func Initialize(repoRoot string) error {
	config := Config{
		RepoRoot: repoRoot,
		Token:    make(map[string]string),
		Log: Log{
			Level:    "info",
			Dir:      "logs",
			Filename: "bridge.log",
			Module:   map[string]string{"app": "info"}},
	}

	color.Blue("1. Configure tokens you care about:")
	for i := 1; ; i++ {
		fmt.Printf("token %d address:", i)
		token := readInput()
		fmt.Printf("chain ID of token %s:", token)
		chainID := readInput()
		config.Token[token] = chainID

		fmt.Print("have next token? Y/n:")
		if !ReadYes() {
			break
		}
	}

	color.Blue("2. Configure bridges:")
	for i := 1; ; i++ {
		color.Green("Start to configure bridge %d", i)
		index, chain := selectChain(i)
		chainID := readChainID(chain)
		addrs := readRpcAddrs(chain)
		contract := readBridgeContract(chain)
		minConfirms := readMinConfirms(chain)
		gasLimit := readGasLimit(chain)

		config.Bridges = append(config.Bridges, &BridgeConfig{
			Name:           chain,
			Addrs:          addrs,
			ChainID:        chainID,
			MinConfirms:    minConfirms,
			GasLimit:       gasLimit,
			BridgeContract: contract,
		})
		config.Log.Module[chain] = "info"
		color.Green("have next bridge to configure? Y/n:")
		if !ReadYes() {
			break
		}
		supportedChain = append(supportedChain[0:index], supportedChain[index+1:]...)
	}

	data, err := toml.Marshal(config)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath.Join(repoRoot, configName), data, 0666); err != nil {
		return err
	}

	return nil
}

func Initialized(repoRoot string) bool {
	return fileutil.Exist(filepath.Join(repoRoot, configName))
}

func readInput() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}

func selectChain(i int) (int, string) {
	var (
		chainIndex int
		err        error
	)

	var s string
	for i, chain := range supportedChain {
		s += fmt.Sprintf("[%d]%s ", i, chain)
	}

	for {
		fmt.Printf("Please select blockchain for bridge %d, %s:", i, s)
		chainIndexStr := readInput()
		chainIndex, err = strconv.Atoi(chainIndexStr)
		if err != nil {
			color.Red("invalid chain index, should be a number")
		} else if chainIndex < 0 || chainIndex >= len(supportedChain) {
			color.Red("invalid chain index, out of range")
		} else {
			break
		}
	}

	return chainIndex, supportedChain[chainIndex]
}

func readChainID(chain string) uint64 {
	var (
		chainID int
		err     error
	)
	for {
		fmt.Printf("chain ID of %s:", chain)
		chainIDStr := readInput()
		chainID, err = strconv.Atoi(chainIDStr)
		if err != nil {
			fmt.Println("invalid chain index, should be a number")
		} else {
			break
		}
	}

	return uint64(chainID)
}

func readRpcAddrs(chain string) []string {
	var addrs []string

	for {
		fmt.Printf("rpc address of %s:", chain)
		addrs = append(addrs, readInput())
		fmt.Print("have next rpc address? Y/n:")
		if !ReadYes() {
			break
		}
	}

	return addrs
}

func readBridgeContract(chain string) string {
	fmt.Printf("bridge contract address of %s:", chain)
	return readInput()
}

func readMinConfirms(chain string) uint64 {
	var (
		minConfirms int
		err         error
	)

	fmt.Printf("minimum blocks to confirm of %s, use 30 by default? Y/n:", chain)
	if ReadYes() {
		return 30
	}

	for {
		fmt.Printf("input minimum blocks to confirm of %s:", chain)
		minConfirmsStr := readInput()
		minConfirms, err = strconv.Atoi(minConfirmsStr)
		if err != nil {
			fmt.Println("invalid input, should be a number")
		} else {
			break
		}
	}

	return uint64(minConfirms)
}

func readGasLimit(chain string) uint64 {
	var (
		gasLimit int
		err      error
	)
	fmt.Printf("gas limit of %s, use 1500000 by default? Y/n:", chain)
	if ReadYes() {
		return 1500000
	}

	for {
		fmt.Printf("input gas limit of %s:", chain)
		gasLimitStr := readInput()
		gasLimit, err = strconv.Atoi(gasLimitStr)
		if err != nil {
			fmt.Println("invalid input, should be a number")
		} else {
			break
		}
	}

	return uint64(gasLimit)
}

func ReadYes() bool {
	s := readInput()
	return strings.EqualFold(s, "y") || strings.EqualFold(s, "")
}
