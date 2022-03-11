package monitor

import (
	"fmt"
	"path/filepath"
	"plugin"

	"github.com/boringdao/bridge/pkg/bridge"
	"github.com/boringdao/bridge/pkg/repo"

	"github.com/sirupsen/logrus"
)

//Load order plugin
func New(repoRoot string, config *repo.EdgeConfig, logger logrus.FieldLogger) (bridge.Mnt, error) {
	pluginPath := config.Plugin

	if !filepath.IsAbs(pluginPath) {
		pluginPath = filepath.Join(repoRoot, pluginPath)
	}

	p, err := plugin.Open(pluginPath)
	if err != nil {
		return nil, fmt.Errorf("plugin open: %s", err)
	}

	m, err := p.Lookup("New")
	if err != nil {
		return nil, fmt.Errorf("plugin lookup: %s", err)
	}

	NewMnt, ok := m.(func(repoRoot string, config *repo.EdgeConfig, logger logrus.FieldLogger) (bridge.Mnt, error))
	if !ok {
		return nil, fmt.Errorf("assert New error")
	}
	return NewMnt(repoRoot, config, logger)
}
