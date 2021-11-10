package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/boringdao/bridge/internal/app"
	"github.com/boringdao/bridge/internal/loggers"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/log"
	"github.com/urfave/cli"
)

var (
	startCMD = cli.Command{
		Name:   "start",
		Usage:  "Start a long-running daemon process",
		Action: start,
	}
)

func start(ctx *cli.Context) error {
	fmt.Println(getVersion(true))

	repoRoot, err := repo.PathRootWithDefault(ctx.GlobalString("repo"))
	if err != nil {
		return err
	}

	repo, err := repo.Load(repoRoot)
	if err != nil {
		return fmt.Errorf("repo load: %w", err)
	}

	err = log.Initialize(
		log.WithReportCaller(repo.Config.Log.ReportCaller),
		log.WithPersist(true),
		log.WithFilePath(filepath.Join(repoRoot, repo.Config.Log.Dir)),
		log.WithFileName(repo.Config.Log.Filename),
		log.WithMaxSize(20*1024*1024),
		log.WithMaxAge(90*24*time.Hour),
		log.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		return fmt.Errorf("log initialize: %w", err)
	}

	loggers.Initialize(repo.Config)

	bridge, err := app.New(repo)
	if err != nil {
		return fmt.Errorf("boring-node new: %w", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	handleShutdown(bridge, &wg)

	if err := bridge.Start(); err != nil {
		return err
	}

	wg.Wait()

	logger.Info("Bridge exits")
	return nil
}

func handleShutdown(bridge *app.Bridge, wg *sync.WaitGroup) {
	var stop = make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM)
	signal.Notify(stop, syscall.SIGINT)

	go func() {
		<-stop
		fmt.Println("received interrupt signal, shutting down...")
		if err := bridge.Stop(); err != nil {
			logger.Error("bridge node stop: ", err)
		}

		wg.Done()
		os.Exit(0)
	}()
}
