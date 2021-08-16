package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/boringdao/bridge/internal/loggers"
	"github.com/boringdao/bridge/internal/monitor"
	"github.com/boringdao/bridge/internal/monitor/chain"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/storage"
	"github.com/boringdao/bridge/pkg/storage/leveldb"
	"github.com/common-nighthawk/go-figure"
	"github.com/sirupsen/logrus"
)

type Bridge struct {
	repo    *repo.Repo
	storage storage.Storage
	mnts    map[uint64]monitor.Mnt
	logger  logrus.FieldLogger
	cocoC   chan *monitor.Coco
	mux     sync.Mutex

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
	for _, config := range repoRoot.Config.Bridges {
		mnt, err := chain.New(repoRoot.Config.RepoRoot, config, loggers.Logger(config.Name))
		if err != nil {
			return nil, err
		}
		mnts[config.ChainID] = mnt
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Bridge{
		repo:    repoRoot,
		mnts:    mnts,
		storage: boringStorage,
		cocoC:   make(chan *monitor.Coco),
		logger:  loggers.Logger(loggers.APP),
		ctx:     ctx,
		cancel:  cancel,
	}, nil
}

func (b *Bridge) Start() error {
	for chainID, mnt := range b.mnts {
		mnt := mnt
		if err := mnt.Start(); err != nil {
			return err
		}
		b.logger.Infof("mnt  for chain ID %d has started", chainID)
		go func() {
			coco := <-mnt.HandleCocoC()
			b.cocoC <- coco
		}()
	}

	go b.listenCocoC()

	b.printLogo()

	return nil
}

func (b *Bridge) Stop() error {
	b.cancel()
	return nil
}

func (b *Bridge) listenCocoC() {
	for {
		select {
		case coco := <-b.cocoC:
			handle := func() {
				b.mux.Lock()
				defer b.mux.Unlock()

				mnt0 := b.mnts[coco.ChainID0.Uint64()]
				mnt1 := b.mnts[coco.ChainID1.Uint64()]
				b.logger.Infof("========> start handle %s to %s transaction...", mnt0.Name(), mnt1.Name())
				defer b.logger.Infof("========> end handle %s to %s transaction...", mnt0.Name(), mnt1.Name())
				if mnt0.HasTx(coco.TxId, coco) {
					b.logger.WithField("tx", coco.TxId).Error("has handled the interchain event")
					return
				}
				var err error
				switch coco.Typ {
				case monitor.Lock:
					err = mnt1.CrossIn(coco.TxId+"#LOCK", coco.Token1, coco.From, coco.To, coco.ChainID0, coco.Amount)
				case monitor.CrossBurn:
					err = mnt1.Unlock(coco.TxId+"#CrossBurn", coco.Token1, coco.From, coco.To, coco.ChainID0, coco.Amount)
				case monitor.Rollback:
					err = mnt1.Rollback(coco.TxId+"#Rollback", coco.Token1, coco.From, coco.To, coco.ChainID0, coco.Amount)
				}
				if err != nil {
					b.logger.Panic(err)
				}
				mnt0.PutTxID(coco.TxId, coco)

			}
			handle()
		case <-b.ctx.Done():
			close(b.cocoC)
			return
		}
	}
}

func (b *Bridge) printLogo() {
	for {
		time.Sleep(100 * time.Millisecond)
		fmt.Println()
		fmt.Println("=======================================================")
		fig := figure.NewColorFigure("Bridge", "slant", "red", true)
		fig.Print()
		fmt.Println()
		fmt.Println("=======================================================")
		fmt.Println()
		return
	}
}
