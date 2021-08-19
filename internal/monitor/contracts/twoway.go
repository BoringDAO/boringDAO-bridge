// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// SwapInParams is an auto generated low-level Go binding around an user-defined struct.
type SwapInParams struct {
	To           common.Address
	Amount1      *big.Int
	FeeAmountFix *big.Int
	RemainAmount *big.Int
	FeeToDev     common.Address
	ChainID      *big.Int
}

// AccessControlABI is the input ABI used to generate the binding from.
const AccessControlABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AccessControlFuncSigs maps the 4-byte function signature to its string representation.
var AccessControlFuncSigs = map[string]string{
	"a217fddf": "DEFAULT_ADMIN_ROLE()",
	"248a9ca3": "getRoleAdmin(bytes32)",
	"2f2ff15d": "grantRole(bytes32,address)",
	"91d14854": "hasRole(bytes32,address)",
	"36568abe": "renounceRole(bytes32,address)",
	"d547741f": "revokeRole(bytes32,address)",
	"01ffc9a7": "supportsInterface(bytes4)",
}

// AccessControl is an auto generated Go binding around an Ethereum contract.
type AccessControl struct {
	AccessControlCaller     // Read-only binding to the contract
	AccessControlTransactor // Write-only binding to the contract
	AccessControlFilterer   // Log filterer for contract events
}

// AccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccessControlSession struct {
	Contract     *AccessControl    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccessControlCallerSession struct {
	Contract *AccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccessControlTransactorSession struct {
	Contract     *AccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccessControlRaw struct {
	Contract *AccessControl // Generic contract binding to access the raw methods on
}

// AccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccessControlCallerRaw struct {
	Contract *AccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// AccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccessControlTransactorRaw struct {
	Contract *AccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccessControl creates a new instance of AccessControl, bound to a specific deployed contract.
func NewAccessControl(address common.Address, backend bind.ContractBackend) (*AccessControl, error) {
	contract, err := bindAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AccessControl{AccessControlCaller: AccessControlCaller{contract: contract}, AccessControlTransactor: AccessControlTransactor{contract: contract}, AccessControlFilterer: AccessControlFilterer{contract: contract}}, nil
}

// NewAccessControlCaller creates a new read-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlCaller(address common.Address, caller bind.ContractCaller) (*AccessControlCaller, error) {
	contract, err := bindAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlCaller{contract: contract}, nil
}

// NewAccessControlTransactor creates a new write-only instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*AccessControlTransactor, error) {
	contract, err := bindAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccessControlTransactor{contract: contract}, nil
}

// NewAccessControlFilterer creates a new log filterer instance of AccessControl, bound to a specific deployed contract.
func NewAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*AccessControlFilterer, error) {
	contract, err := bindAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccessControlFilterer{contract: contract}, nil
}

