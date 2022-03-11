package repo

import (
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

	return &Repo{
		Config: config,
	}, nil
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
