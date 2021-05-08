package eth

import (
	"context"
	"fmt"
	"math/big"
	"regexp"

	"github.com/boringdao/bridge/internal/repo"
	"github.com/boringdao/bridge/pkg/kit/hexutil"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

type EthWrapper struct {
	addrIdx   int
	eth       *repo.Eth
	ethClient *ethclient.Client
	lock      *CrossLock
	session   *CrossLockSession
	logger    logrus.FieldLogger
}

func NewEthWrapper(config *repo.Eth, logger logrus.FieldLogger) (*EthWrapper, error) {
	if len(config.Addrs) == 0 {
		return nil, fmt.Errorf("addrs for bsc session wrapper is empty")
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

	return &EthWrapper{
		addrIdx:   0,
		eth:       config,
		ethClient: etherCli,
		lock:      lock,
		session:   session,
		logger:    logger,
	}, nil
}

func (ew *EthWrapper) HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error) {
	for {
		header, err := ew.ethClient.HeaderByNumber(ctx, number)
		if ew.isNetworkError(err) {
			ew.switchToNextAddr()
			continue
		}
		return header, err
	}
}

func (ew *EthWrapper) FilterLock(opts *bind.FilterOpts) (*CrossLockLockIterator, error) {
	for {
		itrator, err := ew.lock.FilterLock(opts)
		if ew.isNetworkError(err) {
			ew.switchToNextAddr()
			continue
		}
		return itrator, err
	}
}

func (ew *EthWrapper) TxUnlocked(txId string) (bool, error) {
	for {
		result, err := ew.session.TxUnlocked(txId)
		if ew.isNetworkError(err) {
			ew.switchToNextAddr()
			continue
		}
		return result, err
	}
}

func (ew *EthWrapper) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	for {
		result, err := ew.ethClient.SuggestGasPrice(ctx)
		if ew.isNetworkError(err) {
			ew.switchToNextAddr()
			continue
		}
		return result, err
	}
}

func (ew *EthWrapper) NonceAt(ctx context.Context, address common.Address) (uint64, error) {
	for {
		result, err := ew.ethClient.NonceAt(ctx, address, nil)
		if ew.isNetworkError(err) {
			ew.switchToNextAddr()
			continue
		}
		return result, err
	}
}

func (ew *EthWrapper) Unlock(token common.Address, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	for {
		result, err := ew.session.Unlock(token, from, to, amount, txid)
		if ew.isNetworkError(err) {
			ew.switchToNextAddr()
			continue
		}
		return result, err
	}
}

func (ew *EthWrapper) TransactionReceipt(ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	for {
		result, err := ew.ethClient.TransactionReceipt(ctx, txHash)
		if ew.isNetworkError(err) {
			ew.switchToNextAddr()
			continue
		}
		return result, err
	}
}

func (ew *EthWrapper) switchToNextAddr() {
	var err error

	for i := 0; i < len(ew.eth.Addrs); i++ {
		ew.addrIdx++
		if ew.addrIdx == len(ew.eth.Addrs) {
			ew.addrIdx = 0
		}

		ew.logger.Warnf("try to switch to %s", ew.eth.Addrs[ew.addrIdx])

		ew.ethClient, err = ethclient.Dial(ew.eth.Addrs[ew.addrIdx])
		if err != nil {
			continue
		}

		ew.lock, err = NewCrossLock(common.HexToAddress(ew.eth.CrossLockContract), ew.ethClient)
		if err != nil {
			continue
		}

		ew.session.Contract = ew.lock

		ew.logger.Infof("switch to %s successfully", ew.eth.Addrs[ew.addrIdx])

		return
	}

	panic("all eth addrs are not valid")
}

func (ew *EthWrapper) Close() {
	ew.ethClient.Close()
}

func (ew *EthWrapper) isNetworkError(err error) bool {
	if err == nil {
		return false
	}

	ew.logger.Infof("check network error for %s", err.Error())

	return regexp.MustCompile("Post .* EOF").MatchString(err.Error())
}
