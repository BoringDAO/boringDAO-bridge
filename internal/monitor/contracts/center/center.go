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
)

// InParam is an auto generated low-level Go binding around an user-defined struct.
type InParam struct {
	FromChainId *big.Int
	FromToken   common.Address
	From        common.Address
	ToChainId   *big.Int
	ToToken     common.Address
	To          common.Address
	Amount      *big.Int
}

// OutParam is an auto generated low-level Go binding around an user-defined struct.
type OutParam struct {
	FromChainId *big.Int
	FromToken   common.Address
	From        common.Address
	ToChainId   *big.Int
	To          common.Address
	Amount      *big.Int
}

// AccessControlUpgradeableMetaData contains all meta data concerning the AccessControlUpgradeable contract.
var AccessControlUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// AccessControlUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use AccessControlUpgradeableMetaData.ABI instead.
var AccessControlUpgradeableABI = AccessControlUpgradeableMetaData.ABI

// Deprecated: Use AccessControlUpgradeableMetaData.Sigs instead.
// AccessControlUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlUpgradeableFuncSigs = AccessControlUpgradeableMetaData.Sigs

// AccessControlUpgradeable is an auto generated Go binding around an Ethereum contract.
type AccessControlUpgradeable struct {
	AccessControlUpgradeableCaller     // Read-only binding to the contract
	AccessControlUpgradeableTransactor // Write-only binding to the contract
	AccessControlUpgradeableFilterer   // Log filterer for contract events
}

// AccessControlUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlUpgradeableSession struct {
	Contract     *AccessControlUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AccessControlUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlUpgradeableCallerSession struct {
	Contract *AccessControlUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AccessControlUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlUpgradeableTransactorSession struct {
	Contract     *AccessControlUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AccessControlUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlUpgradeableRaw struct {
	Contract *AccessControlUpgradeable // Generic contract binding to access the raw methods on
}

// AccessControlUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlUpgradeableCallerRaw struct {
	Contract *AccessControlUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlUpgradeableTransactorRaw struct {
	Contract *AccessControlUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControlUpgradeable creates a new instance of AccessControlUpgradeable, bound to a specific deployed contract.
func NewAccessControlUpgradeable(address common.Address, backend bind.ContractBackend) (*AccessControlUpgradeable, error) {
	contract, err := bindAccessControlUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeable{AccessControlUpgradeableCaller: AccessControlUpgradeableCaller{contract: contract}, AccessControlUpgradeableTransactor: AccessControlUpgradeableTransactor{contract: contract}, AccessControlUpgradeableFilterer: AccessControlUpgradeableFilterer{contract: contract}}, nil
}

// NewAccessControlUpgradeableCaller creates a new read-only instance of AccessControlUpgradeable, bound to a specific deployed contract.
func NewAccessControlUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*AccessControlUpgradeableCaller, error) {
	contract, err := bindAccessControlUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeableCaller{contract: contract}, nil
}

// NewAccessControlUpgradeableTransactor creates a new write-only instance of AccessControlUpgradeable, bound to a specific deployed contract.
func NewAccessControlUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlUpgradeableTransactor, error) {
	contract, err := bindAccessControlUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeableTransactor{contract: contract}, nil
}

// NewAccessControlUpgradeableFilterer creates a new log filterer instance of AccessControlUpgradeable, bound to a specific deployed contract.
func NewAccessControlUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlUpgradeableFilterer, error) {
	contract, err := bindAccessControlUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeableFilterer{contract: contract}, nil
}

