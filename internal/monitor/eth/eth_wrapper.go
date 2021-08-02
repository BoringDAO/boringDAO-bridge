package eth

import (
	"context"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rlp"

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

type EthWrapper struct {
	addrIdx   int
	eth       *repo.Eth
	ethClient *ethclient.Client
	rpcClient *rpc.Client
	pegProxy  *mnt.PegProxy
	session   *mnt.PegProxySession
	logger    logrus.FieldLogger
}

func NewEthWrapper(config *repo.Eth, logger logrus.FieldLogger) (*EthWrapper, error) {
	if len(config.Addrs) == 0 {
		return nil, fmt.Errorf("addrs for bsc session wrapper is empty")
	}

	rpcClient, err := rpc.DialContext(context.Background(), config.Addrs[0])
	if err != nil {
		return nil, err
	}
	etherCli := ethclient.NewClient(rpcClient)

	pegProxy, err := mnt.NewPegProxy(common.HexToAddress(config.PegBridgeContract), etherCli)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate a PegBridgeContract contract: %w", err)
	}

	privKey, err := crypto.ToECDSA(hexutil.Decode(config.PrivKey))
	if err != nil {
		return nil, err
	}

	chainID, err := etherCli.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		return nil, err
	}
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

	return &EthWrapper{
		addrIdx:   0,
		eth:       config,
		rpcClient: rpcClient,
		ethClient: etherCli,
		pegProxy:  pegProxy,
		session:   session,
		logger:    logger,
	}, nil
}

