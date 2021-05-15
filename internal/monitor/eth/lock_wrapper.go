package eth

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
	"strings"
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
		return nil, fmt.Errorf("dial bsc node: %w", err)
	}

	lock, err := NewCrossLock(common.HexToAddress(config.CrossLockContract), etherCli)
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

func (lw *LockWrapper) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	for {
		header, err := lw.ethClient.HeaderByNumber(ctx, number)
		if lw.isNetworkError(err) {
			lw.switchToNextAddr()
			continue
		}
		return header, err
	}
}

func (lw *LockWrapper) FilterLock(opts *bind.FilterOpts) (*CrossLockLockIterator, error) {
	for {
		itrator, err := lw.lock.FilterLock(opts)
		if lw.isNetworkError(err) {
			lw.switchToNextAddr()
			continue
		}
		return itrator, err
	}
}

func (lw *LockWrapper) TxUnlocked(txId string) (bool, error) {
	for {
		result, err := lw.session.TxUnlocked(txId)
		if lw.isNetworkError(err) {
			lw.switchToNextAddr()
			continue
		}
		return result, err
	}
}

func (lw *LockWrapper) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	for {
		result, err := lw.ethClient.SuggestGasPrice(ctx)
		if lw.isNetworkError(err) {
			lw.switchToNextAddr()
			continue
		}
		return result, err
	}
}

func (lw *LockWrapper) Unlock(token common.Address, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	for {
		result, err := lw.session.Unlock(token, from, to, amount, txid)
		if lw.isNetworkError(err) {
			lw.switchToNextAddr()
			continue
		}
		return result, err
	}
}

func (lw *LockWrapper) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	for {
		result, err := lw.ethClient.TransactionReceipt(ctx, txHash)
		if lw.isNetworkError(err) {
			lw.switchToNextAddr()
			continue
		}
		return result, err
	}
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

	lw.logger.Infof("check network error for %s", err.Error())

	return regexp.MustCompile("Post .* EOF").MatchString(err.Error()) ||
		strings.Contains(err.Error(), "connection reset by peer")
}
