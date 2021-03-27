package bsc

import (
	"context"
	"fmt"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/hexutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math/big"
	"regexp"
)

type BridgeWrapper struct {
	addrIdx   int
	bsc       *repo.Bsc
	ethClient *ethclient.Client
	bridge    *Bridge
	session   *BridgeSession
	logger    logrus.FieldLogger
}

func NewBridgeWrapper(config *repo.Bsc, logger logrus.FieldLogger) (*BridgeWrapper, error) {
	if len(config.Addrs) == 0 {
		return nil, fmt.Errorf("addrs for bsc session wrapper is empty")
	}

	etherCli, err := ethclient.Dial(config.Addrs[0])
	if err != nil {
		return nil, fmt.Errorf("dial bsc node: %w", err)
	}

	bridge, err := NewBridge(common.HexToAddress(config.BridgeContract), etherCli)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate a cross contract: %w", err)
	}

	privKey, err := crypto.ToECDSA(hexutil.Decode(config.PrivKey))
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privKey)
	if config.GasLimit < 800000 {
		auth.GasLimit = 1500000
	} else {
		auth.GasLimit = config.GasLimit
	}

	session := &BridgeSession{
		Contract: bridge,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
		TransactOpts: *auth,
	}

	return &BridgeWrapper{
		addrIdx:   0,
		bsc:       config,
		ethClient: etherCli,
		bridge:    bridge,
		session:   session,
		logger:    logger,
	}, nil
}

func (bw *BridgeWrapper) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	for {
		header, err := bw.ethClient.HeaderByNumber(ctx, number)
		if bw.isNetworkError(err) {
			bw.switchToNextAddr()
			continue
		}
		return header, err
	}
}

func (bw *BridgeWrapper) FilterCrossBurn(opts *bind.FilterOpts) (*BridgeCrossBurnIterator, error) {
	for {
		itrator, err := bw.bridge.FilterCrossBurn(opts)
		if bw.isNetworkError(err) {
			bw.switchToNextAddr()
			continue
		}
		return itrator, err
	}
}

func (bw *BridgeWrapper) TxMinted(txId string) (bool, error) {
	for {
		result, err := bw.session.TxMinted(txId)
		if bw.isNetworkError(err) {
			bw.switchToNextAddr()
			continue
		}
		return result, err
	}
}

func (bw *BridgeWrapper) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	for {
		result, err := bw.ethClient.SuggestGasPrice(ctx)
		if bw.isNetworkError(err) {
			bw.switchToNextAddr()
			continue
		}
		return result, err
	}
}

func (bw *BridgeWrapper) CrossMint(from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	for {
		result, err := bw.session.CrossMint(from, to, amount, txid)
		if err != nil && bw.isNetworkError(err) {
			bw.switchToNextAddr()
			continue
		}
		return result, err
	}
}

func (bw *BridgeWrapper) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	for {
		result, err := bw.ethClient.TransactionReceipt(ctx, txHash)
		if bw.isNetworkError(err) {
			bw.switchToNextAddr()
			continue
		}
		return result, err
	}
}

func (bw *BridgeWrapper) switchToNextAddr() {
	var err error

	for i := 0; i < len(bw.bsc.Addrs); i++ {
		bw.addrIdx++
		if bw.addrIdx == len(bw.bsc.Addrs) {
			bw.addrIdx = 0
		}

		bw.logger.Warnf("try to switch to %s", bw.bsc.Addrs[bw.addrIdx])

		bw.ethClient, err = ethclient.Dial(bw.bsc.Addrs[bw.addrIdx])
		if err != nil {
			continue
		}

		bw.bridge, err = NewBridge(common.HexToAddress(bw.bsc.BridgeContract), bw.ethClient)
		if err != nil {
			continue
		}

		bw.session.Contract = bw.bridge

		bw.logger.Infof("switch to %s successfully", bw.bsc.Addrs[bw.addrIdx])

		return
	}

	panic("all bsc addrs are not valid")
}

func (bw *BridgeWrapper) Close() {
	bw.ethClient.Close()
}

func (bw *BridgeWrapper) isNetworkError(err error) bool {
	if err == nil {
		return false
	}

	bw.logger.Infof("check network error for %s", err.Error())

	return regexp.MustCompile("Post .* EOF").MatchString(err.Error())
}