// bindAccessControlUpgradeable binds a generic wrapper to an already deployed contract.
func bindAccessControlUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlUpgradeable *AccessControlUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlUpgradeable.Contract.AccessControlUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlUpgradeable *AccessControlUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.AccessControlUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlUpgradeable *AccessControlUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.AccessControlUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControlUpgradeable *AccessControlUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControlUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlUpgradeable *AccessControlUpgradeableCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlUpgradeable.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControlUpgradeable.Contract.DEFAULTADMINROLE(&_AccessControlUpgradeable.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControlUpgradeable *AccessControlUpgradeableCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControlUpgradeable.Contract.DEFAULTADMINROLE(&_AccessControlUpgradeable.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlUpgradeable *AccessControlUpgradeableCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessControlUpgradeable.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControlUpgradeable.Contract.GetRoleAdmin(&_AccessControlUpgradeable.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControlUpgradeable *AccessControlUpgradeableCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControlUpgradeable.Contract.GetRoleAdmin(&_AccessControlUpgradeable.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlUpgradeable *AccessControlUpgradeableCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControlUpgradeable.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControlUpgradeable.Contract.HasRole(&_AccessControlUpgradeable.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControlUpgradeable *AccessControlUpgradeableCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControlUpgradeable.Contract.HasRole(&_AccessControlUpgradeable.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControlUpgradeable *AccessControlUpgradeableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AccessControlUpgradeable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControlUpgradeable.Contract.SupportsInterface(&_AccessControlUpgradeable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControlUpgradeable *AccessControlUpgradeableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControlUpgradeable.Contract.SupportsInterface(&_AccessControlUpgradeable.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.GrantRole(&_AccessControlUpgradeable.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.GrantRole(&_AccessControlUpgradeable.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.RenounceRole(&_AccessControlUpgradeable.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.RenounceRole(&_AccessControlUpgradeable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.RevokeRole(&_AccessControlUpgradeable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControlUpgradeable *AccessControlUpgradeableTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControlUpgradeable.Contract.RevokeRole(&_AccessControlUpgradeable.TransactOpts, role, account)
}

// AccessControlUpgradeableRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccessControlUpgradeable contract.
type AccessControlUpgradeableRoleAdminChangedIterator struct {
	Event *AccessControlUpgradeableRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccessControlUpgradeableRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlUpgradeableRoleAdminChanged)
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
		it.Event = new(AccessControlUpgradeableRoleAdminChanged)
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
func (it *AccessControlUpgradeableRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlUpgradeableRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlUpgradeableRoleAdminChanged represents a RoleAdminChanged event raised by the AccessControlUpgradeable contract.
type AccessControlUpgradeableRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AccessControlUpgradeableRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AccessControlUpgradeable.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeableRoleAdminChangedIterator{contract: _AccessControlUpgradeable.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessControlUpgradeableRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AccessControlUpgradeable.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlUpgradeableRoleAdminChanged)
				if err := _AccessControlUpgradeable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) ParseRoleAdminChanged(log types.Log) (*AccessControlUpgradeableRoleAdminChanged, error) {
	event := new(AccessControlUpgradeableRoleAdminChanged)
	if err := _AccessControlUpgradeable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlUpgradeableRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControlUpgradeable contract.
type AccessControlUpgradeableRoleGrantedIterator struct {
	Event *AccessControlUpgradeableRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessControlUpgradeableRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlUpgradeableRoleGranted)
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
		it.Event = new(AccessControlUpgradeableRoleGranted)
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
func (it *AccessControlUpgradeableRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlUpgradeableRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlUpgradeableRoleGranted represents a RoleGranted event raised by the AccessControlUpgradeable contract.
type AccessControlUpgradeableRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlUpgradeableRoleGrantedIterator, error) {

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

	logs, sub, err := _AccessControlUpgradeable.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeableRoleGrantedIterator{contract: _AccessControlUpgradeable.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlUpgradeableRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessControlUpgradeable.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlUpgradeableRoleGranted)
				if err := _AccessControlUpgradeable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) ParseRoleGranted(log types.Log) (*AccessControlUpgradeableRoleGranted, error) {
	event := new(AccessControlUpgradeableRoleGranted)
	if err := _AccessControlUpgradeable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlUpgradeableRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControlUpgradeable contract.
type AccessControlUpgradeableRoleRevokedIterator struct {
	Event *AccessControlUpgradeableRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessControlUpgradeableRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlUpgradeableRoleRevoked)
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
		it.Event = new(AccessControlUpgradeableRoleRevoked)
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
func (it *AccessControlUpgradeableRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlUpgradeableRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlUpgradeableRoleRevoked represents a RoleRevoked event raised by the AccessControlUpgradeable contract.
type AccessControlUpgradeableRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlUpgradeableRoleRevokedIterator, error) {

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

	logs, sub, err := _AccessControlUpgradeable.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlUpgradeableRoleRevokedIterator{contract: _AccessControlUpgradeable.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlUpgradeableRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessControlUpgradeable.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlUpgradeableRoleRevoked)
				if err := _AccessControlUpgradeable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AccessControlUpgradeable *AccessControlUpgradeableFilterer) ParseRoleRevoked(log types.Log) (*AccessControlUpgradeableRoleRevoked, error) {
	event := new(AccessControlUpgradeableRoleRevoked)
	if err := _AccessControlUpgradeable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressUpgradeableMetaData contains all meta data concerning the AddressUpgradeable contract.
var AddressUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220e5fdfd0e1537224197a8c8cd2642fc26c2fbd0477640b29bd628946a8decd02c64736f6c63430008020033",
}

// AddressUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressUpgradeableMetaData.ABI instead.
var AddressUpgradeableABI = AddressUpgradeableMetaData.ABI

// AddressUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressUpgradeableMetaData.Bin instead.
var AddressUpgradeableBin = AddressUpgradeableMetaData.Bin

// DeployAddressUpgradeable deploys a new Ethereum contract, binding an instance of AddressUpgradeable to it.
func DeployAddressUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressUpgradeable, error) {
	parsed, err := AddressUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// AddressUpgradeable is an auto generated Go binding around an Ethereum contract.
type AddressUpgradeable struct {
	AddressUpgradeableCaller     // Read-only binding to the contract
	AddressUpgradeableTransactor // Write-only binding to the contract
	AddressUpgradeableFilterer   // Log filterer for contract events
}

// AddressUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressUpgradeableSession struct {
	Contract     *AddressUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AddressUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressUpgradeableCallerSession struct {
	Contract *AddressUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AddressUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressUpgradeableTransactorSession struct {
	Contract     *AddressUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AddressUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressUpgradeableRaw struct {
	Contract *AddressUpgradeable // Generic contract binding to access the raw methods on
}

// AddressUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressUpgradeableCallerRaw struct {
	Contract *AddressUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// AddressUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressUpgradeableTransactorRaw struct {
	Contract *AddressUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressUpgradeable creates a new instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeable(address common.Address, backend bind.ContractBackend) (*AddressUpgradeable, error) {
	contract, err := bindAddressUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeable{AddressUpgradeableCaller: AddressUpgradeableCaller{contract: contract}, AddressUpgradeableTransactor: AddressUpgradeableTransactor{contract: contract}, AddressUpgradeableFilterer: AddressUpgradeableFilterer{contract: contract}}, nil
}

// NewAddressUpgradeableCaller creates a new read-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*AddressUpgradeableCaller, error) {
	contract, err := bindAddressUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableCaller{contract: contract}, nil
}

// NewAddressUpgradeableTransactor creates a new write-only instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressUpgradeableTransactor, error) {
	contract, err := bindAddressUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableTransactor{contract: contract}, nil
}

// NewAddressUpgradeableFilterer creates a new log filterer instance of AddressUpgradeable, bound to a specific deployed contract.
func NewAddressUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressUpgradeableFilterer, error) {
	contract, err := bindAddressUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressUpgradeableFilterer{contract: contract}, nil
}

// bindAddressUpgradeable binds a generic wrapper to an already deployed contract.
func bindAddressUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.AddressUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.AddressUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressUpgradeable *AddressUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressUpgradeable *AddressUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ContextUpgradeableMetaData contains all meta data concerning the ContextUpgradeable contract.
var ContextUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
}

// ContextUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ContextUpgradeableMetaData.ABI instead.
var ContextUpgradeableABI = ContextUpgradeableMetaData.ABI

// ContextUpgradeable is an auto generated Go binding around an Ethereum contract.
type ContextUpgradeable struct {
	ContextUpgradeableCaller     // Read-only binding to the contract
	ContextUpgradeableTransactor // Write-only binding to the contract
	ContextUpgradeableFilterer   // Log filterer for contract events
}

// ContextUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextUpgradeableSession struct {
	Contract     *ContextUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContextUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextUpgradeableCallerSession struct {
	Contract *ContextUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// ContextUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextUpgradeableTransactorSession struct {
	Contract     *ContextUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// ContextUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextUpgradeableRaw struct {
	Contract *ContextUpgradeable // Generic contract binding to access the raw methods on
}

// ContextUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextUpgradeableCallerRaw struct {
	Contract *ContextUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ContextUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextUpgradeableTransactorRaw struct {
	Contract *ContextUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContextUpgradeable creates a new instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeable(address common.Address, backend bind.ContractBackend) (*ContextUpgradeable, error) {
	contract, err := bindContextUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeable{ContextUpgradeableCaller: ContextUpgradeableCaller{contract: contract}, ContextUpgradeableTransactor: ContextUpgradeableTransactor{contract: contract}, ContextUpgradeableFilterer: ContextUpgradeableFilterer{contract: contract}}, nil
}

// NewContextUpgradeableCaller creates a new read-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ContextUpgradeableCaller, error) {
	contract, err := bindContextUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableCaller{contract: contract}, nil
}

// NewContextUpgradeableTransactor creates a new write-only instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextUpgradeableTransactor, error) {
	contract, err := bindContextUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableTransactor{contract: contract}, nil
}

// NewContextUpgradeableFilterer creates a new log filterer instance of ContextUpgradeable, bound to a specific deployed contract.
func NewContextUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextUpgradeableFilterer, error) {
	contract, err := bindContextUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextUpgradeableFilterer{contract: contract}, nil
}

// bindContextUpgradeable binds a generic wrapper to an already deployed contract.
func bindContextUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.ContextUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.ContextUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContextUpgradeable *ContextUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContextUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContextUpgradeable *ContextUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContextUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ERC165UpgradeableMetaData contains all meta data concerning the ERC165Upgradeable contract.
var ERC165UpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// ERC165UpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC165UpgradeableMetaData.ABI instead.
var ERC165UpgradeableABI = ERC165UpgradeableMetaData.ABI

// Deprecated: Use ERC165UpgradeableMetaData.Sigs instead.
// ERC165UpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var ERC165UpgradeableFuncSigs = ERC165UpgradeableMetaData.Sigs

// ERC165Upgradeable is an auto generated Go binding around an Ethereum contract.
type ERC165Upgradeable struct {
	ERC165UpgradeableCaller     // Read-only binding to the contract
	ERC165UpgradeableTransactor // Write-only binding to the contract
	ERC165UpgradeableFilterer   // Log filterer for contract events
}

// ERC165UpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165UpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165UpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165UpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165UpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165UpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165UpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165UpgradeableSession struct {
	Contract     *ERC165Upgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ERC165UpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165UpgradeableCallerSession struct {
	Contract *ERC165UpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ERC165UpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165UpgradeableTransactorSession struct {
	Contract     *ERC165UpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ERC165UpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165UpgradeableRaw struct {
	Contract *ERC165Upgradeable // Generic contract binding to access the raw methods on
}

// ERC165UpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165UpgradeableCallerRaw struct {
	Contract *ERC165UpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC165UpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165UpgradeableTransactorRaw struct {
	Contract *ERC165UpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165Upgradeable creates a new instance of ERC165Upgradeable, bound to a specific deployed contract.
func NewERC165Upgradeable(address common.Address, backend bind.ContractBackend) (*ERC165Upgradeable, error) {
	contract, err := bindERC165Upgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165Upgradeable{ERC165UpgradeableCaller: ERC165UpgradeableCaller{contract: contract}, ERC165UpgradeableTransactor: ERC165UpgradeableTransactor{contract: contract}, ERC165UpgradeableFilterer: ERC165UpgradeableFilterer{contract: contract}}, nil
}

// NewERC165UpgradeableCaller creates a new read-only instance of ERC165Upgradeable, bound to a specific deployed contract.
func NewERC165UpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ERC165UpgradeableCaller, error) {
	contract, err := bindERC165Upgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165UpgradeableCaller{contract: contract}, nil
}

// NewERC165UpgradeableTransactor creates a new write-only instance of ERC165Upgradeable, bound to a specific deployed contract.
func NewERC165UpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC165UpgradeableTransactor, error) {
	contract, err := bindERC165Upgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165UpgradeableTransactor{contract: contract}, nil
}

// NewERC165UpgradeableFilterer creates a new log filterer instance of ERC165Upgradeable, bound to a specific deployed contract.
func NewERC165UpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC165UpgradeableFilterer, error) {
	contract, err := bindERC165Upgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165UpgradeableFilterer{contract: contract}, nil
}

// bindERC165Upgradeable binds a generic wrapper to an already deployed contract.
func bindERC165Upgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165UpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165Upgradeable *ERC165UpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165Upgradeable.Contract.ERC165UpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165Upgradeable *ERC165UpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165Upgradeable.Contract.ERC165UpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165Upgradeable *ERC165UpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165Upgradeable.Contract.ERC165UpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165Upgradeable *ERC165UpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165Upgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165Upgradeable *ERC165UpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165Upgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165Upgradeable *ERC165UpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165Upgradeable.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165Upgradeable *ERC165UpgradeableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC165Upgradeable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165Upgradeable *ERC165UpgradeableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165Upgradeable.Contract.SupportsInterface(&_ERC165Upgradeable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165Upgradeable *ERC165UpgradeableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165Upgradeable.Contract.SupportsInterface(&_ERC165Upgradeable.CallOpts, interfaceId)
}

// ERC1967UpgradeUpgradeableMetaData contains all meta data concerning the ERC1967UpgradeUpgradeable contract.
var ERC1967UpgradeUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"}]",
}

// ERC1967UpgradeUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC1967UpgradeUpgradeableMetaData.ABI instead.
var ERC1967UpgradeUpgradeableABI = ERC1967UpgradeUpgradeableMetaData.ABI

// ERC1967UpgradeUpgradeable is an auto generated Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeable struct {
	ERC1967UpgradeUpgradeableCaller     // Read-only binding to the contract
	ERC1967UpgradeUpgradeableTransactor // Write-only binding to the contract
	ERC1967UpgradeUpgradeableFilterer   // Log filterer for contract events
}

// ERC1967UpgradeUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC1967UpgradeUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC1967UpgradeUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC1967UpgradeUpgradeableSession struct {
	Contract     *ERC1967UpgradeUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ERC1967UpgradeUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC1967UpgradeUpgradeableCallerSession struct {
	Contract *ERC1967UpgradeUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// ERC1967UpgradeUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC1967UpgradeUpgradeableTransactorSession struct {
	Contract     *ERC1967UpgradeUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// ERC1967UpgradeUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableRaw struct {
	Contract *ERC1967UpgradeUpgradeable // Generic contract binding to access the raw methods on
}

// ERC1967UpgradeUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableCallerRaw struct {
	Contract *ERC1967UpgradeUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC1967UpgradeUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC1967UpgradeUpgradeableTransactorRaw struct {
	Contract *ERC1967UpgradeUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC1967UpgradeUpgradeable creates a new instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeable(address common.Address, backend bind.ContractBackend) (*ERC1967UpgradeUpgradeable, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeable{ERC1967UpgradeUpgradeableCaller: ERC1967UpgradeUpgradeableCaller{contract: contract}, ERC1967UpgradeUpgradeableTransactor: ERC1967UpgradeUpgradeableTransactor{contract: contract}, ERC1967UpgradeUpgradeableFilterer: ERC1967UpgradeUpgradeableFilterer{contract: contract}}, nil
}

// NewERC1967UpgradeUpgradeableCaller creates a new read-only instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ERC1967UpgradeUpgradeableCaller, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableCaller{contract: contract}, nil
}

// NewERC1967UpgradeUpgradeableTransactor creates a new write-only instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC1967UpgradeUpgradeableTransactor, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableTransactor{contract: contract}, nil
}

// NewERC1967UpgradeUpgradeableFilterer creates a new log filterer instance of ERC1967UpgradeUpgradeable, bound to a specific deployed contract.
func NewERC1967UpgradeUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC1967UpgradeUpgradeableFilterer, error) {
	contract, err := bindERC1967UpgradeUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableFilterer{contract: contract}, nil
}

// bindERC1967UpgradeUpgradeable binds a generic wrapper to an already deployed contract.
func bindERC1967UpgradeUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC1967UpgradeUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967UpgradeUpgradeable.Contract.ERC1967UpgradeUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.ERC1967UpgradeUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.ERC1967UpgradeUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC1967UpgradeUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC1967UpgradeUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// ERC1967UpgradeUpgradeableAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableAdminChangedIterator struct {
	Event *ERC1967UpgradeUpgradeableAdminChanged // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeUpgradeableAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableAdminChanged)
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
		it.Event = new(ERC1967UpgradeUpgradeableAdminChanged)
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
func (it *ERC1967UpgradeUpgradeableAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableAdminChanged represents a AdminChanged event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*ERC1967UpgradeUpgradeableAdminChangedIterator, error) {

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableAdminChangedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableAdminChanged) (event.Subscription, error) {

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableAdminChanged)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseAdminChanged(log types.Log) (*ERC1967UpgradeUpgradeableAdminChanged, error) {
	event := new(ERC1967UpgradeUpgradeableAdminChanged)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UpgradeUpgradeableBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableBeaconUpgradedIterator struct {
	Event *ERC1967UpgradeUpgradeableBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeUpgradeableBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableBeaconUpgraded)
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
		it.Event = new(ERC1967UpgradeUpgradeableBeaconUpgraded)
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
func (it *ERC1967UpgradeUpgradeableBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableBeaconUpgraded represents a BeaconUpgraded event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*ERC1967UpgradeUpgradeableBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableBeaconUpgradedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableBeaconUpgraded)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseBeaconUpgraded(log types.Log) (*ERC1967UpgradeUpgradeableBeaconUpgraded, error) {
	event := new(ERC1967UpgradeUpgradeableBeaconUpgraded)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC1967UpgradeUpgradeableUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableUpgradedIterator struct {
	Event *ERC1967UpgradeUpgradeableUpgraded // Event containing the contract specifics and raw log

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
func (it *ERC1967UpgradeUpgradeableUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC1967UpgradeUpgradeableUpgraded)
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
		it.Event = new(ERC1967UpgradeUpgradeableUpgraded)
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
func (it *ERC1967UpgradeUpgradeableUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC1967UpgradeUpgradeableUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC1967UpgradeUpgradeableUpgraded represents a Upgraded event raised by the ERC1967UpgradeUpgradeable contract.
type ERC1967UpgradeUpgradeableUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*ERC1967UpgradeUpgradeableUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ERC1967UpgradeUpgradeableUpgradedIterator{contract: _ERC1967UpgradeUpgradeable.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *ERC1967UpgradeUpgradeableUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _ERC1967UpgradeUpgradeable.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC1967UpgradeUpgradeableUpgraded)
				if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_ERC1967UpgradeUpgradeable *ERC1967UpgradeUpgradeableFilterer) ParseUpgraded(log types.Log) (*ERC1967UpgradeUpgradeableUpgraded, error) {
	event := new(ERC1967UpgradeUpgradeableUpgraded)
	if err := _ERC1967UpgradeUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20UpgradeableMetaData contains all meta data concerning the ERC20Upgradeable contract.
var ERC20UpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"a457c2d7": "decreaseAllowance(address,uint256)",
		"39509351": "increaseAllowance(address,uint256)",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
	Bin: "0x608060405234801561001057600080fd5b506108ee806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012357806370a082311461013657806395d89b4114610149578063a457c2d714610151578063a9059cbb14610164578063dd62ed3e14610177576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ef57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b66101b0565b6040516100c391906107e5565b60405180910390f35b6100df6100da3660046107bc565b610242565b60405190151581526020016100c3565b6035545b6040519081526020016100c3565b6100df61010f366004610781565b610258565b604051601281526020016100c3565b6100df6101313660046107bc565b61030e565b6100f361014436600461072e565b610345565b6100b6610364565b6100df61015f3660046107bc565b610373565b6100df6101723660046107bc565b61040e565b6100f361018536600461074f565b6001600160a01b03918216600090815260346020908152604080832093909416825291909152205490565b6060603680546101bf90610867565b80601f01602080910402602001604051908101604052809291908181526020018280546101eb90610867565b80156102385780601f1061020d57610100808354040283529160200191610238565b820191906000526020600020905b81548152906001019060200180831161021b57829003601f168201915b5050505050905090565b600061024f33848461041b565b50600192915050565b600061026584848461053f565b6001600160a01b0384166000908152603460209081526040808320338452909152902054828110156102ef5760405162461bcd60e51b815260206004820152602860248201527f45524332303a207472616e7366657220616d6f756e74206578636565647320616044820152676c6c6f77616e636560c01b60648201526084015b60405180910390fd5b61030385336102fe8685610850565b61041b565b506001949350505050565b3360008181526034602090815260408083206001600160a01b0387168452909152812054909161024f9185906102fe908690610838565b6001600160a01b0381166000908152603360205260409020545b919050565b6060603780546101bf90610867565b3360009081526034602090815260408083206001600160a01b0386168452909152812054828110156103f55760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084016102e6565b61040433856102fe8685610850565b5060019392505050565b600061024f33848461053f565b6001600160a01b03831661047d5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016102e6565b6001600160a01b0382166104de5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016102e6565b6001600160a01b0383811660008181526034602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b6001600160a01b0383166105a35760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016102e6565b6001600160a01b0382166106055760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016102e6565b6001600160a01b0383166000908152603360205260409020548181101561067d5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016102e6565b6106878282610850565b6001600160a01b0380861660009081526033602052604080822093909355908516815290812080548492906106bd908490610838565b92505081905550826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161070991815260200190565b60405180910390a350505050565b80356001600160a01b038116811461035f57600080fd5b60006020828403121561073f578081fd5b61074882610717565b9392505050565b60008060408385031215610761578081fd5b61076a83610717565b915061077860208401610717565b90509250929050565b600080600060608486031215610795578081fd5b61079e84610717565b92506107ac60208501610717565b9150604084013590509250925092565b600080604083850312156107ce578182fd5b6107d783610717565b946020939093013593505050565b6000602080835283518082850152825b81811015610811578581018301518582016040015282016107f5565b818111156108225783604083870101525b50601f01601f1916929092016040019392505050565b6000821982111561084b5761084b6108a2565b500190565b600082821015610862576108626108a2565b500390565b60028104600182168061087b57607f821691505b6020821081141561089c57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fdfea264697066735822122060c420c9426e05080da8fa019ca2825f92e2eb6323ec4144442cb15072dfdb2864736f6c63430008020033",
}

// ERC20UpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20UpgradeableMetaData.ABI instead.
var ERC20UpgradeableABI = ERC20UpgradeableMetaData.ABI

// Deprecated: Use ERC20UpgradeableMetaData.Sigs instead.
// ERC20UpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var ERC20UpgradeableFuncSigs = ERC20UpgradeableMetaData.Sigs

// ERC20UpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20UpgradeableMetaData.Bin instead.
var ERC20UpgradeableBin = ERC20UpgradeableMetaData.Bin

// DeployERC20Upgradeable deploys a new Ethereum contract, binding an instance of ERC20Upgradeable to it.
func DeployERC20Upgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Upgradeable, error) {
	parsed, err := ERC20UpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20UpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Upgradeable{ERC20UpgradeableCaller: ERC20UpgradeableCaller{contract: contract}, ERC20UpgradeableTransactor: ERC20UpgradeableTransactor{contract: contract}, ERC20UpgradeableFilterer: ERC20UpgradeableFilterer{contract: contract}}, nil
}

// ERC20Upgradeable is an auto generated Go binding around an Ethereum contract.
type ERC20Upgradeable struct {
	ERC20UpgradeableCaller     // Read-only binding to the contract
	ERC20UpgradeableTransactor // Write-only binding to the contract
	ERC20UpgradeableFilterer   // Log filterer for contract events
}

// ERC20UpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20UpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20UpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20UpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20UpgradeableSession struct {
	Contract     *ERC20Upgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20UpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20UpgradeableCallerSession struct {
	Contract *ERC20UpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ERC20UpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20UpgradeableTransactorSession struct {
	Contract     *ERC20UpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ERC20UpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20UpgradeableRaw struct {
	Contract *ERC20Upgradeable // Generic contract binding to access the raw methods on
}

// ERC20UpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20UpgradeableCallerRaw struct {
	Contract *ERC20UpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20UpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20UpgradeableTransactorRaw struct {
	Contract *ERC20UpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Upgradeable creates a new instance of ERC20Upgradeable, bound to a specific deployed contract.
func NewERC20Upgradeable(address common.Address, backend bind.ContractBackend) (*ERC20Upgradeable, error) {
	contract, err := bindERC20Upgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Upgradeable{ERC20UpgradeableCaller: ERC20UpgradeableCaller{contract: contract}, ERC20UpgradeableTransactor: ERC20UpgradeableTransactor{contract: contract}, ERC20UpgradeableFilterer: ERC20UpgradeableFilterer{contract: contract}}, nil
}

// NewERC20UpgradeableCaller creates a new read-only instance of ERC20Upgradeable, bound to a specific deployed contract.
func NewERC20UpgradeableCaller(address common.Address, caller bind.ContractCaller) (*ERC20UpgradeableCaller, error) {
	contract, err := bindERC20Upgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20UpgradeableCaller{contract: contract}, nil
}

// NewERC20UpgradeableTransactor creates a new write-only instance of ERC20Upgradeable, bound to a specific deployed contract.
func NewERC20UpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20UpgradeableTransactor, error) {
	contract, err := bindERC20Upgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20UpgradeableTransactor{contract: contract}, nil
}

// NewERC20UpgradeableFilterer creates a new log filterer instance of ERC20Upgradeable, bound to a specific deployed contract.
func NewERC20UpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20UpgradeableFilterer, error) {
	contract, err := bindERC20Upgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20UpgradeableFilterer{contract: contract}, nil
}

// bindERC20Upgradeable binds a generic wrapper to an already deployed contract.
func bindERC20Upgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20UpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Upgradeable *ERC20UpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Upgradeable.Contract.ERC20UpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Upgradeable *ERC20UpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.ERC20UpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Upgradeable *ERC20UpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.ERC20UpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Upgradeable *ERC20UpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Upgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Upgradeable *ERC20UpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Upgradeable *ERC20UpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Upgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Upgradeable.Contract.Allowance(&_ERC20Upgradeable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Upgradeable.Contract.Allowance(&_ERC20Upgradeable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Upgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Upgradeable.Contract.BalanceOf(&_ERC20Upgradeable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Upgradeable.Contract.BalanceOf(&_ERC20Upgradeable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Upgradeable *ERC20UpgradeableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Upgradeable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Upgradeable *ERC20UpgradeableSession) Decimals() (uint8, error) {
	return _ERC20Upgradeable.Contract.Decimals(&_ERC20Upgradeable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Upgradeable *ERC20UpgradeableCallerSession) Decimals() (uint8, error) {
	return _ERC20Upgradeable.Contract.Decimals(&_ERC20Upgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Upgradeable *ERC20UpgradeableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Upgradeable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Upgradeable *ERC20UpgradeableSession) Name() (string, error) {
	return _ERC20Upgradeable.Contract.Name(&_ERC20Upgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Upgradeable *ERC20UpgradeableCallerSession) Name() (string, error) {
	return _ERC20Upgradeable.Contract.Name(&_ERC20Upgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Upgradeable *ERC20UpgradeableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Upgradeable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Upgradeable *ERC20UpgradeableSession) Symbol() (string, error) {
	return _ERC20Upgradeable.Contract.Symbol(&_ERC20Upgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Upgradeable *ERC20UpgradeableCallerSession) Symbol() (string, error) {
	return _ERC20Upgradeable.Contract.Symbol(&_ERC20Upgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Upgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableSession) TotalSupply() (*big.Int, error) {
	return _ERC20Upgradeable.Contract.TotalSupply(&_ERC20Upgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Upgradeable *ERC20UpgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Upgradeable.Contract.TotalSupply(&_ERC20Upgradeable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.Approve(&_ERC20Upgradeable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.Approve(&_ERC20Upgradeable.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.DecreaseAllowance(&_ERC20Upgradeable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.DecreaseAllowance(&_ERC20Upgradeable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.IncreaseAllowance(&_ERC20Upgradeable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.IncreaseAllowance(&_ERC20Upgradeable.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.Transfer(&_ERC20Upgradeable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.Transfer(&_ERC20Upgradeable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.TransferFrom(&_ERC20Upgradeable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20Upgradeable *ERC20UpgradeableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Upgradeable.Contract.TransferFrom(&_ERC20Upgradeable.TransactOpts, sender, recipient, amount)
}

// ERC20UpgradeableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Upgradeable contract.
type ERC20UpgradeableApprovalIterator struct {
	Event *ERC20UpgradeableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20UpgradeableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20UpgradeableApproval)
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
		it.Event = new(ERC20UpgradeableApproval)
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
func (it *ERC20UpgradeableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20UpgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20UpgradeableApproval represents a Approval event raised by the ERC20Upgradeable contract.
type ERC20UpgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Upgradeable *ERC20UpgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20UpgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Upgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20UpgradeableApprovalIterator{contract: _ERC20Upgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Upgradeable *ERC20UpgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20UpgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Upgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20UpgradeableApproval)
				if err := _ERC20Upgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Upgradeable *ERC20UpgradeableFilterer) ParseApproval(log types.Log) (*ERC20UpgradeableApproval, error) {
	event := new(ERC20UpgradeableApproval)
	if err := _ERC20Upgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20UpgradeableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Upgradeable contract.
type ERC20UpgradeableTransferIterator struct {
	Event *ERC20UpgradeableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20UpgradeableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20UpgradeableTransfer)
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
		it.Event = new(ERC20UpgradeableTransfer)
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
func (it *ERC20UpgradeableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20UpgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20UpgradeableTransfer represents a Transfer event raised by the ERC20Upgradeable contract.
type ERC20UpgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Upgradeable *ERC20UpgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20UpgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Upgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20UpgradeableTransferIterator{contract: _ERC20Upgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Upgradeable *ERC20UpgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20UpgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Upgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20UpgradeableTransfer)
				if err := _ERC20Upgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Upgradeable *ERC20UpgradeableFilterer) ParseTransfer(log types.Log) (*ERC20UpgradeableTransfer, error) {
	event := new(ERC20UpgradeableTransfer)
	if err := _ERC20Upgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccessControlUpgradeableMetaData contains all meta data concerning the IAccessControlUpgradeable contract.
var IAccessControlUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"248a9ca3": "getRoleAdmin(bytes32)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"d547741f": "revokeRole(bytes32,address)",
	},
}

// IAccessControlUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccessControlUpgradeableMetaData.ABI instead.
var IAccessControlUpgradeableABI = IAccessControlUpgradeableMetaData.ABI

// Deprecated: Use IAccessControlUpgradeableMetaData.Sigs instead.
// IAccessControlUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var IAccessControlUpgradeableFuncSigs = IAccessControlUpgradeableMetaData.Sigs

// IAccessControlUpgradeable is an auto generated Go binding around an Ethereum contract.
type IAccessControlUpgradeable struct {
	IAccessControlUpgradeableCaller     // Read-only binding to the contract
	IAccessControlUpgradeableTransactor // Write-only binding to the contract
	IAccessControlUpgradeableFilterer   // Log filterer for contract events
}

// IAccessControlUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccessControlUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccessControlUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccessControlUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccessControlUpgradeableSession struct {
	Contract     *IAccessControlUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IAccessControlUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccessControlUpgradeableCallerSession struct {
	Contract *IAccessControlUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// IAccessControlUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccessControlUpgradeableTransactorSession struct {
	Contract     *IAccessControlUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// IAccessControlUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccessControlUpgradeableRaw struct {
	Contract *IAccessControlUpgradeable // Generic contract binding to access the raw methods on
}

// IAccessControlUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccessControlUpgradeableCallerRaw struct {
	Contract *IAccessControlUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IAccessControlUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccessControlUpgradeableTransactorRaw struct {
	Contract *IAccessControlUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccessControlUpgradeable creates a new instance of IAccessControlUpgradeable, bound to a specific deployed contract.
func NewIAccessControlUpgradeable(address common.Address, backend bind.ContractBackend) (*IAccessControlUpgradeable, error) {
	contract, err := bindIAccessControlUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccessControlUpgradeable{IAccessControlUpgradeableCaller: IAccessControlUpgradeableCaller{contract: contract}, IAccessControlUpgradeableTransactor: IAccessControlUpgradeableTransactor{contract: contract}, IAccessControlUpgradeableFilterer: IAccessControlUpgradeableFilterer{contract: contract}}, nil
}

// NewIAccessControlUpgradeableCaller creates a new read-only instance of IAccessControlUpgradeable, bound to a specific deployed contract.
func NewIAccessControlUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IAccessControlUpgradeableCaller, error) {
	contract, err := bindIAccessControlUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlUpgradeableCaller{contract: contract}, nil
}

// NewIAccessControlUpgradeableTransactor creates a new write-only instance of IAccessControlUpgradeable, bound to a specific deployed contract.
func NewIAccessControlUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccessControlUpgradeableTransactor, error) {
	contract, err := bindIAccessControlUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlUpgradeableTransactor{contract: contract}, nil
}

// NewIAccessControlUpgradeableFilterer creates a new log filterer instance of IAccessControlUpgradeable, bound to a specific deployed contract.
func NewIAccessControlUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccessControlUpgradeableFilterer, error) {
	contract, err := bindIAccessControlUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccessControlUpgradeableFilterer{contract: contract}, nil
}

// bindIAccessControlUpgradeable binds a generic wrapper to an already deployed contract.
func bindIAccessControlUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAccessControlUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControlUpgradeable *IAccessControlUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControlUpgradeable.Contract.IAccessControlUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControlUpgradeable *IAccessControlUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControlUpgradeable.Contract.IAccessControlUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControlUpgradeable *IAccessControlUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControlUpgradeable.Contract.IAccessControlUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControlUpgradeable *IAccessControlUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControlUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControlUpgradeable *IAccessControlUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControlUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControlUpgradeable *IAccessControlUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControlUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControlUpgradeable *IAccessControlUpgradeableCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IAccessControlUpgradeable.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControlUpgradeable *IAccessControlUpgradeableSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControlUpgradeable.Contract.GetRoleAdmin(&_IAccessControlUpgradeable.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControlUpgradeable *IAccessControlUpgradeableCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControlUpgradeable.Contract.GetRoleAdmin(&_IAccessControlUpgradeable.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControlUpgradeable *IAccessControlUpgradeableCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IAccessControlUpgradeable.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControlUpgradeable *IAccessControlUpgradeableSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControlUpgradeable.Contract.HasRole(&_IAccessControlUpgradeable.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControlUpgradeable *IAccessControlUpgradeableCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControlUpgradeable.Contract.HasRole(&_IAccessControlUpgradeable.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControlUpgradeable *IAccessControlUpgradeableTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlUpgradeable.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControlUpgradeable *IAccessControlUpgradeableSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlUpgradeable.Contract.GrantRole(&_IAccessControlUpgradeable.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControlUpgradeable *IAccessControlUpgradeableTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlUpgradeable.Contract.GrantRole(&_IAccessControlUpgradeable.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IAccessControlUpgradeable *IAccessControlUpgradeableTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlUpgradeable.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IAccessControlUpgradeable *IAccessControlUpgradeableSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlUpgradeable.Contract.RenounceRole(&_IAccessControlUpgradeable.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IAccessControlUpgradeable *IAccessControlUpgradeableTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlUpgradeable.Contract.RenounceRole(&_IAccessControlUpgradeable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControlUpgradeable *IAccessControlUpgradeableTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlUpgradeable.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControlUpgradeable *IAccessControlUpgradeableSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlUpgradeable.Contract.RevokeRole(&_IAccessControlUpgradeable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControlUpgradeable *IAccessControlUpgradeableTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControlUpgradeable.Contract.RevokeRole(&_IAccessControlUpgradeable.TransactOpts, role, account)
}

// IBeaconUpgradeableMetaData contains all meta data concerning the IBeaconUpgradeable contract.
var IBeaconUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"5c60da1b": "implementation()",
	},
}

// IBeaconUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IBeaconUpgradeableMetaData.ABI instead.
var IBeaconUpgradeableABI = IBeaconUpgradeableMetaData.ABI

// Deprecated: Use IBeaconUpgradeableMetaData.Sigs instead.
// IBeaconUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var IBeaconUpgradeableFuncSigs = IBeaconUpgradeableMetaData.Sigs

// IBeaconUpgradeable is an auto generated Go binding around an Ethereum contract.
type IBeaconUpgradeable struct {
	IBeaconUpgradeableCaller     // Read-only binding to the contract
	IBeaconUpgradeableTransactor // Write-only binding to the contract
	IBeaconUpgradeableFilterer   // Log filterer for contract events
}

// IBeaconUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBeaconUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBeaconUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBeaconUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBeaconUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBeaconUpgradeableSession struct {
	Contract     *IBeaconUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IBeaconUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBeaconUpgradeableCallerSession struct {
	Contract *IBeaconUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IBeaconUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBeaconUpgradeableTransactorSession struct {
	Contract     *IBeaconUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IBeaconUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBeaconUpgradeableRaw struct {
	Contract *IBeaconUpgradeable // Generic contract binding to access the raw methods on
}

// IBeaconUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBeaconUpgradeableCallerRaw struct {
	Contract *IBeaconUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IBeaconUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBeaconUpgradeableTransactorRaw struct {
	Contract *IBeaconUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBeaconUpgradeable creates a new instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeable(address common.Address, backend bind.ContractBackend) (*IBeaconUpgradeable, error) {
	contract, err := bindIBeaconUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeable{IBeaconUpgradeableCaller: IBeaconUpgradeableCaller{contract: contract}, IBeaconUpgradeableTransactor: IBeaconUpgradeableTransactor{contract: contract}, IBeaconUpgradeableFilterer: IBeaconUpgradeableFilterer{contract: contract}}, nil
}

// NewIBeaconUpgradeableCaller creates a new read-only instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IBeaconUpgradeableCaller, error) {
	contract, err := bindIBeaconUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeableCaller{contract: contract}, nil
}

// NewIBeaconUpgradeableTransactor creates a new write-only instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IBeaconUpgradeableTransactor, error) {
	contract, err := bindIBeaconUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeableTransactor{contract: contract}, nil
}

// NewIBeaconUpgradeableFilterer creates a new log filterer instance of IBeaconUpgradeable, bound to a specific deployed contract.
func NewIBeaconUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IBeaconUpgradeableFilterer, error) {
	contract, err := bindIBeaconUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBeaconUpgradeableFilterer{contract: contract}, nil
}

// bindIBeaconUpgradeable binds a generic wrapper to an already deployed contract.
func bindIBeaconUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBeaconUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBeaconUpgradeable *IBeaconUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBeaconUpgradeable.Contract.IBeaconUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBeaconUpgradeable *IBeaconUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.IBeaconUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBeaconUpgradeable *IBeaconUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.IBeaconUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBeaconUpgradeable *IBeaconUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBeaconUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBeaconUpgradeable *IBeaconUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBeaconUpgradeable *IBeaconUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBeaconUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeaconUpgradeable *IBeaconUpgradeableCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBeaconUpgradeable.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeaconUpgradeable *IBeaconUpgradeableSession) Implementation() (common.Address, error) {
	return _IBeaconUpgradeable.Contract.Implementation(&_IBeaconUpgradeable.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_IBeaconUpgradeable *IBeaconUpgradeableCallerSession) Implementation() (common.Address, error) {
	return _IBeaconUpgradeable.Contract.Implementation(&_IBeaconUpgradeable.CallOpts)
}

// IERC165UpgradeableMetaData contains all meta data concerning the IERC165Upgradeable contract.
var IERC165UpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"01ffc9a7": "supportsInterface(bytes4)",
	},
}

// IERC165UpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC165UpgradeableMetaData.ABI instead.
var IERC165UpgradeableABI = IERC165UpgradeableMetaData.ABI

// Deprecated: Use IERC165UpgradeableMetaData.Sigs instead.
// IERC165UpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var IERC165UpgradeableFuncSigs = IERC165UpgradeableMetaData.Sigs

// IERC165Upgradeable is an auto generated Go binding around an Ethereum contract.
type IERC165Upgradeable struct {
	IERC165UpgradeableCaller     // Read-only binding to the contract
	IERC165UpgradeableTransactor // Write-only binding to the contract
	IERC165UpgradeableFilterer   // Log filterer for contract events
}

// IERC165UpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165UpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165UpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165UpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165UpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165UpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165UpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165UpgradeableSession struct {
	Contract     *IERC165Upgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IERC165UpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165UpgradeableCallerSession struct {
	Contract *IERC165UpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IERC165UpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165UpgradeableTransactorSession struct {
	Contract     *IERC165UpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IERC165UpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165UpgradeableRaw struct {
	Contract *IERC165Upgradeable // Generic contract binding to access the raw methods on
}

// IERC165UpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165UpgradeableCallerRaw struct {
	Contract *IERC165UpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC165UpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165UpgradeableTransactorRaw struct {
	Contract *IERC165UpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165Upgradeable creates a new instance of IERC165Upgradeable, bound to a specific deployed contract.
func NewIERC165Upgradeable(address common.Address, backend bind.ContractBackend) (*IERC165Upgradeable, error) {
	contract, err := bindIERC165Upgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165Upgradeable{IERC165UpgradeableCaller: IERC165UpgradeableCaller{contract: contract}, IERC165UpgradeableTransactor: IERC165UpgradeableTransactor{contract: contract}, IERC165UpgradeableFilterer: IERC165UpgradeableFilterer{contract: contract}}, nil
}

// NewIERC165UpgradeableCaller creates a new read-only instance of IERC165Upgradeable, bound to a specific deployed contract.
func NewIERC165UpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IERC165UpgradeableCaller, error) {
	contract, err := bindIERC165Upgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165UpgradeableCaller{contract: contract}, nil
}

// NewIERC165UpgradeableTransactor creates a new write-only instance of IERC165Upgradeable, bound to a specific deployed contract.
func NewIERC165UpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC165UpgradeableTransactor, error) {
	contract, err := bindIERC165Upgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165UpgradeableTransactor{contract: contract}, nil
}

// NewIERC165UpgradeableFilterer creates a new log filterer instance of IERC165Upgradeable, bound to a specific deployed contract.
func NewIERC165UpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC165UpgradeableFilterer, error) {
	contract, err := bindIERC165Upgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165UpgradeableFilterer{contract: contract}, nil
}

// bindIERC165Upgradeable binds a generic wrapper to an already deployed contract.
func bindIERC165Upgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165UpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165Upgradeable *IERC165UpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165Upgradeable.Contract.IERC165UpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165Upgradeable *IERC165UpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165Upgradeable.Contract.IERC165UpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165Upgradeable *IERC165UpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165Upgradeable.Contract.IERC165UpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165Upgradeable *IERC165UpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165Upgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165Upgradeable *IERC165UpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165Upgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165Upgradeable *IERC165UpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165Upgradeable.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165Upgradeable *IERC165UpgradeableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165Upgradeable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165Upgradeable *IERC165UpgradeableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165Upgradeable.Contract.SupportsInterface(&_IERC165Upgradeable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165Upgradeable *IERC165UpgradeableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165Upgradeable.Contract.SupportsInterface(&_IERC165Upgradeable.CallOpts, interfaceId)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// Deprecated: Use IERC20MetaData.Sigs instead.
// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = IERC20MetaData.Sigs

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataUpgradeableMetaData contains all meta data concerning the IERC20MetadataUpgradeable contract.
var IERC20MetadataUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"313ce567": "decimals()",
		"06fdde03": "name()",
		"95d89b41": "symbol()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20MetadataUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetadataUpgradeableMetaData.ABI instead.
var IERC20MetadataUpgradeableABI = IERC20MetadataUpgradeableMetaData.ABI

// Deprecated: Use IERC20MetadataUpgradeableMetaData.Sigs instead.
// IERC20MetadataUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var IERC20MetadataUpgradeableFuncSigs = IERC20MetadataUpgradeableMetaData.Sigs

// IERC20MetadataUpgradeable is an auto generated Go binding around an Ethereum contract.
type IERC20MetadataUpgradeable struct {
	IERC20MetadataUpgradeableCaller     // Read-only binding to the contract
	IERC20MetadataUpgradeableTransactor // Write-only binding to the contract
	IERC20MetadataUpgradeableFilterer   // Log filterer for contract events
}

// IERC20MetadataUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20MetadataUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20MetadataUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20MetadataUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20MetadataUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20MetadataUpgradeableSession struct {
	Contract     *IERC20MetadataUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IERC20MetadataUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20MetadataUpgradeableCallerSession struct {
	Contract *IERC20MetadataUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// IERC20MetadataUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20MetadataUpgradeableTransactorSession struct {
	Contract     *IERC20MetadataUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// IERC20MetadataUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20MetadataUpgradeableRaw struct {
	Contract *IERC20MetadataUpgradeable // Generic contract binding to access the raw methods on
}

// IERC20MetadataUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20MetadataUpgradeableCallerRaw struct {
	Contract *IERC20MetadataUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20MetadataUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20MetadataUpgradeableTransactorRaw struct {
	Contract *IERC20MetadataUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20MetadataUpgradeable creates a new instance of IERC20MetadataUpgradeable, bound to a specific deployed contract.
func NewIERC20MetadataUpgradeable(address common.Address, backend bind.ContractBackend) (*IERC20MetadataUpgradeable, error) {
	contract, err := bindIERC20MetadataUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataUpgradeable{IERC20MetadataUpgradeableCaller: IERC20MetadataUpgradeableCaller{contract: contract}, IERC20MetadataUpgradeableTransactor: IERC20MetadataUpgradeableTransactor{contract: contract}, IERC20MetadataUpgradeableFilterer: IERC20MetadataUpgradeableFilterer{contract: contract}}, nil
}

// NewIERC20MetadataUpgradeableCaller creates a new read-only instance of IERC20MetadataUpgradeable, bound to a specific deployed contract.
func NewIERC20MetadataUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IERC20MetadataUpgradeableCaller, error) {
	contract, err := bindIERC20MetadataUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataUpgradeableCaller{contract: contract}, nil
}

// NewIERC20MetadataUpgradeableTransactor creates a new write-only instance of IERC20MetadataUpgradeable, bound to a specific deployed contract.
func NewIERC20MetadataUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20MetadataUpgradeableTransactor, error) {
	contract, err := bindIERC20MetadataUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataUpgradeableTransactor{contract: contract}, nil
}

// NewIERC20MetadataUpgradeableFilterer creates a new log filterer instance of IERC20MetadataUpgradeable, bound to a specific deployed contract.
func NewIERC20MetadataUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20MetadataUpgradeableFilterer, error) {
	contract, err := bindIERC20MetadataUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataUpgradeableFilterer{contract: contract}, nil
}

// bindIERC20MetadataUpgradeable binds a generic wrapper to an already deployed contract.
func bindIERC20MetadataUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20MetadataUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20MetadataUpgradeable.Contract.IERC20MetadataUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20MetadataUpgradeable.Contract.IERC20MetadataUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20MetadataUpgradeable.Contract.IERC20MetadataUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20MetadataUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20MetadataUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20MetadataUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20MetadataUpgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20MetadataUpgradeable.Contract.Allowance(&_IERC20MetadataUpgradeable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20MetadataUpgradeable.Contract.Allowance(&_IERC20MetadataUpgradeable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20MetadataUpgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20MetadataUpgradeable.Contract.BalanceOf(&_IERC20MetadataUpgradeable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20MetadataUpgradeable.Contract.BalanceOf(&_IERC20MetadataUpgradeable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20MetadataUpgradeable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableSession) Decimals() (uint8, error) {
	return _IERC20MetadataUpgradeable.Contract.Decimals(&_IERC20MetadataUpgradeable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableCallerSession) Decimals() (uint8, error) {
	return _IERC20MetadataUpgradeable.Contract.Decimals(&_IERC20MetadataUpgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20MetadataUpgradeable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableSession) Name() (string, error) {
	return _IERC20MetadataUpgradeable.Contract.Name(&_IERC20MetadataUpgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableCallerSession) Name() (string, error) {
	return _IERC20MetadataUpgradeable.Contract.Name(&_IERC20MetadataUpgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20MetadataUpgradeable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableSession) Symbol() (string, error) {
	return _IERC20MetadataUpgradeable.Contract.Symbol(&_IERC20MetadataUpgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableCallerSession) Symbol() (string, error) {
	return _IERC20MetadataUpgradeable.Contract.Symbol(&_IERC20MetadataUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20MetadataUpgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableSession) TotalSupply() (*big.Int, error) {
	return _IERC20MetadataUpgradeable.Contract.TotalSupply(&_IERC20MetadataUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20MetadataUpgradeable.Contract.TotalSupply(&_IERC20MetadataUpgradeable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MetadataUpgradeable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MetadataUpgradeable.Contract.Approve(&_IERC20MetadataUpgradeable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MetadataUpgradeable.Contract.Approve(&_IERC20MetadataUpgradeable.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MetadataUpgradeable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MetadataUpgradeable.Contract.Transfer(&_IERC20MetadataUpgradeable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MetadataUpgradeable.Contract.Transfer(&_IERC20MetadataUpgradeable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MetadataUpgradeable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MetadataUpgradeable.Contract.TransferFrom(&_IERC20MetadataUpgradeable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20MetadataUpgradeable.Contract.TransferFrom(&_IERC20MetadataUpgradeable.TransactOpts, sender, recipient, amount)
}

// IERC20MetadataUpgradeableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20MetadataUpgradeable contract.
type IERC20MetadataUpgradeableApprovalIterator struct {
	Event *IERC20MetadataUpgradeableApproval // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataUpgradeableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataUpgradeableApproval)
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
		it.Event = new(IERC20MetadataUpgradeableApproval)
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
func (it *IERC20MetadataUpgradeableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataUpgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataUpgradeableApproval represents a Approval event raised by the IERC20MetadataUpgradeable contract.
type IERC20MetadataUpgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20MetadataUpgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20MetadataUpgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataUpgradeableApprovalIterator{contract: _IERC20MetadataUpgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20MetadataUpgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20MetadataUpgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataUpgradeableApproval)
				if err := _IERC20MetadataUpgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableFilterer) ParseApproval(log types.Log) (*IERC20MetadataUpgradeableApproval, error) {
	event := new(IERC20MetadataUpgradeableApproval)
	if err := _IERC20MetadataUpgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20MetadataUpgradeableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20MetadataUpgradeable contract.
type IERC20MetadataUpgradeableTransferIterator struct {
	Event *IERC20MetadataUpgradeableTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20MetadataUpgradeableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20MetadataUpgradeableTransfer)
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
		it.Event = new(IERC20MetadataUpgradeableTransfer)
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
func (it *IERC20MetadataUpgradeableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20MetadataUpgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20MetadataUpgradeableTransfer represents a Transfer event raised by the IERC20MetadataUpgradeable contract.
type IERC20MetadataUpgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20MetadataUpgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20MetadataUpgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20MetadataUpgradeableTransferIterator{contract: _IERC20MetadataUpgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20MetadataUpgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20MetadataUpgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20MetadataUpgradeableTransfer)
				if err := _IERC20MetadataUpgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20MetadataUpgradeable *IERC20MetadataUpgradeableFilterer) ParseTransfer(log types.Log) (*IERC20MetadataUpgradeableTransfer, error) {
	event := new(IERC20MetadataUpgradeableTransfer)
	if err := _IERC20MetadataUpgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20UpgradeableMetaData contains all meta data concerning the IERC20Upgradeable contract.
var IERC20UpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// IERC20UpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20UpgradeableMetaData.ABI instead.
var IERC20UpgradeableABI = IERC20UpgradeableMetaData.ABI

// Deprecated: Use IERC20UpgradeableMetaData.Sigs instead.
// IERC20UpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var IERC20UpgradeableFuncSigs = IERC20UpgradeableMetaData.Sigs

// IERC20Upgradeable is an auto generated Go binding around an Ethereum contract.
type IERC20Upgradeable struct {
	IERC20UpgradeableCaller     // Read-only binding to the contract
	IERC20UpgradeableTransactor // Write-only binding to the contract
	IERC20UpgradeableFilterer   // Log filterer for contract events
}

// IERC20UpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20UpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20UpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20UpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20UpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20UpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20UpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20UpgradeableSession struct {
	Contract     *IERC20Upgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC20UpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20UpgradeableCallerSession struct {
	Contract *IERC20UpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IERC20UpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20UpgradeableTransactorSession struct {
	Contract     *IERC20UpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IERC20UpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20UpgradeableRaw struct {
	Contract *IERC20Upgradeable // Generic contract binding to access the raw methods on
}

// IERC20UpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20UpgradeableCallerRaw struct {
	Contract *IERC20UpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20UpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20UpgradeableTransactorRaw struct {
	Contract *IERC20UpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Upgradeable creates a new instance of IERC20Upgradeable, bound to a specific deployed contract.
func NewIERC20Upgradeable(address common.Address, backend bind.ContractBackend) (*IERC20Upgradeable, error) {
	contract, err := bindIERC20Upgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Upgradeable{IERC20UpgradeableCaller: IERC20UpgradeableCaller{contract: contract}, IERC20UpgradeableTransactor: IERC20UpgradeableTransactor{contract: contract}, IERC20UpgradeableFilterer: IERC20UpgradeableFilterer{contract: contract}}, nil
}

// NewIERC20UpgradeableCaller creates a new read-only instance of IERC20Upgradeable, bound to a specific deployed contract.
func NewIERC20UpgradeableCaller(address common.Address, caller bind.ContractCaller) (*IERC20UpgradeableCaller, error) {
	contract, err := bindIERC20Upgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableCaller{contract: contract}, nil
}

// NewIERC20UpgradeableTransactor creates a new write-only instance of IERC20Upgradeable, bound to a specific deployed contract.
func NewIERC20UpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20UpgradeableTransactor, error) {
	contract, err := bindIERC20Upgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableTransactor{contract: contract}, nil
}

// NewIERC20UpgradeableFilterer creates a new log filterer instance of IERC20Upgradeable, bound to a specific deployed contract.
func NewIERC20UpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20UpgradeableFilterer, error) {
	contract, err := bindIERC20Upgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableFilterer{contract: contract}, nil
}

// bindIERC20Upgradeable binds a generic wrapper to an already deployed contract.
func bindIERC20Upgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20UpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Upgradeable *IERC20UpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Upgradeable.Contract.IERC20UpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Upgradeable *IERC20UpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.IERC20UpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Upgradeable *IERC20UpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.IERC20UpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Upgradeable *IERC20UpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Upgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Upgradeable *IERC20UpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Upgradeable *IERC20UpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Upgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Upgradeable.Contract.Allowance(&_IERC20Upgradeable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Upgradeable.Contract.Allowance(&_IERC20Upgradeable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Upgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Upgradeable.Contract.BalanceOf(&_IERC20Upgradeable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Upgradeable.Contract.BalanceOf(&_IERC20Upgradeable.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Upgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableSession) TotalSupply() (*big.Int, error) {
	return _IERC20Upgradeable.Contract.TotalSupply(&_IERC20Upgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Upgradeable *IERC20UpgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20Upgradeable.Contract.TotalSupply(&_IERC20Upgradeable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.Approve(&_IERC20Upgradeable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.Approve(&_IERC20Upgradeable.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.Transfer(&_IERC20Upgradeable.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.Transfer(&_IERC20Upgradeable.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.TransferFrom(&_IERC20Upgradeable.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Upgradeable *IERC20UpgradeableTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Upgradeable.Contract.TransferFrom(&_IERC20Upgradeable.TransactOpts, sender, recipient, amount)
}

// IERC20UpgradeableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Upgradeable contract.
type IERC20UpgradeableApprovalIterator struct {
	Event *IERC20UpgradeableApproval // Event containing the contract specifics and raw log

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
func (it *IERC20UpgradeableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20UpgradeableApproval)
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
		it.Event = new(IERC20UpgradeableApproval)
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
func (it *IERC20UpgradeableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20UpgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20UpgradeableApproval represents a Approval event raised by the IERC20Upgradeable contract.
type IERC20UpgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20UpgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Upgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableApprovalIterator{contract: _IERC20Upgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20UpgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Upgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20UpgradeableApproval)
				if err := _IERC20Upgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) ParseApproval(log types.Log) (*IERC20UpgradeableApproval, error) {
	event := new(IERC20UpgradeableApproval)
	if err := _IERC20Upgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20UpgradeableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Upgradeable contract.
type IERC20UpgradeableTransferIterator struct {
	Event *IERC20UpgradeableTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20UpgradeableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20UpgradeableTransfer)
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
		it.Event = new(IERC20UpgradeableTransfer)
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
func (it *IERC20UpgradeableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20UpgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20UpgradeableTransfer represents a Transfer event raised by the IERC20Upgradeable contract.
type IERC20UpgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20UpgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Upgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20UpgradeableTransferIterator{contract: _IERC20Upgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20UpgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Upgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20UpgradeableTransfer)
				if err := _IERC20Upgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Upgradeable *IERC20UpgradeableFilterer) ParseTransfer(log types.Log) (*IERC20UpgradeableTransfer, error) {
	event := new(IERC20UpgradeableTransfer)
	if err := _IERC20Upgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IMBMetaData contains all meta data concerning the IMB contract.
var IMBMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"9dc29fac": "burn(address,uint256)",
		"40c10f19": "mint(address,uint256)",
	},
}

// IMBABI is the input ABI used to generate the binding from.
// Deprecated: Use IMBMetaData.ABI instead.
var IMBABI = IMBMetaData.ABI

// Deprecated: Use IMBMetaData.Sigs instead.
// IMBFuncSigs maps the 4-byte function signature to its string representation.
var IMBFuncSigs = IMBMetaData.Sigs

// IMB is an auto generated Go binding around an Ethereum contract.
type IMB struct {
	IMBCaller     // Read-only binding to the contract
	IMBTransactor // Write-only binding to the contract
	IMBFilterer   // Log filterer for contract events
}

// IMBCaller is an auto generated read-only Go binding around an Ethereum contract.
type IMBCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMBTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IMBTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMBFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IMBFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IMBSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IMBSession struct {
	Contract     *IMB              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMBCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IMBCallerSession struct {
	Contract *IMBCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IMBTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IMBTransactorSession struct {
	Contract     *IMBTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IMBRaw is an auto generated low-level Go binding around an Ethereum contract.
type IMBRaw struct {
	Contract *IMB // Generic contract binding to access the raw methods on
}

// IMBCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IMBCallerRaw struct {
	Contract *IMBCaller // Generic read-only contract binding to access the raw methods on
}

// IMBTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IMBTransactorRaw struct {
	Contract *IMBTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIMB creates a new instance of IMB, bound to a specific deployed contract.
func NewIMB(address common.Address, backend bind.ContractBackend) (*IMB, error) {
	contract, err := bindIMB(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IMB{IMBCaller: IMBCaller{contract: contract}, IMBTransactor: IMBTransactor{contract: contract}, IMBFilterer: IMBFilterer{contract: contract}}, nil
}

// NewIMBCaller creates a new read-only instance of IMB, bound to a specific deployed contract.
func NewIMBCaller(address common.Address, caller bind.ContractCaller) (*IMBCaller, error) {
	contract, err := bindIMB(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IMBCaller{contract: contract}, nil
}

// NewIMBTransactor creates a new write-only instance of IMB, bound to a specific deployed contract.
func NewIMBTransactor(address common.Address, transactor bind.ContractTransactor) (*IMBTransactor, error) {
	contract, err := bindIMB(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IMBTransactor{contract: contract}, nil
}

// NewIMBFilterer creates a new log filterer instance of IMB, bound to a specific deployed contract.
func NewIMBFilterer(address common.Address, filterer bind.ContractFilterer) (*IMBFilterer, error) {
	contract, err := bindIMB(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IMBFilterer{contract: contract}, nil
}

// bindIMB binds a generic wrapper to an already deployed contract.
func bindIMB(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IMBABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMB *IMBRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMB.Contract.IMBCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMB *IMBRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMB.Contract.IMBTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMB *IMBRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMB.Contract.IMBTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IMB *IMBCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IMB.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IMB *IMBTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IMB.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IMB *IMBTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IMB.Contract.contract.Transact(opts, method, params...)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns(bool)
func (_IMB *IMBTransactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMB.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns(bool)
func (_IMB *IMBSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMB.Contract.Burn(&_IMB.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns(bool)
func (_IMB *IMBTransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMB.Contract.Burn(&_IMB.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_IMB *IMBTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMB.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_IMB *IMBSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMB.Contract.Mint(&_IMB.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_IMB *IMBTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IMB.Contract.Mint(&_IMB.TransactOpts, to, amount)
}

// IStakingRewardMetaData contains all meta data concerning the IStakingReward contract.
var IStakingRewardMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"earned\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"notifyRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"onChefReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"008cc262": "earned(address)",
		"246132f9": "notifyRewardAmount(uint256,uint256)",
		"04e6a5e2": "onChefReward(address)",
		"7b0a47ee": "rewardRate()",
	},
}

// IStakingRewardABI is the input ABI used to generate the binding from.
// Deprecated: Use IStakingRewardMetaData.ABI instead.
var IStakingRewardABI = IStakingRewardMetaData.ABI

// Deprecated: Use IStakingRewardMetaData.Sigs instead.
// IStakingRewardFuncSigs maps the 4-byte function signature to its string representation.
var IStakingRewardFuncSigs = IStakingRewardMetaData.Sigs

// IStakingReward is an auto generated Go binding around an Ethereum contract.
type IStakingReward struct {
	IStakingRewardCaller     // Read-only binding to the contract
	IStakingRewardTransactor // Write-only binding to the contract
	IStakingRewardFilterer   // Log filterer for contract events
}

// IStakingRewardCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStakingRewardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingRewardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStakingRewardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingRewardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStakingRewardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStakingRewardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStakingRewardSession struct {
	Contract     *IStakingReward   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStakingRewardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStakingRewardCallerSession struct {
	Contract *IStakingRewardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IStakingRewardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStakingRewardTransactorSession struct {
	Contract     *IStakingRewardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IStakingRewardRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStakingRewardRaw struct {
	Contract *IStakingReward // Generic contract binding to access the raw methods on
}

// IStakingRewardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStakingRewardCallerRaw struct {
	Contract *IStakingRewardCaller // Generic read-only contract binding to access the raw methods on
}

// IStakingRewardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStakingRewardTransactorRaw struct {
	Contract *IStakingRewardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStakingReward creates a new instance of IStakingReward, bound to a specific deployed contract.
func NewIStakingReward(address common.Address, backend bind.ContractBackend) (*IStakingReward, error) {
	contract, err := bindIStakingReward(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStakingReward{IStakingRewardCaller: IStakingRewardCaller{contract: contract}, IStakingRewardTransactor: IStakingRewardTransactor{contract: contract}, IStakingRewardFilterer: IStakingRewardFilterer{contract: contract}}, nil
}

// NewIStakingRewardCaller creates a new read-only instance of IStakingReward, bound to a specific deployed contract.
func NewIStakingRewardCaller(address common.Address, caller bind.ContractCaller) (*IStakingRewardCaller, error) {
	contract, err := bindIStakingReward(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingRewardCaller{contract: contract}, nil
}

// NewIStakingRewardTransactor creates a new write-only instance of IStakingReward, bound to a specific deployed contract.
func NewIStakingRewardTransactor(address common.Address, transactor bind.ContractTransactor) (*IStakingRewardTransactor, error) {
	contract, err := bindIStakingReward(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStakingRewardTransactor{contract: contract}, nil
}

// NewIStakingRewardFilterer creates a new log filterer instance of IStakingReward, bound to a specific deployed contract.
func NewIStakingRewardFilterer(address common.Address, filterer bind.ContractFilterer) (*IStakingRewardFilterer, error) {
	contract, err := bindIStakingReward(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStakingRewardFilterer{contract: contract}, nil
}

// bindIStakingReward binds a generic wrapper to an already deployed contract.
func bindIStakingReward(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStakingRewardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingReward *IStakingRewardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingReward.Contract.IStakingRewardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingReward *IStakingRewardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingReward.Contract.IStakingRewardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingReward *IStakingRewardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingReward.Contract.IStakingRewardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStakingReward *IStakingRewardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStakingReward.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStakingReward *IStakingRewardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStakingReward.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStakingReward *IStakingRewardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStakingReward.Contract.contract.Transact(opts, method, params...)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_IStakingReward *IStakingRewardCaller) Earned(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStakingReward.contract.Call(opts, &out, "earned", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_IStakingReward *IStakingRewardSession) Earned(account common.Address) (*big.Int, error) {
	return _IStakingReward.Contract.Earned(&_IStakingReward.CallOpts, account)
}

// Earned is a free data retrieval call binding the contract method 0x008cc262.
//
// Solidity: function earned(address account) view returns(uint256)
func (_IStakingReward *IStakingRewardCallerSession) Earned(account common.Address) (*big.Int, error) {
	return _IStakingReward.Contract.Earned(&_IStakingReward.CallOpts, account)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_IStakingReward *IStakingRewardCaller) RewardRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStakingReward.contract.Call(opts, &out, "rewardRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_IStakingReward *IStakingRewardSession) RewardRate() (*big.Int, error) {
	return _IStakingReward.Contract.RewardRate(&_IStakingReward.CallOpts)
}

// RewardRate is a free data retrieval call binding the contract method 0x7b0a47ee.
//
// Solidity: function rewardRate() view returns(uint256)
func (_IStakingReward *IStakingRewardCallerSession) RewardRate() (*big.Int, error) {
	return _IStakingReward.Contract.RewardRate(&_IStakingReward.CallOpts)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x246132f9.
//
// Solidity: function notifyRewardAmount(uint256 reward, uint256 duration) returns()
func (_IStakingReward *IStakingRewardTransactor) NotifyRewardAmount(opts *bind.TransactOpts, reward *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _IStakingReward.contract.Transact(opts, "notifyRewardAmount", reward, duration)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x246132f9.
//
// Solidity: function notifyRewardAmount(uint256 reward, uint256 duration) returns()
func (_IStakingReward *IStakingRewardSession) NotifyRewardAmount(reward *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _IStakingReward.Contract.NotifyRewardAmount(&_IStakingReward.TransactOpts, reward, duration)
}

// NotifyRewardAmount is a paid mutator transaction binding the contract method 0x246132f9.
//
// Solidity: function notifyRewardAmount(uint256 reward, uint256 duration) returns()
func (_IStakingReward *IStakingRewardTransactorSession) NotifyRewardAmount(reward *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _IStakingReward.Contract.NotifyRewardAmount(&_IStakingReward.TransactOpts, reward, duration)
}

// OnChefReward is a paid mutator transaction binding the contract method 0x04e6a5e2.
//
// Solidity: function onChefReward(address user) returns()
func (_IStakingReward *IStakingRewardTransactor) OnChefReward(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _IStakingReward.contract.Transact(opts, "onChefReward", user)
}

// OnChefReward is a paid mutator transaction binding the contract method 0x04e6a5e2.
//
// Solidity: function onChefReward(address user) returns()
func (_IStakingReward *IStakingRewardSession) OnChefReward(user common.Address) (*types.Transaction, error) {
	return _IStakingReward.Contract.OnChefReward(&_IStakingReward.TransactOpts, user)
}

// OnChefReward is a paid mutator transaction binding the contract method 0x04e6a5e2.
//
// Solidity: function onChefReward(address user) returns()
func (_IStakingReward *IStakingRewardTransactorSession) OnChefReward(user common.Address) (*types.Transaction, error) {
	return _IStakingReward.Contract.OnChefReward(&_IStakingReward.TransactOpts, user)
}

// ITokenMetaData contains all meta data concerning the IToken contract.
var ITokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"9dc29fac": "burn(address,uint256)",
		"40c10f19": "mint(address,uint256)",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
	},
}

// ITokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ITokenMetaData.ABI instead.
var ITokenABI = ITokenMetaData.ABI

// Deprecated: Use ITokenMetaData.Sigs instead.
// ITokenFuncSigs maps the 4-byte function signature to its string representation.
var ITokenFuncSigs = ITokenMetaData.Sigs

// IToken is an auto generated Go binding around an Ethereum contract.
type IToken struct {
	ITokenCaller     // Read-only binding to the contract
	ITokenTransactor // Write-only binding to the contract
	ITokenFilterer   // Log filterer for contract events
}

// ITokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITokenSession struct {
	Contract     *IToken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITokenCallerSession struct {
	Contract *ITokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ITokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITokenTransactorSession struct {
	Contract     *ITokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITokenRaw struct {
	Contract *IToken // Generic contract binding to access the raw methods on
}

// ITokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITokenCallerRaw struct {
	Contract *ITokenCaller // Generic read-only contract binding to access the raw methods on
}

// ITokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITokenTransactorRaw struct {
	Contract *ITokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIToken creates a new instance of IToken, bound to a specific deployed contract.
func NewIToken(address common.Address, backend bind.ContractBackend) (*IToken, error) {
	contract, err := bindIToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IToken{ITokenCaller: ITokenCaller{contract: contract}, ITokenTransactor: ITokenTransactor{contract: contract}, ITokenFilterer: ITokenFilterer{contract: contract}}, nil
}

// NewITokenCaller creates a new read-only instance of IToken, bound to a specific deployed contract.
func NewITokenCaller(address common.Address, caller bind.ContractCaller) (*ITokenCaller, error) {
	contract, err := bindIToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenCaller{contract: contract}, nil
}

// NewITokenTransactor creates a new write-only instance of IToken, bound to a specific deployed contract.
func NewITokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ITokenTransactor, error) {
	contract, err := bindIToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITokenTransactor{contract: contract}, nil
}

// NewITokenFilterer creates a new log filterer instance of IToken, bound to a specific deployed contract.
func NewITokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ITokenFilterer, error) {
	contract, err := bindIToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITokenFilterer{contract: contract}, nil
}

// bindIToken binds a generic wrapper to an already deployed contract.
func bindIToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IToken *ITokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IToken.Contract.ITokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IToken *ITokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.Contract.ITokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IToken *ITokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IToken.Contract.ITokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IToken *ITokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IToken *ITokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IToken *ITokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IToken *ITokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IToken *ITokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IToken.Contract.Allowance(&_IToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IToken *ITokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IToken.Contract.Allowance(&_IToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IToken *ITokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IToken *ITokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IToken.Contract.BalanceOf(&_IToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IToken *ITokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IToken.Contract.BalanceOf(&_IToken.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IToken *ITokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IToken *ITokenSession) TotalSupply() (*big.Int, error) {
	return _IToken.Contract.TotalSupply(&_IToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IToken *ITokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IToken.Contract.TotalSupply(&_IToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IToken *ITokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IToken *ITokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Approve(&_IToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IToken *ITokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Approve(&_IToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns(bool)
func (_IToken *ITokenTransactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns(bool)
func (_IToken *ITokenSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Burn(&_IToken.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns(bool)
func (_IToken *ITokenTransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Burn(&_IToken.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_IToken *ITokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_IToken *ITokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Mint(&_IToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns(bool)
func (_IToken *ITokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Mint(&_IToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IToken *ITokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IToken *ITokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Transfer(&_IToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IToken *ITokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.Transfer(&_IToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IToken *ITokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IToken *ITokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.TransferFrom(&_IToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IToken *ITokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IToken.Contract.TransferFrom(&_IToken.TransactOpts, sender, recipient, amount)
}

// ITokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IToken contract.
type ITokenApprovalIterator struct {
	Event *ITokenApproval // Event containing the contract specifics and raw log

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
func (it *ITokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenApproval)
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
		it.Event = new(ITokenApproval)
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
func (it *ITokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenApproval represents a Approval event raised by the IToken contract.
type ITokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IToken *ITokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ITokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ITokenApprovalIterator{contract: _IToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IToken *ITokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ITokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenApproval)
				if err := _IToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IToken *ITokenFilterer) ParseApproval(log types.Log) (*ITokenApproval, error) {
	event := new(ITokenApproval)
	if err := _IToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ITokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IToken contract.
type ITokenTransferIterator struct {
	Event *ITokenTransfer // Event containing the contract specifics and raw log

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
func (it *ITokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ITokenTransfer)
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
		it.Event = new(ITokenTransfer)
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
func (it *ITokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ITokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ITokenTransfer represents a Transfer event raised by the IToken contract.
type ITokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IToken *ITokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ITokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ITokenTransferIterator{contract: _IToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IToken *ITokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ITokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ITokenTransfer)
				if err := _IToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IToken *ITokenFilterer) ParseTransfer(log types.Log) (*ITokenTransfer, error) {
	event := new(ITokenTransfer)
	if err := _IToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWETH9MetaData contains all meta data concerning the IWETH9 contract.
var IWETH9MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"dd62ed3e": "allowance(address,address)",
		"095ea7b3": "approve(address,uint256)",
		"70a08231": "balanceOf(address)",
		"d0e30db0": "deposit()",
		"18160ddd": "totalSupply()",
		"a9059cbb": "transfer(address,uint256)",
		"23b872dd": "transferFrom(address,address,uint256)",
		"2e1a7d4d": "withdraw(uint256)",
	},
}

// IWETH9ABI is the input ABI used to generate the binding from.
// Deprecated: Use IWETH9MetaData.ABI instead.
var IWETH9ABI = IWETH9MetaData.ABI

// Deprecated: Use IWETH9MetaData.Sigs instead.
// IWETH9FuncSigs maps the 4-byte function signature to its string representation.
var IWETH9FuncSigs = IWETH9MetaData.Sigs

// IWETH9 is an auto generated Go binding around an Ethereum contract.
type IWETH9 struct {
	IWETH9Caller     // Read-only binding to the contract
	IWETH9Transactor // Write-only binding to the contract
	IWETH9Filterer   // Log filterer for contract events
}

// IWETH9Caller is an auto generated read-only Go binding around an Ethereum contract.
type IWETH9Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IWETH9Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWETH9Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWETH9Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWETH9Session struct {
	Contract     *IWETH9           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETH9CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWETH9CallerSession struct {
	Contract *IWETH9Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IWETH9TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWETH9TransactorSession struct {
	Contract     *IWETH9Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWETH9Raw is an auto generated low-level Go binding around an Ethereum contract.
type IWETH9Raw struct {
	Contract *IWETH9 // Generic contract binding to access the raw methods on
}

// IWETH9CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWETH9CallerRaw struct {
	Contract *IWETH9Caller // Generic read-only contract binding to access the raw methods on
}

// IWETH9TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWETH9TransactorRaw struct {
	Contract *IWETH9Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIWETH9 creates a new instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9(address common.Address, backend bind.ContractBackend) (*IWETH9, error) {
	contract, err := bindIWETH9(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWETH9{IWETH9Caller: IWETH9Caller{contract: contract}, IWETH9Transactor: IWETH9Transactor{contract: contract}, IWETH9Filterer: IWETH9Filterer{contract: contract}}, nil
}

// NewIWETH9Caller creates a new read-only instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Caller(address common.Address, caller bind.ContractCaller) (*IWETH9Caller, error) {
	contract, err := bindIWETH9(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH9Caller{contract: contract}, nil
}

// NewIWETH9Transactor creates a new write-only instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Transactor(address common.Address, transactor bind.ContractTransactor) (*IWETH9Transactor, error) {
	contract, err := bindIWETH9(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWETH9Transactor{contract: contract}, nil
}

// NewIWETH9Filterer creates a new log filterer instance of IWETH9, bound to a specific deployed contract.
func NewIWETH9Filterer(address common.Address, filterer bind.ContractFilterer) (*IWETH9Filterer, error) {
	contract, err := bindIWETH9(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWETH9Filterer{contract: contract}, nil
}

// bindIWETH9 binds a generic wrapper to an already deployed contract.
func bindIWETH9(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWETH9ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH9 *IWETH9Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH9.Contract.IWETH9Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH9 *IWETH9Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.Contract.IWETH9Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH9 *IWETH9Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH9.Contract.IWETH9Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWETH9 *IWETH9CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IWETH9.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWETH9 *IWETH9TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWETH9 *IWETH9TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWETH9.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH9 *IWETH9Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH9 *IWETH9Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWETH9.Contract.Allowance(&_IWETH9.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IWETH9 *IWETH9CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IWETH9.Contract.Allowance(&_IWETH9.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH9 *IWETH9Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH9 *IWETH9Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IWETH9.Contract.BalanceOf(&_IWETH9.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IWETH9 *IWETH9CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IWETH9.Contract.BalanceOf(&_IWETH9.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH9 *IWETH9Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IWETH9.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH9 *IWETH9Session) TotalSupply() (*big.Int, error) {
	return _IWETH9.Contract.TotalSupply(&_IWETH9.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IWETH9 *IWETH9CallerSession) TotalSupply() (*big.Int, error) {
	return _IWETH9.Contract.TotalSupply(&_IWETH9.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Approve(&_IWETH9.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Approve(&_IWETH9.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9Transactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9Session) Deposit() (*types.Transaction, error) {
	return _IWETH9.Contract.Deposit(&_IWETH9.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_IWETH9 *IWETH9TransactorSession) Deposit() (*types.Transaction, error) {
	return _IWETH9.Contract.Deposit(&_IWETH9.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Transfer(&_IWETH9.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Transfer(&_IWETH9.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.TransferFrom(&_IWETH9.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IWETH9 *IWETH9TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.TransferFrom(&_IWETH9.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH9 *IWETH9Transactor) Withdraw(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH9.contract.Transact(opts, "withdraw", arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH9 *IWETH9Session) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Withdraw(&_IWETH9.TransactOpts, arg0)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 ) returns()
func (_IWETH9 *IWETH9TransactorSession) Withdraw(arg0 *big.Int) (*types.Transaction, error) {
	return _IWETH9.Contract.Withdraw(&_IWETH9.TransactOpts, arg0)
}

// IWETH9ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IWETH9 contract.
type IWETH9ApprovalIterator struct {
	Event *IWETH9Approval // Event containing the contract specifics and raw log

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
func (it *IWETH9ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWETH9Approval)
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
		it.Event = new(IWETH9Approval)
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
func (it *IWETH9ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWETH9ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWETH9Approval represents a Approval event raised by the IWETH9 contract.
type IWETH9Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH9 *IWETH9Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IWETH9ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWETH9.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IWETH9ApprovalIterator{contract: _IWETH9.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH9 *IWETH9Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IWETH9Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IWETH9.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWETH9Approval)
				if err := _IWETH9.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IWETH9 *IWETH9Filterer) ParseApproval(log types.Log) (*IWETH9Approval, error) {
	event := new(IWETH9Approval)
	if err := _IWETH9.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IWETH9TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IWETH9 contract.
type IWETH9TransferIterator struct {
	Event *IWETH9Transfer // Event containing the contract specifics and raw log

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
func (it *IWETH9TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IWETH9Transfer)
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
		it.Event = new(IWETH9Transfer)
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
func (it *IWETH9TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IWETH9TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IWETH9Transfer represents a Transfer event raised by the IWETH9 contract.
type IWETH9Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH9 *IWETH9Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IWETH9TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWETH9.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IWETH9TransferIterator{contract: _IWETH9.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH9 *IWETH9Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IWETH9Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IWETH9.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IWETH9Transfer)
				if err := _IWETH9.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IWETH9 *IWETH9Filterer) ParseTransfer(log types.Log) (*IWETH9Transfer, error) {
	event := new(IWETH9Transfer)
	if err := _IWETH9.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// InitializableMetaData contains all meta data concerning the Initializable contract.
var InitializableMetaData = &bind.MetaData{
	ABI: "[]",
}

// InitializableABI is the input ABI used to generate the binding from.
// Deprecated: Use InitializableMetaData.ABI instead.
var InitializableABI = InitializableMetaData.ABI

// Initializable is an auto generated Go binding around an Ethereum contract.
type Initializable struct {
	InitializableCaller     // Read-only binding to the contract
	InitializableTransactor // Write-only binding to the contract
	InitializableFilterer   // Log filterer for contract events
}

// InitializableCaller is an auto generated read-only Go binding around an Ethereum contract.
type InitializableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InitializableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InitializableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InitializableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InitializableSession struct {
	Contract     *Initializable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// InitializableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InitializableCallerSession struct {
	Contract *InitializableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// InitializableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InitializableTransactorSession struct {
	Contract     *InitializableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// InitializableRaw is an auto generated low-level Go binding around an Ethereum contract.
type InitializableRaw struct {
	Contract *Initializable // Generic contract binding to access the raw methods on
}

// InitializableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InitializableCallerRaw struct {
	Contract *InitializableCaller // Generic read-only contract binding to access the raw methods on
}

// InitializableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InitializableTransactorRaw struct {
	Contract *InitializableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInitializable creates a new instance of Initializable, bound to a specific deployed contract.
func NewInitializable(address common.Address, backend bind.ContractBackend) (*Initializable, error) {
	contract, err := bindInitializable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Initializable{InitializableCaller: InitializableCaller{contract: contract}, InitializableTransactor: InitializableTransactor{contract: contract}, InitializableFilterer: InitializableFilterer{contract: contract}}, nil
}

// NewInitializableCaller creates a new read-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableCaller(address common.Address, caller bind.ContractCaller) (*InitializableCaller, error) {
	contract, err := bindInitializable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableCaller{contract: contract}, nil
}

// NewInitializableTransactor creates a new write-only instance of Initializable, bound to a specific deployed contract.
func NewInitializableTransactor(address common.Address, transactor bind.ContractTransactor) (*InitializableTransactor, error) {
	contract, err := bindInitializable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InitializableTransactor{contract: contract}, nil
}

// NewInitializableFilterer creates a new log filterer instance of Initializable, bound to a specific deployed contract.
func NewInitializableFilterer(address common.Address, filterer bind.ContractFilterer) (*InitializableFilterer, error) {
	contract, err := bindInitializable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InitializableFilterer{contract: contract}, nil
}

// bindInitializable binds a generic wrapper to an already deployed contract.
func bindInitializable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InitializableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.InitializableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.InitializableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Initializable *InitializableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Initializable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Initializable *InitializableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Initializable *InitializableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Initializable.Contract.contract.Transact(opts, method, params...)
}

// SafeDecimalMathMetaData contains all meta data concerning the SafeDecimalMath contract.
var SafeDecimalMathMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"PRECISE_UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highPrecisionDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preciseUnit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"864029e7": "PRECISE_UNIT()",
		"9d8e2177": "UNIT()",
		"313ce567": "decimals()",
		"def4419d": "highPrecisionDecimals()",
		"d5e5e6e6": "preciseUnit()",
		"907af6c0": "unit()",
	},
	Bin: "0x61026e61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061006c5760003560e01c8063313ce56714610071578063864029e714610090578063907af6c0146100a65780639d8e2177146100ae578063d5e5e6e6146100b6578063def4419d146100be575b600080fd5b610079601281565b60405160ff90911681526020015b60405180910390f35b6100986100c6565b604051908152602001610087565b6100986100d5565b6100986100e8565b6100986100f4565b610079601b81565b6100d2601b600a610148565b81565b60006100e36012600a610148565b905090565b6100d26012600a610148565b60006100e3601b600a610148565b80825b6001808611610114575061013f565b81870482111561012657610126610222565b8086161561013357918102915b9490941c938002610105565b94509492505050565b6000610157600019848461015e565b9392505050565b60008261016d57506001610157565b8161017a57506000610157565b8160018114610190576002811461019a576101c7565b6001915050610157565b60ff8411156101ab576101ab610222565b6001841b9150848211156101c1576101c1610222565b50610157565b5060208310610133831016604e8410600b84101617156101fa575081810a838111156101f5576101f5610222565b610157565b6102078484846001610102565b80860482111561021957610219610222565b02949350505050565b634e487b7160e01b600052601160045260246000fdfea26469706673582212206ca30230cefedbdee00b77ff200aa631a2a4b260a79d825abefeefb02760eba264736f6c63430008020033",
}

// SafeDecimalMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeDecimalMathMetaData.ABI instead.
var SafeDecimalMathABI = SafeDecimalMathMetaData.ABI

// Deprecated: Use SafeDecimalMathMetaData.Sigs instead.
// SafeDecimalMathFuncSigs maps the 4-byte function signature to its string representation.
var SafeDecimalMathFuncSigs = SafeDecimalMathMetaData.Sigs

// SafeDecimalMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeDecimalMathMetaData.Bin instead.
var SafeDecimalMathBin = SafeDecimalMathMetaData.Bin

// DeploySafeDecimalMath deploys a new Ethereum contract, binding an instance of SafeDecimalMath to it.
func DeploySafeDecimalMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeDecimalMath, error) {
	parsed, err := SafeDecimalMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeDecimalMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeDecimalMath{SafeDecimalMathCaller: SafeDecimalMathCaller{contract: contract}, SafeDecimalMathTransactor: SafeDecimalMathTransactor{contract: contract}, SafeDecimalMathFilterer: SafeDecimalMathFilterer{contract: contract}}, nil
}

// SafeDecimalMath is an auto generated Go binding around an Ethereum contract.
type SafeDecimalMath struct {
	SafeDecimalMathCaller     // Read-only binding to the contract
	SafeDecimalMathTransactor // Write-only binding to the contract
	SafeDecimalMathFilterer   // Log filterer for contract events
}

// SafeDecimalMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeDecimalMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeDecimalMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeDecimalMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeDecimalMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeDecimalMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeDecimalMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeDecimalMathSession struct {
	Contract     *SafeDecimalMath  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeDecimalMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeDecimalMathCallerSession struct {
	Contract *SafeDecimalMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SafeDecimalMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeDecimalMathTransactorSession struct {
	Contract     *SafeDecimalMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SafeDecimalMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeDecimalMathRaw struct {
	Contract *SafeDecimalMath // Generic contract binding to access the raw methods on
}

// SafeDecimalMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeDecimalMathCallerRaw struct {
	Contract *SafeDecimalMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeDecimalMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeDecimalMathTransactorRaw struct {
	Contract *SafeDecimalMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeDecimalMath creates a new instance of SafeDecimalMath, bound to a specific deployed contract.
func NewSafeDecimalMath(address common.Address, backend bind.ContractBackend) (*SafeDecimalMath, error) {
	contract, err := bindSafeDecimalMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeDecimalMath{SafeDecimalMathCaller: SafeDecimalMathCaller{contract: contract}, SafeDecimalMathTransactor: SafeDecimalMathTransactor{contract: contract}, SafeDecimalMathFilterer: SafeDecimalMathFilterer{contract: contract}}, nil
}

// NewSafeDecimalMathCaller creates a new read-only instance of SafeDecimalMath, bound to a specific deployed contract.
func NewSafeDecimalMathCaller(address common.Address, caller bind.ContractCaller) (*SafeDecimalMathCaller, error) {
	contract, err := bindSafeDecimalMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeDecimalMathCaller{contract: contract}, nil
}

// NewSafeDecimalMathTransactor creates a new write-only instance of SafeDecimalMath, bound to a specific deployed contract.
func NewSafeDecimalMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeDecimalMathTransactor, error) {
	contract, err := bindSafeDecimalMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeDecimalMathTransactor{contract: contract}, nil
}

// NewSafeDecimalMathFilterer creates a new log filterer instance of SafeDecimalMath, bound to a specific deployed contract.
func NewSafeDecimalMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeDecimalMathFilterer, error) {
	contract, err := bindSafeDecimalMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeDecimalMathFilterer{contract: contract}, nil
}

// bindSafeDecimalMath binds a generic wrapper to an already deployed contract.
func bindSafeDecimalMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeDecimalMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeDecimalMath *SafeDecimalMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeDecimalMath.Contract.SafeDecimalMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeDecimalMath *SafeDecimalMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeDecimalMath.Contract.SafeDecimalMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeDecimalMath *SafeDecimalMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeDecimalMath.Contract.SafeDecimalMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeDecimalMath *SafeDecimalMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeDecimalMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeDecimalMath *SafeDecimalMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeDecimalMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeDecimalMath *SafeDecimalMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeDecimalMath.Contract.contract.Transact(opts, method, params...)
}

// PRECISEUNIT is a free data retrieval call binding the contract method 0x864029e7.
//
// Solidity: function PRECISE_UNIT() view returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCaller) PRECISEUNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeDecimalMath.contract.Call(opts, &out, "PRECISE_UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRECISEUNIT is a free data retrieval call binding the contract method 0x864029e7.
//
// Solidity: function PRECISE_UNIT() view returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathSession) PRECISEUNIT() (*big.Int, error) {
	return _SafeDecimalMath.Contract.PRECISEUNIT(&_SafeDecimalMath.CallOpts)
}

// PRECISEUNIT is a free data retrieval call binding the contract method 0x864029e7.
//
// Solidity: function PRECISE_UNIT() view returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCallerSession) PRECISEUNIT() (*big.Int, error) {
	return _SafeDecimalMath.Contract.PRECISEUNIT(&_SafeDecimalMath.CallOpts)
}

// UNIT is a free data retrieval call binding the contract method 0x9d8e2177.
//
// Solidity: function UNIT() view returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCaller) UNIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeDecimalMath.contract.Call(opts, &out, "UNIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UNIT is a free data retrieval call binding the contract method 0x9d8e2177.
//
// Solidity: function UNIT() view returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathSession) UNIT() (*big.Int, error) {
	return _SafeDecimalMath.Contract.UNIT(&_SafeDecimalMath.CallOpts)
}

// UNIT is a free data retrieval call binding the contract method 0x9d8e2177.
//
// Solidity: function UNIT() view returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCallerSession) UNIT() (*big.Int, error) {
	return _SafeDecimalMath.Contract.UNIT(&_SafeDecimalMath.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SafeDecimalMath *SafeDecimalMathCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SafeDecimalMath.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SafeDecimalMath *SafeDecimalMathSession) Decimals() (uint8, error) {
	return _SafeDecimalMath.Contract.Decimals(&_SafeDecimalMath.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SafeDecimalMath *SafeDecimalMathCallerSession) Decimals() (uint8, error) {
	return _SafeDecimalMath.Contract.Decimals(&_SafeDecimalMath.CallOpts)
}

// HighPrecisionDecimals is a free data retrieval call binding the contract method 0xdef4419d.
//
// Solidity: function highPrecisionDecimals() view returns(uint8)
func (_SafeDecimalMath *SafeDecimalMathCaller) HighPrecisionDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SafeDecimalMath.contract.Call(opts, &out, "highPrecisionDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// HighPrecisionDecimals is a free data retrieval call binding the contract method 0xdef4419d.
//
// Solidity: function highPrecisionDecimals() view returns(uint8)
func (_SafeDecimalMath *SafeDecimalMathSession) HighPrecisionDecimals() (uint8, error) {
	return _SafeDecimalMath.Contract.HighPrecisionDecimals(&_SafeDecimalMath.CallOpts)
}

// HighPrecisionDecimals is a free data retrieval call binding the contract method 0xdef4419d.
//
// Solidity: function highPrecisionDecimals() view returns(uint8)
func (_SafeDecimalMath *SafeDecimalMathCallerSession) HighPrecisionDecimals() (uint8, error) {
	return _SafeDecimalMath.Contract.HighPrecisionDecimals(&_SafeDecimalMath.CallOpts)
}

// PreciseUnit is a free data retrieval call binding the contract method 0xd5e5e6e6.
//
// Solidity: function preciseUnit() pure returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCaller) PreciseUnit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeDecimalMath.contract.Call(opts, &out, "preciseUnit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreciseUnit is a free data retrieval call binding the contract method 0xd5e5e6e6.
//
// Solidity: function preciseUnit() pure returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathSession) PreciseUnit() (*big.Int, error) {
	return _SafeDecimalMath.Contract.PreciseUnit(&_SafeDecimalMath.CallOpts)
}

// PreciseUnit is a free data retrieval call binding the contract method 0xd5e5e6e6.
//
// Solidity: function preciseUnit() pure returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCallerSession) PreciseUnit() (*big.Int, error) {
	return _SafeDecimalMath.Contract.PreciseUnit(&_SafeDecimalMath.CallOpts)
}

// Unit is a free data retrieval call binding the contract method 0x907af6c0.
//
// Solidity: function unit() pure returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCaller) Unit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SafeDecimalMath.contract.Call(opts, &out, "unit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Unit is a free data retrieval call binding the contract method 0x907af6c0.
//
// Solidity: function unit() pure returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathSession) Unit() (*big.Int, error) {
	return _SafeDecimalMath.Contract.Unit(&_SafeDecimalMath.CallOpts)
}

// Unit is a free data retrieval call binding the contract method 0x907af6c0.
//
// Solidity: function unit() pure returns(uint256)
func (_SafeDecimalMath *SafeDecimalMathCallerSession) Unit() (*big.Int, error) {
	return _SafeDecimalMath.Contract.Unit(&_SafeDecimalMath.CallOpts)
}

// SafeERC20UpgradeableMetaData contains all meta data concerning the SafeERC20Upgradeable contract.
var SafeERC20UpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212208a4ce7a9428c0bc598f6c6184fe62521d566bdd416b562326ebc7fe261d559b864736f6c63430008020033",
}

// SafeERC20UpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeERC20UpgradeableMetaData.ABI instead.
var SafeERC20UpgradeableABI = SafeERC20UpgradeableMetaData.ABI

// SafeERC20UpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeERC20UpgradeableMetaData.Bin instead.
var SafeERC20UpgradeableBin = SafeERC20UpgradeableMetaData.Bin

// DeploySafeERC20Upgradeable deploys a new Ethereum contract, binding an instance of SafeERC20Upgradeable to it.
func DeploySafeERC20Upgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20Upgradeable, error) {
	parsed, err := SafeERC20UpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeERC20UpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20Upgradeable{SafeERC20UpgradeableCaller: SafeERC20UpgradeableCaller{contract: contract}, SafeERC20UpgradeableTransactor: SafeERC20UpgradeableTransactor{contract: contract}, SafeERC20UpgradeableFilterer: SafeERC20UpgradeableFilterer{contract: contract}}, nil
}

// SafeERC20Upgradeable is an auto generated Go binding around an Ethereum contract.
type SafeERC20Upgradeable struct {
	SafeERC20UpgradeableCaller     // Read-only binding to the contract
	SafeERC20UpgradeableTransactor // Write-only binding to the contract
	SafeERC20UpgradeableFilterer   // Log filterer for contract events
}

// SafeERC20UpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20UpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20UpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20UpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20UpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20UpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20UpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20UpgradeableSession struct {
	Contract     *SafeERC20Upgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SafeERC20UpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20UpgradeableCallerSession struct {
	Contract *SafeERC20UpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// SafeERC20UpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20UpgradeableTransactorSession struct {
	Contract     *SafeERC20UpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// SafeERC20UpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20UpgradeableRaw struct {
	Contract *SafeERC20Upgradeable // Generic contract binding to access the raw methods on
}

// SafeERC20UpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20UpgradeableCallerRaw struct {
	Contract *SafeERC20UpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20UpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20UpgradeableTransactorRaw struct {
	Contract *SafeERC20UpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20Upgradeable creates a new instance of SafeERC20Upgradeable, bound to a specific deployed contract.
func NewSafeERC20Upgradeable(address common.Address, backend bind.ContractBackend) (*SafeERC20Upgradeable, error) {
	contract, err := bindSafeERC20Upgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Upgradeable{SafeERC20UpgradeableCaller: SafeERC20UpgradeableCaller{contract: contract}, SafeERC20UpgradeableTransactor: SafeERC20UpgradeableTransactor{contract: contract}, SafeERC20UpgradeableFilterer: SafeERC20UpgradeableFilterer{contract: contract}}, nil
}

// NewSafeERC20UpgradeableCaller creates a new read-only instance of SafeERC20Upgradeable, bound to a specific deployed contract.
func NewSafeERC20UpgradeableCaller(address common.Address, caller bind.ContractCaller) (*SafeERC20UpgradeableCaller, error) {
	contract, err := bindSafeERC20Upgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20UpgradeableCaller{contract: contract}, nil
}

// NewSafeERC20UpgradeableTransactor creates a new write-only instance of SafeERC20Upgradeable, bound to a specific deployed contract.
func NewSafeERC20UpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20UpgradeableTransactor, error) {
	contract, err := bindSafeERC20Upgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20UpgradeableTransactor{contract: contract}, nil
}

// NewSafeERC20UpgradeableFilterer creates a new log filterer instance of SafeERC20Upgradeable, bound to a specific deployed contract.
func NewSafeERC20UpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20UpgradeableFilterer, error) {
	contract, err := bindSafeERC20Upgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20UpgradeableFilterer{contract: contract}, nil
}

// bindSafeERC20Upgradeable binds a generic wrapper to an already deployed contract.
func bindSafeERC20Upgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20UpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20Upgradeable *SafeERC20UpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20Upgradeable.Contract.SafeERC20UpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20Upgradeable *SafeERC20UpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20Upgradeable.Contract.SafeERC20UpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20Upgradeable *SafeERC20UpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20Upgradeable.Contract.SafeERC20UpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20Upgradeable *SafeERC20UpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20Upgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20Upgradeable *SafeERC20UpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20Upgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20Upgradeable *SafeERC20UpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20Upgradeable.Contract.contract.Transact(opts, method, params...)
}

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202e316e3f7d265bf017330cf9df33314eef901547b120602bfd254dd34e285d2e64736f6c63430008020033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// StorageSlotUpgradeableMetaData contains all meta data concerning the StorageSlotUpgradeable contract.
var StorageSlotUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212205c9c4a0d0219bc2b8e52caf3528191596b60fa300e1baa15297a01ab925a034664736f6c63430008020033",
}

// StorageSlotUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use StorageSlotUpgradeableMetaData.ABI instead.
var StorageSlotUpgradeableABI = StorageSlotUpgradeableMetaData.ABI

// StorageSlotUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StorageSlotUpgradeableMetaData.Bin instead.
var StorageSlotUpgradeableBin = StorageSlotUpgradeableMetaData.Bin

// DeployStorageSlotUpgradeable deploys a new Ethereum contract, binding an instance of StorageSlotUpgradeable to it.
func DeployStorageSlotUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StorageSlotUpgradeable, error) {
	parsed, err := StorageSlotUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StorageSlotUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StorageSlotUpgradeable{StorageSlotUpgradeableCaller: StorageSlotUpgradeableCaller{contract: contract}, StorageSlotUpgradeableTransactor: StorageSlotUpgradeableTransactor{contract: contract}, StorageSlotUpgradeableFilterer: StorageSlotUpgradeableFilterer{contract: contract}}, nil
}

// StorageSlotUpgradeable is an auto generated Go binding around an Ethereum contract.
type StorageSlotUpgradeable struct {
	StorageSlotUpgradeableCaller     // Read-only binding to the contract
	StorageSlotUpgradeableTransactor // Write-only binding to the contract
	StorageSlotUpgradeableFilterer   // Log filterer for contract events
}

// StorageSlotUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type StorageSlotUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StorageSlotUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StorageSlotUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StorageSlotUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StorageSlotUpgradeableSession struct {
	Contract     *StorageSlotUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StorageSlotUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StorageSlotUpgradeableCallerSession struct {
	Contract *StorageSlotUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// StorageSlotUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StorageSlotUpgradeableTransactorSession struct {
	Contract     *StorageSlotUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// StorageSlotUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type StorageSlotUpgradeableRaw struct {
	Contract *StorageSlotUpgradeable // Generic contract binding to access the raw methods on
}

// StorageSlotUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StorageSlotUpgradeableCallerRaw struct {
	Contract *StorageSlotUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// StorageSlotUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StorageSlotUpgradeableTransactorRaw struct {
	Contract *StorageSlotUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStorageSlotUpgradeable creates a new instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeable(address common.Address, backend bind.ContractBackend) (*StorageSlotUpgradeable, error) {
	contract, err := bindStorageSlotUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeable{StorageSlotUpgradeableCaller: StorageSlotUpgradeableCaller{contract: contract}, StorageSlotUpgradeableTransactor: StorageSlotUpgradeableTransactor{contract: contract}, StorageSlotUpgradeableFilterer: StorageSlotUpgradeableFilterer{contract: contract}}, nil
}

// NewStorageSlotUpgradeableCaller creates a new read-only instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*StorageSlotUpgradeableCaller, error) {
	contract, err := bindStorageSlotUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeableCaller{contract: contract}, nil
}

// NewStorageSlotUpgradeableTransactor creates a new write-only instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*StorageSlotUpgradeableTransactor, error) {
	contract, err := bindStorageSlotUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeableTransactor{contract: contract}, nil
}

// NewStorageSlotUpgradeableFilterer creates a new log filterer instance of StorageSlotUpgradeable, bound to a specific deployed contract.
func NewStorageSlotUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*StorageSlotUpgradeableFilterer, error) {
	contract, err := bindStorageSlotUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StorageSlotUpgradeableFilterer{contract: contract}, nil
}

// bindStorageSlotUpgradeable binds a generic wrapper to an already deployed contract.
func bindStorageSlotUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StorageSlotUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSlotUpgradeable.Contract.StorageSlotUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.StorageSlotUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.StorageSlotUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StorageSlotUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StorageSlotUpgradeable *StorageSlotUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StorageSlotUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// StringsUpgradeableMetaData contains all meta data concerning the StringsUpgradeable contract.
var StringsUpgradeableMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c70296c0d0fc0e50424a7f2e34667a9944d2fae42dbfd748ebba05656956865364736f6c63430008020033",
}

// StringsUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use StringsUpgradeableMetaData.ABI instead.
var StringsUpgradeableABI = StringsUpgradeableMetaData.ABI

// StringsUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StringsUpgradeableMetaData.Bin instead.
var StringsUpgradeableBin = StringsUpgradeableMetaData.Bin

// DeployStringsUpgradeable deploys a new Ethereum contract, binding an instance of StringsUpgradeable to it.
func DeployStringsUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StringsUpgradeable, error) {
	parsed, err := StringsUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StringsUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StringsUpgradeable{StringsUpgradeableCaller: StringsUpgradeableCaller{contract: contract}, StringsUpgradeableTransactor: StringsUpgradeableTransactor{contract: contract}, StringsUpgradeableFilterer: StringsUpgradeableFilterer{contract: contract}}, nil
}

// StringsUpgradeable is an auto generated Go binding around an Ethereum contract.
type StringsUpgradeable struct {
	StringsUpgradeableCaller     // Read-only binding to the contract
	StringsUpgradeableTransactor // Write-only binding to the contract
	StringsUpgradeableFilterer   // Log filterer for contract events
}

// StringsUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsUpgradeableSession struct {
	Contract     *StringsUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StringsUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsUpgradeableCallerSession struct {
	Contract *StringsUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// StringsUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsUpgradeableTransactorSession struct {
	Contract     *StringsUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// StringsUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsUpgradeableRaw struct {
	Contract *StringsUpgradeable // Generic contract binding to access the raw methods on
}

// StringsUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsUpgradeableCallerRaw struct {
	Contract *StringsUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// StringsUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsUpgradeableTransactorRaw struct {
	Contract *StringsUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStringsUpgradeable creates a new instance of StringsUpgradeable, bound to a specific deployed contract.
func NewStringsUpgradeable(address common.Address, backend bind.ContractBackend) (*StringsUpgradeable, error) {
	contract, err := bindStringsUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StringsUpgradeable{StringsUpgradeableCaller: StringsUpgradeableCaller{contract: contract}, StringsUpgradeableTransactor: StringsUpgradeableTransactor{contract: contract}, StringsUpgradeableFilterer: StringsUpgradeableFilterer{contract: contract}}, nil
}

// NewStringsUpgradeableCaller creates a new read-only instance of StringsUpgradeable, bound to a specific deployed contract.
func NewStringsUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*StringsUpgradeableCaller, error) {
	contract, err := bindStringsUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsUpgradeableCaller{contract: contract}, nil
}

// NewStringsUpgradeableTransactor creates a new write-only instance of StringsUpgradeable, bound to a specific deployed contract.
func NewStringsUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsUpgradeableTransactor, error) {
	contract, err := bindStringsUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsUpgradeableTransactor{contract: contract}, nil
}

// NewStringsUpgradeableFilterer creates a new log filterer instance of StringsUpgradeable, bound to a specific deployed contract.
func NewStringsUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsUpgradeableFilterer, error) {
	contract, err := bindStringsUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsUpgradeableFilterer{contract: contract}, nil
}

// bindStringsUpgradeable binds a generic wrapper to an already deployed contract.
func bindStringsUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StringsUpgradeable *StringsUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StringsUpgradeable.Contract.StringsUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StringsUpgradeable *StringsUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StringsUpgradeable.Contract.StringsUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StringsUpgradeable *StringsUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StringsUpgradeable.Contract.StringsUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StringsUpgradeable *StringsUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StringsUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StringsUpgradeable *StringsUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StringsUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StringsUpgradeable *StringsUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StringsUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// TwoWayCenterMetaData contains all meta data concerning the TwoWayCenter contract.
var TwoWayCenterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structInParam\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"CrossInFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structInParam\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"CrossIned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structInParam\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"CrossOuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structInParam\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"ForwardCrossOutFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structInParam\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"ForwardCrossOuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structInParam\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"Issued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oToken\",\"type\":\"address\"}],\"name\":\"StakingRewardReseted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"Supported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structInParam\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"Withdrawed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structInParam\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"WithdrawedToCenter\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UPGRADER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_centerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_edgeChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_edgeToken\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"fixAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ratioAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structOutParam\",\"name\":\"p\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"crossIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"crossOut\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"decimalDiff\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"eventHeights0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"eventHeights1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"eventIndex0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eventIndex1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"feeToTreasuryRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fixFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structOutParam\",\"name\":\"p\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"forwardCrossOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMsgSender\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"}],\"name\":\"getRoleKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_treasuryTo\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isClosed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isCoin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isInWhilelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"fromToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"toToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structInParam\",\"name\":\"p\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"issue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockBalances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ratioFeesHigh\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ratioFeesLow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ratioFeesMedium\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"remainHigh\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"remainLow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_centerToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_edgeChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_edgeToken\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oToken\",\"type\":\"address\"}],\"name\":\"resetStakingReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewardRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"toChainIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_fixFees\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_ratioFeesHigh\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_ratioFeesMedium\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_ratioFeesLow\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_remains\",\"type\":\"uint256[]\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeTo\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"setFeeToTreasuryRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isCoin\",\"type\":\"bool\"}],\"name\":\"setIsCoin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"setRewardRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"oTokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_srs\",\"type\":\"address[]\"}],\"name\":\"setStakingRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasuryTo\",\"type\":\"address\"}],\"name\":\"setTreasuryTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_isInWhitelist\",\"type\":\"bool\"}],\"name\":\"setWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sr\",\"outputs\":[{\"internalType\":\"contractIStakingReward\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"srs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"toCenterToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"toEdgeToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasuryTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"txHandled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"oToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"a217fddf": "DEFAULT_ADMIN_ROLE()",
		"f72c0d8b": "UPGRADER_ROLE()",
		"a766a392": "addToken(address,uint256,address)",
		"0324ef9c": "calculateFee(address,uint256,uint256)",
		"2c32c360": "calculateReward(address,uint256,uint256)",
		"9a8a0592": "chainId()",
		"1ca97185": "crossIn((uint256,address,address,uint256,address,uint256),string)",
		"df7e600a": "crossOut(address,uint256,address,uint256)",
		"bf814720": "decimalDiff(address)",
		"47e7ef24": "deposit(address,uint256)",
		"fc1aa909": "eventHeights0(uint256,uint256)",
		"05e82506": "eventHeights1(uint256)",
		"0826820e": "eventIndex0(uint256)",
		"15cd2c70": "eventIndex1()",
		"017e7e58": "feeTo()",
		"59ded2d7": "feeToTreasuryRatio(address)",
		"75d47f99": "fixFees(address,uint256)",
		"4ba9b352": "forwardCrossOut((uint256,address,address,uint256,address,uint256),string)",
		"7a6ce2e1": "getMsgSender()",
		"248a9ca3": "getRoleAdmin(bytes32)",
		"3ab796e7": "getRoleKey(address)",
		"2f2ff15d": "grantRole(bytes32,address)",
		"91d14854": "hasRole(bytes32,address)",
		"b4988fd0": "initialize(uint256,address,address)",
		"c2b6b58c": "isClosed()",
		"a611033e": "isCoin(address)",
		"a07b32c0": "isInWhilelist(address,address)",
		"af437845": "issue((uint256,address,address,uint256,address,address,uint256),string)",
		"a6ba8fc1": "lockBalances(address,uint256)",
		"ee5f7d96": "ratioFeesHigh(address,uint256)",
		"e4732500": "ratioFeesLow(address,uint256)",
		"3ff9bb66": "ratioFeesMedium(address,uint256)",
		"260ca4d5": "remainHigh(address)",
		"4297171f": "remainLow(address)",
		"97bb15c9": "removeToken(address,uint256,address)",
		"36568abe": "renounceRole(bytes32,address)",
		"0f4693c6": "resetStakingReward(address)",
		"d547741f": "revokeRole(bytes32,address)",
		"ced2a7ba": "rewardRatio(address)",
		"a5886f2d": "setFee(address,uint256[],uint256[],uint256[],uint256[],uint256[],uint256[])",
		"f46901ed": "setFeeTo(address)",
		"f88d0a2a": "setFeeToTreasuryRatio(address,uint256)",
		"d93e76bb": "setIsCoin(address,bool)",
		"41950ad3": "setRewardRatio(address,uint256)",
		"9773a000": "setStakingRewards(address[],address[])",
		"9d879990": "setThreshold(address,uint256)",
		"08b60fb8": "setTreasuryTo(address)",
		"1bdbce49": "setWhitelist(address,address,bool)",
		"df791d80": "sr()",
		"f3669d80": "srs(address)",
		"01ffc9a7": "supportsInterface(bytes4)",
		"c86ec2bf": "threshold(address)",
		"bafd2d19": "toCenterToken(uint256,address)",
		"527b0f64": "toEdgeToken(address,uint256)",
		"42079671": "treasuryTo()",
		"824df449": "txHandled(string)",
		"3659cfe6": "upgradeTo(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
		"4b2084e3": "withdraw(address,uint256,address,uint256)",
	},
	Bin: "0x60806040523480156200001157600080fd5b50600054610100900460ff16806200002c575060005460ff16155b620000945760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b606482015260840160405180910390fd5b600054610100900460ff16158015620000c0576000805460ff1961ff0019909116610100171660011790555b8015620000d3576000805461ff00191690555b50615a8480620000e46000396000f3fe60806040526004361061038c5760003560e01c8063824df449116101dc578063bf81472011610102578063df7e600a116100a0578063f46901ed1161006f578063f46901ed14610c22578063f72c0d8b14610c42578063f88d0a2a14610c76578063fc1aa90914610c965761038c565b8063df7e600a14610b66578063e473250014610b79578063ee5f7d9614610bb2578063f3669d8014610beb5761038c565b8063ced2a7ba116100dc578063ced2a7ba14610ad7578063d547741f14610b05578063d93e76bb14610b25578063df791d8014610b455761038c565b8063bf81472014610a61578063c2b6b58c14610a8f578063c86ec2bf14610aaa5761038c565b8063a217fddf1161017a578063a766a39211610149578063a766a392146109bf578063af437845146109df578063b4988fd0146109ff578063bafd2d1914610a1f5761038c565b8063a217fddf14610920578063a5886f2d14610935578063a611033e14610955578063a6ba8fc1146109865761038c565b806397bb15c9116101b657806397bb15c91461088e5780639a8a0592146108ae5780639d879990146108c4578063a07b32c0146108e45761038c565b8063824df4491461081257806391d148541461084e5780639773a0001461086e5761038c565b806336568abe116102c157806347e7ef241161025f578063527b0f641161022e578063527b0f641461074957806359ded2d71461078b57806375d47f99146107b95780637a6ce2e1146107f25761038c565b806347e7ef24146106e35780634b2084e3146106f65780634ba9b352146107165780634f1ef286146107365761038c565b80633ff9bb661161029b5780633ff9bb661461063b57806341950ad31461067457806342079671146106945780634297171f146106b55761038c565b806336568abe146105db5780633659cfe6146105fb5780633ab796e71461061b5761038c565b806315cd2c701161032e578063248a9ca311610308578063248a9ca31461053d578063260ca4d51461056d5780632c32c3601461059b5780632f2ff15d146105bb5761038c565b806315cd2c70146104e65780631bdbce49146104fd5780631ca971851461051d5761038c565b806305e825061161036a57806305e825061461043a5780630826820e1461047657806308b60fb8146104a45780630f4693c6146104c65761038c565b8063017e7e581461039157806301ffc9a7146103cf5780630324ef9c146103ff575b600080fd5b34801561039d57600080fd5b50610100546103b2906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156103db57600080fd5b506103ef6103ea36600461525f565b610ccf565b60405190151581526020016103c6565b34801561040b57600080fd5b5061041f61041a366004615180565b610d08565b604080519384526020840192909252908201526060016103c6565b34801561044657600080fd5b50610468610455366004615225565b6101176020526000908152604090205481565b6040519081526020016103c6565b34801561048257600080fd5b50610468610491366004615225565b6101146020526000908152604090205481565b3480156104b057600080fd5b506104c46104bf366004614eb2565b610ef2565b005b3480156104d257600080fd5b506104c46104e1366004614eb2565b610f6b565b3480156104f257600080fd5b506104686101165481565b34801561050957600080fd5b506104c4610518366004614efe565b610ff0565b34801561052957600080fd5b506104c4610538366004615369565b6110bd565b34801561054957600080fd5b50610468610558366004615225565b60009081526065602052604090206001015490565b34801561057957600080fd5b50610468610588366004614eb2565b61010d6020526000908152604090205481565b3480156105a757600080fd5b506104686105b6366004615180565b6116b5565b3480156105c757600080fd5b506104c46105d636600461523d565b611a1e565b3480156105e757600080fd5b506104c46105f636600461523d565b611a4a565b34801561060757600080fd5b506104c4610616366004614eb2565b611ac8565b34801561062757600080fd5b50610468610636366004614eb2565b611aef565b34801561064757600080fd5b506104686106563660046150d9565b61010b60209081526000928352604080842090915290825290205481565b34801561068057600080fd5b506104c461068f3660046150d9565b611b2e565b3480156106a057600080fd5b50610101546103b2906001600160a01b031681565b3480156106c157600080fd5b506104686106d0366004614eb2565b61010e6020526000908152604090205481565b6104c46106f13660046150d9565b611b72565b34801561070257600080fd5b506104c461071136600461513d565b611e20565b34801561072257600080fd5b506104c4610731366004615369565b612245565b6104c461074436600461507a565b6127c5565b34801561075557600080fd5b506103b26107643660046150d9565b6101046020908152600092835260408084209091529082529020546001600160a01b031681565b34801561079757600080fd5b506104686107a6366004614eb2565b6101116020526000908152604090205481565b3480156107c557600080fd5b506104686107d43660046150d9565b61010860209081526000928352604080842090915290825290205481565b3480156107fe57600080fd5b5060408051338082526020820152016103c6565b34801561081e57600080fd5b506103ef61082d366004615287565b80516020818301810180516101068252928201919093012091525460ff1681565b34801561085a57600080fd5b506103ef61086936600461523d565b6127da565b34801561087a57600080fd5b506104c46108893660046151b2565b612805565b34801561089a57600080fd5b506104c46108a9366004615102565b6129d5565b3480156108ba57600080fd5b5061046860ff5481565b3480156108d057600080fd5b506104c46108df3660046150d9565b612b24565b3480156108f057600080fd5b506103ef6108ff366004614ecc565b61010760209081526000928352604080842090915290825290205460ff1681565b34801561092c57600080fd5b50610468600081565b34801561094157600080fd5b506104c4610950366004614f44565b612b55565b34801561096157600080fd5b506103ef610970366004614eb2565b6101126020526000908152604090205460ff1681565b34801561099257600080fd5b506104686109a13660046150d9565b61010960209081526000928352604080842090915290825290205481565b3480156109cb57600080fd5b506104c46109da366004615102565b612f62565b3480156109eb57600080fd5b506104c46109fa3660046152ba565b61315e565b348015610a0b57600080fd5b506104c4610a1a366004615412565b6134c3565b348015610a2b57600080fd5b506103b2610a3a36600461523d565b6101036020908152600092835260408084209091529082529020546001600160a01b031681565b348015610a6d57600080fd5b50610468610a7c366004614eb2565b6101056020526000908152604090205481565b348015610a9b57600080fd5b50610113546103ef9060ff1681565b348015610ab657600080fd5b50610468610ac5366004614eb2565b60fb6020526000908152604090205481565b348015610ae357600080fd5b50610468610af2366004614eb2565b6101106020526000908152604090205481565b348015610b1157600080fd5b506104c4610b2036600461523d565b6135b7565b348015610b3157600080fd5b506104c4610b40366004615044565b6135dd565b348015610b5157600080fd5b50610102546103b2906001600160a01b031681565b6104c4610b7436600461513d565b613630565b348015610b8557600080fd5b50610468610b943660046150d9565b61010c60209081526000928352604080842090915290825290205481565b348015610bbe57600080fd5b50610468610bcd3660046150d9565b61010a60209081526000928352604080842090915290825290205481565b348015610bf757600080fd5b506103b2610c06366004614eb2565b61010f602052600090815260409020546001600160a01b031681565b348015610c2e57600080fd5b506104c4610c3d366004614eb2565b6139ea565b348015610c4e57600080fd5b506104687f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e381565b348015610c8257600080fd5b506104c4610c913660046150d9565b613a5a565b348015610ca257600080fd5b50610468610cb1366004615444565b61011560209081526000928352604080842090915290825290205481565b60006001600160e01b03198216637965db0b60e01b1480610d0057506301ffc9a760e01b6001600160e01b03198316145b90505b919050565b6001600160a01b0383166000908152610107602090815260408083203384529091528120548190819060ff1615610d40575082610ee9565b6001600160a01b038616600090815261010960209081526040808320888452909152812054819080871115610d9b576001600160a01b038916600090815261010c602090815260408083208b84529091529020549250610da8565b610da58782615955565b91505b6001600160a01b038916600090815261010e60205260409020548211610df4576001600160a01b038916600090815261010a602090815260408083208b84529091529020549250610e8e565b6001600160a01b038916600090815261010e602052604090205482118015610e3557506001600160a01b038916600090815261010d60205260409020548211155b15610e66576001600160a01b038916600090815261010b602090815260408083208b84529091529020549250610e8e565b6001600160a01b038916600090815261010c602090815260408083208b845290915290205492505b6001600160a01b0389166000908152610108602090815260408083208b84529091528120549650610ebf8789615955565b9050610ecb8482613ae4565b955085610ed8888a615955565b610ee29190615955565b9450505050505b93509350939050565b610efd6000336127da565b610f225760405162461bcd60e51b8152600401610f1990615698565b60405180910390fd5b6001600160a01b038116610f485760405162461bcd60e51b8152600401610f19906156cf565b61010180546001600160a01b0319166001600160a01b0392909216919091179055565b610f766000336127da565b610f925760405162461bcd60e51b8152600401610f1990615698565b6001600160a01b038116600081815261010f602090815260409182902080546001600160a01b031916905590519182527f4b8830da429e7f8c1dc368e9a9a3dc21aa6e3e08d115679e616b1fedaf420aeb910160405180910390a150565b610ffb6000336127da565b6110175760405162461bcd60e51b8152600401610f1990615698565b6001600160a01b038084166000908152610107602090815260408083209386168352929052205460ff16151581151514156110825760405162461bcd60e51b815260206004820152600b60248201526a6572726f7220737461746560a81b6044820152606401610f19565b6001600160a01b0392831660009081526101076020908152604080832094909516825292909252919020805460ff1916911515919091179055565b8151600090815261010360209081526040808320828601516001600160a01b039081168552925282205416906110f282611aef565b90506110fe81336127da565b61111a5760405162461bcd60e51b8152600401610f19906155c3565b826101068160405161112c9190615486565b9081526040519081900360200190205460ff161561115c5760405162461bcd60e51b8152600401610f199061561f565b6101135460ff16156111995760405162461bcd60e51b815260206004820152600660248201526518db1bdcd95960d21b6044820152606401610f19565b610113805460ff1916600117905560ff546060860151146111ec5760405162461bcd60e51b815260206004820152600d60248201526c31b430b4b724b21032b93937b960991b6044820152606401610f19565b8451600090815261010360209081526040808320828901516001600160a01b039081168552908352818420548116808552610104845282852060608b01518652909352922054909116816112825760405162461bcd60e51b815260206004820152601760248201527f63656e746572546f6b656e206e6f7420737570706f72740000000000000000006044820152606401610f19565b6001600160a01b0381166112d05760405162461bcd60e51b8152602060048201526015602482015274195919d9551bdad95b881b9bdd081cdd5c1c1bdc9d605a1b6044820152606401610f19565b60006112f08860000151848a604001518b608001518c60a001518c613b06565b905080156116715760016101068860405161130b9190615486565b9081526040805160209281900383019020805460ff19169315159390931790925560a08a01516001600160a01b038616600090815261010983528381208c51825290925291812080549091906113629084906157d7565b909155505060a08801516001600160a01b03841660009081526101096020908152604080832060ff548452909152902054106115445760008060006113b0868c606001518d60a00151610d08565b6001600160a01b03881660009081526101126020526040902054929550909350915060ff16156114125760808b01516001600160a01b0386166000908152610105602052604090205461140d919061140890846157ef565b613d94565b611451565b60808b01516001600160a01b03861660009081526101056020526040902054611451919061144090846157ef565b6001600160a01b0388169190613ead565b61145c868484613f03565b6001600160a01b03861660009081526101096020908152604080832060608f0151845290915281208054839290611494908490615955565b925050819055507fe63fe074c53764bbe6e6f716575da08b4e5546ea802aa49ba761f2c1799715416040518060e001604052808d6000015181526020018d602001516001600160a01b031681526020018d604001516001600160a01b031681526020018d606001518152602001876001600160a01b031681526020018d608001516001600160a01b03168152602001838152506040516115349190615722565b60405180910390a1505050611671565b608088015160a08901516040516340c10f1960e01b81526001600160a01b038616926340c10f199261157892600401615577565b602060405180830381600087803b15801561159257600080fd5b505af11580156115a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115ca9190615209565b507f1ecf44265cb8eead95bfe72d8e66611159bbbe6c385212224660bf7ee22874116040518060e001604052808a6000015181526020018a602001516001600160a01b031681526020018a604001516001600160a01b031681526020018a606001518152602001846001600160a01b031681526020018a608001516001600160a01b031681526020018a60a001518152506040516116689190615722565b60405180910390a15b5050506001610116600082825461168891906157d7565b9091555050610116546000908152610117602052604090204390555050610113805460ff19169055505050565b6001600160a01b038316600090815261011060205260408120546116db57506000611a17565b6001600160a01b03841660008181526101096020908152604080832087845282528083205493835261010d909152902054811061171c576000915050611a17565b600061172884836157d7565b6001600160a01b0387166000908152610110602090815260408083205461010a83528184208a8552909252822054929350909161176491613ae4565b6001600160a01b0388166000908152610110602090815260408083205461010b83528184208b855290925282205492935090916117a091613ae4565b6001600160a01b038916600090815261010e60205260409020549091508411611904576001600160a01b038816600090815261010e602052604090205483116117f4576117ed8683613ae4565b94506118ff565b6001600160a01b038816600090815261010d60205260409020548311611888576001600160a01b038816600090815261010e602052604081205461184590849061183f908890615955565b90613ae4565b6001600160a01b038a16600090815261010e60205260408120549192509061187390849061183f9088615955565b905061187f81836157d7565b965050506118ff565b6001600160a01b038816600090815261010e60205260408120546118b390849061183f908890615955565b6001600160a01b038a16600090815261010e602090815260408083205461010d90925282205492935090916118ee91859161183f9190615955565b90506118fa81836157d7565b965050505b61197c565b6001600160a01b038816600090815261010d6020526040902054841161197c576001600160a01b038816600090815261010d6020526040902054831161194e576117ed8682613ae4565b6001600160a01b038816600090815261010d602052604090205461197990829061183f908790615955565b94505b610101546040516370a0823160e01b81526001600160a01b0391821660048201526000918a16906370a082319060240160206040518083038186803b1580156119c457600080fd5b505afa1580156119d8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119fc91906153fa565b9050858111611a11579450611a179350505050565b50505050505b9392505050565b600082815260656020526040902060010154611a3b81335b614215565b611a458383614279565b505050565b6001600160a01b0381163314611aba5760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152608401610f19565b611ac482826142ff565b5050565b611ad181614366565b611aec81604051806020016040528060008152506000614391565b50565b6040516bffffffffffffffffffffffff19606083901b166020820152600090603401604051602081830303815290604052805190602001209050919050565b611b396000336127da565b611b555760405162461bcd60e51b8152600401610f1990615698565b6001600160a01b0390911660009081526101106020526040902055565b60ff546000908152610103602090815260408083206001600160a01b0380871685529252909120541680611bb85760405162461bcd60e51b8152600401610f19906155fa565b6001600160a01b0383166000908152610112602052604090205460ff1615611c1d57348214611c185760405162461bcd60e51b815260206004820152600c60248201526b30b6b7bab73a1032b93937b960a11b6044820152606401610f19565b611c32565b611c326001600160a01b03841633308561454e565b6001600160a01b0383811660009081526101056020526040902054908216906340c10f19903390611c639086615936565b6040518363ffffffff1660e01b8152600401611c80929190615577565b602060405180830381600087803b158015611c9a57600080fd5b505af1158015611cae573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611cd29190615209565b5060ff546001600160a01b03841660009081526101056020526040902054611d079183913390611d029087615936565b614586565b6001600160a01b03831660009081526101056020526040902054611d2b9083615936565b6001600160a01b03821660009081526101096020908152604080832060ff54845290915281208054909190611d619084906157d7565b909155505060ff546001600160a01b038416600090815261010560205260409020547f984a71c9d95fd4794aeba33ae72edfec22053fde75488d63abef9dc69ee795af919085903390611db49087615936565b604051611de494939291909384526001600160a01b03928316602085015291166040830152606082015260800190565b60405180910390a15060016101166000828254611e0191906157d7565b9091555050610116546000908152610117602052604090204390555050565b6101135460ff1615611e5d5760405162461bcd60e51b815260206004820152600660248201526518db1bdcd95960d21b6044820152606401610f19565b610113805460ff191660011790556001600160a01b0384166000908152610109602090815260408083208684529091529020548390821115611ed45760405162461bcd60e51b815260206004820152601060248201526f6e6f7420656e6f756768206c6971756960801b6044820152606401610f19565b6001600160a01b038086166000908152610104602090815260408083208884529091529020541680611f405760405162461bcd60e51b8152602060048201526015602482015274195919d9481d1bdad9481b9bdd081cdd5c1c1bdc9d605a1b6044820152606401610f19565b6000806000611f50898988610d08565b925092509250611f61898484613f03565b604051632770a7eb60e21b81526001600160a01b038a1690639dc29fac90611f8f9033908a90600401615577565b602060405180830381600087803b158015611fa957600080fd5b505af1158015611fbd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fe19190615209565b506001600160a01b0389166000908152610109602090815260408083208b845290915281208054889290612016908490615955565b909155505060ff5488141561212c576001600160a01b0384166000908152610112602052604090205460ff1615612076576001600160a01b0384166000908152610105602052604090205461207190339061140890846157ef565b6120b1565b6001600160a01b038416600090815261010560205260409020546120b19033906120a090846157ef565b6001600160a01b0387169190613ead565b6040805160e08101825260ff5481526001600160a01b03808c1660208301523382840152606082018b90528681166080830152891660a082015260c0810183905290517f7f50f8e7f2e365c46040fcc45f0f1ee1b3f7d8c3befc68282bfa0562f4f1d2b09161211f91615722565b60405180910390a16121a3565b6040805160e08101825260ff5481526001600160a01b03808c1660208301523382840152606082018b90528681166080830152891660a082015260c0810183905290517ff0bc31bd6717a3afcc9a7f44d5b6d8638707a8ee613a6d3b38ec03e9f0392d679161219a91615722565b60405180910390a15b50505050600161011660008282546121bb91906157d7565b90915550506101165460009081526101176020908152604080832043905583835261011490915281208054600192906121f59084906157d7565b9091555050600081815261011460205260408120546122159060016157d7565b6000928352610115602090815260408085209285529190529091204390555050610113805460ff19169055505050565b8151600090815261010360209081526040808320828601516001600160a01b0390811685529252822054169061227a82611aef565b905061228681336127da565b6122a25760405162461bcd60e51b8152600401610f19906155c3565b82610106816040516122b49190615486565b9081526040519081900360200190205460ff16156122e45760405162461bcd60e51b8152600401610f199061561f565b6060850151855181141561232a5760405162461bcd60e51b815260206004820152600d60248201526c31b430b4b724b21032b93937b960991b6044820152606401610f19565b8551600090815261010360209081526040808320828a01516001600160a01b03908116855292529091205416806123a35760405162461bcd60e51b815260206004820152601760248201527f6e6f7420737570706f72742063656e746572546f6b656e0000000000000000006044820152606401610f19565b6001600160a01b0380821660009081526101046020908152604080832060608c0151845290915290205416806124135760405162461bcd60e51b81526020600482015260156024820152743737ba1039bab83837b93a1032b233b2aa37b5b2b760591b6044820152606401610f19565b60006124338960000151848b604001518c608001518d60a001518d613b06565b9050801561272e5760016101068960405161244e9190615486565b9081526040805160209281900383019020805460ff19169315159390931790925560a08b01516001600160a01b038616600090815261010983528381208d51825290925291812080549091906124a59084906157d7565b909155505060a08901516001600160a01b03841660009081526101096020908152604080832060608e01518452909152902054106126015760008060006124f5868d606001518e60a00151610d08565b925092509250612506868484613f03565b806101096000886001600160a01b03166001600160a01b0316815260200190815260200160002060008e60600151815260200190815260200160002060008282546125519190615955565b925050819055507f715d2e6514c297c9e18ba16c3e3972b7562b34f8f185c21a5cc305324039c77a6040518060e001604052808e6000015181526020018e602001516001600160a01b031681526020018e604001516001600160a01b031681526020018e606001518152602001876001600160a01b031681526020018e608001516001600160a01b03168152602001838152506040516125f19190615722565b60405180910390a150505061272e565b608089015160a08a01516040516340c10f1960e01b81526001600160a01b038616926340c10f199261263592600401615577565b602060405180830381600087803b15801561264f57600080fd5b505af1158015612663573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126879190615209565b507fa22f2d032a9fe57cd0476bed912548f17dccb0695666d445bed754d98219b6a06040518060e001604052808b6000015181526020018b602001516001600160a01b031681526020018b604001516001600160a01b031681526020018b606001518152602001846001600160a01b031681526020018b608001516001600160a01b031681526020018b60a001518152506040516127259190615722565b60405180910390a15b5050506001610116600082825461274591906157d7565b909155505061011654600090815261011760209081526040808320439055838352610114909152812080546001929061277f9084906157d7565b90915550506000818152610114602052604081205461279f9060016157d7565b600092835261011560209081526040808520928552919052909120439055505050505050565b6127ce82614366565b611ac482826001614391565b60009182526065602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6128106000336127da565b61282c5760405162461bcd60e51b8152600401610f1990615698565b80518251146128695760405162461bcd60e51b81526020600482015260096024820152680dcdee840dac2e8c6d60bb1b6044820152606401610f19565b60005b8251811015611a455760006001600160a01b03168382815181106128a057634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b031614156128cf5760405162461bcd60e51b8152600401610f19906156cf565b60006001600160a01b03168282815181106128fa57634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b031614156129295760405162461bcd60e51b8152600401610f19906156cf565b81818151811061294957634e487b7160e01b600052603260045260246000fd5b602002602001015161010f600085848151811061297657634e487b7160e01b600052603260045260246000fd5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b0316021790555080806129cd906159d2565b91505061286c565b6129e06000336127da565b6129fc5760405162461bcd60e51b8152600401610f1990615698565b6000828152610103602090815260408083206001600160a01b03858116855292529091205416612a665760405162461bcd60e51b815260206004820152601560248201527418d95b9d195c951bdad95b881b9bdd08195e1a5cdd605a1b6044820152606401610f19565b6001600160a01b0383811660009081526101046020908152604080832086845290915290205416612acf5760405162461bcd60e51b8152602060048201526013602482015272195919d9551bdad95b881b9bdd08195e1a5cdd606a1b6044820152606401610f19565b6000828152610103602090815260408083206001600160a01b039485168452825280832080546001600160a01b0319908116909155959093168252610104815282822093825292909252902080549091169055565b612b2f6000336127da565b612b4b5760405162461bcd60e51b8152600401610f1990615698565b611ac482826145b7565b612b606000336127da565b612b7c5760405162461bcd60e51b8152600401610f1990615698565b8451865114612bc45760405162461bcd60e51b8152602060048201526014602482015273040e8de86d0c2d2dc92c840dcdee840dac2e8c6d60631b6044820152606401610f19565b8451845114612be55760405162461bcd60e51b8152600401610f19906156f5565b8451835114612c065760405162461bcd60e51b8152600401610f19906156f5565b8451825114612c275760405162461bcd60e51b8152600401610f19906156f5565b8051600214612c6b5760405162461bcd60e51b815260206004820152601060248201526f0e4c2dad2dce640dcdee840dac2e8c6d60831b6044820152606401610f19565b60005b8651811015612eb257858181518110612c9757634e487b7160e01b600052603260045260246000fd5b602002602001015161010860008a6001600160a01b03166001600160a01b031681526020019081526020016000206000898481518110612ce757634e487b7160e01b600052603260045260246000fd5b6020026020010151815260200190815260200160002081905550848181518110612d2157634e487b7160e01b600052603260045260246000fd5b602002602001015161010a60008a6001600160a01b03166001600160a01b031681526020019081526020016000206000898481518110612d7157634e487b7160e01b600052603260045260246000fd5b6020026020010151815260200190815260200160002081905550838181518110612dab57634e487b7160e01b600052603260045260246000fd5b602002602001015161010b60008a6001600160a01b03166001600160a01b031681526020019081526020016000206000898481518110612dfb57634e487b7160e01b600052603260045260246000fd5b6020026020010151815260200190815260200160002081905550828181518110612e3557634e487b7160e01b600052603260045260246000fd5b602002602001015161010c60008a6001600160a01b03166001600160a01b031681526020019081526020016000206000898481518110612e8557634e487b7160e01b600052603260045260246000fd5b60200260200101518152602001908152602001600020819055508080612eaa906159d2565b915050612c6e565b5080600081518110612ed457634e487b7160e01b600052603260045260246000fd5b602002602001015161010d6000896001600160a01b03166001600160a01b031681526020019081526020016000208190555080600181518110612f2757634e487b7160e01b600052603260045260246000fd5b602002602001015161010e6000896001600160a01b03166001600160a01b031681526020019081526020016000208190555050505050505050565b612f6d6000336127da565b612f895760405162461bcd60e51b8152600401610f1990615698565b6000828152610103602090815260408083206001600160a01b0385811685529252909120541615612ff05760405162461bcd60e51b815260206004820152601160248201527018d95b9d195c951bdad95b88195e1a5cdd607a1b6044820152606401610f19565b6001600160a01b0383811660009081526101046020908152604080832086845290915290205416156130565760405162461bcd60e51b815260206004820152600f60248201526e195919d9551bdad95b88195e1a5cdd608a1b6044820152606401610f19565b6000828152610103602090815260408083206001600160a01b0380861680865291845282852080549189166001600160a01b031992831681179091558552610104845282852087865290935292208054909116909117905560ff54821415611a4557806001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b1580156130f157600080fd5b505afa158015613105573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131299190615465565b61313490601261596c565b61313f90600a615864565b6001600160a01b03821660009081526101056020526040902055505050565b8151600090815261010360209081526040808320828601516001600160a01b0390811685529252822054169061319382611aef565b905061319f81336127da565b6131bb5760405162461bcd60e51b8152600401610f19906155c3565b82610106816040516131cd9190615486565b9081526040519081900360200190205460ff16156131fd5760405162461bcd60e51b8152600401610f199061561f565b60ff548560600151146132465760405162461bcd60e51b81526020600482015260116024820152700c6d0c2d2dc92c840dcdee840dac2e8c6d607b1b6044820152606401610f19565b8451600090815261010360209081526040808320828901516001600160a01b03908116855292529091205416806132b35760405162461bcd60e51b815260206004820152601160248201527018d95b9d195c951bdad95b88195e1a5cdd607a1b6044820152606401610f19565b6001600160a01b038181166000908152610104602090815260408083208a518452909152902054166133195760405162461bcd60e51b815260206004820152600f60248201526e195919d9551bdad95b88195e1a5cdd608a1b6044820152606401610f19565b600061333987600001518389604001518a60a001518b60c001518b613b06565b9050801561348b576001610106876040516133549190615486565b908152604051908190036020018120805492151560ff199093169290921790915560a088015160c08901516340c10f1960e01b83526001600160a01b038516926340c10f19926133a8929091600401615577565b602060405180830381600087803b1580156133c257600080fd5b505af11580156133d6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906133fa9190615209565b5061341382886000015189604001518a60c00151614586565b60c08701516001600160a01b0383166000908152610109602090815260408083208b5184529091528120805490919061344d9084906157d7565b90915550506040517f082ec0a3f5481f56ed1aa06ed4b5f96a0b82a7f1ce3342041b20b8161fabeb9290613482908990615722565b60405180910390a15b5050600161011660008282546134a191906157d7565b9091555050610116546000908152610117602052604090204390555050505050565b600054610100900460ff16806134dc575060005460ff16155b6134f85760405162461bcd60e51b8152600401610f199061564a565b600054610100900460ff16158015613523576000805460ff1961ff0019909116610100171660011790555b61352b61461d565b6135336146a9565b61353e600033611a1e565b6135687f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e333611a1e565b60ff84905561010080546001600160a01b038086166001600160a01b03199283161790925561010180549285169290911691909117905580156135b1576000805461ff00191690555b50505050565b6000828152606560205260409020600101546135d38133611a36565b611a4583836142ff565b6135e86000336127da565b6136045760405162461bcd60e51b8152600401610f1990615698565b6001600160a01b0391909116600090815261011260205260409020805460ff1916911515919091179055565b8260ff548414156136755760405162461bcd60e51b815260206004820152600f60248201526e3a37a1b430b4b724b21032b93937b960891b6044820152606401610f19565b60ff546000908152610103602090815260408083206001600160a01b03808a16855292529091205416806136bb5760405162461bcd60e51b8152600401610f19906155fa565b6001600160a01b0380821660009081526101046020908152604080832089845290915290205416806136ff5760405162461bcd60e51b8152600401610f19906155fa565b6001600160a01b038716600090815261010560205260409020546137239085615936565b6001600160a01b0383166000908152610109602090815260408083208a845290915290205410156137895760405162461bcd60e51b815260206004820152601060248201526f6e6f7420656e6f756768206c6971756960801b6044820152606401610f19565b6001600160a01b0387166000908152610112602052604090205460ff16156137ee578334146137e95760405162461bcd60e51b815260206004820152600c60248201526b30b6b7bab73a1032b93937b960a11b6044820152606401610f19565b613803565b6138036001600160a01b03881633308761454e565b6001600160a01b03871660009081526101056020526040812054819081906138339086908b9061041a908b615936565b925092509250613844858484613f03565b6001600160a01b038a16600090815261010560205260409020546138689088615936565b6001600160a01b03861660009081526101096020908152604080832060ff5484529091528120805490919061389e9084906157d7565b90915550506001600160a01b0385166000908152610109602090815260408083208c8452909152812080548392906138d7908490615955565b90915550506040805160e08101825260ff5481526001600160a01b03808d1660208301523382840152606082018c905286811660808301528a1660a082015260c0810183905290517fe8ae00f8b99ac422d841599f0cc887e0d20c0bf99271117d943e474f280f11459161394a91615722565b60405180910390a150505050506001610116600082825461396b91906157d7565b90915550506101165460009081526101176020908152604080832043905583835261011490915281208054600192906139a59084906157d7565b9091555050600081815261011460205260408120546139c59060016157d7565b6000928352610115602090815260408085209285529190529091204390555050505050565b6139f56000336127da565b613a115760405162461bcd60e51b8152600401610f1990615698565b6001600160a01b038116613a375760405162461bcd60e51b8152600401610f19906156cf565b61010080546001600160a01b0319166001600160a01b0392909216919091179055565b613a656000336127da565b613a815760405162461bcd60e51b8152600401610f1990615698565b670de0b6b3a7640000811115613ac75760405162461bcd60e51b815260206004820152600b60248201526a3930ba34b79032b93937b960a91b6044820152606401610f19565b6001600160a01b0390911660009081526101116020526040902055565b6000613af26012600a615855565b613afc8484614710565b611a1791906157ef565b6001600160a01b038516600090815260fb6020526040812054613b845760405162461bcd60e51b815260206004820152603060248201527f50726f706f73616c566f74653a207468726573686f6c642073686f756c64206260448201526f0652067726561746572207468616e20360841b6064820152608401610f19565b6001600160a01b038616600090815260fb60209081526040808320549051909291613bbb918b918b918b918b918b918b9101615517565b60408051601f198184030181529181528151602092830120600081815260fc90935291205490915060ff1615613c335760405162461bcd60e51b815260206004820152601860248201527f5f766f74653a3a70726f706f73616c2066696e697368656400000000000000006044820152606401610f19565b600081815260fd6020908152604080832033845290915290205460ff1615613c9d5760405162461bcd60e51b815260206004820152601760248201527f5f766f74653a3a6d73672e73656e64657220766f7465640000000000000000006044820152606401610f19565b600081815260fe6020526040902054613cb790600161471c565b600082815260fe6020818152604080842094855560fd825280842033855282528320805460ff19166001179055918490529052548211613d1057600081815260fc60205260409020805460ff1916600190811790915592505b600081815260fe60209081526040918290205482516001600160a01b038c811682528b811693820193909352918916828401526060820188905233608083015260a082015260c0810184905290517fe458477b94285c9f254a407e2614c39f55230b527ff7f51d76fdf33e251ae9369181900360e00190a150509695505050505050565b80471015613de45760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a20696e73756666696369656e742062616c616e63650000006044820152606401610f19565b6000826001600160a01b03168260405160006040518083038185875af1925050503d8060008114613e31576040519150601f19603f3d011682016040523d82523d6000602084013e613e36565b606091505b5050905080611a455760405162461bcd60e51b815260206004820152603a60248201527f416464726573733a20756e61626c6520746f2073656e642076616c75652c207260448201527f6563697069656e74206d617920686176652072657665727465640000000000006064820152608401610f19565b611a458363a9059cbb60e01b8484604051602401613ecc929190615577565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152614728565b8115613f9257610100546040516340c10f1960e01b81526001600160a01b03858116926340c10f1992613f3e92909116908690600401615577565b602060405180830381600087803b158015613f5857600080fd5b505af1158015613f6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613f909190615209565b505b8015611a45576001600160a01b03838116600090815261010f60205260409020541661404557610101546040516340c10f1960e01b81526001600160a01b03858116926340c10f1992613fed92909116908590600401615577565b602060405180830381600087803b15801561400757600080fd5b505af115801561401b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061403f9190615209565b50611a45565b6001600160a01b0383166000908152610111602052604081205461406a908390613ae4565b905060006140788284615955565b610101546040516340c10f1960e01b81529192506001600160a01b03808816926340c10f19926140ae9216908690600401615577565b602060405180830381600087803b1580156140c857600080fd5b505af11580156140dc573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141009190615209565b506001600160a01b03808616600081815261010f6020526040908190205490516340c10f1960e01b815291926340c10f199261414492909116908590600401615577565b602060405180830381600087803b15801561415e57600080fd5b505af1158015614172573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906141969190615209565b506001600160a01b03858116600090815261010f60205260409081902054905163246132f960e01b81526004810184905262015180602482015291169063246132f990604401600060405180830381600087803b1580156141f657600080fd5b505af115801561420a573d6000803e3d6000fd5b505050505050505050565b61421f82826127da565b611ac457614237816001600160a01b031660146147fa565b6142428360206147fa565b6040516020016142539291906154a2565b60408051601f198184030181529082905262461bcd60e51b8252610f1991600401615590565b61428382826127da565b611ac45760008281526065602090815260408083206001600160a01b03851684529091529020805460ff191660011790556142bb3390565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b61430982826127da565b15611ac45760008281526065602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b7f189ab7a9244df0848122154315af71fe140f3db0fe014031783b0946b8c9d2e3611ac48133611a36565b60006143c47f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b90506143cf846149dc565b6000835111806143dc5750815b156143ed576143eb8484614a81565b505b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd9143805460ff1661454757805460ff191660011781556040516001600160a01b038316602482015261446c90869060440160408051601f198184030181529190526020810180516001600160e01b0316631b2ce7f360e11b179052614a81565b50805460ff191681557f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b038381169116146145095760405162461bcd60e51b815260206004820152602f60248201527f45524331393637557067726164653a207570677261646520627265616b73206660448201526e75727468657220757067726164657360881b6064820152608401610f19565b614512856149dc565b6040516001600160a01b038616907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a25b5050505050565b6040516001600160a01b03808516602483015283166044820152606481018290526135b19085906323b872dd60e01b90608401613ecc565b60006145938585846116b5565b905080156145475761010154614547906001600160a01b038781169116858461454e565b6001600160a01b038216600081815260fb60209081526040918290208054908590558251938452908301819052908201839052907fb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e449060600160405180910390a1505050565b600054610100900460ff1680614636575060005460ff16155b6146525760405162461bcd60e51b8152600401610f199061564a565b600054610100900460ff1615801561467d576000805460ff1961ff0019909116610100171660011790555b614685614b6c565b61468d614b6c565b614695614b6c565b8015611aec576000805461ff001916905550565b600054610100900460ff16806146c2575060005460ff16155b6146de5760405162461bcd60e51b8152600401610f199061564a565b600054610100900460ff16158015614685576000805460ff1961ff00199091166101001716600117905561468d614b6c565b6000611a178284615936565b6000611a1782846157d7565b600061477d826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316614bdf9092919063ffffffff16565b805190915015611a45578080602001905181019061479b9190615209565b611a455760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b6064820152608401610f19565b60606000614809836002615936565b6148149060026157d7565b67ffffffffffffffff81111561483a57634e487b7160e01b600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015614864576020820181803683370190505b509050600360fc1b8160008151811061488d57634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a905350600f60fb1b816001815181106148ca57634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a90535060006148ee846002615936565b6148f99060016157d7565b90505b600181111561498d576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811061493b57634e487b7160e01b600052603260045260246000fd5b1a60f81b82828151811061495f57634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a90535060049490941c93614986816159bb565b90506148fc565b508315611a175760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610f19565b803b614a405760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610f19565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b6060823b614ae05760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610f19565b600080846001600160a01b031684604051614afb9190615486565b600060405180830381855af49150503d8060008114614b36576040519150601f19603f3d011682016040523d82523d6000602084013e614b3b565b606091505b5091509150614b638282604051806060016040528060278152602001615a2860279139614bf6565b95945050505050565b600054610100900460ff1680614b85575060005460ff16155b614ba15760405162461bcd60e51b8152600401610f199061564a565b600054610100900460ff16158015614695576000805460ff1961ff0019909116610100171660011790558015611aec576000805461ff001916905550565b6060614bee8484600085614c2f565b949350505050565b60608315614c05575081611a17565b825115614c155782518084602001fd5b8160405162461bcd60e51b8152600401610f199190615590565b606082471015614c905760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b6064820152608401610f19565b843b614cde5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e74726163740000006044820152606401610f19565b600080866001600160a01b03168587604051614cfa9190615486565b60006040518083038185875af1925050503d8060008114614d37576040519150601f19603f3d011682016040523d82523d6000602084013e614d3c565b606091505b5091509150614d4c828286614bf6565b979650505050505050565b600067ffffffffffffffff831115614d7157614d71615a03565b614d84601f8401601f1916602001615782565b9050828152838383011115614d9857600080fd5b828260208301376000602084830101529392505050565b80356001600160a01b0381168114610d0357600080fd5b600082601f830112614dd6578081fd5b81356020614deb614de6836157b3565b615782565b8281528181019085830183850287018401881015614e07578586fd5b855b85811015614e2c57614e1a82614daf565b84529284019290840190600101614e09565b5090979650505050505050565b600082601f830112614e49578081fd5b81356020614e59614de6836157b3565b8281528181019085830183850287018401881015614e75578586fd5b855b85811015614e2c57813584529284019290840190600101614e77565b600082601f830112614ea3578081fd5b611a1783833560208501614d57565b600060208284031215614ec3578081fd5b611a1782614daf565b60008060408385031215614ede578081fd5b614ee783614daf565b9150614ef560208401614daf565b90509250929050565b600080600060608486031215614f12578081fd5b614f1b84614daf565b9250614f2960208501614daf565b91506040840135614f3981615a19565b809150509250925092565b600080600080600080600060e0888a031215614f5e578283fd5b614f6788614daf565b9650602088013567ffffffffffffffff80821115614f83578485fd5b614f8f8b838c01614e39565b975060408a0135915080821115614fa4578485fd5b614fb08b838c01614e39565b965060608a0135915080821115614fc5578485fd5b614fd18b838c01614e39565b955060808a0135915080821115614fe6578485fd5b614ff28b838c01614e39565b945060a08a0135915080821115615007578384fd5b6150138b838c01614e39565b935060c08a0135915080821115615028578283fd5b506150358a828b01614e39565b91505092959891949750929550565b60008060408385031215615056578182fd5b61505f83614daf565b9150602083013561506f81615a19565b809150509250929050565b6000806040838503121561508c578182fd5b61509583614daf565b9150602083013567ffffffffffffffff8111156150b0578182fd5b8301601f810185136150c0578182fd5b6150cf85823560208401614d57565b9150509250929050565b600080604083850312156150eb578182fd5b6150f483614daf565b946020939093013593505050565b600080600060608486031215615116578081fd5b61511f84614daf565b92506020840135915061513460408501614daf565b90509250925092565b60008060008060808587031215615152578182fd5b61515b85614daf565b93506020850135925061517060408601614daf565b9396929550929360600135925050565b600080600060608486031215615194578081fd5b61519d84614daf565b95602085013595506040909401359392505050565b600080604083850312156151c4578182fd5b823567ffffffffffffffff808211156151db578384fd5b6151e786838701614dc6565b935060208501359150808211156151fc578283fd5b506150cf85828601614dc6565b60006020828403121561521a578081fd5b8151611a1781615a19565b600060208284031215615236578081fd5b5035919050565b6000806040838503121561524f578182fd5b82359150614ef560208401614daf565b600060208284031215615270578081fd5b81356001600160e01b031981168114611a17578182fd5b600060208284031215615298578081fd5b813567ffffffffffffffff8111156152ae578182fd5b614bee84828501614e93565b6000808284036101008112156152ce578283fd5b60e08112156152db578283fd5b506152e660e0615782565b833581526152f660208501614daf565b602082015261530760408501614daf565b60408201526060840135606082015261532260808501614daf565b608082015261533360a08501614daf565b60a082015260c08481013590820152915060e083013567ffffffffffffffff81111561535d578182fd5b6150cf85828601614e93565b60008082840360e081121561537c578283fd5b60c0811215615389578283fd5b5061539460c0615782565b833581526153a460208501614daf565b60208201526153b560408501614daf565b6040820152606084013560608201526153d060808501614daf565b608082015260a08481013590820152915060c083013567ffffffffffffffff81111561535d578182fd5b60006020828403121561540b578081fd5b5051919050565b600080600060608486031215615426578081fd5b8335925061543660208501614daf565b915061513460408501614daf565b60008060408385031215615456578182fd5b50508035926020909101359150565b600060208284031215615476578081fd5b815160ff81168114611a17578182fd5b6000825161549881846020870161598f565b9190910192915050565b60007f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000825283516154da81601785016020880161598f565b7001034b99036b4b9b9b4b733903937b6329607d1b601791840191820152835161550b81602884016020880161598f565b01602801949350505050565b60008782526bffffffffffffffffffffffff19808860601b166020840152808760601b166034840152808660601b1660488401525083605c830152825161556581607c85016020870161598f565b91909101607c01979650505050505050565b6001600160a01b03929092168252602082015260400190565b60006020825282518060208401526155af81604085016020870161598f565b601f01601f19169190910160400192915050565b6020808252601d908201527f54776f5761793a2063616c6c6572206973206e6f742063726f73736572000000604082015260600190565b6020808252600b908201526a1b9bdd081cdd5c1c1bdc9d60aa1b604082015260600190565b602080825260119082015270151ddbd5d85e4e881d1e081b5a5b9d1959607a1b604082015260600190565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201526d191e481a5b9a5d1a585b1a5e995960921b606082015260800190565b6020808252601b908201527f54776f5761793a2063616c6c6572206973206e6f742061646d696e0000000000604082015260600190565b6020808252600c908201526b7a65726f206164647265737360a01b604082015260600190565b602080825260139082015272040e4c2e8d2de8ccaca40dcdee840dac2e8c6d606b1b604082015260600190565b600060e08201905082518252602083015160018060a01b038082166020850152806040860151166040850152606085015160608501528060808601511660808501528060a08601511660a0850152505060c083015160c083015292915050565b604051601f8201601f1916810167ffffffffffffffff811182821017156157ab576157ab615a03565b604052919050565b600067ffffffffffffffff8211156157cd576157cd615a03565b5060209081020190565b600082198211156157ea576157ea6159ed565b500190565b60008261580a57634e487b7160e01b81526012600452602481fd5b500490565b80825b6001808611615821575061584c565b818704821115615833576158336159ed565b8086161561584057918102915b9490941c938002615812565b94509492505050565b6000611a176000198484615872565b6000611a1760001960ff8516845b60008261588157506001611a17565b8161588e57506000611a17565b81600181146158a457600281146158ae576158db565b6001915050611a17565b60ff8411156158bf576158bf6159ed565b6001841b9150848211156158d5576158d56159ed565b50611a17565b5060208310610133831016604e8410600b841016171561590e575081810a83811115615909576159096159ed565b611a17565b61591b848484600161580f565b80860482111561592d5761592d6159ed565b02949350505050565b6000816000190483118215151615615950576159506159ed565b500290565b600082821015615967576159676159ed565b500390565b600060ff821660ff841680821015615986576159866159ed565b90039392505050565b60005b838110156159aa578181015183820152602001615992565b838111156135b15750506000910152565b6000816159ca576159ca6159ed565b506000190190565b60006000198214156159e6576159e66159ed565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b8015158114611aec57600080fdfe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a26469706673582212202f2d69b51cef1ef26d6d1e4050de824df0f7cc8aeda208e02b9f5e4f24d4946864736f6c63430008020033",
}

// TwoWayCenterABI is the input ABI used to generate the binding from.
// Deprecated: Use TwoWayCenterMetaData.ABI instead.
var TwoWayCenterABI = TwoWayCenterMetaData.ABI

// Deprecated: Use TwoWayCenterMetaData.Sigs instead.
// TwoWayCenterFuncSigs maps the 4-byte function signature to its string representation.
var TwoWayCenterFuncSigs = TwoWayCenterMetaData.Sigs

// TwoWayCenterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TwoWayCenterMetaData.Bin instead.
var TwoWayCenterBin = TwoWayCenterMetaData.Bin

// DeployTwoWayCenter deploys a new Ethereum contract, binding an instance of TwoWayCenter to it.
func DeployTwoWayCenter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TwoWayCenter, error) {
	parsed, err := TwoWayCenterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TwoWayCenterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TwoWayCenter{TwoWayCenterCaller: TwoWayCenterCaller{contract: contract}, TwoWayCenterTransactor: TwoWayCenterTransactor{contract: contract}, TwoWayCenterFilterer: TwoWayCenterFilterer{contract: contract}}, nil
}

// TwoWayCenter is an auto generated Go binding around an Ethereum contract.
type TwoWayCenter struct {
	TwoWayCenterCaller     // Read-only binding to the contract
	TwoWayCenterTransactor // Write-only binding to the contract
	TwoWayCenterFilterer   // Log filterer for contract events
}

// TwoWayCenterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TwoWayCenterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoWayCenterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TwoWayCenterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoWayCenterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TwoWayCenterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoWayCenterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TwoWayCenterSession struct {
	Contract     *TwoWayCenter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TwoWayCenterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TwoWayCenterCallerSession struct {
	Contract *TwoWayCenterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TwoWayCenterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TwoWayCenterTransactorSession struct {
	Contract     *TwoWayCenterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TwoWayCenterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TwoWayCenterRaw struct {
	Contract *TwoWayCenter // Generic contract binding to access the raw methods on
}

// TwoWayCenterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TwoWayCenterCallerRaw struct {
	Contract *TwoWayCenterCaller // Generic read-only contract binding to access the raw methods on
}

// TwoWayCenterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TwoWayCenterTransactorRaw struct {
	Contract *TwoWayCenterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTwoWayCenter creates a new instance of TwoWayCenter, bound to a specific deployed contract.
func NewTwoWayCenter(address common.Address, backend bind.ContractBackend) (*TwoWayCenter, error) {
	contract, err := bindTwoWayCenter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TwoWayCenter{TwoWayCenterCaller: TwoWayCenterCaller{contract: contract}, TwoWayCenterTransactor: TwoWayCenterTransactor{contract: contract}, TwoWayCenterFilterer: TwoWayCenterFilterer{contract: contract}}, nil
}

// NewTwoWayCenterCaller creates a new read-only instance of TwoWayCenter, bound to a specific deployed contract.
func NewTwoWayCenterCaller(address common.Address, caller bind.ContractCaller) (*TwoWayCenterCaller, error) {
	contract, err := bindTwoWayCenter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterCaller{contract: contract}, nil
}

// NewTwoWayCenterTransactor creates a new write-only instance of TwoWayCenter, bound to a specific deployed contract.
func NewTwoWayCenterTransactor(address common.Address, transactor bind.ContractTransactor) (*TwoWayCenterTransactor, error) {
	contract, err := bindTwoWayCenter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterTransactor{contract: contract}, nil
}

// NewTwoWayCenterFilterer creates a new log filterer instance of TwoWayCenter, bound to a specific deployed contract.
func NewTwoWayCenterFilterer(address common.Address, filterer bind.ContractFilterer) (*TwoWayCenterFilterer, error) {
	contract, err := bindTwoWayCenter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterFilterer{contract: contract}, nil
}

// bindTwoWayCenter binds a generic wrapper to an already deployed contract.
func bindTwoWayCenter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TwoWayCenterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwoWayCenter *TwoWayCenterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwoWayCenter.Contract.TwoWayCenterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwoWayCenter *TwoWayCenterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.TwoWayCenterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwoWayCenter *TwoWayCenterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.TwoWayCenterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwoWayCenter *TwoWayCenterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwoWayCenter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwoWayCenter *TwoWayCenterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwoWayCenter *TwoWayCenterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TwoWayCenter *TwoWayCenterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TwoWayCenter *TwoWayCenterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TwoWayCenter.Contract.DEFAULTADMINROLE(&_TwoWayCenter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TwoWayCenter *TwoWayCenterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TwoWayCenter.Contract.DEFAULTADMINROLE(&_TwoWayCenter.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_TwoWayCenter *TwoWayCenterCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_TwoWayCenter *TwoWayCenterSession) UPGRADERROLE() ([32]byte, error) {
	return _TwoWayCenter.Contract.UPGRADERROLE(&_TwoWayCenter.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_TwoWayCenter *TwoWayCenterCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _TwoWayCenter.Contract.UPGRADERROLE(&_TwoWayCenter.CallOpts)
}

// CalculateFee is a free data retrieval call binding the contract method 0x0324ef9c.
//
// Solidity: function calculateFee(address oToken, uint256 toChainId, uint256 amount) view returns(uint256 fixAmount, uint256 ratioAmount, uint256 remainAmount)
func (_TwoWayCenter *TwoWayCenterCaller) CalculateFee(opts *bind.CallOpts, oToken common.Address, toChainId *big.Int, amount *big.Int) (struct {
	FixAmount    *big.Int
	RatioAmount  *big.Int
	RemainAmount *big.Int
}, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "calculateFee", oToken, toChainId, amount)

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

// CalculateFee is a free data retrieval call binding the contract method 0x0324ef9c.
//
// Solidity: function calculateFee(address oToken, uint256 toChainId, uint256 amount) view returns(uint256 fixAmount, uint256 ratioAmount, uint256 remainAmount)
func (_TwoWayCenter *TwoWayCenterSession) CalculateFee(oToken common.Address, toChainId *big.Int, amount *big.Int) (struct {
	FixAmount    *big.Int
	RatioAmount  *big.Int
	RemainAmount *big.Int
}, error) {
	return _TwoWayCenter.Contract.CalculateFee(&_TwoWayCenter.CallOpts, oToken, toChainId, amount)
}

// CalculateFee is a free data retrieval call binding the contract method 0x0324ef9c.
//
// Solidity: function calculateFee(address oToken, uint256 toChainId, uint256 amount) view returns(uint256 fixAmount, uint256 ratioAmount, uint256 remainAmount)
func (_TwoWayCenter *TwoWayCenterCallerSession) CalculateFee(oToken common.Address, toChainId *big.Int, amount *big.Int) (struct {
	FixAmount    *big.Int
	RatioAmount  *big.Int
	RemainAmount *big.Int
}, error) {
	return _TwoWayCenter.Contract.CalculateFee(&_TwoWayCenter.CallOpts, oToken, toChainId, amount)
}

// CalculateReward is a free data retrieval call binding the contract method 0x2c32c360.
//
// Solidity: function calculateReward(address oToken, uint256 fromChainId, uint256 amount) view returns(uint256 _reward)
func (_TwoWayCenter *TwoWayCenterCaller) CalculateReward(opts *bind.CallOpts, oToken common.Address, fromChainId *big.Int, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "calculateReward", oToken, fromChainId, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateReward is a free data retrieval call binding the contract method 0x2c32c360.
//
// Solidity: function calculateReward(address oToken, uint256 fromChainId, uint256 amount) view returns(uint256 _reward)
func (_TwoWayCenter *TwoWayCenterSession) CalculateReward(oToken common.Address, fromChainId *big.Int, amount *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.CalculateReward(&_TwoWayCenter.CallOpts, oToken, fromChainId, amount)
}

// CalculateReward is a free data retrieval call binding the contract method 0x2c32c360.
//
// Solidity: function calculateReward(address oToken, uint256 fromChainId, uint256 amount) view returns(uint256 _reward)
func (_TwoWayCenter *TwoWayCenterCallerSession) CalculateReward(oToken common.Address, fromChainId *big.Int, amount *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.CalculateReward(&_TwoWayCenter.CallOpts, oToken, fromChainId, amount)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) ChainId() (*big.Int, error) {
	return _TwoWayCenter.Contract.ChainId(&_TwoWayCenter.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) ChainId() (*big.Int, error) {
	return _TwoWayCenter.Contract.ChainId(&_TwoWayCenter.CallOpts)
}

// DecimalDiff is a free data retrieval call binding the contract method 0xbf814720.
//
// Solidity: function decimalDiff(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) DecimalDiff(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "decimalDiff", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecimalDiff is a free data retrieval call binding the contract method 0xbf814720.
//
// Solidity: function decimalDiff(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) DecimalDiff(arg0 common.Address) (*big.Int, error) {
	return _TwoWayCenter.Contract.DecimalDiff(&_TwoWayCenter.CallOpts, arg0)
}

// DecimalDiff is a free data retrieval call binding the contract method 0xbf814720.
//
// Solidity: function decimalDiff(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) DecimalDiff(arg0 common.Address) (*big.Int, error) {
	return _TwoWayCenter.Contract.DecimalDiff(&_TwoWayCenter.CallOpts, arg0)
}

// EventHeights0 is a free data retrieval call binding the contract method 0xfc1aa909.
//
// Solidity: function eventHeights0(uint256 , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) EventHeights0(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "eventHeights0", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventHeights0 is a free data retrieval call binding the contract method 0xfc1aa909.
//
// Solidity: function eventHeights0(uint256 , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) EventHeights0(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.EventHeights0(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// EventHeights0 is a free data retrieval call binding the contract method 0xfc1aa909.
//
// Solidity: function eventHeights0(uint256 , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) EventHeights0(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.EventHeights0(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// EventHeights1 is a free data retrieval call binding the contract method 0x05e82506.
//
// Solidity: function eventHeights1(uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) EventHeights1(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "eventHeights1", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventHeights1 is a free data retrieval call binding the contract method 0x05e82506.
//
// Solidity: function eventHeights1(uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) EventHeights1(arg0 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.EventHeights1(&_TwoWayCenter.CallOpts, arg0)
}

// EventHeights1 is a free data retrieval call binding the contract method 0x05e82506.
//
// Solidity: function eventHeights1(uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) EventHeights1(arg0 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.EventHeights1(&_TwoWayCenter.CallOpts, arg0)
}

// EventIndex0 is a free data retrieval call binding the contract method 0x0826820e.
//
// Solidity: function eventIndex0(uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) EventIndex0(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "eventIndex0", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventIndex0 is a free data retrieval call binding the contract method 0x0826820e.
//
// Solidity: function eventIndex0(uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) EventIndex0(arg0 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.EventIndex0(&_TwoWayCenter.CallOpts, arg0)
}

// EventIndex0 is a free data retrieval call binding the contract method 0x0826820e.
//
// Solidity: function eventIndex0(uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) EventIndex0(arg0 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.EventIndex0(&_TwoWayCenter.CallOpts, arg0)
}

// EventIndex1 is a free data retrieval call binding the contract method 0x15cd2c70.
//
// Solidity: function eventIndex1() view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) EventIndex1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "eventIndex1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EventIndex1 is a free data retrieval call binding the contract method 0x15cd2c70.
//
// Solidity: function eventIndex1() view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) EventIndex1() (*big.Int, error) {
	return _TwoWayCenter.Contract.EventIndex1(&_TwoWayCenter.CallOpts)
}

// EventIndex1 is a free data retrieval call binding the contract method 0x15cd2c70.
//
// Solidity: function eventIndex1() view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) EventIndex1() (*big.Int, error) {
	return _TwoWayCenter.Contract.EventIndex1(&_TwoWayCenter.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_TwoWayCenter *TwoWayCenterCaller) FeeTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "feeTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_TwoWayCenter *TwoWayCenterSession) FeeTo() (common.Address, error) {
	return _TwoWayCenter.Contract.FeeTo(&_TwoWayCenter.CallOpts)
}

// FeeTo is a free data retrieval call binding the contract method 0x017e7e58.
//
// Solidity: function feeTo() view returns(address)
func (_TwoWayCenter *TwoWayCenterCallerSession) FeeTo() (common.Address, error) {
	return _TwoWayCenter.Contract.FeeTo(&_TwoWayCenter.CallOpts)
}

// FeeToTreasuryRatio is a free data retrieval call binding the contract method 0x59ded2d7.
//
// Solidity: function feeToTreasuryRatio(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) FeeToTreasuryRatio(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "feeToTreasuryRatio", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeToTreasuryRatio is a free data retrieval call binding the contract method 0x59ded2d7.
//
// Solidity: function feeToTreasuryRatio(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) FeeToTreasuryRatio(arg0 common.Address) (*big.Int, error) {
	return _TwoWayCenter.Contract.FeeToTreasuryRatio(&_TwoWayCenter.CallOpts, arg0)
}

// FeeToTreasuryRatio is a free data retrieval call binding the contract method 0x59ded2d7.
//
// Solidity: function feeToTreasuryRatio(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) FeeToTreasuryRatio(arg0 common.Address) (*big.Int, error) {
	return _TwoWayCenter.Contract.FeeToTreasuryRatio(&_TwoWayCenter.CallOpts, arg0)
}

// FixFees is a free data retrieval call binding the contract method 0x75d47f99.
//
// Solidity: function fixFees(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) FixFees(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "fixFees", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FixFees is a free data retrieval call binding the contract method 0x75d47f99.
//
// Solidity: function fixFees(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) FixFees(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.FixFees(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// FixFees is a free data retrieval call binding the contract method 0x75d47f99.
//
// Solidity: function fixFees(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) FixFees(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.FixFees(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// GetMsgSender is a free data retrieval call binding the contract method 0x7a6ce2e1.
//
// Solidity: function getMsgSender() view returns(address, address)
func (_TwoWayCenter *TwoWayCenterCaller) GetMsgSender(opts *bind.CallOpts) (common.Address, common.Address, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "getMsgSender")

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
func (_TwoWayCenter *TwoWayCenterSession) GetMsgSender() (common.Address, common.Address, error) {
	return _TwoWayCenter.Contract.GetMsgSender(&_TwoWayCenter.CallOpts)
}

// GetMsgSender is a free data retrieval call binding the contract method 0x7a6ce2e1.
//
// Solidity: function getMsgSender() view returns(address, address)
func (_TwoWayCenter *TwoWayCenterCallerSession) GetMsgSender() (common.Address, common.Address, error) {
	return _TwoWayCenter.Contract.GetMsgSender(&_TwoWayCenter.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TwoWayCenter *TwoWayCenterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TwoWayCenter *TwoWayCenterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TwoWayCenter.Contract.GetRoleAdmin(&_TwoWayCenter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TwoWayCenter *TwoWayCenterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TwoWayCenter.Contract.GetRoleAdmin(&_TwoWayCenter.CallOpts, role)
}

// GetRoleKey is a free data retrieval call binding the contract method 0x3ab796e7.
//
// Solidity: function getRoleKey(address toToken) pure returns(bytes32 key)
func (_TwoWayCenter *TwoWayCenterCaller) GetRoleKey(opts *bind.CallOpts, toToken common.Address) ([32]byte, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "getRoleKey", toToken)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleKey is a free data retrieval call binding the contract method 0x3ab796e7.
//
// Solidity: function getRoleKey(address toToken) pure returns(bytes32 key)
func (_TwoWayCenter *TwoWayCenterSession) GetRoleKey(toToken common.Address) ([32]byte, error) {
	return _TwoWayCenter.Contract.GetRoleKey(&_TwoWayCenter.CallOpts, toToken)
}

// GetRoleKey is a free data retrieval call binding the contract method 0x3ab796e7.
//
// Solidity: function getRoleKey(address toToken) pure returns(bytes32 key)
func (_TwoWayCenter *TwoWayCenterCallerSession) GetRoleKey(toToken common.Address) ([32]byte, error) {
	return _TwoWayCenter.Contract.GetRoleKey(&_TwoWayCenter.CallOpts, toToken)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TwoWayCenter *TwoWayCenterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TwoWayCenter *TwoWayCenterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TwoWayCenter.Contract.HasRole(&_TwoWayCenter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TwoWayCenter *TwoWayCenterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TwoWayCenter.Contract.HasRole(&_TwoWayCenter.CallOpts, role, account)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() view returns(bool)
func (_TwoWayCenter *TwoWayCenterCaller) IsClosed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "isClosed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() view returns(bool)
func (_TwoWayCenter *TwoWayCenterSession) IsClosed() (bool, error) {
	return _TwoWayCenter.Contract.IsClosed(&_TwoWayCenter.CallOpts)
}

// IsClosed is a free data retrieval call binding the contract method 0xc2b6b58c.
//
// Solidity: function isClosed() view returns(bool)
func (_TwoWayCenter *TwoWayCenterCallerSession) IsClosed() (bool, error) {
	return _TwoWayCenter.Contract.IsClosed(&_TwoWayCenter.CallOpts)
}

// IsCoin is a free data retrieval call binding the contract method 0xa611033e.
//
// Solidity: function isCoin(address ) view returns(bool)
func (_TwoWayCenter *TwoWayCenterCaller) IsCoin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "isCoin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCoin is a free data retrieval call binding the contract method 0xa611033e.
//
// Solidity: function isCoin(address ) view returns(bool)
func (_TwoWayCenter *TwoWayCenterSession) IsCoin(arg0 common.Address) (bool, error) {
	return _TwoWayCenter.Contract.IsCoin(&_TwoWayCenter.CallOpts, arg0)
}

// IsCoin is a free data retrieval call binding the contract method 0xa611033e.
//
// Solidity: function isCoin(address ) view returns(bool)
func (_TwoWayCenter *TwoWayCenterCallerSession) IsCoin(arg0 common.Address) (bool, error) {
	return _TwoWayCenter.Contract.IsCoin(&_TwoWayCenter.CallOpts, arg0)
}

// IsInWhilelist is a free data retrieval call binding the contract method 0xa07b32c0.
//
// Solidity: function isInWhilelist(address , address ) view returns(bool)
func (_TwoWayCenter *TwoWayCenterCaller) IsInWhilelist(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "isInWhilelist", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInWhilelist is a free data retrieval call binding the contract method 0xa07b32c0.
//
// Solidity: function isInWhilelist(address , address ) view returns(bool)
func (_TwoWayCenter *TwoWayCenterSession) IsInWhilelist(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _TwoWayCenter.Contract.IsInWhilelist(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// IsInWhilelist is a free data retrieval call binding the contract method 0xa07b32c0.
//
// Solidity: function isInWhilelist(address , address ) view returns(bool)
func (_TwoWayCenter *TwoWayCenterCallerSession) IsInWhilelist(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _TwoWayCenter.Contract.IsInWhilelist(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// LockBalances is a free data retrieval call binding the contract method 0xa6ba8fc1.
//
// Solidity: function lockBalances(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) LockBalances(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "lockBalances", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockBalances is a free data retrieval call binding the contract method 0xa6ba8fc1.
//
// Solidity: function lockBalances(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) LockBalances(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.LockBalances(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// LockBalances is a free data retrieval call binding the contract method 0xa6ba8fc1.
//
// Solidity: function lockBalances(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) LockBalances(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.LockBalances(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// RatioFeesHigh is a free data retrieval call binding the contract method 0xee5f7d96.
//
// Solidity: function ratioFeesHigh(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) RatioFeesHigh(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "ratioFeesHigh", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RatioFeesHigh is a free data retrieval call binding the contract method 0xee5f7d96.
//
// Solidity: function ratioFeesHigh(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) RatioFeesHigh(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.RatioFeesHigh(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// RatioFeesHigh is a free data retrieval call binding the contract method 0xee5f7d96.
//
// Solidity: function ratioFeesHigh(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) RatioFeesHigh(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.RatioFeesHigh(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// RatioFeesLow is a free data retrieval call binding the contract method 0xe4732500.
//
// Solidity: function ratioFeesLow(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) RatioFeesLow(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "ratioFeesLow", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RatioFeesLow is a free data retrieval call binding the contract method 0xe4732500.
//
// Solidity: function ratioFeesLow(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) RatioFeesLow(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.RatioFeesLow(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// RatioFeesLow is a free data retrieval call binding the contract method 0xe4732500.
//
// Solidity: function ratioFeesLow(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) RatioFeesLow(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.RatioFeesLow(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// RatioFeesMedium is a free data retrieval call binding the contract method 0x3ff9bb66.
//
// Solidity: function ratioFeesMedium(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) RatioFeesMedium(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "ratioFeesMedium", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RatioFeesMedium is a free data retrieval call binding the contract method 0x3ff9bb66.
//
// Solidity: function ratioFeesMedium(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) RatioFeesMedium(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.RatioFeesMedium(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// RatioFeesMedium is a free data retrieval call binding the contract method 0x3ff9bb66.
//
// Solidity: function ratioFeesMedium(address , uint256 ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) RatioFeesMedium(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayCenter.Contract.RatioFeesMedium(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// RemainHigh is a free data retrieval call binding the contract method 0x260ca4d5.
//
// Solidity: function remainHigh(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) RemainHigh(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "remainHigh", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainHigh is a free data retrieval call binding the contract method 0x260ca4d5.
//
// Solidity: function remainHigh(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) RemainHigh(arg0 common.Address) (*big.Int, error) {
	return _TwoWayCenter.Contract.RemainHigh(&_TwoWayCenter.CallOpts, arg0)
}

// RemainHigh is a free data retrieval call binding the contract method 0x260ca4d5.
//
// Solidity: function remainHigh(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) RemainHigh(arg0 common.Address) (*big.Int, error) {
	return _TwoWayCenter.Contract.RemainHigh(&_TwoWayCenter.CallOpts, arg0)
}

// RemainLow is a free data retrieval call binding the contract method 0x4297171f.
//
// Solidity: function remainLow(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) RemainLow(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "remainLow", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainLow is a free data retrieval call binding the contract method 0x4297171f.
//
// Solidity: function remainLow(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) RemainLow(arg0 common.Address) (*big.Int, error) {
	return _TwoWayCenter.Contract.RemainLow(&_TwoWayCenter.CallOpts, arg0)
}

// RemainLow is a free data retrieval call binding the contract method 0x4297171f.
//
// Solidity: function remainLow(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) RemainLow(arg0 common.Address) (*big.Int, error) {
	return _TwoWayCenter.Contract.RemainLow(&_TwoWayCenter.CallOpts, arg0)
}

// RewardRatio is a free data retrieval call binding the contract method 0xced2a7ba.
//
// Solidity: function rewardRatio(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) RewardRatio(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "rewardRatio", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardRatio is a free data retrieval call binding the contract method 0xced2a7ba.
//
// Solidity: function rewardRatio(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) RewardRatio(arg0 common.Address) (*big.Int, error) {
	return _TwoWayCenter.Contract.RewardRatio(&_TwoWayCenter.CallOpts, arg0)
}

// RewardRatio is a free data retrieval call binding the contract method 0xced2a7ba.
//
// Solidity: function rewardRatio(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) RewardRatio(arg0 common.Address) (*big.Int, error) {
	return _TwoWayCenter.Contract.RewardRatio(&_TwoWayCenter.CallOpts, arg0)
}

// Sr is a free data retrieval call binding the contract method 0xdf791d80.
//
// Solidity: function sr() view returns(address)
func (_TwoWayCenter *TwoWayCenterCaller) Sr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "sr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Sr is a free data retrieval call binding the contract method 0xdf791d80.
//
// Solidity: function sr() view returns(address)
func (_TwoWayCenter *TwoWayCenterSession) Sr() (common.Address, error) {
	return _TwoWayCenter.Contract.Sr(&_TwoWayCenter.CallOpts)
}

// Sr is a free data retrieval call binding the contract method 0xdf791d80.
//
// Solidity: function sr() view returns(address)
func (_TwoWayCenter *TwoWayCenterCallerSession) Sr() (common.Address, error) {
	return _TwoWayCenter.Contract.Sr(&_TwoWayCenter.CallOpts)
}

// Srs is a free data retrieval call binding the contract method 0xf3669d80.
//
// Solidity: function srs(address ) view returns(address)
func (_TwoWayCenter *TwoWayCenterCaller) Srs(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "srs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Srs is a free data retrieval call binding the contract method 0xf3669d80.
//
// Solidity: function srs(address ) view returns(address)
func (_TwoWayCenter *TwoWayCenterSession) Srs(arg0 common.Address) (common.Address, error) {
	return _TwoWayCenter.Contract.Srs(&_TwoWayCenter.CallOpts, arg0)
}

// Srs is a free data retrieval call binding the contract method 0xf3669d80.
//
// Solidity: function srs(address ) view returns(address)
func (_TwoWayCenter *TwoWayCenterCallerSession) Srs(arg0 common.Address) (common.Address, error) {
	return _TwoWayCenter.Contract.Srs(&_TwoWayCenter.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TwoWayCenter *TwoWayCenterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TwoWayCenter *TwoWayCenterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TwoWayCenter.Contract.SupportsInterface(&_TwoWayCenter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TwoWayCenter *TwoWayCenterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TwoWayCenter.Contract.SupportsInterface(&_TwoWayCenter.CallOpts, interfaceId)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCaller) Threshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "threshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _TwoWayCenter.Contract.Threshold(&_TwoWayCenter.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_TwoWayCenter *TwoWayCenterCallerSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _TwoWayCenter.Contract.Threshold(&_TwoWayCenter.CallOpts, arg0)
}

// ToCenterToken is a free data retrieval call binding the contract method 0xbafd2d19.
//
// Solidity: function toCenterToken(uint256 , address ) view returns(address)
func (_TwoWayCenter *TwoWayCenterCaller) ToCenterToken(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "toCenterToken", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ToCenterToken is a free data retrieval call binding the contract method 0xbafd2d19.
//
// Solidity: function toCenterToken(uint256 , address ) view returns(address)
func (_TwoWayCenter *TwoWayCenterSession) ToCenterToken(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _TwoWayCenter.Contract.ToCenterToken(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// ToCenterToken is a free data retrieval call binding the contract method 0xbafd2d19.
//
// Solidity: function toCenterToken(uint256 , address ) view returns(address)
func (_TwoWayCenter *TwoWayCenterCallerSession) ToCenterToken(arg0 *big.Int, arg1 common.Address) (common.Address, error) {
	return _TwoWayCenter.Contract.ToCenterToken(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// ToEdgeToken is a free data retrieval call binding the contract method 0x527b0f64.
//
// Solidity: function toEdgeToken(address , uint256 ) view returns(address)
func (_TwoWayCenter *TwoWayCenterCaller) ToEdgeToken(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "toEdgeToken", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ToEdgeToken is a free data retrieval call binding the contract method 0x527b0f64.
//
// Solidity: function toEdgeToken(address , uint256 ) view returns(address)
func (_TwoWayCenter *TwoWayCenterSession) ToEdgeToken(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _TwoWayCenter.Contract.ToEdgeToken(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// ToEdgeToken is a free data retrieval call binding the contract method 0x527b0f64.
//
// Solidity: function toEdgeToken(address , uint256 ) view returns(address)
func (_TwoWayCenter *TwoWayCenterCallerSession) ToEdgeToken(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _TwoWayCenter.Contract.ToEdgeToken(&_TwoWayCenter.CallOpts, arg0, arg1)
}

// TreasuryTo is a free data retrieval call binding the contract method 0x42079671.
//
// Solidity: function treasuryTo() view returns(address)
func (_TwoWayCenter *TwoWayCenterCaller) TreasuryTo(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "treasuryTo")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TreasuryTo is a free data retrieval call binding the contract method 0x42079671.
//
// Solidity: function treasuryTo() view returns(address)
func (_TwoWayCenter *TwoWayCenterSession) TreasuryTo() (common.Address, error) {
	return _TwoWayCenter.Contract.TreasuryTo(&_TwoWayCenter.CallOpts)
}

// TreasuryTo is a free data retrieval call binding the contract method 0x42079671.
//
// Solidity: function treasuryTo() view returns(address)
func (_TwoWayCenter *TwoWayCenterCallerSession) TreasuryTo() (common.Address, error) {
	return _TwoWayCenter.Contract.TreasuryTo(&_TwoWayCenter.CallOpts)
}

// TxHandled is a free data retrieval call binding the contract method 0x824df449.
//
// Solidity: function txHandled(string ) view returns(bool)
func (_TwoWayCenter *TwoWayCenterCaller) TxHandled(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _TwoWayCenter.contract.Call(opts, &out, "txHandled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TxHandled is a free data retrieval call binding the contract method 0x824df449.
//
// Solidity: function txHandled(string ) view returns(bool)
func (_TwoWayCenter *TwoWayCenterSession) TxHandled(arg0 string) (bool, error) {
	return _TwoWayCenter.Contract.TxHandled(&_TwoWayCenter.CallOpts, arg0)
}

// TxHandled is a free data retrieval call binding the contract method 0x824df449.
//
// Solidity: function txHandled(string ) view returns(bool)
func (_TwoWayCenter *TwoWayCenterCallerSession) TxHandled(arg0 string) (bool, error) {
	return _TwoWayCenter.Contract.TxHandled(&_TwoWayCenter.CallOpts, arg0)
}

// AddToken is a paid mutator transaction binding the contract method 0xa766a392.
//
// Solidity: function addToken(address _centerToken, uint256 _edgeChainId, address _edgeToken) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) AddToken(opts *bind.TransactOpts, _centerToken common.Address, _edgeChainId *big.Int, _edgeToken common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "addToken", _centerToken, _edgeChainId, _edgeToken)
}

// AddToken is a paid mutator transaction binding the contract method 0xa766a392.
//
// Solidity: function addToken(address _centerToken, uint256 _edgeChainId, address _edgeToken) returns()
func (_TwoWayCenter *TwoWayCenterSession) AddToken(_centerToken common.Address, _edgeChainId *big.Int, _edgeToken common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.AddToken(&_TwoWayCenter.TransactOpts, _centerToken, _edgeChainId, _edgeToken)
}

// AddToken is a paid mutator transaction binding the contract method 0xa766a392.
//
// Solidity: function addToken(address _centerToken, uint256 _edgeChainId, address _edgeToken) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) AddToken(_centerToken common.Address, _edgeChainId *big.Int, _edgeToken common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.AddToken(&_TwoWayCenter.TransactOpts, _centerToken, _edgeChainId, _edgeToken)
}

// CrossIn is a paid mutator transaction binding the contract method 0x1ca97185.
//
// Solidity: function crossIn((uint256,address,address,uint256,address,uint256) p, string txid) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) CrossIn(opts *bind.TransactOpts, p OutParam, txid string) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "crossIn", p, txid)
}

// CrossIn is a paid mutator transaction binding the contract method 0x1ca97185.
//
// Solidity: function crossIn((uint256,address,address,uint256,address,uint256) p, string txid) returns()
func (_TwoWayCenter *TwoWayCenterSession) CrossIn(p OutParam, txid string) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.CrossIn(&_TwoWayCenter.TransactOpts, p, txid)
}

// CrossIn is a paid mutator transaction binding the contract method 0x1ca97185.
//
// Solidity: function crossIn((uint256,address,address,uint256,address,uint256) p, string txid) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) CrossIn(p OutParam, txid string) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.CrossIn(&_TwoWayCenter.TransactOpts, p, txid)
}

// CrossOut is a paid mutator transaction binding the contract method 0xdf7e600a.
//
// Solidity: function crossOut(address fromToken, uint256 toChainId, address to, uint256 amount) payable returns()
func (_TwoWayCenter *TwoWayCenterTransactor) CrossOut(opts *bind.TransactOpts, fromToken common.Address, toChainId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "crossOut", fromToken, toChainId, to, amount)
}

// CrossOut is a paid mutator transaction binding the contract method 0xdf7e600a.
//
// Solidity: function crossOut(address fromToken, uint256 toChainId, address to, uint256 amount) payable returns()
func (_TwoWayCenter *TwoWayCenterSession) CrossOut(fromToken common.Address, toChainId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.CrossOut(&_TwoWayCenter.TransactOpts, fromToken, toChainId, to, amount)
}

// CrossOut is a paid mutator transaction binding the contract method 0xdf7e600a.
//
// Solidity: function crossOut(address fromToken, uint256 toChainId, address to, uint256 amount) payable returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) CrossOut(fromToken common.Address, toChainId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.CrossOut(&_TwoWayCenter.TransactOpts, fromToken, toChainId, to, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_TwoWayCenter *TwoWayCenterTransactor) Deposit(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "deposit", token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_TwoWayCenter *TwoWayCenterSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.Deposit(&_TwoWayCenter.TransactOpts, token, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address token, uint256 amount) payable returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) Deposit(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.Deposit(&_TwoWayCenter.TransactOpts, token, amount)
}

// ForwardCrossOut is a paid mutator transaction binding the contract method 0x4ba9b352.
//
// Solidity: function forwardCrossOut((uint256,address,address,uint256,address,uint256) p, string txid) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) ForwardCrossOut(opts *bind.TransactOpts, p OutParam, txid string) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "forwardCrossOut", p, txid)
}

// ForwardCrossOut is a paid mutator transaction binding the contract method 0x4ba9b352.
//
// Solidity: function forwardCrossOut((uint256,address,address,uint256,address,uint256) p, string txid) returns()
func (_TwoWayCenter *TwoWayCenterSession) ForwardCrossOut(p OutParam, txid string) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.ForwardCrossOut(&_TwoWayCenter.TransactOpts, p, txid)
}

// ForwardCrossOut is a paid mutator transaction binding the contract method 0x4ba9b352.
//
// Solidity: function forwardCrossOut((uint256,address,address,uint256,address,uint256) p, string txid) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) ForwardCrossOut(p OutParam, txid string) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.ForwardCrossOut(&_TwoWayCenter.TransactOpts, p, txid)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TwoWayCenter *TwoWayCenterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.GrantRole(&_TwoWayCenter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.GrantRole(&_TwoWayCenter.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xb4988fd0.
//
// Solidity: function initialize(uint256 _chainId, address _feeTo, address _treasuryTo) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) Initialize(opts *bind.TransactOpts, _chainId *big.Int, _feeTo common.Address, _treasuryTo common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "initialize", _chainId, _feeTo, _treasuryTo)
}

// Initialize is a paid mutator transaction binding the contract method 0xb4988fd0.
//
// Solidity: function initialize(uint256 _chainId, address _feeTo, address _treasuryTo) returns()
func (_TwoWayCenter *TwoWayCenterSession) Initialize(_chainId *big.Int, _feeTo common.Address, _treasuryTo common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.Initialize(&_TwoWayCenter.TransactOpts, _chainId, _feeTo, _treasuryTo)
}

// Initialize is a paid mutator transaction binding the contract method 0xb4988fd0.
//
// Solidity: function initialize(uint256 _chainId, address _feeTo, address _treasuryTo) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) Initialize(_chainId *big.Int, _feeTo common.Address, _treasuryTo common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.Initialize(&_TwoWayCenter.TransactOpts, _chainId, _feeTo, _treasuryTo)
}

// Issue is a paid mutator transaction binding the contract method 0xaf437845.
//
// Solidity: function issue((uint256,address,address,uint256,address,address,uint256) p, string txid) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) Issue(opts *bind.TransactOpts, p InParam, txid string) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "issue", p, txid)
}

// Issue is a paid mutator transaction binding the contract method 0xaf437845.
//
// Solidity: function issue((uint256,address,address,uint256,address,address,uint256) p, string txid) returns()
func (_TwoWayCenter *TwoWayCenterSession) Issue(p InParam, txid string) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.Issue(&_TwoWayCenter.TransactOpts, p, txid)
}

// Issue is a paid mutator transaction binding the contract method 0xaf437845.
//
// Solidity: function issue((uint256,address,address,uint256,address,address,uint256) p, string txid) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) Issue(p InParam, txid string) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.Issue(&_TwoWayCenter.TransactOpts, p, txid)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x97bb15c9.
//
// Solidity: function removeToken(address _centerToken, uint256 _edgeChainId, address _edgeToken) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) RemoveToken(opts *bind.TransactOpts, _centerToken common.Address, _edgeChainId *big.Int, _edgeToken common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "removeToken", _centerToken, _edgeChainId, _edgeToken)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x97bb15c9.
//
// Solidity: function removeToken(address _centerToken, uint256 _edgeChainId, address _edgeToken) returns()
func (_TwoWayCenter *TwoWayCenterSession) RemoveToken(_centerToken common.Address, _edgeChainId *big.Int, _edgeToken common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.RemoveToken(&_TwoWayCenter.TransactOpts, _centerToken, _edgeChainId, _edgeToken)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x97bb15c9.
//
// Solidity: function removeToken(address _centerToken, uint256 _edgeChainId, address _edgeToken) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) RemoveToken(_centerToken common.Address, _edgeChainId *big.Int, _edgeToken common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.RemoveToken(&_TwoWayCenter.TransactOpts, _centerToken, _edgeChainId, _edgeToken)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TwoWayCenter *TwoWayCenterSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.RenounceRole(&_TwoWayCenter.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.RenounceRole(&_TwoWayCenter.TransactOpts, role, account)
}

// ResetStakingReward is a paid mutator transaction binding the contract method 0x0f4693c6.
//
// Solidity: function resetStakingReward(address _oToken) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) ResetStakingReward(opts *bind.TransactOpts, _oToken common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "resetStakingReward", _oToken)
}

// ResetStakingReward is a paid mutator transaction binding the contract method 0x0f4693c6.
//
// Solidity: function resetStakingReward(address _oToken) returns()
func (_TwoWayCenter *TwoWayCenterSession) ResetStakingReward(_oToken common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.ResetStakingReward(&_TwoWayCenter.TransactOpts, _oToken)
}

// ResetStakingReward is a paid mutator transaction binding the contract method 0x0f4693c6.
//
// Solidity: function resetStakingReward(address _oToken) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) ResetStakingReward(_oToken common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.ResetStakingReward(&_TwoWayCenter.TransactOpts, _oToken)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TwoWayCenter *TwoWayCenterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.RevokeRole(&_TwoWayCenter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.RevokeRole(&_TwoWayCenter.TransactOpts, role, account)
}

// SetFee is a paid mutator transaction binding the contract method 0xa5886f2d.
//
// Solidity: function setFee(address token, uint256[] toChainIds, uint256[] _fixFees, uint256[] _ratioFeesHigh, uint256[] _ratioFeesMedium, uint256[] _ratioFeesLow, uint256[] _remains) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) SetFee(opts *bind.TransactOpts, token common.Address, toChainIds []*big.Int, _fixFees []*big.Int, _ratioFeesHigh []*big.Int, _ratioFeesMedium []*big.Int, _ratioFeesLow []*big.Int, _remains []*big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "setFee", token, toChainIds, _fixFees, _ratioFeesHigh, _ratioFeesMedium, _ratioFeesLow, _remains)
}

// SetFee is a paid mutator transaction binding the contract method 0xa5886f2d.
//
// Solidity: function setFee(address token, uint256[] toChainIds, uint256[] _fixFees, uint256[] _ratioFeesHigh, uint256[] _ratioFeesMedium, uint256[] _ratioFeesLow, uint256[] _remains) returns()
func (_TwoWayCenter *TwoWayCenterSession) SetFee(token common.Address, toChainIds []*big.Int, _fixFees []*big.Int, _ratioFeesHigh []*big.Int, _ratioFeesMedium []*big.Int, _ratioFeesLow []*big.Int, _remains []*big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetFee(&_TwoWayCenter.TransactOpts, token, toChainIds, _fixFees, _ratioFeesHigh, _ratioFeesMedium, _ratioFeesLow, _remains)
}

// SetFee is a paid mutator transaction binding the contract method 0xa5886f2d.
//
// Solidity: function setFee(address token, uint256[] toChainIds, uint256[] _fixFees, uint256[] _ratioFeesHigh, uint256[] _ratioFeesMedium, uint256[] _ratioFeesLow, uint256[] _remains) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) SetFee(token common.Address, toChainIds []*big.Int, _fixFees []*big.Int, _ratioFeesHigh []*big.Int, _ratioFeesMedium []*big.Int, _ratioFeesLow []*big.Int, _remains []*big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetFee(&_TwoWayCenter.TransactOpts, token, toChainIds, _fixFees, _ratioFeesHigh, _ratioFeesMedium, _ratioFeesLow, _remains)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) SetFeeTo(opts *bind.TransactOpts, _feeTo common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "setFeeTo", _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_TwoWayCenter *TwoWayCenterSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetFeeTo(&_TwoWayCenter.TransactOpts, _feeTo)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address _feeTo) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) SetFeeTo(_feeTo common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetFeeTo(&_TwoWayCenter.TransactOpts, _feeTo)
}

// SetFeeToTreasuryRatio is a paid mutator transaction binding the contract method 0xf88d0a2a.
//
// Solidity: function setFeeToTreasuryRatio(address _oToken, uint256 _ratio) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) SetFeeToTreasuryRatio(opts *bind.TransactOpts, _oToken common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "setFeeToTreasuryRatio", _oToken, _ratio)
}

// SetFeeToTreasuryRatio is a paid mutator transaction binding the contract method 0xf88d0a2a.
//
// Solidity: function setFeeToTreasuryRatio(address _oToken, uint256 _ratio) returns()
func (_TwoWayCenter *TwoWayCenterSession) SetFeeToTreasuryRatio(_oToken common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetFeeToTreasuryRatio(&_TwoWayCenter.TransactOpts, _oToken, _ratio)
}

// SetFeeToTreasuryRatio is a paid mutator transaction binding the contract method 0xf88d0a2a.
//
// Solidity: function setFeeToTreasuryRatio(address _oToken, uint256 _ratio) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) SetFeeToTreasuryRatio(_oToken common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetFeeToTreasuryRatio(&_TwoWayCenter.TransactOpts, _oToken, _ratio)
}

// SetIsCoin is a paid mutator transaction binding the contract method 0xd93e76bb.
//
// Solidity: function setIsCoin(address token, bool _isCoin) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) SetIsCoin(opts *bind.TransactOpts, token common.Address, _isCoin bool) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "setIsCoin", token, _isCoin)
}

// SetIsCoin is a paid mutator transaction binding the contract method 0xd93e76bb.
//
// Solidity: function setIsCoin(address token, bool _isCoin) returns()
func (_TwoWayCenter *TwoWayCenterSession) SetIsCoin(token common.Address, _isCoin bool) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetIsCoin(&_TwoWayCenter.TransactOpts, token, _isCoin)
}

// SetIsCoin is a paid mutator transaction binding the contract method 0xd93e76bb.
//
// Solidity: function setIsCoin(address token, bool _isCoin) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) SetIsCoin(token common.Address, _isCoin bool) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetIsCoin(&_TwoWayCenter.TransactOpts, token, _isCoin)
}

// SetRewardRatio is a paid mutator transaction binding the contract method 0x41950ad3.
//
// Solidity: function setRewardRatio(address _oToken, uint256 _ratio) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) SetRewardRatio(opts *bind.TransactOpts, _oToken common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "setRewardRatio", _oToken, _ratio)
}

// SetRewardRatio is a paid mutator transaction binding the contract method 0x41950ad3.
//
// Solidity: function setRewardRatio(address _oToken, uint256 _ratio) returns()
func (_TwoWayCenter *TwoWayCenterSession) SetRewardRatio(_oToken common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetRewardRatio(&_TwoWayCenter.TransactOpts, _oToken, _ratio)
}

// SetRewardRatio is a paid mutator transaction binding the contract method 0x41950ad3.
//
// Solidity: function setRewardRatio(address _oToken, uint256 _ratio) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) SetRewardRatio(_oToken common.Address, _ratio *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetRewardRatio(&_TwoWayCenter.TransactOpts, _oToken, _ratio)
}

// SetStakingRewards is a paid mutator transaction binding the contract method 0x9773a000.
//
// Solidity: function setStakingRewards(address[] oTokens, address[] _srs) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) SetStakingRewards(opts *bind.TransactOpts, oTokens []common.Address, _srs []common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "setStakingRewards", oTokens, _srs)
}

// SetStakingRewards is a paid mutator transaction binding the contract method 0x9773a000.
//
// Solidity: function setStakingRewards(address[] oTokens, address[] _srs) returns()
func (_TwoWayCenter *TwoWayCenterSession) SetStakingRewards(oTokens []common.Address, _srs []common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetStakingRewards(&_TwoWayCenter.TransactOpts, oTokens, _srs)
}

// SetStakingRewards is a paid mutator transaction binding the contract method 0x9773a000.
//
// Solidity: function setStakingRewards(address[] oTokens, address[] _srs) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) SetStakingRewards(oTokens []common.Address, _srs []common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetStakingRewards(&_TwoWayCenter.TransactOpts, oTokens, _srs)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) SetThreshold(opts *bind.TransactOpts, token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "setThreshold", token0, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_TwoWayCenter *TwoWayCenterSession) SetThreshold(token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetThreshold(&_TwoWayCenter.TransactOpts, token0, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) SetThreshold(token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetThreshold(&_TwoWayCenter.TransactOpts, token0, _threshold)
}

// SetTreasuryTo is a paid mutator transaction binding the contract method 0x08b60fb8.
//
// Solidity: function setTreasuryTo(address _treasuryTo) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) SetTreasuryTo(opts *bind.TransactOpts, _treasuryTo common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "setTreasuryTo", _treasuryTo)
}

// SetTreasuryTo is a paid mutator transaction binding the contract method 0x08b60fb8.
//
// Solidity: function setTreasuryTo(address _treasuryTo) returns()
func (_TwoWayCenter *TwoWayCenterSession) SetTreasuryTo(_treasuryTo common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetTreasuryTo(&_TwoWayCenter.TransactOpts, _treasuryTo)
}

// SetTreasuryTo is a paid mutator transaction binding the contract method 0x08b60fb8.
//
// Solidity: function setTreasuryTo(address _treasuryTo) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) SetTreasuryTo(_treasuryTo common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetTreasuryTo(&_TwoWayCenter.TransactOpts, _treasuryTo)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x1bdbce49.
//
// Solidity: function setWhitelist(address oToken, address user, bool _isInWhitelist) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) SetWhitelist(opts *bind.TransactOpts, oToken common.Address, user common.Address, _isInWhitelist bool) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "setWhitelist", oToken, user, _isInWhitelist)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x1bdbce49.
//
// Solidity: function setWhitelist(address oToken, address user, bool _isInWhitelist) returns()
func (_TwoWayCenter *TwoWayCenterSession) SetWhitelist(oToken common.Address, user common.Address, _isInWhitelist bool) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetWhitelist(&_TwoWayCenter.TransactOpts, oToken, user, _isInWhitelist)
}

// SetWhitelist is a paid mutator transaction binding the contract method 0x1bdbce49.
//
// Solidity: function setWhitelist(address oToken, address user, bool _isInWhitelist) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) SetWhitelist(oToken common.Address, user common.Address, _isInWhitelist bool) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.SetWhitelist(&_TwoWayCenter.TransactOpts, oToken, user, _isInWhitelist)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TwoWayCenter *TwoWayCenterSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.UpgradeTo(&_TwoWayCenter.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.UpgradeTo(&_TwoWayCenter.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TwoWayCenter *TwoWayCenterTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TwoWayCenter *TwoWayCenterSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.UpgradeToAndCall(&_TwoWayCenter.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.UpgradeToAndCall(&_TwoWayCenter.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4b2084e3.
//
// Solidity: function withdraw(address oToken, uint256 toChainId, address to, uint256 amount) returns()
func (_TwoWayCenter *TwoWayCenterTransactor) Withdraw(opts *bind.TransactOpts, oToken common.Address, toChainId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.contract.Transact(opts, "withdraw", oToken, toChainId, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4b2084e3.
//
// Solidity: function withdraw(address oToken, uint256 toChainId, address to, uint256 amount) returns()
func (_TwoWayCenter *TwoWayCenterSession) Withdraw(oToken common.Address, toChainId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.Withdraw(&_TwoWayCenter.TransactOpts, oToken, toChainId, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4b2084e3.
//
// Solidity: function withdraw(address oToken, uint256 toChainId, address to, uint256 amount) returns()
func (_TwoWayCenter *TwoWayCenterTransactorSession) Withdraw(oToken common.Address, toChainId *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TwoWayCenter.Contract.Withdraw(&_TwoWayCenter.TransactOpts, oToken, toChainId, to, amount)
}

// TwoWayCenterAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the TwoWayCenter contract.
type TwoWayCenterAdminChangedIterator struct {
	Event *TwoWayCenterAdminChanged // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterAdminChanged)
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
		it.Event = new(TwoWayCenterAdminChanged)
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
func (it *TwoWayCenterAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterAdminChanged represents a AdminChanged event raised by the TwoWayCenter contract.
type TwoWayCenterAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*TwoWayCenterAdminChangedIterator, error) {

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterAdminChangedIterator{contract: _TwoWayCenter.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TwoWayCenterAdminChanged) (event.Subscription, error) {

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterAdminChanged)
				if err := _TwoWayCenter.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TwoWayCenter *TwoWayCenterFilterer) ParseAdminChanged(log types.Log) (*TwoWayCenterAdminChanged, error) {
	event := new(TwoWayCenterAdminChanged)
	if err := _TwoWayCenter.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the TwoWayCenter contract.
type TwoWayCenterBeaconUpgradedIterator struct {
	Event *TwoWayCenterBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterBeaconUpgraded)
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
		it.Event = new(TwoWayCenterBeaconUpgraded)
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
func (it *TwoWayCenterBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterBeaconUpgraded represents a BeaconUpgraded event raised by the TwoWayCenter contract.
type TwoWayCenterBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*TwoWayCenterBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterBeaconUpgradedIterator{contract: _TwoWayCenter.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *TwoWayCenterBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterBeaconUpgraded)
				if err := _TwoWayCenter.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TwoWayCenter *TwoWayCenterFilterer) ParseBeaconUpgraded(log types.Log) (*TwoWayCenterBeaconUpgraded, error) {
	event := new(TwoWayCenterBeaconUpgraded)
	if err := _TwoWayCenter.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterCrossInFailedIterator is returned from FilterCrossInFailed and is used to iterate over the raw logs and unpacked data for CrossInFailed events raised by the TwoWayCenter contract.
type TwoWayCenterCrossInFailedIterator struct {
	Event *TwoWayCenterCrossInFailed // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterCrossInFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterCrossInFailed)
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
		it.Event = new(TwoWayCenterCrossInFailed)
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
func (it *TwoWayCenterCrossInFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterCrossInFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterCrossInFailed represents a CrossInFailed event raised by the TwoWayCenter contract.
type TwoWayCenterCrossInFailed struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCrossInFailed is a free log retrieval operation binding the contract event 0x1ecf44265cb8eead95bfe72d8e66611159bbbe6c385212224660bf7ee2287411.
//
// Solidity: event CrossInFailed((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterCrossInFailed(opts *bind.FilterOpts) (*TwoWayCenterCrossInFailedIterator, error) {

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "CrossInFailed")
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterCrossInFailedIterator{contract: _TwoWayCenter.contract, event: "CrossInFailed", logs: logs, sub: sub}, nil
}

// WatchCrossInFailed is a free log subscription operation binding the contract event 0x1ecf44265cb8eead95bfe72d8e66611159bbbe6c385212224660bf7ee2287411.
//
// Solidity: event CrossInFailed((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchCrossInFailed(opts *bind.WatchOpts, sink chan<- *TwoWayCenterCrossInFailed) (event.Subscription, error) {

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "CrossInFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterCrossInFailed)
				if err := _TwoWayCenter.contract.UnpackLog(event, "CrossInFailed", log); err != nil {
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

// ParseCrossInFailed is a log parse operation binding the contract event 0x1ecf44265cb8eead95bfe72d8e66611159bbbe6c385212224660bf7ee2287411.
//
// Solidity: event CrossInFailed((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) ParseCrossInFailed(log types.Log) (*TwoWayCenterCrossInFailed, error) {
	event := new(TwoWayCenterCrossInFailed)
	if err := _TwoWayCenter.contract.UnpackLog(event, "CrossInFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterCrossInedIterator is returned from FilterCrossIned and is used to iterate over the raw logs and unpacked data for CrossIned events raised by the TwoWayCenter contract.
type TwoWayCenterCrossInedIterator struct {
	Event *TwoWayCenterCrossIned // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterCrossInedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterCrossIned)
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
		it.Event = new(TwoWayCenterCrossIned)
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
func (it *TwoWayCenterCrossInedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterCrossInedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterCrossIned represents a CrossIned event raised by the TwoWayCenter contract.
type TwoWayCenterCrossIned struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCrossIned is a free log retrieval operation binding the contract event 0xe63fe074c53764bbe6e6f716575da08b4e5546ea802aa49ba761f2c179971541.
//
// Solidity: event CrossIned((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterCrossIned(opts *bind.FilterOpts) (*TwoWayCenterCrossInedIterator, error) {

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "CrossIned")
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterCrossInedIterator{contract: _TwoWayCenter.contract, event: "CrossIned", logs: logs, sub: sub}, nil
}

// WatchCrossIned is a free log subscription operation binding the contract event 0xe63fe074c53764bbe6e6f716575da08b4e5546ea802aa49ba761f2c179971541.
//
// Solidity: event CrossIned((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchCrossIned(opts *bind.WatchOpts, sink chan<- *TwoWayCenterCrossIned) (event.Subscription, error) {

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "CrossIned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterCrossIned)
				if err := _TwoWayCenter.contract.UnpackLog(event, "CrossIned", log); err != nil {
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

// ParseCrossIned is a log parse operation binding the contract event 0xe63fe074c53764bbe6e6f716575da08b4e5546ea802aa49ba761f2c179971541.
//
// Solidity: event CrossIned((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) ParseCrossIned(log types.Log) (*TwoWayCenterCrossIned, error) {
	event := new(TwoWayCenterCrossIned)
	if err := _TwoWayCenter.contract.UnpackLog(event, "CrossIned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterCrossOutedIterator is returned from FilterCrossOuted and is used to iterate over the raw logs and unpacked data for CrossOuted events raised by the TwoWayCenter contract.
type TwoWayCenterCrossOutedIterator struct {
	Event *TwoWayCenterCrossOuted // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterCrossOutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterCrossOuted)
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
		it.Event = new(TwoWayCenterCrossOuted)
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
func (it *TwoWayCenterCrossOutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterCrossOutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterCrossOuted represents a CrossOuted event raised by the TwoWayCenter contract.
type TwoWayCenterCrossOuted struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCrossOuted is a free log retrieval operation binding the contract event 0xe8ae00f8b99ac422d841599f0cc887e0d20c0bf99271117d943e474f280f1145.
//
// Solidity: event CrossOuted((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterCrossOuted(opts *bind.FilterOpts) (*TwoWayCenterCrossOutedIterator, error) {

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "CrossOuted")
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterCrossOutedIterator{contract: _TwoWayCenter.contract, event: "CrossOuted", logs: logs, sub: sub}, nil
}

// WatchCrossOuted is a free log subscription operation binding the contract event 0xe8ae00f8b99ac422d841599f0cc887e0d20c0bf99271117d943e474f280f1145.
//
// Solidity: event CrossOuted((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchCrossOuted(opts *bind.WatchOpts, sink chan<- *TwoWayCenterCrossOuted) (event.Subscription, error) {

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "CrossOuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterCrossOuted)
				if err := _TwoWayCenter.contract.UnpackLog(event, "CrossOuted", log); err != nil {
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

// ParseCrossOuted is a log parse operation binding the contract event 0xe8ae00f8b99ac422d841599f0cc887e0d20c0bf99271117d943e474f280f1145.
//
// Solidity: event CrossOuted((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) ParseCrossOuted(log types.Log) (*TwoWayCenterCrossOuted, error) {
	event := new(TwoWayCenterCrossOuted)
	if err := _TwoWayCenter.contract.UnpackLog(event, "CrossOuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the TwoWayCenter contract.
type TwoWayCenterDepositedIterator struct {
	Event *TwoWayCenterDeposited // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterDeposited)
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
		it.Event = new(TwoWayCenterDeposited)
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
func (it *TwoWayCenterDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterDeposited represents a Deposited event raised by the TwoWayCenter contract.
type TwoWayCenterDeposited struct {
	FromChainId *big.Int
	FromToken   common.Address
	From        common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x984a71c9d95fd4794aeba33ae72edfec22053fde75488d63abef9dc69ee795af.
//
// Solidity: event Deposited(uint256 fromChainId, address fromToken, address from, uint256 amount)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterDeposited(opts *bind.FilterOpts) (*TwoWayCenterDepositedIterator, error) {

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterDepositedIterator{contract: _TwoWayCenter.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x984a71c9d95fd4794aeba33ae72edfec22053fde75488d63abef9dc69ee795af.
//
// Solidity: event Deposited(uint256 fromChainId, address fromToken, address from, uint256 amount)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *TwoWayCenterDeposited) (event.Subscription, error) {

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterDeposited)
				if err := _TwoWayCenter.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_TwoWayCenter *TwoWayCenterFilterer) ParseDeposited(log types.Log) (*TwoWayCenterDeposited, error) {
	event := new(TwoWayCenterDeposited)
	if err := _TwoWayCenter.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterForwardCrossOutFailedIterator is returned from FilterForwardCrossOutFailed and is used to iterate over the raw logs and unpacked data for ForwardCrossOutFailed events raised by the TwoWayCenter contract.
type TwoWayCenterForwardCrossOutFailedIterator struct {
	Event *TwoWayCenterForwardCrossOutFailed // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterForwardCrossOutFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterForwardCrossOutFailed)
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
		it.Event = new(TwoWayCenterForwardCrossOutFailed)
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
func (it *TwoWayCenterForwardCrossOutFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterForwardCrossOutFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterForwardCrossOutFailed represents a ForwardCrossOutFailed event raised by the TwoWayCenter contract.
type TwoWayCenterForwardCrossOutFailed struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterForwardCrossOutFailed is a free log retrieval operation binding the contract event 0xa22f2d032a9fe57cd0476bed912548f17dccb0695666d445bed754d98219b6a0.
//
// Solidity: event ForwardCrossOutFailed((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterForwardCrossOutFailed(opts *bind.FilterOpts) (*TwoWayCenterForwardCrossOutFailedIterator, error) {

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "ForwardCrossOutFailed")
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterForwardCrossOutFailedIterator{contract: _TwoWayCenter.contract, event: "ForwardCrossOutFailed", logs: logs, sub: sub}, nil
}

// WatchForwardCrossOutFailed is a free log subscription operation binding the contract event 0xa22f2d032a9fe57cd0476bed912548f17dccb0695666d445bed754d98219b6a0.
//
// Solidity: event ForwardCrossOutFailed((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchForwardCrossOutFailed(opts *bind.WatchOpts, sink chan<- *TwoWayCenterForwardCrossOutFailed) (event.Subscription, error) {

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "ForwardCrossOutFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterForwardCrossOutFailed)
				if err := _TwoWayCenter.contract.UnpackLog(event, "ForwardCrossOutFailed", log); err != nil {
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

// ParseForwardCrossOutFailed is a log parse operation binding the contract event 0xa22f2d032a9fe57cd0476bed912548f17dccb0695666d445bed754d98219b6a0.
//
// Solidity: event ForwardCrossOutFailed((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) ParseForwardCrossOutFailed(log types.Log) (*TwoWayCenterForwardCrossOutFailed, error) {
	event := new(TwoWayCenterForwardCrossOutFailed)
	if err := _TwoWayCenter.contract.UnpackLog(event, "ForwardCrossOutFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterForwardCrossOutedIterator is returned from FilterForwardCrossOuted and is used to iterate over the raw logs and unpacked data for ForwardCrossOuted events raised by the TwoWayCenter contract.
type TwoWayCenterForwardCrossOutedIterator struct {
	Event *TwoWayCenterForwardCrossOuted // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterForwardCrossOutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterForwardCrossOuted)
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
		it.Event = new(TwoWayCenterForwardCrossOuted)
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
func (it *TwoWayCenterForwardCrossOutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterForwardCrossOutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterForwardCrossOuted represents a ForwardCrossOuted event raised by the TwoWayCenter contract.
type TwoWayCenterForwardCrossOuted struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterForwardCrossOuted is a free log retrieval operation binding the contract event 0x715d2e6514c297c9e18ba16c3e3972b7562b34f8f185c21a5cc305324039c77a.
//
// Solidity: event ForwardCrossOuted((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterForwardCrossOuted(opts *bind.FilterOpts) (*TwoWayCenterForwardCrossOutedIterator, error) {

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "ForwardCrossOuted")
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterForwardCrossOutedIterator{contract: _TwoWayCenter.contract, event: "ForwardCrossOuted", logs: logs, sub: sub}, nil
}

// WatchForwardCrossOuted is a free log subscription operation binding the contract event 0x715d2e6514c297c9e18ba16c3e3972b7562b34f8f185c21a5cc305324039c77a.
//
// Solidity: event ForwardCrossOuted((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchForwardCrossOuted(opts *bind.WatchOpts, sink chan<- *TwoWayCenterForwardCrossOuted) (event.Subscription, error) {

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "ForwardCrossOuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterForwardCrossOuted)
				if err := _TwoWayCenter.contract.UnpackLog(event, "ForwardCrossOuted", log); err != nil {
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

// ParseForwardCrossOuted is a log parse operation binding the contract event 0x715d2e6514c297c9e18ba16c3e3972b7562b34f8f185c21a5cc305324039c77a.
//
// Solidity: event ForwardCrossOuted((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) ParseForwardCrossOuted(log types.Log) (*TwoWayCenterForwardCrossOuted, error) {
	event := new(TwoWayCenterForwardCrossOuted)
	if err := _TwoWayCenter.contract.UnpackLog(event, "ForwardCrossOuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterIssuedIterator is returned from FilterIssued and is used to iterate over the raw logs and unpacked data for Issued events raised by the TwoWayCenter contract.
type TwoWayCenterIssuedIterator struct {
	Event *TwoWayCenterIssued // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterIssuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterIssued)
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
		it.Event = new(TwoWayCenterIssued)
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
func (it *TwoWayCenterIssuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterIssuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterIssued represents a Issued event raised by the TwoWayCenter contract.
type TwoWayCenterIssued struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIssued is a free log retrieval operation binding the contract event 0x082ec0a3f5481f56ed1aa06ed4b5f96a0b82a7f1ce3342041b20b8161fabeb92.
//
// Solidity: event Issued((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterIssued(opts *bind.FilterOpts) (*TwoWayCenterIssuedIterator, error) {

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "Issued")
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterIssuedIterator{contract: _TwoWayCenter.contract, event: "Issued", logs: logs, sub: sub}, nil
}

// WatchIssued is a free log subscription operation binding the contract event 0x082ec0a3f5481f56ed1aa06ed4b5f96a0b82a7f1ce3342041b20b8161fabeb92.
//
// Solidity: event Issued((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchIssued(opts *bind.WatchOpts, sink chan<- *TwoWayCenterIssued) (event.Subscription, error) {

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "Issued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterIssued)
				if err := _TwoWayCenter.contract.UnpackLog(event, "Issued", log); err != nil {
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

// ParseIssued is a log parse operation binding the contract event 0x082ec0a3f5481f56ed1aa06ed4b5f96a0b82a7f1ce3342041b20b8161fabeb92.
//
// Solidity: event Issued((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) ParseIssued(log types.Log) (*TwoWayCenterIssued, error) {
	event := new(TwoWayCenterIssued)
	if err := _TwoWayCenter.contract.UnpackLog(event, "Issued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the TwoWayCenter contract.
type TwoWayCenterProposalVotedIterator struct {
	Event *TwoWayCenterProposalVoted // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterProposalVoted)
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
		it.Event = new(TwoWayCenterProposalVoted)
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
func (it *TwoWayCenterProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterProposalVoted represents a ProposalVoted event raised by the TwoWayCenter contract.
type TwoWayCenterProposalVoted struct {
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
func (_TwoWayCenter *TwoWayCenterFilterer) FilterProposalVoted(opts *bind.FilterOpts) (*TwoWayCenterProposalVotedIterator, error) {

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterProposalVotedIterator{contract: _TwoWayCenter.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0xe458477b94285c9f254a407e2614c39f55230b527ff7f51d76fdf33e251ae936.
//
// Solidity: event ProposalVoted(address token, address from, address to, uint256 amount, address proposer, uint256 count, uint256 threshold)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *TwoWayCenterProposalVoted) (event.Subscription, error) {

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterProposalVoted)
				if err := _TwoWayCenter.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
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
func (_TwoWayCenter *TwoWayCenterFilterer) ParseProposalVoted(log types.Log) (*TwoWayCenterProposalVoted, error) {
	event := new(TwoWayCenterProposalVoted)
	if err := _TwoWayCenter.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TwoWayCenter contract.
type TwoWayCenterRoleAdminChangedIterator struct {
	Event *TwoWayCenterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterRoleAdminChanged)
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
		it.Event = new(TwoWayCenterRoleAdminChanged)
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
func (it *TwoWayCenterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterRoleAdminChanged represents a RoleAdminChanged event raised by the TwoWayCenter contract.
type TwoWayCenterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TwoWayCenterRoleAdminChangedIterator, error) {

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

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterRoleAdminChangedIterator{contract: _TwoWayCenter.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TwoWayCenterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterRoleAdminChanged)
				if err := _TwoWayCenter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_TwoWayCenter *TwoWayCenterFilterer) ParseRoleAdminChanged(log types.Log) (*TwoWayCenterRoleAdminChanged, error) {
	event := new(TwoWayCenterRoleAdminChanged)
	if err := _TwoWayCenter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TwoWayCenter contract.
type TwoWayCenterRoleGrantedIterator struct {
	Event *TwoWayCenterRoleGranted // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterRoleGranted)
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
		it.Event = new(TwoWayCenterRoleGranted)
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
func (it *TwoWayCenterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterRoleGranted represents a RoleGranted event raised by the TwoWayCenter contract.
type TwoWayCenterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TwoWayCenterRoleGrantedIterator, error) {

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

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterRoleGrantedIterator{contract: _TwoWayCenter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TwoWayCenterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterRoleGranted)
				if err := _TwoWayCenter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_TwoWayCenter *TwoWayCenterFilterer) ParseRoleGranted(log types.Log) (*TwoWayCenterRoleGranted, error) {
	event := new(TwoWayCenterRoleGranted)
	if err := _TwoWayCenter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TwoWayCenter contract.
type TwoWayCenterRoleRevokedIterator struct {
	Event *TwoWayCenterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterRoleRevoked)
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
		it.Event = new(TwoWayCenterRoleRevoked)
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
func (it *TwoWayCenterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterRoleRevoked represents a RoleRevoked event raised by the TwoWayCenter contract.
type TwoWayCenterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TwoWayCenterRoleRevokedIterator, error) {

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

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterRoleRevokedIterator{contract: _TwoWayCenter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TwoWayCenterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterRoleRevoked)
				if err := _TwoWayCenter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_TwoWayCenter *TwoWayCenterFilterer) ParseRoleRevoked(log types.Log) (*TwoWayCenterRoleRevoked, error) {
	event := new(TwoWayCenterRoleRevoked)
	if err := _TwoWayCenter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterStakingRewardResetedIterator is returned from FilterStakingRewardReseted and is used to iterate over the raw logs and unpacked data for StakingRewardReseted events raised by the TwoWayCenter contract.
type TwoWayCenterStakingRewardResetedIterator struct {
	Event *TwoWayCenterStakingRewardReseted // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterStakingRewardResetedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterStakingRewardReseted)
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
		it.Event = new(TwoWayCenterStakingRewardReseted)
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
func (it *TwoWayCenterStakingRewardResetedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterStakingRewardResetedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterStakingRewardReseted represents a StakingRewardReseted event raised by the TwoWayCenter contract.
type TwoWayCenterStakingRewardReseted struct {
	OToken common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStakingRewardReseted is a free log retrieval operation binding the contract event 0x4b8830da429e7f8c1dc368e9a9a3dc21aa6e3e08d115679e616b1fedaf420aeb.
//
// Solidity: event StakingRewardReseted(address oToken)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterStakingRewardReseted(opts *bind.FilterOpts) (*TwoWayCenterStakingRewardResetedIterator, error) {

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "StakingRewardReseted")
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterStakingRewardResetedIterator{contract: _TwoWayCenter.contract, event: "StakingRewardReseted", logs: logs, sub: sub}, nil
}

// WatchStakingRewardReseted is a free log subscription operation binding the contract event 0x4b8830da429e7f8c1dc368e9a9a3dc21aa6e3e08d115679e616b1fedaf420aeb.
//
// Solidity: event StakingRewardReseted(address oToken)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchStakingRewardReseted(opts *bind.WatchOpts, sink chan<- *TwoWayCenterStakingRewardReseted) (event.Subscription, error) {

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "StakingRewardReseted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterStakingRewardReseted)
				if err := _TwoWayCenter.contract.UnpackLog(event, "StakingRewardReseted", log); err != nil {
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
func (_TwoWayCenter *TwoWayCenterFilterer) ParseStakingRewardReseted(log types.Log) (*TwoWayCenterStakingRewardReseted, error) {
	event := new(TwoWayCenterStakingRewardReseted)
	if err := _TwoWayCenter.contract.UnpackLog(event, "StakingRewardReseted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterSupportedIterator is returned from FilterSupported and is used to iterate over the raw logs and unpacked data for Supported events raised by the TwoWayCenter contract.
type TwoWayCenterSupportedIterator struct {
	Event *TwoWayCenterSupported // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterSupportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterSupported)
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
		it.Event = new(TwoWayCenterSupported)
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
func (it *TwoWayCenterSupportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterSupportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterSupported represents a Supported event raised by the TwoWayCenter contract.
type TwoWayCenterSupported struct {
	Token   common.Address
	ChainId *big.Int
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSupported is a free log retrieval operation binding the contract event 0xd7cadba2609cba4364d862761a89c823e956d810d9947b1f158ad1aa9c2affbc.
//
// Solidity: event Supported(address token, uint256 chainId, bool status)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterSupported(opts *bind.FilterOpts) (*TwoWayCenterSupportedIterator, error) {

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "Supported")
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterSupportedIterator{contract: _TwoWayCenter.contract, event: "Supported", logs: logs, sub: sub}, nil
}

// WatchSupported is a free log subscription operation binding the contract event 0xd7cadba2609cba4364d862761a89c823e956d810d9947b1f158ad1aa9c2affbc.
//
// Solidity: event Supported(address token, uint256 chainId, bool status)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchSupported(opts *bind.WatchOpts, sink chan<- *TwoWayCenterSupported) (event.Subscription, error) {

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "Supported")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterSupported)
				if err := _TwoWayCenter.contract.UnpackLog(event, "Supported", log); err != nil {
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
func (_TwoWayCenter *TwoWayCenterFilterer) ParseSupported(log types.Log) (*TwoWayCenterSupported, error) {
	event := new(TwoWayCenterSupported)
	if err := _TwoWayCenter.contract.UnpackLog(event, "Supported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the TwoWayCenter contract.
type TwoWayCenterThresholdChangedIterator struct {
	Event *TwoWayCenterThresholdChanged // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterThresholdChanged)
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
		it.Event = new(TwoWayCenterThresholdChanged)
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
func (it *TwoWayCenterThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterThresholdChanged represents a ThresholdChanged event raised by the TwoWayCenter contract.
type TwoWayCenterThresholdChanged struct {
	Token        common.Address
	OldThreshold *big.Int
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*TwoWayCenterThresholdChangedIterator, error) {

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterThresholdChangedIterator{contract: _TwoWayCenter.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *TwoWayCenterThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterThresholdChanged)
				if err := _TwoWayCenter.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_TwoWayCenter *TwoWayCenterFilterer) ParseThresholdChanged(log types.Log) (*TwoWayCenterThresholdChanged, error) {
	event := new(TwoWayCenterThresholdChanged)
	if err := _TwoWayCenter.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TwoWayCenter contract.
type TwoWayCenterUpgradedIterator struct {
	Event *TwoWayCenterUpgraded // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterUpgraded)
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
		it.Event = new(TwoWayCenterUpgraded)
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
func (it *TwoWayCenterUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterUpgraded represents a Upgraded event raised by the TwoWayCenter contract.
type TwoWayCenterUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TwoWayCenterUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterUpgradedIterator{contract: _TwoWayCenter.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TwoWayCenterUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterUpgraded)
				if err := _TwoWayCenter.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_TwoWayCenter *TwoWayCenterFilterer) ParseUpgraded(log types.Log) (*TwoWayCenterUpgraded, error) {
	event := new(TwoWayCenterUpgraded)
	if err := _TwoWayCenter.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterWithdrawedIterator is returned from FilterWithdrawed and is used to iterate over the raw logs and unpacked data for Withdrawed events raised by the TwoWayCenter contract.
type TwoWayCenterWithdrawedIterator struct {
	Event *TwoWayCenterWithdrawed // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterWithdrawedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterWithdrawed)
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
		it.Event = new(TwoWayCenterWithdrawed)
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
func (it *TwoWayCenterWithdrawedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterWithdrawedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterWithdrawed represents a Withdrawed event raised by the TwoWayCenter contract.
type TwoWayCenterWithdrawed struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawed is a free log retrieval operation binding the contract event 0xf0bc31bd6717a3afcc9a7f44d5b6d8638707a8ee613a6d3b38ec03e9f0392d67.
//
// Solidity: event Withdrawed((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterWithdrawed(opts *bind.FilterOpts) (*TwoWayCenterWithdrawedIterator, error) {

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "Withdrawed")
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterWithdrawedIterator{contract: _TwoWayCenter.contract, event: "Withdrawed", logs: logs, sub: sub}, nil
}

// WatchWithdrawed is a free log subscription operation binding the contract event 0xf0bc31bd6717a3afcc9a7f44d5b6d8638707a8ee613a6d3b38ec03e9f0392d67.
//
// Solidity: event Withdrawed((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchWithdrawed(opts *bind.WatchOpts, sink chan<- *TwoWayCenterWithdrawed) (event.Subscription, error) {

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "Withdrawed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterWithdrawed)
				if err := _TwoWayCenter.contract.UnpackLog(event, "Withdrawed", log); err != nil {
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

// ParseWithdrawed is a log parse operation binding the contract event 0xf0bc31bd6717a3afcc9a7f44d5b6d8638707a8ee613a6d3b38ec03e9f0392d67.
//
// Solidity: event Withdrawed((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) ParseWithdrawed(log types.Log) (*TwoWayCenterWithdrawed, error) {
	event := new(TwoWayCenterWithdrawed)
	if err := _TwoWayCenter.contract.UnpackLog(event, "Withdrawed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayCenterWithdrawedToCenterIterator is returned from FilterWithdrawedToCenter and is used to iterate over the raw logs and unpacked data for WithdrawedToCenter events raised by the TwoWayCenter contract.
type TwoWayCenterWithdrawedToCenterIterator struct {
	Event *TwoWayCenterWithdrawedToCenter // Event containing the contract specifics and raw log

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
func (it *TwoWayCenterWithdrawedToCenterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCenterWithdrawedToCenter)
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
		it.Event = new(TwoWayCenterWithdrawedToCenter)
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
func (it *TwoWayCenterWithdrawedToCenterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCenterWithdrawedToCenterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCenterWithdrawedToCenter represents a WithdrawedToCenter event raised by the TwoWayCenter contract.
type TwoWayCenterWithdrawedToCenter struct {
	P   InParam
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWithdrawedToCenter is a free log retrieval operation binding the contract event 0x7f50f8e7f2e365c46040fcc45f0f1ee1b3f7d8c3befc68282bfa0562f4f1d2b0.
//
// Solidity: event WithdrawedToCenter((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) FilterWithdrawedToCenter(opts *bind.FilterOpts) (*TwoWayCenterWithdrawedToCenterIterator, error) {

	logs, sub, err := _TwoWayCenter.contract.FilterLogs(opts, "WithdrawedToCenter")
	if err != nil {
		return nil, err
	}
	return &TwoWayCenterWithdrawedToCenterIterator{contract: _TwoWayCenter.contract, event: "WithdrawedToCenter", logs: logs, sub: sub}, nil
}

// WatchWithdrawedToCenter is a free log subscription operation binding the contract event 0x7f50f8e7f2e365c46040fcc45f0f1ee1b3f7d8c3befc68282bfa0562f4f1d2b0.
//
// Solidity: event WithdrawedToCenter((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) WatchWithdrawedToCenter(opts *bind.WatchOpts, sink chan<- *TwoWayCenterWithdrawedToCenter) (event.Subscription, error) {

	logs, sub, err := _TwoWayCenter.contract.WatchLogs(opts, "WithdrawedToCenter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCenterWithdrawedToCenter)
				if err := _TwoWayCenter.contract.UnpackLog(event, "WithdrawedToCenter", log); err != nil {
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

// ParseWithdrawedToCenter is a log parse operation binding the contract event 0x7f50f8e7f2e365c46040fcc45f0f1ee1b3f7d8c3befc68282bfa0562f4f1d2b0.
//
// Solidity: event WithdrawedToCenter((uint256,address,address,uint256,address,address,uint256) p)
func (_TwoWayCenter *TwoWayCenterFilterer) ParseWithdrawedToCenter(log types.Log) (*TwoWayCenterWithdrawedToCenter, error) {
	event := new(TwoWayCenterWithdrawedToCenter)
	if err := _TwoWayCenter.contract.UnpackLog(event, "WithdrawedToCenter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayProposalVoteMetaData contains all meta data concerning the TwoWayProposalVote contract.
var TwoWayProposalVoteMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c86ec2bf": "threshold(address)",
	},
	Bin: "0x608060405234801561001057600080fd5b5060be8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063c86ec2bf14602d575b600080fd5b604a6038366004605c565b60006020819052908152604090205481565b60405190815260200160405180910390f35b600060208284031215606c578081fd5b81356001600160a01b03811681146081578182fd5b939250505056fea26469706673582212203ebeae6abb7ad644c842ef4b9a4162e78d9163905096933ae2be82cde002887e64736f6c63430008020033",
}

// TwoWayProposalVoteABI is the input ABI used to generate the binding from.
// Deprecated: Use TwoWayProposalVoteMetaData.ABI instead.
var TwoWayProposalVoteABI = TwoWayProposalVoteMetaData.ABI

// Deprecated: Use TwoWayProposalVoteMetaData.Sigs instead.
// TwoWayProposalVoteFuncSigs maps the 4-byte function signature to its string representation.
var TwoWayProposalVoteFuncSigs = TwoWayProposalVoteMetaData.Sigs

// TwoWayProposalVoteBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TwoWayProposalVoteMetaData.Bin instead.
var TwoWayProposalVoteBin = TwoWayProposalVoteMetaData.Bin

// DeployTwoWayProposalVote deploys a new Ethereum contract, binding an instance of TwoWayProposalVote to it.
func DeployTwoWayProposalVote(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TwoWayProposalVote, error) {
	parsed, err := TwoWayProposalVoteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TwoWayProposalVoteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TwoWayProposalVote{TwoWayProposalVoteCaller: TwoWayProposalVoteCaller{contract: contract}, TwoWayProposalVoteTransactor: TwoWayProposalVoteTransactor{contract: contract}, TwoWayProposalVoteFilterer: TwoWayProposalVoteFilterer{contract: contract}}, nil
}

// TwoWayProposalVote is an auto generated Go binding around an Ethereum contract.
type TwoWayProposalVote struct {
	TwoWayProposalVoteCaller     // Read-only binding to the contract
	TwoWayProposalVoteTransactor // Write-only binding to the contract
	TwoWayProposalVoteFilterer   // Log filterer for contract events
}

// TwoWayProposalVoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type TwoWayProposalVoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoWayProposalVoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TwoWayProposalVoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoWayProposalVoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TwoWayProposalVoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoWayProposalVoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TwoWayProposalVoteSession struct {
	Contract     *TwoWayProposalVote // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TwoWayProposalVoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TwoWayProposalVoteCallerSession struct {
	Contract *TwoWayProposalVoteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TwoWayProposalVoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TwoWayProposalVoteTransactorSession struct {
	Contract     *TwoWayProposalVoteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TwoWayProposalVoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type TwoWayProposalVoteRaw struct {
	Contract *TwoWayProposalVote // Generic contract binding to access the raw methods on
}

// TwoWayProposalVoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TwoWayProposalVoteCallerRaw struct {
	Contract *TwoWayProposalVoteCaller // Generic read-only contract binding to access the raw methods on
}

// TwoWayProposalVoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TwoWayProposalVoteTransactorRaw struct {
	Contract *TwoWayProposalVoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTwoWayProposalVote creates a new instance of TwoWayProposalVote, bound to a specific deployed contract.
func NewTwoWayProposalVote(address common.Address, backend bind.ContractBackend) (*TwoWayProposalVote, error) {
	contract, err := bindTwoWayProposalVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TwoWayProposalVote{TwoWayProposalVoteCaller: TwoWayProposalVoteCaller{contract: contract}, TwoWayProposalVoteTransactor: TwoWayProposalVoteTransactor{contract: contract}, TwoWayProposalVoteFilterer: TwoWayProposalVoteFilterer{contract: contract}}, nil
}

// NewTwoWayProposalVoteCaller creates a new read-only instance of TwoWayProposalVote, bound to a specific deployed contract.
func NewTwoWayProposalVoteCaller(address common.Address, caller bind.ContractCaller) (*TwoWayProposalVoteCaller, error) {
	contract, err := bindTwoWayProposalVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TwoWayProposalVoteCaller{contract: contract}, nil
}

// NewTwoWayProposalVoteTransactor creates a new write-only instance of TwoWayProposalVote, bound to a specific deployed contract.
func NewTwoWayProposalVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*TwoWayProposalVoteTransactor, error) {
	contract, err := bindTwoWayProposalVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TwoWayProposalVoteTransactor{contract: contract}, nil
}

// NewTwoWayProposalVoteFilterer creates a new log filterer instance of TwoWayProposalVote, bound to a specific deployed contract.
func NewTwoWayProposalVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*TwoWayProposalVoteFilterer, error) {
	contract, err := bindTwoWayProposalVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TwoWayProposalVoteFilterer{contract: contract}, nil
}

// bindTwoWayProposalVote binds a generic wrapper to an already deployed contract.
func bindTwoWayProposalVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TwoWayProposalVoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwoWayProposalVote *TwoWayProposalVoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwoWayProposalVote.Contract.TwoWayProposalVoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwoWayProposalVote *TwoWayProposalVoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwoWayProposalVote.Contract.TwoWayProposalVoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwoWayProposalVote *TwoWayProposalVoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwoWayProposalVote.Contract.TwoWayProposalVoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwoWayProposalVote *TwoWayProposalVoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwoWayProposalVote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwoWayProposalVote *TwoWayProposalVoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwoWayProposalVote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwoWayProposalVote *TwoWayProposalVoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwoWayProposalVote.Contract.contract.Transact(opts, method, params...)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_TwoWayProposalVote *TwoWayProposalVoteCaller) Threshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayProposalVote.contract.Call(opts, &out, "threshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_TwoWayProposalVote *TwoWayProposalVoteSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _TwoWayProposalVote.Contract.Threshold(&_TwoWayProposalVote.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_TwoWayProposalVote *TwoWayProposalVoteCallerSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _TwoWayProposalVote.Contract.Threshold(&_TwoWayProposalVote.CallOpts, arg0)
}

// TwoWayProposalVoteProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the TwoWayProposalVote contract.
type TwoWayProposalVoteProposalVotedIterator struct {
	Event *TwoWayProposalVoteProposalVoted // Event containing the contract specifics and raw log

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
func (it *TwoWayProposalVoteProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayProposalVoteProposalVoted)
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
		it.Event = new(TwoWayProposalVoteProposalVoted)
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
func (it *TwoWayProposalVoteProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayProposalVoteProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayProposalVoteProposalVoted represents a ProposalVoted event raised by the TwoWayProposalVote contract.
type TwoWayProposalVoteProposalVoted struct {
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
func (_TwoWayProposalVote *TwoWayProposalVoteFilterer) FilterProposalVoted(opts *bind.FilterOpts) (*TwoWayProposalVoteProposalVotedIterator, error) {

	logs, sub, err := _TwoWayProposalVote.contract.FilterLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return &TwoWayProposalVoteProposalVotedIterator{contract: _TwoWayProposalVote.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0xe458477b94285c9f254a407e2614c39f55230b527ff7f51d76fdf33e251ae936.
//
// Solidity: event ProposalVoted(address token, address from, address to, uint256 amount, address proposer, uint256 count, uint256 threshold)
func (_TwoWayProposalVote *TwoWayProposalVoteFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *TwoWayProposalVoteProposalVoted) (event.Subscription, error) {

	logs, sub, err := _TwoWayProposalVote.contract.WatchLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayProposalVoteProposalVoted)
				if err := _TwoWayProposalVote.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
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
func (_TwoWayProposalVote *TwoWayProposalVoteFilterer) ParseProposalVoted(log types.Log) (*TwoWayProposalVoteProposalVoted, error) {
	event := new(TwoWayProposalVoteProposalVoted)
	if err := _TwoWayProposalVote.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayProposalVoteThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the TwoWayProposalVote contract.
type TwoWayProposalVoteThresholdChangedIterator struct {
	Event *TwoWayProposalVoteThresholdChanged // Event containing the contract specifics and raw log

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
func (it *TwoWayProposalVoteThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayProposalVoteThresholdChanged)
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
		it.Event = new(TwoWayProposalVoteThresholdChanged)
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
func (it *TwoWayProposalVoteThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayProposalVoteThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayProposalVoteThresholdChanged represents a ThresholdChanged event raised by the TwoWayProposalVote contract.
type TwoWayProposalVoteThresholdChanged struct {
	Token        common.Address
	OldThreshold *big.Int
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_TwoWayProposalVote *TwoWayProposalVoteFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*TwoWayProposalVoteThresholdChangedIterator, error) {

	logs, sub, err := _TwoWayProposalVote.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &TwoWayProposalVoteThresholdChangedIterator{contract: _TwoWayProposalVote.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_TwoWayProposalVote *TwoWayProposalVoteFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *TwoWayProposalVoteThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _TwoWayProposalVote.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayProposalVoteThresholdChanged)
				if err := _TwoWayProposalVote.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_TwoWayProposalVote *TwoWayProposalVoteFilterer) ParseThresholdChanged(log types.Log) (*TwoWayProposalVoteThresholdChanged, error) {
	event := new(TwoWayProposalVoteThresholdChanged)
	if err := _TwoWayProposalVote.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UUPSUpgradeableMetaData contains all meta data concerning the UUPSUpgradeable contract.
var UUPSUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3659cfe6": "upgradeTo(address)",
		"4f1ef286": "upgradeToAndCall(address,bytes)",
	},
}

// UUPSUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use UUPSUpgradeableMetaData.ABI instead.
var UUPSUpgradeableABI = UUPSUpgradeableMetaData.ABI

// Deprecated: Use UUPSUpgradeableMetaData.Sigs instead.
// UUPSUpgradeableFuncSigs maps the 4-byte function signature to its string representation.
var UUPSUpgradeableFuncSigs = UUPSUpgradeableMetaData.Sigs

// UUPSUpgradeable is an auto generated Go binding around an Ethereum contract.
type UUPSUpgradeable struct {
	UUPSUpgradeableCaller     // Read-only binding to the contract
	UUPSUpgradeableTransactor // Write-only binding to the contract
	UUPSUpgradeableFilterer   // Log filterer for contract events
}

// UUPSUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type UUPSUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UUPSUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UUPSUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UUPSUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UUPSUpgradeableSession struct {
	Contract     *UUPSUpgradeable  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UUPSUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UUPSUpgradeableCallerSession struct {
	Contract *UUPSUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UUPSUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UUPSUpgradeableTransactorSession struct {
	Contract     *UUPSUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UUPSUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type UUPSUpgradeableRaw struct {
	Contract *UUPSUpgradeable // Generic contract binding to access the raw methods on
}

// UUPSUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UUPSUpgradeableCallerRaw struct {
	Contract *UUPSUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// UUPSUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UUPSUpgradeableTransactorRaw struct {
	Contract *UUPSUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUUPSUpgradeable creates a new instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeable(address common.Address, backend bind.ContractBackend) (*UUPSUpgradeable, error) {
	contract, err := bindUUPSUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeable{UUPSUpgradeableCaller: UUPSUpgradeableCaller{contract: contract}, UUPSUpgradeableTransactor: UUPSUpgradeableTransactor{contract: contract}, UUPSUpgradeableFilterer: UUPSUpgradeableFilterer{contract: contract}}, nil
}

// NewUUPSUpgradeableCaller creates a new read-only instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*UUPSUpgradeableCaller, error) {
	contract, err := bindUUPSUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableCaller{contract: contract}, nil
}

// NewUUPSUpgradeableTransactor creates a new write-only instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*UUPSUpgradeableTransactor, error) {
	contract, err := bindUUPSUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableTransactor{contract: contract}, nil
}

// NewUUPSUpgradeableFilterer creates a new log filterer instance of UUPSUpgradeable, bound to a specific deployed contract.
func NewUUPSUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*UUPSUpgradeableFilterer, error) {
	contract, err := bindUUPSUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableFilterer{contract: contract}, nil
}

// bindUUPSUpgradeable binds a generic wrapper to an already deployed contract.
func bindUUPSUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UUPSUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UUPSUpgradeable *UUPSUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UUPSUpgradeable.Contract.UUPSUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UUPSUpgradeable *UUPSUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UUPSUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UUPSUpgradeable *UUPSUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UUPSUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UUPSUpgradeable *UUPSUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UUPSUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UUPSUpgradeable *UUPSUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UUPSUpgradeable *UUPSUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSUpgradeable.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSUpgradeable *UUPSUpgradeableSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeTo(&_UUPSUpgradeable.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeTo(&_UUPSUpgradeable.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSUpgradeable.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSUpgradeable *UUPSUpgradeableSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeToAndCall(&_UUPSUpgradeable.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_UUPSUpgradeable *UUPSUpgradeableTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _UUPSUpgradeable.Contract.UpgradeToAndCall(&_UUPSUpgradeable.TransactOpts, newImplementation, data)
}

// UUPSUpgradeableAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableAdminChangedIterator struct {
	Event *UUPSUpgradeableAdminChanged // Event containing the contract specifics and raw log

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
func (it *UUPSUpgradeableAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableAdminChanged)
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
		it.Event = new(UUPSUpgradeableAdminChanged)
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
func (it *UUPSUpgradeableAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableAdminChanged represents a AdminChanged event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*UUPSUpgradeableAdminChangedIterator, error) {

	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableAdminChangedIterator{contract: _UUPSUpgradeable.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableAdminChanged) (event.Subscription, error) {

	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableAdminChanged)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseAdminChanged(log types.Log) (*UUPSUpgradeableAdminChanged, error) {
	event := new(UUPSUpgradeableAdminChanged)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UUPSUpgradeableBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableBeaconUpgradedIterator struct {
	Event *UUPSUpgradeableBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *UUPSUpgradeableBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableBeaconUpgraded)
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
		it.Event = new(UUPSUpgradeableBeaconUpgraded)
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
func (it *UUPSUpgradeableBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableBeaconUpgraded represents a BeaconUpgraded event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*UUPSUpgradeableBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableBeaconUpgradedIterator{contract: _UUPSUpgradeable.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableBeaconUpgraded)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseBeaconUpgraded(log types.Log) (*UUPSUpgradeableBeaconUpgraded, error) {
	event := new(UUPSUpgradeableBeaconUpgraded)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UUPSUpgradeableUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the UUPSUpgradeable contract.
type UUPSUpgradeableUpgradedIterator struct {
	Event *UUPSUpgradeableUpgraded // Event containing the contract specifics and raw log

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
func (it *UUPSUpgradeableUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UUPSUpgradeableUpgraded)
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
		it.Event = new(UUPSUpgradeableUpgraded)
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
func (it *UUPSUpgradeableUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UUPSUpgradeableUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UUPSUpgradeableUpgraded represents a Upgraded event raised by the UUPSUpgradeable contract.
type UUPSUpgradeableUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*UUPSUpgradeableUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &UUPSUpgradeableUpgradedIterator{contract: _UUPSUpgradeable.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *UUPSUpgradeableUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _UUPSUpgradeable.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UUPSUpgradeableUpgraded)
				if err := _UUPSUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_UUPSUpgradeable *UUPSUpgradeableFilterer) ParseUpgraded(log types.Log) (*UUPSUpgradeableUpgraded, error) {
	event := new(UUPSUpgradeableUpgraded)
	if err := _UUPSUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
