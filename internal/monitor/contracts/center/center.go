// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package center

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// InParam is an auto generated low-level Go binding around an user-defined struct.
type InParam struct {
	FromChainId *big.Int
	FromToken   common.Address
	From        []byte
	ToChainId   *big.Int
	ToToken     common.Address
	To          []byte
	Amount      *big.Int
}

// OutParam is an auto generated low-level Go binding around an user-defined struct.
type OutParam struct {
	FromChainId *big.Int
	FromToken   common.Address
	From        []byte
	ToChainId   *big.Int
	To          []byte
	Amount      *big.Int
}

// CenterMetaData contains all meta data concerning the Center contract.
var CenterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addToken\",\"inputs\":[{\"name\":\"_centerToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_edgeChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_edgeToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"calculateFee\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"oToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"fixAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ratioAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"remainAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"calculateReward\",\"inputs\":[{\"name\":\"oToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"_reward\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"crossIn\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"internalType\":\"structOutParam\",\"components\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"txid\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"crossOut\",\"inputs\":[{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"decimalDiff\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"eventHeights0\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eventHeights1\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eventIndex0\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eventIndex1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeTo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeToTreasuryRatio\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fixFees\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"forwardCrossOut\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"internalType\":\"structOutParam\",\"components\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"txid\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getMsgSender\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleKey\",\"inputs\":[{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_feeTo\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_treasuryTo\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isCoin\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isInWhilelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"issue\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"internalType\":\"structInParam\",\"components\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"txid\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockBalances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"locked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ratioFeesHigh\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ratioFeesLow\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ratioFeesMedium\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"remainHigh\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"remainLow\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeToken\",\"inputs\":[{\"name\":\"_centerToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_edgeChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_edgeToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resetStakingReward\",\"inputs\":[{\"name\":\"_oToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"rewardRatio\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setFee\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_fixFees\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_ratioFeesHigh\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_ratioFeesMedium\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_ratioFeesLow\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_remains\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeTo\",\"inputs\":[{\"name\":\"_feeTo\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeToTreasuryRatio\",\"inputs\":[{\"name\":\"_oToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ratio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsCoin\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isCoin\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRewardRatio\",\"inputs\":[{\"name\":\"_oToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_ratio\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setStakingRewards\",\"inputs\":[{\"name\":\"oTokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_srs\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThreshold\",\"inputs\":[{\"name\":\"token0\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setTreasuryTo\",\"inputs\":[{\"name\":\"_treasuryTo\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWhitelist\",\"inputs\":[{\"name\":\"oToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isInWhitelist\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sr\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIStakingReward\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"srs\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"threshold\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"toCenterToken\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"toEdgeToken\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"treasuryTo\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"txHandled\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"oToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"CoinSeted\",\"inputs\":[{\"name\":\"coin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CrossInFailed\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInParam\",\"components\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CrossIned\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInParam\",\"components\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CrossOuted\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInParam\",\"components\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ForwardCrossOutFailed\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInParam\",\"components\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ForwardCrossOuted\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInParam\",\"components\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Issued\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInParam\",\"components\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalVoted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"to\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"StakingRewardReseted\",\"inputs\":[{\"name\":\"oToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Supported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ThresholdChanged\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"oldThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdrawed\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInParam\",\"components\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawedToCenter\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInParam\",\"components\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// CenterABI is the input ABI used to generate the binding from.
// Deprecated: Use CenterMetaData.ABI instead.
var CenterABI = CenterMetaData.ABI

// Center is an auto generated Go binding around an Ethereum contract.
type Center struct {
	CenterCaller     // Read-only binding to the contract
	CenterTransactor // Write-only binding to the contract
	CenterFilterer   // Log filterer for contract events
}

// CenterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CenterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CenterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CenterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CenterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CenterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CenterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CenterSession struct {
	Contract     *Center           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CenterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CenterCallerSession struct {
	Contract *CenterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CenterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CenterTransactorSession struct {
	Contract     *CenterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CenterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CenterRaw struct {
	Contract *Center // Generic contract binding to access the raw methods on
}

// CenterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CenterCallerRaw struct {
	Contract *CenterCaller // Generic read-only contract binding to access the raw methods on
}

// CenterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CenterTransactorRaw struct {
	Contract *CenterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCenter creates a new instance of Center, bound to a specific deployed contract.
func NewCenter(address common.Address, backend bind.ContractBackend) (*Center, error) {
	contract, err := bindCenter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Center{CenterCaller: CenterCaller{contract: contract}, CenterTransactor: CenterTransactor{contract: contract}, CenterFilterer: CenterFilterer{contract: contract}}, nil
}

// NewCenterCaller creates a new read-only instance of Center, bound to a specific deployed contract.
func NewCenterCaller(address common.Address, caller bind.ContractCaller) (*CenterCaller, error) {
	contract, err := bindCenter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CenterCaller{contract: contract}, nil
}

// NewCenterTransactor creates a new write-only instance of Center, bound to a specific deployed contract.
func NewCenterTransactor(address common.Address, transactor bind.ContractTransactor) (*CenterTransactor, error) {
	contract, err := bindCenter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CenterTransactor{contract: contract}, nil
}

// NewCenterFilterer creates a new log filterer instance of Center, bound to a specific deployed contract.
func NewCenterFilterer(address common.Address, filterer bind.ContractFilterer) (*CenterFilterer, error) {
	contract, err := bindCenter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CenterFilterer{contract: contract}, nil
}

// bindCenter binds a generic wrapper to an already deployed contract.
func bindCenter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CenterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Center *CenterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Center.Contract.CenterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Center *CenterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Center.Contract.CenterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Center *CenterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Center.Contract.CenterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Center *CenterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Center.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Center *CenterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Center.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Center *CenterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Center.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Center *CenterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Center *CenterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Center.Contract.DEFAULTADMINROLE(&_Center.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Center *CenterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Center.Contract.DEFAULTADMINROLE(&_Center.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Center *CenterCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Center *CenterSession) UPGRADERROLE() ([32]byte, error) {
	return _Center.Contract.UPGRADERROLE(&_Center.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Center *CenterCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _Center.Contract.UPGRADERROLE(&_Center.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Center *CenterCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Center *CenterSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Center.Contract.UPGRADEINTERFACEVERSION(&_Center.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Center *CenterCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Center.Contract.UPGRADEINTERFACEVERSION(&_Center.CallOpts)
}

// CalculateFee is a free data retrieval call binding the contract method 0xe0e193a2.
//
// Solidity: function calculateFee(address from, address oToken, uint256 toChainId, uint256 amount) view returns(uint256 fixAmount, uint256 ratioAmount, uint256 remainAmount)
func (_Center *CenterCaller) CalculateFee(opts *bind.CallOpts, from common.Address, oToken common.Address, toChainId *big.Int, amount *big.Int) (struct {
	FixAmount    *big.Int
	RatioAmount  *big.Int
	RemainAmount *big.Int
}, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "calculateFee", from, oToken, toChainId, amount)

	outstruct := new(struct {
		FixAmount    *big.Int
		RatioAmount  *big.Int
		RemainAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FixAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RatioAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RemainAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateFee is a free data retrieval call binding the contract method 0xe0e193a2.
//
// Solidity: function calculateFee(address from, address oToken, uint256 toChainId, uint256 amount) view returns(uint256 fixAmount, uint256 ratioAmount, uint256 remainAmount)
func (_Center *CenterSession) CalculateFee(from common.Address, oToken common.Address, toChainId *big.Int, amount *big.Int) (struct {
	FixAmount    *big.Int
	RatioAmount  *big.Int
	RemainAmount *big.Int
}, error) {
	return _Center.Contract.CalculateFee(&_Center.CallOpts, from, oToken, toChainId, amount)
}

// CalculateFee is a free data retrieval call binding the contract method 0xe0e193a2.
//
// Solidity: function calculateFee(address from, address oToken, uint256 toChainId, uint256 amount) view returns(uint256 fixAmount, uint256 ratioAmount, uint256 remainAmount)
func (_Center *CenterCallerSession) CalculateFee(from common.Address, oToken common.Address, toChainId *big.Int, amount *big.Int) (struct {
	FixAmount    *big.Int
	RatioAmount  *big.Int
	RemainAmount *big.Int
}, error) {
	return _Center.Contract.CalculateFee(&_Center.CallOpts, from, oToken, toChainId, amount)
}

// CalculateReward is a free data retrieval call binding the contract method 0x2c32c360.
//
// Solidity: function calculateReward(address oToken, uint256 fromChainId, uint256 amount) view returns(uint256 _reward)
func (_Center *CenterCaller) CalculateReward(opts *bind.CallOpts, oToken common.Address, fromChainId *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "calculateReward", oToken, fromChainId, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateReward is a free data retrieval call binding the contract method 0x2c32c360.
//
// Solidity: function calculateReward(address oToken, uint256 fromChainId, uint256 amount) view returns(uint256 _reward)
func (_Center *CenterSession) CalculateReward(oToken common.Address, fromChainId *big.Int, amount *big.Int) (*big.Int, error) {
	return _Center.Contract.CalculateReward(&_Center.CallOpts, oToken, fromChainId, amount)
}

// CalculateReward is a free data retrieval call binding the contract method 0x2c32c360.
//
// Solidity: function calculateReward(address oToken, uint256 fromChainId, uint256 amount) view returns(uint256 _reward)
func (_Center *CenterCallerSession) CalculateReward(oToken common.Address, fromChainId *big.Int, amount *big.Int) (*big.Int, error) {
	return _Center.Contract.CalculateReward(&_Center.CallOpts, oToken, fromChainId, amount)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Center *CenterCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Center *CenterSession) ChainId() (*big.Int, error) {
	return _Center.Contract.ChainId(&_Center.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Center *CenterCallerSession) ChainId() (*big.Int, error) {
	return _Center.Contract.ChainId(&_Center.CallOpts)
}

// DecimalDiff is a free data retrieval call binding the contract method 0xbf814720.
//
// Solidity: function decimalDiff(address ) view returns(uint256)
func (_Center *CenterCaller) DecimalDiff(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "decimalDiff", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalDiff is a free data retrieval call binding the contract method 0xbf814720.
//
// Solidity: function decimalDiff(address ) view returns(uint256)
func (_Center *CenterSession) DecimalDiff(arg0 common.Address) (*big.Int, error) {
	return _Center.Contract.DecimalDiff(&_Center.CallOpts, arg0)
}

// DecimalDiff is a free data retrieval call binding the contract method 0xbf814720.
//
// Solidity: function decimalDiff(address ) view returns(uint256)
func (_Center *CenterCallerSession) DecimalDiff(arg0 common.Address) (*big.Int, error) {
	return _Center.Contract.DecimalDiff(&_Center.CallOpts, arg0)
}

// EventHeights0 is a free data retrieval call binding the contract method 0xfc1aa909.
//
// Solidity: function eventHeights0(uint256 , uint256 ) view returns(uint256)
func (_Center *CenterCaller) EventHeights0(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "eventHeights0", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventHeights0 is a free data retrieval call binding the contract method 0xfc1aa909.
//
// Solidity: function eventHeights0(uint256 , uint256 ) view returns(uint256)
func (_Center *CenterSession) EventHeights0(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Center.Contract.EventHeights0(&_Center.CallOpts, arg0, arg1)
}

// EventHeights0 is a free data retrieval call binding the contract method 0xfc1aa909.
//
// Solidity: function eventHeights0(uint256 , uint256 ) view returns(uint256)
func (_Center *CenterCallerSession) EventHeights0(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Center.Contract.EventHeights0(&_Center.CallOpts, arg0, arg1)
}

// EventHeights1 is a free data retrieval call binding the contract method 0x05e82506.
//
// Solidity: function eventHeights1(uint256 ) view returns(uint256)
func (_Center *CenterCaller) EventHeights1(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "eventHeights1", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventHeights1 is a free data retrieval call binding the contract method 0x05e82506.
//
// Solidity: function eventHeights1(uint256 ) view returns(uint256)
func (_Center *CenterSession) EventHeights1(arg0 *big.Int) (*big.Int, error) {
	return _Center.Contract.EventHeights1(&_Center.CallOpts, arg0)
}

// EventHeights1 is a free data retrieval call binding the contract method 0x05e82506.
//
// Solidity: function eventHeights1(uint256 ) view returns(uint256)
func (_Center *CenterCallerSession) EventHeights1(arg0 *big.Int) (*big.Int, error) {
	return _Center.Contract.EventHeights1(&_Center.CallOpts, arg0)
}

// EventIndex0 is a free data retrieval call binding the contract method 0x0826820e.
//
// Solidity: function eventIndex0(uint256 ) view returns(uint256)
func (_Center *CenterCaller) EventIndex0(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "eventIndex0", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventIndex0 is a free data retrieval call binding the contract method 0x0826820e.
//
// Solidity: function eventIndex0(uint256 ) view returns(uint256)
func (_Center *CenterSession) EventIndex0(arg0 *big.Int) (*big.Int, error) {
	return _Center.Contract.EventIndex0(&_Center.CallOpts, arg0)
}

// EventIndex0 is a free data retrieval call binding the contract method 0x0826820e.
//
// Solidity: function eventIndex0(uint256 ) view returns(uint256)
func (_Center *CenterCallerSession) EventIndex0(arg0 *big.Int) (*big.Int, error) {
	return _Center.Contract.EventIndex0(&_Center.CallOpts, arg0)
}

// EventIndex1 is a free data retrieval call binding the contract method 0x15cd2c70.
//
// Solidity: function eventIndex1() view returns(uint256)
func (_Center *CenterCaller) EventIndex1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "eventIndex1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventIndex1 is a free data retrieval call binding the contract method 0x15cd2c70.
//
// Solidity: function eventIndex1() view returns(uint256)
func (_Center *CenterSession) EventIndex1() (*big.Int, error) {
	return _Center.Contract.EventIndex1(&_Center.CallOpts)
}

// EventIndex1 is a free data retrieval call binding the contract method 0x15cd2c70.
//
// Solidity: function eventIndex1() view returns(uint256)
func (_Center *CenterCallerSession) EventIndex1() (*big.Int, error) {
	return _Center.Contract.EventIndex1(&_Center.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Center *CenterCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Center *CenterSession) FeeTo() (common.Address, error) {
	return _Center.Contract.FeeTo(&_Center.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_Center *CenterCallerSession) FeeTo() (common.Address, error) {
	return _Center.Contract.FeeTo(&_Center.CallOpts)
}

// FeeToTreasuryRatio is a free data retrieval call binding the contract method 0x59ded2d7.
//
// Solidity: function feeToTreasuryRatio(address ) view returns(uint256)
func (_Center *CenterCaller) FeeToTreasuryRatio(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "feeToTreasuryRatio", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeToTreasuryRatio is a free data retrieval call binding the contract method 0x59ded2d7.
//
// Solidity: function feeToTreasuryRatio(address ) view returns(uint256)
func (_Center *CenterSession) FeeToTreasuryRatio(arg0 common.Address) (*big.Int, error) {
	return _Center.Contract.FeeToTreasuryRatio(&_Center.CallOpts, arg0)
}

// FeeToTreasuryRatio is a free data retrieval call binding the contract method 0x59ded2d7.
//
// Solidity: function feeToTreasuryRatio(address ) view returns(uint256)
func (_Center *CenterCallerSession) FeeToTreasuryRatio(arg0 common.Address) (*big.Int, error) {
	return _Center.Contract.FeeToTreasuryRatio(&_Center.CallOpts, arg0)
}

// FixFees is a free data retrieval call binding the contract method 0x75d47f99.
//
// Solidity: function fixFees(address , uint256 ) view returns(uint256)
func (_Center *CenterCaller) FixFees(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "fixFees", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FixFees is a free data retrieval call binding the contract method 0x75d47f99.
//
// Solidity: function fixFees(address , uint256 ) view returns(uint256)
func (_Center *CenterSession) FixFees(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Center.Contract.FixFees(&_Center.CallOpts, arg0, arg1)
}

// FixFees is a free data retrieval call binding the contract method 0x75d47f99.
//
// Solidity: function fixFees(address , uint256 ) view returns(uint256)
func (_Center *CenterCallerSession) FixFees(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Center.Contract.FixFees(&_Center.CallOpts, arg0, arg1)
}

// GetMsgSender is a free data retrieval call binding the contract method 0x7a6ce2e1.
//
// Solidity: function getMsgSender() view returns(address, address)
func (_Center *CenterCaller) GetMsgSender(opts *bind.CallOpts) (common.Address, common.Address, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "getMsgSender")

	if err != nil {
		return *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetMsgSender is a free data retrieval call binding the contract method 0x7a6ce2e1.
//
// Solidity: function getMsgSender() view returns(address, address)
func (_Center *CenterSession) GetMsgSender() (common.Address, common.Address, error) {
	return _Center.Contract.GetMsgSender(&_Center.CallOpts)
}

// GetMsgSender is a free data retrieval call binding the contract method 0x7a6ce2e1.
//
// Solidity: function getMsgSender() view returns(address, address)
func (_Center *CenterCallerSession) GetMsgSender() (common.Address, common.Address, error) {
	return _Center.Contract.GetMsgSender(&_Center.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Center *CenterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Center *CenterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Center.Contract.GetRoleAdmin(&_Center.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Center *CenterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Center.Contract.GetRoleAdmin(&_Center.CallOpts, role)
}

// GetRoleKey is a free data retrieval call binding the contract method 0x3ab796e7.
//
// Solidity: function getRoleKey(address toToken) pure returns(bytes32 key)
func (_Center *CenterCaller) GetRoleKey(opts *bind.CallOpts, toToken common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "getRoleKey", toToken)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleKey is a free data retrieval call binding the contract method 0x3ab796e7.
//
// Solidity: function getRoleKey(address toToken) pure returns(bytes32 key)
func (_Center *CenterSession) GetRoleKey(toToken common.Address) ([32]byte, error) {
	return _Center.Contract.GetRoleKey(&_Center.CallOpts, toToken)
}

// GetRoleKey is a free data retrieval call binding the contract method 0x3ab796e7.
//
// Solidity: function getRoleKey(address toToken) pure returns(bytes32 key)
func (_Center *CenterCallerSession) GetRoleKey(toToken common.Address) ([32]byte, error) {
	return _Center.Contract.GetRoleKey(&_Center.CallOpts, toToken)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Center *CenterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Center *CenterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Center.Contract.HasRole(&_Center.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Center *CenterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Center.Contract.HasRole(&_Center.CallOpts, role, account)
}

// IsCoin is a free data retrieval call binding the contract method 0xa611033e.
//
// Solidity: function isCoin(address ) view returns(bool)
func (_Center *CenterCaller) IsCoin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "isCoin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCoin is a free data retrieval call binding the contract method 0xa611033e.
//
// Solidity: function isCoin(address ) view returns(bool)
func (_Center *CenterSession) IsCoin(arg0 common.Address) (bool, error) {
	return _Center.Contract.IsCoin(&_Center.CallOpts, arg0)
}

// IsCoin is a free data retrieval call binding the contract method 0xa611033e.
//
// Solidity: function isCoin(address ) view returns(bool)
func (_Center *CenterCallerSession) IsCoin(arg0 common.Address) (bool, error) {
	return _Center.Contract.IsCoin(&_Center.CallOpts, arg0)
}

// IsInWhilelist is a free data retrieval call binding the contract method 0xa07b32c0.
//
// Solidity: function isInWhilelist(address , address ) view returns(bool)
func (_Center *CenterCaller) IsInWhilelist(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "isInWhilelist", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInWhilelist is a free data retrieval call binding the contract method 0xa07b32c0.
//
// Solidity: function isInWhilelist(address , address ) view returns(bool)
func (_Center *CenterSession) IsInWhilelist(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Center.Contract.IsInWhilelist(&_Center.CallOpts, arg0, arg1)
}

// IsInWhilelist is a free data retrieval call binding the contract method 0xa07b32c0.
//
// Solidity: function isInWhilelist(address , address ) view returns(bool)
func (_Center *CenterCallerSession) IsInWhilelist(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Center.Contract.IsInWhilelist(&_Center.CallOpts, arg0, arg1)
}

// LockBalances is a free data retrieval call binding the contract method 0xa6ba8fc1.
//
// Solidity: function lockBalances(address , uint256 ) view returns(uint256)
func (_Center *CenterCaller) LockBalances(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "lockBalances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockBalances is a free data retrieval call binding the contract method 0xa6ba8fc1.
//
// Solidity: function lockBalances(address , uint256 ) view returns(uint256)
func (_Center *CenterSession) LockBalances(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Center.Contract.LockBalances(&_Center.CallOpts, arg0, arg1)
}

// LockBalances is a free data retrieval call binding the contract method 0xa6ba8fc1.
//
// Solidity: function lockBalances(address , uint256 ) view returns(uint256)
func (_Center *CenterCallerSession) LockBalances(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Center.Contract.LockBalances(&_Center.CallOpts, arg0, arg1)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Center *CenterCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "locked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Center *CenterSession) Locked() (bool, error) {
	return _Center.Contract.Locked(&_Center.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Center *CenterCallerSession) Locked() (bool, error) {
	return _Center.Contract.Locked(&_Center.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Center *CenterCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Center *CenterSession) ProxiableUUID() ([32]byte, error) {
	return _Center.Contract.ProxiableUUID(&_Center.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Center *CenterCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Center.Contract.ProxiableUUID(&_Center.CallOpts)
}

// RatioFeesHigh is a free data retrieval call binding the contract method 0xee5f7d96.
//
// Solidity: function ratioFeesHigh(address , uint256 ) view returns(uint256)
func (_Center *CenterCaller) RatioFeesHigh(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "ratioFeesHigh", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RatioFeesHigh is a free data retrieval call binding the contract method 0xee5f7d96.
//
// Solidity: function ratioFeesHigh(address , uint256 ) view returns(uint256)
func (_Center *CenterSession) RatioFeesHigh(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Center.Contract.RatioFeesHigh(&_Center.CallOpts, arg0, arg1)
}

// RatioFeesHigh is a free data retrieval call binding the contract method 0xee5f7d96.
//
// Solidity: function ratioFeesHigh(address , uint256 ) view returns(uint256)
func (_Center *CenterCallerSession) RatioFeesHigh(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Center.Contract.RatioFeesHigh(&_Center.CallOpts, arg0, arg1)
}

// RatioFeesLow is a free data retrieval call binding the contract method 0xe4732500.
//
// Solidity: function ratioFeesLow(address , uint256 ) view returns(uint256)
func (_Center *CenterCaller) RatioFeesLow(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "ratioFeesLow", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RatioFeesLow is a free data retrieval call binding the contract method 0xe4732500.
//
// Solidity: function ratioFeesLow(address , uint256 ) view returns(uint256)
func (_Center *CenterSession) RatioFeesLow(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Center.Contract.RatioFeesLow(&_Center.CallOpts, arg0, arg1)
}

// RatioFeesLow is a free data retrieval call binding the contract method 0xe4732500.
//
// Solidity: function ratioFeesLow(address , uint256 ) view returns(uint256)
func (_Center *CenterCallerSession) RatioFeesLow(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Center.Contract.RatioFeesLow(&_Center.CallOpts, arg0, arg1)
}

// RatioFeesMedium is a free data retrieval call binding the contract method 0x3ff9bb66.
//
// Solidity: function ratioFeesMedium(address , uint256 ) view returns(uint256)
func (_Center *CenterCaller) RatioFeesMedium(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "ratioFeesMedium", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RatioFeesMedium is a free data retrieval call binding the contract method 0x3ff9bb66.
//
// Solidity: function ratioFeesMedium(address , uint256 ) view returns(uint256)
func (_Center *CenterSession) RatioFeesMedium(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Center.Contract.RatioFeesMedium(&_Center.CallOpts, arg0, arg1)
}

// RatioFeesMedium is a free data retrieval call binding the contract method 0x3ff9bb66.
//
// Solidity: function ratioFeesMedium(address , uint256 ) view returns(uint256)
func (_Center *CenterCallerSession) RatioFeesMedium(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Center.Contract.RatioFeesMedium(&_Center.CallOpts, arg0, arg1)
}

// RemainHigh is a free data retrieval call binding the contract method 0x260ca4d5.
//
// Solidity: function remainHigh(address ) view returns(uint256)
func (_Center *CenterCaller) RemainHigh(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "remainHigh", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainHigh is a free data retrieval call binding the contract method 0x260ca4d5.
//
// Solidity: function remainHigh(address ) view returns(uint256)
func (_Center *CenterSession) RemainHigh(arg0 common.Address) (*big.Int, error) {
	return _Center.Contract.RemainHigh(&_Center.CallOpts, arg0)
}

// RemainHigh is a free data retrieval call binding the contract method 0x260ca4d5.
//
// Solidity: function remainHigh(address ) view returns(uint256)
func (_Center *CenterCallerSession) RemainHigh(arg0 common.Address) (*big.Int, error) {
	return _Center.Contract.RemainHigh(&_Center.CallOpts, arg0)
}

// RemainLow is a free data retrieval call binding the contract method 0x4297171f.
//
// Solidity: function remainLow(address ) view returns(uint256)
func (_Center *CenterCaller) RemainLow(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "remainLow", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainLow is a free data retrieval call binding the contract method 0x4297171f.
//
// Solidity: function remainLow(address ) view returns(uint256)
func (_Center *CenterSession) RemainLow(arg0 common.Address) (*big.Int, error) {
	return _Center.Contract.RemainLow(&_Center.CallOpts, arg0)
}

// RemainLow is a free data retrieval call binding the contract method 0x4297171f.
//
// Solidity: function remainLow(address ) view returns(uint256)
func (_Center *CenterCallerSession) RemainLow(arg0 common.Address) (*big.Int, error) {
	return _Center.Contract.RemainLow(&_Center.CallOpts, arg0)
}

// RewardRatio is a free data retrieval call binding the contract method 0xced2a7ba.
//
// Solidity: function rewardRatio(address ) view returns(uint256)
func (_Center *CenterCaller) RewardRatio(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "rewardRatio", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRatio is a free data retrieval call binding the contract method 0xced2a7ba.
//
// Solidity: function rewardRatio(address ) view returns(uint256)
func (_Center *CenterSession) RewardRatio(arg0 common.Address) (*big.Int, error) {
	return _Center.Contract.RewardRatio(&_Center.CallOpts, arg0)
}

// RewardRatio is a free data retrieval call binding the contract method 0xced2a7ba.
//
// Solidity: function rewardRatio(address ) view returns(uint256)
func (_Center *CenterCallerSession) RewardRatio(arg0 common.Address) (*big.Int, error) {
	return _Center.Contract.RewardRatio(&_Center.CallOpts, arg0)
}

// Sr is a free data retrieval call binding the contract method 0xdf791d80.
//
// Solidity: function sr() view returns(address)
func (_Center *CenterCaller) Sr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "sr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sr is a free data retrieval call binding the contract method 0xdf791d80.
//
// Solidity: function sr() view returns(address)
func (_Center *CenterSession) Sr() (common.Address, error) {
	return _Center.Contract.Sr(&_Center.CallOpts)
}

// Sr is a free data retrieval call binding the contract method 0xdf791d80.
//
// Solidity: function sr() view returns(address)
func (_Center *CenterCallerSession) Sr() (common.Address, error) {
	return _Center.Contract.Sr(&_Center.CallOpts)
}

// Srs is a free data retrieval call binding the contract method 0xf3669d80.
//
// Solidity: function srs(address ) view returns(address)
func (_Center *CenterCaller) Srs(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "srs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Srs is a free data retrieval call binding the contract method 0xf3669d80.
//
// Solidity: function srs(address ) view returns(address)
func (_Center *CenterSession) Srs(arg0 common.Address) (common.Address, error) {
	return _Center.Contract.Srs(&_Center.CallOpts, arg0)
}

// Srs is a free data retrieval call binding the contract method 0xf3669d80.
//
// Solidity: function srs(address ) view returns(address)
func (_Center *CenterCallerSession) Srs(arg0 common.Address) (common.Address, error) {
	return _Center.Contract.Srs(&_Center.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Center *CenterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Center *CenterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Center.Contract.SupportsInterface(&_Center.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Center *CenterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Center.Contract.SupportsInterface(&_Center.CallOpts, interfaceId)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_Center *CenterCaller) Threshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "threshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_Center *CenterSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _Center.Contract.Threshold(&_Center.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_Center *CenterCallerSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _Center.Contract.Threshold(&_Center.CallOpts, arg0)
}

// ToCenterToken is a free data retrieval call binding the contract method 0xbafd2d19.
//
// Solidity: function toCenterToken(uint256 , address ) view returns(address)
func (_Center *CenterCaller) ToCenterToken(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "toCenterToken", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ToCenterToken is a free data retrieval call binding the contract method 0xbafd2d19.
//
// Solidity: function toCenterToken(uint256 , address ) view returns(address)
func (_Center *CenterSession) ToCenterToken(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _Center.Contract.ToCenterToken(&_Center.CallOpts, arg0, arg1)
}

// ToCenterToken is a free data retrieval call binding the contract method 0xbafd2d19.
//
// Solidity: function toCenterToken(uint256 , address ) view returns(address)
func (_Center *CenterCallerSession) ToCenterToken(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _Center.Contract.ToCenterToken(&_Center.CallOpts, arg0, arg1)
}

// ToEdgeToken is a free data retrieval call binding the contract method 0x527b0f64.
//
// Solidity: function toEdgeToken(address , uint256 ) view returns(address)
func (_Center *CenterCaller) ToEdgeToken(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "toEdgeToken", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ToEdgeToken is a free data retrieval call binding the contract method 0x527b0f64.
//
// Solidity: function toEdgeToken(address , uint256 ) view returns(address)
func (_Center *CenterSession) ToEdgeToken(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Center.Contract.ToEdgeToken(&_Center.CallOpts, arg0, arg1)
}

// ToEdgeToken is a free data retrieval call binding the contract method 0x527b0f64.
//
// Solidity: function toEdgeToken(address , uint256 ) view returns(address)
func (_Center *CenterCallerSession) ToEdgeToken(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Center.Contract.ToEdgeToken(&_Center.CallOpts, arg0, arg1)
}

// TreasuryTo is a free data retrieval call binding the contract method 0x42079671.
//
// Solidity: function treasuryTo() view returns(address)
func (_Center *CenterCaller) TreasuryTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "treasuryTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryTo is a free data retrieval call binding the contract method 0x42079671.
//
// Solidity: function treasuryTo() view returns(address)
func (_Center *CenterSession) TreasuryTo() (common.Address, error) {
	return _Center.Contract.TreasuryTo(&_Center.CallOpts)
}

// TreasuryTo is a free data retrieval call binding the contract method 0x42079671.
//
// Solidity: function treasuryTo() view returns(address)
func (_Center *CenterCallerSession) TreasuryTo() (common.Address, error) {
	return _Center.Contract.TreasuryTo(&_Center.CallOpts)
}

// TxHandled is a free data retrieval call binding the contract method 0x824df449.
//
// Solidity: function txHandled(string ) view returns(bool)
func (_Center *CenterCaller) TxHandled(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Center.contract.Call(opts, &out, "txHandled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TxHandled is a free data retrieval call binding the contract method 0x824df449.
//
// Solidity: function txHandled(string ) view returns(bool)
func (_Center *CenterSession) TxHandled(arg0 string) (bool, error) {
	return _Center.Contract.TxHandled(&_Center.CallOpts, arg0)
}

// TxHandled is a free data retrieval call binding the contract method 0x824df449.
//
// Solidity: function txHandled(string ) view returns(bool)
func (_Center *CenterCallerSession) TxHandled(arg0 string) (bool, error) {
	return _Center.Contract.TxHandled(&_Center.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0xa766a392.
//
// Solidity: function addToken(address _centerToken, uint256 _edgeChainId, address _edgeToken) returns()
func (_Center *CenterTransactor) AddToken(opts *bind.TransactOpts, _centerToken common.Address, _edgeChainId *big.Int, _edgeToken common.Address) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "addToken", _centerToken, _edgeChainId, _edgeToken)
}

// AddToken is a paid mutator transaction binding the contract method 0xa766a392.
//
// Solidity: function addToken(address _centerToken, uint256 _edgeChainId, address _edgeToken) returns()
func (_Center *CenterSession) AddToken(_centerToken common.Address, _edgeChainId *big.Int, _edgeToken common.Address) (*types.Transaction, error) {
	return _Center.Contract.AddToken(&_Center.TransactOpts, _centerToken, _edgeChainId, _edgeToken)
}

// AddToken is a paid mutator transaction binding the contract method 0xa766a392.
//
// Solidity: function addToken(address _centerToken, uint256 _edgeChainId, address _edgeToken) returns()
func (_Center *CenterTransactorSession) AddToken(_centerToken common.Address, _edgeChainId *big.Int, _edgeToken common.Address) (*types.Transaction, error) {
	return _Center.Contract.AddToken(&_Center.TransactOpts, _centerToken, _edgeChainId, _edgeToken)
}

// CrossIn is a paid mutator transaction binding the contract method 0x2cb0abde.
//
// Solidity: function crossIn((uint256,address,bytes,uint256,bytes,uint256) p, string txid) returns()
func (_Center *CenterTransactor) CrossIn(opts *bind.TransactOpts, p OutParam, txid string) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "crossIn", p, txid)
}

// CrossIn is a paid mutator transaction binding the contract method 0x2cb0abde.
//
// Solidity: function crossIn((uint256,address,bytes,uint256,bytes,uint256) p, string txid) returns()
func (_Center *CenterSession) CrossIn(p OutParam, txid string) (*types.Transaction, error) {
	return _Center.Contract.CrossIn(&_Center.TransactOpts, p, txid)
}

// CrossIn is a paid mutator transaction binding the contract method 0x2cb0abde.
//
// Solidity: function crossIn((uint256,address,bytes,uint256,bytes,uint256) p, string txid) returns()
func (_Center *CenterTransactorSession) CrossIn(p OutParam, txid string) (*types.Transaction, error) {
	return _Center.Contract.CrossIn(&_Center.TransactOpts, p, txid)
}

// CrossOut is a paid mutator transaction binding the contract method 0xe73d88d3.
//
// Solidity: function crossOut(address fromToken, uint256 toChainId, bytes to, uint256 amount) payable returns()
func (_Center *CenterTransactor) CrossOut(opts *bind.TransactOpts, fromToken common.Address, toChainId *big.Int, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "crossOut", fromToken, toChainId, to, amount)
}

// CrossOut is a paid mutator transaction binding the contract method 0xe73d88d3.
//
// Solidity: function crossOut(address fromToken, uint256 toChainId, bytes to, uint256 amount) payable returns()
func (_Center *CenterSession) CrossOut(fromToken common.Address, toChainId *big.Int, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _Center.Contract.CrossOut(&_Center.TransactOpts, fromToken, toChainId, to, amount)
}

// CrossOut is a paid mutator transaction binding the contract method 0xe73d88d3.
//
// Solidity: function crossOut(address fromToken, uint256 toChainId, bytes to, uint256 amount) payable returns()
func (_Center *CenterTransactorSession) CrossOut(fromToken common.Address, toChainId *big.Int, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _Center.Contract.CrossOut(&_Center.TransactOpts, fromToken, toChainId, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_Center *CenterTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_Center *CenterSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Center.Contract.Deposit(&_Center.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_Center *CenterTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Center.Contract.Deposit(&_Center.TransactOpts, token, amount)
}

// ForwardCrossOut is a paid mutator transaction binding the contract method 0xa17cc003.
//
// Solidity: function forwardCrossOut((uint256,address,bytes,uint256,bytes,uint256) p, string txid) returns()
func (_Center *CenterTransactor) ForwardCrossOut(opts *bind.TransactOpts, p OutParam, txid string) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "forwardCrossOut", p, txid)
}

// ForwardCrossOut is a paid mutator transaction binding the contract method 0xa17cc003.
//
// Solidity: function forwardCrossOut((uint256,address,bytes,uint256,bytes,uint256) p, string txid) returns()
func (_Center *CenterSession) ForwardCrossOut(p OutParam, txid string) (*types.Transaction, error) {
	return _Center.Contract.ForwardCrossOut(&_Center.TransactOpts, p, txid)
}

// ForwardCrossOut is a paid mutator transaction binding the contract method 0xa17cc003.
//
// Solidity: function forwardCrossOut((uint256,address,bytes,uint256,bytes,uint256) p, string txid) returns()
func (_Center *CenterTransactorSession) ForwardCrossOut(p OutParam, txid string) (*types.Transaction, error) {
	return _Center.Contract.ForwardCrossOut(&_Center.TransactOpts, p, txid)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Center *CenterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Center *CenterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Center.Contract.GrantRole(&_Center.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Center *CenterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Center.Contract.GrantRole(&_Center.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xb4988fd0.
//
// Solidity: function initialize(uint256 _chainId, address _feeTo, address _treasuryTo) returns()
func (_Center *CenterTransactor) Initialize(opts *bind.TransactOpts, _chainId *big.Int, _feeTo common.Address, _treasuryTo common.Address) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "initialize", _chainId, _feeTo, _treasuryTo)
}

// Initialize is a paid mutator transaction binding the contract method 0xb4988fd0.
//
// Solidity: function initialize(uint256 _chainId, address _feeTo, address _treasuryTo) returns()
func (_Center *CenterSession) Initialize(_chainId *big.Int, _feeTo common.Address, _treasuryTo common.Address) (*types.Transaction, error) {
	return _Center.Contract.Initialize(&_Center.TransactOpts, _chainId, _feeTo, _treasuryTo)
}

// Initialize is a paid mutator transaction binding the contract method 0xb4988fd0.
//
// Solidity: function initialize(uint256 _chainId, address _feeTo, address _treasuryTo) returns()
func (_Center *CenterTransactorSession) Initialize(_chainId *big.Int, _feeTo common.Address, _treasuryTo common.Address) (*types.Transaction, error) {
	return _Center.Contract.Initialize(&_Center.TransactOpts, _chainId, _feeTo, _treasuryTo)
}

// Issue is a paid mutator transaction binding the contract method 0x1987ba0c.
//
// Solidity: function issue((uint256,address,bytes,uint256,address,bytes,uint256) p, string txid) returns()
func (_Center *CenterTransactor) Issue(opts *bind.TransactOpts, p InParam, txid string) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "issue", p, txid)
}

// Issue is a paid mutator transaction binding the contract method 0x1987ba0c.
//
// Solidity: function issue((uint256,address,bytes,uint256,address,bytes,uint256) p, string txid) returns()
func (_Center *CenterSession) Issue(p InParam, txid string) (*types.Transaction, error) {
	return _Center.Contract.Issue(&_Center.TransactOpts, p, txid)
}

// Issue is a paid mutator transaction binding the contract method 0x1987ba0c.
//
// Solidity: function issue((uint256,address,bytes,uint256,address,bytes,uint256) p, string txid) returns()
func (_Center *CenterTransactorSession) Issue(p InParam, txid string) (*types.Transaction, error) {
	return _Center.Contract.Issue(&_Center.TransactOpts, p, txid)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x97bb15c9.
//
// Solidity: function removeToken(address _centerToken, uint256 _edgeChainId, address _edgeToken) returns()
func (_Center *CenterTransactor) RemoveToken(opts *bind.TransactOpts, _centerToken common.Address, _edgeChainId *big.Int, _edgeToken common.Address) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "removeToken", _centerToken, _edgeChainId, _edgeToken)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x97bb15c9.
//
// Solidity: function removeToken(address _centerToken, uint256 _edgeChainId, address _edgeToken) returns()
func (_Center *CenterSession) RemoveToken(_centerToken common.Address, _edgeChainId *big.Int, _edgeToken common.Address) (*types.Transaction, error) {
	return _Center.Contract.RemoveToken(&_Center.TransactOpts, _centerToken, _edgeChainId, _edgeToken)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x97bb15c9.
//
// Solidity: function removeToken(address _centerToken, uint256 _edgeChainId, address _edgeToken) returns()
func (_Center *CenterTransactorSession) RemoveToken(_centerToken common.Address, _edgeChainId *big.Int, _edgeToken common.Address) (*types.Transaction, error) {
	return _Center.Contract.RemoveToken(&_Center.TransactOpts, _centerToken, _edgeChainId, _edgeToken)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Center *CenterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Center *CenterSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Center.Contract.RenounceRole(&_Center.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Center *CenterTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Center.Contract.RenounceRole(&_Center.TransactOpts, role, callerConfirmation)
}

// ResetStakingReward is a paid mutator transaction binding the contract method 0x0f4693c6.
//
// Solidity: function resetStakingReward(address _oToken) returns()
func (_Center *CenterTransactor) ResetStakingReward(opts *bind.TransactOpts, _oToken common.Address) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "resetStakingReward", _oToken)
}

// ResetStakingReward is a paid mutator transaction binding the contract method 0x0f4693c6.
//
// Solidity: function resetStakingReward(address _oToken) returns()
func (_Center *CenterSession) ResetStakingReward(_oToken common.Address) (*types.Transaction, error) {
	return _Center.Contract.ResetStakingReward(&_Center.TransactOpts, _oToken)
}

// ResetStakingReward is a paid mutator transaction binding the contract method 0x0f4693c6.
//
// Solidity: function resetStakingReward(address _oToken) returns()
func (_Center *CenterTransactorSession) ResetStakingReward(_oToken common.Address) (*types.Transaction, error) {
	return _Center.Contract.ResetStakingReward(&_Center.TransactOpts, _oToken)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Center *CenterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Center *CenterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Center.Contract.RevokeRole(&_Center.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Center *CenterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Center.Contract.RevokeRole(&_Center.TransactOpts, role, account)
}

// SetFee is a paid mutator transaction binding the contract method 0xa5886f2d.
//
// Solidity: function setFee(address token, uint256[] toChainIds, uint256[] _fixFees, uint256[] _ratioFeesHigh, uint256[] _ratioFeesMedium, uint256[] _ratioFeesLow, uint256[] _remains) returns()
func (_Center *CenterTransactor) SetFee(opts *bind.TransactOpts, token common.Address, toChainIds []*big.Int, _fixFees []*big.Int, _ratioFeesHigh []*big.Int, _ratioFeesMedium []*big.Int, _ratioFeesLow []*big.Int, _remains []*big.Int) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "setFee", token, toChainIds, _fixFees, _ratioFeesHigh, _ratioFeesMedium, _ratioFeesLow, _remains)
}

// SetFee is a paid mutator transaction binding the contract method 0xa5886f2d.
//
// Solidity: function setFee(address token, uint256[] toChainIds, uint256[] _fixFees, uint256[] _ratioFeesHigh, uint256[] _ratioFeesMedium, uint256[] _ratioFeesLow, uint256[] _remains) returns()
func (_Center *CenterSession) SetFee(token common.Address, toChainIds []*big.Int, _fixFees []*big.Int, _ratioFeesHigh []*big.Int, _ratioFeesMedium []*big.Int, _ratioFeesLow []*big.Int, _remains []*big.Int) (*types.Transaction, error) {
	return _Center.Contract.SetFee(&_Center.TransactOpts, token, toChainIds, _fixFees, _ratioFeesHigh, _ratioFeesMedium, _ratioFeesLow, _remains)
}

// SetFee is a paid mutator transaction binding the contract method 0xa5886f2d.
//
// Solidity: function setFee(address token, uint256[] toChainIds, uint256[] _fixFees, uint256[] _ratioFeesHigh, uint256[] _ratioFeesMedium, uint256[] _ratioFeesLow, uint256[] _remains) returns()
func (_Center *CenterTransactorSession) SetFee(token common.Address, toChainIds []*big.Int, _fixFees []*big.Int, _ratioFeesHigh []*big.Int, _ratioFeesMedium []*big.Int, _ratioFeesLow []*big.Int, _remains []*big.Int) (*types.Transaction, error) {
	return _Center.Contract.SetFee(&_Center.TransactOpts, token, toChainIds, _fixFees, _ratioFeesHigh, _ratioFeesMedium, _ratioFeesLow, _remains)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Center *CenterTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Center *CenterSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Center.Contract.SetFeeTo(&_Center.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_Center *CenterTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _Center.Contract.SetFeeTo(&_Center.TransactOpts, _feeTo)
}

// SetFeeToTreasuryRatio is a paid mutator transaction binding the contract method 0xf88d0a2a.
//
// Solidity: function setFeeToTreasuryRatio(address _oToken, uint256 _ratio) returns()
func (_Center *CenterTransactor) SetFeeToTreasuryRatio(opts *bind.TransactOpts, _oToken common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "setFeeToTreasuryRatio", _oToken, _ratio)
}

// SetFeeToTreasuryRatio is a paid mutator transaction binding the contract method 0xf88d0a2a.
//
// Solidity: function setFeeToTreasuryRatio(address _oToken, uint256 _ratio) returns()
func (_Center *CenterSession) SetFeeToTreasuryRatio(_oToken common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _Center.Contract.SetFeeToTreasuryRatio(&_Center.TransactOpts, _oToken, _ratio)
}

// SetFeeToTreasuryRatio is a paid mutator transaction binding the contract method 0xf88d0a2a.
//
// Solidity: function setFeeToTreasuryRatio(address _oToken, uint256 _ratio) returns()
func (_Center *CenterTransactorSession) SetFeeToTreasuryRatio(_oToken common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _Center.Contract.SetFeeToTreasuryRatio(&_Center.TransactOpts, _oToken, _ratio)
}

// SetIsCoin is a paid mutator transaction binding the contract method 0xd93e76bb.
//
// Solidity: function setIsCoin(address token, bool _isCoin) returns()
func (_Center *CenterTransactor) SetIsCoin(opts *bind.TransactOpts, token common.Address, _isCoin bool) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "setIsCoin", token, _isCoin)
}

// SetIsCoin is a paid mutator transaction binding the contract method 0xd93e76bb.
//
// Solidity: function setIsCoin(address token, bool _isCoin) returns()
func (_Center *CenterSession) SetIsCoin(token common.Address, _isCoin bool) (*types.Transaction, error) {
	return _Center.Contract.SetIsCoin(&_Center.TransactOpts, token, _isCoin)
}

// SetIsCoin is a paid mutator transaction binding the contract method 0xd93e76bb.
//
// Solidity: function setIsCoin(address token, bool _isCoin) returns()
func (_Center *CenterTransactorSession) SetIsCoin(token common.Address, _isCoin bool) (*types.Transaction, error) {
	return _Center.Contract.SetIsCoin(&_Center.TransactOpts, token, _isCoin)
}

// SetRewardRatio is a paid mutator transaction binding the contract method 0x41950ad3.
//
// Solidity: function setRewardRatio(address _oToken, uint256 _ratio) returns()
func (_Center *CenterTransactor) SetRewardRatio(opts *bind.TransactOpts, _oToken common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "setRewardRatio", _oToken, _ratio)
}

// SetRewardRatio is a paid mutator transaction binding the contract method 0x41950ad3.
//
// Solidity: function setRewardRatio(address _oToken, uint256 _ratio) returns()
func (_Center *CenterSession) SetRewardRatio(_oToken common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _Center.Contract.SetRewardRatio(&_Center.TransactOpts, _oToken, _ratio)
}

// SetRewardRatio is a paid mutator transaction binding the contract method 0x41950ad3.
//
// Solidity: function setRewardRatio(address _oToken, uint256 _ratio) returns()
func (_Center *CenterTransactorSession) SetRewardRatio(_oToken common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _Center.Contract.SetRewardRatio(&_Center.TransactOpts, _oToken, _ratio)
}

// SetStakingRewards is a paid mutator transaction binding the contract method 0x9773a000.
//
// Solidity: function setStakingRewards(address[] oTokens, address[] _srs) returns()
func (_Center *CenterTransactor) SetStakingRewards(opts *bind.TransactOpts, oTokens []common.Address, _srs []common.Address) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "setStakingRewards", oTokens, _srs)
}

// SetStakingRewards is a paid mutator transaction binding the contract method 0x9773a000.
//
// Solidity: function setStakingRewards(address[] oTokens, address[] _srs) returns()
func (_Center *CenterSession) SetStakingRewards(oTokens []common.Address, _srs []common.Address) (*types.Transaction, error) {
	return _Center.Contract.SetStakingRewards(&_Center.TransactOpts, oTokens, _srs)
}

// SetStakingRewards is a paid mutator transaction binding the contract method 0x9773a000.
//
// Solidity: function setStakingRewards(address[] oTokens, address[] _srs) returns()
func (_Center *CenterTransactorSession) SetStakingRewards(oTokens []common.Address, _srs []common.Address) (*types.Transaction, error) {
	return _Center.Contract.SetStakingRewards(&_Center.TransactOpts, oTokens, _srs)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_Center *CenterTransactor) SetThreshold(opts *bind.TransactOpts, token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "setThreshold", token0, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_Center *CenterSession) SetThreshold(token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Center.Contract.SetThreshold(&_Center.TransactOpts, token0, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_Center *CenterTransactorSession) SetThreshold(token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Center.Contract.SetThreshold(&_Center.TransactOpts, token0, _threshold)
}

// SetTreasuryTo is a paid mutator transaction binding the contract method 0x08b60fb8.
//
// Solidity: function setTreasuryTo(address _treasuryTo) returns()
func (_Center *CenterTransactor) SetTreasuryTo(opts *bind.TransactOpts, _treasuryTo common.Address) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "setTreasuryTo", _treasuryTo)
}

// SetTreasuryTo is a paid mutator transaction binding the contract method 0x08b60fb8.
//
// Solidity: function setTreasuryTo(address _treasuryTo) returns()
func (_Center *CenterSession) SetTreasuryTo(_treasuryTo common.Address) (*types.Transaction, error) {
	return _Center.Contract.SetTreasuryTo(&_Center.TransactOpts, _treasuryTo)
}

// SetTreasuryTo is a paid mutator transaction binding the contract method 0x08b60fb8.
//
// Solidity: function setTreasuryTo(address _treasuryTo) returns()
func (_Center *CenterTransactorSession) SetTreasuryTo(_treasuryTo common.Address) (*types.Transaction, error) {
	return _Center.Contract.SetTreasuryTo(&_Center.TransactOpts, _treasuryTo)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x1bdbce49.
//
// Solidity: function setWhitelist(address oToken, address user, bool _isInWhitelist) returns()
func (_Center *CenterTransactor) SetWhitelist(opts *bind.TransactOpts, oToken common.Address, user common.Address, _isInWhitelist bool) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "setWhitelist", oToken, user, _isInWhitelist)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x1bdbce49.
//
// Solidity: function setWhitelist(address oToken, address user, bool _isInWhitelist) returns()
func (_Center *CenterSession) SetWhitelist(oToken common.Address, user common.Address, _isInWhitelist bool) (*types.Transaction, error) {
	return _Center.Contract.SetWhitelist(&_Center.TransactOpts, oToken, user, _isInWhitelist)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x1bdbce49.
//
// Solidity: function setWhitelist(address oToken, address user, bool _isInWhitelist) returns()
func (_Center *CenterTransactorSession) SetWhitelist(oToken common.Address, user common.Address, _isInWhitelist bool) (*types.Transaction, error) {
	return _Center.Contract.SetWhitelist(&_Center.TransactOpts, oToken, user, _isInWhitelist)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Center *CenterTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Center *CenterSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Center.Contract.UpgradeToAndCall(&_Center.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Center *CenterTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Center.Contract.UpgradeToAndCall(&_Center.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9483e91a.
//
// Solidity: function withdraw(address oToken, uint256 toChainId, bytes to, uint256 amount) returns()
func (_Center *CenterTransactor) Withdraw(opts *bind.TransactOpts, oToken common.Address, toChainId *big.Int, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _Center.contract.Transact(opts, "withdraw", oToken, toChainId, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9483e91a.
//
// Solidity: function withdraw(address oToken, uint256 toChainId, bytes to, uint256 amount) returns()
func (_Center *CenterSession) Withdraw(oToken common.Address, toChainId *big.Int, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _Center.Contract.Withdraw(&_Center.TransactOpts, oToken, toChainId, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x9483e91a.
//
// Solidity: function withdraw(address oToken, uint256 toChainId, bytes to, uint256 amount) returns()
func (_Center *CenterTransactorSession) Withdraw(oToken common.Address, toChainId *big.Int, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _Center.Contract.Withdraw(&_Center.TransactOpts, oToken, toChainId, to, amount)
}

// CenterCoinSetedIterator is returned from FilterCoinSeted and is used to iterate over the raw logs and unpacked data for CoinSeted events raised by the Center contract.
type CenterCoinSetedIterator struct {
	Event *CenterCoinSeted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterCoinSetedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterCoinSeted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterCoinSeted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterCoinSetedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterCoinSetedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterCoinSeted represents a CoinSeted event raised by the Center contract.
type CenterCoinSeted struct {
	Coin common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCoinSeted is a free log retrieval operation binding the contract event 0x068ea4c4aed0572edc3ba8586d929cd0d8022e8bc1794748e4ff980ce55c3292.
//
// Solidity: event CoinSeted(address coin)
func (_Center *CenterFilterer) FilterCoinSeted(opts *bind.FilterOpts) (*CenterCoinSetedIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "CoinSeted")
	if err != nil {
		return nil, err
	}
	return &CenterCoinSetedIterator{contract: _Center.contract, event: "CoinSeted", logs: logs, sub: sub}, nil
}

// WatchCoinSeted is a free log subscription operation binding the contract event 0x068ea4c4aed0572edc3ba8586d929cd0d8022e8bc1794748e4ff980ce55c3292.
//
// Solidity: event CoinSeted(address coin)
func (_Center *CenterFilterer) WatchCoinSeted(opts *bind.WatchOpts, sink chan<- *CenterCoinSeted) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "CoinSeted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterCoinSeted)
				if err := _Center.contract.UnpackLog(event, "CoinSeted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCoinSeted is a log parse operation binding the contract event 0x068ea4c4aed0572edc3ba8586d929cd0d8022e8bc1794748e4ff980ce55c3292.
//
// Solidity: event CoinSeted(address coin)
func (_Center *CenterFilterer) ParseCoinSeted(log types.Log) (*CenterCoinSeted, error) {
	event := new(CenterCoinSeted)
	if err := _Center.contract.UnpackLog(event, "CoinSeted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterCrossInFailedIterator is returned from FilterCrossInFailed and is used to iterate over the raw logs and unpacked data for CrossInFailed events raised by the Center contract.
type CenterCrossInFailedIterator struct {
	Event *CenterCrossInFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterCrossInFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterCrossInFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterCrossInFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterCrossInFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterCrossInFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterCrossInFailed represents a CrossInFailed event raised by the Center contract.
type CenterCrossInFailed struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCrossInFailed is a free log retrieval operation binding the contract event 0x8d6ab4790984b48ccd527098208bc9944798ae24bbe40995ecb672728a4c0203.
//
// Solidity: event CrossInFailed((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) FilterCrossInFailed(opts *bind.FilterOpts) (*CenterCrossInFailedIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "CrossInFailed")
	if err != nil {
		return nil, err
	}
	return &CenterCrossInFailedIterator{contract: _Center.contract, event: "CrossInFailed", logs: logs, sub: sub}, nil
}

// WatchCrossInFailed is a free log subscription operation binding the contract event 0x8d6ab4790984b48ccd527098208bc9944798ae24bbe40995ecb672728a4c0203.
//
// Solidity: event CrossInFailed((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) WatchCrossInFailed(opts *bind.WatchOpts, sink chan<- *CenterCrossInFailed) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "CrossInFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterCrossInFailed)
				if err := _Center.contract.UnpackLog(event, "CrossInFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossInFailed is a log parse operation binding the contract event 0x8d6ab4790984b48ccd527098208bc9944798ae24bbe40995ecb672728a4c0203.
//
// Solidity: event CrossInFailed((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) ParseCrossInFailed(log types.Log) (*CenterCrossInFailed, error) {
	event := new(CenterCrossInFailed)
	if err := _Center.contract.UnpackLog(event, "CrossInFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterCrossInedIterator is returned from FilterCrossIned and is used to iterate over the raw logs and unpacked data for CrossIned events raised by the Center contract.
type CenterCrossInedIterator struct {
	Event *CenterCrossIned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterCrossInedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterCrossIned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterCrossIned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterCrossInedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterCrossInedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterCrossIned represents a CrossIned event raised by the Center contract.
type CenterCrossIned struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCrossIned is a free log retrieval operation binding the contract event 0x186a37dd020447374a89e19ce3bc90447d34d9af95d221f64366bb54be38617b.
//
// Solidity: event CrossIned((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) FilterCrossIned(opts *bind.FilterOpts) (*CenterCrossInedIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "CrossIned")
	if err != nil {
		return nil, err
	}
	return &CenterCrossInedIterator{contract: _Center.contract, event: "CrossIned", logs: logs, sub: sub}, nil
}

// WatchCrossIned is a free log subscription operation binding the contract event 0x186a37dd020447374a89e19ce3bc90447d34d9af95d221f64366bb54be38617b.
//
// Solidity: event CrossIned((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) WatchCrossIned(opts *bind.WatchOpts, sink chan<- *CenterCrossIned) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "CrossIned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterCrossIned)
				if err := _Center.contract.UnpackLog(event, "CrossIned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossIned is a log parse operation binding the contract event 0x186a37dd020447374a89e19ce3bc90447d34d9af95d221f64366bb54be38617b.
//
// Solidity: event CrossIned((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) ParseCrossIned(log types.Log) (*CenterCrossIned, error) {
	event := new(CenterCrossIned)
	if err := _Center.contract.UnpackLog(event, "CrossIned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterCrossOutedIterator is returned from FilterCrossOuted and is used to iterate over the raw logs and unpacked data for CrossOuted events raised by the Center contract.
type CenterCrossOutedIterator struct {
	Event *CenterCrossOuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterCrossOutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterCrossOuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterCrossOuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterCrossOutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterCrossOutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterCrossOuted represents a CrossOuted event raised by the Center contract.
type CenterCrossOuted struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCrossOuted is a free log retrieval operation binding the contract event 0xd741980ccff5de68cf85c9ac157c9d947ab1a72db0c0d6035325cc7ccb567d22.
//
// Solidity: event CrossOuted((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) FilterCrossOuted(opts *bind.FilterOpts) (*CenterCrossOutedIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "CrossOuted")
	if err != nil {
		return nil, err
	}
	return &CenterCrossOutedIterator{contract: _Center.contract, event: "CrossOuted", logs: logs, sub: sub}, nil
}

// WatchCrossOuted is a free log subscription operation binding the contract event 0xd741980ccff5de68cf85c9ac157c9d947ab1a72db0c0d6035325cc7ccb567d22.
//
// Solidity: event CrossOuted((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) WatchCrossOuted(opts *bind.WatchOpts, sink chan<- *CenterCrossOuted) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "CrossOuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterCrossOuted)
				if err := _Center.contract.UnpackLog(event, "CrossOuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossOuted is a log parse operation binding the contract event 0xd741980ccff5de68cf85c9ac157c9d947ab1a72db0c0d6035325cc7ccb567d22.
//
// Solidity: event CrossOuted((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) ParseCrossOuted(log types.Log) (*CenterCrossOuted, error) {
	event := new(CenterCrossOuted)
	if err := _Center.contract.UnpackLog(event, "CrossOuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Center contract.
type CenterDepositedIterator struct {
	Event *CenterDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterDeposited represents a Deposited event raised by the Center contract.
type CenterDeposited struct {
	FromChainId *big.Int
	FromToken   common.Address
	From        common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x984a71c9d95fd4794aeba33ae72edfec22053fde75488d63abef9dc69ee795af.
//
// Solidity: event Deposited(uint256 fromChainId, address fromToken, address from, uint256 amount)
func (_Center *CenterFilterer) FilterDeposited(opts *bind.FilterOpts) (*CenterDepositedIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &CenterDepositedIterator{contract: _Center.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x984a71c9d95fd4794aeba33ae72edfec22053fde75488d63abef9dc69ee795af.
//
// Solidity: event Deposited(uint256 fromChainId, address fromToken, address from, uint256 amount)
func (_Center *CenterFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *CenterDeposited) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterDeposited)
				if err := _Center.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x984a71c9d95fd4794aeba33ae72edfec22053fde75488d63abef9dc69ee795af.
//
// Solidity: event Deposited(uint256 fromChainId, address fromToken, address from, uint256 amount)
func (_Center *CenterFilterer) ParseDeposited(log types.Log) (*CenterDeposited, error) {
	event := new(CenterDeposited)
	if err := _Center.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterForwardCrossOutFailedIterator is returned from FilterForwardCrossOutFailed and is used to iterate over the raw logs and unpacked data for ForwardCrossOutFailed events raised by the Center contract.
type CenterForwardCrossOutFailedIterator struct {
	Event *CenterForwardCrossOutFailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterForwardCrossOutFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterForwardCrossOutFailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterForwardCrossOutFailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterForwardCrossOutFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterForwardCrossOutFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterForwardCrossOutFailed represents a ForwardCrossOutFailed event raised by the Center contract.
type CenterForwardCrossOutFailed struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterForwardCrossOutFailed is a free log retrieval operation binding the contract event 0x61e468e1aee35aac0f5f87ea7ca89ba291e44c32695fb303ee864613d30ba3c3.
//
// Solidity: event ForwardCrossOutFailed((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) FilterForwardCrossOutFailed(opts *bind.FilterOpts) (*CenterForwardCrossOutFailedIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "ForwardCrossOutFailed")
	if err != nil {
		return nil, err
	}
	return &CenterForwardCrossOutFailedIterator{contract: _Center.contract, event: "ForwardCrossOutFailed", logs: logs, sub: sub}, nil
}

// WatchForwardCrossOutFailed is a free log subscription operation binding the contract event 0x61e468e1aee35aac0f5f87ea7ca89ba291e44c32695fb303ee864613d30ba3c3.
//
// Solidity: event ForwardCrossOutFailed((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) WatchForwardCrossOutFailed(opts *bind.WatchOpts, sink chan<- *CenterForwardCrossOutFailed) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "ForwardCrossOutFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterForwardCrossOutFailed)
				if err := _Center.contract.UnpackLog(event, "ForwardCrossOutFailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseForwardCrossOutFailed is a log parse operation binding the contract event 0x61e468e1aee35aac0f5f87ea7ca89ba291e44c32695fb303ee864613d30ba3c3.
//
// Solidity: event ForwardCrossOutFailed((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) ParseForwardCrossOutFailed(log types.Log) (*CenterForwardCrossOutFailed, error) {
	event := new(CenterForwardCrossOutFailed)
	if err := _Center.contract.UnpackLog(event, "ForwardCrossOutFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterForwardCrossOutedIterator is returned from FilterForwardCrossOuted and is used to iterate over the raw logs and unpacked data for ForwardCrossOuted events raised by the Center contract.
type CenterForwardCrossOutedIterator struct {
	Event *CenterForwardCrossOuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterForwardCrossOutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterForwardCrossOuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterForwardCrossOuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterForwardCrossOutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterForwardCrossOutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterForwardCrossOuted represents a ForwardCrossOuted event raised by the Center contract.
type CenterForwardCrossOuted struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterForwardCrossOuted is a free log retrieval operation binding the contract event 0xc3b2c8626c76e00d71763942beacf23e2f41074b27e054d8a9bd09fecbedd8d4.
//
// Solidity: event ForwardCrossOuted((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) FilterForwardCrossOuted(opts *bind.FilterOpts) (*CenterForwardCrossOutedIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "ForwardCrossOuted")
	if err != nil {
		return nil, err
	}
	return &CenterForwardCrossOutedIterator{contract: _Center.contract, event: "ForwardCrossOuted", logs: logs, sub: sub}, nil
}

// WatchForwardCrossOuted is a free log subscription operation binding the contract event 0xc3b2c8626c76e00d71763942beacf23e2f41074b27e054d8a9bd09fecbedd8d4.
//
// Solidity: event ForwardCrossOuted((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) WatchForwardCrossOuted(opts *bind.WatchOpts, sink chan<- *CenterForwardCrossOuted) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "ForwardCrossOuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterForwardCrossOuted)
				if err := _Center.contract.UnpackLog(event, "ForwardCrossOuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseForwardCrossOuted is a log parse operation binding the contract event 0xc3b2c8626c76e00d71763942beacf23e2f41074b27e054d8a9bd09fecbedd8d4.
//
// Solidity: event ForwardCrossOuted((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) ParseForwardCrossOuted(log types.Log) (*CenterForwardCrossOuted, error) {
	event := new(CenterForwardCrossOuted)
	if err := _Center.contract.UnpackLog(event, "ForwardCrossOuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Center contract.
type CenterInitializedIterator struct {
	Event *CenterInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterInitialized represents a Initialized event raised by the Center contract.
type CenterInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Center *CenterFilterer) FilterInitialized(opts *bind.FilterOpts) (*CenterInitializedIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CenterInitializedIterator{contract: _Center.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Center *CenterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CenterInitialized) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterInitialized)
				if err := _Center.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Center *CenterFilterer) ParseInitialized(log types.Log) (*CenterInitialized, error) {
	event := new(CenterInitialized)
	if err := _Center.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterIssuedIterator is returned from FilterIssued and is used to iterate over the raw logs and unpacked data for Issued events raised by the Center contract.
type CenterIssuedIterator struct {
	Event *CenterIssued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterIssued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterIssued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterIssued represents a Issued event raised by the Center contract.
type CenterIssued struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIssued is a free log retrieval operation binding the contract event 0xc2a9778caeda1799613927ccd8fb1611c8dad0bfe2dc0b03d4f82ef2afe801eb.
//
// Solidity: event Issued((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) FilterIssued(opts *bind.FilterOpts) (*CenterIssuedIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "Issued")
	if err != nil {
		return nil, err
	}
	return &CenterIssuedIterator{contract: _Center.contract, event: "Issued", logs: logs, sub: sub}, nil
}

// WatchIssued is a free log subscription operation binding the contract event 0xc2a9778caeda1799613927ccd8fb1611c8dad0bfe2dc0b03d4f82ef2afe801eb.
//
// Solidity: event Issued((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) WatchIssued(opts *bind.WatchOpts, sink chan<- *CenterIssued) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "Issued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterIssued)
				if err := _Center.contract.UnpackLog(event, "Issued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseIssued is a log parse operation binding the contract event 0xc2a9778caeda1799613927ccd8fb1611c8dad0bfe2dc0b03d4f82ef2afe801eb.
//
// Solidity: event Issued((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) ParseIssued(log types.Log) (*CenterIssued, error) {
	event := new(CenterIssued)
	if err := _Center.contract.UnpackLog(event, "Issued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the Center contract.
type CenterProposalVotedIterator struct {
	Event *CenterProposalVoted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterProposalVoted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterProposalVoted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterProposalVoted represents a ProposalVoted event raised by the Center contract.
type CenterProposalVoted struct {
	Token     common.Address
	From      []byte
	To        []byte
	Amount    *big.Int
	Proposer  common.Address
	Count     *big.Int
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProposalVoted is a free log retrieval operation binding the contract event 0xaa8d9692f9529994d8599d4a8aad3f9cb2e1290ffe484cf531edfb62de17258c.
//
// Solidity: event ProposalVoted(address token, bytes from, bytes to, uint256 amount, address proposer, uint256 count, uint256 threshold)
func (_Center *CenterFilterer) FilterProposalVoted(opts *bind.FilterOpts) (*CenterProposalVotedIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return &CenterProposalVotedIterator{contract: _Center.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0xaa8d9692f9529994d8599d4a8aad3f9cb2e1290ffe484cf531edfb62de17258c.
//
// Solidity: event ProposalVoted(address token, bytes from, bytes to, uint256 amount, address proposer, uint256 count, uint256 threshold)
func (_Center *CenterFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *CenterProposalVoted) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterProposalVoted)
				if err := _Center.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalVoted is a log parse operation binding the contract event 0xaa8d9692f9529994d8599d4a8aad3f9cb2e1290ffe484cf531edfb62de17258c.
//
// Solidity: event ProposalVoted(address token, bytes from, bytes to, uint256 amount, address proposer, uint256 count, uint256 threshold)
func (_Center *CenterFilterer) ParseProposalVoted(log types.Log) (*CenterProposalVoted, error) {
	event := new(CenterProposalVoted)
	if err := _Center.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Center contract.
type CenterRoleAdminChangedIterator struct {
	Event *CenterRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterRoleAdminChanged represents a RoleAdminChanged event raised by the Center contract.
type CenterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Center *CenterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CenterRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Center.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CenterRoleAdminChangedIterator{contract: _Center.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Center *CenterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CenterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Center.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterRoleAdminChanged)
				if err := _Center.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Center *CenterFilterer) ParseRoleAdminChanged(log types.Log) (*CenterRoleAdminChanged, error) {
	event := new(CenterRoleAdminChanged)
	if err := _Center.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Center contract.
type CenterRoleGrantedIterator struct {
	Event *CenterRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterRoleGranted represents a RoleGranted event raised by the Center contract.
type CenterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Center *CenterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CenterRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Center.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CenterRoleGrantedIterator{contract: _Center.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Center *CenterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CenterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Center.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterRoleGranted)
				if err := _Center.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Center *CenterFilterer) ParseRoleGranted(log types.Log) (*CenterRoleGranted, error) {
	event := new(CenterRoleGranted)
	if err := _Center.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Center contract.
type CenterRoleRevokedIterator struct {
	Event *CenterRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterRoleRevoked represents a RoleRevoked event raised by the Center contract.
type CenterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Center *CenterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CenterRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Center.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CenterRoleRevokedIterator{contract: _Center.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Center *CenterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CenterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Center.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterRoleRevoked)
				if err := _Center.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Center *CenterFilterer) ParseRoleRevoked(log types.Log) (*CenterRoleRevoked, error) {
	event := new(CenterRoleRevoked)
	if err := _Center.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterStakingRewardResetedIterator is returned from FilterStakingRewardReseted and is used to iterate over the raw logs and unpacked data for StakingRewardReseted events raised by the Center contract.
type CenterStakingRewardResetedIterator struct {
	Event *CenterStakingRewardReseted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterStakingRewardResetedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterStakingRewardReseted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterStakingRewardReseted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterStakingRewardResetedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterStakingRewardResetedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterStakingRewardReseted represents a StakingRewardReseted event raised by the Center contract.
type CenterStakingRewardReseted struct {
	OToken common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakingRewardReseted is a free log retrieval operation binding the contract event 0x4b8830da429e7f8c1dc368e9a9a3dc21aa6e3e08d115679e616b1fedaf420aeb.
//
// Solidity: event StakingRewardReseted(address oToken)
func (_Center *CenterFilterer) FilterStakingRewardReseted(opts *bind.FilterOpts) (*CenterStakingRewardResetedIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "StakingRewardReseted")
	if err != nil {
		return nil, err
	}
	return &CenterStakingRewardResetedIterator{contract: _Center.contract, event: "StakingRewardReseted", logs: logs, sub: sub}, nil
}

// WatchStakingRewardReseted is a free log subscription operation binding the contract event 0x4b8830da429e7f8c1dc368e9a9a3dc21aa6e3e08d115679e616b1fedaf420aeb.
//
// Solidity: event StakingRewardReseted(address oToken)
func (_Center *CenterFilterer) WatchStakingRewardReseted(opts *bind.WatchOpts, sink chan<- *CenterStakingRewardReseted) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "StakingRewardReseted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterStakingRewardReseted)
				if err := _Center.contract.UnpackLog(event, "StakingRewardReseted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStakingRewardReseted is a log parse operation binding the contract event 0x4b8830da429e7f8c1dc368e9a9a3dc21aa6e3e08d115679e616b1fedaf420aeb.
//
// Solidity: event StakingRewardReseted(address oToken)
func (_Center *CenterFilterer) ParseStakingRewardReseted(log types.Log) (*CenterStakingRewardReseted, error) {
	event := new(CenterStakingRewardReseted)
	if err := _Center.contract.UnpackLog(event, "StakingRewardReseted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterSupportedIterator is returned from FilterSupported and is used to iterate over the raw logs and unpacked data for Supported events raised by the Center contract.
type CenterSupportedIterator struct {
	Event *CenterSupported // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterSupportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterSupported)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterSupported)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterSupportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterSupportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterSupported represents a Supported event raised by the Center contract.
type CenterSupported struct {
	Token   common.Address
	ChainId *big.Int
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSupported is a free log retrieval operation binding the contract event 0xd7cadba2609cba4364d862761a89c823e956d810d9947b1f158ad1aa9c2affbc.
//
// Solidity: event Supported(address token, uint256 chainId, bool status)
func (_Center *CenterFilterer) FilterSupported(opts *bind.FilterOpts) (*CenterSupportedIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "Supported")
	if err != nil {
		return nil, err
	}
	return &CenterSupportedIterator{contract: _Center.contract, event: "Supported", logs: logs, sub: sub}, nil
}

// WatchSupported is a free log subscription operation binding the contract event 0xd7cadba2609cba4364d862761a89c823e956d810d9947b1f158ad1aa9c2affbc.
//
// Solidity: event Supported(address token, uint256 chainId, bool status)
func (_Center *CenterFilterer) WatchSupported(opts *bind.WatchOpts, sink chan<- *CenterSupported) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "Supported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterSupported)
				if err := _Center.contract.UnpackLog(event, "Supported", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSupported is a log parse operation binding the contract event 0xd7cadba2609cba4364d862761a89c823e956d810d9947b1f158ad1aa9c2affbc.
//
// Solidity: event Supported(address token, uint256 chainId, bool status)
func (_Center *CenterFilterer) ParseSupported(log types.Log) (*CenterSupported, error) {
	event := new(CenterSupported)
	if err := _Center.contract.UnpackLog(event, "Supported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the Center contract.
type CenterThresholdChangedIterator struct {
	Event *CenterThresholdChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterThresholdChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterThresholdChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterThresholdChanged represents a ThresholdChanged event raised by the Center contract.
type CenterThresholdChanged struct {
	Token        common.Address
	OldThreshold *big.Int
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_Center *CenterFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*CenterThresholdChangedIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &CenterThresholdChangedIterator{contract: _Center.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_Center *CenterFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *CenterThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterThresholdChanged)
				if err := _Center.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseThresholdChanged is a log parse operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_Center *CenterFilterer) ParseThresholdChanged(log types.Log) (*CenterThresholdChanged, error) {
	event := new(CenterThresholdChanged)
	if err := _Center.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Center contract.
type CenterUpgradedIterator struct {
	Event *CenterUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterUpgraded represents a Upgraded event raised by the Center contract.
type CenterUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Center *CenterFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*CenterUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Center.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &CenterUpgradedIterator{contract: _Center.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Center *CenterFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *CenterUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Center.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterUpgraded)
				if err := _Center.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Center *CenterFilterer) ParseUpgraded(log types.Log) (*CenterUpgraded, error) {
	event := new(CenterUpgraded)
	if err := _Center.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterWithdrawedIterator is returned from FilterWithdrawed and is used to iterate over the raw logs and unpacked data for Withdrawed events raised by the Center contract.
type CenterWithdrawedIterator struct {
	Event *CenterWithdrawed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterWithdrawed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterWithdrawed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterWithdrawed represents a Withdrawed event raised by the Center contract.
type CenterWithdrawed struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawed is a free log retrieval operation binding the contract event 0x8ee630ab76a39b3fb8b9ef6e86730269aaf0cc092ec53430810a1a2423590b73.
//
// Solidity: event Withdrawed((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) FilterWithdrawed(opts *bind.FilterOpts) (*CenterWithdrawedIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "Withdrawed")
	if err != nil {
		return nil, err
	}
	return &CenterWithdrawedIterator{contract: _Center.contract, event: "Withdrawed", logs: logs, sub: sub}, nil
}

// WatchWithdrawed is a free log subscription operation binding the contract event 0x8ee630ab76a39b3fb8b9ef6e86730269aaf0cc092ec53430810a1a2423590b73.
//
// Solidity: event Withdrawed((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) WatchWithdrawed(opts *bind.WatchOpts, sink chan<- *CenterWithdrawed) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "Withdrawed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterWithdrawed)
				if err := _Center.contract.UnpackLog(event, "Withdrawed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawed is a log parse operation binding the contract event 0x8ee630ab76a39b3fb8b9ef6e86730269aaf0cc092ec53430810a1a2423590b73.
//
// Solidity: event Withdrawed((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) ParseWithdrawed(log types.Log) (*CenterWithdrawed, error) {
	event := new(CenterWithdrawed)
	if err := _Center.contract.UnpackLog(event, "Withdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CenterWithdrawedToCenterIterator is returned from FilterWithdrawedToCenter and is used to iterate over the raw logs and unpacked data for WithdrawedToCenter events raised by the Center contract.
type CenterWithdrawedToCenterIterator struct {
	Event *CenterWithdrawedToCenter // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CenterWithdrawedToCenterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CenterWithdrawedToCenter)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CenterWithdrawedToCenter)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CenterWithdrawedToCenterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CenterWithdrawedToCenterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CenterWithdrawedToCenter represents a WithdrawedToCenter event raised by the Center contract.
type CenterWithdrawedToCenter struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawedToCenter is a free log retrieval operation binding the contract event 0xd95e7902305a71cae6bac50839d79f21ce0aa62ff6bf7e2bb1c8dbd81755a941.
//
// Solidity: event WithdrawedToCenter((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) FilterWithdrawedToCenter(opts *bind.FilterOpts) (*CenterWithdrawedToCenterIterator, error) {

	logs, sub, err := _Center.contract.FilterLogs(opts, "WithdrawedToCenter")
	if err != nil {
		return nil, err
	}
	return &CenterWithdrawedToCenterIterator{contract: _Center.contract, event: "WithdrawedToCenter", logs: logs, sub: sub}, nil
}

// WatchWithdrawedToCenter is a free log subscription operation binding the contract event 0xd95e7902305a71cae6bac50839d79f21ce0aa62ff6bf7e2bb1c8dbd81755a941.
//
// Solidity: event WithdrawedToCenter((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) WatchWithdrawedToCenter(opts *bind.WatchOpts, sink chan<- *CenterWithdrawedToCenter) (event.Subscription, error) {

	logs, sub, err := _Center.contract.WatchLogs(opts, "WithdrawedToCenter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CenterWithdrawedToCenter)
				if err := _Center.contract.UnpackLog(event, "WithdrawedToCenter", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawedToCenter is a log parse operation binding the contract event 0xd95e7902305a71cae6bac50839d79f21ce0aa62ff6bf7e2bb1c8dbd81755a941.
//
// Solidity: event WithdrawedToCenter((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Center *CenterFilterer) ParseWithdrawedToCenter(log types.Log) (*CenterWithdrawedToCenter, error) {
	event := new(CenterWithdrawedToCenter)
	if err := _Center.contract.UnpackLog(event, "WithdrawedToCenter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
