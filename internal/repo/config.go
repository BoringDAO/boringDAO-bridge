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
	RepoRoot string `json:"repo_root"`
	Title    string `json:"title"`
	Eth      `json:"eth"`
	Bsc      `json:"bsc"`
	Log      `json:"log"`
	Token    map[string]string `json:"token" json:"token"`
}

type Log struct {
	Level        string    `toml:"level" json:"level"`
	Dir          string    `toml:"dir" json:"dir"`
	Filename     string    `toml:"filename" json:"filename"`
	ReportCaller bool      `mapstructure:"report_caller" json:"report_caller"`
	Module       LogModule `toml:"module" json:"module"`
}

type LogModule struct {
	BSC string `toml:"bsc" json:"bsc"`
	ETH string `toml:"eth" json:"eth"`
	APP string `toml:"app" json:"app"`
}

type Eth struct {
	Addrs             []string `toml:"addrs" json:"addrs"`
	MinConfirms       uint64   `toml:"minConfirms" json:"minConfirms"`
	PrivKey           string   `toml:"privKey" json:"privKey"`
	GasLimit          uint64   `toml:"gasLimit" json:"gasLimit"`
	Height            uint64   `toml:"height" json:"height"`
	CrossLockContract string   `toml:"crossLockContract" json:"crossLockContract"`
}

type Bsc struct {
	Addrs          []string `toml:"addrs" json:"addrs"`
	MinConfirms    uint64   `toml:"minConfirms" json:"minConfirms"`
	PrivKey        string   `toml:"privKey" json:"privKey"`
	GasLimit       uint64   `toml:"gasLimit" json:"gasLimit"`
	Height         uint64   `toml:"height" json:"height"`
	BridgeContract string   `toml:"bridgeContract" json:"bridgeContract"`
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
			Module: LogModule{
				APP: "info",
				ETH: "info",
			},
		},
		Eth: Eth{
			GasLimit: 1500000,
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
