// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// CrossLockABI is the input ABI used to generate the binding from.
const CrossLockABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Lock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_roleFlag\",\"type\":\"bytes32\"}],\"name\":\"addSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"token0Addrs\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"token1Addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chainIDs\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"_roleFlags\",\"type\":\"bytes32[]\"}],\"name\":\"addSupportTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposalOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isFinished\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isExist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"removeRoleFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"removeSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chainIDs\",\"type\":\"uint256[]\"}],\"name\":\"removeSupportTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roleFlag\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"supportToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"txUnlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CrossLock is an auto generated Go binding around an Ethereum contract.
type CrossLock struct {
	CrossLockCaller     // Read-only binding to the contract
	CrossLockTransactor // Write-only binding to the contract
	CrossLockFilterer   // Log filterer for contract events
}

// CrossLockCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrossLockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossLockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrossLockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossLockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrossLockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrossLockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrossLockSession struct {
	Contract     *CrossLock        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrossLockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrossLockCallerSession struct {
	Contract *CrossLockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// CrossLockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrossLockTransactorSession struct {
	Contract     *CrossLockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// CrossLockRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrossLockRaw struct {
	Contract *CrossLock // Generic contract binding to access the raw methods on
}

// CrossLockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrossLockCallerRaw struct {
	Contract *CrossLockCaller // Generic read-only contract binding to access the raw methods on
}

// CrossLockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrossLockTransactorRaw struct {
	Contract *CrossLockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrossLock creates a new instance of CrossLock, bound to a specific deployed contract.
func NewCrossLock(address common.Address, backend bind.ContractBackend) (*CrossLock, error) {
	contract, err := bindCrossLock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrossLock{CrossLockCaller: CrossLockCaller{contract: contract}, CrossLockTransactor: CrossLockTransactor{contract: contract}, CrossLockFilterer: CrossLockFilterer{contract: contract}}, nil
}

// NewCrossLockCaller creates a new read-only instance of CrossLock, bound to a specific deployed contract.
func NewCrossLockCaller(address common.Address, caller bind.ContractCaller) (*CrossLockCaller, error) {
	contract, err := bindCrossLock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrossLockCaller{contract: contract}, nil
}

// NewCrossLockTransactor creates a new write-only instance of CrossLock, bound to a specific deployed contract.
func NewCrossLockTransactor(address common.Address, transactor bind.ContractTransactor) (*CrossLockTransactor, error) {
	contract, err := bindCrossLock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrossLockTransactor{contract: contract}, nil
}

// NewCrossLockFilterer creates a new log filterer instance of CrossLock, bound to a specific deployed contract.
func NewCrossLockFilterer(address common.Address, filterer bind.ContractFilterer) (*CrossLockFilterer, error) {
	contract, err := bindCrossLock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrossLockFilterer{contract: contract}, nil
}

