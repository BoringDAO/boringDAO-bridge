package repo

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"

	"github.com/boringdao/bridge/pkg/kit/fileutil"
	"github.com/gobuffalo/packd"
	"github.com/gobuffalo/packr"
)

const (
	packPath = "../../config"
)

func Initialize(repoRoot string) error {
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

func Initialized(repoRoot string) bool {
	return fileutil.Exist(filepath.Join(repoRoot, configName))
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

func readInput() string {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	return input.Text()
}
