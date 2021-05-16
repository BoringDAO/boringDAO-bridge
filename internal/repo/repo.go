package repo

import (
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
	if len(config.Token) == 0 {
		return fmt.Errorf("token is not configured")
	}

	if config.Eth.CrossLockContract == "" {
		return fmt.Errorf("eth cross lock contract is not configured")
	}

	if len(config.Eth.Addrs) == 0 {
		return fmt.Errorf("eth addrs are not configured")
	}

	if config.Eth.MinConfirms < 15 {
		return fmt.Errorf("eth minconfirms should be at least 15")
	}

	if config.Bsc.BridgeContract == "" {
		return fmt.Errorf("bsc bridge contract is not configured")
	}

	if len(config.Bsc.Addrs) == 0 {
		return fmt.Errorf("bsc addrs are not configured")
	}

	if config.Bsc.MinConfirms < 15 {
		return fmt.Errorf("bsc minconfirms should be at least 15")
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
