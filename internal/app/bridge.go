package app

import (
	"context"
	"fmt"
	"math/big"
	"sync"
	"time"

	center_chain "github.com/boringdao/bridge/internal/monitor/center"
	"github.com/common-nighthawk/go-figure"

	"github.com/boringdao/bridge/internal/loggers"
	"github.com/boringdao/bridge/internal/monitor"
	"github.com/boringdao/bridge/internal/monitor/chain"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"
	"github.com/sirupsen/logrus"
)

type Bridge struct {
	repo      *repo.Repo
	storage   storage.Storage
	mnts      map[uint64]monitor.Mnt
	center    *center_chain.Monitor
	logger    logrus.FieldLogger
	edgeCocoC chan *monitor.Coco
	mntCocoC  map[uint64]chan *monitor.Coco
	mux       sync.Mutex

	ctx    context.Context
	cancel context.CancelFunc
}

func New(repoRoot *repo.Repo) (*Bridge, error) {
	storagePath := repo.GetStoragePath(repoRoot.Config.RepoRoot, "app")
	boringStorage, err := leveldb.New(storagePath)
	if err != nil {
		return nil, err
	}

	mnts := make(map[uint64]monitor.Mnt)
	mntCocoC := make(map[uint64]chan *monitor.Coco)
	for _, config := range repoRoot.Config.Edges {
		mnt, err := chain.New(repoRoot.Config.RepoRoot, config, loggers.Logger(config.Name))
		if err != nil {
			return nil, err
		}
		mnts[config.ChainID] = mnt
		mntCocoC[config.ChainID] = make(chan *monitor.Coco, 1024)
	}

	center, err := center_chain.New(repoRoot.Config.RepoRoot, repoRoot.Config.Center, loggers.Logger(repoRoot.Config.Center.Name))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Bridge{
		repo:      repoRoot,
		mnts:      mnts,
		center:    center,
		storage:   boringStorage,
		edgeCocoC: make(chan *monitor.Coco, 1024),
		mntCocoC:  mntCocoC,
		logger:    loggers.Logger(loggers.APP),
		ctx:       ctx,
		cancel:    cancel,
	}, nil
}

func (b *Bridge) Start() error {
	for chainID, mnt := range b.mnts {
		mnt := mnt
		chainID := chainID
		go func() {
			if err := mnt.Start(); err != nil {
				b.logger.Errorf("mnt %s start error:%w", mnt.Name(), err)
				return
			}
			b.logger.Infof("mnt %s for chain ID %d has started", mnt.Name(), chainID)
			for coco := range mnt.HandleCocoC() {
				b.edgeCocoC <- coco
			}
		}()
	}
	if err := b.center.Start(); err != nil {
		return err
	}
	b.logger.Infof("mnt %s for chain ID %d has started", b.center.Name(), b.center.ChainId())

	go func() {
		for coco := range b.center.HandleCocoC() {
			b.mntCocoC[coco.ToChainId.Uint64()] <- coco
		}
	}()

	go b.listenEdgeCocoC()
	for chainId, _ := range b.mntCocoC {
		go b.listenCenterCocoC(chainId)
	}
	b.printLogo()

	return nil
}

func (b *Bridge) Stop() error {
	b.cancel()
	return nil
}

func (b *Bridge) listenEdgeCocoC() {
	for {
		select {
		case coco := <-b.edgeCocoC:
			edge := b.mnts[coco.FromChainId.Uint64()]

			b.logger.Infof("========> start handle %s to %s transaction...", edge.Name(), b.center.Name())
			if edge.HasTx(coco.TxId, coco) {
				b.logger.WithField("tx", coco.TxId).Error("has handled the interchain event")
				return
			}
			var err error
			switch coco.Typ {
			case monitor.Deposited:
				err = b.center.Issue(coco.FromToken, coco.ToToken, coco.From, coco.From, coco.FromChainId, big.NewInt(int64(b.center.ChainId())), coco.Amount, fmt.Sprintf("%s#Deposited", coco.TxId))
			case monitor.CrossOuted:
				if coco.ToChainId.Uint64() == b.center.ChainId() {
					err = b.center.CrossIn(coco.FromToken, coco.From, coco.To, coco.FromChainId, coco.ToChainId, coco.Amount, fmt.Sprintf("%s#CrossOuted", coco.TxId))
				} else {
					err = b.center.ForwardCrossOut(coco.FromToken, coco.From, coco.To, coco.FromChainId, coco.ToChainId, coco.Amount, fmt.Sprintf("%s#CrossOuted", coco.TxId))
				}
			}
			if err != nil {
				b.logger.Panic(err)
			}
			edge.PutTxID(coco.TxId, coco)
			b.logger.Infof("========> end handle %s to %s transaction...", edge.Name(), b.center.Name())
		case <-b.ctx.Done():
			close(b.edgeCocoC)
			return
		}
	}
}

func (b *Bridge) listenCenterCocoC(chainId uint64) {
	for {
		select {
		case coco := <-b.mntCocoC[chainId]:
			edge := b.mnts[coco.ToChainId.Uint64()]

			b.logger.Infof("========> start handle %s to %s transaction...", b.center.Name(), edge.Name())
			if b.center.HasTx(coco.TxId, coco) {
				b.logger.WithField("tx", coco.TxId).Error("has handled the interchain event")
				return
			}
			var err error
			switch coco.Typ {
			case monitor.CrossOuted:
				err = edge.CrossIn(coco.FromToken, coco.ToToken, coco.From, coco.To, coco.FromChainId, coco.ToChainId, coco.Amount, fmt.Sprintf("%s#CrossIn", coco.TxId))
			case monitor.ForwardCrossOuted:
				err = edge.CrossIn(coco.FromToken, coco.ToToken, coco.From, coco.To, coco.FromChainId, coco.ToChainId, coco.Amount, fmt.Sprintf("%s#ForwardCrossOuted", coco.TxId))
			case monitor.Withdrawed:
				err = edge.CrossIn(coco.FromToken, coco.ToToken, coco.From, coco.To, coco.FromChainId, coco.ToChainId, coco.Amount, fmt.Sprintf("%s#Withdrawed", coco.TxId))
			}
			if err != nil {
				b.logger.Panic(err)
			}
			b.center.PutTxID(coco.TxId, coco)
			b.logger.Infof("========> end handle %s to %s transaction...", b.center.Name(), edge.Name())
		case <-b.ctx.Done():
			close(b.mntCocoC[chainId])
			return
		}
	}
}

func (b *Bridge) printLogo() {
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println()
		fmt.Println("=======================================================")
		fig := figure.NewColorFigure("TwoWay", "slant", "red", true)
		fig.Print()
		fmt.Println()
		fmt.Println("=======================================================")
		fmt.Println()
		return
	}
}
