package chain

import (
	"context"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	mnt "github.com/boringdao/bridge/internal/monitor/contracts"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rlp"

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

type Wrapper struct {
	addrIdx   int
	config    *repo.BridgeConfig
	ethClient *ethclient.Client
	rpcClient *rpc.Client
	pegProxy  *mnt.PegProxy
	session   *mnt.PegProxySession
	logger    logrus.FieldLogger
}

func NewWrapper(config *repo.BridgeConfig, logger logrus.FieldLogger) (*Wrapper, error) {
	if len(config.Addrs) == 0 {
		return nil, fmt.Errorf("addrs for %s session wrapper is empty", config.Name)
	}

	rpcClient, err := rpc.DialContext(context.Background(), config.Addrs[0])
	if err != nil {
		return nil, err
	}
	etherCli := ethclient.NewClient(rpcClient)

	pegProxy, err := mnt.NewPegProxy(common.HexToAddress(config.TwoWayContract), etherCli)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate a TwoWayContract contract: %w", err)
	}

	privKey, err := crypto.ToECDSA(hexutil.Decode(config.PrivKey))
	if err != nil {
		return nil, err
	}

	chainID, err := etherCli.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	if chainID.Uint64() != config.ChainID {
		return nil, fmt.Errorf("%s blockchain chainID is %d but config is %d", config.Name, chainID.Uint64(), config.ChainID)
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

	return &Wrapper{
		addrIdx:   0,
		config:    config,
		rpcClient: rpcClient,
		ethClient: etherCli,
		pegProxy:  pegProxy,
		session:   session,
		logger:    logger,
	}, nil
}

func (w *Wrapper) HeaderByNumber(ctx context.Context, number *big.Int) *types.Header {
	var header *types.Header
	var err error

	if err := retry.Retry(func(attempt uint) error {
		header, err = w.ethClient.HeaderByNumber(ctx, number)
		if err != nil {
			w.logger.Warnf("HeaderByNumber: %s", err.Error())

			if w.isNetworkError(err) {
				w.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		w.logger.Panic(err)
	}

	return header
}

func (w *Wrapper) FilterCrossBurn(opts *bind.FilterOpts) *mnt.PegProxyCrossBurnIterator {
	var iterator *mnt.PegProxyCrossBurnIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = w.pegProxy.FilterCrossBurn(opts)
		if err != nil {
			w.logger.Warnf("FilterCrossBurn: %s", err.Error())

			if w.isNetworkError(err) {
				w.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		w.logger.Panic(err)
	}

	return iterator
}

func (w *Wrapper) FilterRollback(opts *bind.FilterOpts) *mnt.PegProxyRollbackIterator {
	var iterator *mnt.PegProxyRollbackIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = w.pegProxy.FilterRollback(opts)
		if err != nil {
			w.logger.Warnf("FilterRollback: %s", err.Error())

			if w.isNetworkError(err) {
				w.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		w.logger.Panic(err)
	}

	return iterator
}

func (w *Wrapper) FilterLock(opts *bind.FilterOpts) *mnt.PegProxyLockIterator {
	var iterator *mnt.PegProxyLockIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = w.pegProxy.FilterLock(opts)
		if err != nil {
			w.logger.Warnf("FilterLock: %s", err.Error())

			if w.isNetworkError(err) {
				w.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		w.logger.Panic(err)
	}

	return iterator
}

func (w *Wrapper) TxUnlocked(txId string) bool {
	var result bool
	var err error

	if err := retry.Retry(func(attempt uint) error {
		result, err = w.session.TxUnlocked(txId)
		if err != nil {
			w.logger.Warnf("TxUnlocked: %s", err.Error())

			if w.isNetworkError(err) {
				w.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		w.logger.Panic(err)
	}

	return result
}

func (w *Wrapper) TxRollbacked(txId string) bool {
	var result bool
	var err error

	if err := retry.Retry(func(attempt uint) error {
		result, err = w.session.TxRollbacked(txId)
		if err != nil {
			w.logger.Warnf("TxRollbacked: %s", err.Error())

			if w.isNetworkError(err) {
				w.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		w.logger.Panic(err)
	}

	return result
}

func (w *Wrapper) TxMinted(txId string) bool {
	var result bool
	var err error

	if err := retry.Retry(func(attempt uint) error {
		result, err = w.session.TxMinted(txId)
		if err != nil {
			w.logger.Warnf("TxMinted: %s", err.Error())

			if w.isNetworkError(err) {
				w.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		w.logger.Panic(err)
	}

	return result
}

func (w *Wrapper) SuggestGasPrice(ctx context.Context) *big.Int {
	var result *big.Int
	var err error

	if err := retry.Retry(func(attempt uint) error {
		result, err = w.ethClient.SuggestGasPrice(ctx)
		if err != nil {
			w.logger.Warnf("SuggestGasPrice: %s", err.Error())

			if w.isNetworkError(err) {
				w.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		w.logger.Panic(err)
	}

	return result
}

func (w *Wrapper) CrossIn(token, from, to common.Address, chainID, amount *big.Int, txid string) (*types.Transaction, common.Hash) {
	var tx *types.Transaction
	var err error
	var hash common.Hash

	if err := retry.Retry(func(attempt uint) error {
		price := w.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		w.session.TransactOpts.GasPrice = gasPrice.BigInt()

		var newHash common.Hash
		tx, newHash, err = w.crossIn(token, from, to, chainID, amount, txid)
		if tx != nil {
			hash = newHash
		}
		if err != nil {
			w.logger.Warnf("CrossIn: %s", err.Error())

			if w.isNetworkError(err) {
				w.switchToNextAddr()
				return err
			}
		}

		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		w.logger.Panic(err)
	}

	return tx, hash
}

func (w *Wrapper) Rollback(token common.Address, from common.Address, chainID, amount *big.Int, txid string) (*types.Transaction, common.Hash) {
	var tx *types.Transaction
	var err error
	var hash common.Hash

	if err := retry.Retry(func(attempt uint) error {
		price := w.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		w.session.TransactOpts.GasPrice = gasPrice.BigInt()
		var newHash common.Hash
		tx, newHash, err = w.rollback(token, from, chainID, amount, txid)
		if tx != nil {
			hash = newHash
		}
		if err != nil {
			w.logger.Warnf("Rollback: %s", err.Error())

			if w.isNetworkError(err) {
				w.switchToNextAddr()
			}

		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		w.logger.Panic(err)
	}

	return tx, hash
}

func (w *Wrapper) Unlock(token common.Address, from common.Address, to common.Address, chainID, amount *big.Int, txid string) (*types.Transaction, common.Hash) {
	var tx *types.Transaction
	var err error
	var hash common.Hash

	if err := retry.Retry(func(attempt uint) error {
		price := w.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		w.session.TransactOpts.GasPrice = gasPrice.BigInt()
		var newHash common.Hash
		tx, newHash, err = w.unlock(token, from, to, chainID, amount, txid)
		if tx != nil {
			hash = newHash
		}
		if err != nil {
			w.logger.Warnf("Unlock: %s", err.Error())

			if w.isNetworkError(err) {
				w.switchToNextAddr()
				return err
			}
		}

		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		w.logger.Panic(err)
	}

	return tx, hash
}

func (w *Wrapper) unlock(token, from, to common.Address, chainID, amount *big.Int, txid string) (*types.Transaction, common.Hash, error) {
	parsed, err := abi.JSON(strings.NewReader(mnt.PegProxyABI))
	if err != nil {
		return nil, common.Hash{}, err
	}
	input, err := parsed.Pack("unlock", token, chainID, from, to, amount, txid)
	if err != nil {
		return nil, common.Hash{}, err
	}

	opts := w.session.TransactOpts
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = w.ethClient.PendingNonceAt(ensureContext(opts.Context), opts.From)
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}

	// Create the transaction, sign it and schedule it for execution
	rawTx := types.NewTransaction(nonce, common.HexToAddress(w.config.TwoWayContract), value, opts.GasLimit, opts.GasPrice, input)
	signedTx, err := opts.Signer(opts.From, rawTx)
	if err != nil {
		return nil, common.Hash{}, err
	}
	data, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return nil, common.Hash{}, err
	}

	var txHash common.Hash
	err = w.rpcClient.CallContext(ensureContext(opts.Context), &txHash, "eth_sendRawTransaction", hexutil.Encode(data))

	return signedTx, txHash, err
}

func (w *Wrapper) crossIn(token common.Address, from common.Address, to common.Address, chainID, amount *big.Int, txid string) (*types.Transaction, common.Hash, error) {
	parsed, err := abi.JSON(strings.NewReader(mnt.PegProxyABI))
	if err != nil {
		return nil, common.Hash{}, err
	}
	input, err := parsed.Pack("crossIn", token, chainID, from, to, amount, txid)
	if err != nil {
		return nil, common.Hash{}, err
	}

	opts := w.session.TransactOpts
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = w.ethClient.PendingNonceAt(ensureContext(opts.Context), opts.From)
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}

	// Create the transaction, sign it and schedule it for execution
	rawTx := types.NewTransaction(nonce, common.HexToAddress(w.config.TwoWayContract), value, opts.GasLimit, opts.GasPrice, input)
	signedTx, err := opts.Signer(opts.From, rawTx)
	if err != nil {
		return nil, common.Hash{}, err
	}
	data, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return nil, common.Hash{}, err
	}

	var txHash common.Hash
	err = w.rpcClient.CallContext(ensureContext(opts.Context), &txHash, "eth_sendRawTransaction", hexutil.Encode(data))

	return signedTx, txHash, err
}

func (w *Wrapper) rollback(token common.Address, from common.Address, chainID, amount *big.Int, txid string) (*types.Transaction, common.Hash, error) {
	parsed, err := abi.JSON(strings.NewReader(mnt.PegProxyABI))
	if err != nil {
		return nil, common.Hash{}, err
	}
	input, err := parsed.Pack("rollback", token, chainID, from, amount, txid)
	if err != nil {
		return nil, common.Hash{}, err
	}

	opts := w.session.TransactOpts
	value := opts.Value
	if value == nil {
		value = new(big.Int)
	}
	var nonce uint64
	if opts.Nonce == nil {
		nonce, err = w.ethClient.PendingNonceAt(ensureContext(opts.Context), opts.From)
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
	} else {
		nonce = opts.Nonce.Uint64()
	}

	// Create the transaction, sign it and schedule it for execution
	rawTx := types.NewTransaction(nonce, common.HexToAddress(w.config.TwoWayContract), value, opts.GasLimit, opts.GasPrice, input)
	signedTx, err := opts.Signer(opts.From, rawTx)
	if err != nil {
		return nil, common.Hash{}, err
	}
	data, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		return nil, common.Hash{}, err
	}

	var txHash common.Hash
	err = w.rpcClient.CallContext(ensureContext(opts.Context), &txHash, "eth_sendRawTransaction", hexutil.Encode(data))

	return signedTx, txHash, err
}

func (w *Wrapper) TransactionReceiptsLimitedRetry(ctx context.Context, txHashes []common.Hash) (*types.Receipt, error) {
	var result *types.Receipt
	var err error

	if err := retry.Retry(func(attempt uint) error {
		for _, txHash := range txHashes {
			result, err = w.ethClient.TransactionReceipt(ctx, txHash)
			if err != nil {
				w.logger.Warnf("TransactionReceipt: %s, %s", err.Error(), txHash.String())

				if w.isNetworkError(err) {
					w.switchToNextAddr()
				}
			}
			if err == nil {
				return nil
			}
		}
		return err
	}, strategy.Wait(10*time.Second), strategy.Limit(30)); err != nil {
		w.logger.Warnf("retry TransactionReceipt: %s", err.Error())
	}

	return result, err
}

func (w *Wrapper) TransactionReceipt(ctx context.Context, txHash common.Hash) *types.Receipt {
	var result *types.Receipt
	var err error

	if err := retry.Retry(func(attempt uint) error {
		result, err = w.ethClient.TransactionReceipt(ctx, txHash)
		if err != nil {
			w.logger.Warnf("TransactionReceipt: %s", err.Error())

			if w.isNetworkError(err) {
				w.switchToNextAddr()
			}
		}
		return err
	}, strategy.Wait(10*time.Second)); err != nil {
		w.logger.Panic(err)
	}

	return result
}

func (w *Wrapper) switchToNextAddr() {
	var err error
	var rpcClient *rpc.Client
	for i := 0; i < len(w.config.Addrs); i++ {
		w.addrIdx++
		if w.addrIdx == len(w.config.Addrs) {
			w.addrIdx = 0
		}

		w.logger.Warnf("try to switch to %s", w.config.Addrs[w.addrIdx])

		rpcClient, err = rpc.DialContext(context.Background(), w.config.Addrs[w.addrIdx])
		if err != nil {
			continue
		}
		w.ethClient = ethclient.NewClient(rpcClient)

		w.pegProxy, err = mnt.NewPegProxy(common.HexToAddress(w.config.TwoWayContract), w.ethClient)
		if err != nil {
			continue
		}

		w.session.Contract = w.pegProxy

		w.logger.Infof("switch to %s successfully", w.config.Addrs[w.addrIdx])

		return
	}

	panic("all eth addrs are not valid")
}

func (w *Wrapper) Close() {
	w.ethClient.Close()
}

func (w *Wrapper) isNetworkError(err error) bool {
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
