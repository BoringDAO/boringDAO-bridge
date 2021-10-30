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
	var chainIDs []uint64
	for _, bConfig := range repoRoot.Config.Bridges {
		chainIDs = append(chainIDs, bConfig.ChainID)
	}

	monitors := make(map[uint64]monitor.IMonitor)

	for _, bConfig := range repoRoot.Config.Bridges {
		monitor, err := monitor.New(repoRoot.Config.RepoRoot, bConfig, repoRoot.Config.Nft, chainIDs, loggers.Logger(bConfig.Name))
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

		go b.listenBridgeCrossOutC(mnt)
		go b.listenBridgeFinishedCocoC(mnt)

		b.logger.Infof("monitor for chain ID %d has started", chainID)
	}

	return nil
}

func (b *Bridge) Stop() error {
	b.cancel()
	return nil
}

func (b *Bridge) listenBridgeCrossOutC(mnt monitor.IMonitor) {
	cocoC := mnt.ListenCrossOutC()
	for {
		select {
		case coco := <-cocoC:
			targetMnt, ok := b.monitors[coco.TargetChainID.Uint64()]
			if !ok {
				b.logger.WithFields(logrus.Fields{
					"source chain ID": coco.SourceChainID,
					"target chain ID": coco.TargetChainID.String(),
					"txId":            coco.TxId,
				}).Error("unexpected target chain ID in cross out coco")
				continue
			}

			targetMnt.HandleCrossIn(coco)

		case <-b.ctx.Done():
			close(cocoC)
			return
		}
	}
}

func (b *Bridge) listenBridgeFinishedCocoC(mnt monitor.IMonitor) {
	cocoC := mnt.ListenFinishedCocoC()
	for {
		select {
		case coco := <-cocoC:
			sourceMnt, ok := b.monitors[coco.SourceChainID]
			if !ok {
				b.logger.WithFields(logrus.Fields{
					"source chain ID": coco.SourceChainID,
					"target chain ID": coco.TargetChainID.String(),
					"txId":            coco.TxId,
				}).Error("unexpected source chain ID in finished coco")
				continue
			}

			sourceMnt.HandleFinishedCoco(coco)

		case <-b.ctx.Done():
			close(cocoC)
			return
		}
	}
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