func (lw *EthWrapper) HeaderByNumber(ctx context.Context, number *big.Int) *types.Header {
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

func (lw *EthWrapper) FilterCrossBurn(opts *bind.FilterOpts) *mnt.PegProxyCrossBurnIterator {
	var iterator *mnt.PegProxyCrossBurnIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = lw.pegProxy.FilterCrossBurn(opts)
		if err != nil {
			lw.logger.Warnf("FilterCrossBurn: %s", err.Error())

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

func (lw *EthWrapper) FilterRollback(opts *bind.FilterOpts) *mnt.PegProxyRollbackIterator {
	var iterator *mnt.PegProxyRollbackIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = lw.pegProxy.FilterRollback(opts)
		if err != nil {
			lw.logger.Warnf("FilterRollback: %s", err.Error())

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

func (lw *EthWrapper) FilterLock(opts *bind.FilterOpts) *mnt.PegProxyLockIterator {
	var iterator *mnt.PegProxyLockIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = lw.pegProxy.FilterLock(opts)
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

func (lw *EthWrapper) TxUnlocked(txId string) bool {
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

func (lw *EthWrapper) TxRollbacked(txId string) bool {
	var result bool
	var err error

	if err := retry.Retry(func(attempt uint) error {
		result, err = lw.session.TxRollbacked(txId)
		if err != nil {
			lw.logger.Warnf("TxRollbacked: %s", err.Error())

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

func (lw *EthWrapper) TxMinted(txId string) bool {
	var result bool
	var err error

	if err := retry.Retry(func(attempt uint) error {
		result, err = lw.session.TxMinted(txId)
		if err != nil {
			lw.logger.Warnf("TxMinted: %s", err.Error())

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

func (lw *EthWrapper) SuggestGasPrice(ctx context.Context) *big.Int {
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

func (lw *EthWrapper) CrossIn(token, from, to common.Address, chainID, amount *big.Int, txid string) (*types.Transaction, common.Hash) {
	var tx *types.Transaction
	var err error
	var hash common.Hash

	if err := retry.Retry(func(attempt uint) error {
		price := lw.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		lw.session.TransactOpts.GasPrice = gasPrice.BigInt()

		var newHash common.Hash
		tx, newHash, err = lw.crossIn(token, from, to, chainID, amount, txid)
		if tx != nil {
			hash = newHash
		}
		if err != nil {
			lw.logger.Warnf("CrossIn: %s", err.Error())

			if lw.isNetworkError(err) {
				lw.switchToNextAddr()
				return err
			}
		}

		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		lw.logger.Panic(err)
	}

	return tx, hash
}

func (lw *EthWrapper) Rollback(token common.Address, from common.Address, chainID, amount *big.Int, txid string) (*types.Transaction, common.Hash) {
	var tx *types.Transaction
	var err error
	var hash common.Hash

	if err := retry.Retry(func(attempt uint) error {
		price := lw.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		lw.session.TransactOpts.GasPrice = gasPrice.BigInt()
		var newHash common.Hash
		tx, newHash, err = lw.rollback(token, from, chainID, amount, txid)
		if tx != nil {
			hash = newHash
		}
		if err != nil {
			lw.logger.Warnf("Rollback: %s", err.Error())

			if lw.isNetworkError(err) {
				lw.switchToNextAddr()
			}

		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		lw.logger.Panic(err)
	}

	return tx, hash
}

func (lw *EthWrapper) Unlock(token common.Address, from common.Address, to common.Address, chainID, amount *big.Int, txid string) (*types.Transaction, common.Hash) {
	var tx *types.Transaction
	var err error
	var hash common.Hash

	if err := retry.Retry(func(attempt uint) error {
		price := lw.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		lw.session.TransactOpts.GasPrice = gasPrice.BigInt()
		var newHash common.Hash
		tx, newHash, err = lw.unlock(token, from, to, chainID, amount, txid)
		if tx != nil {
			hash = newHash
		}
		if err != nil {
			lw.logger.Warnf("Unlock: %s", err.Error())

			if lw.isNetworkError(err) {
				lw.switchToNextAddr()
				return err
			}
		}

		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		lw.logger.Panic(err)
	}

	return tx, hash
}

func (bw *EthWrapper) unlock(token, from, to common.Address, chainID, amount *big.Int, txid string) (*types.Transaction, common.Hash, error) {
	parsed, err := abi.JSON(strings.NewReader(mnt.PegProxyABI))
	if err != nil {
		return nil, common.Hash{}, err
	}
	input, err := parsed.Pack("unlock", token, chainID, from, to, amount, txid)
	if err != nil {
		return nil, common.Hash{}, err
	}

	opts := bw.session.TransactOpts
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = bw.ethClient.PendingNonceAt(ensureContext(opts.Context), opts.From)
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}

	// Create the transaction, sign it and schedule it for execution
	rawTx := types.NewTransaction(nonce, common.HexToAddress(bw.eth.PegBridgeContract), value, opts.GasLimit, opts.GasPrice, input)
	signedTx, err := opts.Signer(opts.From, rawTx)
	if err != nil {
		return nil, common.Hash{}, err
	}
	data, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return nil, common.Hash{}, err
	}

	var txHash common.Hash
	err = bw.rpcClient.CallContext(ensureContext(opts.Context), &txHash, "eth_sendRawTransaction", hexutil.Encode(data))

	return signedTx, txHash, err
}

func (bw *EthWrapper) crossIn(token common.Address, from common.Address, to common.Address, chainID, amount *big.Int, txid string) (*types.Transaction, common.Hash, error) {
	parsed, err := abi.JSON(strings.NewReader(mnt.PegProxyABI))
	if err != nil {
		return nil, common.Hash{}, err
	}
	input, err := parsed.Pack("crossIn", token, chainID, from, to, amount, txid)
	if err != nil {
		return nil, common.Hash{}, err
	}

	opts := bw.session.TransactOpts
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = bw.ethClient.PendingNonceAt(ensureContext(opts.Context), opts.From)
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}

	// Create the transaction, sign it and schedule it for execution
	rawTx := types.NewTransaction(nonce, common.HexToAddress(bw.eth.PegBridgeContract), value, opts.GasLimit, opts.GasPrice, input)
	signedTx, err := opts.Signer(opts.From, rawTx)
	if err != nil {
		return nil, common.Hash{}, err
	}
	data, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return nil, common.Hash{}, err
	}

	var txHash common.Hash
	err = bw.rpcClient.CallContext(ensureContext(opts.Context), &txHash, "eth_sendRawTransaction", hexutil.Encode(data))

	return signedTx, txHash, err
}

func (bw *EthWrapper) rollback(token common.Address, from common.Address, chainID, amount *big.Int, txid string) (*types.Transaction, common.Hash, error) {
	parsed, err := abi.JSON(strings.NewReader(mnt.PegProxyABI))
	if err != nil {
		return nil, common.Hash{}, err
	}
	input, err := parsed.Pack("rollback", token, chainID, from, amount, txid)
	if err != nil {
		return nil, common.Hash{}, err
	}

	opts := bw.session.TransactOpts
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = bw.ethClient.PendingNonceAt(ensureContext(opts.Context), opts.From)
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}

	// Create the transaction, sign it and schedule it for execution
	rawTx := types.NewTransaction(nonce, common.HexToAddress(bw.eth.PegBridgeContract), value, opts.GasLimit, opts.GasPrice, input)
	signedTx, err := opts.Signer(opts.From, rawTx)
	if err != nil {
		return nil, common.Hash{}, err
	}
	data, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return nil, common.Hash{}, err
	}

	var txHash common.Hash
	err = bw.rpcClient.CallContext(ensureContext(opts.Context), &txHash, "eth_sendRawTransaction", hexutil.Encode(data))

	return signedTx, txHash, err
}

func (lw *EthWrapper) TransactionReceiptsLimitedRetry(ctx context.Context, txHashes []common.Hash) (*types.Receipt, error) {
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

func (lw *EthWrapper) TransactionReceipt(ctx context.Context, txHash common.Hash) *types.Receipt {
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

func (lw *EthWrapper) switchToNextAddr() {
	var err error
	var rpcClient *rpc.Client
	for i := 0; i < len(lw.eth.Addrs); i++ {
		lw.addrIdx++
		if lw.addrIdx == len(lw.eth.Addrs) {
			lw.addrIdx = 0
		}

		lw.logger.Warnf("try to switch to %s", lw.eth.Addrs[lw.addrIdx])

		rpcClient, err = rpc.DialContext(context.Background(), lw.eth.Addrs[lw.addrIdx])
		if err != nil {
			continue
		}
		lw.ethClient = ethclient.NewClient(rpcClient)

		lw.pegProxy, err = mnt.NewPegProxy(common.HexToAddress(lw.eth.PegBridgeContract), lw.ethClient)
		if err != nil {
			continue
		}

		lw.session.Contract = lw.pegProxy

		lw.logger.Infof("switch to %s successfully", lw.eth.Addrs[lw.addrIdx])

		return
	}

	panic("all eth addrs are not valid")
}

func (lw *EthWrapper) Close() {
	lw.ethClient.Close()
}

func (lw *EthWrapper) isNetworkError(err error) bool {
	if err == nil {
		return false
	}

	return regexp.MustCompile("Post .* EOF").MatchString(err.Error()) ||
		strings.Contains(err.Error(), "connection reset by peer") ||
		strings.Contains(err.Error(), "TLS handshake timeout")
}

// ensureContext is a helper method to ensure a context is not nil, even if the
// user specified it as such.
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}
