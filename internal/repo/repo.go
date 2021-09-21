package repo

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

type Repo struct {
	Config *Config
}

func Load(repoRoot string) (*Repo, error) {
	config, err := UnmarshalConfig(repoRoot)
	if err != nil {
		return nil, err
	}

	if err := checkConfig(config); err != nil {
		return nil, err
	}

	return &Repo{
		Config: config,
	}, nil
}

func checkConfig(config *Config) error {
	if len(config.Nft.Tokens) == 0 {
		return fmt.Errorf("supported original nft token is not configured")
	}

	for _, bConfig := range config.Bridges {
		//if bConfig.Name != "bsc" && bConfig.Name != "okex" && bConfig.Name != "avax" && bConfig.Name != "harmony" {
		//	return fmt.Errorf("unknown blockchain %s, current only bsc, okex and avax are supported", bConfig.Name)
		//}

		if bConfig.BridgeContract == "" {
			return fmt.Errorf("%s bridge contract is not configured", bConfig.Name)
		}

		if len(bConfig.Addrs) == 0 {
			return fmt.Errorf("%s addrs are not configured", bConfig.Name)
		}

		if bConfig.Name == "bsc" && bConfig.MinConfirms < 15 {
			fmt.Println("Warning: bsc minconfirms should be at least 15, please change it if it's in prod environment")
		}
	}

	fmt.Println("Nft configuration:")
	prettyPrint(config.Nft)

	for _, bConfig := range config.Bridges {
		fmt.Println(fmt.Sprintf("%s configuration:", bConfig.Name))
		prettyPrint(bConfig)
	}

	return nil
}

func GetStoragePath(repoRoot string, subPath ...string) string {
	p := filepath.Join(repoRoot, "storage")
	for _, s := range subPath {
		p = filepath.Join(p, s)
	}

	return p
}

func LoadWithNotKey(repoRoot string) (*Repo, error) {
	config, err := UnmarshalConfig(repoRoot)
	if err != nil {
		return nil, err
	}

	return &Repo{
		Config: config,
	}, nil
}

func prettyPrint(input interface{}) {
	data, err := json.MarshalIndent(input, "", "    ")
	if err != nil {
		panic(fmt.Errorf("Fail to marshal: %v", err))
	}

	fmt.Println(string(data))
}
