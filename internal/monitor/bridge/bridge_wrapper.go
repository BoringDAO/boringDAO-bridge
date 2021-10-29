package bridge

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
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type BridgeWrapper struct {
	addrIdx   int
	config    *repo.BridgeConfig
	rpcClient *rpc.Client
	ethClient *ethclient.Client
	bridge    *Bridge
	session   *BridgeSession
	logger    logrus.FieldLogger
}

func NewBridgeWrapper(config *repo.BridgeConfig, logger logrus.FieldLogger) (*BridgeWrapper, error) {
	if len(config.Addrs) == 0 {
		return nil, fmt.Errorf("addrs for bsc session wrapper is empty")
	}

	rpcClient, err := rpc.DialContext(context.Background(), config.Addrs[0])
	if err != nil {
		return nil, fmt.Errorf("dial bridge chainID %d node: %w", config.ChainID, err)
	}

	etherCli := ethclient.NewClient(rpcClient)

	bridge, err := NewBridge(common.HexToAddress(config.BridgeContract), etherCli)
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
		rpcClient: rpcClient,
		ethClient: etherCli,
		bridge:    bridge,
		session:   session,
		logger:    logger,
	}, nil
}

func (bw *BridgeWrapper) BlockNumber(ctx context.Context) uint64 {
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

func (bw *BridgeWrapper) CrossMint(ethToken common.Address, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, common.Hash) {
	var tx *types.Transaction
	var err error
	var hash common.Hash

	if err := retry.Retry(func(attempt uint) error {
		price := bw.SuggestGasPrice(context.TODO())
		gasPrice := decimal.NewFromBigInt(price, 0).Mul(decimal.NewFromFloat(1.2))
		bw.session.TransactOpts.GasPrice = gasPrice.BigInt()

		if bw.config.ChainID == 65 || bw.config.ChainID == 66 {
			tx, hash, err = bw.okChainCrossMint(ethToken, from, to, amount, txid)
		} else {
			tx, err = bw.session.CrossMint(ethToken, from, to, amount, txid)
			if tx != nil {
				hash = tx.Hash()
			}
		}
		if err != nil {
			bw.logger.Warnf("CrossMint: %s", err.Error())
			if errors.Is(err, core.ErrNonceTooLow) {
				tx = nil
				hash = common.Hash{}
				return nil
			}

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

func (bw *BridgeWrapper) okChainCrossMint(token0 common.Address, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, common.Hash, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, common.Hash{}, err
	}
	input, err := parsed.Pack("crossMint", token0, from, to, amount, txid)
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
	rawTx := types.NewTransaction(nonce, common.HexToAddress(bw.config.BridgeContract), value, opts.GasLimit, opts.GasPrice, input)
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

// ensureContext is a helper method to ensure a context is not nil, even if the
// user specified it as such.
func ensureContext(ctx context.Context) context.Context {
	if ctx == nil {
		return context.TODO()
	}
	return ctx
}
