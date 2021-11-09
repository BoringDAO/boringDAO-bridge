package repo

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/boringdao/bridge/pkg/kit/fileutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fatih/color"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr"
	"github.com/pelletier/go-toml"
)

const (
	packPath = "../../config"
)

var supportedChain = []string{"eth", "bsc", "okex", "polygon", "avalanche", "harmony", "heco", "fantom", "xdai"}

func Initialize(repoRoot string, interactive bool) error {
	if interactive {
		return initializeInteractively(repoRoot)
	}

	box := packr.NewBox(packPath)
	if err := box.Walk(func(s string, file packd.File) error {
		p := filepath.Join(repoRoot, s)
		dir := filepath.Dir(p)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			err := os.MkdirAll(dir, 0755)
			if err != nil {
				return err
			}
		}

		return ioutil.WriteFile(p, []byte(file.String()), 0644)
	}); err != nil {
		return err
	}

	return nil
}

func initializeInteractively(repoRoot string) error {
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
		token := ReadEvmAddress(fmt.Sprintf("token %d address:", i))
		chainID := readChainID(fmt.Sprintf("chain ID of token %s:", token))
		config.Token[token] = strconv.Itoa(int(chainID))

		fmt.Print("have next token? Y/n:")
		if !ReadYes() {
			break
		}
	}

	color.Blue("2. Configure bridges:")
	for i := 1; ; i++ {
		color.Green("Start to configure bridge %d", i)
		index, chain := selectChain(fmt.Sprintf("[1/6] Please select blockchain for bridge %d", i))
		chainID := readChainID(fmt.Sprintf("[2/6] Please input chain ID of  %s:", chain))
		addrs := readRpcAddrs(fmt.Sprintf("[3/6] Please input rpc address of %s:", chain))
		contract := ReadEvmAddress(fmt.Sprintf("[4/6] Please input bridge contract address of %s:", chain))
		minConfirms := readMinConfirms(chain, "[5/6]")
		gasLimit := readGasLimit(chain, "[6/6]")

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

func selectChain(msg string) (int, string) {
	var (
		chainIndex int
		err        error
	)

	var s string
	for i, chain := range supportedChain {
		s += fmt.Sprintf("[%d]%s ", i, chain)
	}

	for {
		fmt.Printf("%s, %s:", msg, s)
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

func readChainID(msg string) uint64 {
	var (
		chainID int
		err     error
	)
	for {
		fmt.Printf(msg)
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

func readRpcAddrs(msg string) []string {
	var addrs []string

	for {
		fmt.Printf(msg)
		addrs = append(addrs, readInput())
		fmt.Print("have next rpc address? Y/n:")
		if !ReadYes() {
			break
		}
	}

	return addrs
}

func readMinConfirms(chain, step string) uint64 {
	var (
		minConfirms int
		err         error
	)

	fmt.Printf("%s minimum blocks to confirm of %s, use 30 by default? Y/n:", step, chain)
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

func readGasLimit(chain, step string) uint64 {
	var (
		gasLimit int
		err      error
	)
	fmt.Printf("%s gas limit of %s, use 1500000 by default? Y/n:", step, chain)
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

func ReadEvmAddress(msg string) string {
	for {
		fmt.Print(msg)
		addr := readInput()
		if strings.Compare(common.HexToAddress(addr).String(), addr) != 0 {
			fmt.Println("invalid evm address, please try again")
		} else {
			return addr
		}
	}
}

func ReadYes() bool {
	s := readInput()
	return strings.EqualFold(s, "y") || strings.EqualFold(s, "")
}
