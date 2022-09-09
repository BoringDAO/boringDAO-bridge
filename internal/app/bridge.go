package app

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/boringdao/bridge/internal/loggers"
	"github.com/boringdao/bridge/internal/monitor"
	"github.com/boringdao/bridge/internal/monitor/filter"
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

	tokens := make(map[string]string)
	for tokenAddr, originChainId := range repoRoot.Config.Token {
		tokens[strings.TrimSpace(tokenAddr)] = strings.TrimSpace(originChainId)
	}

	for _, bConfig := range repoRoot.Config.Bridges {
		if bConfig.IsFilter {
			mnt, err := filter.New(repoRoot.Config.RepoRoot, bConfig, tokens, chainIDs, loggers.Logger(bConfig.Name))
			if err != nil {
				return nil, err
			}
			monitors[bConfig.ChainID] = mnt
		} else {
			mnt, err := monitor.New(repoRoot.Config.RepoRoot, bConfig, tokens, chainIDs, loggers.Logger(bConfig.Name))
			if err != nil {
				return nil, err
			}
			monitors[bConfig.ChainID] = mnt
		}

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
			targetMnt, ok := b.monitors[coco.ToChainId.Uint64()]
			if !ok {
				b.logger.WithFields(logrus.Fields{
					"from chain ID": coco.FromChainId.String(),
					"to chain ID":   coco.ToChainId.String(),
					"txId":          coco.TxId,
				}).Error("unexpected to chain ID in cross out coco")
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
			sourceMnt, ok := b.monitors[coco.FromChainId.Uint64()]
			if !ok {
				b.logger.WithFields(logrus.Fields{
					"from chain ID": coco.FromChainId.String(),
					"to chain ID":   coco.ToChainId.String(),
					"txId":          coco.TxId,
				}).Error("unexpected from chain ID in finished coco")
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
