package eth

import (
	"context"
	"errors"
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
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type LockWrapper struct {
	addrIdx   int
	eth       *repo.Eth
	ethClient *ethclient.Client
	lock      *CrossLock
	session   *CrossLockSession
	logger    logrus.FieldLogger
}

func NewLockWrapper(config *repo.Eth, logger logrus.FieldLogger) (*LockWrapper, error) {
	if len(config.Addrs) == 0 {
		return nil, fmt.Errorf("addrs for eth session wrapper is empty")
	}

	etherCli, err := ethclient.Dial(config.Addrs[0])
	if err != nil {
		return nil, fmt.Errorf("dial eth node: %w", err)
	}

	lock, err := NewCrossLock(common.HexToAddress(config.CrossLockContract), etherCli)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate a cross contract: %w", err)
	}

	privKey, err := crypto.ToECDSA(hexutil.Decode(config.PrivKey))
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

	session := &CrossLockSession{
		Contract: lock,
		CallOpts: bind.CallOpts{
			Pending: false,
		},
		TransactOpts: *auth,
	}

	return &LockWrapper{
		addrIdx:   0,
		eth:       config,
		ethClient: etherCli,
		lock:      lock,
		session:   session,
		logger:    logger,
	}, nil
}

func (lw *LockWrapper) HeaderByNumber(ctx context.Context, number *big.Int) *types.Header {
	var header *types.Header
	var err error

	if err := retry.Retry(func(attempt uint) error {
		header, err = lw.ethClient.HeaderByNumber(ctx, number)
		if err != nil {
			lw.logger.Warnf("HeaderByNumber: %s", err.Error())

			if lw.isNetworkError(err) {
				lw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		lw.logger.Panic(err)
	}

	return header
}

func (lw *LockWrapper) FilterLock(opts *bind.FilterOpts) *CrossLockLockIterator {
	var iterator *CrossLockLockIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = lw.lock.FilterLock(opts)
		if err != nil {
			lw.logger.Warnf("FilterLock: %s", err.Error())

			if lw.isNetworkError(err) {
				lw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		lw.logger.Panic(err)
	}

	return iterator
}

func (lw *LockWrapper) TxUnlocked(txId string) bool {
	var result bool
	var err error

	if err := retry.Retry(func(attempt uint) error {
		result, err = lw.session.TxUnlocked(txId)
		if err != nil {
			lw.logger.Warnf("TxUnlocked: %s", err.Error())

			if lw.isNetworkError(err) {
				lw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		lw.logger.Panic(err)
	}

	return result
}

func (lw *LockWrapper) SuggestGasPrice(ctx context.Context) *big.Int {
	var result *big.Int
	var err error

	if err := retry.Retry(func(attempt uint) error {
		result, err = lw.ethClient.SuggestGasPrice(ctx)
		if err != nil {
			lw.logger.Warnf("SuggestGasPrice: %s", err.Error())

			if lw.isNetworkError(err) {
				lw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		lw.logger.Panic(err)
	}

	return result
}

func (lw *LockWrapper) Unlock(token common.Address, chainID *big.Int, from common.Address, to common.Address, amount *big.Int, txid string) *types.Transaction {
	var result *types.Transaction
	var err error

	if err := retry.Retry(func(attempt uint) error {
		price := lw.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		lw.session.TransactOpts.GasPrice = gasPrice.BigInt()

		result, err = lw.session.Unlock(token, chainID, from, to, amount, txid)
		if err != nil {
			lw.logger.Warnf("Unlock: %s", err.Error())
			if errors.Is(err, core.ErrNonceTooLow) {
				result = nil
				return nil
			}

			if lw.isNetworkError(err) {
				lw.switchToNextAddr()
			}

		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		lw.logger.Panic(err)
	}

	return result
}

func (lw *LockWrapper) TransactionReceiptsLimitedRetry(ctx context.Context, txHashes []common.Hash) (*types.Receipt, error) {
	var result *types.Receipt
	var err error

	if err := retry.Retry(func(attempt uint) error {
		for _, txHash := range txHashes {
			result, err = lw.ethClient.TransactionReceipt(ctx, txHash)
			if err != nil {
				lw.logger.Warnf("TransactionReceipt: %s", err.Error())

				if lw.isNetworkError(err) {
					lw.switchToNextAddr()
				}
			}
			if err == nil {
				return nil
			}
		}
		return err
	}, strategy.Wait(10*time.Second), strategy.Limit(30)); err != nil {
		lw.logger.Warnf("retry TransactionReceipt: %s", err.Error())
	}

	return result, err
}

func (lw *LockWrapper) TransactionReceipt(ctx context.Context, txHash common.Hash) *types.Receipt {
	var result *types.Receipt
	var err error

	if err := retry.Retry(func(attempt uint) error {
		result, err = lw.ethClient.TransactionReceipt(ctx, txHash)
		if err != nil {
			lw.logger.Warnf("TransactionReceipt: %s", err.Error())

			if lw.isNetworkError(err) {
				lw.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		lw.logger.Panic(err)
	}

	return result
}

func (lw *LockWrapper) switchToNextAddr() {
	var err error

	for i := 0; i < len(lw.eth.Addrs); i++ {
		lw.addrIdx++
		if lw.addrIdx == len(lw.eth.Addrs) {
			lw.addrIdx = 0
		}

		lw.logger.Warnf("try to switch to %s", lw.eth.Addrs[lw.addrIdx])

		lw.ethClient, err = ethclient.Dial(lw.eth.Addrs[lw.addrIdx])
		if err != nil {
			continue
		}

		lw.lock, err = NewCrossLock(common.HexToAddress(lw.eth.CrossLockContract), lw.ethClient)
		if err != nil {
			continue
		}

		lw.session.Contract = lw.lock

		lw.logger.Infof("switch to %s successfully", lw.eth.Addrs[lw.addrIdx])

		return
	}

	panic("all eth addrs are not valid")
}

func (lw *LockWrapper) Close() {
	lw.ethClient.Close()
}

func (lw *LockWrapper) isNetworkError(err error) bool {
	if err == nil {
		return false
	}

	return regexp.MustCompile("Post .* EOF").MatchString(err.Error()) ||
		strings.Contains(err.Error(), "connection reset by peer") ||
		strings.Contains(err.Error(), "TLS handshake timeout")
}
