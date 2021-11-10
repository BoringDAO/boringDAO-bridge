package monitor

import (
	"context"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sirupsen/logrus"
)

type UsdtWrapper struct {
	addrIdx   int
	config    *repo.StatisticsConfig
	rpcClient *rpc.Client
	ethClient *ethclient.Client
	usdt      *Usdt
	session   *UsdtSession
	logger    logrus.FieldLogger
}

func NewUsdtWrapper(config *repo.StatisticsConfig, logger logrus.FieldLogger) (*UsdtWrapper, error) {
	if len(config.Addrs) == 0 {
		return nil, fmt.Errorf("addrs for bsc session wrapper is empty")
	}

	rpcClient, err := rpc.DialContext(context.Background(), config.Addrs[0])
	if err != nil {
		return nil, fmt.Errorf("dial bridge chainID %d node: %w", config.ChainID, err)
	}

	etherCli := ethclient.NewClient(rpcClient)

	bridge, err := NewUsdt(common.HexToAddress(config.UsdtContract), etherCli)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate a cross contract: %w", err)
	}

	privKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, new(big.Int).SetUint64(config.ChainID))
	if err != nil {
		return nil, fmt.Errorf("failed to NewKeyedTransactorWithChainID %d: %w", config.ChainID, err)
	}

	if config.GasLimit < 800000 {
		auth.GasLimit = 1500000
	} else {
		auth.GasLimit = config.GasLimit
	}

	session := &UsdtSession{
		Contract: bridge,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
		TransactOpts: *auth,
	}

	return &UsdtWrapper{
		addrIdx:   0,
		config:    config,
		rpcClient: rpcClient,
		ethClient: etherCli,
		usdt:      bridge,
		session:   session,
		logger:    logger,
	}, nil
}

func (bw *UsdtWrapper) BlockNumber(ctx context.Context) uint64 {
	var number uint64
	var err error

	if err := retry.Retry(func(attempt uint) error {
		number, err = bw.ethClient.BlockNumber(ctx)
		if err != nil {
			bw.logger.Warnf("BlockNumber: %s", err.Error())

			if bw.isNetworkError(err) {
				bw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		bw.logger.Panic(err)
	}

	return number
}

func (bw *UsdtWrapper) HeaderByNumber(ctx context.Context, num *big.Int) *types.Header {
	var header *types.Header
	var err error

	if err := retry.Retry(func(attempt uint) error {
		header, err = bw.ethClient.HeaderByNumber(ctx, num)
		if err != nil {
			bw.logger.Warnf("HeaderByNumber at block %s: %s", num.String(), err.Error())

			if bw.isNetworkError(err) {
				bw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		bw.logger.Panic(err)
	}

	return header
}

func (bw *UsdtWrapper) FilterTransfer(opts *bind.FilterOpts, swapPair common.Address) (*UsdtTransferIterator, *UsdtTransferIterator) {
	var (
		iteratorAdd *UsdtTransferIterator
		iteratorSub *UsdtTransferIterator
		err         error
	)

	if err := retry.Retry(func(attempt uint) error {
		iteratorAdd, err = bw.usdt.FilterTransfer(opts, nil, []common.Address{swapPair})
		if err != nil {
			bw.logger.Warnf("FilterTransfer to %s: %s", swapPair.String(), err.Error())

			if bw.isNetworkError(err) {
				bw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		bw.logger.Panic(err)
	}

	if err := retry.Retry(func(attempt uint) error {
		iteratorSub, err = bw.usdt.FilterTransfer(opts, []common.Address{swapPair}, nil)
		if err != nil {
			bw.logger.Warnf("FilterTransfer from %s: %s", swapPair.String(), err.Error())

			if bw.isNetworkError(err) {
				bw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		bw.logger.Panic(err)
	}

	return iteratorAdd, iteratorSub
}

func (bw *UsdtWrapper) switchToNextAddr() {
	var err error

	for i := 0; i < len(bw.config.Addrs); i++ {
		bw.addrIdx++
		if bw.addrIdx == len(bw.config.Addrs) {
			bw.addrIdx = 0
		}

		bw.logger.Warnf("try to switch to %s", bw.config.Addrs[bw.addrIdx])

		bw.ethClient, err = ethclient.Dial(bw.config.Addrs[bw.addrIdx])
		if err != nil {
			continue
		}

		bw.usdt, err = NewUsdt(common.HexToAddress(bw.config.UsdtContract), bw.ethClient)
		if err != nil {
			continue
		}

		bw.session.Contract = bw.usdt

		bw.logger.Infof("switch to %s successfully", bw.config.Addrs[bw.addrIdx])

		return
	}

	panic("all bridge addrs are not valid")
}

func (bw *UsdtWrapper) Close() {
	bw.ethClient.Close()
}

func (bw *UsdtWrapper) isNetworkError(err error) bool {
	if err == nil {
		return false
	}

	return regexp.MustCompile("Post .* EOF").MatchString(err.Error()) ||
		strings.Contains(err.Error(), "connection reset by peer") ||
		strings.Contains(err.Error(), "TLS handshake timeout") ||
		strings.Contains(err.Error(), "502")
}
