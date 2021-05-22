package bsc_okex

import (
	"context"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/hexutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type BridgeWrapper struct {
	addrIdx   int
	config    *repo.BridgeConfig
	ethClient *ethclient.Client
	bridge    *Bridge
	session   *BridgeSession
	logger    logrus.FieldLogger
}

func NewBridgeWrapper(config *repo.BridgeConfig, logger logrus.FieldLogger) (*BridgeWrapper, error) {
	if len(config.Addrs) == 0 {
		return nil, fmt.Errorf("addrs for bsc session wrapper is empty")
	}

	etherCli, err := ethclient.Dial(config.Addrs[0])
	if err != nil {
		return nil, fmt.Errorf("dial bridge chainID %d node: %w", config.ChainID, err)
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
		config:    config,
		ethClient: etherCli,
		bridge:    bridge,
		session:   session,
		logger:    logger,
	}, nil
}

func (bw *BridgeWrapper) HeaderByNumber(ctx context.Context, number *big.Int) *types.Header {
	var header *types.Header
	var err error

	if err := retry.Retry(func(attempt uint) error {
		header, err = bw.ethClient.HeaderByNumber(ctx, number)
		if err != nil {
			bw.logger.Warnf("HeaderByNumber: %s", err.Error())

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

func (bw *BridgeWrapper) FilterCrossBurn(opts *bind.FilterOpts) *BridgeCrossBurnIterator {
	var iterator *BridgeCrossBurnIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = bw.bridge.FilterCrossBurn(opts)
		if err != nil {
			bw.logger.Warnf("FilterCrossBurn: %s", err.Error())

			if bw.isNetworkError(err) {
				bw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		bw.logger.Panic(err)
	}

	return iterator
}

func (bw *BridgeWrapper) TxMinted(txId string) bool {
	var unlocked bool
	var err error

	if err := retry.Retry(func(attempt uint) error {
		unlocked, err = bw.session.TxMinted(txId)
		if err != nil {
			bw.logger.Warnf("TxMinted: %s", err.Error())

			if bw.isNetworkError(err) {
				bw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		bw.logger.Panic(err)
	}

	return unlocked
}

func (bw *BridgeWrapper) SuggestGasPrice(ctx context.Context) *big.Int {
	var result *big.Int
	var err error

	if err := retry.Retry(func(attempt uint) error {
		result, err = bw.ethClient.SuggestGasPrice(ctx)
		if err != nil {
			bw.logger.Warnf("SuggestGasPrice: %s", err.Error())

			if bw.isNetworkError(err) {
				bw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		bw.logger.Panic(err)
	}

	return result
}

func (bw *BridgeWrapper) CrossMint(ethToken common.Address, from common.Address, to common.Address, amount *big.Int, txid string) *types.Transaction {
	var tx *types.Transaction
	var err error

	if err := retry.Retry(func(attempt uint) error {
		price := bw.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		bw.session.TransactOpts.GasPrice = gasPrice.BigInt()

		tx, err = bw.session.CrossMint(ethToken, from, to, amount, txid)
		if err != nil {
			bw.logger.Warnf("CrossMint: %s", err.Error())

			if bw.isNetworkError(err) {
				bw.switchToNextAddr()
				return err
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		bw.logger.Panic(err)
	}

	return tx
}

func (bw *BridgeWrapper) TransactionReceiptsLimitedRetry(ctx context.Context, txHashes []common.Hash) (*types.Receipt, error) {
	var receipt *types.Receipt
	var err error

	if err := retry.Retry(func(attempt uint) error {
		for _, txHash := range txHashes {
			receipt, err = bw.ethClient.TransactionReceipt(ctx, txHash)
			if err != nil {
				bw.logger.Warnf("TransactionReceipt: %s", err.Error())

				if bw.isNetworkError(err) {
					bw.switchToNextAddr()
				}
			}

			if err == nil {
				return nil
			}
		}
		return err
	}, strategy.Wait(10*time.Second), strategy.Limit(30)); err != nil {
		bw.logger.Warnf("retry TransactionReceipt: %s", err.Error())
	}

	return receipt, err
}

func (bw *BridgeWrapper) TransactionReceipt(ctx context.Context, txHash common.Hash) *types.Receipt {
	var receipt *types.Receipt
	var err error

	if err := retry.Retry(func(attempt uint) error {
		receipt, err = bw.ethClient.TransactionReceipt(ctx, txHash)
		if err != nil {
			bw.logger.Warnf("TransactionReceipt: %s", err.Error())

			if bw.isNetworkError(err) {
				bw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		bw.logger.Panic(err)
	}

	return receipt
}

func (bw *BridgeWrapper) switchToNextAddr() {
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

		bw.bridge, err = NewBridge(common.HexToAddress(bw.config.BridgeContract), bw.ethClient)
		if err != nil {
			continue
		}

		bw.session.Contract = bw.bridge

		bw.logger.Infof("switch to %s successfully", bw.config.Addrs[bw.addrIdx])

		return
	}

	panic("all bridge addrs are not valid")
}

func (bw *BridgeWrapper) Close() {
	bw.ethClient.Close()
}

func (bw *BridgeWrapper) isNetworkError(err error) bool {
	if err == nil {
		return false
	}

	return regexp.MustCompile("Post .* EOF").MatchString(err.Error()) ||
		strings.Contains(err.Error(), "connection reset by peer") ||
		strings.Contains(err.Error(), "TLS handshake timeout")
}
