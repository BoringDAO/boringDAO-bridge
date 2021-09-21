// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package monitor

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

// BridgetokenInfo is an auto generated low-level Go binding around an user-defined struct.
type BridgetokenInfo struct {
	Token   common.Address
	ChainId *big.Int
	TokenId *big.Int
}

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"CrossIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken0\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"originTokenUri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetChainID\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"CrossOut\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BURNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"internalType\":\"structBridge.tokenInfo\",\"name\":\"origin\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"originTokenUri\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"crossIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"crossOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isCurrentChain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenIdInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"txMinted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// BridgeFuncSigs maps the 4-byte function signature to its string representation.
var BridgeFuncSigs = map[string]string{
	"282c51f3": "BURNER_ROLE()",
	"56cf02d9": "CROSSER_ROLE()",
	"d5391393": "MINTER_ROLE()",
	"e63ab1e9": "PAUSER_ROLE()",
	"2d4f1460": "crossIn((address,uint256,uint256),string,address,address,string)",
	"df7e600a": "crossOut(address,uint256,address,uint256)",
	"e6da3850": "isCurrentChain(address)",
	"2a4f1621": "supportToken(address)",
	"aa5a310d": "tokenIdInfo(uint256)",
	"10c27402": "txMinted(string)",
}

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x608060405234801561001057600080fd5b50610596806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c8063aa5a310d11610066578063aa5a310d14610111578063d539139314610133578063df7e600a1461013b578063e63ab1e91461014e578063e6da3850146101565761009e565b806310c27402146100a3578063282c51f3146100cc5780632a4f1621146100e15780632d4f1460146100f457806356cf02d914610109575b600080fd5b6100b66100b1366004610356565b610169565b6040516100c391906104ee565b60405180910390f35b6100d4610189565b6040516100c391906104f9565b6100d46100ef3660046102ef565b6101a0565b610107610102366004610391565b6101b2565b005b6100d46101b9565b61012461011f366004610458565b6101c5565b6040516100c3939291906104cd565b6100d46101f0565b610107610149366004610311565b6101fc565b6100d4610248565b6100b66101643660046102ef565b610254565b805160208183018101805160038252928201919093012091525460ff1681565b60405161019590610488565b604051809103902081565b60016020526000908152604090205481565b5050505050565b60405161019590610470565b6000602081905290815260409020805460018201546002909201546001600160a01b03909116919083565b604051610195906104b6565b6001600160a01b0384166000908152600160205260409020548490849081146102405760405162461bcd60e51b815260040161023790610502565b60405180910390fd5b505050505050565b6040516101959061049f565b60026020526000908152604090205460ff1681565b80356001600160a01b038116811461028057600080fd5b92915050565b600082601f830112610296578081fd5b813567ffffffffffffffff8111156102ac578182fd5b6102bf601f8201601f1916602001610539565b91508082528360208285010111156102d657600080fd5b8060208401602084013760009082016020015292915050565b600060208284031215610300578081fd5b61030a8383610269565b9392505050565b60008060008060808587031215610326578283fd5b6103308686610269565b9350602085013592506103468660408701610269565b9396929550929360600135925050565b600060208284031215610367578081fd5b813567ffffffffffffffff81111561037d578182fd5b61038984828501610286565b949350505050565b600080600080600085870360e08112156103a9578182fd5b60608112156103b6578182fd5b506103c16060610539565b6103cb8888610269565b8152602087013560208201526040870135604082015280955050606086013567ffffffffffffffff808211156103ff578283fd5b61040b89838a01610286565b955061041a8960808a01610269565b94506104298960a08a01610269565b935060c088013591508082111561043e578283fd5b5061044b88828901610286565b9150509295509295909350565b600060208284031215610469578081fd5b5035919050565b6b43524f535345525f524f4c4560a01b8152600c0190565b6a4255524e45525f524f4c4560a81b8152600b0190565b6a5041555345525f524f4c4560a81b8152600b0190565b6a4d494e5445525f524f4c4560a81b8152600b0190565b6001600160a01b039390931683526020830191909152604082015260600190565b901515815260200190565b90815260200190565b60208082526017908201527f4c6f636b3a3a4e6f7420537570706f727420546f6b656e000000000000000000604082015260600190565b60405181810167ffffffffffffffff8111828210171561055857600080fd5b60405291905056fea264697066735822122034f20ee5564e5f52aca1de4410482d4c4159b48535d5ad31a9daab200443283664736f6c634300060a0033"

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

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

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) BURNERROLE() ([32]byte, error) {
	return _Bridge.Contract.BURNERROLE(&_Bridge.CallOpts)
}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) BURNERROLE() ([32]byte, error) {
	return _Bridge.Contract.BURNERROLE(&_Bridge.CallOpts)
}

