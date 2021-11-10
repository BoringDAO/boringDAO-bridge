package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/boringdao/bridge/internal/app"
	"github.com/boringdao/bridge/internal/loggers"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/hexutil"
	"github.com/boringdao/bridge/pkg/kit/log"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/howeyc/gopass"
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

	ethKey, err := ethKey()
	if err != nil {
		return err
	}
	for _, config := range repo.Config.Bridges {
		config.PrivKey = ethKey
	}

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

func ethKey() (string, error) {
	key, err := gopass.GetPasswdPrompt("Please input eth/bridge private key: ", true, os.Stdin, os.Stdout)
	priv, err := crypto.ToECDSA(hexutil.Decode(string(key)))
	if err != nil || priv == nil {
		return "", fmt.Errorf("eth private key format error:%w", err)
	}

	addr := repo.ReadEvmAddress("Please input the address of your private key:")
	keyAddr := crypto.PubkeyToAddress(priv.PublicKey).String()

	if strings.Compare(addr, keyAddr) != 0 {
		return "", fmt.Errorf("the address cannot match the private key, please check and try again")
	}

	return string(key), nil
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
