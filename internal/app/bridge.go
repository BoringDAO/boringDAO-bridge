package app

import (
	"context"
	"fmt"
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
	repo        *repo.Repo
	storage     storage.Storage
	mnts        map[uint64]monitor.Mnt
	center      *center_chain.Monitor
	logger      logrus.FieldLogger
	edgeCocoC   chan *monitor.Coco
	centerCocoC chan *monitor.Coco
	mux         sync.Mutex

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
	for _, config := range repoRoot.Config.Edges {
		mnt, err := chain.New(repoRoot.Config.RepoRoot, config, loggers.Logger(config.Name))
		if err != nil {
			return nil, err
		}
		mnts[config.ChainID] = mnt
	}

	center, err := center_chain.New(repoRoot.Config.RepoRoot, repoRoot.Config.Center, loggers.Logger(repoRoot.Config.Center.Name))
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Bridge{
		repo:        repoRoot,
		mnts:        mnts,
		center:      center,
		storage:     boringStorage,
		edgeCocoC:   make(chan *monitor.Coco),
		centerCocoC: make(chan *monitor.Coco),
		logger:      loggers.Logger(loggers.APP),
		ctx:         ctx,
		cancel:      cancel,
	}, nil
}

func (b *Bridge) Start() error {
	for chainID, mnt := range b.mnts {
		mnt := mnt
		if err := mnt.Start(); err != nil {
			return err
		}
		b.logger.Infof("mnt %s for chain ID %d has started", mnt.Name(), chainID)
		go func() {
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
			b.centerCocoC <- coco
		}
	}()

	go b.listenEdgeCocoC()
	go b.listenCenterCocoC()

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
			defer b.logger.Infof("========> end handle %s to %s transaction...", edge.Name(), b.center.Name())
			if edge.HasTx(coco.TxId, coco) {
				b.logger.WithField("tx", coco.TxId).Error("has handled the interchain event")
				return
			}
			var err error
			switch coco.Typ {
			case monitor.Deposited:
				err = b.center.Issue(coco.FromToken, coco.ToToken, coco.From, coco.To, coco.FromChainId, coco.ToChainId, coco.Amount, fmt.Sprintf("%s#Deposited", coco.TxId))
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
		case <-b.ctx.Done():
			close(b.edgeCocoC)
			return
		}
	}
}

func (b *Bridge) listenCenterCocoC() {
	for {
		select {
		case coco := <-b.centerCocoC:
			edge := b.mnts[coco.ToChainId.Uint64()]

			b.logger.Infof("========> start handle %s to %s transaction...", b.center.Name(), edge.Name())
			defer b.logger.Infof("========> end handle %s to %s transaction...", b.center.Name(), edge.Name())
			if b.center.HasTx(coco.TxId, coco) {
				b.logger.WithField("tx", coco.TxId).Error("has handled the interchain event")
				return
			}
			var err error
			switch coco.Typ {
			case monitor.CrossOuted:
				err = edge.CrossIn(coco.FromToken, coco.ToToken, coco.From, coco.To, coco.FromChainId, coco.ToChainId, coco.Amount, fmt.Sprintf("%s#CrossIn", coco.TxId))
			case monitor.Withdrawed:
				err = edge.CrossIn(coco.FromToken, coco.ToToken, coco.From, coco.To, coco.FromChainId, coco.ToChainId, coco.Amount, fmt.Sprintf("%s#Withdrawed", coco.TxId))
			}
			if err != nil {
				b.logger.Panic(err)
			}
			b.center.PutTxID(coco.TxId, coco)

		case <-b.ctx.Done():
			close(b.centerCocoC)
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
