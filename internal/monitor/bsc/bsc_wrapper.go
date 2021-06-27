package bsc

import (
	"context"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	mnt "github.com/boringdao/bridge/internal/monitor"
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

type BscWrapper struct {
	addrIdx   int
	bsc       *repo.Bsc
	ethClient *ethclient.Client
	pegProxy  *mnt.PegProxy
	session   *mnt.PegProxySession
	logger    logrus.FieldLogger
}

func NewBscWrapper(config *repo.Bsc, logger logrus.FieldLogger) (*BscWrapper, error) {
	if len(config.Addrs) == 0 {
		return nil, fmt.Errorf("addrs for bsc session wrapper is empty")
	}

	etherCli, err := ethclient.Dial(config.Addrs[0])
	if err != nil {
		return nil, fmt.Errorf("dial bsc node: %w", err)
	}

	pegProxy, err := mnt.NewPegProxy(common.HexToAddress(config.PegBridgeContract), etherCli)
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

	session := &mnt.PegProxySession{
		Contract: pegProxy,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
		TransactOpts: *auth,
	}

	return &BscWrapper{
		addrIdx:   0,
		bsc:       config,
		ethClient: etherCli,
		pegProxy:  pegProxy,
		session:   session,
		logger:    logger,
	}, nil
}

func (bw *BscWrapper) HeaderByNumber(ctx context.Context, number *big.Int) *types.Header {
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

func (bw *BscWrapper) FilterLock(opts *bind.FilterOpts) *mnt.PegProxyLockIterator {
	var iterator *mnt.PegProxyLockIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = bw.pegProxy.FilterLock(opts)
		if err != nil {
			bw.logger.Warnf("FilterLock: %s", err.Error())

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

func (bw *BscWrapper) FilterCrossBurn(opts *bind.FilterOpts) *mnt.PegProxyCrossBurnIterator {
	var iterator *mnt.PegProxyCrossBurnIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = bw.pegProxy.FilterCrossBurn(opts)
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

func (bw *BscWrapper) FilterRollback(opts *bind.FilterOpts) *mnt.PegProxyRollbackIterator {
	var iterator *mnt.PegProxyRollbackIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = bw.pegProxy.FilterRollback(opts)
		if err != nil {
			bw.logger.Warnf("FilterRollback: %s", err.Error())

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

func (bw *BscWrapper) TxMinted(txId string) bool {
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

func (bw *BscWrapper) SuggestGasPrice(ctx context.Context) *big.Int {
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

func (bw *BscWrapper) CrossIn(token, from, to common.Address, chainId, amount *big.Int, txid string) (*types.Transaction, common.Hash) {
	var tx *types.Transaction
	var err error
	var hash common.Hash

	if err := retry.Retry(func(attempt uint) error {
		price := bw.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		bw.session.TransactOpts.GasPrice = gasPrice.BigInt()

		tx, err = bw.session.CrossIn(token, chainId, from, to, amount, txid)
		if tx != nil {
			hash = tx.Hash()
		}
		if err != nil {
			bw.logger.Warnf("CrossIn: %s", err.Error())

			if bw.isNetworkError(err) {
				bw.switchToNextAddr()
				return err
			}
		}

		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		bw.logger.Panic(err)
	}

	return tx, hash
}

func (bw *BscWrapper) Unlock(token common.Address, from common.Address, to common.Address, chainID, amount *big.Int, txid string) *types.Transaction {
	var result *types.Transaction
	var err error

	if err := retry.Retry(func(attempt uint) error {
		price := bw.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		bw.session.TransactOpts.GasPrice = gasPrice.BigInt()

		result, err = bw.session.Unlock(token, chainID, from, to, amount, txid)
		if err != nil {
			bw.logger.Warnf("Unlock: %s", err.Error())

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

func (bw *BscWrapper) Rollback(token, from common.Address, chainID, amount *big.Int, txid string) (*types.Transaction, common.Hash) {
	var result *types.Transaction
	var err error
	var hash common.Hash

	if err := retry.Retry(func(attempt uint) error {
		price := bw.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		bw.session.TransactOpts.GasPrice = gasPrice.BigInt()

		result, err = bw.session.Rollback(token, chainID, from, amount, txid)
		if result != nil {
			hash = result.Hash()
		}
		if err != nil {
			bw.logger.Warnf("Rollback: %s", err.Error())

			if bw.isNetworkError(err) {
				bw.switchToNextAddr()
			}

		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		bw.logger.Panic(err)
	}

	return result, hash
}

func (bw *BscWrapper) TransactionReceiptsLimitedRetry(ctx context.Context, txHashes []common.Hash) (*types.Receipt, error) {
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

func (bw *BscWrapper) TransactionReceipt(ctx context.Context, txHash common.Hash) *types.Receipt {
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

func (bw *BscWrapper) TxRollbacked(txId string) bool {
	var result bool
	var err error

	if err := retry.Retry(func(attempt uint) error {
		result, err = bw.session.TxRollbacked(txId)
		if err != nil {
			bw.logger.Warnf("TxRollbacked: %s", err.Error())

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

func (bw *BscWrapper) TxUnlocked(txId string) bool {
	var result bool
	var err error

	if err := retry.Retry(func(attempt uint) error {
		result, err = bw.session.TxUnlocked(txId)
		if err != nil {
			bw.logger.Warnf("TxUnlocked: %s", err.Error())

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

func (bw *BscWrapper) switchToNextAddr() {
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

		bw.pegProxy, err = mnt.NewPegProxy(common.HexToAddress(bw.bsc.PegBridgeContract), bw.ethClient)
		if err != nil {
			continue
		}

		bw.session.Contract = bw.pegProxy

		bw.logger.Infof("switch to %s successfully", bw.bsc.Addrs[bw.addrIdx])

		return
	}

	panic("all bridge addrs are not valid")
}

func (bw *BscWrapper) Close() {
	bw.ethClient.Close()
}

func (bw *BscWrapper) isNetworkError(err error) bool {
	if err == nil {
		return false
	}

	return regexp.MustCompile("Post .* EOF").MatchString(err.Error()) ||
		strings.Contains(err.Error(), "connection reset by peer") ||
		strings.Contains(err.Error(), "TLS handshake timeout")
}
