package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/boringdao/bridge/internal/loggers"
	"github.com/boringdao/bridge/internal/monitor/bridge"
	"github.com/boringdao/bridge/internal/monitor/bsc_okex"
	"github.com/boringdao/bridge/internal/monitor/eth"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"
	"github.com/common-nighthawk/go-figure"
	"github.com/sirupsen/logrus"
)

type MntPair struct {
	ethMnt    *eth.Monitor
	bridgeMnt bridge.Monitor
}

type Bridge struct {
	repo     *repo.Repo
	mntPairs map[uint64]*MntPair
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

	mntPairs := make(map[uint64]*MntPair)
	if len(repoRoot.Config.Bsc.Tokens) != 0 {
		ethMnt, err := eth.New(repoRoot.Config, repoRoot.Config.Bsc.ChainID, loggers.Logger(loggers.ETH))
		if err != nil {
			return nil, err
		}
		bscMnt, err := bsc_okex.New(repoRoot.Config.RepoRoot, &repoRoot.Config.Bsc, loggers.Logger(loggers.BSC))
		if err != nil {
			return nil, err
		}
		mntPairs[repoRoot.Config.Bsc.ChainID] = &MntPair{
			ethMnt:    ethMnt,
			bridgeMnt: bscMnt,
		}
	}
	if len(repoRoot.Config.Okex.Tokens) != 0 {
		ethMnt, err := eth.New(repoRoot.Config, repoRoot.Config.Okex.ChainID, loggers.Logger(loggers.ETH))
		if err != nil {
			return nil, err
		}
		okexMnt, err := bsc_okex.New(repoRoot.Config.RepoRoot, &repoRoot.Config.Okex, loggers.Logger(loggers.OKEX))
		if err != nil {
			return nil, err
		}
		mntPairs[repoRoot.Config.Okex.ChainID] = &MntPair{
			ethMnt:    ethMnt,
			bridgeMnt: okexMnt,
		}
	}

	ctx, cancel := context.WithCancel(context.Background())

	return &Bridge{
		repo:     repoRoot,
		mntPairs: mntPairs,
		storage:  boringStorage,
		logger:   loggers.Logger(loggers.APP),
		ctx:      ctx,
		cancel:   cancel,
	}, nil
}

func (b *Bridge) Start() error {
	for chainID, mntPair := range b.mntPairs {
		if err := mntPair.ethMnt.Start(); err != nil {
			return err
		}

		if err := mntPair.bridgeMnt.Start(); err != nil {
			return err
		}

		go b.listenEthCocoC(mntPair)
		go b.listenBridgeCocoC(mntPair)

		b.logger.Infof("mnt pair for chain ID %d has started", chainID)
	}

	return nil
}

func (b *Bridge) Stop() error {
	b.cancel()
	return nil
}

func (b *Bridge) listenEthCocoC(mntPair *MntPair) {
	cocoC := mntPair.ethMnt.HandleCocoC()
	for {
		select {
		case coco := <-cocoC:
			handle := func() {
				b.logger.Infof("========> start handle Chain %d transaction...", mntPair.ethMnt.GetChainID())
				defer b.logger.Infof("========> end handle Chain %d transaction...", mntPair.ethMnt.GetChainID())
				if mntPair.ethMnt.HasTx(coco.TxId) {
					b.logger.WithField("tx", coco.TxId).Error("has handled the interchain event")
					return
				}
				err := mntPair.bridgeMnt.CrossMint(coco.EthToken, coco.TxId, coco.Sender, coco.Recipient, coco.Amount)
				if err != nil {
					b.logger.Panic(err)
				}
				mntPair.ethMnt.PutTxID(coco.TxId, coco)
			}
			handle()
		case <-b.ctx.Done():
			close(cocoC)
			return
		}
	}
}

func (b *Bridge) listenBridgeCocoC(mntPair *MntPair) {
	cocoC := mntPair.bridgeMnt.HandleCocoC()
	for {
		select {
		case coco := <-cocoC:
			handle := func() {
				b.mux.Lock()
				defer b.mux.Unlock()
				b.logger.Infof("========> start handle eth transaction...")
				defer b.logger.Infof("========> end handle eth transaction...")
				if mntPair.bridgeMnt.HasTx(coco.TxId) {
					b.logger.WithField("tx", coco.TxId).Error("has handled the interchain event")
					return
				}

				err := mntPair.ethMnt.UnlockBor(coco.TxId, coco.EthToken, coco.ChainID, coco.Sender, coco.Recipient, coco.Amount)
				if err != nil {
					b.logger.Panic(err)
				}

				mntPair.bridgeMnt.PutTxID(coco.TxId, coco)
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
