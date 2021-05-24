// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CrossBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"CrossMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockFeeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockFeeRatio\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockFeeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockFeeRatio\",\"type\":\"uint256\"}],\"name\":\"FeeChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeeToAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeeToRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_roleFlag\",\"type\":\"bytes32\"}],\"name\":\"addSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"token0Addrs\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"token1Addrs\",\"type\":\"address[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_roleFlags\",\"type\":\"bytes32[]\"}],\"name\":\"addSupportTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"crossType\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"crossBurn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"crossMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"feeToLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"len\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getChainID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"getFeeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lockFeeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposalOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isExist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removeFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"}],\"name\":\"removeRoleFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"}],\"name\":\"removeSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"}],\"name\":\"removeSupportTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"roleFlag\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_lockFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockFeeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unlockFeeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unlockFeeRatio\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportToken\",\"outputs\":[{\"internalType\":\"contractIToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"txMinted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"unlockFeeRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// CalculateFee is a free data retrieval call binding the contract method 0x0324ef9c.
//
// Solidity: function calculateFee(address token, uint256 amount, uint256 crossType) view returns(uint256 feeAmount, uint256 remainAmount)
func (_Bridge *BridgeCaller) CalculateFee(opts *bind.CallOpts, token common.Address, amount *big.Int, crossType *big.Int) (struct {
	FeeAmount    *big.Int
	RemainAmount *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "calculateFee", token, amount, crossType)

	outstruct := new(struct {
		FeeAmount    *big.Int
		RemainAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.RemainAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateFee is a free data retrieval call binding the contract method 0x0324ef9c.
//
// Solidity: function calculateFee(address token, uint256 amount, uint256 crossType) view returns(uint256 feeAmount, uint256 remainAmount)
func (_Bridge *BridgeSession) CalculateFee(token common.Address, amount *big.Int, crossType *big.Int) (struct {
	FeeAmount    *big.Int
	RemainAmount *big.Int
}, error) {
	return _Bridge.Contract.CalculateFee(&_Bridge.CallOpts, token, amount, crossType)
}

// CalculateFee is a free data retrieval call binding the contract method 0x0324ef9c.
//
// Solidity: function calculateFee(address token, uint256 amount, uint256 crossType) view returns(uint256 feeAmount, uint256 remainAmount)
func (_Bridge *BridgeCallerSession) CalculateFee(token common.Address, amount *big.Int, crossType *big.Int) (struct {
	FeeAmount    *big.Int
	RemainAmount *big.Int
}, error) {
	return _Bridge.Contract.CalculateFee(&_Bridge.CallOpts, token, amount, crossType)
}

// FeeToLength is a free data retrieval call binding the contract method 0xc0f1abdd.
//
// Solidity: function feeToLength(address token) view returns(uint256 len)
func (_Bridge *BridgeCaller) FeeToLength(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "feeToLength", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeToLength is a free data retrieval call binding the contract method 0xc0f1abdd.
//
// Solidity: function feeToLength(address token) view returns(uint256 len)
func (_Bridge *BridgeSession) FeeToLength(token common.Address) (*big.Int, error) {
	return _Bridge.Contract.FeeToLength(&_Bridge.CallOpts, token)
}

// FeeToLength is a free data retrieval call binding the contract method 0xc0f1abdd.
//
// Solidity: function feeToLength(address token) view returns(uint256 len)
func (_Bridge *BridgeCallerSession) FeeToLength(token common.Address) (*big.Int, error) {
	return _Bridge.Contract.FeeToLength(&_Bridge.CallOpts, token)
}

// GetChainID is a free data retrieval call binding the contract method 0x564b81ef.
//
// Solidity: function getChainID() view returns(uint256)
func (_Bridge *BridgeCaller) GetChainID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getChainID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetChainID is a free data retrieval call binding the contract method 0x564b81ef.
//
// Solidity: function getChainID() view returns(uint256)
func (_Bridge *BridgeSession) GetChainID() (*big.Int, error) {
	return _Bridge.Contract.GetChainID(&_Bridge.CallOpts)
}

// GetChainID is a free data retrieval call binding the contract method 0x564b81ef.
//
// Solidity: function getChainID() view returns(uint256)
func (_Bridge *BridgeCallerSession) GetChainID() (*big.Int, error) {
	return _Bridge.Contract.GetChainID(&_Bridge.CallOpts)
}

// GetFeeTo is a free data retrieval call binding the contract method 0x5e08378c.
//
// Solidity: function getFeeTo(address token, uint256 i) view returns(address account)
func (_Bridge *BridgeCaller) GetFeeTo(opts *bind.CallOpts, token common.Address, i *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getFeeTo", token, i)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeTo is a free data retrieval call binding the contract method 0x5e08378c.
//
// Solidity: function getFeeTo(address token, uint256 i) view returns(address account)
func (_Bridge *BridgeSession) GetFeeTo(token common.Address, i *big.Int) (common.Address, error) {
	return _Bridge.Contract.GetFeeTo(&_Bridge.CallOpts, token, i)
}

// GetFeeTo is a free data retrieval call binding the contract method 0x5e08378c.
//
// Solidity: function getFeeTo(address token, uint256 i) view returns(address account)
func (_Bridge *BridgeCallerSession) GetFeeTo(token common.Address, i *big.Int) (common.Address, error) {
	return _Bridge.Contract.GetFeeTo(&_Bridge.CallOpts, token, i)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bridge.Contract.GetRoleMember(&_Bridge.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bridge.Contract.GetRoleMember(&_Bridge.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bridge.Contract.GetRoleMemberCount(&_Bridge.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bridge.Contract.GetRoleMemberCount(&_Bridge.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// LockFeeAmount is a free data retrieval call binding the contract method 0x71b47e18.
//
// Solidity: function lockFeeAmount(address ) view returns(uint256)
func (_Bridge *BridgeCaller) LockFeeAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "lockFeeAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockFeeAmount is a free data retrieval call binding the contract method 0x71b47e18.
//
// Solidity: function lockFeeAmount(address ) view returns(uint256)
func (_Bridge *BridgeSession) LockFeeAmount(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.LockFeeAmount(&_Bridge.CallOpts, arg0)
}

// LockFeeAmount is a free data retrieval call binding the contract method 0x71b47e18.
//
// Solidity: function lockFeeAmount(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) LockFeeAmount(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.LockFeeAmount(&_Bridge.CallOpts, arg0)
}

// LockFeeRatio is a free data retrieval call binding the contract method 0xf21b8af9.
//
// Solidity: function lockFeeRatio(address ) view returns(uint256)
func (_Bridge *BridgeCaller) LockFeeRatio(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "lockFeeRatio", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockFeeRatio is a free data retrieval call binding the contract method 0xf21b8af9.
//
// Solidity: function lockFeeRatio(address ) view returns(uint256)
func (_Bridge *BridgeSession) LockFeeRatio(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.LockFeeRatio(&_Bridge.CallOpts, arg0)
}

// LockFeeRatio is a free data retrieval call binding the contract method 0xf21b8af9.
//
// Solidity: function lockFeeRatio(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) LockFeeRatio(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.LockFeeRatio(&_Bridge.CallOpts, arg0)
}

// ProposalOf is a free data retrieval call binding the contract method 0xf48c0abf.
//
// Solidity: function proposalOf(bytes32 ) view returns(address tokenTo, address from, address to, uint256 amount, uint256 count, string txid, bool isFinished, bool isExist)
func (_Bridge *BridgeCaller) ProposalOf(opts *bind.CallOpts, arg0 [32]byte) (struct {
	TokenTo    common.Address
	From       common.Address
	To         common.Address
	Amount     *big.Int
	Count      *big.Int
	Txid       string
	IsFinished bool
	IsExist    bool
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "proposalOf", arg0)

	outstruct := new(struct {
		TokenTo    common.Address
		From       common.Address
		To         common.Address
		Amount     *big.Int
		Count      *big.Int
		Txid       string
		IsFinished bool
		IsExist    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenTo = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.From = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.To = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Count = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Txid = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.IsFinished = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.IsExist = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// ProposalOf is a free data retrieval call binding the contract method 0xf48c0abf.
//
// Solidity: function proposalOf(bytes32 ) view returns(address tokenTo, address from, address to, uint256 amount, uint256 count, string txid, bool isFinished, bool isExist)
func (_Bridge *BridgeSession) ProposalOf(arg0 [32]byte) (struct {
	TokenTo    common.Address
	From       common.Address
	To         common.Address
	Amount     *big.Int
	Count      *big.Int
	Txid       string
	IsFinished bool
	IsExist    bool
}, error) {
	return _Bridge.Contract.ProposalOf(&_Bridge.CallOpts, arg0)
}

// ProposalOf is a free data retrieval call binding the contract method 0xf48c0abf.
//
// Solidity: function proposalOf(bytes32 ) view returns(address tokenTo, address from, address to, uint256 amount, uint256 count, string txid, bool isFinished, bool isExist)
func (_Bridge *BridgeCallerSession) ProposalOf(arg0 [32]byte) (struct {
	TokenTo    common.Address
	From       common.Address
	To         common.Address
	Amount     *big.Int
	Count      *big.Int
	Txid       string
	IsFinished bool
	IsExist    bool
}, error) {
	return _Bridge.Contract.ProposalOf(&_Bridge.CallOpts, arg0)
}

// RoleFlag is a free data retrieval call binding the contract method 0x338a5663.
//
// Solidity: function roleFlag(address ) view returns(bytes32)
func (_Bridge *BridgeCaller) RoleFlag(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "roleFlag", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RoleFlag is a free data retrieval call binding the contract method 0x338a5663.
//
// Solidity: function roleFlag(address ) view returns(bytes32)
func (_Bridge *BridgeSession) RoleFlag(arg0 common.Address) ([32]byte, error) {
	return _Bridge.Contract.RoleFlag(&_Bridge.CallOpts, arg0)
}

// RoleFlag is a free data retrieval call binding the contract method 0x338a5663.
//
// Solidity: function roleFlag(address ) view returns(bytes32)
func (_Bridge *BridgeCallerSession) RoleFlag(arg0 common.Address) ([32]byte, error) {
	return _Bridge.Contract.RoleFlag(&_Bridge.CallOpts, arg0)
}

// SupportToken is a free data retrieval call binding the contract method 0x2a4f1621.
//
// Solidity: function supportToken(address ) view returns(address)
func (_Bridge *BridgeCaller) SupportToken(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "supportToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportToken is a free data retrieval call binding the contract method 0x2a4f1621.
//
// Solidity: function supportToken(address ) view returns(address)
func (_Bridge *BridgeSession) SupportToken(arg0 common.Address) (common.Address, error) {
	return _Bridge.Contract.SupportToken(&_Bridge.CallOpts, arg0)
}

// SupportToken is a free data retrieval call binding the contract method 0x2a4f1621.
//
// Solidity: function supportToken(address ) view returns(address)
func (_Bridge *BridgeCallerSession) SupportToken(arg0 common.Address) (common.Address, error) {
	return _Bridge.Contract.SupportToken(&_Bridge.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_Bridge *BridgeCaller) Threshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "threshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_Bridge *BridgeSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.Threshold(&_Bridge.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.Threshold(&_Bridge.CallOpts, arg0)
}

// TxMinted is a free data retrieval call binding the contract method 0x10c27402.
//
// Solidity: function txMinted(string ) view returns(bool)
func (_Bridge *BridgeCaller) TxMinted(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "txMinted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TxMinted is a free data retrieval call binding the contract method 0x10c27402.
//
// Solidity: function txMinted(string ) view returns(bool)
func (_Bridge *BridgeSession) TxMinted(arg0 string) (bool, error) {
	return _Bridge.Contract.TxMinted(&_Bridge.CallOpts, arg0)
}

// TxMinted is a free data retrieval call binding the contract method 0x10c27402.
//
// Solidity: function txMinted(string ) view returns(bool)
func (_Bridge *BridgeCallerSession) TxMinted(arg0 string) (bool, error) {
	return _Bridge.Contract.TxMinted(&_Bridge.CallOpts, arg0)
}

// UnlockFeeAmount is a free data retrieval call binding the contract method 0x09db763e.
//
// Solidity: function unlockFeeAmount(address ) view returns(uint256)
func (_Bridge *BridgeCaller) UnlockFeeAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "unlockFeeAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockFeeAmount is a free data retrieval call binding the contract method 0x09db763e.
//
// Solidity: function unlockFeeAmount(address ) view returns(uint256)
func (_Bridge *BridgeSession) UnlockFeeAmount(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.UnlockFeeAmount(&_Bridge.CallOpts, arg0)
}

// UnlockFeeAmount is a free data retrieval call binding the contract method 0x09db763e.
//
// Solidity: function unlockFeeAmount(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) UnlockFeeAmount(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.UnlockFeeAmount(&_Bridge.CallOpts, arg0)
}

// UnlockFeeRatio is a free data retrieval call binding the contract method 0x12d365dd.
//
// Solidity: function unlockFeeRatio(address ) view returns(uint256)
func (_Bridge *BridgeCaller) UnlockFeeRatio(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "unlockFeeRatio", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnlockFeeRatio is a free data retrieval call binding the contract method 0x12d365dd.
//
// Solidity: function unlockFeeRatio(address ) view returns(uint256)
func (_Bridge *BridgeSession) UnlockFeeRatio(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.UnlockFeeRatio(&_Bridge.CallOpts, arg0)
}

// UnlockFeeRatio is a free data retrieval call binding the contract method 0x12d365dd.
//
// Solidity: function unlockFeeRatio(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) UnlockFeeRatio(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.UnlockFeeRatio(&_Bridge.CallOpts, arg0)
}

// AddFeeTo is a paid mutator transaction binding the contract method 0xa0052546.
//
// Solidity: function addFeeTo(address token0, address account) returns()
func (_Bridge *BridgeTransactor) AddFeeTo(opts *bind.TransactOpts, token0 common.Address, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addFeeTo", token0, account)
}

// AddFeeTo is a paid mutator transaction binding the contract method 0xa0052546.
//
// Solidity: function addFeeTo(address token0, address account) returns()
func (_Bridge *BridgeSession) AddFeeTo(token0 common.Address, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddFeeTo(&_Bridge.TransactOpts, token0, account)
}

// AddFeeTo is a paid mutator transaction binding the contract method 0xa0052546.
//
// Solidity: function addFeeTo(address token0, address account) returns()
func (_Bridge *BridgeTransactorSession) AddFeeTo(token0 common.Address, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AddFeeTo(&_Bridge.TransactOpts, token0, account)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x013e91e6.
//
// Solidity: function addSupportToken(address token0, address token1, bytes32 _roleFlag) returns()
func (_Bridge *BridgeTransactor) AddSupportToken(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, _roleFlag [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addSupportToken", token0, token1, _roleFlag)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x013e91e6.
//
// Solidity: function addSupportToken(address token0, address token1, bytes32 _roleFlag) returns()
func (_Bridge *BridgeSession) AddSupportToken(token0 common.Address, token1 common.Address, _roleFlag [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.AddSupportToken(&_Bridge.TransactOpts, token0, token1, _roleFlag)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x013e91e6.
//
// Solidity: function addSupportToken(address token0, address token1, bytes32 _roleFlag) returns()
func (_Bridge *BridgeTransactorSession) AddSupportToken(token0 common.Address, token1 common.Address, _roleFlag [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.AddSupportToken(&_Bridge.TransactOpts, token0, token1, _roleFlag)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0x00964fff.
//
// Solidity: function addSupportTokens(address[] token0Addrs, address[] token1Addrs, bytes32[] _roleFlags) returns()
func (_Bridge *BridgeTransactor) AddSupportTokens(opts *bind.TransactOpts, token0Addrs []common.Address, token1Addrs []common.Address, _roleFlags [][32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "addSupportTokens", token0Addrs, token1Addrs, _roleFlags)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0x00964fff.
//
// Solidity: function addSupportTokens(address[] token0Addrs, address[] token1Addrs, bytes32[] _roleFlags) returns()
func (_Bridge *BridgeSession) AddSupportTokens(token0Addrs []common.Address, token1Addrs []common.Address, _roleFlags [][32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.AddSupportTokens(&_Bridge.TransactOpts, token0Addrs, token1Addrs, _roleFlags)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0x00964fff.
//
// Solidity: function addSupportTokens(address[] token0Addrs, address[] token1Addrs, bytes32[] _roleFlags) returns()
func (_Bridge *BridgeTransactorSession) AddSupportTokens(token0Addrs []common.Address, token1Addrs []common.Address, _roleFlags [][32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.AddSupportTokens(&_Bridge.TransactOpts, token0Addrs, token1Addrs, _roleFlags)
}

// CrossBurn is a paid mutator transaction binding the contract method 0xcbde3c73.
//
// Solidity: function crossBurn(address token0, address to, uint256 amount) returns()
func (_Bridge *BridgeTransactor) CrossBurn(opts *bind.TransactOpts, token0 common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "crossBurn", token0, to, amount)
}

// CrossBurn is a paid mutator transaction binding the contract method 0xcbde3c73.
//
// Solidity: function crossBurn(address token0, address to, uint256 amount) returns()
func (_Bridge *BridgeSession) CrossBurn(token0 common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.CrossBurn(&_Bridge.TransactOpts, token0, to, amount)
}

// CrossBurn is a paid mutator transaction binding the contract method 0xcbde3c73.
//
// Solidity: function crossBurn(address token0, address to, uint256 amount) returns()
func (_Bridge *BridgeTransactorSession) CrossBurn(token0 common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.CrossBurn(&_Bridge.TransactOpts, token0, to, amount)
}

// CrossMint is a paid mutator transaction binding the contract method 0x98a7582f.
//
// Solidity: function crossMint(address token0, address from, address to, uint256 amount, string txid) returns()
func (_Bridge *BridgeTransactor) CrossMint(opts *bind.TransactOpts, token0 common.Address, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "crossMint", token0, from, to, amount, txid)
}

// CrossMint is a paid mutator transaction binding the contract method 0x98a7582f.
//
// Solidity: function crossMint(address token0, address from, address to, uint256 amount, string txid) returns()
func (_Bridge *BridgeSession) CrossMint(token0 common.Address, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _Bridge.Contract.CrossMint(&_Bridge.TransactOpts, token0, from, to, amount, txid)
}

// CrossMint is a paid mutator transaction binding the contract method 0x98a7582f.
//
// Solidity: function crossMint(address token0, address from, address to, uint256 amount, string txid) returns()
func (_Bridge *BridgeTransactorSession) CrossMint(token0 common.Address, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _Bridge.Contract.CrossMint(&_Bridge.TransactOpts, token0, from, to, amount, txid)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// RemoveFeeTo is a paid mutator transaction binding the contract method 0x7b58453d.
//
// Solidity: function removeFeeTo(address token0, address account) returns()
func (_Bridge *BridgeTransactor) RemoveFeeTo(opts *bind.TransactOpts, token0 common.Address, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeFeeTo", token0, account)
}

// RemoveFeeTo is a paid mutator transaction binding the contract method 0x7b58453d.
//
// Solidity: function removeFeeTo(address token0, address account) returns()
func (_Bridge *BridgeSession) RemoveFeeTo(token0 common.Address, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveFeeTo(&_Bridge.TransactOpts, token0, account)
}

// RemoveFeeTo is a paid mutator transaction binding the contract method 0x7b58453d.
//
// Solidity: function removeFeeTo(address token0, address account) returns()
func (_Bridge *BridgeTransactorSession) RemoveFeeTo(token0 common.Address, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveFeeTo(&_Bridge.TransactOpts, token0, account)
}

// RemoveRoleFlag is a paid mutator transaction binding the contract method 0x73338918.
//
// Solidity: function removeRoleFlag(address token0) returns()
func (_Bridge *BridgeTransactor) RemoveRoleFlag(opts *bind.TransactOpts, token0 common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeRoleFlag", token0)
}

// RemoveRoleFlag is a paid mutator transaction binding the contract method 0x73338918.
//
// Solidity: function removeRoleFlag(address token0) returns()
func (_Bridge *BridgeSession) RemoveRoleFlag(token0 common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveRoleFlag(&_Bridge.TransactOpts, token0)
}

// RemoveRoleFlag is a paid mutator transaction binding the contract method 0x73338918.
//
// Solidity: function removeRoleFlag(address token0) returns()
func (_Bridge *BridgeTransactorSession) RemoveRoleFlag(token0 common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveRoleFlag(&_Bridge.TransactOpts, token0)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xe2769cfa.
//
// Solidity: function removeSupportToken(address token0) returns()
func (_Bridge *BridgeTransactor) RemoveSupportToken(opts *bind.TransactOpts, token0 common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeSupportToken", token0)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xe2769cfa.
//
// Solidity: function removeSupportToken(address token0) returns()
func (_Bridge *BridgeSession) RemoveSupportToken(token0 common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveSupportToken(&_Bridge.TransactOpts, token0)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xe2769cfa.
//
// Solidity: function removeSupportToken(address token0) returns()
func (_Bridge *BridgeTransactorSession) RemoveSupportToken(token0 common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveSupportToken(&_Bridge.TransactOpts, token0)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x0daff621.
//
// Solidity: function removeSupportTokens(address[] addrs) returns()
func (_Bridge *BridgeTransactor) RemoveSupportTokens(opts *bind.TransactOpts, addrs []common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "removeSupportTokens", addrs)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x0daff621.
//
// Solidity: function removeSupportTokens(address[] addrs) returns()
func (_Bridge *BridgeSession) RemoveSupportTokens(addrs []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveSupportTokens(&_Bridge.TransactOpts, addrs)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x0daff621.
//
// Solidity: function removeSupportTokens(address[] addrs) returns()
func (_Bridge *BridgeTransactorSession) RemoveSupportTokens(addrs []common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RemoveSupportTokens(&_Bridge.TransactOpts, addrs)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// SetFee is a paid mutator transaction binding the contract method 0xd1346b54.
//
// Solidity: function setFee(address token0, uint256 _lockFeeAmount, uint256 _lockFeeRatio, uint256 _unlockFeeAmount, uint256 _unlockFeeRatio) returns()
func (_Bridge *BridgeTransactor) SetFee(opts *bind.TransactOpts, token0 common.Address, _lockFeeAmount *big.Int, _lockFeeRatio *big.Int, _unlockFeeAmount *big.Int, _unlockFeeRatio *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setFee", token0, _lockFeeAmount, _lockFeeRatio, _unlockFeeAmount, _unlockFeeRatio)
}

// SetFee is a paid mutator transaction binding the contract method 0xd1346b54.
//
// Solidity: function setFee(address token0, uint256 _lockFeeAmount, uint256 _lockFeeRatio, uint256 _unlockFeeAmount, uint256 _unlockFeeRatio) returns()
func (_Bridge *BridgeSession) SetFee(token0 common.Address, _lockFeeAmount *big.Int, _lockFeeRatio *big.Int, _unlockFeeAmount *big.Int, _unlockFeeRatio *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetFee(&_Bridge.TransactOpts, token0, _lockFeeAmount, _lockFeeRatio, _unlockFeeAmount, _unlockFeeRatio)
}

// SetFee is a paid mutator transaction binding the contract method 0xd1346b54.
//
// Solidity: function setFee(address token0, uint256 _lockFeeAmount, uint256 _lockFeeRatio, uint256 _unlockFeeAmount, uint256 _unlockFeeRatio) returns()
func (_Bridge *BridgeTransactorSession) SetFee(token0 common.Address, _lockFeeAmount *big.Int, _lockFeeRatio *big.Int, _unlockFeeAmount *big.Int, _unlockFeeRatio *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetFee(&_Bridge.TransactOpts, token0, _lockFeeAmount, _lockFeeRatio, _unlockFeeAmount, _unlockFeeRatio)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token1, uint256 _threshold) returns()
func (_Bridge *BridgeTransactor) SetThreshold(opts *bind.TransactOpts, token1 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setThreshold", token1, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token1, uint256 _threshold) returns()
func (_Bridge *BridgeSession) SetThreshold(token1 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetThreshold(&_Bridge.TransactOpts, token1, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token1, uint256 _threshold) returns()
func (_Bridge *BridgeTransactorSession) SetThreshold(token1 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.SetThreshold(&_Bridge.TransactOpts, token1, _threshold)
}

// BridgeCrossBurnIterator is returned from FilterCrossBurn and is used to iterate over the raw logs and unpacked data for CrossBurn events raised by the Bridge contract.
type BridgeCrossBurnIterator struct {
	Event *BridgeCrossBurn // Event containing the contract specifics and raw log

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
func (it *BridgeCrossBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossBurn)
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
		it.Event = new(BridgeCrossBurn)
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
func (it *BridgeCrossBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossBurn represents a CrossBurn event raised by the Bridge contract.
type BridgeCrossBurn struct {
	Token0  common.Address
	Token1  common.Address
	ChainID *big.Int
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossBurn is a free log retrieval operation binding the contract event 0x0c0ccff1e07d6497e77b2c10d52915eab0e2ba09bb3169ce6f1b90653ca99a87.
//
// Solidity: event CrossBurn(address token0, address token1, uint256 chainID, address from, address to, uint256 amount)
func (_Bridge *BridgeFilterer) FilterCrossBurn(opts *bind.FilterOpts) (*BridgeCrossBurnIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "CrossBurn")
	if err != nil {
		return nil, err
	}
	return &BridgeCrossBurnIterator{contract: _Bridge.contract, event: "CrossBurn", logs: logs, sub: sub}, nil
}

// WatchCrossBurn is a free log subscription operation binding the contract event 0x0c0ccff1e07d6497e77b2c10d52915eab0e2ba09bb3169ce6f1b90653ca99a87.
//
// Solidity: event CrossBurn(address token0, address token1, uint256 chainID, address from, address to, uint256 amount)
func (_Bridge *BridgeFilterer) WatchCrossBurn(opts *bind.WatchOpts, sink chan<- *BridgeCrossBurn) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "CrossBurn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossBurn)
				if err := _Bridge.contract.UnpackLog(event, "CrossBurn", log); err != nil {
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

// ParseCrossBurn is a log parse operation binding the contract event 0x0c0ccff1e07d6497e77b2c10d52915eab0e2ba09bb3169ce6f1b90653ca99a87.
//
// Solidity: event CrossBurn(address token0, address token1, uint256 chainID, address from, address to, uint256 amount)
func (_Bridge *BridgeFilterer) ParseCrossBurn(log types.Log) (*BridgeCrossBurn, error) {
	event := new(BridgeCrossBurn)
	if err := _Bridge.contract.UnpackLog(event, "CrossBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossMintIterator is returned from FilterCrossMint and is used to iterate over the raw logs and unpacked data for CrossMint events raised by the Bridge contract.
type BridgeCrossMintIterator struct {
	Event *BridgeCrossMint // Event containing the contract specifics and raw log

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
func (it *BridgeCrossMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossMint)
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
		it.Event = new(BridgeCrossMint)
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
func (it *BridgeCrossMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossMint represents a CrossMint event raised by the Bridge contract.
type BridgeCrossMint struct {
	Token0  common.Address
	Token1  common.Address
	ChainID *big.Int
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Txid    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCrossMint is a free log retrieval operation binding the contract event 0x528f691cb4ff9535caaf69e54c50406adf913ba29b79a3e2dc3d865a3ef7ae41.
//
// Solidity: event CrossMint(address token0, address token1, uint256 chainID, address from, address to, uint256 amount, string txid)
func (_Bridge *BridgeFilterer) FilterCrossMint(opts *bind.FilterOpts) (*BridgeCrossMintIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "CrossMint")
	if err != nil {
		return nil, err
	}
	return &BridgeCrossMintIterator{contract: _Bridge.contract, event: "CrossMint", logs: logs, sub: sub}, nil
}

// WatchCrossMint is a free log subscription operation binding the contract event 0x528f691cb4ff9535caaf69e54c50406adf913ba29b79a3e2dc3d865a3ef7ae41.
//
// Solidity: event CrossMint(address token0, address token1, uint256 chainID, address from, address to, uint256 amount, string txid)
func (_Bridge *BridgeFilterer) WatchCrossMint(opts *bind.WatchOpts, sink chan<- *BridgeCrossMint) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "CrossMint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossMint)
				if err := _Bridge.contract.UnpackLog(event, "CrossMint", log); err != nil {
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

// ParseCrossMint is a log parse operation binding the contract event 0x528f691cb4ff9535caaf69e54c50406adf913ba29b79a3e2dc3d865a3ef7ae41.
//
// Solidity: event CrossMint(address token0, address token1, uint256 chainID, address from, address to, uint256 amount, string txid)
func (_Bridge *BridgeFilterer) ParseCrossMint(log types.Log) (*BridgeCrossMint, error) {
	event := new(BridgeCrossMint)
	if err := _Bridge.contract.UnpackLog(event, "CrossMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeChangeIterator is returned from FilterFeeChange and is used to iterate over the raw logs and unpacked data for FeeChange events raised by the Bridge contract.
type BridgeFeeChangeIterator struct {
	Event *BridgeFeeChange // Event containing the contract specifics and raw log

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
func (it *BridgeFeeChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeChange)
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
		it.Event = new(BridgeFeeChange)
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
func (it *BridgeFeeChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeChange represents a FeeChange event raised by the Bridge contract.
type BridgeFeeChange struct {
	Token           common.Address
	LockFeeAmount   *big.Int
	LockFeeRatio    *big.Int
	UnlockFeeAmount *big.Int
	UnlockFeeRatio  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFeeChange is a free log retrieval operation binding the contract event 0x0ada58b95297d7845bc4aec0fd7ccab1aebd1ed4bebdc0902741bb9d193ef640.
//
// Solidity: event FeeChange(address token, uint256 lockFeeAmount, uint256 lockFeeRatio, uint256 unlockFeeAmount, uint256 unlockFeeRatio)
func (_Bridge *BridgeFilterer) FilterFeeChange(opts *bind.FilterOpts) (*BridgeFeeChangeIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "FeeChange")
	if err != nil {
		return nil, err
	}
	return &BridgeFeeChangeIterator{contract: _Bridge.contract, event: "FeeChange", logs: logs, sub: sub}, nil
}

// WatchFeeChange is a free log subscription operation binding the contract event 0x0ada58b95297d7845bc4aec0fd7ccab1aebd1ed4bebdc0902741bb9d193ef640.
//
// Solidity: event FeeChange(address token, uint256 lockFeeAmount, uint256 lockFeeRatio, uint256 unlockFeeAmount, uint256 unlockFeeRatio)
func (_Bridge *BridgeFilterer) WatchFeeChange(opts *bind.WatchOpts, sink chan<- *BridgeFeeChange) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "FeeChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeChange)
				if err := _Bridge.contract.UnpackLog(event, "FeeChange", log); err != nil {
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

// ParseFeeChange is a log parse operation binding the contract event 0x0ada58b95297d7845bc4aec0fd7ccab1aebd1ed4bebdc0902741bb9d193ef640.
//
// Solidity: event FeeChange(address token, uint256 lockFeeAmount, uint256 lockFeeRatio, uint256 unlockFeeAmount, uint256 unlockFeeRatio)
func (_Bridge *BridgeFilterer) ParseFeeChange(log types.Log) (*BridgeFeeChange, error) {
	event := new(BridgeFeeChange)
	if err := _Bridge.contract.UnpackLog(event, "FeeChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeToAddedIterator is returned from FilterFeeToAdded and is used to iterate over the raw logs and unpacked data for FeeToAdded events raised by the Bridge contract.
type BridgeFeeToAddedIterator struct {
	Event *BridgeFeeToAdded // Event containing the contract specifics and raw log

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
func (it *BridgeFeeToAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeToAdded)
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
		it.Event = new(BridgeFeeToAdded)
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
func (it *BridgeFeeToAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeToAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeToAdded represents a FeeToAdded event raised by the Bridge contract.
type BridgeFeeToAdded struct {
	Token   common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeToAdded is a free log retrieval operation binding the contract event 0xaab99c312372d476211a5508724bd953c6cffb12ea46f5b1a9a38b6091eb616c.
//
// Solidity: event FeeToAdded(address token, address account)
func (_Bridge *BridgeFilterer) FilterFeeToAdded(opts *bind.FilterOpts) (*BridgeFeeToAddedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "FeeToAdded")
	if err != nil {
		return nil, err
	}
	return &BridgeFeeToAddedIterator{contract: _Bridge.contract, event: "FeeToAdded", logs: logs, sub: sub}, nil
}

// WatchFeeToAdded is a free log subscription operation binding the contract event 0xaab99c312372d476211a5508724bd953c6cffb12ea46f5b1a9a38b6091eb616c.
//
// Solidity: event FeeToAdded(address token, address account)
func (_Bridge *BridgeFilterer) WatchFeeToAdded(opts *bind.WatchOpts, sink chan<- *BridgeFeeToAdded) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "FeeToAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeToAdded)
				if err := _Bridge.contract.UnpackLog(event, "FeeToAdded", log); err != nil {
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

// ParseFeeToAdded is a log parse operation binding the contract event 0xaab99c312372d476211a5508724bd953c6cffb12ea46f5b1a9a38b6091eb616c.
//
// Solidity: event FeeToAdded(address token, address account)
func (_Bridge *BridgeFilterer) ParseFeeToAdded(log types.Log) (*BridgeFeeToAdded, error) {
	event := new(BridgeFeeToAdded)
	if err := _Bridge.contract.UnpackLog(event, "FeeToAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeFeeToRemovedIterator is returned from FilterFeeToRemoved and is used to iterate over the raw logs and unpacked data for FeeToRemoved events raised by the Bridge contract.
type BridgeFeeToRemovedIterator struct {
	Event *BridgeFeeToRemoved // Event containing the contract specifics and raw log

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
func (it *BridgeFeeToRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeFeeToRemoved)
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
		it.Event = new(BridgeFeeToRemoved)
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
func (it *BridgeFeeToRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeFeeToRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeFeeToRemoved represents a FeeToRemoved event raised by the Bridge contract.
type BridgeFeeToRemoved struct {
	Token   common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeToRemoved is a free log retrieval operation binding the contract event 0x1162c56e1face129565df34ecfdd3ca4f645d76f083be3a48f5d253b2b490b3a.
//
// Solidity: event FeeToRemoved(address token, address account)
func (_Bridge *BridgeFilterer) FilterFeeToRemoved(opts *bind.FilterOpts) (*BridgeFeeToRemovedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "FeeToRemoved")
	if err != nil {
		return nil, err
	}
	return &BridgeFeeToRemovedIterator{contract: _Bridge.contract, event: "FeeToRemoved", logs: logs, sub: sub}, nil
}

// WatchFeeToRemoved is a free log subscription operation binding the contract event 0x1162c56e1face129565df34ecfdd3ca4f645d76f083be3a48f5d253b2b490b3a.
//
// Solidity: event FeeToRemoved(address token, address account)
func (_Bridge *BridgeFilterer) WatchFeeToRemoved(opts *bind.WatchOpts, sink chan<- *BridgeFeeToRemoved) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "FeeToRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeFeeToRemoved)
				if err := _Bridge.contract.UnpackLog(event, "FeeToRemoved", log); err != nil {
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

// ParseFeeToRemoved is a log parse operation binding the contract event 0x1162c56e1face129565df34ecfdd3ca4f645d76f083be3a48f5d253b2b490b3a.
//
// Solidity: event FeeToRemoved(address token, address account)
func (_Bridge *BridgeFilterer) ParseFeeToRemoved(log types.Log) (*BridgeFeeToRemoved, error) {
	event := new(BridgeFeeToRemoved)
	if err := _Bridge.contract.UnpackLog(event, "FeeToRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the Bridge contract.
type BridgeProposalVotedIterator struct {
	Event *BridgeProposalVoted // Event containing the contract specifics and raw log

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
func (it *BridgeProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeProposalVoted)
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
		it.Event = new(BridgeProposalVoted)
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
func (it *BridgeProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeProposalVoted represents a ProposalVoted event raised by the Bridge contract.
type BridgeProposalVoted struct {
	Token     common.Address
	From      common.Address
	To        common.Address
	Amount    *big.Int
	Proposer  common.Address
	Count     *big.Int
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterProposalVoted is a free log retrieval operation binding the contract event 0xe458477b94285c9f254a407e2614c39f55230b527ff7f51d76fdf33e251ae936.
//
// Solidity: event ProposalVoted(address token, address from, address to, uint256 amount, address proposer, uint256 count, uint256 threshold)
func (_Bridge *BridgeFilterer) FilterProposalVoted(opts *bind.FilterOpts) (*BridgeProposalVotedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return &BridgeProposalVotedIterator{contract: _Bridge.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0xe458477b94285c9f254a407e2614c39f55230b527ff7f51d76fdf33e251ae936.
//
// Solidity: event ProposalVoted(address token, address from, address to, uint256 amount, address proposer, uint256 count, uint256 threshold)
func (_Bridge *BridgeFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *BridgeProposalVoted) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeProposalVoted)
				if err := _Bridge.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
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

// ParseProposalVoted is a log parse operation binding the contract event 0xe458477b94285c9f254a407e2614c39f55230b527ff7f51d76fdf33e251ae936.
//
// Solidity: event ProposalVoted(address token, address from, address to, uint256 amount, address proposer, uint256 count, uint256 threshold)
func (_Bridge *BridgeFilterer) ParseProposalVoted(log types.Log) (*BridgeProposalVoted, error) {
	event := new(BridgeProposalVoted)
	if err := _Bridge.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bridge contract.
type BridgeRoleAdminChangedIterator struct {
	Event *BridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleAdminChanged)
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
		it.Event = new(BridgeRoleAdminChanged)
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
func (it *BridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleAdminChanged represents a RoleAdminChanged event raised by the Bridge contract.
type BridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleAdminChangedIterator{contract: _Bridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleAdminChanged)
				if err := _Bridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeRoleAdminChanged, error) {
	event := new(BridgeRoleAdminChanged)
	if err := _Bridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bridge contract.
type BridgeRoleGrantedIterator struct {
	Event *BridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *BridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleGranted)
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
		it.Event = new(BridgeRoleGranted)
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
func (it *BridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleGranted represents a RoleGranted event raised by the Bridge contract.
type BridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleGrantedIterator{contract: _Bridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleGranted)
				if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseRoleGranted(log types.Log) (*BridgeRoleGranted, error) {
	event := new(BridgeRoleGranted)
	if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bridge contract.
type BridgeRoleRevokedIterator struct {
	Event *BridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleRevoked)
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
		it.Event = new(BridgeRoleRevoked)
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
func (it *BridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleRevoked represents a RoleRevoked event raised by the Bridge contract.
type BridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleRevokedIterator{contract: _Bridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleRevoked)
				if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseRoleRevoked(log types.Log) (*BridgeRoleRevoked, error) {
	event := new(BridgeRoleRevoked)
	if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the Bridge contract.
type BridgeThresholdChangedIterator struct {
	Event *BridgeThresholdChanged // Event containing the contract specifics and raw log

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
func (it *BridgeThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeThresholdChanged)
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
		it.Event = new(BridgeThresholdChanged)
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
func (it *BridgeThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeThresholdChanged represents a ThresholdChanged event raised by the Bridge contract.
type BridgeThresholdChanged struct {
	Token        common.Address
	OldThreshold *big.Int
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_Bridge *BridgeFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*BridgeThresholdChangedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &BridgeThresholdChangedIterator{contract: _Bridge.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_Bridge *BridgeFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *BridgeThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeThresholdChanged)
				if err := _Bridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseThresholdChanged(log types.Log) (*BridgeThresholdChanged, error) {
	event := new(BridgeThresholdChanged)
	if err := _Bridge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