// bindAccessControl binds a generic wrapper to an already deployed contract.
func bindAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.AccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.AccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AccessControl *AccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AccessControl *AccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AccessControl *AccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AccessControl.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AccessControl.Contract.DEFAULTADMINROLE(&_AccessControl.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControl.Contract.GetRoleAdmin(&_AccessControl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AccessControl *AccessControlCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AccessControl.Contract.GetRoleAdmin(&_AccessControl.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AccessControl *AccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AccessControl.Contract.HasRole(&_AccessControl.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AccessControl.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControl.Contract.SupportsInterface(&_AccessControl.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AccessControl *AccessControlCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AccessControl.Contract.SupportsInterface(&_AccessControl.CallOpts, interfaceId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.GrantRole(&_AccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RenounceRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AccessControl *AccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AccessControl.Contract.RevokeRole(&_AccessControl.TransactOpts, role, account)
}

// AccessControlRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AccessControl contract.
type AccessControlRoleAdminChangedIterator struct {
	Event *AccessControlRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleAdminChanged)
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
		it.Event = new(AccessControlRoleAdminChanged)
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
func (it *AccessControlRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleAdminChanged represents a RoleAdminChanged event raised by the AccessControl contract.
type AccessControlRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AccessControlRoleAdminChangedIterator, error) {

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

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleAdminChangedIterator{contract: _AccessControl.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AccessControl *AccessControlFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AccessControlRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleAdminChanged)
				if err := _AccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_AccessControl *AccessControlFilterer) ParseRoleAdminChanged(log types.Log) (*AccessControlRoleAdminChanged, error) {
	event := new(AccessControlRoleAdminChanged)
	if err := _AccessControl.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AccessControl contract.
type AccessControlRoleGrantedIterator struct {
	Event *AccessControlRoleGranted // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleGranted)
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
		it.Event = new(AccessControlRoleGranted)
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
func (it *AccessControlRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleGranted represents a RoleGranted event raised by the AccessControl contract.
type AccessControlRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlRoleGrantedIterator, error) {

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

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleGrantedIterator{contract: _AccessControl.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AccessControlRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleGranted)
				if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_AccessControl *AccessControlFilterer) ParseRoleGranted(log types.Log) (*AccessControlRoleGranted, error) {
	event := new(AccessControlRoleGranted)
	if err := _AccessControl.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AccessControlRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AccessControl contract.
type AccessControlRoleRevokedIterator struct {
	Event *AccessControlRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AccessControlRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccessControlRoleRevoked)
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
		it.Event = new(AccessControlRoleRevoked)
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
func (it *AccessControlRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccessControlRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccessControlRoleRevoked represents a RoleRevoked event raised by the AccessControl contract.
type AccessControlRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AccessControlRoleRevokedIterator, error) {

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

	logs, sub, err := _AccessControl.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AccessControlRoleRevokedIterator{contract: _AccessControl.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AccessControl *AccessControlFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AccessControlRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AccessControl.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccessControlRoleRevoked)
				if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_AccessControl *AccessControlFilterer) ParseRoleRevoked(log types.Log) (*AccessControlRoleRevoked, error) {
	event := new(AccessControlRoleRevoked)
	if err := _AccessControl.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122049d99c19d14b1a6174d62de806cb3ffc3418dc4a5fd612040a58deb48132c57564736f6c63430008000033"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ERC165ABI is the input ABI used to generate the binding from.
const ERC165ABI = "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC165FuncSigs maps the 4-byte function signature to its string representation.
var ERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// ERC165 is an auto generated Go binding around an Ethereum contract.
type ERC165 struct {
	ERC165Caller     // Read-only binding to the contract
	ERC165Transactor // Write-only binding to the contract
	ERC165Filterer   // Log filterer for contract events
}

// ERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC165Session struct {
	Contract     *ERC165           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC165CallerSession struct {
	Contract *ERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC165TransactorSession struct {
	Contract     *ERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC165Raw struct {
	Contract *ERC165 // Generic contract binding to access the raw methods on
}

// ERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC165CallerRaw struct {
	Contract *ERC165Caller // Generic read-only contract binding to access the raw methods on
}

// ERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC165TransactorRaw struct {
	Contract *ERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC165 creates a new instance of ERC165, bound to a specific deployed contract.
func NewERC165(address common.Address, backend bind.ContractBackend) (*ERC165, error) {
	contract, err := bindERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC165{ERC165Caller: ERC165Caller{contract: contract}, ERC165Transactor: ERC165Transactor{contract: contract}, ERC165Filterer: ERC165Filterer{contract: contract}}, nil
}

// NewERC165Caller creates a new read-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Caller(address common.Address, caller bind.ContractCaller) (*ERC165Caller, error) {
	contract, err := bindERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Caller{contract: contract}, nil
}

// NewERC165Transactor creates a new write-only instance of ERC165, bound to a specific deployed contract.
func NewERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC165Transactor, error) {
	contract, err := bindERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC165Transactor{contract: contract}, nil
}

// NewERC165Filterer creates a new log filterer instance of ERC165, bound to a specific deployed contract.
func NewERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC165Filterer, error) {
	contract, err := bindERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC165Filterer{contract: contract}, nil
}

// bindERC165 binds a generic wrapper to an already deployed contract.
func bindERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.ERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.ERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC165 *ERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC165 *ERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC165 *ERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ERC165 *ERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ERC165.Contract.SupportsInterface(&_ERC165.CallOpts, interfaceId)
}

// IAccessControlABI is the input ABI used to generate the binding from.
const IAccessControlABI = "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IAccessControlFuncSigs maps the 4-byte function signature to its string representation.
var IAccessControlFuncSigs = map[string]string{
	"248a9ca3": "getRoleAdmin(bytes32)",
	"2f2ff15d": "grantRole(bytes32,address)",
	"91d14854": "hasRole(bytes32,address)",
	"36568abe": "renounceRole(bytes32,address)",
	"d547741f": "revokeRole(bytes32,address)",
}

// IAccessControl is an auto generated Go binding around an Ethereum contract.
type IAccessControl struct {
	IAccessControlCaller     // Read-only binding to the contract
	IAccessControlTransactor // Write-only binding to the contract
	IAccessControlFilterer   // Log filterer for contract events
}

// IAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccessControlSession struct {
	Contract     *IAccessControl   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccessControlCallerSession struct {
	Contract *IAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccessControlTransactorSession struct {
	Contract     *IAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccessControlRaw struct {
	Contract *IAccessControl // Generic contract binding to access the raw methods on
}

// IAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccessControlCallerRaw struct {
	Contract *IAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// IAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccessControlTransactorRaw struct {
	Contract *IAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccessControl creates a new instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControl(address common.Address, backend bind.ContractBackend) (*IAccessControl, error) {
	contract, err := bindIAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccessControl{IAccessControlCaller: IAccessControlCaller{contract: contract}, IAccessControlTransactor: IAccessControlTransactor{contract: contract}, IAccessControlFilterer: IAccessControlFilterer{contract: contract}}, nil
}

// NewIAccessControlCaller creates a new read-only instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlCaller(address common.Address, caller bind.ContractCaller) (*IAccessControlCaller, error) {
	contract, err := bindIAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlCaller{contract: contract}, nil
}

// NewIAccessControlTransactor creates a new write-only instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccessControlTransactor, error) {
	contract, err := bindIAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccessControlTransactor{contract: contract}, nil
}

// NewIAccessControlFilterer creates a new log filterer instance of IAccessControl, bound to a specific deployed contract.
func NewIAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccessControlFilterer, error) {
	contract, err := bindIAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccessControlFilterer{contract: contract}, nil
}

// bindIAccessControl binds a generic wrapper to an already deployed contract.
func bindIAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControl *IAccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControl.Contract.IAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControl *IAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.Contract.IAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControl *IAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControl.Contract.IAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccessControl *IAccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccessControl *IAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccessControl *IAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccessControl.Contract.contract.Transact(opts, method, params...)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControl.Contract.GetRoleAdmin(&_IAccessControl.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_IAccessControl *IAccessControlCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _IAccessControl.Contract.GetRoleAdmin(&_IAccessControl.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _IAccessControl.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControl.Contract.HasRole(&_IAccessControl.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_IAccessControl *IAccessControlCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _IAccessControl.Contract.HasRole(&_IAccessControl.CallOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.GrantRole(&_IAccessControl.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.GrantRole(&_IAccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RenounceRole(&_IAccessControl.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RenounceRole(&_IAccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRole(&_IAccessControl.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_IAccessControl *IAccessControlTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _IAccessControl.Contract.RevokeRole(&_IAccessControl.TransactOpts, role, account)
}

// IBoringTokenABI is the input ABI used to generate the binding from.
const IBoringTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IBoringTokenFuncSigs maps the 4-byte function signature to its string representation.
var IBoringTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"9dc29fac": "burn(address,uint256)",
	"40c10f19": "mint(address,uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IBoringToken is an auto generated Go binding around an Ethereum contract.
type IBoringToken struct {
	IBoringTokenCaller     // Read-only binding to the contract
	IBoringTokenTransactor // Write-only binding to the contract
	IBoringTokenFilterer   // Log filterer for contract events
}

// IBoringTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBoringTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBoringTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBoringTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBoringTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBoringTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBoringTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBoringTokenSession struct {
	Contract     *IBoringToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBoringTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBoringTokenCallerSession struct {
	Contract *IBoringTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IBoringTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBoringTokenTransactorSession struct {
	Contract     *IBoringTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IBoringTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBoringTokenRaw struct {
	Contract *IBoringToken // Generic contract binding to access the raw methods on
}

// IBoringTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBoringTokenCallerRaw struct {
	Contract *IBoringTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IBoringTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBoringTokenTransactorRaw struct {
	Contract *IBoringTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBoringToken creates a new instance of IBoringToken, bound to a specific deployed contract.
func NewIBoringToken(address common.Address, backend bind.ContractBackend) (*IBoringToken, error) {
	contract, err := bindIBoringToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBoringToken{IBoringTokenCaller: IBoringTokenCaller{contract: contract}, IBoringTokenTransactor: IBoringTokenTransactor{contract: contract}, IBoringTokenFilterer: IBoringTokenFilterer{contract: contract}}, nil
}

// NewIBoringTokenCaller creates a new read-only instance of IBoringToken, bound to a specific deployed contract.
func NewIBoringTokenCaller(address common.Address, caller bind.ContractCaller) (*IBoringTokenCaller, error) {
	contract, err := bindIBoringToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBoringTokenCaller{contract: contract}, nil
}

// NewIBoringTokenTransactor creates a new write-only instance of IBoringToken, bound to a specific deployed contract.
func NewIBoringTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IBoringTokenTransactor, error) {
	contract, err := bindIBoringToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBoringTokenTransactor{contract: contract}, nil
}

// NewIBoringTokenFilterer creates a new log filterer instance of IBoringToken, bound to a specific deployed contract.
func NewIBoringTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IBoringTokenFilterer, error) {
	contract, err := bindIBoringToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBoringTokenFilterer{contract: contract}, nil
}

// bindIBoringToken binds a generic wrapper to an already deployed contract.
func bindIBoringToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBoringTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBoringToken *IBoringTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBoringToken.Contract.IBoringTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBoringToken *IBoringTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBoringToken.Contract.IBoringTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBoringToken *IBoringTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBoringToken.Contract.IBoringTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBoringToken *IBoringTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBoringToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBoringToken *IBoringTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBoringToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBoringToken *IBoringTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBoringToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IBoringToken *IBoringTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBoringToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IBoringToken *IBoringTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IBoringToken.Contract.Allowance(&_IBoringToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IBoringToken *IBoringTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IBoringToken.Contract.Allowance(&_IBoringToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IBoringToken *IBoringTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IBoringToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IBoringToken *IBoringTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IBoringToken.Contract.BalanceOf(&_IBoringToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IBoringToken *IBoringTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IBoringToken.Contract.BalanceOf(&_IBoringToken.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IBoringToken *IBoringTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IBoringToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IBoringToken *IBoringTokenSession) TotalSupply() (*big.Int, error) {
	return _IBoringToken.Contract.TotalSupply(&_IBoringToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IBoringToken *IBoringTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _IBoringToken.Contract.TotalSupply(&_IBoringToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IBoringToken *IBoringTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IBoringToken *IBoringTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.Contract.Approve(&_IBoringToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IBoringToken *IBoringTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.Contract.Approve(&_IBoringToken.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_IBoringToken *IBoringTokenTransactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_IBoringToken *IBoringTokenSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.Contract.Burn(&_IBoringToken.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_IBoringToken *IBoringTokenTransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.Contract.Burn(&_IBoringToken.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IBoringToken *IBoringTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IBoringToken *IBoringTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.Contract.Mint(&_IBoringToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_IBoringToken *IBoringTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.Contract.Mint(&_IBoringToken.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IBoringToken *IBoringTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IBoringToken *IBoringTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.Contract.Transfer(&_IBoringToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IBoringToken *IBoringTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.Contract.Transfer(&_IBoringToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IBoringToken *IBoringTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IBoringToken *IBoringTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.Contract.TransferFrom(&_IBoringToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IBoringToken *IBoringTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IBoringToken.Contract.TransferFrom(&_IBoringToken.TransactOpts, sender, recipient, amount)
}

// IBoringTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IBoringToken contract.
type IBoringTokenApprovalIterator struct {
	Event *IBoringTokenApproval // Event containing the contract specifics and raw log

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
func (it *IBoringTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBoringTokenApproval)
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
		it.Event = new(IBoringTokenApproval)
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
func (it *IBoringTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBoringTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBoringTokenApproval represents a Approval event raised by the IBoringToken contract.
type IBoringTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IBoringToken *IBoringTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IBoringTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IBoringToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IBoringTokenApprovalIterator{contract: _IBoringToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IBoringToken *IBoringTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IBoringTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IBoringToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBoringTokenApproval)
				if err := _IBoringToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IBoringToken *IBoringTokenFilterer) ParseApproval(log types.Log) (*IBoringTokenApproval, error) {
	event := new(IBoringTokenApproval)
	if err := _IBoringToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IBoringTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IBoringToken contract.
type IBoringTokenTransferIterator struct {
	Event *IBoringTokenTransfer // Event containing the contract specifics and raw log

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
func (it *IBoringTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBoringTokenTransfer)
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
		it.Event = new(IBoringTokenTransfer)
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
func (it *IBoringTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBoringTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBoringTokenTransfer represents a Transfer event raised by the IBoringToken contract.
type IBoringTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IBoringToken *IBoringTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IBoringTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBoringToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IBoringTokenTransferIterator{contract: _IBoringToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IBoringToken *IBoringTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IBoringTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IBoringToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBoringTokenTransfer)
				if err := _IBoringToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IBoringToken *IBoringTokenFilterer) ParseTransfer(log types.Log) (*IBoringTokenTransfer, error) {
	event := new(IBoringTokenTransfer)
	if err := _IBoringToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC165ABI is the input ABI used to generate the binding from.
const IERC165ABI = "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IERC165FuncSigs maps the 4-byte function signature to its string representation.
var IERC165FuncSigs = map[string]string{
	"01ffc9a7": "supportsInterface(bytes4)",
}

// IERC165 is an auto generated Go binding around an Ethereum contract.
type IERC165 struct {
	IERC165Caller     // Read-only binding to the contract
	IERC165Transactor // Write-only binding to the contract
	IERC165Filterer   // Log filterer for contract events
}

// IERC165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC165Session struct {
	Contract     *IERC165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC165CallerSession struct {
	Contract *IERC165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IERC165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC165TransactorSession struct {
	Contract     *IERC165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IERC165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC165Raw struct {
	Contract *IERC165 // Generic contract binding to access the raw methods on
}

// IERC165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC165CallerRaw struct {
	Contract *IERC165Caller // Generic read-only contract binding to access the raw methods on
}

// IERC165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC165TransactorRaw struct {
	Contract *IERC165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC165 creates a new instance of IERC165, bound to a specific deployed contract.
func NewIERC165(address common.Address, backend bind.ContractBackend) (*IERC165, error) {
	contract, err := bindIERC165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC165{IERC165Caller: IERC165Caller{contract: contract}, IERC165Transactor: IERC165Transactor{contract: contract}, IERC165Filterer: IERC165Filterer{contract: contract}}, nil
}

// NewIERC165Caller creates a new read-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Caller(address common.Address, caller bind.ContractCaller) (*IERC165Caller, error) {
	contract, err := bindIERC165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Caller{contract: contract}, nil
}

// NewIERC165Transactor creates a new write-only instance of IERC165, bound to a specific deployed contract.
func NewIERC165Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC165Transactor, error) {
	contract, err := bindIERC165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC165Transactor{contract: contract}, nil
}

// NewIERC165Filterer creates a new log filterer instance of IERC165, bound to a specific deployed contract.
func NewIERC165Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC165Filterer, error) {
	contract, err := bindIERC165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC165Filterer{contract: contract}, nil
}

// bindIERC165 binds a generic wrapper to an already deployed contract.
func bindIERC165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC165ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.IERC165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.IERC165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC165 *IERC165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC165 *IERC165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC165 *IERC165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IERC165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IERC165 *IERC165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IERC165.Contract.SupportsInterface(&_IERC165.CallOpts, interfaceId)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

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

// ISwapPairABI is the input ABI used to generate the binding from.
const ISwapPairABI = "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"chainids\",\"type\":\"uint256[]\"}],\"name\":\"addChainIDs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"diff0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmountFix\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"feeToDev\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"internalType\":\"structSwapInParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"swapIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"chainid\",\"type\":\"uint256\"}],\"name\":\"swapOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"update\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISwapPairFuncSigs maps the 4-byte function signature to its string representation.
var ISwapPairFuncSigs = map[string]string{
	"19301e75": "addChainIDs(uint256[])",
	"779afa37": "burn(address,address,uint256,address,uint256)",
	"bcb67e14": "diff0()",
	"8ea7495c": "getReserves(uint256)",
	"6a627842": "mint(address)",
	"79f658ba": "swapIn((address,uint256,uint256,uint256,address,uint256))",
	"739dc0cf": "swapOut(address,uint256,uint256)",
	"0dfe1681": "token0()",
	"a2e62045": "update()",
}

// ISwapPair is an auto generated Go binding around an Ethereum contract.
type ISwapPair struct {
	ISwapPairCaller     // Read-only binding to the contract
	ISwapPairTransactor // Write-only binding to the contract
	ISwapPairFilterer   // Log filterer for contract events
}

// ISwapPairCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISwapPairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapPairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISwapPairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapPairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISwapPairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISwapPairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISwapPairSession struct {
	Contract     *ISwapPair        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISwapPairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISwapPairCallerSession struct {
	Contract *ISwapPairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ISwapPairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISwapPairTransactorSession struct {
	Contract     *ISwapPairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ISwapPairRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISwapPairRaw struct {
	Contract *ISwapPair // Generic contract binding to access the raw methods on
}

// ISwapPairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISwapPairCallerRaw struct {
	Contract *ISwapPairCaller // Generic read-only contract binding to access the raw methods on
}

// ISwapPairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISwapPairTransactorRaw struct {
	Contract *ISwapPairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISwapPair creates a new instance of ISwapPair, bound to a specific deployed contract.
func NewISwapPair(address common.Address, backend bind.ContractBackend) (*ISwapPair, error) {
	contract, err := bindISwapPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISwapPair{ISwapPairCaller: ISwapPairCaller{contract: contract}, ISwapPairTransactor: ISwapPairTransactor{contract: contract}, ISwapPairFilterer: ISwapPairFilterer{contract: contract}}, nil
}

// NewISwapPairCaller creates a new read-only instance of ISwapPair, bound to a specific deployed contract.
func NewISwapPairCaller(address common.Address, caller bind.ContractCaller) (*ISwapPairCaller, error) {
	contract, err := bindISwapPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapPairCaller{contract: contract}, nil
}

// NewISwapPairTransactor creates a new write-only instance of ISwapPair, bound to a specific deployed contract.
func NewISwapPairTransactor(address common.Address, transactor bind.ContractTransactor) (*ISwapPairTransactor, error) {
	contract, err := bindISwapPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISwapPairTransactor{contract: contract}, nil
}

// NewISwapPairFilterer creates a new log filterer instance of ISwapPair, bound to a specific deployed contract.
func NewISwapPairFilterer(address common.Address, filterer bind.ContractFilterer) (*ISwapPairFilterer, error) {
	contract, err := bindISwapPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISwapPairFilterer{contract: contract}, nil
}

// bindISwapPair binds a generic wrapper to an already deployed contract.
func bindISwapPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISwapPairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapPair *ISwapPairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapPair.Contract.ISwapPairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapPair *ISwapPairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapPair.Contract.ISwapPairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapPair *ISwapPairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapPair.Contract.ISwapPairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISwapPair *ISwapPairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISwapPair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISwapPair *ISwapPairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapPair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISwapPair *ISwapPairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISwapPair.Contract.contract.Transact(opts, method, params...)
}

// GetReserves is a free data retrieval call binding the contract method 0x8ea7495c.
//
// Solidity: function getReserves(uint256 chainid) view returns(uint256, uint256)
func (_ISwapPair *ISwapPairCaller) GetReserves(opts *bind.CallOpts, chainid *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ISwapPair.contract.Call(opts, &out, "getReserves", chainid)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetReserves is a free data retrieval call binding the contract method 0x8ea7495c.
//
// Solidity: function getReserves(uint256 chainid) view returns(uint256, uint256)
func (_ISwapPair *ISwapPairSession) GetReserves(chainid *big.Int) (*big.Int, *big.Int, error) {
	return _ISwapPair.Contract.GetReserves(&_ISwapPair.CallOpts, chainid)
}

// GetReserves is a free data retrieval call binding the contract method 0x8ea7495c.
//
// Solidity: function getReserves(uint256 chainid) view returns(uint256, uint256)
func (_ISwapPair *ISwapPairCallerSession) GetReserves(chainid *big.Int) (*big.Int, *big.Int, error) {
	return _ISwapPair.Contract.GetReserves(&_ISwapPair.CallOpts, chainid)
}

// AddChainIDs is a paid mutator transaction binding the contract method 0x19301e75.
//
// Solidity: function addChainIDs(uint256[] chainids) returns()
func (_ISwapPair *ISwapPairTransactor) AddChainIDs(opts *bind.TransactOpts, chainids []*big.Int) (*types.Transaction, error) {
	return _ISwapPair.contract.Transact(opts, "addChainIDs", chainids)
}

// AddChainIDs is a paid mutator transaction binding the contract method 0x19301e75.
//
// Solidity: function addChainIDs(uint256[] chainids) returns()
func (_ISwapPair *ISwapPairSession) AddChainIDs(chainids []*big.Int) (*types.Transaction, error) {
	return _ISwapPair.Contract.AddChainIDs(&_ISwapPair.TransactOpts, chainids)
}

// AddChainIDs is a paid mutator transaction binding the contract method 0x19301e75.
//
// Solidity: function addChainIDs(uint256[] chainids) returns()
func (_ISwapPair *ISwapPairTransactorSession) AddChainIDs(chainids []*big.Int) (*types.Transaction, error) {
	return _ISwapPair.Contract.AddChainIDs(&_ISwapPair.TransactOpts, chainids)
}

// Burn is a paid mutator transaction binding the contract method 0x779afa37.
//
// Solidity: function burn(address from, address to, uint256 amount, address feeTo, uint256 feeAmount) returns(uint256, uint256[], uint256[])
func (_ISwapPair *ISwapPairTransactor) Burn(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int, feeTo common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _ISwapPair.contract.Transact(opts, "burn", from, to, amount, feeTo, feeAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x779afa37.
//
// Solidity: function burn(address from, address to, uint256 amount, address feeTo, uint256 feeAmount) returns(uint256, uint256[], uint256[])
func (_ISwapPair *ISwapPairSession) Burn(from common.Address, to common.Address, amount *big.Int, feeTo common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _ISwapPair.Contract.Burn(&_ISwapPair.TransactOpts, from, to, amount, feeTo, feeAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x779afa37.
//
// Solidity: function burn(address from, address to, uint256 amount, address feeTo, uint256 feeAmount) returns(uint256, uint256[], uint256[])
func (_ISwapPair *ISwapPairTransactorSession) Burn(from common.Address, to common.Address, amount *big.Int, feeTo common.Address, feeAmount *big.Int) (*types.Transaction, error) {
	return _ISwapPair.Contract.Burn(&_ISwapPair.TransactOpts, from, to, amount, feeTo, feeAmount)
}

// Diff0 is a paid mutator transaction binding the contract method 0xbcb67e14.
//
// Solidity: function diff0() returns(uint256)
func (_ISwapPair *ISwapPairTransactor) Diff0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapPair.contract.Transact(opts, "diff0")
}

// Diff0 is a paid mutator transaction binding the contract method 0xbcb67e14.
//
// Solidity: function diff0() returns(uint256)
func (_ISwapPair *ISwapPairSession) Diff0() (*types.Transaction, error) {
	return _ISwapPair.Contract.Diff0(&_ISwapPair.TransactOpts)
}

// Diff0 is a paid mutator transaction binding the contract method 0xbcb67e14.
//
// Solidity: function diff0() returns(uint256)
func (_ISwapPair *ISwapPairTransactorSession) Diff0() (*types.Transaction, error) {
	return _ISwapPair.Contract.Diff0(&_ISwapPair.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_ISwapPair *ISwapPairTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _ISwapPair.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_ISwapPair *ISwapPairSession) Mint(to common.Address) (*types.Transaction, error) {
	return _ISwapPair.Contract.Mint(&_ISwapPair.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256 liquidity)
func (_ISwapPair *ISwapPairTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _ISwapPair.Contract.Mint(&_ISwapPair.TransactOpts, to)
}

// SwapIn is a paid mutator transaction binding the contract method 0x79f658ba.
//
// Solidity: function swapIn((address,uint256,uint256,uint256,address,uint256) params) returns()
func (_ISwapPair *ISwapPairTransactor) SwapIn(opts *bind.TransactOpts, params SwapInParams) (*types.Transaction, error) {
	return _ISwapPair.contract.Transact(opts, "swapIn", params)
}

// SwapIn is a paid mutator transaction binding the contract method 0x79f658ba.
//
// Solidity: function swapIn((address,uint256,uint256,uint256,address,uint256) params) returns()
func (_ISwapPair *ISwapPairSession) SwapIn(params SwapInParams) (*types.Transaction, error) {
	return _ISwapPair.Contract.SwapIn(&_ISwapPair.TransactOpts, params)
}

// SwapIn is a paid mutator transaction binding the contract method 0x79f658ba.
//
// Solidity: function swapIn((address,uint256,uint256,uint256,address,uint256) params) returns()
func (_ISwapPair *ISwapPairTransactorSession) SwapIn(params SwapInParams) (*types.Transaction, error) {
	return _ISwapPair.Contract.SwapIn(&_ISwapPair.TransactOpts, params)
}

// SwapOut is a paid mutator transaction binding the contract method 0x739dc0cf.
//
// Solidity: function swapOut(address to, uint256 amount, uint256 chainid) returns()
func (_ISwapPair *ISwapPairTransactor) SwapOut(opts *bind.TransactOpts, to common.Address, amount *big.Int, chainid *big.Int) (*types.Transaction, error) {
	return _ISwapPair.contract.Transact(opts, "swapOut", to, amount, chainid)
}

// SwapOut is a paid mutator transaction binding the contract method 0x739dc0cf.
//
// Solidity: function swapOut(address to, uint256 amount, uint256 chainid) returns()
func (_ISwapPair *ISwapPairSession) SwapOut(to common.Address, amount *big.Int, chainid *big.Int) (*types.Transaction, error) {
	return _ISwapPair.Contract.SwapOut(&_ISwapPair.TransactOpts, to, amount, chainid)
}

// SwapOut is a paid mutator transaction binding the contract method 0x739dc0cf.
//
// Solidity: function swapOut(address to, uint256 amount, uint256 chainid) returns()
func (_ISwapPair *ISwapPairTransactorSession) SwapOut(to common.Address, amount *big.Int, chainid *big.Int) (*types.Transaction, error) {
	return _ISwapPair.Contract.SwapOut(&_ISwapPair.TransactOpts, to, amount, chainid)
}

// Token0 is a paid mutator transaction binding the contract method 0x0dfe1681.
//
// Solidity: function token0() returns(address)
func (_ISwapPair *ISwapPairTransactor) Token0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapPair.contract.Transact(opts, "token0")
}

// Token0 is a paid mutator transaction binding the contract method 0x0dfe1681.
//
// Solidity: function token0() returns(address)
func (_ISwapPair *ISwapPairSession) Token0() (*types.Transaction, error) {
	return _ISwapPair.Contract.Token0(&_ISwapPair.TransactOpts)
}

// Token0 is a paid mutator transaction binding the contract method 0x0dfe1681.
//
// Solidity: function token0() returns(address)
func (_ISwapPair *ISwapPairTransactorSession) Token0() (*types.Transaction, error) {
	return _ISwapPair.Contract.Token0(&_ISwapPair.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_ISwapPair *ISwapPairTransactor) Update(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISwapPair.contract.Transact(opts, "update")
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_ISwapPair *ISwapPairSession) Update() (*types.Transaction, error) {
	return _ISwapPair.Contract.Update(&_ISwapPair.TransactOpts)
}

// Update is a paid mutator transaction binding the contract method 0xa2e62045.
//
// Solidity: function update() returns()
func (_ISwapPair *ISwapPairTransactorSession) Update() (*types.Transaction, error) {
	return _ISwapPair.Contract.Update(&_ISwapPair.TransactOpts)
}

// MathABI is the input ABI used to generate the binding from.
const MathABI = "[]"

// MathBin is the compiled bytecode used for deploying new contracts.
var MathBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220fb6fb08566f43c9096c1b8ba10fa7532ab86306b2ae85793bb854e40ef96a58564736f6c63430008000033"

// DeployMath deploys a new Ethereum contract, binding an instance of Math to it.
func DeployMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Math, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// Math is an auto generated Go binding around an Ethereum contract.
type Math struct {
	MathCaller     // Read-only binding to the contract
	MathTransactor // Write-only binding to the contract
	MathFilterer   // Log filterer for contract events
}

// MathCaller is an auto generated read-only Go binding around an Ethereum contract.
type MathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MathSession struct {
	Contract     *Math             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MathCallerSession struct {
	Contract *MathCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MathTransactorSession struct {
	Contract     *MathTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MathRaw is an auto generated low-level Go binding around an Ethereum contract.
type MathRaw struct {
	Contract *Math // Generic contract binding to access the raw methods on
}

// MathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MathCallerRaw struct {
	Contract *MathCaller // Generic read-only contract binding to access the raw methods on
}

// MathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MathTransactorRaw struct {
	Contract *MathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMath creates a new instance of Math, bound to a specific deployed contract.
func NewMath(address common.Address, backend bind.ContractBackend) (*Math, error) {
	contract, err := bindMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Math{MathCaller: MathCaller{contract: contract}, MathTransactor: MathTransactor{contract: contract}, MathFilterer: MathFilterer{contract: contract}}, nil
}

// NewMathCaller creates a new read-only instance of Math, bound to a specific deployed contract.
func NewMathCaller(address common.Address, caller bind.ContractCaller) (*MathCaller, error) {
	contract, err := bindMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MathCaller{contract: contract}, nil
}

// NewMathTransactor creates a new write-only instance of Math, bound to a specific deployed contract.
func NewMathTransactor(address common.Address, transactor bind.ContractTransactor) (*MathTransactor, error) {
	contract, err := bindMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MathTransactor{contract: contract}, nil
}

// NewMathFilterer creates a new log filterer instance of Math, bound to a specific deployed contract.
func NewMathFilterer(address common.Address, filterer bind.ContractFilterer) (*MathFilterer, error) {
	contract, err := bindMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MathFilterer{contract: contract}, nil
}

// bindMath binds a generic wrapper to an already deployed contract.
func bindMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.MathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.MathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Math *MathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Math.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Math *MathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Math.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Math *MathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Math.Contract.contract.Transact(opts, method, params...)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"715018a6": "renounceOwnership()",
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Ownable *OwnableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Ownable.Contract.RenounceOwnership(&_Ownable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalVoteABI is the input ABI used to generate the binding from.
const ProposalVoteABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ProposalVoteFuncSigs maps the 4-byte function signature to its string representation.
var ProposalVoteFuncSigs = map[string]string{
	"c86ec2bf": "threshold(address)",
}

// ProposalVoteBin is the compiled bytecode used for deploying new contracts.
var ProposalVoteBin = "0x608060405234801561001057600080fd5b5060cd8061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060285760003560e01c8063c86ec2bf14602d575b600080fd5b603c60383660046062565b6050565b60405160479190608e565b60405180910390f35b60006020819052908152604090205481565b6000602082840312156072578081fd5b81356001600160a01b03811681146087578182fd5b9392505050565b9081526020019056fea2646970667358221220030fc3379e0e2d06b14bbe46be3620b919c72c79bc07c8f00c386ec68666306964736f6c63430008000033"

// DeployProposalVote deploys a new Ethereum contract, binding an instance of ProposalVote to it.
func DeployProposalVote(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ProposalVote, error) {
	parsed, err := abi.JSON(strings.NewReader(ProposalVoteABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ProposalVoteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ProposalVote{ProposalVoteCaller: ProposalVoteCaller{contract: contract}, ProposalVoteTransactor: ProposalVoteTransactor{contract: contract}, ProposalVoteFilterer: ProposalVoteFilterer{contract: contract}}, nil
}

// ProposalVote is an auto generated Go binding around an Ethereum contract.
type ProposalVote struct {
	ProposalVoteCaller     // Read-only binding to the contract
	ProposalVoteTransactor // Write-only binding to the contract
	ProposalVoteFilterer   // Log filterer for contract events
}

// ProposalVoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProposalVoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalVoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProposalVoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalVoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProposalVoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalVoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProposalVoteSession struct {
	Contract     *ProposalVote     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProposalVoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProposalVoteCallerSession struct {
	Contract *ProposalVoteCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ProposalVoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProposalVoteTransactorSession struct {
	Contract     *ProposalVoteTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ProposalVoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProposalVoteRaw struct {
	Contract *ProposalVote // Generic contract binding to access the raw methods on
}

// ProposalVoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProposalVoteCallerRaw struct {
	Contract *ProposalVoteCaller // Generic read-only contract binding to access the raw methods on
}

// ProposalVoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProposalVoteTransactorRaw struct {
	Contract *ProposalVoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProposalVote creates a new instance of ProposalVote, bound to a specific deployed contract.
func NewProposalVote(address common.Address, backend bind.ContractBackend) (*ProposalVote, error) {
	contract, err := bindProposalVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ProposalVote{ProposalVoteCaller: ProposalVoteCaller{contract: contract}, ProposalVoteTransactor: ProposalVoteTransactor{contract: contract}, ProposalVoteFilterer: ProposalVoteFilterer{contract: contract}}, nil
}

// NewProposalVoteCaller creates a new read-only instance of ProposalVote, bound to a specific deployed contract.
func NewProposalVoteCaller(address common.Address, caller bind.ContractCaller) (*ProposalVoteCaller, error) {
	contract, err := bindProposalVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalVoteCaller{contract: contract}, nil
}

// NewProposalVoteTransactor creates a new write-only instance of ProposalVote, bound to a specific deployed contract.
func NewProposalVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*ProposalVoteTransactor, error) {
	contract, err := bindProposalVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalVoteTransactor{contract: contract}, nil
}

// NewProposalVoteFilterer creates a new log filterer instance of ProposalVote, bound to a specific deployed contract.
func NewProposalVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*ProposalVoteFilterer, error) {
	contract, err := bindProposalVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProposalVoteFilterer{contract: contract}, nil
}

// bindProposalVote binds a generic wrapper to an already deployed contract.
func bindProposalVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProposalVoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposalVote *ProposalVoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposalVote.Contract.ProposalVoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposalVote *ProposalVoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalVote.Contract.ProposalVoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposalVote *ProposalVoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposalVote.Contract.ProposalVoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ProposalVote *ProposalVoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ProposalVote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ProposalVote *ProposalVoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ProposalVote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ProposalVote *ProposalVoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ProposalVote.Contract.contract.Transact(opts, method, params...)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_ProposalVote *ProposalVoteCaller) Threshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ProposalVote.contract.Call(opts, &out, "threshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_ProposalVote *ProposalVoteSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _ProposalVote.Contract.Threshold(&_ProposalVote.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_ProposalVote *ProposalVoteCallerSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _ProposalVote.Contract.Threshold(&_ProposalVote.CallOpts, arg0)
}

// ProposalVoteProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the ProposalVote contract.
type ProposalVoteProposalVotedIterator struct {
	Event *ProposalVoteProposalVoted // Event containing the contract specifics and raw log

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
func (it *ProposalVoteProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalVoteProposalVoted)
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
		it.Event = new(ProposalVoteProposalVoted)
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
func (it *ProposalVoteProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalVoteProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalVoteProposalVoted represents a ProposalVoted event raised by the ProposalVote contract.
type ProposalVoteProposalVoted struct {
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
func (_ProposalVote *ProposalVoteFilterer) FilterProposalVoted(opts *bind.FilterOpts) (*ProposalVoteProposalVotedIterator, error) {

	logs, sub, err := _ProposalVote.contract.FilterLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return &ProposalVoteProposalVotedIterator{contract: _ProposalVote.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0xe458477b94285c9f254a407e2614c39f55230b527ff7f51d76fdf33e251ae936.
//
// Solidity: event ProposalVoted(address token, address from, address to, uint256 amount, address proposer, uint256 count, uint256 threshold)
func (_ProposalVote *ProposalVoteFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *ProposalVoteProposalVoted) (event.Subscription, error) {

	logs, sub, err := _ProposalVote.contract.WatchLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalVoteProposalVoted)
				if err := _ProposalVote.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
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
func (_ProposalVote *ProposalVoteFilterer) ParseProposalVoted(log types.Log) (*ProposalVoteProposalVoted, error) {
	event := new(ProposalVoteProposalVoted)
	if err := _ProposalVote.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalVoteThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the ProposalVote contract.
type ProposalVoteThresholdChangedIterator struct {
	Event *ProposalVoteThresholdChanged // Event containing the contract specifics and raw log

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
func (it *ProposalVoteThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalVoteThresholdChanged)
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
		it.Event = new(ProposalVoteThresholdChanged)
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
func (it *ProposalVoteThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalVoteThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalVoteThresholdChanged represents a ThresholdChanged event raised by the ProposalVote contract.
type ProposalVoteThresholdChanged struct {
	Token        common.Address
	OldThreshold *big.Int
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_ProposalVote *ProposalVoteFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*ProposalVoteThresholdChangedIterator, error) {

	logs, sub, err := _ProposalVote.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &ProposalVoteThresholdChangedIterator{contract: _ProposalVote.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_ProposalVote *ProposalVoteFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *ProposalVoteThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _ProposalVote.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalVoteThresholdChanged)
				if err := _ProposalVote.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_ProposalVote *ProposalVoteFilterer) ParseThresholdChanged(log types.Log) (*ProposalVoteThresholdChanged, error) {
	event := new(ProposalVoteThresholdChanged)
	if err := _ProposalVote.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SafeDecimalMathABI is the input ABI used to generate the binding from.
const SafeDecimalMathABI = "[{\"inputs\":[],\"name\":\"PRECISE_UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highPrecisionDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"preciseUnit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// SafeDecimalMathFuncSigs maps the 4-byte function signature to its string representation.
var SafeDecimalMathFuncSigs = map[string]string{
	"864029e7": "PRECISE_UNIT()",
	"9d8e2177": "UNIT()",
	"313ce567": "decimals()",
	"def4419d": "highPrecisionDecimals()",
	"d5e5e6e6": "preciseUnit()",
	"907af6c0": "unit()",
}

// SafeDecimalMathBin is the compiled bytecode used for deploying new contracts.
var SafeDecimalMathBin = "0x61028d61003a600b82828239805160001a60731461002d57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061006c5760003560e01c8063313ce56714610071578063864029e71461008f578063907af6c0146100a45780639d8e2177146100ac578063d5e5e6e6146100b4578063def4419d146100bc575b600080fd5b6100796100c4565b6040516100869190610113565b60405180910390f35b6100976100c9565b604051610086919061010a565b6100976100d8565b6100976100eb565b6100976100f7565b610079610105565b601281565b6100d5601b600a610167565b81565b60006100e66012600a610167565b905090565b6100d56012600a610167565b60006100e6601b600a610167565b601b81565b90815260200190565b60ff91909116815260200190565b80825b6001808611610133575061015e565b81870482111561014557610145610241565b8086161561015257918102915b9490941c938002610124565b94509492505050565b6000610176600019848461017d565b9392505050565b60008261018c57506001610176565b8161019957506000610176565b81600181146101af57600281146101b9576101e6565b6001915050610176565b60ff8411156101ca576101ca610241565b6001841b9150848211156101e0576101e0610241565b50610176565b5060208310610133831016604e8410600b8410161715610219575081810a8381111561021457610214610241565b610176565b6102268484846001610121565b80860482111561023857610238610241565b02949350505050565b634e487b7160e01b600052601160045260246000fdfea2646970667358221220d72eaf15dcdfe746c219b695de330cefcc606ad3970883e2b2720929eb25ad4b64736f6c63430008000033"

// DeploySafeDecimalMath deploys a new Ethereum contract, binding an instance of SafeDecimalMath to it.
func DeploySafeDecimalMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeDecimalMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeDecimalMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeDecimalMathBin), backend)
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

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
var SafeERC20Bin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220cfb1ce2d51d20c6f2700bb5dc0b4c5d2214acb8018601a5f342c5173a793900b64736f6c63430008000033"

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b61e01022d7bba09a2e793bc1fcf07ab2b7d6faf7a362d777b6eb6b7cf6e9d1b64736f6c63430008000033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
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

// StringsABI is the input ABI used to generate the binding from.
const StringsABI = "[]"

// StringsBin is the compiled bytecode used for deploying new contracts.
var StringsBin = "0x60566037600b82828239805160001a607314602a57634e487b7160e01b600052600060045260246000fd5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220af7217fd26198da328f3e6ddfc2ed1cd69c4c6e8a7f17a869009d88196558e8c64736f6c63430008000033"

// DeployStrings deploys a new Ethereum contract, binding an instance of Strings to it.
func DeployStrings(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Strings, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StringsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// Strings is an auto generated Go binding around an Ethereum contract.
type Strings struct {
	StringsCaller     // Read-only binding to the contract
	StringsTransactor // Write-only binding to the contract
	StringsFilterer   // Log filterer for contract events
}

// StringsCaller is an auto generated read-only Go binding around an Ethereum contract.
type StringsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StringsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StringsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StringsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StringsSession struct {
	Contract     *Strings          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StringsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StringsCallerSession struct {
	Contract *StringsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StringsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StringsTransactorSession struct {
	Contract     *StringsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StringsRaw is an auto generated low-level Go binding around an Ethereum contract.
type StringsRaw struct {
	Contract *Strings // Generic contract binding to access the raw methods on
}

// StringsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StringsCallerRaw struct {
	Contract *StringsCaller // Generic read-only contract binding to access the raw methods on
}

// StringsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StringsTransactorRaw struct {
	Contract *StringsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStrings creates a new instance of Strings, bound to a specific deployed contract.
func NewStrings(address common.Address, backend bind.ContractBackend) (*Strings, error) {
	contract, err := bindStrings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Strings{StringsCaller: StringsCaller{contract: contract}, StringsTransactor: StringsTransactor{contract: contract}, StringsFilterer: StringsFilterer{contract: contract}}, nil
}

// NewStringsCaller creates a new read-only instance of Strings, bound to a specific deployed contract.
func NewStringsCaller(address common.Address, caller bind.ContractCaller) (*StringsCaller, error) {
	contract, err := bindStrings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StringsCaller{contract: contract}, nil
}

// NewStringsTransactor creates a new write-only instance of Strings, bound to a specific deployed contract.
func NewStringsTransactor(address common.Address, transactor bind.ContractTransactor) (*StringsTransactor, error) {
	contract, err := bindStrings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StringsTransactor{contract: contract}, nil
}

// NewStringsFilterer creates a new log filterer instance of Strings, bound to a specific deployed contract.
func NewStringsFilterer(address common.Address, filterer bind.ContractFilterer) (*StringsFilterer, error) {
	contract, err := bindStrings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StringsFilterer{contract: contract}, nil
}

// bindStrings binds a generic wrapper to an already deployed contract.
func bindStrings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StringsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.StringsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.StringsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Strings *StringsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Strings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Strings *StringsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Strings *StringsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Strings.Contract.contract.Transact(opts, method, params...)
}

// TwoWayABI is the input ABI used to generate the binding from.
const TwoWayABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToDev\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CrossBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeRatio\",\"type\":\"uint256\"}],\"name\":\"FeeChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeeToChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeeToRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Lock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"Rollback\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"Rollbacked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"ThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chianID0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"Unlock\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"CROSSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"chainIDs\",\"type\":\"uint256[]\"}],\"name\":\"addChainIDs\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"chainIDs\",\"type\":\"uint256[]\"}],\"name\":\"addPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"addSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"token0s\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"token1s\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chainIDs\",\"type\":\"uint256[]\"}],\"name\":\"addSupportTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmountFix\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmountRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateRemoveFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"crossIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"crossOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeAmountM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeRatioM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToDev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxToken0AmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"getMaxToken1AmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pairs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"removeFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lpAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"chainids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount1s\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removePair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"}],\"name\":\"removeSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"token0s\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"chainIDs\",\"type\":\"uint256[]\"}],\"name\":\"removeSupportTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"rollback\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRatio\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setFeeToDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feeAmount\",\"type\":\"uint256\"}],\"name\":\"setRemoveFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_inactived\",\"type\":\"bool\"}],\"name\":\"setUnlockFeeOn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"supportToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"threshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"txMinted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"txRollbacked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"txUnlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unlockFeeOn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TwoWayFuncSigs maps the 4-byte function signature to its string representation.
var TwoWayFuncSigs = map[string]string{
	"56cf02d9": "CROSSER_ROLE()",
	"a217fddf": "DEFAULT_ADMIN_ROLE()",
	"27f8dd9b": "addChainIDs(address,uint256[])",
	"1b879378": "addLiquidity(address,uint256,address)",
	"31b9e5a3": "addPair(address,address,uint256[])",
	"1b322be4": "addSupportToken(address,address,uint256)",
	"ee8a6acf": "addSupportTokens(address[],address[],uint256[])",
	"0324ef9c": "calculateFee(address,uint256,uint256)",
	"2d63c99a": "calculateRemoveFee(address,uint256)",
	"cf28f67c": "crossIn(address,uint256,address,address,uint256,string)",
	"df7e600a": "crossOut(address,uint256,address,uint256)",
	"64078943": "feeAmountM(address,uint256)",
	"9c8d70be": "feeRatioM(address,uint256)",
	"db697be4": "feeTo(address,uint256)",
	"100be3bd": "feeToDev()",
	"934ac707": "getMaxToken0AmountOut(address,uint256)",
	"56ecaeb7": "getMaxToken1AmountOut(address,uint256)",
	"248a9ca3": "getRoleAdmin(bytes32)",
	"2f2ff15d": "grantRole(bytes32,address)",
	"91d14854": "hasRole(bytes32,address)",
	"fe33b302": "pairs(address)",
	"0b19993f": "removeFeeAmount(address)",
	"47fdbc8e": "removeLiquidity(address,uint256,address)",
	"af6c9c1d": "removePair(address)",
	"f8f8d5c0": "removeSupportToken(address,uint256)",
	"03507ba5": "removeSupportTokens(address[],uint256[])",
	"36568abe": "renounceRole(bytes32,address)",
	"d547741f": "revokeRole(bytes32,address)",
	"81660f5d": "rollback(address,uint256,address,uint256,string)",
	"f06410f8": "setFee(address,uint256,uint256,uint256)",
	"c369c773": "setFeeToDev(address)",
	"3e1cd7e4": "setRemoveFee(address,uint256)",
	"9d879990": "setThreshold(address,uint256)",
	"98d2f6b4": "setUnlockFeeOn(address,uint256,bool)",
	"d692f4c5": "supportToken(address,uint256)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"c86ec2bf": "threshold(address)",
	"10c27402": "txMinted(string)",
	"53ad72e5": "txRollbacked(string)",
	"967145af": "txUnlocked(string)",
	"715ec45c": "unlock(address,uint256,address,address,uint256,string)",
	"68b1188f": "unlockFeeOn(address,uint256)",
}

// TwoWayBin is the compiled bytecode used for deploying new contracts.
var TwoWayBin = "0x60806040523480156200001157600080fd5b5060405162003f1b38038062003f1b83398101604081905262000034916200012e565b600980546001600160a01b0319166001600160a01b0383161790556200005c60003362000063565b506200015e565b6200006f828262000073565b5050565b6200007f8282620000ff565b6200006f5760008281526004602090815260408083206001600160a01b03851684529091529020805460ff19166001179055620000bb6200012a565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b60009182526004602090815260408084206001600160a01b0393909316845291905290205460ff1690565b3390565b60006020828403121562000140578081fd5b81516001600160a01b038116811462000157578182fd5b9392505050565b613dad806200016e6000396000f3fe608060405234801561001057600080fd5b506004361061025e5760003560e01c8063715ec45c11610146578063c369c773116100c3578063db697be411610087578063db697be41461054c578063df7e600a1461055f578063ee8a6acf14610572578063f06410f814610585578063f8f8d5c014610598578063fe33b302146105ab5761025e565b8063c369c773146104ed578063c86ec2bf14610500578063cf28f67c14610513578063d547741f14610526578063d692f4c5146105395761025e565b806398d2f6b41161010a57806398d2f6b4146104995780639c8d70be146104ac5780639d879990146104bf578063a217fddf146104d2578063af6c9c1d146104da5761025e565b8063715ec45c1461043a57806381660f5d1461044d57806391d1485414610460578063934ac70714610473578063967145af146104865761025e565b80632d63c99a116101df57806347fdbc8e116101a357806347fdbc8e146103c457806353ad72e5146103e657806356cf02d9146103f957806356ecaeb714610401578063640789431461041457806368b1188f146104275761025e565b80632d63c99a146103575780632f2ff15d1461037857806331b9e5a31461038b57806336568abe1461039e5780633e1cd7e4146103b15761025e565b806310c274021161022657806310c27402146102f85780631b322be41461030b5780631b8793781461031e578063248a9ca31461033157806327f8dd9b146103445761025e565b806301ffc9a7146102635780630324ef9c1461028c57806303507ba5146102ae5780630b19993f146102c3578063100be3bd146102e3575b600080fd5b610276610271366004612ffd565b6105be565b604051610283919061343a565b60405180910390f35b61029f61029a366004612e6c565b6105eb565b60405161028393929190613af9565b6102c16102bc366004612f50565b610667565b005b6102d66102d1366004612b9a565b61070c565b6040516102839190613445565b6102eb61071e565b6040516102839190613243565b610276610306366004613025565b61072d565b6102c1610319366004612c10565b61074d565b6102d661032c366004612cc0565b610810565b6102d661033f366004612fc3565b6108e6565b6102c1610352366004612c4b565b6108fb565b61036a610365366004612c97565b6109bd565b604051610283929190613aeb565b6102c1610386366004612fdb565b610a24565b6102c1610399366004612bb4565b610a48565b6102c16103ac366004612fdb565b610af5565b6102c16103bf366004612c97565b610b3b565b6103d76103d2366004612cc0565b610b6c565b60405161028393929190613ac0565b6102766103f4366004613025565b610d2f565b6102d6610d4f565b6102d661040f366004612c97565b610d62565b6102d6610422366004612c97565b610dfe565b610276610435366004612c97565b610e1b565b6102c1610448366004612cfb565b610e3b565b6102c161045b366004612dbd565b611108565b61027661046e366004612fdb565b6113b3565b6102d6610481366004612c97565b6113de565b610276610494366004613025565b61147a565b6102c16104a7366004612e2d565b61149a565b6102d66104ba366004612c97565b611588565b6102c16104cd366004612c97565b6115a5565b6102d66115d6565b6102c16104e8366004612b9a565b6115db565b6102c16104fb366004612b9a565b611660565b6102d661050e366004612b9a565b611693565b6102c1610521366004612cfb565b6116a5565b6102c1610534366004612fdb565b61195b565b6102eb610547366004612c97565b61197a565b6102eb61055a366004612c97565b6119a0565b6102c161056d366004612d7a565b6119c6565b6102c1610580366004612ed6565b611db9565b6102c1610593366004612e9e565b611e9f565b6102c16105a6366004612c97565b611ed2565b6102eb6105b9366004612b9a565b611f6e565b60006001600160e01b03198216637965db0b60e01b14806105e357506105e382611f89565b90505b919050565b6001600160a01b038316600081815260056020908152604080832086845282528083205493835260068252808320868452909152812054819061062e8582611fa2565b925061063a8385613b5d565b85101561064657600093505b826106518587613cc8565b61065b9190613cc8565b91505093509350939050565b80518251146106915760405162461bcd60e51b8152600401610688906134e4565b60405180910390fd5b60005b8251811015610707576106f58382815181106106c057634e487b7160e01b600052603260045260246000fd5b60200260200101518383815181106106e857634e487b7160e01b600052603260045260246000fd5b6020026020010151611ed2565b806106ff81613d22565b915050610694565b505050565b60076020526000908152604090205481565b6009546001600160a01b031681565b8051602081830181018051600b8252928201919093012091525460ff1681565b6107586000336113b3565b6107745760405162461bcd60e51b8152600401610688906137b9565b6001600160a01b038381166000908152600a6020908152604080832085845290915290205416156107b75760405162461bcd60e51b8152600401610688906135b1565b6001600160a01b039283166000818152600a60209081526040808320858452825280832080546001600160a01b0319169690971695909517909555908152600f845282812091815292529020805460ff19166001179055565b6001600160a01b038084166000908152600e60205260408120549091168061084a5760405162461bcd60e51b8152600401610688906136f1565b61085f6001600160a01b038616338387611fcb565b6040516335313c2160e11b81526001600160a01b03821690636a6278429061088b908690600401613243565b602060405180830381600087803b1580156108a557600080fd5b505af11580156108b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108dd9190613058565b95945050505050565b60009081526004602052604090206001015490565b6109066000336113b3565b6109225760405162461bcd60e51b8152600401610688906137b9565b6001600160a01b038083166000908152600e6020526040902054168061095a5760405162461bcd60e51b81526004016106889061354f565b6040516319301e7560e01b81526001600160a01b038216906319301e7590610986908590600401613427565b600060405180830381600087803b1580156109a057600080fd5b505af11580156109b4573d6000803e3d6000fd5b50505050505050565b6001600160a01b038216600090815260076020526040812054819083116109f65760405162461bcd60e51b81526004016106889061361f565b6001600160a01b0384166000908152600760205260409020549150610a1b8284613cc8565b90509250929050565b610a2d826108e6565b610a3e81610a39612023565b612027565b610707838361208b565b610a536000336113b3565b610a6f5760405162461bcd60e51b8152600401610688906137b9565b6001600160a01b038381166000908152600e60205260409020541615610aa75760405162461bcd60e51b81526004016106889061374b565b6001600160a01b038381166000908152600e60205260409081902080546001600160a01b0319169285169283179055516319301e7560e01b81526319301e7590610986908490600401613427565b610afd612023565b6001600160a01b0316816001600160a01b031614610b2d5760405162461bcd60e51b8152600401610688906139f5565b610b378282612112565b5050565b610b466000336113b3565b610b625760405162461bcd60e51b8152600401610688906137b9565b610b378282612197565b600060608060008511610b915760405162461bcd60e51b8152600401610688906134c3565b6001600160a01b038087166000908152600e60205260409020541680610bc95760405162461bcd60e51b8152600401610688906136f1565b6040516370a0823160e01b81526000906001600160a01b038316906370a0823190610bf8903390600401613243565b60206040518083038186803b158015610c1057600080fd5b505afa158015610c24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c489190613058565b905086811015610c6a5760405162461bcd60e51b8152600401610688906137f0565b6001600160a01b038089166000908152600760205260409081902054600954915163779afa3760e01b815290928581169263779afa3792610cba9233928d928f92919091169088906004016132be565b600060405180830381600087803b158015610cd457600080fd5b505af1158015610ce8573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610d109190810190613070565b91975095509350610d23898887876121b3565b50505093509350939050565b8051602081830181018051600d8252928201919093012091525460ff1681565b6b43524f535345525f524f4c4560a01b81565b6001600160a01b038083166000908152600e60205260408082205490516323a9d25760e21b81529192169082908290638ea7495c90610da5908790600401613445565b604080518083038186803b158015610dbc57600080fd5b505afa158015610dd0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610df491906130d0565b9695505050505050565b600560209081526000928352604080842090915290825290205481565b600f60209081526000928352604080842090915290825290205460ff1681565b6001600160a01b038087166000908152600a602090815260408083208984529091529020548791879116610e815760405162461bcd60e51b815260040161068890613817565b610e9a6b43524f535345525f524f4c4560a01b336113b3565b610eb65760405162461bcd60e51b81526004016106889061357a565b82600c81604051610ec791906131b2565b9081526040519081900360200190205460ff1615610ef75760405162461bcd60e51b815260040161068890613964565b6000610f068a89898989612300565b6001600160a01b03808c166000908152600e60209081526040808320548151632f2d9f8560e21b815291519596509294929093169263bcb67e14926004808301939282900301818787803b158015610f5d57600080fd5b505af1158015610f71573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f959190613058565b610f9f9088613b75565b905081156110fb576001600c87604051610fb991906131b2565b9081526040805160209281900383019020805460ff1916931515939093179092556001600160a01b038d166000908152600f82528281208d8252909152205460ff16156110795761100c8b8b838b6124ba565b6001600160a01b03808c166000908152600e602052604080822054815163a2e6204560e01b8152915193169263a2e620459260048084019391929182900301818387803b15801561105c57600080fd5b505af1158015611070573d6000803e3d6000fd5b5050505061108d565b61108d6001600160a01b038c168983612526565b6001600160a01b03808c166000908152600a602090815260408083208e8452909152908190205490517fe8a9cddb11d86358ad5c2fdd6359f0a6f41de0d399033319becf1364409dacbc926110f2928f9291169046908f908f908f9089908f90613366565b60405180910390a15b5050505050505050505050565b6001600160a01b038086166000908152600a60209081526040808320888452909152902054869186911661114e5760405162461bcd60e51b815260040161068890613817565b6111676b43524f535345525f524f4c4560a01b336113b3565b6111835760405162461bcd60e51b81526004016106889061357a565b82600d8160405161119491906131b2565b9081526040519081900360200190205460ff16156111c45760405162461bcd60e51b81526004016106889061371c565b60006111d38988898989612300565b905080156113a8576001600d866040516111ed91906131b2565b9081526040805160209281900383018120805460ff1916941515949094179093556001600160a01b038c81166000908152600e845282812054632f2d9f8560e21b865292516112b6958d95949093169363bcb67e1493600480820194929392918390030190829087803b15801561126357600080fd5b505af1158015611277573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061129b9190613058565b6112a59089613b75565b6001600160a01b038c169190612526565b7f9420db1a0c5e4f45cc0dac05f17e8eb893645981e3c01e8d14120353a487d9e98988600e60008d6001600160a01b03166001600160a01b0316815260200190815260200160002060009054906101000a90046001600160a01b03166001600160a01b031663bcb67e146040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561134c57600080fd5b505af1158015611360573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113849190613058565b61138e908a613b75565b8860405161139f94939291906132f1565b60405180910390a15b505050505050505050565b60009182526004602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6001600160a01b038083166000908152600e60205260408082205490516323a9d25760e21b81529192169082908290638ea7495c90611421908790600401613445565b604080518083038186803b15801561143857600080fd5b505afa15801561144c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061147091906130d0565b5095945050505050565b8051602081830181018051600c8252928201919093012091525460ff1681565b6114a56000336113b3565b6114c15760405162461bcd60e51b8152600401610688906137b9565b6001600160a01b038084166000908152600a6020908152604080832086845290915290205484918491166115075760405162461bcd60e51b815260040161068890613817565b6001600160a01b0385166000908152600f6020908152604080832087845290915290205460ff16151583151514156115515760405162461bcd60e51b815260040161068890613525565b50506001600160a01b03929092166000908152600f6020908152604080832093835292905220805460ff1916911515919091179055565b600660209081526000928352604080842090915290825290205481565b6115b06000336113b3565b6115cc5760405162461bcd60e51b8152600401610688906137b9565b610b378282612545565b600081565b6115e66000336113b3565b6116025760405162461bcd60e51b8152600401610688906137b9565b6001600160a01b038181166000908152600e6020526040902054166116395760405162461bcd60e51b815260040161068890613461565b6001600160a01b03166000908152600e6020526040902080546001600160a01b0319169055565b61166b6000336113b3565b6116875760405162461bcd60e51b8152600401610688906137b9565b611690816125a5565b50565b60006020819052908152604090205481565b6116be6b43524f535345525f524f4c4560a01b336113b3565b6116da5760405162461bcd60e51b81526004016106889061357a565b80600b816040516116eb91906131b2565b9081526040519081900360200190205460ff161561171b5760405162461bcd60e51b815260040161068890613649565b600061172a8887878787612300565b90508015611951576001600b8460405161174491906131b2565b9081526040805160209281900383018120805460ff1916941515949094179093556001600160a01b038b81166000908152600e845282812054632f2d9f8560e21b8652925192909116939092849263bcb67e1492600480820193929182900301818787803b1580156117b557600080fd5b505af11580156117c9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117ed9190613058565b6117f79087613b75565b905060006118058b8b6113de565b905080821115611881576001600160a01b03808c166000908152600a602090815260408083208e8452909152908190205490517f56a316ee63540abb5f8fd44be7a6eb2016a90fc9bb3a34ffe13948a44c49a61692611874928f9291169046908f908f908f908f908f90613366565b60405180910390a16110fb565b60008061188f8d8d866105eb565b925050915060006040518060c001604052808c6001600160a01b03168152602001868152602001848152602001838152602001600960009054906101000a90046001600160a01b03166001600160a01b031681526020018e8152509050856001600160a01b03166379f658ba826040518263ffffffff1660e01b81526004016119189190613a73565b600060405180830381600087803b15801561193257600080fd5b505af1158015611946573d6000803e3d6000fd5b505050505050505050505b5050505050505050565b611964826108e6565b61197081610a39612023565b6107078383612112565b600a6020908152600092835260408084209091529082529020546001600160a01b031681565b60086020908152600092835260408084209091529082529020546001600160a01b031681565b6001600160a01b038085166000908152600a602090815260408083208784529091529020548591859116611a0c5760405162461bcd60e51b815260040161068890613817565b60008311611a2c5760405162461bcd60e51b81526004016106889061391f565b6001600160a01b038416611a525760405162461bcd60e51b815260040161068890613991565b6001600160a01b038087166000908152600e60209081526040808320548151632f2d9f8560e21b81529151941693849263bcb67e14926004808201939182900301818787803b158015611aa457600080fd5b505af1158015611ab8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611adc9190613058565b611ae68989610d62565b611af09190613b75565b90506000611afe86836125c7565b90508015611c8657611b1b6001600160a01b038a16338584611fcb565b60405163739dc0cf60e01b81526001600160a01b0384169063739dc0cf90611b4b908a9085908d906004016133e0565b600060405180830381600087803b158015611b6557600080fd5b505af1158015611b79573d6000803e3d6000fd5b505050507f79aee25f54411be6ea1ab53c9d3cc3e245c301d481091deda6343da926932c0c89600a60008c6001600160a01b03166001600160a01b0316815260200190815260200160002060008b815260200190815260200160002060009054906101000a90046001600160a01b0316468b338c896001600160a01b031663bcb67e146040518163ffffffff1660e01b8152600401602060405180830381600087803b158015611c2857600080fd5b505af1158015611c3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c609190613058565b611c6a9089613ca9565b604051611c7d9796959493929190613324565b60405180910390a15b818611156113a857611caf3330611c9d848a613cc8565b6001600160a01b038d16929190611fcb565b7f7cb5e2d54d5587c3a3448631884061009b0e2c30e37922a4dd4aed50e11dd7f389600a60008c6001600160a01b03166001600160a01b0316815260200190815260200160002060008b815260200190815260200160002060009054906101000a90046001600160a01b0316468b338c896001600160a01b031663bcb67e146040518163ffffffff1660e01b8152600401602060405180830381600087803b158015611d5a57600080fd5b505af1158015611d6e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d929190613058565b611d9c898f613cc8565b611da69190613ca9565b60405161139f9796959493929190613324565b8151835114611dda5760405162461bcd60e51b815260040161068890613782565b8051835114611dfb5760405162461bcd60e51b8152600401610688906134e4565b60005b8351811015611e9957611e87848281518110611e2a57634e487b7160e01b600052603260045260246000fd5b6020026020010151848381518110611e5257634e487b7160e01b600052603260045260246000fd5b6020026020010151848481518110611e7a57634e487b7160e01b600052603260045260246000fd5b602002602001015161074d565b80611e9181613d22565b915050611dfe565b50505050565b611eaa6000336113b3565b611ec65760405162461bcd60e51b8152600401610688906137b9565b611e99848484846125dd565b611edd6000336113b3565b611ef95760405162461bcd60e51b8152600401610688906137b9565b6001600160a01b038281166000908152600a6020908152604080832085845290915290205416611f3b5760405162461bcd60e51b8152600401610688906139be565b6001600160a01b039091166000908152600a602090815260408083209383529290522080546001600160a01b0319169055565b600e602052600090815260409020546001600160a01b031681565b6001600160e01b031981166301ffc9a760e01b14919050565b6000611fb06012600a613bdb565b611fba8484612682565b611fc49190613b75565b9392505050565b611e99846323b872dd60e01b858585604051602401611fec9392919061329a565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915261268e565b3390565b61203182826113b3565b610b3757612049816001600160a01b0316601461271d565b61205483602061271d565b6040516020016120659291906131ce565b60408051601f198184030181529082905262461bcd60e51b82526106889160040161344e565b61209582826113b3565b610b375760008281526004602090815260408083206001600160a01b03851684529091529020805460ff191660011790556120ce612023565b6001600160a01b0316816001600160a01b0316837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45050565b61211c82826113b3565b15610b375760008281526004602090815260408083206001600160a01b03851684529091529020805460ff19169055612153612023565b6001600160a01b0316816001600160a01b0316837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45050565b6001600160a01b03909116600090815260076020526040902055565b60005b82518110156122f95760008282815181106121e157634e487b7160e01b600052603260045260246000fd5b602002602001015111156122e7576001600160a01b0385166000908152600a6020526040812084517f79aee25f54411be6ea1ab53c9d3cc3e245c301d481091deda6343da926932c0c9288929187908690811061224e57634e487b7160e01b600052603260045260246000fd5b6020026020010151815260200190815260200160002060009054906101000a90046001600160a01b03164686858151811061229957634e487b7160e01b600052603260045260246000fd5b602002602001015133898888815181106122c357634e487b7160e01b600052603260045260246000fd5b60200260200101516040516122de9796959493929190613324565b60405180910390a15b806122f181613d22565b9150506121b6565b5050505050565b6001600160a01b0385166000908152602081905260408120546123355760405162461bcd60e51b815260040161068890613885565b6001600160a01b038616600090815260208181526040808320549051909291612368918a918a918a918a918a9101613159565b60408051601f1981840301815291815281516020928301206000818152600190935291205490915060ff16156123b05760405162461bcd60e51b8152600401610688906135e8565b600081815260026020908152604080832033845290915290205460ff16156123ea5760405162461bcd60e51b8152600401610688906136ba565b6000818152600360205260409020546124049060016128cf565b60008281526003602081815260408084209485556002825280842033855282528320805460ff1916600117905591849052905254821161245c576000818152600160208190526040909120805460ff19168217905592505b600081815260036020526040908190205490517fe458477b94285c9f254a407e2614c39f55230b527ff7f51d76fdf33e251ae936916124a7918b918b918b918b913391908a90613257565b60405180910390a1505095945050505050565b60008060006124ca8787876105eb565b919450925090506124e56001600160a01b0388168583612526565b6001600160a01b038088166000818152600e602052604090205461250a921684612526565b82156109b4576009546109b4906001600160a01b038981169116855b6107078363a9059cbb60e01b8484604051602401611fec9291906133c7565b6001600160a01b0382166000908152602081905260409081902080549083905590517fb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e4490612598908590849086906133e0565b60405180910390a1505050565b600980546001600160a01b0319166001600160a01b0392909216919091179055565b60008183106125d65781611fc4565b5090919050565b670de0b6b3a76400008111156126055760405162461bcd60e51b815260040161068890613a44565b6001600160a01b03841660008181526005602090815260408083208784528252808320869055928252600681528282208683529052819020829055517fe7de0268825882caac9be515100046b260e4bb88ef28dcf7e4d99a3b9b75378290612674908690869086908690613401565b60405180910390a150505050565b6000611fc48284613ca9565b60006126e3826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166128db9092919063ffffffff16565b80519091501561070757808060200190518101906127019190612fa7565b6107075760405162461bcd60e51b8152600401610688906138d5565b6060600061272c836002613ca9565b612737906002613b5d565b67ffffffffffffffff81111561275d57634e487b7160e01b600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015612787576020820181803683370190505b509050600360fc1b816000815181106127b057634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a905350600f60fb1b816001815181106127ed57634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a9053506000612811846002613ca9565b61281c906001613b5d565b90505b60018111156128b0576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811061285e57634e487b7160e01b600052603260045260246000fd5b1a60f81b82828151811061288257634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a90535060049490941c936128a981613d0b565b905061281f565b508315611fc45760405162461bcd60e51b81526004016106889061348e565b6000611fc48284613b5d565b60606128ea84846000856128f2565b949350505050565b6060824710156129145760405162461bcd60e51b815260040161068890613674565b61291d856129b2565b6129395760405162461bcd60e51b81526004016106889061384e565b600080866001600160a01b0316858760405161295591906131b2565b60006040518083038185875af1925050503d8060008114612992576040519150601f19603f3d011682016040523d82523d6000602084013e612997565b606091505b50915091506129a78282866129b8565b979650505050505050565b3b151590565b606083156129c7575081611fc4565b8251156129d75782518084602001fd5b8160405162461bcd60e51b8152600401610688919061344e565b80356001600160a01b03811681146105e657600080fd5b600082601f830112612a18578081fd5b81356020612a2d612a2883613b39565b613b0f565b8281528181019085830183850287018401881015612a49578586fd5b855b85811015612a6e57612a5c826129f1565b84529284019290840190600101612a4b565b5090979650505050505050565b600082601f830112612a8b578081fd5b81356020612a9b612a2883613b39565b8281528181019085830183850287018401881015612ab7578586fd5b855b85811015612a6e57813584529284019290840190600101612ab9565b600082601f830112612ae5578081fd5b81516020612af5612a2883613b39565b8281528181019085830183850287018401881015612b11578586fd5b855b85811015612a6e57815184529284019290840190600101612b13565b600082601f830112612b3f578081fd5b813567ffffffffffffffff811115612b5957612b59613d53565b612b6c601f8201601f1916602001613b0f565b818152846020838601011115612b80578283fd5b816020850160208301379081016020019190915292915050565b600060208284031215612bab578081fd5b611fc4826129f1565b600080600060608486031215612bc8578182fd5b612bd1846129f1565b9250612bdf602085016129f1565b9150604084013567ffffffffffffffff811115612bfa578182fd5b612c0686828701612a7b565b9150509250925092565b600080600060608486031215612c24578283fd5b612c2d846129f1565b9250612c3b602085016129f1565b9150604084013590509250925092565b60008060408385031215612c5d578182fd5b612c66836129f1565b9150602083013567ffffffffffffffff811115612c81578182fd5b612c8d85828601612a7b565b9150509250929050565b60008060408385031215612ca9578182fd5b612cb2836129f1565b946020939093013593505050565b600080600060608486031215612cd4578283fd5b612cdd846129f1565b925060208401359150612cf2604085016129f1565b90509250925092565b60008060008060008060c08789031215612d13578182fd5b612d1c876129f1565b955060208701359450612d31604088016129f1565b9350612d3f606088016129f1565b92506080870135915060a087013567ffffffffffffffff811115612d61578182fd5b612d6d89828a01612b2f565b9150509295509295509295565b60008060008060808587031215612d8f578182fd5b612d98856129f1565b935060208501359250612dad604086016129f1565b9396929550929360600135925050565b600080600080600060a08688031215612dd4578283fd5b612ddd866129f1565b945060208601359350612df2604087016129f1565b925060608601359150608086013567ffffffffffffffff811115612e14578182fd5b612e2088828901612b2f565b9150509295509295909350565b600080600060608486031215612e41578081fd5b612e4a846129f1565b9250602084013591506040840135612e6181613d69565b809150509250925092565b600080600060608486031215612e80578081fd5b612e89846129f1565b95602085013595506040909401359392505050565b60008060008060808587031215612eb3578182fd5b612ebc856129f1565b966020860135965060408601359560600135945092505050565b600080600060608486031215612eea578081fd5b833567ffffffffffffffff80821115612f01578283fd5b612f0d87838801612a08565b94506020860135915080821115612f22578283fd5b612f2e87838801612a08565b93506040860135915080821115612f43578283fd5b50612c0686828701612a7b565b60008060408385031215612f62578182fd5b823567ffffffffffffffff80821115612f79578384fd5b612f8586838701612a08565b93506020850135915080821115612f9a578283fd5b50612c8d85828601612a7b565b600060208284031215612fb8578081fd5b8151611fc481613d69565b600060208284031215612fd4578081fd5b5035919050565b60008060408385031215612fed578182fd5b82359150610a1b602084016129f1565b60006020828403121561300e578081fd5b81356001600160e01b031981168114611fc4578182fd5b600060208284031215613036578081fd5b813567ffffffffffffffff81111561304c578182fd5b6128ea84828501612b2f565b600060208284031215613069578081fd5b5051919050565b600080600060608486031215613084578081fd5b83519250602084015167ffffffffffffffff808211156130a2578283fd5b6130ae87838801612ad5565b935060408601519150808211156130c3578283fd5b50612c0686828701612ad5565b600080604083850312156130e2578182fd5b505080516020909101519092909150565b6000815180845260208085019450808401835b8381101561312257815187529582019590820190600101613106565b509495945050505050565b60008151808452613145816020860160208601613cdf565b601f01601f19169290920160200192915050565b60006bffffffffffffffffffffffff19808860601b168352808760601b166014840152808660601b1660288401525083603c83015282516131a181605c850160208701613cdf565b91909101605c019695505050505050565b600082516131c4818460208701613cdf565b9190910192915050565b60007f416363657373436f6e74726f6c3a206163636f756e742000000000000000000082528351613206816017850160208801613cdf565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351613237816028840160208801613cdf565b01602801949350505050565b6001600160a01b0391909116815260200190565b6001600160a01b039788168152958716602087015293861660408601526060850192909252909316608083015260a082019290925260c081019190915260e00190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b6001600160a01b039586168152938516602085015260408401929092529092166060820152608081019190915260a00190565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090610df49083018461312d565b6001600160a01b0397881681529587166020870152604086019490945260608501929092528416608084015290921660a082015260c081019190915260e00190565b6001600160a01b038981168252888116602083015260408201889052606082018790528581166080830152841660a082015260c0810183905261010060e082018190526000906133b88382018561312d565b9b9a5050505050505050505050565b6001600160a01b03929092168252602082015260400190565b6001600160a01b039390931683526020830191909152604082015260600190565b6001600160a01b0394909416845260208401929092526040830152606082015260800190565b600060208252611fc460208301846130f3565b901515815260200190565b90815260200190565b600060208252611fc4602083018461312d565b6020808252601390820152721d1bdad95b881b9bdd081cdd5c1c1bdc9d1959606a1b604082015260600190565b6020808252818101527f537472696e67733a20686578206c656e67746820696e73756666696369656e74604082015260600190565b60208082526007908201526607a65726f206c760cc1b604082015260600190565b60208082526021908201527f54776f5761793a20636861696e494473206c656e677468206e6f74206d6174636040820152600d60fb1b606082015260800190565b60208082526010908201526f646f6e74206e656564206368616e676560801b604082015260600190565b6020808252601190820152703737ba1039bab83837b93a103a37b5b2b760791b604082015260600190565b6020808252601d908201527f54776f5761793a2063616c6c6572206973206e6f742063726f73736572000000604082015260600190565b6020808252601e908201527f54776f5761793a20546f6b6520616c726561647920537570706f727465640000604082015260600190565b60208082526018908201527f5f766f74653a3a70726f706f73616c2066696e69736865640000000000000000604082015260600190565b60208082526010908201526f3737ba1032b737bab3b4103a37b5b2b760811b604082015260600190565b602080825260119082015270151ddbd5d85e4e881d1e081b5a5b9d1959607a1b604082015260600190565b60208082526026908201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6040820152651c8818d85b1b60d21b606082015260800190565b60208082526017908201527f5f766f74653a3a6d73672e73656e64657220766f746564000000000000000000604082015260600190565b6020808252601190820152703737ba1039b7bab83837b93a103830b4b960791b604082015260600190565b602080825260159082015274151ddbd5d85e4e881d1e081c9bdb1b189858dad959605a1b604082015260600190565b60208082526017908201527f746f6b656e20616c726561647920737570706f72746564000000000000000000604082015260600190565b6020808252601e908201527f54776f5761793a20746f6b656e206c656e677468206e6f74206d617463680000604082015260600190565b6020808252601b908201527f54776f5761793a2063616c6c6572206973206e6f742061646d696e0000000000604082015260600190565b6020808252600d908201526c04e6f7420656e6f756768206c7609c1b604082015260600190565b6020808252601e908201527f54776f5761793a206e6f7420737570706f7274207468697320746f6b656e0000604082015260600190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b60208082526030908201527f50726f706f73616c566f74653a207468726573686f6c642073686f756c64206260408201526f0652067726561746572207468616e20360841b606082015260800190565b6020808252602a908201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6040820152691bdd081cdd58d8d9595960b21b606082015260800190565b60208082526025908201527f54776f5761793a20616d6f756e74206d75737420626520677265617465722074604082015264068616e20360dc1b606082015260800190565b602080825260139082015272151ddbd5d85e4e881d1e081d5b9b1bd8dad959606a1b604082015260600190565b60208082526013908201527254776f5761793a20746f20697320656d70747960681b604082015260600190565b6020808252601a908201527f54776f5761793a20746f6b65206e6f7420737570706f72746564000000000000604082015260600190565b6020808252602f908201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560408201526e103937b632b9903337b91039b2b63360891b606082015260800190565b602080825260159082015274199959481c985d1a5bc81b9bdd0818dbdc9c9958dd605a1b604082015260600190565b81516001600160a01b0390811682526020808401519083015260408084015190830152606080840151908301526080808401519091169082015260a0918201519181019190915260c00190565b600084825260606020830152613ad960608301856130f3565b8281036040840152610df481856130f3565b918252602082015260400190565b9283526020830191909152604082015260600190565b60405181810167ffffffffffffffff81118282101715613b3157613b31613d53565b604052919050565b600067ffffffffffffffff821115613b5357613b53613d53565b5060209081020190565b60008219821115613b7057613b70613d3d565b500190565b600082613b9057634e487b7160e01b81526012600452602481fd5b500490565b80825b6001808611613ba75750613bd2565b818704821115613bb957613bb9613d3d565b80861615613bc657918102915b9490941c938002613b98565b94509492505050565b6000611fc46000198484600082613bf457506001611fc4565b81613c0157506000611fc4565b8160018114613c175760028114613c2157613c4e565b6001915050611fc4565b60ff841115613c3257613c32613d3d565b6001841b915084821115613c4857613c48613d3d565b50611fc4565b5060208310610133831016604e8410600b8410161715613c81575081810a83811115613c7c57613c7c613d3d565b611fc4565b613c8e8484846001613b95565b808604821115613ca057613ca0613d3d565b02949350505050565b6000816000190483118215151615613cc357613cc3613d3d565b500290565b600082821015613cda57613cda613d3d565b500390565b60005b83811015613cfa578181015183820152602001613ce2565b83811115611e995750506000910152565b600081613d1a57613d1a613d3d565b506000190190565b6000600019821415613d3657613d36613d3d565b5060010190565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fd5b801515811461169057600080fdfea26469706673582212204c5736dfebfa56b5541381aad7da4a585ed1015b9b36a3a7e8579bb6b56e7a5964736f6c63430008000033"

// DeployTwoWay deploys a new Ethereum contract, binding an instance of TwoWay to it.
func DeployTwoWay(auth *bind.TransactOpts, backend bind.ContractBackend, _feeToDev common.Address) (common.Address, *types.Transaction, *TwoWay, error) {
	parsed, err := abi.JSON(strings.NewReader(TwoWayABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TwoWayBin), backend, _feeToDev)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TwoWay{TwoWayCaller: TwoWayCaller{contract: contract}, TwoWayTransactor: TwoWayTransactor{contract: contract}, TwoWayFilterer: TwoWayFilterer{contract: contract}}, nil
}

// TwoWay is an auto generated Go binding around an Ethereum contract.
type TwoWay struct {
	TwoWayCaller     // Read-only binding to the contract
	TwoWayTransactor // Write-only binding to the contract
	TwoWayFilterer   // Log filterer for contract events
}

// TwoWayCaller is an auto generated read-only Go binding around an Ethereum contract.
type TwoWayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoWayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TwoWayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoWayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TwoWayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoWaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TwoWaySession struct {
	Contract     *TwoWay           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TwoWayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TwoWayCallerSession struct {
	Contract *TwoWayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TwoWayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TwoWayTransactorSession struct {
	Contract     *TwoWayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TwoWayRaw is an auto generated low-level Go binding around an Ethereum contract.
type TwoWayRaw struct {
	Contract *TwoWay // Generic contract binding to access the raw methods on
}

// TwoWayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TwoWayCallerRaw struct {
	Contract *TwoWayCaller // Generic read-only contract binding to access the raw methods on
}

// TwoWayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TwoWayTransactorRaw struct {
	Contract *TwoWayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTwoWay creates a new instance of TwoWay, bound to a specific deployed contract.
func NewTwoWay(address common.Address, backend bind.ContractBackend) (*TwoWay, error) {
	contract, err := bindTwoWay(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TwoWay{TwoWayCaller: TwoWayCaller{contract: contract}, TwoWayTransactor: TwoWayTransactor{contract: contract}, TwoWayFilterer: TwoWayFilterer{contract: contract}}, nil
}

// NewTwoWayCaller creates a new read-only instance of TwoWay, bound to a specific deployed contract.
func NewTwoWayCaller(address common.Address, caller bind.ContractCaller) (*TwoWayCaller, error) {
	contract, err := bindTwoWay(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TwoWayCaller{contract: contract}, nil
}

// NewTwoWayTransactor creates a new write-only instance of TwoWay, bound to a specific deployed contract.
func NewTwoWayTransactor(address common.Address, transactor bind.ContractTransactor) (*TwoWayTransactor, error) {
	contract, err := bindTwoWay(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TwoWayTransactor{contract: contract}, nil
}

// NewTwoWayFilterer creates a new log filterer instance of TwoWay, bound to a specific deployed contract.
func NewTwoWayFilterer(address common.Address, filterer bind.ContractFilterer) (*TwoWayFilterer, error) {
	contract, err := bindTwoWay(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TwoWayFilterer{contract: contract}, nil
}

// bindTwoWay binds a generic wrapper to an already deployed contract.
func bindTwoWay(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TwoWayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwoWay *TwoWayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwoWay.Contract.TwoWayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwoWay *TwoWayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwoWay.Contract.TwoWayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwoWay *TwoWayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwoWay.Contract.TwoWayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwoWay *TwoWayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwoWay.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwoWay *TwoWayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwoWay.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwoWay *TwoWayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwoWay.Contract.contract.Transact(opts, method, params...)
}

// CROSSERROLE is a free data retrieval call binding the contract method 0x56cf02d9.
//
// Solidity: function CROSSER_ROLE() view returns(bytes32)
func (_TwoWay *TwoWayCaller) CROSSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "CROSSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CROSSERROLE is a free data retrieval call binding the contract method 0x56cf02d9.
//
// Solidity: function CROSSER_ROLE() view returns(bytes32)
func (_TwoWay *TwoWaySession) CROSSERROLE() ([32]byte, error) {
	return _TwoWay.Contract.CROSSERROLE(&_TwoWay.CallOpts)
}

// CROSSERROLE is a free data retrieval call binding the contract method 0x56cf02d9.
//
// Solidity: function CROSSER_ROLE() view returns(bytes32)
func (_TwoWay *TwoWayCallerSession) CROSSERROLE() ([32]byte, error) {
	return _TwoWay.Contract.CROSSERROLE(&_TwoWay.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TwoWay *TwoWayCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TwoWay *TwoWaySession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TwoWay.Contract.DEFAULTADMINROLE(&_TwoWay.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_TwoWay *TwoWayCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _TwoWay.Contract.DEFAULTADMINROLE(&_TwoWay.CallOpts)
}

// CalculateFee is a free data retrieval call binding the contract method 0x0324ef9c.
//
// Solidity: function calculateFee(address token, uint256 chainID, uint256 amount) view returns(uint256 feeAmountFix, uint256 feeAmountRatio, uint256 remainAmount)
func (_TwoWay *TwoWayCaller) CalculateFee(opts *bind.CallOpts, token common.Address, chainID *big.Int, amount *big.Int) (struct {
	FeeAmountFix   *big.Int
	FeeAmountRatio *big.Int
	RemainAmount   *big.Int
}, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "calculateFee", token, chainID, amount)

	outstruct := new(struct {
		FeeAmountFix   *big.Int
		FeeAmountRatio *big.Int
		RemainAmount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeAmountFix = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeeAmountRatio = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RemainAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateFee is a free data retrieval call binding the contract method 0x0324ef9c.
//
// Solidity: function calculateFee(address token, uint256 chainID, uint256 amount) view returns(uint256 feeAmountFix, uint256 feeAmountRatio, uint256 remainAmount)
func (_TwoWay *TwoWaySession) CalculateFee(token common.Address, chainID *big.Int, amount *big.Int) (struct {
	FeeAmountFix   *big.Int
	FeeAmountRatio *big.Int
	RemainAmount   *big.Int
}, error) {
	return _TwoWay.Contract.CalculateFee(&_TwoWay.CallOpts, token, chainID, amount)
}

// CalculateFee is a free data retrieval call binding the contract method 0x0324ef9c.
//
// Solidity: function calculateFee(address token, uint256 chainID, uint256 amount) view returns(uint256 feeAmountFix, uint256 feeAmountRatio, uint256 remainAmount)
func (_TwoWay *TwoWayCallerSession) CalculateFee(token common.Address, chainID *big.Int, amount *big.Int) (struct {
	FeeAmountFix   *big.Int
	FeeAmountRatio *big.Int
	RemainAmount   *big.Int
}, error) {
	return _TwoWay.Contract.CalculateFee(&_TwoWay.CallOpts, token, chainID, amount)
}

// CalculateRemoveFee is a free data retrieval call binding the contract method 0x2d63c99a.
//
// Solidity: function calculateRemoveFee(address token, uint256 amount) view returns(uint256 feeAmount, uint256 remainAmount)
func (_TwoWay *TwoWayCaller) CalculateRemoveFee(opts *bind.CallOpts, token common.Address, amount *big.Int) (struct {
	FeeAmount    *big.Int
	RemainAmount *big.Int
}, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "calculateRemoveFee", token, amount)

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

// CalculateRemoveFee is a free data retrieval call binding the contract method 0x2d63c99a.
//
// Solidity: function calculateRemoveFee(address token, uint256 amount) view returns(uint256 feeAmount, uint256 remainAmount)
func (_TwoWay *TwoWaySession) CalculateRemoveFee(token common.Address, amount *big.Int) (struct {
	FeeAmount    *big.Int
	RemainAmount *big.Int
}, error) {
	return _TwoWay.Contract.CalculateRemoveFee(&_TwoWay.CallOpts, token, amount)
}

// CalculateRemoveFee is a free data retrieval call binding the contract method 0x2d63c99a.
//
// Solidity: function calculateRemoveFee(address token, uint256 amount) view returns(uint256 feeAmount, uint256 remainAmount)
func (_TwoWay *TwoWayCallerSession) CalculateRemoveFee(token common.Address, amount *big.Int) (struct {
	FeeAmount    *big.Int
	RemainAmount *big.Int
}, error) {
	return _TwoWay.Contract.CalculateRemoveFee(&_TwoWay.CallOpts, token, amount)
}

// FeeAmountM is a free data retrieval call binding the contract method 0x64078943.
//
// Solidity: function feeAmountM(address , uint256 ) view returns(uint256)
func (_TwoWay *TwoWayCaller) FeeAmountM(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "feeAmountM", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeAmountM is a free data retrieval call binding the contract method 0x64078943.
//
// Solidity: function feeAmountM(address , uint256 ) view returns(uint256)
func (_TwoWay *TwoWaySession) FeeAmountM(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWay.Contract.FeeAmountM(&_TwoWay.CallOpts, arg0, arg1)
}

// FeeAmountM is a free data retrieval call binding the contract method 0x64078943.
//
// Solidity: function feeAmountM(address , uint256 ) view returns(uint256)
func (_TwoWay *TwoWayCallerSession) FeeAmountM(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWay.Contract.FeeAmountM(&_TwoWay.CallOpts, arg0, arg1)
}

// FeeRatioM is a free data retrieval call binding the contract method 0x9c8d70be.
//
// Solidity: function feeRatioM(address , uint256 ) view returns(uint256)
func (_TwoWay *TwoWayCaller) FeeRatioM(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "feeRatioM", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRatioM is a free data retrieval call binding the contract method 0x9c8d70be.
//
// Solidity: function feeRatioM(address , uint256 ) view returns(uint256)
func (_TwoWay *TwoWaySession) FeeRatioM(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWay.Contract.FeeRatioM(&_TwoWay.CallOpts, arg0, arg1)
}

// FeeRatioM is a free data retrieval call binding the contract method 0x9c8d70be.
//
// Solidity: function feeRatioM(address , uint256 ) view returns(uint256)
func (_TwoWay *TwoWayCallerSession) FeeRatioM(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWay.Contract.FeeRatioM(&_TwoWay.CallOpts, arg0, arg1)
}

// FeeTo is a free data retrieval call binding the contract method 0xdb697be4.
//
// Solidity: function feeTo(address , uint256 ) view returns(address)
func (_TwoWay *TwoWayCaller) FeeTo(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "feeTo", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0xdb697be4.
//
// Solidity: function feeTo(address , uint256 ) view returns(address)
func (_TwoWay *TwoWaySession) FeeTo(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _TwoWay.Contract.FeeTo(&_TwoWay.CallOpts, arg0, arg1)
}

// FeeTo is a free data retrieval call binding the contract method 0xdb697be4.
//
// Solidity: function feeTo(address , uint256 ) view returns(address)
func (_TwoWay *TwoWayCallerSession) FeeTo(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _TwoWay.Contract.FeeTo(&_TwoWay.CallOpts, arg0, arg1)
}

// FeeToDev is a free data retrieval call binding the contract method 0x100be3bd.
//
// Solidity: function feeToDev() view returns(address)
func (_TwoWay *TwoWayCaller) FeeToDev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "feeToDev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToDev is a free data retrieval call binding the contract method 0x100be3bd.
//
// Solidity: function feeToDev() view returns(address)
func (_TwoWay *TwoWaySession) FeeToDev() (common.Address, error) {
	return _TwoWay.Contract.FeeToDev(&_TwoWay.CallOpts)
}

// FeeToDev is a free data retrieval call binding the contract method 0x100be3bd.
//
// Solidity: function feeToDev() view returns(address)
func (_TwoWay *TwoWayCallerSession) FeeToDev() (common.Address, error) {
	return _TwoWay.Contract.FeeToDev(&_TwoWay.CallOpts)
}

// GetMaxToken0AmountOut is a free data retrieval call binding the contract method 0x934ac707.
//
// Solidity: function getMaxToken0AmountOut(address token0, uint256 chainID) view returns(uint256)
func (_TwoWay *TwoWayCaller) GetMaxToken0AmountOut(opts *bind.CallOpts, token0 common.Address, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "getMaxToken0AmountOut", token0, chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxToken0AmountOut is a free data retrieval call binding the contract method 0x934ac707.
//
// Solidity: function getMaxToken0AmountOut(address token0, uint256 chainID) view returns(uint256)
func (_TwoWay *TwoWaySession) GetMaxToken0AmountOut(token0 common.Address, chainID *big.Int) (*big.Int, error) {
	return _TwoWay.Contract.GetMaxToken0AmountOut(&_TwoWay.CallOpts, token0, chainID)
}

// GetMaxToken0AmountOut is a free data retrieval call binding the contract method 0x934ac707.
//
// Solidity: function getMaxToken0AmountOut(address token0, uint256 chainID) view returns(uint256)
func (_TwoWay *TwoWayCallerSession) GetMaxToken0AmountOut(token0 common.Address, chainID *big.Int) (*big.Int, error) {
	return _TwoWay.Contract.GetMaxToken0AmountOut(&_TwoWay.CallOpts, token0, chainID)
}

// GetMaxToken1AmountOut is a free data retrieval call binding the contract method 0x56ecaeb7.
//
// Solidity: function getMaxToken1AmountOut(address token0, uint256 chainID) view returns(uint256)
func (_TwoWay *TwoWayCaller) GetMaxToken1AmountOut(opts *bind.CallOpts, token0 common.Address, chainID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "getMaxToken1AmountOut", token0, chainID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxToken1AmountOut is a free data retrieval call binding the contract method 0x56ecaeb7.
//
// Solidity: function getMaxToken1AmountOut(address token0, uint256 chainID) view returns(uint256)
func (_TwoWay *TwoWaySession) GetMaxToken1AmountOut(token0 common.Address, chainID *big.Int) (*big.Int, error) {
	return _TwoWay.Contract.GetMaxToken1AmountOut(&_TwoWay.CallOpts, token0, chainID)
}

// GetMaxToken1AmountOut is a free data retrieval call binding the contract method 0x56ecaeb7.
//
// Solidity: function getMaxToken1AmountOut(address token0, uint256 chainID) view returns(uint256)
func (_TwoWay *TwoWayCallerSession) GetMaxToken1AmountOut(token0 common.Address, chainID *big.Int) (*big.Int, error) {
	return _TwoWay.Contract.GetMaxToken1AmountOut(&_TwoWay.CallOpts, token0, chainID)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TwoWay *TwoWayCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TwoWay *TwoWaySession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TwoWay.Contract.GetRoleAdmin(&_TwoWay.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_TwoWay *TwoWayCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _TwoWay.Contract.GetRoleAdmin(&_TwoWay.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TwoWay *TwoWayCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TwoWay *TwoWaySession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TwoWay.Contract.HasRole(&_TwoWay.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_TwoWay *TwoWayCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _TwoWay.Contract.HasRole(&_TwoWay.CallOpts, role, account)
}

// Pairs is a free data retrieval call binding the contract method 0xfe33b302.
//
// Solidity: function pairs(address ) view returns(address)
func (_TwoWay *TwoWayCaller) Pairs(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "pairs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Pairs is a free data retrieval call binding the contract method 0xfe33b302.
//
// Solidity: function pairs(address ) view returns(address)
func (_TwoWay *TwoWaySession) Pairs(arg0 common.Address) (common.Address, error) {
	return _TwoWay.Contract.Pairs(&_TwoWay.CallOpts, arg0)
}

// Pairs is a free data retrieval call binding the contract method 0xfe33b302.
//
// Solidity: function pairs(address ) view returns(address)
func (_TwoWay *TwoWayCallerSession) Pairs(arg0 common.Address) (common.Address, error) {
	return _TwoWay.Contract.Pairs(&_TwoWay.CallOpts, arg0)
}

// RemoveFeeAmount is a free data retrieval call binding the contract method 0x0b19993f.
//
// Solidity: function removeFeeAmount(address ) view returns(uint256)
func (_TwoWay *TwoWayCaller) RemoveFeeAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "removeFeeAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemoveFeeAmount is a free data retrieval call binding the contract method 0x0b19993f.
//
// Solidity: function removeFeeAmount(address ) view returns(uint256)
func (_TwoWay *TwoWaySession) RemoveFeeAmount(arg0 common.Address) (*big.Int, error) {
	return _TwoWay.Contract.RemoveFeeAmount(&_TwoWay.CallOpts, arg0)
}

// RemoveFeeAmount is a free data retrieval call binding the contract method 0x0b19993f.
//
// Solidity: function removeFeeAmount(address ) view returns(uint256)
func (_TwoWay *TwoWayCallerSession) RemoveFeeAmount(arg0 common.Address) (*big.Int, error) {
	return _TwoWay.Contract.RemoveFeeAmount(&_TwoWay.CallOpts, arg0)
}

// SupportToken is a free data retrieval call binding the contract method 0xd692f4c5.
//
// Solidity: function supportToken(address , uint256 ) view returns(address)
func (_TwoWay *TwoWayCaller) SupportToken(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "supportToken", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportToken is a free data retrieval call binding the contract method 0xd692f4c5.
//
// Solidity: function supportToken(address , uint256 ) view returns(address)
func (_TwoWay *TwoWaySession) SupportToken(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _TwoWay.Contract.SupportToken(&_TwoWay.CallOpts, arg0, arg1)
}

// SupportToken is a free data retrieval call binding the contract method 0xd692f4c5.
//
// Solidity: function supportToken(address , uint256 ) view returns(address)
func (_TwoWay *TwoWayCallerSession) SupportToken(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _TwoWay.Contract.SupportToken(&_TwoWay.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TwoWay *TwoWayCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TwoWay *TwoWaySession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TwoWay.Contract.SupportsInterface(&_TwoWay.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_TwoWay *TwoWayCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _TwoWay.Contract.SupportsInterface(&_TwoWay.CallOpts, interfaceId)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_TwoWay *TwoWayCaller) Threshold(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "threshold", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_TwoWay *TwoWaySession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _TwoWay.Contract.Threshold(&_TwoWay.CallOpts, arg0)
}

// Threshold is a free data retrieval call binding the contract method 0xc86ec2bf.
//
// Solidity: function threshold(address ) view returns(uint256)
func (_TwoWay *TwoWayCallerSession) Threshold(arg0 common.Address) (*big.Int, error) {
	return _TwoWay.Contract.Threshold(&_TwoWay.CallOpts, arg0)
}

// TxMinted is a free data retrieval call binding the contract method 0x10c27402.
//
// Solidity: function txMinted(string ) view returns(bool)
func (_TwoWay *TwoWayCaller) TxMinted(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "txMinted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TxMinted is a free data retrieval call binding the contract method 0x10c27402.
//
// Solidity: function txMinted(string ) view returns(bool)
func (_TwoWay *TwoWaySession) TxMinted(arg0 string) (bool, error) {
	return _TwoWay.Contract.TxMinted(&_TwoWay.CallOpts, arg0)
}

// TxMinted is a free data retrieval call binding the contract method 0x10c27402.
//
// Solidity: function txMinted(string ) view returns(bool)
func (_TwoWay *TwoWayCallerSession) TxMinted(arg0 string) (bool, error) {
	return _TwoWay.Contract.TxMinted(&_TwoWay.CallOpts, arg0)
}

// TxRollbacked is a free data retrieval call binding the contract method 0x53ad72e5.
//
// Solidity: function txRollbacked(string ) view returns(bool)
func (_TwoWay *TwoWayCaller) TxRollbacked(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "txRollbacked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TxRollbacked is a free data retrieval call binding the contract method 0x53ad72e5.
//
// Solidity: function txRollbacked(string ) view returns(bool)
func (_TwoWay *TwoWaySession) TxRollbacked(arg0 string) (bool, error) {
	return _TwoWay.Contract.TxRollbacked(&_TwoWay.CallOpts, arg0)
}

// TxRollbacked is a free data retrieval call binding the contract method 0x53ad72e5.
//
// Solidity: function txRollbacked(string ) view returns(bool)
func (_TwoWay *TwoWayCallerSession) TxRollbacked(arg0 string) (bool, error) {
	return _TwoWay.Contract.TxRollbacked(&_TwoWay.CallOpts, arg0)
}

// TxUnlocked is a free data retrieval call binding the contract method 0x967145af.
//
// Solidity: function txUnlocked(string ) view returns(bool)
func (_TwoWay *TwoWayCaller) TxUnlocked(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "txUnlocked", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TxUnlocked is a free data retrieval call binding the contract method 0x967145af.
//
// Solidity: function txUnlocked(string ) view returns(bool)
func (_TwoWay *TwoWaySession) TxUnlocked(arg0 string) (bool, error) {
	return _TwoWay.Contract.TxUnlocked(&_TwoWay.CallOpts, arg0)
}

// TxUnlocked is a free data retrieval call binding the contract method 0x967145af.
//
// Solidity: function txUnlocked(string ) view returns(bool)
func (_TwoWay *TwoWayCallerSession) TxUnlocked(arg0 string) (bool, error) {
	return _TwoWay.Contract.TxUnlocked(&_TwoWay.CallOpts, arg0)
}

// UnlockFeeOn is a free data retrieval call binding the contract method 0x68b1188f.
//
// Solidity: function unlockFeeOn(address , uint256 ) view returns(bool)
func (_TwoWay *TwoWayCaller) UnlockFeeOn(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _TwoWay.contract.Call(opts, &out, "unlockFeeOn", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UnlockFeeOn is a free data retrieval call binding the contract method 0x68b1188f.
//
// Solidity: function unlockFeeOn(address , uint256 ) view returns(bool)
func (_TwoWay *TwoWaySession) UnlockFeeOn(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _TwoWay.Contract.UnlockFeeOn(&_TwoWay.CallOpts, arg0, arg1)
}

// UnlockFeeOn is a free data retrieval call binding the contract method 0x68b1188f.
//
// Solidity: function unlockFeeOn(address , uint256 ) view returns(bool)
func (_TwoWay *TwoWayCallerSession) UnlockFeeOn(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _TwoWay.Contract.UnlockFeeOn(&_TwoWay.CallOpts, arg0, arg1)
}

// AddChainIDs is a paid mutator transaction binding the contract method 0x27f8dd9b.
//
// Solidity: function addChainIDs(address token, uint256[] chainIDs) returns()
func (_TwoWay *TwoWayTransactor) AddChainIDs(opts *bind.TransactOpts, token common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "addChainIDs", token, chainIDs)
}

// AddChainIDs is a paid mutator transaction binding the contract method 0x27f8dd9b.
//
// Solidity: function addChainIDs(address token, uint256[] chainIDs) returns()
func (_TwoWay *TwoWaySession) AddChainIDs(token common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.AddChainIDs(&_TwoWay.TransactOpts, token, chainIDs)
}

// AddChainIDs is a paid mutator transaction binding the contract method 0x27f8dd9b.
//
// Solidity: function addChainIDs(address token, uint256[] chainIDs) returns()
func (_TwoWay *TwoWayTransactorSession) AddChainIDs(token common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.AddChainIDs(&_TwoWay.TransactOpts, token, chainIDs)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1b879378.
//
// Solidity: function addLiquidity(address token0, uint256 amount, address to) returns(uint256 liquidity)
func (_TwoWay *TwoWayTransactor) AddLiquidity(opts *bind.TransactOpts, token0 common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "addLiquidity", token0, amount, to)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1b879378.
//
// Solidity: function addLiquidity(address token0, uint256 amount, address to) returns(uint256 liquidity)
func (_TwoWay *TwoWaySession) AddLiquidity(token0 common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _TwoWay.Contract.AddLiquidity(&_TwoWay.TransactOpts, token0, amount, to)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1b879378.
//
// Solidity: function addLiquidity(address token0, uint256 amount, address to) returns(uint256 liquidity)
func (_TwoWay *TwoWayTransactorSession) AddLiquidity(token0 common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _TwoWay.Contract.AddLiquidity(&_TwoWay.TransactOpts, token0, amount, to)
}

// AddPair is a paid mutator transaction binding the contract method 0x31b9e5a3.
//
// Solidity: function addPair(address token, address pair, uint256[] chainIDs) returns()
func (_TwoWay *TwoWayTransactor) AddPair(opts *bind.TransactOpts, token common.Address, pair common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "addPair", token, pair, chainIDs)
}

// AddPair is a paid mutator transaction binding the contract method 0x31b9e5a3.
//
// Solidity: function addPair(address token, address pair, uint256[] chainIDs) returns()
func (_TwoWay *TwoWaySession) AddPair(token common.Address, pair common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.AddPair(&_TwoWay.TransactOpts, token, pair, chainIDs)
}

// AddPair is a paid mutator transaction binding the contract method 0x31b9e5a3.
//
// Solidity: function addPair(address token, address pair, uint256[] chainIDs) returns()
func (_TwoWay *TwoWayTransactorSession) AddPair(token common.Address, pair common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.AddPair(&_TwoWay.TransactOpts, token, pair, chainIDs)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x1b322be4.
//
// Solidity: function addSupportToken(address token0, address token1, uint256 chainID) returns()
func (_TwoWay *TwoWayTransactor) AddSupportToken(opts *bind.TransactOpts, token0 common.Address, token1 common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "addSupportToken", token0, token1, chainID)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x1b322be4.
//
// Solidity: function addSupportToken(address token0, address token1, uint256 chainID) returns()
func (_TwoWay *TwoWaySession) AddSupportToken(token0 common.Address, token1 common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.AddSupportToken(&_TwoWay.TransactOpts, token0, token1, chainID)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0x1b322be4.
//
// Solidity: function addSupportToken(address token0, address token1, uint256 chainID) returns()
func (_TwoWay *TwoWayTransactorSession) AddSupportToken(token0 common.Address, token1 common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.AddSupportToken(&_TwoWay.TransactOpts, token0, token1, chainID)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0xee8a6acf.
//
// Solidity: function addSupportTokens(address[] token0s, address[] token1s, uint256[] chainIDs) returns()
func (_TwoWay *TwoWayTransactor) AddSupportTokens(opts *bind.TransactOpts, token0s []common.Address, token1s []common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "addSupportTokens", token0s, token1s, chainIDs)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0xee8a6acf.
//
// Solidity: function addSupportTokens(address[] token0s, address[] token1s, uint256[] chainIDs) returns()
func (_TwoWay *TwoWaySession) AddSupportTokens(token0s []common.Address, token1s []common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.AddSupportTokens(&_TwoWay.TransactOpts, token0s, token1s, chainIDs)
}

// AddSupportTokens is a paid mutator transaction binding the contract method 0xee8a6acf.
//
// Solidity: function addSupportTokens(address[] token0s, address[] token1s, uint256[] chainIDs) returns()
func (_TwoWay *TwoWayTransactorSession) AddSupportTokens(token0s []common.Address, token1s []common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.AddSupportTokens(&_TwoWay.TransactOpts, token0s, token1s, chainIDs)
}

// CrossIn is a paid mutator transaction binding the contract method 0xcf28f67c.
//
// Solidity: function crossIn(address token0, uint256 chainID, address from, address to, uint256 amount, string txid) returns()
func (_TwoWay *TwoWayTransactor) CrossIn(opts *bind.TransactOpts, token0 common.Address, chainID *big.Int, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "crossIn", token0, chainID, from, to, amount, txid)
}

// CrossIn is a paid mutator transaction binding the contract method 0xcf28f67c.
//
// Solidity: function crossIn(address token0, uint256 chainID, address from, address to, uint256 amount, string txid) returns()
func (_TwoWay *TwoWaySession) CrossIn(token0 common.Address, chainID *big.Int, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _TwoWay.Contract.CrossIn(&_TwoWay.TransactOpts, token0, chainID, from, to, amount, txid)
}

// CrossIn is a paid mutator transaction binding the contract method 0xcf28f67c.
//
// Solidity: function crossIn(address token0, uint256 chainID, address from, address to, uint256 amount, string txid) returns()
func (_TwoWay *TwoWayTransactorSession) CrossIn(token0 common.Address, chainID *big.Int, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _TwoWay.Contract.CrossIn(&_TwoWay.TransactOpts, token0, chainID, from, to, amount, txid)
}

// CrossOut is a paid mutator transaction binding the contract method 0xdf7e600a.
//
// Solidity: function crossOut(address token0, uint256 chainID, address to, uint256 amount) returns()
func (_TwoWay *TwoWayTransactor) CrossOut(opts *bind.TransactOpts, token0 common.Address, chainID *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "crossOut", token0, chainID, to, amount)
}

// CrossOut is a paid mutator transaction binding the contract method 0xdf7e600a.
//
// Solidity: function crossOut(address token0, uint256 chainID, address to, uint256 amount) returns()
func (_TwoWay *TwoWaySession) CrossOut(token0 common.Address, chainID *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.CrossOut(&_TwoWay.TransactOpts, token0, chainID, to, amount)
}

// CrossOut is a paid mutator transaction binding the contract method 0xdf7e600a.
//
// Solidity: function crossOut(address token0, uint256 chainID, address to, uint256 amount) returns()
func (_TwoWay *TwoWayTransactorSession) CrossOut(token0 common.Address, chainID *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.CrossOut(&_TwoWay.TransactOpts, token0, chainID, to, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TwoWay *TwoWayTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TwoWay *TwoWaySession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWay.Contract.GrantRole(&_TwoWay.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_TwoWay *TwoWayTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWay.Contract.GrantRole(&_TwoWay.TransactOpts, role, account)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x47fdbc8e.
//
// Solidity: function removeLiquidity(address token0, uint256 lpAmount, address to) returns(uint256 amount0, uint256[] chainids, uint256[] amount1s)
func (_TwoWay *TwoWayTransactor) RemoveLiquidity(opts *bind.TransactOpts, token0 common.Address, lpAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "removeLiquidity", token0, lpAmount, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x47fdbc8e.
//
// Solidity: function removeLiquidity(address token0, uint256 lpAmount, address to) returns(uint256 amount0, uint256[] chainids, uint256[] amount1s)
func (_TwoWay *TwoWaySession) RemoveLiquidity(token0 common.Address, lpAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _TwoWay.Contract.RemoveLiquidity(&_TwoWay.TransactOpts, token0, lpAmount, to)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x47fdbc8e.
//
// Solidity: function removeLiquidity(address token0, uint256 lpAmount, address to) returns(uint256 amount0, uint256[] chainids, uint256[] amount1s)
func (_TwoWay *TwoWayTransactorSession) RemoveLiquidity(token0 common.Address, lpAmount *big.Int, to common.Address) (*types.Transaction, error) {
	return _TwoWay.Contract.RemoveLiquidity(&_TwoWay.TransactOpts, token0, lpAmount, to)
}

// RemovePair is a paid mutator transaction binding the contract method 0xaf6c9c1d.
//
// Solidity: function removePair(address token) returns()
func (_TwoWay *TwoWayTransactor) RemovePair(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "removePair", token)
}

// RemovePair is a paid mutator transaction binding the contract method 0xaf6c9c1d.
//
// Solidity: function removePair(address token) returns()
func (_TwoWay *TwoWaySession) RemovePair(token common.Address) (*types.Transaction, error) {
	return _TwoWay.Contract.RemovePair(&_TwoWay.TransactOpts, token)
}

// RemovePair is a paid mutator transaction binding the contract method 0xaf6c9c1d.
//
// Solidity: function removePair(address token) returns()
func (_TwoWay *TwoWayTransactorSession) RemovePair(token common.Address) (*types.Transaction, error) {
	return _TwoWay.Contract.RemovePair(&_TwoWay.TransactOpts, token)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xf8f8d5c0.
//
// Solidity: function removeSupportToken(address token0, uint256 chainID) returns()
func (_TwoWay *TwoWayTransactor) RemoveSupportToken(opts *bind.TransactOpts, token0 common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "removeSupportToken", token0, chainID)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xf8f8d5c0.
//
// Solidity: function removeSupportToken(address token0, uint256 chainID) returns()
func (_TwoWay *TwoWaySession) RemoveSupportToken(token0 common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.RemoveSupportToken(&_TwoWay.TransactOpts, token0, chainID)
}

// RemoveSupportToken is a paid mutator transaction binding the contract method 0xf8f8d5c0.
//
// Solidity: function removeSupportToken(address token0, uint256 chainID) returns()
func (_TwoWay *TwoWayTransactorSession) RemoveSupportToken(token0 common.Address, chainID *big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.RemoveSupportToken(&_TwoWay.TransactOpts, token0, chainID)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x03507ba5.
//
// Solidity: function removeSupportTokens(address[] token0s, uint256[] chainIDs) returns()
func (_TwoWay *TwoWayTransactor) RemoveSupportTokens(opts *bind.TransactOpts, token0s []common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "removeSupportTokens", token0s, chainIDs)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x03507ba5.
//
// Solidity: function removeSupportTokens(address[] token0s, uint256[] chainIDs) returns()
func (_TwoWay *TwoWaySession) RemoveSupportTokens(token0s []common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.RemoveSupportTokens(&_TwoWay.TransactOpts, token0s, chainIDs)
}

// RemoveSupportTokens is a paid mutator transaction binding the contract method 0x03507ba5.
//
// Solidity: function removeSupportTokens(address[] token0s, uint256[] chainIDs) returns()
func (_TwoWay *TwoWayTransactorSession) RemoveSupportTokens(token0s []common.Address, chainIDs []*big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.RemoveSupportTokens(&_TwoWay.TransactOpts, token0s, chainIDs)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TwoWay *TwoWayTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TwoWay *TwoWaySession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWay.Contract.RenounceRole(&_TwoWay.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_TwoWay *TwoWayTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWay.Contract.RenounceRole(&_TwoWay.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TwoWay *TwoWayTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TwoWay *TwoWaySession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWay.Contract.RevokeRole(&_TwoWay.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_TwoWay *TwoWayTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _TwoWay.Contract.RevokeRole(&_TwoWay.TransactOpts, role, account)
}

// Rollback is a paid mutator transaction binding the contract method 0x81660f5d.
//
// Solidity: function rollback(address token0, uint256 chainID, address from, uint256 amount, string txid) returns()
func (_TwoWay *TwoWayTransactor) Rollback(opts *bind.TransactOpts, token0 common.Address, chainID *big.Int, from common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "rollback", token0, chainID, from, amount, txid)
}

// Rollback is a paid mutator transaction binding the contract method 0x81660f5d.
//
// Solidity: function rollback(address token0, uint256 chainID, address from, uint256 amount, string txid) returns()
func (_TwoWay *TwoWaySession) Rollback(token0 common.Address, chainID *big.Int, from common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _TwoWay.Contract.Rollback(&_TwoWay.TransactOpts, token0, chainID, from, amount, txid)
}

// Rollback is a paid mutator transaction binding the contract method 0x81660f5d.
//
// Solidity: function rollback(address token0, uint256 chainID, address from, uint256 amount, string txid) returns()
func (_TwoWay *TwoWayTransactorSession) Rollback(token0 common.Address, chainID *big.Int, from common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _TwoWay.Contract.Rollback(&_TwoWay.TransactOpts, token0, chainID, from, amount, txid)
}

// SetFee is a paid mutator transaction binding the contract method 0xf06410f8.
//
// Solidity: function setFee(address token0, uint256 chainID, uint256 feeAmount, uint256 feeRatio) returns()
func (_TwoWay *TwoWayTransactor) SetFee(opts *bind.TransactOpts, token0 common.Address, chainID *big.Int, feeAmount *big.Int, feeRatio *big.Int) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "setFee", token0, chainID, feeAmount, feeRatio)
}

// SetFee is a paid mutator transaction binding the contract method 0xf06410f8.
//
// Solidity: function setFee(address token0, uint256 chainID, uint256 feeAmount, uint256 feeRatio) returns()
func (_TwoWay *TwoWaySession) SetFee(token0 common.Address, chainID *big.Int, feeAmount *big.Int, feeRatio *big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.SetFee(&_TwoWay.TransactOpts, token0, chainID, feeAmount, feeRatio)
}

// SetFee is a paid mutator transaction binding the contract method 0xf06410f8.
//
// Solidity: function setFee(address token0, uint256 chainID, uint256 feeAmount, uint256 feeRatio) returns()
func (_TwoWay *TwoWayTransactorSession) SetFee(token0 common.Address, chainID *big.Int, feeAmount *big.Int, feeRatio *big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.SetFee(&_TwoWay.TransactOpts, token0, chainID, feeAmount, feeRatio)
}

// SetFeeToDev is a paid mutator transaction binding the contract method 0xc369c773.
//
// Solidity: function setFeeToDev(address account) returns()
func (_TwoWay *TwoWayTransactor) SetFeeToDev(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "setFeeToDev", account)
}

// SetFeeToDev is a paid mutator transaction binding the contract method 0xc369c773.
//
// Solidity: function setFeeToDev(address account) returns()
func (_TwoWay *TwoWaySession) SetFeeToDev(account common.Address) (*types.Transaction, error) {
	return _TwoWay.Contract.SetFeeToDev(&_TwoWay.TransactOpts, account)
}

// SetFeeToDev is a paid mutator transaction binding the contract method 0xc369c773.
//
// Solidity: function setFeeToDev(address account) returns()
func (_TwoWay *TwoWayTransactorSession) SetFeeToDev(account common.Address) (*types.Transaction, error) {
	return _TwoWay.Contract.SetFeeToDev(&_TwoWay.TransactOpts, account)
}

// SetRemoveFee is a paid mutator transaction binding the contract method 0x3e1cd7e4.
//
// Solidity: function setRemoveFee(address token0, uint256 _feeAmount) returns()
func (_TwoWay *TwoWayTransactor) SetRemoveFee(opts *bind.TransactOpts, token0 common.Address, _feeAmount *big.Int) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "setRemoveFee", token0, _feeAmount)
}

// SetRemoveFee is a paid mutator transaction binding the contract method 0x3e1cd7e4.
//
// Solidity: function setRemoveFee(address token0, uint256 _feeAmount) returns()
func (_TwoWay *TwoWaySession) SetRemoveFee(token0 common.Address, _feeAmount *big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.SetRemoveFee(&_TwoWay.TransactOpts, token0, _feeAmount)
}

// SetRemoveFee is a paid mutator transaction binding the contract method 0x3e1cd7e4.
//
// Solidity: function setRemoveFee(address token0, uint256 _feeAmount) returns()
func (_TwoWay *TwoWayTransactorSession) SetRemoveFee(token0 common.Address, _feeAmount *big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.SetRemoveFee(&_TwoWay.TransactOpts, token0, _feeAmount)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token, uint256 _threshold) returns()
func (_TwoWay *TwoWayTransactor) SetThreshold(opts *bind.TransactOpts, token common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "setThreshold", token, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token, uint256 _threshold) returns()
func (_TwoWay *TwoWaySession) SetThreshold(token common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.SetThreshold(&_TwoWay.TransactOpts, token, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token, uint256 _threshold) returns()
func (_TwoWay *TwoWayTransactorSession) SetThreshold(token common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _TwoWay.Contract.SetThreshold(&_TwoWay.TransactOpts, token, _threshold)
}

// SetUnlockFeeOn is a paid mutator transaction binding the contract method 0x98d2f6b4.
//
// Solidity: function setUnlockFeeOn(address token0, uint256 chainId, bool _inactived) returns()
func (_TwoWay *TwoWayTransactor) SetUnlockFeeOn(opts *bind.TransactOpts, token0 common.Address, chainId *big.Int, _inactived bool) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "setUnlockFeeOn", token0, chainId, _inactived)
}

// SetUnlockFeeOn is a paid mutator transaction binding the contract method 0x98d2f6b4.
//
// Solidity: function setUnlockFeeOn(address token0, uint256 chainId, bool _inactived) returns()
func (_TwoWay *TwoWaySession) SetUnlockFeeOn(token0 common.Address, chainId *big.Int, _inactived bool) (*types.Transaction, error) {
	return _TwoWay.Contract.SetUnlockFeeOn(&_TwoWay.TransactOpts, token0, chainId, _inactived)
}

// SetUnlockFeeOn is a paid mutator transaction binding the contract method 0x98d2f6b4.
//
// Solidity: function setUnlockFeeOn(address token0, uint256 chainId, bool _inactived) returns()
func (_TwoWay *TwoWayTransactorSession) SetUnlockFeeOn(token0 common.Address, chainId *big.Int, _inactived bool) (*types.Transaction, error) {
	return _TwoWay.Contract.SetUnlockFeeOn(&_TwoWay.TransactOpts, token0, chainId, _inactived)
}

// Unlock is a paid mutator transaction binding the contract method 0x715ec45c.
//
// Solidity: function unlock(address token0, uint256 chainID, address from, address to, uint256 amount, string txid) returns()
func (_TwoWay *TwoWayTransactor) Unlock(opts *bind.TransactOpts, token0 common.Address, chainID *big.Int, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _TwoWay.contract.Transact(opts, "unlock", token0, chainID, from, to, amount, txid)
}

// Unlock is a paid mutator transaction binding the contract method 0x715ec45c.
//
// Solidity: function unlock(address token0, uint256 chainID, address from, address to, uint256 amount, string txid) returns()
func (_TwoWay *TwoWaySession) Unlock(token0 common.Address, chainID *big.Int, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _TwoWay.Contract.Unlock(&_TwoWay.TransactOpts, token0, chainID, from, to, amount, txid)
}

// Unlock is a paid mutator transaction binding the contract method 0x715ec45c.
//
// Solidity: function unlock(address token0, uint256 chainID, address from, address to, uint256 amount, string txid) returns()
func (_TwoWay *TwoWayTransactorSession) Unlock(token0 common.Address, chainID *big.Int, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _TwoWay.Contract.Unlock(&_TwoWay.TransactOpts, token0, chainID, from, to, amount, txid)
}

// TwoWayCrossBurnIterator is returned from FilterCrossBurn and is used to iterate over the raw logs and unpacked data for CrossBurn events raised by the TwoWay contract.
type TwoWayCrossBurnIterator struct {
	Event *TwoWayCrossBurn // Event containing the contract specifics and raw log

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
func (it *TwoWayCrossBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayCrossBurn)
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
		it.Event = new(TwoWayCrossBurn)
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
func (it *TwoWayCrossBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayCrossBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayCrossBurn represents a CrossBurn event raised by the TwoWay contract.
type TwoWayCrossBurn struct {
	Token0   common.Address
	Token1   common.Address
	ChainID0 *big.Int
	ChainID1 *big.Int
	From     common.Address
	To       common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCrossBurn is a free log retrieval operation binding the contract event 0x79aee25f54411be6ea1ab53c9d3cc3e245c301d481091deda6343da926932c0c.
//
// Solidity: event CrossBurn(address token0, address token1, uint256 chainID0, uint256 chainID1, address from, address to, uint256 amount)
func (_TwoWay *TwoWayFilterer) FilterCrossBurn(opts *bind.FilterOpts) (*TwoWayCrossBurnIterator, error) {

	logs, sub, err := _TwoWay.contract.FilterLogs(opts, "CrossBurn")
	if err != nil {
		return nil, err
	}
	return &TwoWayCrossBurnIterator{contract: _TwoWay.contract, event: "CrossBurn", logs: logs, sub: sub}, nil
}

// WatchCrossBurn is a free log subscription operation binding the contract event 0x79aee25f54411be6ea1ab53c9d3cc3e245c301d481091deda6343da926932c0c.
//
// Solidity: event CrossBurn(address token0, address token1, uint256 chainID0, uint256 chainID1, address from, address to, uint256 amount)
func (_TwoWay *TwoWayFilterer) WatchCrossBurn(opts *bind.WatchOpts, sink chan<- *TwoWayCrossBurn) (event.Subscription, error) {

	logs, sub, err := _TwoWay.contract.WatchLogs(opts, "CrossBurn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayCrossBurn)
				if err := _TwoWay.contract.UnpackLog(event, "CrossBurn", log); err != nil {
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

// ParseCrossBurn is a log parse operation binding the contract event 0x79aee25f54411be6ea1ab53c9d3cc3e245c301d481091deda6343da926932c0c.
//
// Solidity: event CrossBurn(address token0, address token1, uint256 chainID0, uint256 chainID1, address from, address to, uint256 amount)
func (_TwoWay *TwoWayFilterer) ParseCrossBurn(log types.Log) (*TwoWayCrossBurn, error) {
	event := new(TwoWayCrossBurn)
	if err := _TwoWay.contract.UnpackLog(event, "CrossBurn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayFeeChangeIterator is returned from FilterFeeChange and is used to iterate over the raw logs and unpacked data for FeeChange events raised by the TwoWay contract.
type TwoWayFeeChangeIterator struct {
	Event *TwoWayFeeChange // Event containing the contract specifics and raw log

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
func (it *TwoWayFeeChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayFeeChange)
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
		it.Event = new(TwoWayFeeChange)
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
func (it *TwoWayFeeChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayFeeChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayFeeChange represents a FeeChange event raised by the TwoWay contract.
type TwoWayFeeChange struct {
	Token     common.Address
	ChainID   *big.Int
	FeeAmount *big.Int
	FeeRatio  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeChange is a free log retrieval operation binding the contract event 0xe7de0268825882caac9be515100046b260e4bb88ef28dcf7e4d99a3b9b753782.
//
// Solidity: event FeeChange(address token, uint256 chainID, uint256 feeAmount, uint256 feeRatio)
func (_TwoWay *TwoWayFilterer) FilterFeeChange(opts *bind.FilterOpts) (*TwoWayFeeChangeIterator, error) {

	logs, sub, err := _TwoWay.contract.FilterLogs(opts, "FeeChange")
	if err != nil {
		return nil, err
	}
	return &TwoWayFeeChangeIterator{contract: _TwoWay.contract, event: "FeeChange", logs: logs, sub: sub}, nil
}

// WatchFeeChange is a free log subscription operation binding the contract event 0xe7de0268825882caac9be515100046b260e4bb88ef28dcf7e4d99a3b9b753782.
//
// Solidity: event FeeChange(address token, uint256 chainID, uint256 feeAmount, uint256 feeRatio)
func (_TwoWay *TwoWayFilterer) WatchFeeChange(opts *bind.WatchOpts, sink chan<- *TwoWayFeeChange) (event.Subscription, error) {

	logs, sub, err := _TwoWay.contract.WatchLogs(opts, "FeeChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayFeeChange)
				if err := _TwoWay.contract.UnpackLog(event, "FeeChange", log); err != nil {
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

// ParseFeeChange is a log parse operation binding the contract event 0xe7de0268825882caac9be515100046b260e4bb88ef28dcf7e4d99a3b9b753782.
//
// Solidity: event FeeChange(address token, uint256 chainID, uint256 feeAmount, uint256 feeRatio)
func (_TwoWay *TwoWayFilterer) ParseFeeChange(log types.Log) (*TwoWayFeeChange, error) {
	event := new(TwoWayFeeChange)
	if err := _TwoWay.contract.UnpackLog(event, "FeeChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayFeeToChangedIterator is returned from FilterFeeToChanged and is used to iterate over the raw logs and unpacked data for FeeToChanged events raised by the TwoWay contract.
type TwoWayFeeToChangedIterator struct {
	Event *TwoWayFeeToChanged // Event containing the contract specifics and raw log

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
func (it *TwoWayFeeToChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayFeeToChanged)
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
		it.Event = new(TwoWayFeeToChanged)
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
func (it *TwoWayFeeToChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayFeeToChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayFeeToChanged represents a FeeToChanged event raised by the TwoWay contract.
type TwoWayFeeToChanged struct {
	Token   common.Address
	ChainID *big.Int
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeToChanged is a free log retrieval operation binding the contract event 0x9963d2a39ea3765553bdc18f1222f7791b9c625670d73a9feb41fdf53158e4a4.
//
// Solidity: event FeeToChanged(address token, uint256 chainID, address account)
func (_TwoWay *TwoWayFilterer) FilterFeeToChanged(opts *bind.FilterOpts) (*TwoWayFeeToChangedIterator, error) {

	logs, sub, err := _TwoWay.contract.FilterLogs(opts, "FeeToChanged")
	if err != nil {
		return nil, err
	}
	return &TwoWayFeeToChangedIterator{contract: _TwoWay.contract, event: "FeeToChanged", logs: logs, sub: sub}, nil
}

// WatchFeeToChanged is a free log subscription operation binding the contract event 0x9963d2a39ea3765553bdc18f1222f7791b9c625670d73a9feb41fdf53158e4a4.
//
// Solidity: event FeeToChanged(address token, uint256 chainID, address account)
func (_TwoWay *TwoWayFilterer) WatchFeeToChanged(opts *bind.WatchOpts, sink chan<- *TwoWayFeeToChanged) (event.Subscription, error) {

	logs, sub, err := _TwoWay.contract.WatchLogs(opts, "FeeToChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayFeeToChanged)
				if err := _TwoWay.contract.UnpackLog(event, "FeeToChanged", log); err != nil {
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

// ParseFeeToChanged is a log parse operation binding the contract event 0x9963d2a39ea3765553bdc18f1222f7791b9c625670d73a9feb41fdf53158e4a4.
//
// Solidity: event FeeToChanged(address token, uint256 chainID, address account)
func (_TwoWay *TwoWayFilterer) ParseFeeToChanged(log types.Log) (*TwoWayFeeToChanged, error) {
	event := new(TwoWayFeeToChanged)
	if err := _TwoWay.contract.UnpackLog(event, "FeeToChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayFeeToRemovedIterator is returned from FilterFeeToRemoved and is used to iterate over the raw logs and unpacked data for FeeToRemoved events raised by the TwoWay contract.
type TwoWayFeeToRemovedIterator struct {
	Event *TwoWayFeeToRemoved // Event containing the contract specifics and raw log

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
func (it *TwoWayFeeToRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayFeeToRemoved)
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
		it.Event = new(TwoWayFeeToRemoved)
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
func (it *TwoWayFeeToRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayFeeToRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayFeeToRemoved represents a FeeToRemoved event raised by the TwoWay contract.
type TwoWayFeeToRemoved struct {
	Token   common.Address
	ChainID *big.Int
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeToRemoved is a free log retrieval operation binding the contract event 0x47ed70551b214b8794c1a97dea263c5d1cf3d03ea32f3d62333cc54035f57039.
//
// Solidity: event FeeToRemoved(address token, uint256 chainID, address account)
func (_TwoWay *TwoWayFilterer) FilterFeeToRemoved(opts *bind.FilterOpts) (*TwoWayFeeToRemovedIterator, error) {

	logs, sub, err := _TwoWay.contract.FilterLogs(opts, "FeeToRemoved")
	if err != nil {
		return nil, err
	}
	return &TwoWayFeeToRemovedIterator{contract: _TwoWay.contract, event: "FeeToRemoved", logs: logs, sub: sub}, nil
}

// WatchFeeToRemoved is a free log subscription operation binding the contract event 0x47ed70551b214b8794c1a97dea263c5d1cf3d03ea32f3d62333cc54035f57039.
//
// Solidity: event FeeToRemoved(address token, uint256 chainID, address account)
func (_TwoWay *TwoWayFilterer) WatchFeeToRemoved(opts *bind.WatchOpts, sink chan<- *TwoWayFeeToRemoved) (event.Subscription, error) {

	logs, sub, err := _TwoWay.contract.WatchLogs(opts, "FeeToRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayFeeToRemoved)
				if err := _TwoWay.contract.UnpackLog(event, "FeeToRemoved", log); err != nil {
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

// ParseFeeToRemoved is a log parse operation binding the contract event 0x47ed70551b214b8794c1a97dea263c5d1cf3d03ea32f3d62333cc54035f57039.
//
// Solidity: event FeeToRemoved(address token, uint256 chainID, address account)
func (_TwoWay *TwoWayFilterer) ParseFeeToRemoved(log types.Log) (*TwoWayFeeToRemoved, error) {
	event := new(TwoWayFeeToRemoved)
	if err := _TwoWay.contract.UnpackLog(event, "FeeToRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayLockIterator is returned from FilterLock and is used to iterate over the raw logs and unpacked data for Lock events raised by the TwoWay contract.
type TwoWayLockIterator struct {
	Event *TwoWayLock // Event containing the contract specifics and raw log

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
func (it *TwoWayLockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayLock)
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
		it.Event = new(TwoWayLock)
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
func (it *TwoWayLockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayLockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayLock represents a Lock event raised by the TwoWay contract.
type TwoWayLock struct {
	Token0   common.Address
	Token1   common.Address
	ChainID0 *big.Int
	ChainID1 *big.Int
	From     common.Address
	To       common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLock is a free log retrieval operation binding the contract event 0x7cb5e2d54d5587c3a3448631884061009b0e2c30e37922a4dd4aed50e11dd7f3.
//
// Solidity: event Lock(address token0, address token1, uint256 chainID0, uint256 chainID1, address from, address to, uint256 amount)
func (_TwoWay *TwoWayFilterer) FilterLock(opts *bind.FilterOpts) (*TwoWayLockIterator, error) {

	logs, sub, err := _TwoWay.contract.FilterLogs(opts, "Lock")
	if err != nil {
		return nil, err
	}
	return &TwoWayLockIterator{contract: _TwoWay.contract, event: "Lock", logs: logs, sub: sub}, nil
}

// WatchLock is a free log subscription operation binding the contract event 0x7cb5e2d54d5587c3a3448631884061009b0e2c30e37922a4dd4aed50e11dd7f3.
//
// Solidity: event Lock(address token0, address token1, uint256 chainID0, uint256 chainID1, address from, address to, uint256 amount)
func (_TwoWay *TwoWayFilterer) WatchLock(opts *bind.WatchOpts, sink chan<- *TwoWayLock) (event.Subscription, error) {

	logs, sub, err := _TwoWay.contract.WatchLogs(opts, "Lock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayLock)
				if err := _TwoWay.contract.UnpackLog(event, "Lock", log); err != nil {
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

// ParseLock is a log parse operation binding the contract event 0x7cb5e2d54d5587c3a3448631884061009b0e2c30e37922a4dd4aed50e11dd7f3.
//
// Solidity: event Lock(address token0, address token1, uint256 chainID0, uint256 chainID1, address from, address to, uint256 amount)
func (_TwoWay *TwoWayFilterer) ParseLock(log types.Log) (*TwoWayLock, error) {
	event := new(TwoWayLock)
	if err := _TwoWay.contract.UnpackLog(event, "Lock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the TwoWay contract.
type TwoWayProposalVotedIterator struct {
	Event *TwoWayProposalVoted // Event containing the contract specifics and raw log

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
func (it *TwoWayProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayProposalVoted)
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
		it.Event = new(TwoWayProposalVoted)
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
func (it *TwoWayProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayProposalVoted represents a ProposalVoted event raised by the TwoWay contract.
type TwoWayProposalVoted struct {
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
func (_TwoWay *TwoWayFilterer) FilterProposalVoted(opts *bind.FilterOpts) (*TwoWayProposalVotedIterator, error) {

	logs, sub, err := _TwoWay.contract.FilterLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return &TwoWayProposalVotedIterator{contract: _TwoWay.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0xe458477b94285c9f254a407e2614c39f55230b527ff7f51d76fdf33e251ae936.
//
// Solidity: event ProposalVoted(address token, address from, address to, uint256 amount, address proposer, uint256 count, uint256 threshold)
func (_TwoWay *TwoWayFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *TwoWayProposalVoted) (event.Subscription, error) {

	logs, sub, err := _TwoWay.contract.WatchLogs(opts, "ProposalVoted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayProposalVoted)
				if err := _TwoWay.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
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
func (_TwoWay *TwoWayFilterer) ParseProposalVoted(log types.Log) (*TwoWayProposalVoted, error) {
	event := new(TwoWayProposalVoted)
	if err := _TwoWay.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the TwoWay contract.
type TwoWayRoleAdminChangedIterator struct {
	Event *TwoWayRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *TwoWayRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayRoleAdminChanged)
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
		it.Event = new(TwoWayRoleAdminChanged)
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
func (it *TwoWayRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayRoleAdminChanged represents a RoleAdminChanged event raised by the TwoWay contract.
type TwoWayRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TwoWay *TwoWayFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*TwoWayRoleAdminChangedIterator, error) {

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

	logs, sub, err := _TwoWay.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &TwoWayRoleAdminChangedIterator{contract: _TwoWay.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_TwoWay *TwoWayFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *TwoWayRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _TwoWay.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayRoleAdminChanged)
				if err := _TwoWay.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_TwoWay *TwoWayFilterer) ParseRoleAdminChanged(log types.Log) (*TwoWayRoleAdminChanged, error) {
	event := new(TwoWayRoleAdminChanged)
	if err := _TwoWay.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the TwoWay contract.
type TwoWayRoleGrantedIterator struct {
	Event *TwoWayRoleGranted // Event containing the contract specifics and raw log

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
func (it *TwoWayRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayRoleGranted)
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
		it.Event = new(TwoWayRoleGranted)
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
func (it *TwoWayRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayRoleGranted represents a RoleGranted event raised by the TwoWay contract.
type TwoWayRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TwoWay *TwoWayFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TwoWayRoleGrantedIterator, error) {

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

	logs, sub, err := _TwoWay.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TwoWayRoleGrantedIterator{contract: _TwoWay.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_TwoWay *TwoWayFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *TwoWayRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TwoWay.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayRoleGranted)
				if err := _TwoWay.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_TwoWay *TwoWayFilterer) ParseRoleGranted(log types.Log) (*TwoWayRoleGranted, error) {
	event := new(TwoWayRoleGranted)
	if err := _TwoWay.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the TwoWay contract.
type TwoWayRoleRevokedIterator struct {
	Event *TwoWayRoleRevoked // Event containing the contract specifics and raw log

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
func (it *TwoWayRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayRoleRevoked)
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
		it.Event = new(TwoWayRoleRevoked)
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
func (it *TwoWayRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayRoleRevoked represents a RoleRevoked event raised by the TwoWay contract.
type TwoWayRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TwoWay *TwoWayFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*TwoWayRoleRevokedIterator, error) {

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

	logs, sub, err := _TwoWay.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &TwoWayRoleRevokedIterator{contract: _TwoWay.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_TwoWay *TwoWayFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *TwoWayRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _TwoWay.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayRoleRevoked)
				if err := _TwoWay.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_TwoWay *TwoWayFilterer) ParseRoleRevoked(log types.Log) (*TwoWayRoleRevoked, error) {
	event := new(TwoWayRoleRevoked)
	if err := _TwoWay.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayRollbackIterator is returned from FilterRollback and is used to iterate over the raw logs and unpacked data for Rollback events raised by the TwoWay contract.
type TwoWayRollbackIterator struct {
	Event *TwoWayRollback // Event containing the contract specifics and raw log

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
func (it *TwoWayRollbackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayRollback)
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
		it.Event = new(TwoWayRollback)
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
func (it *TwoWayRollbackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayRollbackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayRollback represents a Rollback event raised by the TwoWay contract.
type TwoWayRollback struct {
	Token0   common.Address
	Token1   common.Address
	ChainID0 *big.Int
	ChainID1 *big.Int
	From     common.Address
	To       common.Address
	Amount   *big.Int
	Txid     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRollback is a free log retrieval operation binding the contract event 0x56a316ee63540abb5f8fd44be7a6eb2016a90fc9bb3a34ffe13948a44c49a616.
//
// Solidity: event Rollback(address token0, address token1, uint256 chainID0, uint256 chainID1, address from, address to, uint256 amount, string txid)
func (_TwoWay *TwoWayFilterer) FilterRollback(opts *bind.FilterOpts) (*TwoWayRollbackIterator, error) {

	logs, sub, err := _TwoWay.contract.FilterLogs(opts, "Rollback")
	if err != nil {
		return nil, err
	}
	return &TwoWayRollbackIterator{contract: _TwoWay.contract, event: "Rollback", logs: logs, sub: sub}, nil
}

// WatchRollback is a free log subscription operation binding the contract event 0x56a316ee63540abb5f8fd44be7a6eb2016a90fc9bb3a34ffe13948a44c49a616.
//
// Solidity: event Rollback(address token0, address token1, uint256 chainID0, uint256 chainID1, address from, address to, uint256 amount, string txid)
func (_TwoWay *TwoWayFilterer) WatchRollback(opts *bind.WatchOpts, sink chan<- *TwoWayRollback) (event.Subscription, error) {

	logs, sub, err := _TwoWay.contract.WatchLogs(opts, "Rollback")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayRollback)
				if err := _TwoWay.contract.UnpackLog(event, "Rollback", log); err != nil {
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

// ParseRollback is a log parse operation binding the contract event 0x56a316ee63540abb5f8fd44be7a6eb2016a90fc9bb3a34ffe13948a44c49a616.
//
// Solidity: event Rollback(address token0, address token1, uint256 chainID0, uint256 chainID1, address from, address to, uint256 amount, string txid)
func (_TwoWay *TwoWayFilterer) ParseRollback(log types.Log) (*TwoWayRollback, error) {
	event := new(TwoWayRollback)
	if err := _TwoWay.contract.UnpackLog(event, "Rollback", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayRollbackedIterator is returned from FilterRollbacked and is used to iterate over the raw logs and unpacked data for Rollbacked events raised by the TwoWay contract.
type TwoWayRollbackedIterator struct {
	Event *TwoWayRollbacked // Event containing the contract specifics and raw log

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
func (it *TwoWayRollbackedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayRollbacked)
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
		it.Event = new(TwoWayRollbacked)
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
func (it *TwoWayRollbackedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayRollbackedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayRollbacked represents a Rollbacked event raised by the TwoWay contract.
type TwoWayRollbacked struct {
	Token0 common.Address
	From   common.Address
	Amount *big.Int
	Txid   string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRollbacked is a free log retrieval operation binding the contract event 0x9420db1a0c5e4f45cc0dac05f17e8eb893645981e3c01e8d14120353a487d9e9.
//
// Solidity: event Rollbacked(address token0, address from, uint256 amount, string txid)
func (_TwoWay *TwoWayFilterer) FilterRollbacked(opts *bind.FilterOpts) (*TwoWayRollbackedIterator, error) {

	logs, sub, err := _TwoWay.contract.FilterLogs(opts, "Rollbacked")
	if err != nil {
		return nil, err
	}
	return &TwoWayRollbackedIterator{contract: _TwoWay.contract, event: "Rollbacked", logs: logs, sub: sub}, nil
}

// WatchRollbacked is a free log subscription operation binding the contract event 0x9420db1a0c5e4f45cc0dac05f17e8eb893645981e3c01e8d14120353a487d9e9.
//
// Solidity: event Rollbacked(address token0, address from, uint256 amount, string txid)
func (_TwoWay *TwoWayFilterer) WatchRollbacked(opts *bind.WatchOpts, sink chan<- *TwoWayRollbacked) (event.Subscription, error) {

	logs, sub, err := _TwoWay.contract.WatchLogs(opts, "Rollbacked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayRollbacked)
				if err := _TwoWay.contract.UnpackLog(event, "Rollbacked", log); err != nil {
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

// ParseRollbacked is a log parse operation binding the contract event 0x9420db1a0c5e4f45cc0dac05f17e8eb893645981e3c01e8d14120353a487d9e9.
//
// Solidity: event Rollbacked(address token0, address from, uint256 amount, string txid)
func (_TwoWay *TwoWayFilterer) ParseRollbacked(log types.Log) (*TwoWayRollbacked, error) {
	event := new(TwoWayRollbacked)
	if err := _TwoWay.contract.UnpackLog(event, "Rollbacked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayThresholdChangedIterator is returned from FilterThresholdChanged and is used to iterate over the raw logs and unpacked data for ThresholdChanged events raised by the TwoWay contract.
type TwoWayThresholdChangedIterator struct {
	Event *TwoWayThresholdChanged // Event containing the contract specifics and raw log

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
func (it *TwoWayThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayThresholdChanged)
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
		it.Event = new(TwoWayThresholdChanged)
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
func (it *TwoWayThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayThresholdChanged represents a ThresholdChanged event raised by the TwoWay contract.
type TwoWayThresholdChanged struct {
	Token        common.Address
	OldThreshold *big.Int
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterThresholdChanged is a free log retrieval operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_TwoWay *TwoWayFilterer) FilterThresholdChanged(opts *bind.FilterOpts) (*TwoWayThresholdChangedIterator, error) {

	logs, sub, err := _TwoWay.contract.FilterLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &TwoWayThresholdChangedIterator{contract: _TwoWay.contract, event: "ThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchThresholdChanged is a free log subscription operation binding the contract event 0xb18e91516e037486aa6fa38f56a8aac933fd127180efe9c9745ecbf660a78e44.
//
// Solidity: event ThresholdChanged(address token, uint256 oldThreshold, uint256 newThreshold)
func (_TwoWay *TwoWayFilterer) WatchThresholdChanged(opts *bind.WatchOpts, sink chan<- *TwoWayThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _TwoWay.contract.WatchLogs(opts, "ThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayThresholdChanged)
				if err := _TwoWay.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
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
func (_TwoWay *TwoWayFilterer) ParseThresholdChanged(log types.Log) (*TwoWayThresholdChanged, error) {
	event := new(TwoWayThresholdChanged)
	if err := _TwoWay.contract.UnpackLog(event, "ThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayUnlockIterator is returned from FilterUnlock and is used to iterate over the raw logs and unpacked data for Unlock events raised by the TwoWay contract.
type TwoWayUnlockIterator struct {
	Event *TwoWayUnlock // Event containing the contract specifics and raw log

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
func (it *TwoWayUnlockIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayUnlock)
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
		it.Event = new(TwoWayUnlock)
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
func (it *TwoWayUnlockIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayUnlockIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayUnlock represents a Unlock event raised by the TwoWay contract.
type TwoWayUnlock struct {
	Token0   common.Address
	Token1   common.Address
	ChianID0 *big.Int
	ChainID1 *big.Int
	From     common.Address
	To       common.Address
	Amount   *big.Int
	Txid     string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUnlock is a free log retrieval operation binding the contract event 0xe8a9cddb11d86358ad5c2fdd6359f0a6f41de0d399033319becf1364409dacbc.
//
// Solidity: event Unlock(address token0, address token1, uint256 chianID0, uint256 chainID1, address from, address to, uint256 amount, string txid)
func (_TwoWay *TwoWayFilterer) FilterUnlock(opts *bind.FilterOpts) (*TwoWayUnlockIterator, error) {

	logs, sub, err := _TwoWay.contract.FilterLogs(opts, "Unlock")
	if err != nil {
		return nil, err
	}
	return &TwoWayUnlockIterator{contract: _TwoWay.contract, event: "Unlock", logs: logs, sub: sub}, nil
}

// WatchUnlock is a free log subscription operation binding the contract event 0xe8a9cddb11d86358ad5c2fdd6359f0a6f41de0d399033319becf1364409dacbc.
//
// Solidity: event Unlock(address token0, address token1, uint256 chianID0, uint256 chainID1, address from, address to, uint256 amount, string txid)
func (_TwoWay *TwoWayFilterer) WatchUnlock(opts *bind.WatchOpts, sink chan<- *TwoWayUnlock) (event.Subscription, error) {

	logs, sub, err := _TwoWay.contract.WatchLogs(opts, "Unlock")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayUnlock)
				if err := _TwoWay.contract.UnpackLog(event, "Unlock", log); err != nil {
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

// ParseUnlock is a log parse operation binding the contract event 0xe8a9cddb11d86358ad5c2fdd6359f0a6f41de0d399033319becf1364409dacbc.
//
// Solidity: event Unlock(address token0, address token1, uint256 chianID0, uint256 chainID1, address from, address to, uint256 amount, string txid)
func (_TwoWay *TwoWayFilterer) ParseUnlock(log types.Log) (*TwoWayUnlock, error) {
	event := new(TwoWayUnlock)
	if err := _TwoWay.contract.UnpackLog(event, "Unlock", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayTollABI is the input ABI used to generate the binding from.
const TwoWayTollABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeToDev\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeRatio\",\"type\":\"uint256\"}],\"name\":\"FeeChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeeToChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"FeeToRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmountFix\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeAmountRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateRemoveFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeAmountM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeRatioM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"feeTo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeToDev\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"removeFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TwoWayTollFuncSigs maps the 4-byte function signature to its string representation.
var TwoWayTollFuncSigs = map[string]string{
	"0324ef9c": "calculateFee(address,uint256,uint256)",
	"2d63c99a": "calculateRemoveFee(address,uint256)",
	"64078943": "feeAmountM(address,uint256)",
	"9c8d70be": "feeRatioM(address,uint256)",
	"db697be4": "feeTo(address,uint256)",
	"100be3bd": "feeToDev()",
	"0b19993f": "removeFeeAmount(address)",
}

// TwoWayTollBin is the compiled bytecode used for deploying new contracts.
var TwoWayTollBin = "0x608060405234801561001057600080fd5b5060405161063438038061063483398101604081905261002f91610054565b600480546001600160a01b0319166001600160a01b0392909216919091179055610082565b600060208284031215610065578081fd5b81516001600160a01b038116811461007b578182fd5b9392505050565b6105a3806100916000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80632d63c99a1161005b5780632d63c99a146100e257806364078943146101035780639c8d70be14610116578063db697be4146101295761007d565b80630324ef9c146100825780630b19993f146100ad578063100be3bd146100cd575b600080fd5b610095610090366004610338565b61013c565b6040516100a4939291906103bf565b60405180910390f35b6100c06100bb3660046102f5565b6101b6565b6040516100a491906103a8565b6100d56101c8565b6040516100a4919061036a565b6100f56100f036600461030f565b6101d7565b6040516100a49291906103b1565b6100c061011136600461030f565b610247565b6100c061012436600461030f565b610261565b6100d561013736600461030f565b61027e565b6001600160a01b03831660008181526020818152604080832086845282528083205493835260018252808320868452909152812054819061017d85826102a4565b925061018983856103d5565b85101561019557600093505b826101a08587610540565b6101aa9190610540565b91505093509350939050565b60026020526000908152604090205481565b6004546001600160a01b031681565b6001600160a01b038216600090815260026020526040812054819083116102195760405162461bcd60e51b81526004016102109061037e565b60405180910390fd5b6001600160a01b038416600090815260026020526040902054915061023e8284610540565b90509250929050565b600060208181529281526040808220909352908152205481565b600160209081526000928352604080842090915290825290205481565b60036020908152600092835260408084209091529082529020546001600160a01b031681565b60006102b26012600a610453565b6102bc84846102cd565b6102c691906103ed565b9392505050565b60006102c68284610521565b80356001600160a01b03811681146102f057600080fd5b919050565b600060208284031215610306578081fd5b6102c6826102d9565b60008060408385031215610321578081fd5b61032a836102d9565b946020939093013593505050565b60008060006060848603121561034c578081fd5b610355846102d9565b95602085013595506040909401359392505050565b6001600160a01b0391909116815260200190565b60208082526010908201526f3737ba1032b737bab3b4103a37b5b2b760811b604082015260600190565b90815260200190565b918252602082015260400190565b9283526020830191909152604082015260600190565b600082198211156103e8576103e8610557565b500190565b60008261040857634e487b7160e01b81526012600452602481fd5b500490565b80825b600180861161041f575061044a565b81870482111561043157610431610557565b8086161561043e57918102915b9490941c938002610410565b94509492505050565b60006102c6600019848460008261046c575060016102c6565b81610479575060006102c6565b816001811461048f5760028114610499576104c6565b60019150506102c6565b60ff8411156104aa576104aa610557565b6001841b9150848211156104c0576104c0610557565b506102c6565b5060208310610133831016604e8410600b84101617156104f9575081810a838111156104f4576104f4610557565b6102c6565b610506848484600161040d565b80860482111561051857610518610557565b02949350505050565b600081600019048311821515161561053b5761053b610557565b500290565b60008282101561055257610552610557565b500390565b634e487b7160e01b600052601160045260246000fdfea26469706673582212205ff3be19c2933c8c6ea5c75feddb280d6cf04e39f6b2f98fc92a0e405b773fc764736f6c63430008000033"

// DeployTwoWayToll deploys a new Ethereum contract, binding an instance of TwoWayToll to it.
func DeployTwoWayToll(auth *bind.TransactOpts, backend bind.ContractBackend, _feeToDev common.Address) (common.Address, *types.Transaction, *TwoWayToll, error) {
	parsed, err := abi.JSON(strings.NewReader(TwoWayTollABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TwoWayTollBin), backend, _feeToDev)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TwoWayToll{TwoWayTollCaller: TwoWayTollCaller{contract: contract}, TwoWayTollTransactor: TwoWayTollTransactor{contract: contract}, TwoWayTollFilterer: TwoWayTollFilterer{contract: contract}}, nil
}

// TwoWayToll is an auto generated Go binding around an Ethereum contract.
type TwoWayToll struct {
	TwoWayTollCaller     // Read-only binding to the contract
	TwoWayTollTransactor // Write-only binding to the contract
	TwoWayTollFilterer   // Log filterer for contract events
}

// TwoWayTollCaller is an auto generated read-only Go binding around an Ethereum contract.
type TwoWayTollCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoWayTollTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TwoWayTollTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoWayTollFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TwoWayTollFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TwoWayTollSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TwoWayTollSession struct {
	Contract     *TwoWayToll       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TwoWayTollCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TwoWayTollCallerSession struct {
	Contract *TwoWayTollCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TwoWayTollTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TwoWayTollTransactorSession struct {
	Contract     *TwoWayTollTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TwoWayTollRaw is an auto generated low-level Go binding around an Ethereum contract.
type TwoWayTollRaw struct {
	Contract *TwoWayToll // Generic contract binding to access the raw methods on
}

// TwoWayTollCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TwoWayTollCallerRaw struct {
	Contract *TwoWayTollCaller // Generic read-only contract binding to access the raw methods on
}

// TwoWayTollTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TwoWayTollTransactorRaw struct {
	Contract *TwoWayTollTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTwoWayToll creates a new instance of TwoWayToll, bound to a specific deployed contract.
func NewTwoWayToll(address common.Address, backend bind.ContractBackend) (*TwoWayToll, error) {
	contract, err := bindTwoWayToll(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TwoWayToll{TwoWayTollCaller: TwoWayTollCaller{contract: contract}, TwoWayTollTransactor: TwoWayTollTransactor{contract: contract}, TwoWayTollFilterer: TwoWayTollFilterer{contract: contract}}, nil
}

// NewTwoWayTollCaller creates a new read-only instance of TwoWayToll, bound to a specific deployed contract.
func NewTwoWayTollCaller(address common.Address, caller bind.ContractCaller) (*TwoWayTollCaller, error) {
	contract, err := bindTwoWayToll(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TwoWayTollCaller{contract: contract}, nil
}

// NewTwoWayTollTransactor creates a new write-only instance of TwoWayToll, bound to a specific deployed contract.
func NewTwoWayTollTransactor(address common.Address, transactor bind.ContractTransactor) (*TwoWayTollTransactor, error) {
	contract, err := bindTwoWayToll(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TwoWayTollTransactor{contract: contract}, nil
}

// NewTwoWayTollFilterer creates a new log filterer instance of TwoWayToll, bound to a specific deployed contract.
func NewTwoWayTollFilterer(address common.Address, filterer bind.ContractFilterer) (*TwoWayTollFilterer, error) {
	contract, err := bindTwoWayToll(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TwoWayTollFilterer{contract: contract}, nil
}

// bindTwoWayToll binds a generic wrapper to an already deployed contract.
func bindTwoWayToll(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TwoWayTollABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwoWayToll *TwoWayTollRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwoWayToll.Contract.TwoWayTollCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwoWayToll *TwoWayTollRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwoWayToll.Contract.TwoWayTollTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwoWayToll *TwoWayTollRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwoWayToll.Contract.TwoWayTollTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TwoWayToll *TwoWayTollCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TwoWayToll.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TwoWayToll *TwoWayTollTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TwoWayToll.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TwoWayToll *TwoWayTollTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TwoWayToll.Contract.contract.Transact(opts, method, params...)
}

// CalculateFee is a free data retrieval call binding the contract method 0x0324ef9c.
//
// Solidity: function calculateFee(address token, uint256 chainID, uint256 amount) view returns(uint256 feeAmountFix, uint256 feeAmountRatio, uint256 remainAmount)
func (_TwoWayToll *TwoWayTollCaller) CalculateFee(opts *bind.CallOpts, token common.Address, chainID *big.Int, amount *big.Int) (struct {
	FeeAmountFix   *big.Int
	FeeAmountRatio *big.Int
	RemainAmount   *big.Int
}, error) {
	var out []interface{}
	err := _TwoWayToll.contract.Call(opts, &out, "calculateFee", token, chainID, amount)

	outstruct := new(struct {
		FeeAmountFix   *big.Int
		FeeAmountRatio *big.Int
		RemainAmount   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FeeAmountFix = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FeeAmountRatio = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.RemainAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CalculateFee is a free data retrieval call binding the contract method 0x0324ef9c.
//
// Solidity: function calculateFee(address token, uint256 chainID, uint256 amount) view returns(uint256 feeAmountFix, uint256 feeAmountRatio, uint256 remainAmount)
func (_TwoWayToll *TwoWayTollSession) CalculateFee(token common.Address, chainID *big.Int, amount *big.Int) (struct {
	FeeAmountFix   *big.Int
	FeeAmountRatio *big.Int
	RemainAmount   *big.Int
}, error) {
	return _TwoWayToll.Contract.CalculateFee(&_TwoWayToll.CallOpts, token, chainID, amount)
}

// CalculateFee is a free data retrieval call binding the contract method 0x0324ef9c.
//
// Solidity: function calculateFee(address token, uint256 chainID, uint256 amount) view returns(uint256 feeAmountFix, uint256 feeAmountRatio, uint256 remainAmount)
func (_TwoWayToll *TwoWayTollCallerSession) CalculateFee(token common.Address, chainID *big.Int, amount *big.Int) (struct {
	FeeAmountFix   *big.Int
	FeeAmountRatio *big.Int
	RemainAmount   *big.Int
}, error) {
	return _TwoWayToll.Contract.CalculateFee(&_TwoWayToll.CallOpts, token, chainID, amount)
}

// CalculateRemoveFee is a free data retrieval call binding the contract method 0x2d63c99a.
//
// Solidity: function calculateRemoveFee(address token, uint256 amount) view returns(uint256 feeAmount, uint256 remainAmount)
func (_TwoWayToll *TwoWayTollCaller) CalculateRemoveFee(opts *bind.CallOpts, token common.Address, amount *big.Int) (struct {
	FeeAmount    *big.Int
	RemainAmount *big.Int
}, error) {
	var out []interface{}
	err := _TwoWayToll.contract.Call(opts, &out, "calculateRemoveFee", token, amount)

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

// CalculateRemoveFee is a free data retrieval call binding the contract method 0x2d63c99a.
//
// Solidity: function calculateRemoveFee(address token, uint256 amount) view returns(uint256 feeAmount, uint256 remainAmount)
func (_TwoWayToll *TwoWayTollSession) CalculateRemoveFee(token common.Address, amount *big.Int) (struct {
	FeeAmount    *big.Int
	RemainAmount *big.Int
}, error) {
	return _TwoWayToll.Contract.CalculateRemoveFee(&_TwoWayToll.CallOpts, token, amount)
}

// CalculateRemoveFee is a free data retrieval call binding the contract method 0x2d63c99a.
//
// Solidity: function calculateRemoveFee(address token, uint256 amount) view returns(uint256 feeAmount, uint256 remainAmount)
func (_TwoWayToll *TwoWayTollCallerSession) CalculateRemoveFee(token common.Address, amount *big.Int) (struct {
	FeeAmount    *big.Int
	RemainAmount *big.Int
}, error) {
	return _TwoWayToll.Contract.CalculateRemoveFee(&_TwoWayToll.CallOpts, token, amount)
}

// FeeAmountM is a free data retrieval call binding the contract method 0x64078943.
//
// Solidity: function feeAmountM(address , uint256 ) view returns(uint256)
func (_TwoWayToll *TwoWayTollCaller) FeeAmountM(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayToll.contract.Call(opts, &out, "feeAmountM", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeAmountM is a free data retrieval call binding the contract method 0x64078943.
//
// Solidity: function feeAmountM(address , uint256 ) view returns(uint256)
func (_TwoWayToll *TwoWayTollSession) FeeAmountM(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayToll.Contract.FeeAmountM(&_TwoWayToll.CallOpts, arg0, arg1)
}

// FeeAmountM is a free data retrieval call binding the contract method 0x64078943.
//
// Solidity: function feeAmountM(address , uint256 ) view returns(uint256)
func (_TwoWayToll *TwoWayTollCallerSession) FeeAmountM(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayToll.Contract.FeeAmountM(&_TwoWayToll.CallOpts, arg0, arg1)
}

// FeeRatioM is a free data retrieval call binding the contract method 0x9c8d70be.
//
// Solidity: function feeRatioM(address , uint256 ) view returns(uint256)
func (_TwoWayToll *TwoWayTollCaller) FeeRatioM(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayToll.contract.Call(opts, &out, "feeRatioM", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRatioM is a free data retrieval call binding the contract method 0x9c8d70be.
//
// Solidity: function feeRatioM(address , uint256 ) view returns(uint256)
func (_TwoWayToll *TwoWayTollSession) FeeRatioM(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayToll.Contract.FeeRatioM(&_TwoWayToll.CallOpts, arg0, arg1)
}

// FeeRatioM is a free data retrieval call binding the contract method 0x9c8d70be.
//
// Solidity: function feeRatioM(address , uint256 ) view returns(uint256)
func (_TwoWayToll *TwoWayTollCallerSession) FeeRatioM(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _TwoWayToll.Contract.FeeRatioM(&_TwoWayToll.CallOpts, arg0, arg1)
}

// FeeTo is a free data retrieval call binding the contract method 0xdb697be4.
//
// Solidity: function feeTo(address , uint256 ) view returns(address)
func (_TwoWayToll *TwoWayTollCaller) FeeTo(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TwoWayToll.contract.Call(opts, &out, "feeTo", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeTo is a free data retrieval call binding the contract method 0xdb697be4.
//
// Solidity: function feeTo(address , uint256 ) view returns(address)
func (_TwoWayToll *TwoWayTollSession) FeeTo(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _TwoWayToll.Contract.FeeTo(&_TwoWayToll.CallOpts, arg0, arg1)
}

// FeeTo is a free data retrieval call binding the contract method 0xdb697be4.
//
// Solidity: function feeTo(address , uint256 ) view returns(address)
func (_TwoWayToll *TwoWayTollCallerSession) FeeTo(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _TwoWayToll.Contract.FeeTo(&_TwoWayToll.CallOpts, arg0, arg1)
}

// FeeToDev is a free data retrieval call binding the contract method 0x100be3bd.
//
// Solidity: function feeToDev() view returns(address)
func (_TwoWayToll *TwoWayTollCaller) FeeToDev(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TwoWayToll.contract.Call(opts, &out, "feeToDev")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeToDev is a free data retrieval call binding the contract method 0x100be3bd.
//
// Solidity: function feeToDev() view returns(address)
func (_TwoWayToll *TwoWayTollSession) FeeToDev() (common.Address, error) {
	return _TwoWayToll.Contract.FeeToDev(&_TwoWayToll.CallOpts)
}

// FeeToDev is a free data retrieval call binding the contract method 0x100be3bd.
//
// Solidity: function feeToDev() view returns(address)
func (_TwoWayToll *TwoWayTollCallerSession) FeeToDev() (common.Address, error) {
	return _TwoWayToll.Contract.FeeToDev(&_TwoWayToll.CallOpts)
}

// RemoveFeeAmount is a free data retrieval call binding the contract method 0x0b19993f.
//
// Solidity: function removeFeeAmount(address ) view returns(uint256)
func (_TwoWayToll *TwoWayTollCaller) RemoveFeeAmount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TwoWayToll.contract.Call(opts, &out, "removeFeeAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemoveFeeAmount is a free data retrieval call binding the contract method 0x0b19993f.
//
// Solidity: function removeFeeAmount(address ) view returns(uint256)
func (_TwoWayToll *TwoWayTollSession) RemoveFeeAmount(arg0 common.Address) (*big.Int, error) {
	return _TwoWayToll.Contract.RemoveFeeAmount(&_TwoWayToll.CallOpts, arg0)
}

// RemoveFeeAmount is a free data retrieval call binding the contract method 0x0b19993f.
//
// Solidity: function removeFeeAmount(address ) view returns(uint256)
func (_TwoWayToll *TwoWayTollCallerSession) RemoveFeeAmount(arg0 common.Address) (*big.Int, error) {
	return _TwoWayToll.Contract.RemoveFeeAmount(&_TwoWayToll.CallOpts, arg0)
}

// TwoWayTollFeeChangeIterator is returned from FilterFeeChange and is used to iterate over the raw logs and unpacked data for FeeChange events raised by the TwoWayToll contract.
type TwoWayTollFeeChangeIterator struct {
	Event *TwoWayTollFeeChange // Event containing the contract specifics and raw log

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
func (it *TwoWayTollFeeChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayTollFeeChange)
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
		it.Event = new(TwoWayTollFeeChange)
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
func (it *TwoWayTollFeeChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayTollFeeChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayTollFeeChange represents a FeeChange event raised by the TwoWayToll contract.
type TwoWayTollFeeChange struct {
	Token     common.Address
	ChainID   *big.Int
	FeeAmount *big.Int
	FeeRatio  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFeeChange is a free log retrieval operation binding the contract event 0xe7de0268825882caac9be515100046b260e4bb88ef28dcf7e4d99a3b9b753782.
//
// Solidity: event FeeChange(address token, uint256 chainID, uint256 feeAmount, uint256 feeRatio)
func (_TwoWayToll *TwoWayTollFilterer) FilterFeeChange(opts *bind.FilterOpts) (*TwoWayTollFeeChangeIterator, error) {

	logs, sub, err := _TwoWayToll.contract.FilterLogs(opts, "FeeChange")
	if err != nil {
		return nil, err
	}
	return &TwoWayTollFeeChangeIterator{contract: _TwoWayToll.contract, event: "FeeChange", logs: logs, sub: sub}, nil
}

// WatchFeeChange is a free log subscription operation binding the contract event 0xe7de0268825882caac9be515100046b260e4bb88ef28dcf7e4d99a3b9b753782.
//
// Solidity: event FeeChange(address token, uint256 chainID, uint256 feeAmount, uint256 feeRatio)
func (_TwoWayToll *TwoWayTollFilterer) WatchFeeChange(opts *bind.WatchOpts, sink chan<- *TwoWayTollFeeChange) (event.Subscription, error) {

	logs, sub, err := _TwoWayToll.contract.WatchLogs(opts, "FeeChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayTollFeeChange)
				if err := _TwoWayToll.contract.UnpackLog(event, "FeeChange", log); err != nil {
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

// ParseFeeChange is a log parse operation binding the contract event 0xe7de0268825882caac9be515100046b260e4bb88ef28dcf7e4d99a3b9b753782.
//
// Solidity: event FeeChange(address token, uint256 chainID, uint256 feeAmount, uint256 feeRatio)
func (_TwoWayToll *TwoWayTollFilterer) ParseFeeChange(log types.Log) (*TwoWayTollFeeChange, error) {
	event := new(TwoWayTollFeeChange)
	if err := _TwoWayToll.contract.UnpackLog(event, "FeeChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayTollFeeToChangedIterator is returned from FilterFeeToChanged and is used to iterate over the raw logs and unpacked data for FeeToChanged events raised by the TwoWayToll contract.
type TwoWayTollFeeToChangedIterator struct {
	Event *TwoWayTollFeeToChanged // Event containing the contract specifics and raw log

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
func (it *TwoWayTollFeeToChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayTollFeeToChanged)
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
		it.Event = new(TwoWayTollFeeToChanged)
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
func (it *TwoWayTollFeeToChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayTollFeeToChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayTollFeeToChanged represents a FeeToChanged event raised by the TwoWayToll contract.
type TwoWayTollFeeToChanged struct {
	Token   common.Address
	ChainID *big.Int
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeToChanged is a free log retrieval operation binding the contract event 0x9963d2a39ea3765553bdc18f1222f7791b9c625670d73a9feb41fdf53158e4a4.
//
// Solidity: event FeeToChanged(address token, uint256 chainID, address account)
func (_TwoWayToll *TwoWayTollFilterer) FilterFeeToChanged(opts *bind.FilterOpts) (*TwoWayTollFeeToChangedIterator, error) {

	logs, sub, err := _TwoWayToll.contract.FilterLogs(opts, "FeeToChanged")
	if err != nil {
		return nil, err
	}
	return &TwoWayTollFeeToChangedIterator{contract: _TwoWayToll.contract, event: "FeeToChanged", logs: logs, sub: sub}, nil
}

// WatchFeeToChanged is a free log subscription operation binding the contract event 0x9963d2a39ea3765553bdc18f1222f7791b9c625670d73a9feb41fdf53158e4a4.
//
// Solidity: event FeeToChanged(address token, uint256 chainID, address account)
func (_TwoWayToll *TwoWayTollFilterer) WatchFeeToChanged(opts *bind.WatchOpts, sink chan<- *TwoWayTollFeeToChanged) (event.Subscription, error) {

	logs, sub, err := _TwoWayToll.contract.WatchLogs(opts, "FeeToChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayTollFeeToChanged)
				if err := _TwoWayToll.contract.UnpackLog(event, "FeeToChanged", log); err != nil {
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

// ParseFeeToChanged is a log parse operation binding the contract event 0x9963d2a39ea3765553bdc18f1222f7791b9c625670d73a9feb41fdf53158e4a4.
//
// Solidity: event FeeToChanged(address token, uint256 chainID, address account)
func (_TwoWayToll *TwoWayTollFilterer) ParseFeeToChanged(log types.Log) (*TwoWayTollFeeToChanged, error) {
	event := new(TwoWayTollFeeToChanged)
	if err := _TwoWayToll.contract.UnpackLog(event, "FeeToChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TwoWayTollFeeToRemovedIterator is returned from FilterFeeToRemoved and is used to iterate over the raw logs and unpacked data for FeeToRemoved events raised by the TwoWayToll contract.
type TwoWayTollFeeToRemovedIterator struct {
	Event *TwoWayTollFeeToRemoved // Event containing the contract specifics and raw log

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
func (it *TwoWayTollFeeToRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TwoWayTollFeeToRemoved)
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
		it.Event = new(TwoWayTollFeeToRemoved)
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
func (it *TwoWayTollFeeToRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TwoWayTollFeeToRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TwoWayTollFeeToRemoved represents a FeeToRemoved event raised by the TwoWayToll contract.
type TwoWayTollFeeToRemoved struct {
	Token   common.Address
	ChainID *big.Int
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterFeeToRemoved is a free log retrieval operation binding the contract event 0x47ed70551b214b8794c1a97dea263c5d1cf3d03ea32f3d62333cc54035f57039.
//
// Solidity: event FeeToRemoved(address token, uint256 chainID, address account)
func (_TwoWayToll *TwoWayTollFilterer) FilterFeeToRemoved(opts *bind.FilterOpts) (*TwoWayTollFeeToRemovedIterator, error) {

	logs, sub, err := _TwoWayToll.contract.FilterLogs(opts, "FeeToRemoved")
	if err != nil {
		return nil, err
	}
	return &TwoWayTollFeeToRemovedIterator{contract: _TwoWayToll.contract, event: "FeeToRemoved", logs: logs, sub: sub}, nil
}

// WatchFeeToRemoved is a free log subscription operation binding the contract event 0x47ed70551b214b8794c1a97dea263c5d1cf3d03ea32f3d62333cc54035f57039.
//
// Solidity: event FeeToRemoved(address token, uint256 chainID, address account)
func (_TwoWayToll *TwoWayTollFilterer) WatchFeeToRemoved(opts *bind.WatchOpts, sink chan<- *TwoWayTollFeeToRemoved) (event.Subscription, error) {

	logs, sub, err := _TwoWayToll.contract.WatchLogs(opts, "FeeToRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TwoWayTollFeeToRemoved)
				if err := _TwoWayToll.contract.UnpackLog(event, "FeeToRemoved", log); err != nil {
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

// ParseFeeToRemoved is a log parse operation binding the contract event 0x47ed70551b214b8794c1a97dea263c5d1cf3d03ea32f3d62333cc54035f57039.
//
// Solidity: event FeeToRemoved(address token, uint256 chainID, address account)
func (_TwoWayToll *TwoWayTollFilterer) ParseFeeToRemoved(log types.Log) (*TwoWayTollFeeToRemoved, error) {
	event := new(TwoWayTollFeeToRemoved)
	if err := _TwoWayToll.contract.UnpackLog(event, "FeeToRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