// CROSSERROLE is a free data retrieval call binding the contract method 0x56cf02d9.
//
// Solidity: function CROSSER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) CROSSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "CROSSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CROSSERROLE is a free data retrieval call binding the contract method 0x56cf02d9.
//
// Solidity: function CROSSER_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) CROSSERROLE() ([32]byte, error) {
	return _Bridge.Contract.CROSSERROLE(&_Bridge.CallOpts)
}

// CROSSERROLE is a free data retrieval call binding the contract method 0x56cf02d9.
//
// Solidity: function CROSSER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) CROSSERROLE() ([32]byte, error) {
	return _Bridge.Contract.CROSSERROLE(&_Bridge.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) MINTERROLE() ([32]byte, error) {
	return _Bridge.Contract.MINTERROLE(&_Bridge.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) MINTERROLE() ([32]byte, error) {
	return _Bridge.Contract.MINTERROLE(&_Bridge.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) PAUSERROLE() ([32]byte, error) {
	return _Bridge.Contract.PAUSERROLE(&_Bridge.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Bridge.Contract.PAUSERROLE(&_Bridge.CallOpts)
}

// IsCurrentChain is a free data retrieval call binding the contract method 0xe6da3850.
//
// Solidity: function isCurrentChain(address ) view returns(bool)
func (_Bridge *BridgeCaller) IsCurrentChain(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isCurrentChain", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCurrentChain is a free data retrieval call binding the contract method 0xe6da3850.
//
// Solidity: function isCurrentChain(address ) view returns(bool)
func (_Bridge *BridgeSession) IsCurrentChain(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.IsCurrentChain(&_Bridge.CallOpts, arg0)
}

// IsCurrentChain is a free data retrieval call binding the contract method 0xe6da3850.
//
// Solidity: function isCurrentChain(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) IsCurrentChain(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.IsCurrentChain(&_Bridge.CallOpts, arg0)
}

// SupportToken is a free data retrieval call binding the contract method 0x2a4f1621.
//
// Solidity: function supportToken(address ) view returns(uint256)
func (_Bridge *BridgeCaller) SupportToken(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "supportToken", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupportToken is a free data retrieval call binding the contract method 0x2a4f1621.
//
// Solidity: function supportToken(address ) view returns(uint256)
func (_Bridge *BridgeSession) SupportToken(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.SupportToken(&_Bridge.CallOpts, arg0)
}

// SupportToken is a free data retrieval call binding the contract method 0x2a4f1621.
//
// Solidity: function supportToken(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) SupportToken(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.SupportToken(&_Bridge.CallOpts, arg0)
}

// TokenIdInfo is a free data retrieval call binding the contract method 0xaa5a310d.
//
// Solidity: function tokenIdInfo(uint256 ) view returns(address token, uint256 chainId, uint256 tokenId)
func (_Bridge *BridgeCaller) TokenIdInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Token   common.Address
	ChainId *big.Int
	TokenId *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "tokenIdInfo", arg0)

	outstruct := new(struct {
		Token   common.Address
		ChainId *big.Int
		TokenId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ChainId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TokenId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenIdInfo is a free data retrieval call binding the contract method 0xaa5a310d.
//
// Solidity: function tokenIdInfo(uint256 ) view returns(address token, uint256 chainId, uint256 tokenId)
func (_Bridge *BridgeSession) TokenIdInfo(arg0 *big.Int) (struct {
	Token   common.Address
	ChainId *big.Int
	TokenId *big.Int
}, error) {
	return _Bridge.Contract.TokenIdInfo(&_Bridge.CallOpts, arg0)
}

// TokenIdInfo is a free data retrieval call binding the contract method 0xaa5a310d.
//
// Solidity: function tokenIdInfo(uint256 ) view returns(address token, uint256 chainId, uint256 tokenId)
func (_Bridge *BridgeCallerSession) TokenIdInfo(arg0 *big.Int) (struct {
	Token   common.Address
	ChainId *big.Int
	TokenId *big.Int
}, error) {
	return _Bridge.Contract.TokenIdInfo(&_Bridge.CallOpts, arg0)
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

// CrossIn is a paid mutator transaction binding the contract method 0x2d4f1460.
//
// Solidity: function crossIn((address,uint256,uint256) origin, string originTokenUri, address from, address to, string txid) returns()
func (_Bridge *BridgeTransactor) CrossIn(opts *bind.TransactOpts, origin BridgetokenInfo, originTokenUri string, from common.Address, to common.Address, txid string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "crossIn", origin, originTokenUri, from, to, txid)
}

// CrossIn is a paid mutator transaction binding the contract method 0x2d4f1460.
//
// Solidity: function crossIn((address,uint256,uint256) origin, string originTokenUri, address from, address to, string txid) returns()
func (_Bridge *BridgeSession) CrossIn(origin BridgetokenInfo, originTokenUri string, from common.Address, to common.Address, txid string) (*types.Transaction, error) {
	return _Bridge.Contract.CrossIn(&_Bridge.TransactOpts, origin, originTokenUri, from, to, txid)
}

// CrossIn is a paid mutator transaction binding the contract method 0x2d4f1460.
//
// Solidity: function crossIn((address,uint256,uint256) origin, string originTokenUri, address from, address to, string txid) returns()
func (_Bridge *BridgeTransactorSession) CrossIn(origin BridgetokenInfo, originTokenUri string, from common.Address, to common.Address, txid string) (*types.Transaction, error) {
	return _Bridge.Contract.CrossIn(&_Bridge.TransactOpts, origin, originTokenUri, from, to, txid)
}

// CrossOut is a paid mutator transaction binding the contract method 0xdf7e600a.
//
// Solidity: function crossOut(address token0, uint256 chainID, address to, uint256 tokenId) returns()
func (_Bridge *BridgeTransactor) CrossOut(opts *bind.TransactOpts, token0 common.Address, chainID *big.Int, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "crossOut", token0, chainID, to, tokenId)
}

// CrossOut is a paid mutator transaction binding the contract method 0xdf7e600a.
//
// Solidity: function crossOut(address token0, uint256 chainID, address to, uint256 tokenId) returns()
func (_Bridge *BridgeSession) CrossOut(token0 common.Address, chainID *big.Int, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.CrossOut(&_Bridge.TransactOpts, token0, chainID, to, tokenId)
}

// CrossOut is a paid mutator transaction binding the contract method 0xdf7e600a.
//
// Solidity: function crossOut(address token0, uint256 chainID, address to, uint256 tokenId) returns()
func (_Bridge *BridgeTransactorSession) CrossOut(token0 common.Address, chainID *big.Int, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.CrossOut(&_Bridge.TransactOpts, token0, chainID, to, tokenId)
}

// BridgeCrossInIterator is returned from FilterCrossIn and is used to iterate over the raw logs and unpacked data for CrossIn events raised by the Bridge contract.
type BridgeCrossInIterator struct {
	Event *BridgeCrossIn // Event containing the contract specifics and raw log

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
func (it *BridgeCrossInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossIn)
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
		it.Event = new(BridgeCrossIn)
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
func (it *BridgeCrossInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossIn represents a CrossIn event raised by the Bridge contract.
type BridgeCrossIn struct {
	TargetTokenId *big.Int
	From          common.Address
	To            common.Address
	Txid          string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCrossIn is a free log retrieval operation binding the contract event 0x71533e072bdd847386062709b3aaed09429ad75a8851b6c03a113464c9e7bae2.
//
// Solidity: event CrossIn(uint256 targetTokenId, address from, address to, string txid)
func (_Bridge *BridgeFilterer) FilterCrossIn(opts *bind.FilterOpts) (*BridgeCrossInIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "CrossIn")
	if err != nil {
		return nil, err
	}
	return &BridgeCrossInIterator{contract: _Bridge.contract, event: "CrossIn", logs: logs, sub: sub}, nil
}

// WatchCrossIn is a free log subscription operation binding the contract event 0x71533e072bdd847386062709b3aaed09429ad75a8851b6c03a113464c9e7bae2.
//
// Solidity: event CrossIn(uint256 targetTokenId, address from, address to, string txid)
func (_Bridge *BridgeFilterer) WatchCrossIn(opts *bind.WatchOpts, sink chan<- *BridgeCrossIn) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "CrossIn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossIn)
				if err := _Bridge.contract.UnpackLog(event, "CrossIn", log); err != nil {
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

// ParseCrossIn is a log parse operation binding the contract event 0x71533e072bdd847386062709b3aaed09429ad75a8851b6c03a113464c9e7bae2.
//
// Solidity: event CrossIn(uint256 targetTokenId, address from, address to, string txid)
func (_Bridge *BridgeFilterer) ParseCrossIn(log types.Log) (*BridgeCrossIn, error) {
	event := new(BridgeCrossIn)
	if err := _Bridge.contract.UnpackLog(event, "CrossIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeCrossOutIterator is returned from FilterCrossOut and is used to iterate over the raw logs and unpacked data for CrossOut events raised by the Bridge contract.
type BridgeCrossOutIterator struct {
	Event *BridgeCrossOut // Event containing the contract specifics and raw log

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
func (it *BridgeCrossOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeCrossOut)
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
		it.Event = new(BridgeCrossOut)
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
func (it *BridgeCrossOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeCrossOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeCrossOut represents a CrossOut event raised by the Bridge contract.
type BridgeCrossOut struct {
	OriginToken0   common.Address
	OriginChainId  *big.Int
	OriginTokenId  *big.Int
	OriginTokenUri string
	TargetChainID  *big.Int
	From           common.Address
	To             common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCrossOut is a free log retrieval operation binding the contract event 0x4104a7d518c3ed0c0bdedf54d7d6da571c258742233f72e444f5273bf2deaa88.
//
// Solidity: event CrossOut(address originToken0, uint256 originChainId, uint256 originTokenId, string originTokenUri, uint256 targetChainID, address from, address to)
func (_Bridge *BridgeFilterer) FilterCrossOut(opts *bind.FilterOpts) (*BridgeCrossOutIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "CrossOut")
	if err != nil {
		return nil, err
	}
	return &BridgeCrossOutIterator{contract: _Bridge.contract, event: "CrossOut", logs: logs, sub: sub}, nil
}

// WatchCrossOut is a free log subscription operation binding the contract event 0x4104a7d518c3ed0c0bdedf54d7d6da571c258742233f72e444f5273bf2deaa88.
//
// Solidity: event CrossOut(address originToken0, uint256 originChainId, uint256 originTokenId, string originTokenUri, uint256 targetChainID, address from, address to)
func (_Bridge *BridgeFilterer) WatchCrossOut(opts *bind.WatchOpts, sink chan<- *BridgeCrossOut) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "CrossOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeCrossOut)
				if err := _Bridge.contract.UnpackLog(event, "CrossOut", log); err != nil {
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

// ParseCrossOut is a log parse operation binding the contract event 0x4104a7d518c3ed0c0bdedf54d7d6da571c258742233f72e444f5273bf2deaa88.
//
// Solidity: event CrossOut(address originToken0, uint256 originChainId, uint256 originTokenId, string originTokenUri, uint256 targetChainID, address from, address to)
func (_Bridge *BridgeFilterer) ParseCrossOut(log types.Log) (*BridgeCrossOut, error) {
	event := new(BridgeCrossOut)
	if err := _Bridge.contract.UnpackLog(event, "CrossOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
