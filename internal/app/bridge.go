package app

import (
	"context"
	"fmt"
	"time"

	"github.com/boringdao/bridge/internal/monitor/bsc"
	"github.com/boringdao/bridge/internal/monitor/eth"

	"github.com/boringdao/bridge/internal/loggers"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"

	"github.com/common-nighthawk/go-figure"
	"github.com/sirupsen/logrus"
)

type Bridge struct {
	repo    *repo.Repo
	ethMnt  *eth.Monitor
	bscMnt  *bsc.Monitor
	storage storage.Storage
	logger  logrus.FieldLogger

	ctx    context.Context
	cancel context.CancelFunc
}

func New(repoRoot *repo.Repo) (*Bridge, error) {
	storagePath := repo.GetStoragePath(repoRoot.Config.RepoRoot, "app")
	boringStorage, err := leveldb.New(storagePath)
	if err != nil {
		return nil, err
	}

	ethMnt, err := eth.New(repoRoot.Config, loggers.Logger(loggers.ETH))
	if err != nil {
		return nil, err
	}

	bscMnt, err := bsc.New(repoRoot.Config, loggers.Logger(loggers.BSC))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Bridge{
		repo:    repoRoot,
		ethMnt:  ethMnt,
		bscMnt:  bscMnt,
		storage: boringStorage,
		logger:  loggers.Logger(loggers.APP),
		ctx:     ctx,
		cancel:  cancel,
	}, nil
}

func (b *Bridge) Start() error {
	err := b.ethMnt.Start()
	if err != nil {
		return err
	}

	err = b.bscMnt.Start()
	if err != nil {
		return err
	}

	go b.listenEthCocoC()

	go b.listenBscCocoC()

	return nil
}

func (b *Bridge) Stop() error {
	b.cancel()
	return nil
}

func (b *Bridge) listenEthCocoC() {
	cocoC := b.ethMnt.HandleCocoC()
	for {
		select {
		case coco := <-cocoC:
			handle := func() {
				b.logger.Infof("========> start handle bsc transaction...")
				defer b.logger.Infof("========> end handle bsc transaction...")
				err := b.bscMnt.CrossMint(coco.Recipient, coco.Amount)
				if err != nil {
					b.logger.Panic(err)
				}
				b.ethMnt.PutTxID(coco.TxId, coco)

			}
			handle()
		case <-b.ctx.Done():
			close(cocoC)
			return
		}
	}
}

func (b *Bridge) listenBscCocoC() {
	cocoC := b.bscMnt.HandleCocoC()
	for {
		select {
		case coco := <-cocoC:
			handle := func() {
				b.logger.Infof("========> start handle eth transaction...")
				defer b.logger.Infof("========> end handle eth transaction...")
				err := b.ethMnt.UnlockBor(b.ethMnt.BorAddr, coco.Recipient, coco.Amount)
				if err != nil {
					b.logger.Panic(err)
				}

				b.bscMnt.PutTxID(coco.TxId, coco)

			}
			handle()
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
