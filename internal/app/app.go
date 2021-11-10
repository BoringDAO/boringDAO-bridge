package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/boringdao/bridge/internal/monitor"

	"github.com/boringdao/bridge/internal/loggers"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"
	"github.com/common-nighthawk/go-figure"
	"github.com/sirupsen/logrus"
)

type Bridge struct {
	repo     *repo.Repo
	monitors map[uint64]monitor.IMonitor
	storage  storage.Storage
	logger   logrus.FieldLogger
	mux      sync.Mutex

	ctx    context.Context
	cancel context.CancelFunc
}

func New(repoRoot *repo.Repo) (*Bridge, error) {
	storagePath := repo.GetStoragePath(repoRoot.Config.RepoRoot, "app")
	boringStorage, err := leveldb.New(storagePath)
	if err != nil {
		return nil, err
	}

	monitors := make(map[uint64]monitor.IMonitor)

	for _, bConfig := range repoRoot.Config.Bridges {
		monitor, err := monitor.New(repoRoot.Config.RepoRoot, bConfig, loggers.Logger(bConfig.Name))
		if err != nil {
			return nil, err
		}
		monitors[bConfig.ChainID] = monitor
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Bridge{
		repo:     repoRoot,
		monitors: monitors,
		storage:  boringStorage,
		logger:   loggers.Logger(loggers.APP),
		ctx:      ctx,
		cancel:   cancel,
	}, nil
}

func (b *Bridge) Start() error {
	for chainID, mnt := range b.monitors {
		if err := mnt.Start(); err != nil {
			return err
		}

		b.logger.Infof("monitor for chain ID %d has started", chainID)
	}

	return nil
}

func (b *Bridge) Stop() error {
	b.cancel()
	return nil
}

func (b *Bridge) printLogo() {
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println()
		fmt.Println("=======================================================")
		fig := figure.NewFigure("Bridge", "slant", true)
		fig.Print()
		fmt.Println()
		fmt.Println("=======================================================")
		fmt.Println()
		return
	}
}
