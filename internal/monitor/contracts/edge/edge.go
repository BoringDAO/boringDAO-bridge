// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package edge

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

// EdgeMetaData contains all meta data concerning the Edge contract.
var EdgeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addSupport\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"chainSupported\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"changeMultiSupport\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_chainIds\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"statuses\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"changeSupport\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"crossIn\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"internalType\":\"structInParam\",\"components\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"txid\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"crossOut\",\"inputs\":[{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"decimalDiff\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"eventHeights0\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eventHeights1\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eventIndex0\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"eventIndex1\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleKey\",\"inputs\":[{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isCoin\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"locked\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeSupportToken\",\"inputs\":[{\"name\":\"tokenAddr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsCoin\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isCoin\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setThreshold\",\"inputs\":[{\"name\":\"token0\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_threshold\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"threshold\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenSupported\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"txHandled\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"CoinSeted\",\"inputs\":[{\"name\":\"coin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CrossIned\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInParam\",\"components\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CrossOuted\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structOutParam\",\"components\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"toChainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"to\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposited\",\"inputs\":[{\"name\":\"fromChainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fromToken\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProposalVoted\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"from\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"to\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"count\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"threshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Supported\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ThresholdChanged\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"oldThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newThreshold\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// EdgeABI is the input ABI used to generate the binding from.
// Deprecated: Use EdgeMetaData.ABI instead.
var EdgeABI = EdgeMetaData.ABI

// Edge is an auto generated Go binding around an Ethereum contract.
type Edge struct {
	EdgeCaller     // Read-only binding to the contract
	EdgeTransactor // Write-only binding to the contract
	EdgeFilterer   // Log filterer for contract events
}

// EdgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type EdgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EdgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EdgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EdgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EdgeSession struct {
	Contract     *Edge             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EdgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EdgeCallerSession struct {
	Contract *EdgeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EdgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EdgeTransactorSession struct {
	Contract     *EdgeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EdgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type EdgeRaw struct {
	Contract *Edge // Generic contract binding to access the raw methods on
}

// EdgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EdgeCallerRaw struct {
	Contract *EdgeCaller // Generic read-only contract binding to access the raw methods on
}

// EdgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EdgeTransactorRaw struct {
	Contract *EdgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEdge creates a new instance of Edge, bound to a specific deployed contract.
func NewEdge(address common.Address, backend bind.ContractBackend) (*Edge, error) {
	contract, err := bindEdge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Edge{EdgeCaller: EdgeCaller{contract: contract}, EdgeTransactor: EdgeTransactor{contract: contract}, EdgeFilterer: EdgeFilterer{contract: contract}}, nil
}

// NewEdgeCaller creates a new read-only instance of Edge, bound to a specific deployed contract.
func NewEdgeCaller(address common.Address, caller bind.ContractCaller) (*EdgeCaller, error) {
	contract, err := bindEdge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EdgeCaller{contract: contract}, nil
}

// NewEdgeTransactor creates a new write-only instance of Edge, bound to a specific deployed contract.
func NewEdgeTransactor(address common.Address, transactor bind.ContractTransactor) (*EdgeTransactor, error) {
	contract, err := bindEdge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EdgeTransactor{contract: contract}, nil
}

// NewEdgeFilterer creates a new log filterer instance of Edge, bound to a specific deployed contract.
func NewEdgeFilterer(address common.Address, filterer bind.ContractFilterer) (*EdgeFilterer, error) {
	contract, err := bindEdge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EdgeFilterer{contract: contract}, nil
}

// bindEdge binds a generic wrapper to an already deployed contract.
func bindEdge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EdgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Edge *EdgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Edge.Contract.EdgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Edge *EdgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Edge.Contract.EdgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Edge *EdgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Edge.Contract.EdgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Edge *EdgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Edge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Edge *EdgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Edge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Edge *EdgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Edge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Edge *EdgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Edge *EdgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Edge.Contract.DEFAULTADMINROLE(&_Edge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Edge *EdgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Edge.Contract.DEFAULTADMINROLE(&_Edge.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Edge *EdgeCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Edge *EdgeSession) UPGRADERROLE() ([32]byte, error) {
	return _Edge.Contract.UPGRADERROLE(&_Edge.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_Edge *EdgeCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _Edge.Contract.UPGRADERROLE(&_Edge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Edge *EdgeCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Edge *EdgeSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Edge.Contract.UPGRADEINTERFACEVERSION(&_Edge.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_Edge *EdgeCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _Edge.Contract.UPGRADEINTERFACEVERSION(&_Edge.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Edge *EdgeCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Edge *EdgeSession) ChainId() (*big.Int, error) {
	return _Edge.Contract.ChainId(&_Edge.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_Edge *EdgeCallerSession) ChainId() (*big.Int, error) {
	return _Edge.Contract.ChainId(&_Edge.CallOpts)
}

// ChainSupported is a free data retrieval call binding the contract method 0xfe7a82e3.
//
// Solidity: function chainSupported(address , uint256 ) view returns(bool)
func (_Edge *EdgeCaller) ChainSupported(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "chainSupported", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ChainSupported is a free data retrieval call binding the contract method 0xfe7a82e3.
//
// Solidity: function chainSupported(address , uint256 ) view returns(bool)
func (_Edge *EdgeSession) ChainSupported(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Edge.Contract.ChainSupported(&_Edge.CallOpts, arg0, arg1)
}

// ChainSupported is a free data retrieval call binding the contract method 0xfe7a82e3.
//
// Solidity: function chainSupported(address , uint256 ) view returns(bool)
func (_Edge *EdgeCallerSession) ChainSupported(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Edge.Contract.ChainSupported(&_Edge.CallOpts, arg0, arg1)
}

// DecimalDiff is a free data retrieval call binding the contract method 0xbf814720.
//
// Solidity: function decimalDiff(address ) view returns(uint256)
func (_Edge *EdgeCaller) DecimalDiff(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "decimalDiff", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalDiff is a free data retrieval call binding the contract method 0xbf814720.
//
// Solidity: function decimalDiff(address ) view returns(uint256)
func (_Edge *EdgeSession) DecimalDiff(arg0 common.Address) (*big.Int, error) {
	return _Edge.Contract.DecimalDiff(&_Edge.CallOpts, arg0)
}

// DecimalDiff is a free data retrieval call binding the contract method 0xbf814720.
//
// Solidity: function decimalDiff(address ) view returns(uint256)
func (_Edge *EdgeCallerSession) DecimalDiff(arg0 common.Address) (*big.Int, error) {
	return _Edge.Contract.DecimalDiff(&_Edge.CallOpts, arg0)
}

// EventHeights0 is a free data retrieval call binding the contract method 0xfc1aa909.
//
// Solidity: function eventHeights0(uint256 , uint256 ) view returns(uint256)
func (_Edge *EdgeCaller) EventHeights0(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "eventHeights0", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventHeights0 is a free data retrieval call binding the contract method 0xfc1aa909.
//
// Solidity: function eventHeights0(uint256 , uint256 ) view returns(uint256)
func (_Edge *EdgeSession) EventHeights0(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Edge.Contract.EventHeights0(&_Edge.CallOpts, arg0, arg1)
}

// EventHeights0 is a free data retrieval call binding the contract method 0xfc1aa909.
//
// Solidity: function eventHeights0(uint256 , uint256 ) view returns(uint256)
func (_Edge *EdgeCallerSession) EventHeights0(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Edge.Contract.EventHeights0(&_Edge.CallOpts, arg0, arg1)
}

// EventHeights1 is a free data retrieval call binding the contract method 0x05e82506.
//
// Solidity: function eventHeights1(uint256 ) view returns(uint256)
func (_Edge *EdgeCaller) EventHeights1(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "eventHeights1", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventHeights1 is a free data retrieval call binding the contract method 0x05e82506.
//
// Solidity: function eventHeights1(uint256 ) view returns(uint256)
func (_Edge *EdgeSession) EventHeights1(arg0 *big.Int) (*big.Int, error) {
	return _Edge.Contract.EventHeights1(&_Edge.CallOpts, arg0)
}

// EventHeights1 is a free data retrieval call binding the contract method 0x05e82506.
//
// Solidity: function eventHeights1(uint256 ) view returns(uint256)
func (_Edge *EdgeCallerSession) EventHeights1(arg0 *big.Int) (*big.Int, error) {
	return _Edge.Contract.EventHeights1(&_Edge.CallOpts, arg0)
}

// EventIndex0 is a free data retrieval call binding the contract method 0x0826820e.
//
// Solidity: function eventIndex0(uint256 ) view returns(uint256)
func (_Edge *EdgeCaller) EventIndex0(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "eventIndex0", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventIndex0 is a free data retrieval call binding the contract method 0x0826820e.
//
// Solidity: function eventIndex0(uint256 ) view returns(uint256)
func (_Edge *EdgeSession) EventIndex0(arg0 *big.Int) (*big.Int, error) {
	return _Edge.Contract.EventIndex0(&_Edge.CallOpts, arg0)
}

// EventIndex0 is a free data retrieval call binding the contract method 0x0826820e.
//
// Solidity: function eventIndex0(uint256 ) view returns(uint256)
func (_Edge *EdgeCallerSession) EventIndex0(arg0 *big.Int) (*big.Int, error) {
	return _Edge.Contract.EventIndex0(&_Edge.CallOpts, arg0)
}

// EventIndex1 is a free data retrieval call binding the contract method 0x15cd2c70.
//
// Solidity: function eventIndex1() view returns(uint256)
func (_Edge *EdgeCaller) EventIndex1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "eventIndex1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventIndex1 is a free data retrieval call binding the contract method 0x15cd2c70.
//
// Solidity: function eventIndex1() view returns(uint256)
func (_Edge *EdgeSession) EventIndex1() (*big.Int, error) {
	return _Edge.Contract.EventIndex1(&_Edge.CallOpts)
}

// EventIndex1 is a free data retrieval call binding the contract method 0x15cd2c70.
//
// Solidity: function eventIndex1() view returns(uint256)
func (_Edge *EdgeCallerSession) EventIndex1() (*big.Int, error) {
	return _Edge.Contract.EventIndex1(&_Edge.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Edge *EdgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Edge *EdgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Edge.Contract.GetRoleAdmin(&_Edge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Edge *EdgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Edge.Contract.GetRoleAdmin(&_Edge.CallOpts, role)
}

// GetRoleKey is a free data retrieval call binding the contract method 0x3ab796e7.
//
// Solidity: function getRoleKey(address toToken) pure returns(bytes32 key)
func (_Edge *EdgeCaller) GetRoleKey(opts *bind.CallOpts, toToken common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "getRoleKey", toToken)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleKey is a free data retrieval call binding the contract method 0x3ab796e7.
//
// Solidity: function getRoleKey(address toToken) pure returns(bytes32 key)
func (_Edge *EdgeSession) GetRoleKey(toToken common.Address) ([32]byte, error) {
	return _Edge.Contract.GetRoleKey(&_Edge.CallOpts, toToken)
}

// GetRoleKey is a free data retrieval call binding the contract method 0x3ab796e7.
//
// Solidity: function getRoleKey(address toToken) pure returns(bytes32 key)
func (_Edge *EdgeCallerSession) GetRoleKey(toToken common.Address) ([32]byte, error) {
	return _Edge.Contract.GetRoleKey(&_Edge.CallOpts, toToken)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Edge *EdgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Edge *EdgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Edge.Contract.HasRole(&_Edge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Edge *EdgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Edge.Contract.HasRole(&_Edge.CallOpts, role, account)
}

// IsCoin is a free data retrieval call binding the contract method 0xa611033e.
//
// Solidity: function isCoin(address ) view returns(bool)
func (_Edge *EdgeCaller) IsCoin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "isCoin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCoin is a free data retrieval call binding the contract method 0xa611033e.
//
// Solidity: function isCoin(address ) view returns(bool)
func (_Edge *EdgeSession) IsCoin(arg0 common.Address) (bool, error) {
	return _Edge.Contract.IsCoin(&_Edge.CallOpts, arg0)
}

// IsCoin is a free data retrieval call binding the contract method 0xa611033e.
//
// Solidity: function isCoin(address ) view returns(bool)
func (_Edge *EdgeCallerSession) IsCoin(arg0 common.Address) (bool, error) {
	return _Edge.Contract.IsCoin(&_Edge.CallOpts, arg0)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Edge *EdgeCaller) Locked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "locked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Edge *EdgeSession) Locked() (bool, error) {
	return _Edge.Contract.Locked(&_Edge.CallOpts)
}

// Locked is a free data retrieval call binding the contract method 0xcf309012.
//
// Solidity: function locked() view returns(bool)
func (_Edge *EdgeCallerSession) Locked() (bool, error) {
	return _Edge.Contract.Locked(&_Edge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Edge *EdgeCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Edge *EdgeSession) ProxiableUUID() ([32]byte, error) {
	return _Edge.Contract.ProxiableUUID(&_Edge.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Edge *EdgeCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Edge.Contract.ProxiableUUID(&_Edge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Edge *EdgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Edge *EdgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Edge.Contract.SupportsInterface(&_Edge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Edge *EdgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Edge.Contract.SupportsInterface(&_Edge.CallOpts, interfaceId)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_Edge *EdgeCaller) Threshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "threshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_Edge *EdgeSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _Edge.Contract.Threshold(&_Edge.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_Edge *EdgeCallerSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _Edge.Contract.Threshold(&_Edge.CallOpts, arg0)
}

// TokenSupported is a free data retrieval call binding the contract method 0x062143f0.
//
// Solidity: function tokenSupported(address ) view returns(bool)
func (_Edge *EdgeCaller) TokenSupported(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "tokenSupported", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TokenSupported is a free data retrieval call binding the contract method 0x062143f0.
//
// Solidity: function tokenSupported(address ) view returns(bool)
func (_Edge *EdgeSession) TokenSupported(arg0 common.Address) (bool, error) {
	return _Edge.Contract.TokenSupported(&_Edge.CallOpts, arg0)
}

// TokenSupported is a free data retrieval call binding the contract method 0x062143f0.
//
// Solidity: function tokenSupported(address ) view returns(bool)
func (_Edge *EdgeCallerSession) TokenSupported(arg0 common.Address) (bool, error) {
	return _Edge.Contract.TokenSupported(&_Edge.CallOpts, arg0)
}

// TxHandled is a free data retrieval call binding the contract method 0x824df449.
//
// Solidity: function txHandled(string ) view returns(bool)
func (_Edge *EdgeCaller) TxHandled(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Edge.contract.Call(opts, &out, "txHandled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TxHandled is a free data retrieval call binding the contract method 0x824df449.
//
// Solidity: function txHandled(string ) view returns(bool)
func (_Edge *EdgeSession) TxHandled(arg0 string) (bool, error) {
	return _Edge.Contract.TxHandled(&_Edge.CallOpts, arg0)
}

// TxHandled is a free data retrieval call binding the contract method 0x824df449.
//
// Solidity: function txHandled(string ) view returns(bool)
func (_Edge *EdgeCallerSession) TxHandled(arg0 string) (bool, error) {
	return _Edge.Contract.TxHandled(&_Edge.CallOpts, arg0)
}

// AddSupport is a paid mutator transaction binding the contract method 0xbc201a9d.
//
// Solidity: function addSupport(address token) returns()
func (_Edge *EdgeTransactor) AddSupport(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Edge.contract.Transact(opts, "addSupport", token)
}

// AddSupport is a paid mutator transaction binding the contract method 0xbc201a9d.
//
// Solidity: function addSupport(address token) returns()
func (_Edge *EdgeSession) AddSupport(token common.Address) (*types.Transaction, error) {
	return _Edge.Contract.AddSupport(&_Edge.TransactOpts, token)
}

// AddSupport is a paid mutator transaction binding the contract method 0xbc201a9d.
//
// Solidity: function addSupport(address token) returns()
func (_Edge *EdgeTransactorSession) AddSupport(token common.Address) (*types.Transaction, error) {
	return _Edge.Contract.AddSupport(&_Edge.TransactOpts, token)
}

// ChangeMultiSupport is a paid mutator transaction binding the contract method 0x224b28fe.
//
// Solidity: function changeMultiSupport(address[] tokens, uint256[] _chainIds, bool[] statuses) returns()
func (_Edge *EdgeTransactor) ChangeMultiSupport(opts *bind.TransactOpts, tokens []common.Address, _chainIds []*big.Int, statuses []bool) (*types.Transaction, error) {
	return _Edge.contract.Transact(opts, "changeMultiSupport", tokens, _chainIds, statuses)
}

// ChangeMultiSupport is a paid mutator transaction binding the contract method 0x224b28fe.
//
// Solidity: function changeMultiSupport(address[] tokens, uint256[] _chainIds, bool[] statuses) returns()
func (_Edge *EdgeSession) ChangeMultiSupport(tokens []common.Address, _chainIds []*big.Int, statuses []bool) (*types.Transaction, error) {
	return _Edge.Contract.ChangeMultiSupport(&_Edge.TransactOpts, tokens, _chainIds, statuses)
}

// ChangeMultiSupport is a paid mutator transaction binding the contract method 0x224b28fe.
//
// Solidity: function changeMultiSupport(address[] tokens, uint256[] _chainIds, bool[] statuses) returns()
func (_Edge *EdgeTransactorSession) ChangeMultiSupport(tokens []common.Address, _chainIds []*big.Int, statuses []bool) (*types.Transaction, error) {
	return _Edge.Contract.ChangeMultiSupport(&_Edge.TransactOpts, tokens, _chainIds, statuses)
}

// ChangeSupport is a paid mutator transaction binding the contract method 0x5818c39e.
//
// Solidity: function changeSupport(address token, uint256 _chainId, bool status) returns()
func (_Edge *EdgeTransactor) ChangeSupport(opts *bind.TransactOpts, token common.Address, _chainId *big.Int, status bool) (*types.Transaction, error) {
	return _Edge.contract.Transact(opts, "changeSupport", token, _chainId, status)
}

// ChangeSupport is a paid mutator transaction binding the contract method 0x5818c39e.
//
// Solidity: function changeSupport(address token, uint256 _chainId, bool status) returns()
func (_Edge *EdgeSession) ChangeSupport(token common.Address, _chainId *big.Int, status bool) (*types.Transaction, error) {
	return _Edge.Contract.ChangeSupport(&_Edge.TransactOpts, token, _chainId, status)
}

// ChangeSupport is a paid mutator transaction binding the contract method 0x5818c39e.
//
// Solidity: function changeSupport(address token, uint256 _chainId, bool status) returns()
func (_Edge *EdgeTransactorSession) ChangeSupport(token common.Address, _chainId *big.Int, status bool) (*types.Transaction, error) {
	return _Edge.Contract.ChangeSupport(&_Edge.TransactOpts, token, _chainId, status)
}

// CrossIn is a paid mutator transaction binding the contract method 0xaab8d87e.
//
// Solidity: function crossIn((uint256,address,bytes,uint256,address,bytes,uint256) p, string txid) returns()
func (_Edge *EdgeTransactor) CrossIn(opts *bind.TransactOpts, p InParam, txid string) (*types.Transaction, error) {
	return _Edge.contract.Transact(opts, "crossIn", p, txid)
}

// CrossIn is a paid mutator transaction binding the contract method 0xaab8d87e.
//
// Solidity: function crossIn((uint256,address,bytes,uint256,address,bytes,uint256) p, string txid) returns()
func (_Edge *EdgeSession) CrossIn(p InParam, txid string) (*types.Transaction, error) {
	return _Edge.Contract.CrossIn(&_Edge.TransactOpts, p, txid)
}

// CrossIn is a paid mutator transaction binding the contract method 0xaab8d87e.
//
// Solidity: function crossIn((uint256,address,bytes,uint256,address,bytes,uint256) p, string txid) returns()
func (_Edge *EdgeTransactorSession) CrossIn(p InParam, txid string) (*types.Transaction, error) {
	return _Edge.Contract.CrossIn(&_Edge.TransactOpts, p, txid)
}

// CrossOut is a paid mutator transaction binding the contract method 0xe73d88d3.
//
// Solidity: function crossOut(address fromToken, uint256 toChainId, bytes to, uint256 amount) payable returns()
func (_Edge *EdgeTransactor) CrossOut(opts *bind.TransactOpts, fromToken common.Address, toChainId *big.Int, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _Edge.contract.Transact(opts, "crossOut", fromToken, toChainId, to, amount)
}

// CrossOut is a paid mutator transaction binding the contract method 0xe73d88d3.
//
// Solidity: function crossOut(address fromToken, uint256 toChainId, bytes to, uint256 amount) payable returns()
func (_Edge *EdgeSession) CrossOut(fromToken common.Address, toChainId *big.Int, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _Edge.Contract.CrossOut(&_Edge.TransactOpts, fromToken, toChainId, to, amount)
}

// CrossOut is a paid mutator transaction binding the contract method 0xe73d88d3.
//
// Solidity: function crossOut(address fromToken, uint256 toChainId, bytes to, uint256 amount) payable returns()
func (_Edge *EdgeTransactorSession) CrossOut(fromToken common.Address, toChainId *big.Int, to []byte, amount *big.Int) (*types.Transaction, error) {
	return _Edge.Contract.CrossOut(&_Edge.TransactOpts, fromToken, toChainId, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_Edge *EdgeTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Edge.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_Edge *EdgeSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Edge.Contract.Deposit(&_Edge.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_Edge *EdgeTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Edge.Contract.Deposit(&_Edge.TransactOpts, token, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Edge *EdgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Edge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Edge *EdgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Edge.Contract.GrantRole(&_Edge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Edge *EdgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Edge.Contract.GrantRole(&_Edge.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _chainId) returns()
func (_Edge *EdgeTransactor) Initialize(opts *bind.TransactOpts, _chainId *big.Int) (*types.Transaction, error) {
	return _Edge.contract.Transact(opts, "initialize", _chainId)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _chainId) returns()
func (_Edge *EdgeSession) Initialize(_chainId *big.Int) (*types.Transaction, error) {
	return _Edge.Contract.Initialize(&_Edge.TransactOpts, _chainId)
}

// Initialize is a paid mutator transaction binding the contract method 0xfe4b84df.
//
// Solidity: function initialize(uint256 _chainId) returns()
func (_Edge *EdgeTransactorSession) Initialize(_chainId *big.Int) (*types.Transaction, error) {
	return _Edge.Contract.Initialize(&_Edge.TransactOpts, _chainId)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xe2769cfa.
//
// Solidity: function removeSupportToken(address tokenAddr) returns()
func (_Edge *EdgeTransactor) RemoveSupportToken(opts *bind.TransactOpts, tokenAddr common.Address) (*types.Transaction, error) {
	return _Edge.contract.Transact(opts, "removeSupportToken", tokenAddr)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xe2769cfa.
//
// Solidity: function removeSupportToken(address tokenAddr) returns()
func (_Edge *EdgeSession) RemoveSupportToken(tokenAddr common.Address) (*types.Transaction, error) {
	return _Edge.Contract.RemoveSupportToken(&_Edge.TransactOpts, tokenAddr)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xe2769cfa.
//
// Solidity: function removeSupportToken(address tokenAddr) returns()
func (_Edge *EdgeTransactorSession) RemoveSupportToken(tokenAddr common.Address) (*types.Transaction, error) {
	return _Edge.Contract.RemoveSupportToken(&_Edge.TransactOpts, tokenAddr)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Edge *EdgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Edge.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Edge *EdgeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Edge.Contract.RenounceRole(&_Edge.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Edge *EdgeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Edge.Contract.RenounceRole(&_Edge.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Edge *EdgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Edge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Edge *EdgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Edge.Contract.RevokeRole(&_Edge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Edge *EdgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Edge.Contract.RevokeRole(&_Edge.TransactOpts, role, account)
}

// SetIsCoin is a paid mutator transaction binding the contract method 0xd93e76bb.
//
// Solidity: function setIsCoin(address token, bool _isCoin) returns()
func (_Edge *EdgeTransactor) SetIsCoin(opts *bind.TransactOpts, token common.Address, _isCoin bool) (*types.Transaction, error) {
	return _Edge.contract.Transact(opts, "setIsCoin", token, _isCoin)
}

// SetIsCoin is a paid mutator transaction binding the contract method 0xd93e76bb.
//
// Solidity: function setIsCoin(address token, bool _isCoin) returns()
func (_Edge *EdgeSession) SetIsCoin(token common.Address, _isCoin bool) (*types.Transaction, error) {
	return _Edge.Contract.SetIsCoin(&_Edge.TransactOpts, token, _isCoin)
}

// SetIsCoin is a paid mutator transaction binding the contract method 0xd93e76bb.
//
// Solidity: function setIsCoin(address token, bool _isCoin) returns()
func (_Edge *EdgeTransactorSession) SetIsCoin(token common.Address, _isCoin bool) (*types.Transaction, error) {
	return _Edge.Contract.SetIsCoin(&_Edge.TransactOpts, token, _isCoin)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_Edge *EdgeTransactor) SetThreshold(opts *bind.TransactOpts, token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Edge.contract.Transact(opts, "setThreshold", token0, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_Edge *EdgeSession) SetThreshold(token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Edge.Contract.SetThreshold(&_Edge.TransactOpts, token0, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_Edge *EdgeTransactorSession) SetThreshold(token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _Edge.Contract.SetThreshold(&_Edge.TransactOpts, token0, _threshold)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Edge *EdgeTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Edge.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Edge *EdgeSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Edge.Contract.UpgradeToAndCall(&_Edge.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Edge *EdgeTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Edge.Contract.UpgradeToAndCall(&_Edge.TransactOpts, newImplementation, data)
}

// EdgeCoinSetedIterator is returned from FilterCoinSeted and is used to iterate over the raw logs and unpacked data for CoinSeted events raised by the Edge contract.
type EdgeCoinSetedIterator struct {
	Event *EdgeCoinSeted // Event containing the contract specifics and raw log

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
func (it *EdgeCoinSetedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeCoinSeted)
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
		it.Event = new(EdgeCoinSeted)
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
func (it *EdgeCoinSetedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeCoinSetedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeCoinSeted represents a CoinSeted event raised by the Edge contract.
type EdgeCoinSeted struct {
	Coin common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterCoinSeted is a free log retrieval operation binding the contract event 0x068ea4c4aed0572edc3ba8586d929cd0d8022e8bc1794748e4ff980ce55c3292.
//
// Solidity: event CoinSeted(address coin)
func (_Edge *EdgeFilterer) FilterCoinSeted(opts *bind.FilterOpts) (*EdgeCoinSetedIterator, error) {

	logs, sub, err := _Edge.contract.FilterLogs(opts, "CoinSeted")
	if err != nil {
		return nil, err
	}
	return &EdgeCoinSetedIterator{contract: _Edge.contract, event: "CoinSeted", logs: logs, sub: sub}, nil
}

// WatchCoinSeted is a free log subscription operation binding the contract event 0x068ea4c4aed0572edc3ba8586d929cd0d8022e8bc1794748e4ff980ce55c3292.
//
// Solidity: event CoinSeted(address coin)
func (_Edge *EdgeFilterer) WatchCoinSeted(opts *bind.WatchOpts, sink chan<- *EdgeCoinSeted) (event.Subscription, error) {

	logs, sub, err := _Edge.contract.WatchLogs(opts, "CoinSeted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeCoinSeted)
				if err := _Edge.contract.UnpackLog(event, "CoinSeted", log); err != nil {
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
func (_Edge *EdgeFilterer) ParseCoinSeted(log types.Log) (*EdgeCoinSeted, error) {
	event := new(EdgeCoinSeted)
	if err := _Edge.contract.UnpackLog(event, "CoinSeted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeCrossInedIterator is returned from FilterCrossIned and is used to iterate over the raw logs and unpacked data for CrossIned events raised by the Edge contract.
type EdgeCrossInedIterator struct {
	Event *EdgeCrossIned // Event containing the contract specifics and raw log

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
func (it *EdgeCrossInedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeCrossIned)
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
		it.Event = new(EdgeCrossIned)
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
func (it *EdgeCrossInedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeCrossInedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeCrossIned represents a CrossIned event raised by the Edge contract.
type EdgeCrossIned struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCrossIned is a free log retrieval operation binding the contract event 0x186a37dd020447374a89e19ce3bc90447d34d9af95d221f64366bb54be38617b.
//
// Solidity: event CrossIned((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Edge *EdgeFilterer) FilterCrossIned(opts *bind.FilterOpts) (*EdgeCrossInedIterator, error) {

	logs, sub, err := _Edge.contract.FilterLogs(opts, "CrossIned")
	if err != nil {
		return nil, err
	}
	return &EdgeCrossInedIterator{contract: _Edge.contract, event: "CrossIned", logs: logs, sub: sub}, nil
}

// WatchCrossIned is a free log subscription operation binding the contract event 0x186a37dd020447374a89e19ce3bc90447d34d9af95d221f64366bb54be38617b.
//
// Solidity: event CrossIned((uint256,address,bytes,uint256,address,bytes,uint256) p)
func (_Edge *EdgeFilterer) WatchCrossIned(opts *bind.WatchOpts, sink chan<- *EdgeCrossIned) (event.Subscription, error) {

	logs, sub, err := _Edge.contract.WatchLogs(opts, "CrossIned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeCrossIned)
				if err := _Edge.contract.UnpackLog(event, "CrossIned", log); err != nil {
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
func (_Edge *EdgeFilterer) ParseCrossIned(log types.Log) (*EdgeCrossIned, error) {
	event := new(EdgeCrossIned)
	if err := _Edge.contract.UnpackLog(event, "CrossIned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeCrossOutedIterator is returned from FilterCrossOuted and is used to iterate over the raw logs and unpacked data for CrossOuted events raised by the Edge contract.
type EdgeCrossOutedIterator struct {
	Event *EdgeCrossOuted // Event containing the contract specifics and raw log

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
func (it *EdgeCrossOutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeCrossOuted)
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
		it.Event = new(EdgeCrossOuted)
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
func (it *EdgeCrossOutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeCrossOutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeCrossOuted represents a CrossOuted event raised by the Edge contract.
type EdgeCrossOuted struct {
	P   OutParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCrossOuted is a free log retrieval operation binding the contract event 0xd90ceb336939d3ee8cc4c59c89d040144a23f52a81732fc71b25f8fede43fa4d.
//
// Solidity: event CrossOuted((uint256,address,bytes,uint256,bytes,uint256) p)
func (_Edge *EdgeFilterer) FilterCrossOuted(opts *bind.FilterOpts) (*EdgeCrossOutedIterator, error) {

	logs, sub, err := _Edge.contract.FilterLogs(opts, "CrossOuted")
	if err != nil {
		return nil, err
	}
	return &EdgeCrossOutedIterator{contract: _Edge.contract, event: "CrossOuted", logs: logs, sub: sub}, nil
}

// WatchCrossOuted is a free log subscription operation binding the contract event 0xd90ceb336939d3ee8cc4c59c89d040144a23f52a81732fc71b25f8fede43fa4d.
//
// Solidity: event CrossOuted((uint256,address,bytes,uint256,bytes,uint256) p)
func (_Edge *EdgeFilterer) WatchCrossOuted(opts *bind.WatchOpts, sink chan<- *EdgeCrossOuted) (event.Subscription, error) {

	logs, sub, err := _Edge.contract.WatchLogs(opts, "CrossOuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeCrossOuted)
				if err := _Edge.contract.UnpackLog(event, "CrossOuted", log); err != nil {
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

// ParseCrossOuted is a log parse operation binding the contract event 0xd90ceb336939d3ee8cc4c59c89d040144a23f52a81732fc71b25f8fede43fa4d.
//
// Solidity: event CrossOuted((uint256,address,bytes,uint256,bytes,uint256) p)
func (_Edge *EdgeFilterer) ParseCrossOuted(log types.Log) (*EdgeCrossOuted, error) {
	event := new(EdgeCrossOuted)
	if err := _Edge.contract.UnpackLog(event, "CrossOuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Edge contract.
type EdgeDepositedIterator struct {
	Event *EdgeDeposited // Event containing the contract specifics and raw log

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
func (it *EdgeDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeDeposited)
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
		it.Event = new(EdgeDeposited)
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
func (it *EdgeDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeDeposited represents a Deposited event raised by the Edge contract.
type EdgeDeposited struct {
	FromChainId *big.Int
	FromToken   common.Address
	From        common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x984a71c9d95fd4794aeba33ae72edfec22053fde75488d63abef9dc69ee795af.
//
// Solidity: event Deposited(uint256 fromChainId, address fromToken, address from, uint256 amount)
func (_Edge *EdgeFilterer) FilterDeposited(opts *bind.FilterOpts) (*EdgeDepositedIterator, error) {

	logs, sub, err := _Edge.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &EdgeDepositedIterator{contract: _Edge.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x984a71c9d95fd4794aeba33ae72edfec22053fde75488d63abef9dc69ee795af.
//
// Solidity: event Deposited(uint256 fromChainId, address fromToken, address from, uint256 amount)
func (_Edge *EdgeFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *EdgeDeposited) (event.Subscription, error) {

	logs, sub, err := _Edge.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeDeposited)
				if err := _Edge.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_Edge *EdgeFilterer) ParseDeposited(log types.Log) (*EdgeDeposited, error) {
	event := new(EdgeDeposited)
	if err := _Edge.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Edge contract.
type EdgeInitializedIterator struct {
	Event *EdgeInitialized // Event containing the contract specifics and raw log

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
func (it *EdgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeInitialized)
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
		it.Event = new(EdgeInitialized)
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
func (it *EdgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeInitialized represents a Initialized event raised by the Edge contract.
type EdgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Edge *EdgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*EdgeInitializedIterator, error) {

	logs, sub, err := _Edge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &EdgeInitializedIterator{contract: _Edge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Edge *EdgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *EdgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Edge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeInitialized)
				if err := _Edge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Edge *EdgeFilterer) ParseInitialized(log types.Log) (*EdgeInitialized, error) {
	event := new(EdgeInitialized)
	if err := _Edge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the Edge contract.
type EdgeProposalVotedIterator struct {
	Event *EdgeProposalVoted // Event containing the contract specifics and raw log

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
func (it *EdgeProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeProposalVoted)
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
		it.Event = new(EdgeProposalVoted)
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
func (it *EdgeProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeProposalVoted represents a ProposalVoted event raised by the Edge contract.
type EdgeProposalVoted struct {
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
func (_Edge *EdgeFilterer) FilterProposalVoted(opts *bind.FilterOpts) (*EdgeProposalVotedIterator, error) {

	logs, sub, err := _Edge.contract.FilterLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return &EdgeProposalVotedIterator{contract: _Edge.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0xaa8d9692f9529994d8599d4a8aad3f9cb2e1290ffe484cf531edfb62de17258c.
//
// Solidity: event ProposalVoted(address token, bytes from, bytes to, uint256 amount, address proposer, uint256 count, uint256 threshold)
func (_Edge *EdgeFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *EdgeProposalVoted) (event.Subscription, error) {

	logs, sub, err := _Edge.contract.WatchLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeProposalVoted)
				if err := _Edge.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
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
func (_Edge *EdgeFilterer) ParseProposalVoted(log types.Log) (*EdgeProposalVoted, error) {
	event := new(EdgeProposalVoted)
	if err := _Edge.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Edge contract.
type EdgeRoleAdminChangedIterator struct {
	Event *EdgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *EdgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeRoleAdminChanged)
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
		it.Event = new(EdgeRoleAdminChanged)
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
func (it *EdgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeRoleAdminChanged represents a RoleAdminChanged event raised by the Edge contract.
type EdgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Edge *EdgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*EdgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Edge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &EdgeRoleAdminChangedIterator{contract: _Edge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Edge *EdgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *EdgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Edge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeRoleAdminChanged)
				if err := _Edge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Edge *EdgeFilterer) ParseRoleAdminChanged(log types.Log) (*EdgeRoleAdminChanged, error) {
	event := new(EdgeRoleAdminChanged)
	if err := _Edge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Edge contract.
type EdgeRoleGrantedIterator struct {
	Event *EdgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *EdgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeRoleGranted)
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
		it.Event = new(EdgeRoleGranted)
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
func (it *EdgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeRoleGranted represents a RoleGranted event raised by the Edge contract.
type EdgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Edge *EdgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EdgeRoleGrantedIterator, error) {

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

	logs, sub, err := _Edge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EdgeRoleGrantedIterator{contract: _Edge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Edge *EdgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *EdgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Edge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeRoleGranted)
				if err := _Edge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Edge *EdgeFilterer) ParseRoleGranted(log types.Log) (*EdgeRoleGranted, error) {
	event := new(EdgeRoleGranted)
	if err := _Edge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Edge contract.
type EdgeRoleRevokedIterator struct {
	Event *EdgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *EdgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeRoleRevoked)
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
		it.Event = new(EdgeRoleRevoked)
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
func (it *EdgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeRoleRevoked represents a RoleRevoked event raised by the Edge contract.
type EdgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Edge *EdgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*EdgeRoleRevokedIterator, error) {

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

	logs, sub, err := _Edge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &EdgeRoleRevokedIterator{contract: _Edge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Edge *EdgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *EdgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Edge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeRoleRevoked)
				if err := _Edge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Edge *EdgeFilterer) ParseRoleRevoked(log types.Log) (*EdgeRoleRevoked, error) {
	event := new(EdgeRoleRevoked)
	if err := _Edge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeSupportedIterator is returned from FilterSupported and is used to iterate over the raw logs and unpacked data for Supported events raised by the Edge contract.
type EdgeSupportedIterator struct {
	Event *EdgeSupported // Event containing the contract specifics and raw log

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
func (it *EdgeSupportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeSupported)
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
		it.Event = new(EdgeSupported)
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
func (it *EdgeSupportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeSupportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeSupported represents a Supported event raised by the Edge contract.
type EdgeSupported struct {
	Token   common.Address
	ChainId *big.Int
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSupported is a free log retrieval operation binding the contract event 0xd7cadba2609cba4364d862761a89c823e956d810d9947b1f158ad1aa9c2affbc.
//
// Solidity: event Supported(address token, uint256 chainId, bool status)
func (_Edge *EdgeFilterer) FilterSupported(opts *bind.FilterOpts) (*EdgeSupportedIterator, error) {

	logs, sub, err := _Edge.contract.FilterLogs(opts, "Supported")
	if err != nil {
		return nil, err
	}
	return &EdgeSupportedIterator{contract: _Edge.contract, event: "Supported", logs: logs, sub: sub}, nil
}

// WatchSupported is a free log subscription operation binding the contract event 0xd7cadba2609cba4364d862761a89c823e956d810d9947b1f158ad1aa9c2affbc.
//
// Solidity: event Supported(address token, uint256 chainId, bool status)
func (_Edge *EdgeFilterer) WatchSupported(opts *bind.WatchOpts, sink chan<- *EdgeSupported) (event.Subscription, error) {

	logs, sub, err := _Edge.contract.WatchLogs(opts, "Supported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeSupported)
				if err := _Edge.contract.UnpackLog(event, "Supported", log); err != nil {
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
func (_Edge *EdgeFilterer) ParseSupported(log types.Log) (*EdgeSupported, error) {
	event := new(EdgeSupported)
	if err := _Edge.contract.UnpackLog(event, "Supported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the Edge contract.
type EdgeThresholdChangedIterator struct {
	Event *EdgeThresholdChanged // Event containing the contract specifics and raw log

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
func (it *EdgeThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeThresholdChanged)
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
		it.Event = new(EdgeThresholdChanged)
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
func (it *EdgeThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeThresholdChanged represents a ThresholdChanged event raised by the Edge contract.
type EdgeThresholdChanged struct {
	Token        common.Address
	OldThreshold *big.Int
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_Edge *EdgeFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*EdgeThresholdChangedIterator, error) {

	logs, sub, err := _Edge.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &EdgeThresholdChangedIterator{contract: _Edge.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_Edge *EdgeFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *EdgeThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _Edge.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeThresholdChanged)
				if err := _Edge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_Edge *EdgeFilterer) ParseThresholdChanged(log types.Log) (*EdgeThresholdChanged, error) {
	event := new(EdgeThresholdChanged)
	if err := _Edge.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EdgeUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Edge contract.
type EdgeUpgradedIterator struct {
	Event *EdgeUpgraded // Event containing the contract specifics and raw log

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
func (it *EdgeUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EdgeUpgraded)
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
		it.Event = new(EdgeUpgraded)
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
func (it *EdgeUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EdgeUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EdgeUpgraded represents a Upgraded event raised by the Edge contract.
type EdgeUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Edge *EdgeFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*EdgeUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Edge.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &EdgeUpgradedIterator{contract: _Edge.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Edge *EdgeFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *EdgeUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Edge.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EdgeUpgraded)
				if err := _Edge.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Edge *EdgeFilterer) ParseUpgraded(log types.Log) (*EdgeUpgraded, error) {
	event := new(EdgeUpgraded)
	if err := _Edge.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
