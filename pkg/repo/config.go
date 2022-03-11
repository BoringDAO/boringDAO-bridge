package repo

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const (
	// defaultPathName is the default config dir name
	defaultPathName = ".bridge"
	// defaultPathRoot is the path to the default config dir location.
	defaultPathRoot = "~/" + defaultPathName
	// envDir is the environment variable used to change the path root.
	envDir = "BRIDGE_PATH"
	// Config name
	configName = "bridge.toml"
)

type Config struct {
	RepoRoot string        `json:"repo_root"`
	Title    string        `json:"title"`
	Edges    []*EdgeConfig `json:"edges"`
	KeyFile  string        `json:"keyFile"`
	Center   *CenterConfig `json:"center"`
	Log      `json:"log"`
}

type Log struct {
	Level        string            `toml:"level" json:"level"`
	Dir          string            `toml:"dir" json:"dir"`
	Filename     string            `toml:"filename" json:"filename"`
	ReportCaller bool              `mapstructure:"report_caller" json:"report_caller"`
	Module       map[string]string `toml:"module" json:"module"`
}

type CenterConfig struct {
	Name                    string            `toml:"name" json:"name"`
	Addrs                   []string          `toml:"addrs" json:"addrs"`
	ChainID                 uint64            `toml:"chainID" json:"chainID"`
	MinConfirms             uint64            `toml:"minConfirms" json:"minConfirms"`
	PrivKey                 string            `toml:"privKey" json:"privKey"`
	GasLimit                uint64            `toml:"gasLimit" json:"gasLimit"`
	GasFeeRate              float64           `toml:"gasFeeRate" json:"gasFeeRate"`
	Index                   map[uint64]uint64 `toml:"index" json:"index"`
	WithdrawedHeight        uint64            `mapstructure:"withdrawedHeight" json:"withdrawedHeight"`
	CrossOutedHeight        uint64            `mapstructure:"crossOutedHeight" json:"crossOutedHeight"`
	ForwardCrossOutedHeight uint64            `mapstructure:"forwardCrossOutedHeight" json:"forwardCrossOutedHeight"`
	CenterContract          string            `toml:"centerContract" json:"centerContract"`
}

type EdgeConfig struct {
	IsFilter         bool              `toml:"isFilter" json:"isFilter"`
	Name             string            `toml:"name" json:"name"`
	Addrs            []string          `toml:"addrs" json:"addrs"`
	ChainID          uint64            `toml:"chainID" json:"chainID"`
	MinConfirms      uint64            `toml:"minConfirms" json:"minConfirms"`
	PrivKey          string            `toml:"privKey" json:"privKey"`
	GasLimit         uint64            `toml:"gasLimit" json:"gasLimit"`
	GasFeeRate       float64           `toml:"gasFeeRate" json:"gasFeeRate"`
	Index            map[uint64]uint64 `toml:"index" json:"index"`
	DepositedHeight  uint64            `mapstructure:"depositedHeight" json:"depositedHeight"`
	CrossOutedHeight uint64            `mapstructure:"crossOutedHeight" json:"crossOutedHeight"`
	EdgeContract     string            `toml:"EdgeContract" json:"EdgeContract"`
	Plugin           string            `toml:"plugin" json:"plugin"`
}

func (c *Config) Bytes() ([]byte, error) {
	ret, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

type Port struct {
	PProf int64 `toml:"pprof" json:"pprof"`
}

func DefaultConfig() (*Config, error) {
	return &Config{
		Title: "Boring configuration file",
		Log: Log{
			Level:    "info",
			Dir:      "logs",
			Filename: "bridge.log",
			Module:   make(map[string]string),
		},
	}, nil
}

func UnmarshalConfig(repoRoot string) (*Config, error) {
	viper.SetConfigFile(filepath.Join(repoRoot, configName))
	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("BORING")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	config, err := DefaultConfig()
	if err != nil {
		return nil, err
	}

	if err := viper.Unmarshal(config); err != nil {
		return nil, err
	}

	config.RepoRoot = repoRoot

	return config, nil
}

func ReadConfig(path, configType string, config interface{}) error {
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType(configType)
	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(config); err != nil {
		return err
	}

	return nil
}

func PathRoot() (string, error) {
	dir := os.Getenv(envDir)
	var err error
	if len(dir) == 0 {
		dir, err = homedir.Expand(defaultPathRoot)
	}
	return dir, err
}

func PathRootWithDefault(path string) (string, error) {
	if len(path) == 0 {
		return PathRoot()
	}

	return path, nil
}
