package center_chain

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/boringdao/bridge/pkg/repo"

	"github.com/boringdao/bridge/internal/monitor/contracts/center"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/strategy"
	"github.com/boringdao/bridge/pkg/kit/hexutil"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/sirupsen/logrus"
)

type Wrapper struct {
	addrIdx   int
	config    *repo.CenterConfig
	ethClient *ethclient.Client
	rpcClient *rpc.Client
	twoWay    *center.TwoWayCenter
	session   *center.TwoWayCenterSession
	logger    logrus.FieldLogger
}

func NewWrapper(config *repo.CenterConfig, logger logrus.FieldLogger) (*Wrapper, error) {
	if len(config.Addrs) == 0 {
		return nil, fmt.Errorf("addrs for %s session wrapper is empty", config.Name)
	}

	rpcClient, err := rpc.DialContext(context.Background(), config.Addrs[0])
	if err != nil {
		return nil, err
	}
	etherCli := ethclient.NewClient(rpcClient)

	twoWay, err := center.NewTwoWayCenter(common.HexToAddress(config.CenterContract), etherCli)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate a TwoWayCenter contract: %w", err)
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

	session := &center.TwoWayCenterSession{
		Contract: twoWay,
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
		twoWay:    twoWay,
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

func (w *Wrapper) FilterWithdrawed(opts *bind.FilterOpts) *center.TwoWayCenterWithdrawedIterator {
	var iterator *center.TwoWayCenterWithdrawedIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = w.twoWay.FilterWithdrawed(opts)
		if err != nil {
			w.logger.Warnf("FilterWithdrawed: %s", err.Error())

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

func (w *Wrapper) FilterCenterCrossOuted(opts *bind.FilterOpts) *center.TwoWayCenterCrossOutedIterator {
	var iterator *center.TwoWayCenterCrossOutedIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = w.twoWay.FilterCrossOuted(opts)

		if err != nil {
			w.logger.Warnf("FilterCenterCrossOuted: %s", err.Error())

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

func (w *Wrapper) FilterForwardCrossOuted(opts *bind.FilterOpts) *center.TwoWayCenterForwardCrossOutedIterator {
	var iterator *center.TwoWayCenterForwardCrossOutedIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = w.twoWay.FilterForwardCrossOuted(opts)

		if err != nil {
			w.logger.Warnf("FilterForwardCrossOuted: %s", err.Error())

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

func (w *Wrapper) FilterCrossInFailed(opts *bind.FilterOpts) *center.TwoWayCenterCrossInFailedIterator {
	var iterator *center.TwoWayCenterCrossInFailedIterator
	var err error

	if err := retry.Retry(func(attempt uint) error {
		iterator, err = w.twoWay.FilterCrossInFailed(opts)
		if err != nil {
			w.logger.Warnf("FilterCrossInFailed: %s", err.Error())

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

func (w *Wrapper) TxHandled(txId string) bool {
	var result bool
	var err error

	if err := retry.Retry(func(attempt uint) error {
		result, err = w.session.TxHandled(txId)
		if err != nil {
			w.logger.Warnf("TxHandled: %s", err.Error())

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

func (w *Wrapper) ForwardCrossOut(fromToken common.Address, from, to common.Address, fromChainID, toChainID, amount *big.Int, txid string) (*types.Transaction, common.Hash) {
	var tx *types.Transaction
	var err error
	var hash common.Hash

	if err := retry.Retry(func(attempt uint) error {
		//price := w.SuggestGasPrice(context.TODO())
		//gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		//w.session.TransactOpts.GasPrice = gasPrice.BigInt()

		var newHash common.Hash
		tx, newHash, err = w.forwardCrossOut(fromToken, from, to, fromChainID, toChainID, amount, txid)
		if tx != nil {
			hash = newHash
		}
		if err != nil {
			w.logger.Warnf("CrossIn: %s", err.Error())
			if errors.Is(err, core.ErrNonceTooLow) {
				tx = nil
				hash = common.Hash{}
				return nil
			}

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

func (w *Wrapper) forwardCrossOut(fromToken common.Address, from, to common.Address, fromChainID, toChainID, amount *big.Int, txid string) (*types.Transaction, common.Hash, error) {
	parsed, err := abi.JSON(strings.NewReader(center.TwoWayCenterMetaData.ABI))
	if err != nil {
		return nil, common.Hash{}, err
	}
	out := center.OutParam{
		FromChainId: fromChainID,
		FromToken:   fromToken,
		From:        from,
		ToChainId:   toChainID,
		To:          to,
		Amount:      amount,
	}
	input, err := parsed.Pack("forwardCrossOut", out, txid)
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
		confirmNonce, err := w.ethClient.NonceAt(ensureContext(opts.Context), opts.From, nil)
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
		if confirmNonce > opts.Nonce.Uint64() {
			nonce, err = w.ethClient.PendingNonceAt(ensureContext(opts.Context), opts.From)
			if err != nil {
				return nil, common.Hash{}, fmt.Errorf("failed to retrieve account nonce: %v", err)
			}
		} else {
			nonce = opts.Nonce.Uint64()
		}
	}

	// Create the transaction, sign it and schedule it for execution
	rawTx := types.NewTransaction(nonce, common.HexToAddress(w.config.CenterContract), value, opts.GasLimit, opts.GasPrice, input)
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

func (w *Wrapper) CrossIn(fromToken common.Address, from, to common.Address, fromChainID, toChainID, amount *big.Int, txid string) (*types.Transaction, common.Hash) {
	var tx *types.Transaction
	var err error
	var hash common.Hash

	if err := retry.Retry(func(attempt uint) error {
		//price := w.SuggestGasPrice(context.TODO())
		//gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		//w.session.TransactOpts.GasPrice = gasPrice.BigInt()

		var newHash common.Hash
		tx, newHash, err = w.crossIn(fromToken, from, to, fromChainID, toChainID, amount, txid)
		if tx != nil {
			hash = newHash
		}
		if err != nil {
			w.logger.Warnf("CrossIn: %s", err.Error())
			if errors.Is(err, core.ErrNonceTooLow) {
				tx = nil
				hash = common.Hash{}
				return nil
			}

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

func (w *Wrapper) crossIn(fromToken common.Address, from, to common.Address, fromChainID, toChainID, amount *big.Int, txid string) (*types.Transaction, common.Hash, error) {
	parsed, err := abi.JSON(strings.NewReader(center.TwoWayCenterMetaData.ABI))
	if err != nil {
		return nil, common.Hash{}, err
	}
	out := center.OutParam{
		FromChainId: fromChainID,
		FromToken:   fromToken,
		From:        from,
		ToChainId:   toChainID,
		To:          to,
		Amount:      amount,
	}
	input, err := parsed.Pack("crossIn", out, txid)
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
		confirmNonce, err := w.ethClient.NonceAt(ensureContext(opts.Context), opts.From, nil)
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
		if confirmNonce > opts.Nonce.Uint64() {
			nonce, err = w.ethClient.PendingNonceAt(ensureContext(opts.Context), opts.From)
			if err != nil {
				return nil, common.Hash{}, fmt.Errorf("failed to retrieve account nonce: %v", err)
			}
		} else {
			nonce = opts.Nonce.Uint64()
		}
	}

	// Create the transaction, sign it and schedule it for execution
	rawTx := types.NewTransaction(nonce, common.HexToAddress(w.config.CenterContract), value, opts.GasLimit, opts.GasPrice, input)
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

func (w *Wrapper) Issue(fromToken, toToken common.Address, from, to common.Address, fromChainID, toChainID, amount *big.Int, txid string) (*types.Transaction, common.Hash) {
	var tx *types.Transaction
	var err error
	var hash common.Hash

	if err := retry.Retry(func(attempt uint) error {
		//price := w.SuggestGasPrice(context.TODO())
		//gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		//w.session.TransactOpts.GasPrice = gasPrice.BigInt()

		var newHash common.Hash
		tx, newHash, err = w.issue(fromToken, toToken, from, to, fromChainID, toChainID, amount, txid)
		if tx != nil {
			hash = newHash
		}
		if err != nil {
			w.logger.Warnf("Issue: %s", err.Error())
			if errors.Is(err, core.ErrNonceTooLow) {
				tx = nil
				hash = common.Hash{}
				return nil
			}

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

func (w *Wrapper) issue(fromToken, toToken common.Address, from, to common.Address, fromChainID, toChainID, amount *big.Int, txid string) (*types.Transaction, common.Hash, error) {
	parsed, err := abi.JSON(strings.NewReader(center.TwoWayCenterMetaData.ABI))
	if err != nil {
		return nil, common.Hash{}, err
	}
	in := center.InParam{
		FromChainId: fromChainID,
		FromToken:   fromToken,
		From:        from,
		ToChainId:   toChainID,
		ToToken:     toToken,
		To:          to,
		Amount:      amount,
	}
	input, err := parsed.Pack("issue", in, txid)
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
		confirmNonce, err := w.ethClient.NonceAt(ensureContext(opts.Context), opts.From, nil)
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("failed to retrieve account nonce: %v", err)
		}
		if confirmNonce > opts.Nonce.Uint64() {
			nonce, err = w.ethClient.PendingNonceAt(ensureContext(opts.Context), opts.From)
			if err != nil {
				return nil, common.Hash{}, fmt.Errorf("failed to retrieve account nonce: %v", err)
			}
		} else {
			nonce = opts.Nonce.Uint64()
		}
	}

	// Create the transaction, sign it and schedule it for execution
	rawTx := types.NewTransaction(nonce, common.HexToAddress(w.config.CenterContract), value, opts.GasLimit, opts.GasPrice, input)
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

func (w *Wrapper) RollbackCrossIn(fromToken, toToken common.Address, from, to common.Address, fromChainID, toChainID, amount *big.Int, txid string) (*types.Transaction, common.Hash) {
	var tx *types.Transaction
	var err error
	var hash common.Hash

	if err := retry.Retry(func(attempt uint) error {
		//price := w.SuggestGasPrice(context.TODO())
		//gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		//w.session.TransactOpts.GasPrice = gasPrice.BigInt()

		var newHash common.Hash
		tx, newHash, err = w.rollbackCrossIn(fromToken, toToken, from, to, fromChainID, toChainID, amount, txid)
		if tx != nil {
			hash = newHash
		}
		if err != nil {
			w.logger.Warnf("CrossIn: %s", err.Error())
			if errors.Is(err, core.ErrNonceTooLow) {
				tx = nil
				hash = common.Hash{}
				return nil
			}

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

func (w *Wrapper) rollbackCrossIn(fromToken, toToken common.Address, from, to common.Address, fromChainID, toChainID, amount *big.Int, txid string) (*types.Transaction, common.Hash, error) {
	parsed, err := abi.JSON(strings.NewReader(center.TwoWayCenterMetaData.ABI))
	if err != nil {
		return nil, common.Hash{}, err
	}
	in := center.InParam{
		FromChainId: fromChainID,
		FromToken:   fromToken,
		From:        from,
		ToChainId:   toChainID,
		ToToken:     toToken,
		To:          to,
		Amount:      amount,
	}
	input, err := parsed.Pack("rollbackCrossIn", in, txid)
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
	rawTx := types.NewTransaction(nonce, common.HexToAddress(w.config.CenterContract), value, opts.GasLimit, opts.GasPrice, input)
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
		if len(txHashes) > 2 {
			w.switchToNextAddr()
		}
		return err
	}, strategy.Wait(10*time.Second), strategy.Limit(9)); err != nil {
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
		w.rpcClient = rpcClient
		w.twoWay, err = center.NewTwoWayCenter(common.HexToAddress(w.config.CenterContract), w.ethClient)
		if err != nil {
			continue
		}

		w.session.Contract = w.twoWay

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
		strings.Contains(err.Error(), "TLS handshake timeout") ||
		strings.Contains(err.Error(), "too many requests") ||
		strings.Contains(err.Error(), "Internal error")
}

// ensureContext is a helper method to ensure a context is not nil, even if the
// user specified it as such.
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}