// bindCrossLock binds a generic wrapper to an already deployed contract.
func bindCrossLock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrossLockABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossLock *CrossLockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossLock.Contract.CrossLockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossLock *CrossLockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossLock.Contract.CrossLockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossLock *CrossLockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossLock.Contract.CrossLockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrossLock *CrossLockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrossLock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrossLock *CrossLockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrossLock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrossLock *CrossLockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrossLock.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossLock *CrossLockCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CrossLock.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossLock *CrossLockSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossLock.Contract.DEFAULTADMINROLE(&_CrossLock.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CrossLock *CrossLockCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CrossLock.Contract.DEFAULTADMINROLE(&_CrossLock.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossLock *CrossLockCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CrossLock.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossLock *CrossLockSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossLock.Contract.GetRoleAdmin(&_CrossLock.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CrossLock *CrossLockCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CrossLock.Contract.GetRoleAdmin(&_CrossLock.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CrossLock *CrossLockCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrossLock.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CrossLock *CrossLockSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _CrossLock.Contract.GetRoleMember(&_CrossLock.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CrossLock *CrossLockCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _CrossLock.Contract.GetRoleMember(&_CrossLock.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CrossLock *CrossLockCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CrossLock.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CrossLock *CrossLockSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _CrossLock.Contract.GetRoleMemberCount(&_CrossLock.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CrossLock *CrossLockCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _CrossLock.Contract.GetRoleMemberCount(&_CrossLock.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossLock *CrossLockCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CrossLock.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossLock *CrossLockSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossLock.Contract.HasRole(&_CrossLock.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CrossLock *CrossLockCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CrossLock.Contract.HasRole(&_CrossLock.CallOpts, role, account)
}

// ProposalOf is a free data retrieval call binding the contract method 0xf48c0abf.
//
// Solidity: function proposalOf(bytes32 ) view returns(address tokenTo, address from, address to, uint256 amount, uint256 count, string txid, bool isFinished, bool isExist)
func (_CrossLock *CrossLockCaller) ProposalOf(opts *bind.CallOpts, arg0 [32]byte) (struct {
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
	err := _CrossLock.contract.Call(opts, &out, "proposalOf", arg0)

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
func (_CrossLock *CrossLockSession) ProposalOf(arg0 [32]byte) (struct {
	TokenTo    common.Address
	From       common.Address
	To         common.Address
	Amount     *big.Int
	Count      *big.Int
	Txid       string
	IsFinished bool
	IsExist    bool
}, error) {
	return _CrossLock.Contract.ProposalOf(&_CrossLock.CallOpts, arg0)
}

// ProposalOf is a free data retrieval call binding the contract method 0xf48c0abf.
//
// Solidity: function proposalOf(bytes32 ) view returns(address tokenTo, address from, address to, uint256 amount, uint256 count, string txid, bool isFinished, bool isExist)
func (_CrossLock *CrossLockCallerSession) ProposalOf(arg0 [32]byte) (struct {
	TokenTo    common.Address
	From       common.Address
	To         common.Address
	Amount     *big.Int
	Count      *big.Int
	Txid       string
	IsFinished bool
	IsExist    bool
}, error) {
	return _CrossLock.Contract.ProposalOf(&_CrossLock.CallOpts, arg0)
}

// RoleFlag is a free data retrieval call binding the contract method 0xe848ec59.
//
// Solidity: function roleFlag(address , uint256 ) view returns(bytes32)
func (_CrossLock *CrossLockCaller) RoleFlag(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _CrossLock.contract.Call(opts, &out, "roleFlag", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RoleFlag is a free data retrieval call binding the contract method 0xe848ec59.
//
// Solidity: function roleFlag(address , uint256 ) view returns(bytes32)
func (_CrossLock *CrossLockSession) RoleFlag(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _CrossLock.Contract.RoleFlag(&_CrossLock.CallOpts, arg0, arg1)
}

// RoleFlag is a free data retrieval call binding the contract method 0xe848ec59.
//
// Solidity: function roleFlag(address , uint256 ) view returns(bytes32)
func (_CrossLock *CrossLockCallerSession) RoleFlag(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _CrossLock.Contract.RoleFlag(&_CrossLock.CallOpts, arg0, arg1)
}

// SupportToken is a free data retrieval call binding the contract method 0xd692f4c5.
//
// Solidity: function supportToken(address , uint256 ) view returns(address)
func (_CrossLock *CrossLockCaller) SupportToken(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CrossLock.contract.Call(opts, &out, "supportToken", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportToken is a free data retrieval call binding the contract method 0xd692f4c5.
//
// Solidity: function supportToken(address , uint256 ) view returns(address)
func (_CrossLock *CrossLockSession) SupportToken(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _CrossLock.Contract.SupportToken(&_CrossLock.CallOpts, arg0, arg1)
}

// SupportToken is a free data retrieval call binding the contract method 0xd692f4c5.
//
// Solidity: function supportToken(address , uint256 ) view returns(address)
func (_CrossLock *CrossLockCallerSession) SupportToken(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _CrossLock.Contract.SupportToken(&_CrossLock.CallOpts, arg0, arg1)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_CrossLock *CrossLockCaller) Threshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrossLock.contract.Call(opts, &out, "threshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_CrossLock *CrossLockSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _CrossLock.Contract.Threshold(&_CrossLock.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_CrossLock *CrossLockCallerSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _CrossLock.Contract.Threshold(&_CrossLock.CallOpts, arg0)
}

// TxUnlocked is a free data retrieval call binding the contract method 0x967145af.
//
// Solidity: function txUnlocked(string ) view returns(bool)
func (_CrossLock *CrossLockCaller) TxUnlocked(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _CrossLock.contract.Call(opts, &out, "txUnlocked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TxUnlocked is a free data retrieval call binding the contract method 0x967145af.
//
// Solidity: function txUnlocked(string ) view returns(bool)
func (_CrossLock *CrossLockSession) TxUnlocked(arg0 string) (bool, error) {
	return _CrossLock.Contract.TxUnlocked(&_CrossLock.CallOpts, arg0)
}

// TxUnlocked is a free data retrieval call binding the contract method 0x967145af.
//
// Solidity: function txUnlocked(string ) view returns(bool)
func (_CrossLock *CrossLockCallerSession) TxUnlocked(arg0 string) (bool, error) {
	return _CrossLock.Contract.TxUnlocked(&_CrossLock.CallOpts, arg0)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x9fddf460.
//
// Solidity: function addSupportToken(address token0, address token1, uint256 chainID, bytes32 _roleFlag) returns()
func (_CrossLock *CrossLockTransactor) AddSupportToken(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, chainID *big.Int, _roleFlag [32]byte) (*types.Transaction, error) {
	return _CrossLock.contract.Transact(opts, "addSupportToken", token0, token1, chainID, _roleFlag)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x9fddf460.
//
// Solidity: function addSupportToken(address token0, address token1, uint256 chainID, bytes32 _roleFlag) returns()
func (_CrossLock *CrossLockSession) AddSupportToken(token0 common.Address, token1 common.Address, chainID *big.Int, _roleFlag [32]byte) (*types.Transaction, error) {
	return _CrossLock.Contract.AddSupportToken(&_CrossLock.TransactOpts, token0, token1, chainID, _roleFlag)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x9fddf460.
//
// Solidity: function addSupportToken(address token0, address token1, uint256 chainID, bytes32 _roleFlag) returns()
func (_CrossLock *CrossLockTransactorSession) AddSupportToken(token0 common.Address, token1 common.Address, chainID *big.Int, _roleFlag [32]byte) (*types.Transaction, error) {
	return _CrossLock.Contract.AddSupportToken(&_CrossLock.TransactOpts, token0, token1, chainID, _roleFlag)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0x268f3698.
//
// Solidity: function addSupportTokens(address[] token0Addrs, address[] token1Addrs, uint256[] chainIDs, bytes32[] _roleFlags) returns()
func (_CrossLock *CrossLockTransactor) AddSupportTokens(opts *bind.TransactOpts, token0Addrs []common.Address, token1Addrs []common.Address, chainIDs []*big.Int, _roleFlags [][32]byte) (*types.Transaction, error) {
	return _CrossLock.contract.Transact(opts, "addSupportTokens", token0Addrs, token1Addrs, chainIDs, _roleFlags)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0x268f3698.
//
// Solidity: function addSupportTokens(address[] token0Addrs, address[] token1Addrs, uint256[] chainIDs, bytes32[] _roleFlags) returns()
func (_CrossLock *CrossLockSession) AddSupportTokens(token0Addrs []common.Address, token1Addrs []common.Address, chainIDs []*big.Int, _roleFlags [][32]byte) (*types.Transaction, error) {
	return _CrossLock.Contract.AddSupportTokens(&_CrossLock.TransactOpts, token0Addrs, token1Addrs, chainIDs, _roleFlags)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0x268f3698.
//
// Solidity: function addSupportTokens(address[] token0Addrs, address[] token1Addrs, uint256[] chainIDs, bytes32[] _roleFlags) returns()
func (_CrossLock *CrossLockTransactorSession) AddSupportTokens(token0Addrs []common.Address, token1Addrs []common.Address, chainIDs []*big.Int, _roleFlags [][32]byte) (*types.Transaction, error) {
	return _CrossLock.Contract.AddSupportTokens(&_CrossLock.TransactOpts, token0Addrs, token1Addrs, chainIDs, _roleFlags)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossLock *CrossLockTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossLock.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossLock *CrossLockSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossLock.Contract.GrantRole(&_CrossLock.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CrossLock *CrossLockTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossLock.Contract.GrantRole(&_CrossLock.TransactOpts, role, account)
}

// Lock is a paid mutator transaction binding the contract method 0xeaac2925.
//
// Solidity: function lock(address token0, uint256 chainID, address to, uint256 amount) returns()
func (_CrossLock *CrossLockTransactor) Lock(opts *bind.TransactOpts, token0 common.Address, chainID *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossLock.contract.Transact(opts, "lock", token0, chainID, to, amount)
}

// Lock is a paid mutator transaction binding the contract method 0xeaac2925.
//
// Solidity: function lock(address token0, uint256 chainID, address to, uint256 amount) returns()
func (_CrossLock *CrossLockSession) Lock(token0 common.Address, chainID *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossLock.Contract.Lock(&_CrossLock.TransactOpts, token0, chainID, to, amount)
}

// Lock is a paid mutator transaction binding the contract method 0xeaac2925.
//
// Solidity: function lock(address token0, uint256 chainID, address to, uint256 amount) returns()
func (_CrossLock *CrossLockTransactorSession) Lock(token0 common.Address, chainID *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrossLock.Contract.Lock(&_CrossLock.TransactOpts, token0, chainID, to, amount)
}

// RemoveRoleFlag is a paid mutator transaction binding the contract method 0xf38d9932.
//
// Solidity: function removeRoleFlag(address token0, uint256 chainID) returns()
func (_CrossLock *CrossLockTransactor) RemoveRoleFlag(opts *bind.TransactOpts, token0 common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _CrossLock.contract.Transact(opts, "removeRoleFlag", token0, chainID)
}

// RemoveRoleFlag is a paid mutator transaction binding the contract method 0xf38d9932.
//
// Solidity: function removeRoleFlag(address token0, uint256 chainID) returns()
func (_CrossLock *CrossLockSession) RemoveRoleFlag(token0 common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _CrossLock.Contract.RemoveRoleFlag(&_CrossLock.TransactOpts, token0, chainID)
}

// RemoveRoleFlag is a paid mutator transaction binding the contract method 0xf38d9932.
//
// Solidity: function removeRoleFlag(address token0, uint256 chainID) returns()
func (_CrossLock *CrossLockTransactorSession) RemoveRoleFlag(token0 common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _CrossLock.Contract.RemoveRoleFlag(&_CrossLock.TransactOpts, token0, chainID)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xf8f8d5c0.
//
// Solidity: function removeSupportToken(address token0, uint256 chainID) returns()
func (_CrossLock *CrossLockTransactor) RemoveSupportToken(opts *bind.TransactOpts, token0 common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _CrossLock.contract.Transact(opts, "removeSupportToken", token0, chainID)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xf8f8d5c0.
//
// Solidity: function removeSupportToken(address token0, uint256 chainID) returns()
func (_CrossLock *CrossLockSession) RemoveSupportToken(token0 common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _CrossLock.Contract.RemoveSupportToken(&_CrossLock.TransactOpts, token0, chainID)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xf8f8d5c0.
//
// Solidity: function removeSupportToken(address token0, uint256 chainID) returns()
func (_CrossLock *CrossLockTransactorSession) RemoveSupportToken(token0 common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _CrossLock.Contract.RemoveSupportToken(&_CrossLock.TransactOpts, token0, chainID)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x03507ba5.
//
// Solidity: function removeSupportTokens(address[] addrs, uint256[] chainIDs) returns()
func (_CrossLock *CrossLockTransactor) RemoveSupportTokens(opts *bind.TransactOpts, addrs []common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossLock.contract.Transact(opts, "removeSupportTokens", addrs, chainIDs)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x03507ba5.
//
// Solidity: function removeSupportTokens(address[] addrs, uint256[] chainIDs) returns()
func (_CrossLock *CrossLockSession) RemoveSupportTokens(addrs []common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossLock.Contract.RemoveSupportTokens(&_CrossLock.TransactOpts, addrs, chainIDs)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x03507ba5.
//
// Solidity: function removeSupportTokens(address[] addrs, uint256[] chainIDs) returns()
func (_CrossLock *CrossLockTransactorSession) RemoveSupportTokens(addrs []common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _CrossLock.Contract.RemoveSupportTokens(&_CrossLock.TransactOpts, addrs, chainIDs)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossLock *CrossLockTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossLock.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossLock *CrossLockSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossLock.Contract.RenounceRole(&_CrossLock.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_CrossLock *CrossLockTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossLock.Contract.RenounceRole(&_CrossLock.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossLock *CrossLockTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossLock.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossLock *CrossLockSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossLock.Contract.RevokeRole(&_CrossLock.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CrossLock *CrossLockTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CrossLock.Contract.RevokeRole(&_CrossLock.TransactOpts, role, account)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_CrossLock *CrossLockTransactor) SetThreshold(opts *bind.TransactOpts, token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _CrossLock.contract.Transact(opts, "setThreshold", token0, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_CrossLock *CrossLockSession) SetThreshold(token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _CrossLock.Contract.SetThreshold(&_CrossLock.TransactOpts, token0, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_CrossLock *CrossLockTransactorSession) SetThreshold(token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _CrossLock.Contract.SetThreshold(&_CrossLock.TransactOpts, token0, _threshold)
}

// Unlock is a paid mutator transaction binding the contract method 0x715ec45c.
//
// Solidity: function unlock(address token0, uint256 chainID, address from, address to, uint256 amount, string txid) returns()
func (_CrossLock *CrossLockTransactor) Unlock(opts *bind.TransactOpts, token0 common.Address, chainID *big.Int, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _CrossLock.contract.Transact(opts, "unlock", token0, chainID, from, to, amount, txid)
}

// Unlock is a paid mutator transaction binding the contract method 0x715ec45c.
//
// Solidity: function unlock(address token0, uint256 chainID, address from, address to, uint256 amount, string txid) returns()
func (_CrossLock *CrossLockSession) Unlock(token0 common.Address, chainID *big.Int, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _CrossLock.Contract.Unlock(&_CrossLock.TransactOpts, token0, chainID, from, to, amount, txid)
}

// Unlock is a paid mutator transaction binding the contract method 0x715ec45c.
//
// Solidity: function unlock(address token0, uint256 chainID, address from, address to, uint256 amount, string txid) returns()
func (_CrossLock *CrossLockTransactorSession) Unlock(token0 common.Address, chainID *big.Int, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _CrossLock.Contract.Unlock(&_CrossLock.TransactOpts, token0, chainID, from, to, amount, txid)
}

// CrossLockLockIterator is returned from FilterLock and is used to iterate over the raw logs and unpacked data for Lock events raised by the CrossLock contract.
type CrossLockLockIterator struct {
	Event *CrossLockLock // Event containing the contract specifics and raw log

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
func (it *CrossLockLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossLockLock)
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
		it.Event = new(CrossLockLock)
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
func (it *CrossLockLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossLockLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossLockLock represents a Lock event raised by the CrossLock contract.
type CrossLockLock struct {
	Token0  common.Address
	Token1  common.Address
	ChainID *big.Int
	Locker  common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLock is a free log retrieval operation binding the contract event 0x4620c97568351e0eab160aa39863643b388e20b445120a289cb346344a237122.
//
// Solidity: event Lock(address token0, address token1, uint256 chainID, address locker, address to, uint256 amount)
func (_CrossLock *CrossLockFilterer) FilterLock(opts *bind.FilterOpts) (*CrossLockLockIterator, error) {

	logs, sub, err := _CrossLock.contract.FilterLogs(opts, "Lock")
	if err != nil {
		return nil, err
	}
	return &CrossLockLockIterator{contract: _CrossLock.contract, event: "Lock", logs: logs, sub: sub}, nil
}

// WatchLock is a free log subscription operation binding the contract event 0x4620c97568351e0eab160aa39863643b388e20b445120a289cb346344a237122.
//
// Solidity: event Lock(address token0, address token1, uint256 chainID, address locker, address to, uint256 amount)
func (_CrossLock *CrossLockFilterer) WatchLock(opts *bind.WatchOpts, sink chan<- *CrossLockLock) (event.Subscription, error) {

	logs, sub, err := _CrossLock.contract.WatchLogs(opts, "Lock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossLockLock)
				if err := _CrossLock.contract.UnpackLog(event, "Lock", log); err != nil {
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

// ParseLock is a log parse operation binding the contract event 0x4620c97568351e0eab160aa39863643b388e20b445120a289cb346344a237122.
//
// Solidity: event Lock(address token0, address token1, uint256 chainID, address locker, address to, uint256 amount)
func (_CrossLock *CrossLockFilterer) ParseLock(log types.Log) (*CrossLockLock, error) {
	event := new(CrossLockLock)
	if err := _CrossLock.contract.UnpackLog(event, "Lock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossLockProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the CrossLock contract.
type CrossLockProposalVotedIterator struct {
	Event *CrossLockProposalVoted // Event containing the contract specifics and raw log

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
func (it *CrossLockProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossLockProposalVoted)
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
		it.Event = new(CrossLockProposalVoted)
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
func (it *CrossLockProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossLockProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossLockProposalVoted represents a ProposalVoted event raised by the CrossLock contract.
type CrossLockProposalVoted struct {
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
func (_CrossLock *CrossLockFilterer) FilterProposalVoted(opts *bind.FilterOpts) (*CrossLockProposalVotedIterator, error) {

	logs, sub, err := _CrossLock.contract.FilterLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return &CrossLockProposalVotedIterator{contract: _CrossLock.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0xe458477b94285c9f254a407e2614c39f55230b527ff7f51d76fdf33e251ae936.
//
// Solidity: event ProposalVoted(address token, address from, address to, uint256 amount, address proposer, uint256 count, uint256 threshold)
func (_CrossLock *CrossLockFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *CrossLockProposalVoted) (event.Subscription, error) {

	logs, sub, err := _CrossLock.contract.WatchLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossLockProposalVoted)
				if err := _CrossLock.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
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
func (_CrossLock *CrossLockFilterer) ParseProposalVoted(log types.Log) (*CrossLockProposalVoted, error) {
	event := new(CrossLockProposalVoted)
	if err := _CrossLock.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossLockRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CrossLock contract.
type CrossLockRoleAdminChangedIterator struct {
	Event *CrossLockRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CrossLockRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossLockRoleAdminChanged)
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
		it.Event = new(CrossLockRoleAdminChanged)
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
func (it *CrossLockRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossLockRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossLockRoleAdminChanged represents a RoleAdminChanged event raised by the CrossLock contract.
type CrossLockRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossLock *CrossLockFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CrossLockRoleAdminChangedIterator, error) {

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

	logs, sub, err := _CrossLock.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CrossLockRoleAdminChangedIterator{contract: _CrossLock.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CrossLock *CrossLockFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CrossLockRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CrossLock.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossLockRoleAdminChanged)
				if err := _CrossLock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CrossLock *CrossLockFilterer) ParseRoleAdminChanged(log types.Log) (*CrossLockRoleAdminChanged, error) {
	event := new(CrossLockRoleAdminChanged)
	if err := _CrossLock.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossLockRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CrossLock contract.
type CrossLockRoleGrantedIterator struct {
	Event *CrossLockRoleGranted // Event containing the contract specifics and raw log

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
func (it *CrossLockRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossLockRoleGranted)
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
		it.Event = new(CrossLockRoleGranted)
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
func (it *CrossLockRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossLockRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossLockRoleGranted represents a RoleGranted event raised by the CrossLock contract.
type CrossLockRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossLock *CrossLockFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossLockRoleGrantedIterator, error) {

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

	logs, sub, err := _CrossLock.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossLockRoleGrantedIterator{contract: _CrossLock.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossLock *CrossLockFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CrossLockRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossLock.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossLockRoleGranted)
				if err := _CrossLock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CrossLock *CrossLockFilterer) ParseRoleGranted(log types.Log) (*CrossLockRoleGranted, error) {
	event := new(CrossLockRoleGranted)
	if err := _CrossLock.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossLockRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CrossLock contract.
type CrossLockRoleRevokedIterator struct {
	Event *CrossLockRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CrossLockRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossLockRoleRevoked)
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
		it.Event = new(CrossLockRoleRevoked)
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
func (it *CrossLockRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossLockRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossLockRoleRevoked represents a RoleRevoked event raised by the CrossLock contract.
type CrossLockRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossLock *CrossLockFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CrossLockRoleRevokedIterator, error) {

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

	logs, sub, err := _CrossLock.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CrossLockRoleRevokedIterator{contract: _CrossLock.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CrossLock *CrossLockFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CrossLockRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CrossLock.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossLockRoleRevoked)
				if err := _CrossLock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CrossLock *CrossLockFilterer) ParseRoleRevoked(log types.Log) (*CrossLockRoleRevoked, error) {
	event := new(CrossLockRoleRevoked)
	if err := _CrossLock.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossLockThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the CrossLock contract.
type CrossLockThresholdChangedIterator struct {
	Event *CrossLockThresholdChanged // Event containing the contract specifics and raw log

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
func (it *CrossLockThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossLockThresholdChanged)
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
		it.Event = new(CrossLockThresholdChanged)
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
func (it *CrossLockThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossLockThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossLockThresholdChanged represents a ThresholdChanged event raised by the CrossLock contract.
type CrossLockThresholdChanged struct {
	Token        common.Address
	OldThreshold *big.Int
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_CrossLock *CrossLockFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*CrossLockThresholdChangedIterator, error) {

	logs, sub, err := _CrossLock.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &CrossLockThresholdChangedIterator{contract: _CrossLock.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_CrossLock *CrossLockFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *CrossLockThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _CrossLock.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossLockThresholdChanged)
				if err := _CrossLock.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_CrossLock *CrossLockFilterer) ParseThresholdChanged(log types.Log) (*CrossLockThresholdChanged, error) {
	event := new(CrossLockThresholdChanged)
	if err := _CrossLock.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrossLockUnlockIterator is returned from FilterUnlock and is used to iterate over the raw logs and unpacked data for Unlock events raised by the CrossLock contract.
type CrossLockUnlockIterator struct {
	Event *CrossLockUnlock // Event containing the contract specifics and raw log

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
func (it *CrossLockUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrossLockUnlock)
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
		it.Event = new(CrossLockUnlock)
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
func (it *CrossLockUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrossLockUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrossLockUnlock represents a Unlock event raised by the CrossLock contract.
type CrossLockUnlock struct {
	Token0  common.Address
	Token1  common.Address
	ChainID *big.Int
	From    common.Address
	To      common.Address
	Amount  *big.Int
	Txid    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnlock is a free log retrieval operation binding the contract event 0x0c6a7c66882059d74054ca1f9deb21718c82c28d890c3eeeb7c35a50839abf00.
//
// Solidity: event Unlock(address token0, address token1, uint256 chainID, address from, address to, uint256 amount, string txid)
func (_CrossLock *CrossLockFilterer) FilterUnlock(opts *bind.FilterOpts) (*CrossLockUnlockIterator, error) {

	logs, sub, err := _CrossLock.contract.FilterLogs(opts, "Unlock")
	if err != nil {
		return nil, err
	}
	return &CrossLockUnlockIterator{contract: _CrossLock.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0x0c6a7c66882059d74054ca1f9deb21718c82c28d890c3eeeb7c35a50839abf00.
//
// Solidity: event Unlock(address token0, address token1, uint256 chainID, address from, address to, uint256 amount, string txid)
func (_CrossLock *CrossLockFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *CrossLockUnlock) (event.Subscription, error) {

	logs, sub, err := _CrossLock.contract.WatchLogs(opts, "Unlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrossLockUnlock)
				if err := _CrossLock.contract.UnpackLog(event, "Unlock", log); err != nil {
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

// ParseUnlock is a log parse operation binding the contract event 0x0c6a7c66882059d74054ca1f9deb21718c82c28d890c3eeeb7c35a50839abf00.
//
// Solidity: event Unlock(address token0, address token1, uint256 chainID, address from, address to, uint256 amount, string txid)
func (_CrossLock *CrossLockFilterer) ParseUnlock(log types.Log) (*CrossLockUnlock, error) {
	event := new(CrossLockUnlock)
	if err := _CrossLock.contract.UnpackLog(event, "Unlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
