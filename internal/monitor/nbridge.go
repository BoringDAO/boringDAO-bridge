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

// NBridgeTokenInfo is an auto generated low-level Go binding around an user-defined struct.
type NBridgeTokenInfo struct {
	TokenType     *big.Int
	MirrorAddress common.Address
	MirrorChainId *big.Int
	IsSupported   bool
}

// NBridgeABI is the input ABI used to generate the binding from.
const NBridgeABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CrossIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CrossOut\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_chainIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mirrorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mirrorChainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"internalType\":\"structNBridge.TokenInfo[]\",\"name\":\"tis\",\"type\":\"tuple[]\"}],\"name\":\"addMultiSupportTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mirrorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mirrorChainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"internalType\":\"structNBridge.TokenInfo\",\"name\":\"ti\",\"type\":\"tuple\"}],\"name\":\"addSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_originToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_originChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"name\":\"crossIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"crossOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mirrorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mirrorChainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"txHandled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NBridgeFuncSigs maps the 4-byte function signature to its string representation.
var NBridgeFuncSigs = map[string]string{
	"11dc39d3": "addMultiSupportTokens(uint256[],address[],(uint256,address,uint256,bool)[])",
	"ebf3c976": "addSupportToken(uint256,address,(uint256,address,uint256,bool))",
	"9a8a0592": "chainId()",
	"ab01c038": "crossIn(address,uint256,uint256,uint256,address,address,uint256,string)",
	"df7e600a": "crossOut(address,uint256,address,uint256)",
	"9d879990": "setThreshold(address,uint256)",
	"15431c48": "supportedTokens(uint256,address)",
	"824df449": "txHandled(string)",
}

// NBridgeBin is the compiled bytecode used for deploying new contracts.
var NBridgeBin = "0x608060405234801561001057600080fd5b50604051610c44380380610c4483398101604081905261002f91610037565b600055610050565b60006020828403121561004957600080fd5b5051919050565b610be58061005f6000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80639d8799901161005b5780639d87999014610171578063ab01c03814610183578063df7e600a14610196578063ebf3c976146101a957600080fd5b806311dc39d31461008d57806315431c48146100a2578063824df4491461011c5780639a8a05921461015a575b600080fd5b6100a061009b3660046108fe565b6101bc565b005b6100ed6100b0366004610a1f565b6001602081815260009384526040808520909152918352912080549181015460028201546003909201546001600160a01b03909116919060ff1684565b604080519485526001600160a01b03909316602085015291830152151560608201526080015b60405180910390f35b61014a61012a3660046109e2565b805160208183018101805160028252928201919093012091525460ff1681565b6040519015158152602001610113565b61016360005481565b604051908152602001610113565b6100a061017f3660046107fa565b5050565b6100a0610191366004610868565b610280565b6100a06101a4366004610824565b6103c5565b6100a06101b7366004610a4b565b610514565b8151815181146102065760405162461bcd60e51b815260206004820152601060248201526f0d8cadccee8d040dcdee840dac2e8c6d60831b60448201526064015b60405180910390fd5b60005b818110156102795761026785828151811061022657610226610b83565b602002602001015185838151811061024057610240610b83565b602002602001015185848151811061025a5761025a610b83565b6020026020010151610514565b8061027181610b5a565b915050610209565b5050505050565b806002816040516102919190610a88565b9081526040519081900360200190205460ff16156102de5760405162461bcd60e51b815260206004820152600a6024820152691d1e081a185b991b195960b21b60448201526064016101fd565b600054861461031f5760405162461bcd60e51b815260206004820152600d60248201526c31b430b4b724b21032b93937b960991b60448201526064016101fd565b60008881526001602081815260408084206001600160a01b03808f16865290835293819020815160808101835281548152938101549094169183019190915260028301549082015260039091015460ff161515606082018190526103b95760405162461bcd60e51b81526020600482015260116024820152703737ba1039bab83837b93a103a37b5b2b760791b60448201526064016101fd565b50505050505050505050565b6000805481526001602081815260408084206001600160a01b03808a16865290835293819020815160808101835281548152938101549094169183019190915260028301549082015260039091015460ff161515606082018190526104605760405162461bcd60e51b81526020600482015260116024820152703737ba1039bab83837b93a103a37b5b2b760791b60448201526064016101fd565b8051600114156104b6577f549f0ee31ecfdd93c4d6c1591746373d18f77138d790e2424e759e4f3b88798a85600054600054873388886040516104a99796959493929190610ac3565b60405180910390a1610279565b805160021415610279577f549f0ee31ecfdd93c4d6c1591746373d18f77138d790e2424e759e4f3b88798a81602001518260400151600054873388886040516105059796959493929190610ac3565b60405180910390a15050505050565b6001600160a01b0382166105605760405162461bcd60e51b81526020600482015260136024820152726572726f7220746f6b656e206164647265737360681b60448201526064016101fd565b606081015115156001146105a55760405162461bcd60e51b815260206004820152600c60248201526b3830b930b6b99032b93937b960a11b60448201526064016101fd565b60009283526001602081815260408086206001600160a01b039586168752825294859020835181559083015191810180546001600160a01b03191692909416919091179092559182015160028201556060909101516003909101805460ff1916911515919091179055565b80356001600160a01b038116811461062757600080fd5b919050565b600082601f83011261063d57600080fd5b8135602061065261064d83610b36565b610b05565b80838252828201915082860187848660051b890101111561067257600080fd5b60005b858110156106985761068682610610565b84529284019290840190600101610675565b5090979650505050505050565b600082601f8301126106b657600080fd5b813560206106c661064d83610b36565b80838252828201915082860187848660071b89010111156106e657600080fd5b60005b85811015610698576106fb8983610780565b845292840192608091909101906001016106e9565b600082601f83011261072157600080fd5b813567ffffffffffffffff81111561073b5761073b610b99565b61074e601f8201601f1916602001610b05565b81815284602083860101111561076357600080fd5b816020850160208301376000918101602001919091529392505050565b60006080828403121561079257600080fd5b6040516080810181811067ffffffffffffffff821117156107b5576107b5610b99565b604052823581529050806107cb60208401610610565b602082015260408301356040820152606083013580151581146107ed57600080fd5b6060919091015292915050565b6000806040838503121561080d57600080fd5b61081683610610565b946020939093013593505050565b6000806000806080858703121561083a57600080fd5b61084385610610565b93506020850135925061085860408601610610565b9396929550929360600135925050565b600080600080600080600080610100898b03121561088557600080fd5b61088e89610610565b97506020890135965060408901359550606089013594506108b160808a01610610565b93506108bf60a08a01610610565b925060c0890135915060e089013567ffffffffffffffff8111156108e257600080fd5b6108ee8b828c01610710565b9150509295985092959890939650565b60008060006060848603121561091357600080fd5b833567ffffffffffffffff8082111561092b57600080fd5b818601915086601f83011261093f57600080fd5b8135602061094f61064d83610b36565b8083825282820191508286018b848660051b890101111561096f57600080fd5b600096505b84871015610992578035835260019690960195918301918301610974565b50975050870135925050808211156109a957600080fd5b6109b58783880161062c565b935060408601359150808211156109cb57600080fd5b506109d8868287016106a5565b9150509250925092565b6000602082840312156109f457600080fd5b813567ffffffffffffffff811115610a0b57600080fd5b610a1784828501610710565b949350505050565b60008060408385031215610a3257600080fd5b82359150610a4260208401610610565b90509250929050565b600080600060c08486031215610a6057600080fd5b83359250610a7060208501610610565b9150610a7f8560408601610780565b90509250925092565b6000825160005b81811015610aa95760208186018101518583015201610a8f565b81811115610ab8576000828501525b509190910192915050565b6001600160a01b0397881681526020810196909652604086019490945260608501929092528416608084015290921660a082015260c081019190915260e00190565b604051601f8201601f1916810167ffffffffffffffff81118282101715610b2e57610b2e610b99565b604052919050565b600067ffffffffffffffff821115610b5057610b50610b99565b5060051b60200190565b6000600019821415610b7c57634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea26469706673582212203127b834b6c19f381a844c81fe9664c7a331b362b0141db6399348bb1f3146e564736f6c63430008060033"

// DeployNBridge deploys a new Ethereum contract, binding an instance of NBridge to it.
func DeployNBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _chainId *big.Int) (common.Address, *types.Transaction, *NBridge, error) {
	parsed, err := abi.JSON(strings.NewReader(NBridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NBridgeBin), backend, _chainId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NBridge{NBridgeCaller: NBridgeCaller{contract: contract}, NBridgeTransactor: NBridgeTransactor{contract: contract}, NBridgeFilterer: NBridgeFilterer{contract: contract}}, nil
}

// NBridge is an auto generated Go binding around an Ethereum contract.
type NBridge struct {
	NBridgeCaller     // Read-only binding to the contract
	NBridgeTransactor // Write-only binding to the contract
	NBridgeFilterer   // Log filterer for contract events
}

// NBridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NBridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NBridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NBridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NBridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NBridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NBridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NBridgeSession struct {
	Contract     *NBridge          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NBridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NBridgeCallerSession struct {
	Contract *NBridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// NBridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NBridgeTransactorSession struct {
	Contract     *NBridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// NBridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NBridgeRaw struct {
	Contract *NBridge // Generic contract binding to access the raw methods on
}

// NBridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NBridgeCallerRaw struct {
	Contract *NBridgeCaller // Generic read-only contract binding to access the raw methods on
}

// NBridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NBridgeTransactorRaw struct {
	Contract *NBridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNBridge creates a new instance of NBridge, bound to a specific deployed contract.
func NewNBridge(address common.Address, backend bind.ContractBackend) (*NBridge, error) {
	contract, err := bindNBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NBridge{NBridgeCaller: NBridgeCaller{contract: contract}, NBridgeTransactor: NBridgeTransactor{contract: contract}, NBridgeFilterer: NBridgeFilterer{contract: contract}}, nil
}

// NewNBridgeCaller creates a new read-only instance of NBridge, bound to a specific deployed contract.
func NewNBridgeCaller(address common.Address, caller bind.ContractCaller) (*NBridgeCaller, error) {
	contract, err := bindNBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NBridgeCaller{contract: contract}, nil
}

// NewNBridgeTransactor creates a new write-only instance of NBridge, bound to a specific deployed contract.
func NewNBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*NBridgeTransactor, error) {
	contract, err := bindNBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NBridgeTransactor{contract: contract}, nil
}

// NewNBridgeFilterer creates a new log filterer instance of NBridge, bound to a specific deployed contract.
func NewNBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*NBridgeFilterer, error) {
	contract, err := bindNBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NBridgeFilterer{contract: contract}, nil
}

// bindNBridge binds a generic wrapper to an already deployed contract.
func bindNBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NBridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NBridge *NBridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NBridge.Contract.NBridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NBridge *NBridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NBridge.Contract.NBridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NBridge *NBridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NBridge.Contract.NBridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NBridge *NBridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NBridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NBridge *NBridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NBridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NBridge *NBridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NBridge.Contract.contract.Transact(opts, method, params...)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_NBridge *NBridgeCaller) ChainId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NBridge.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_NBridge *NBridgeSession) ChainId() (*big.Int, error) {
	return _NBridge.Contract.ChainId(&_NBridge.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint256)
func (_NBridge *NBridgeCallerSession) ChainId() (*big.Int, error) {
	return _NBridge.Contract.ChainId(&_NBridge.CallOpts)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x15431c48.
//
// Solidity: function supportedTokens(uint256 , address ) view returns(uint256 tokenType, address mirrorAddress, uint256 mirrorChainId, bool isSupported)
func (_NBridge *NBridgeCaller) SupportedTokens(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	TokenType     *big.Int
	MirrorAddress common.Address
	MirrorChainId *big.Int
	IsSupported   bool
}, error) {
	var out []interface{}
	err := _NBridge.contract.Call(opts, &out, "supportedTokens", arg0, arg1)

	outstruct := new(struct {
		TokenType     *big.Int
		MirrorAddress common.Address
		MirrorChainId *big.Int
		IsSupported   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenType = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MirrorAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.MirrorChainId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsSupported = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// SupportedTokens is a free data retrieval call binding the contract method 0x15431c48.
//
// Solidity: function supportedTokens(uint256 , address ) view returns(uint256 tokenType, address mirrorAddress, uint256 mirrorChainId, bool isSupported)
func (_NBridge *NBridgeSession) SupportedTokens(arg0 *big.Int, arg1 common.Address) (struct {
	TokenType     *big.Int
	MirrorAddress common.Address
	MirrorChainId *big.Int
	IsSupported   bool
}, error) {
	return _NBridge.Contract.SupportedTokens(&_NBridge.CallOpts, arg0, arg1)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x15431c48.
//
// Solidity: function supportedTokens(uint256 , address ) view returns(uint256 tokenType, address mirrorAddress, uint256 mirrorChainId, bool isSupported)
func (_NBridge *NBridgeCallerSession) SupportedTokens(arg0 *big.Int, arg1 common.Address) (struct {
	TokenType     *big.Int
	MirrorAddress common.Address
	MirrorChainId *big.Int
	IsSupported   bool
}, error) {
	return _NBridge.Contract.SupportedTokens(&_NBridge.CallOpts, arg0, arg1)
}

// TxHandled is a free data retrieval call binding the contract method 0x824df449.
//
// Solidity: function txHandled(string ) view returns(bool)
func (_NBridge *NBridgeCaller) TxHandled(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _NBridge.contract.Call(opts, &out, "txHandled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TxHandled is a free data retrieval call binding the contract method 0x824df449.
//
// Solidity: function txHandled(string ) view returns(bool)
func (_NBridge *NBridgeSession) TxHandled(arg0 string) (bool, error) {
	return _NBridge.Contract.TxHandled(&_NBridge.CallOpts, arg0)
}

// TxHandled is a free data retrieval call binding the contract method 0x824df449.
//
// Solidity: function txHandled(string ) view returns(bool)
func (_NBridge *NBridgeCallerSession) TxHandled(arg0 string) (bool, error) {
	return _NBridge.Contract.TxHandled(&_NBridge.CallOpts, arg0)
}

// AddMultiSupportTokens is a paid mutator transaction binding the contract method 0x11dc39d3.
//
// Solidity: function addMultiSupportTokens(uint256[] _chainIds, address[] _tokens, (uint256,address,uint256,bool)[] tis) returns()
func (_NBridge *NBridgeTransactor) AddMultiSupportTokens(opts *bind.TransactOpts, _chainIds []*big.Int, _tokens []common.Address, tis []NBridgeTokenInfo) (*types.Transaction, error) {
	return _NBridge.contract.Transact(opts, "addMultiSupportTokens", _chainIds, _tokens, tis)
}

// AddMultiSupportTokens is a paid mutator transaction binding the contract method 0x11dc39d3.
//
// Solidity: function addMultiSupportTokens(uint256[] _chainIds, address[] _tokens, (uint256,address,uint256,bool)[] tis) returns()
func (_NBridge *NBridgeSession) AddMultiSupportTokens(_chainIds []*big.Int, _tokens []common.Address, tis []NBridgeTokenInfo) (*types.Transaction, error) {
	return _NBridge.Contract.AddMultiSupportTokens(&_NBridge.TransactOpts, _chainIds, _tokens, tis)
}

// AddMultiSupportTokens is a paid mutator transaction binding the contract method 0x11dc39d3.
//
// Solidity: function addMultiSupportTokens(uint256[] _chainIds, address[] _tokens, (uint256,address,uint256,bool)[] tis) returns()
func (_NBridge *NBridgeTransactorSession) AddMultiSupportTokens(_chainIds []*big.Int, _tokens []common.Address, tis []NBridgeTokenInfo) (*types.Transaction, error) {
	return _NBridge.Contract.AddMultiSupportTokens(&_NBridge.TransactOpts, _chainIds, _tokens, tis)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0xebf3c976.
//
// Solidity: function addSupportToken(uint256 _chainId, address _token, (uint256,address,uint256,bool) ti) returns()
func (_NBridge *NBridgeTransactor) AddSupportToken(opts *bind.TransactOpts, _chainId *big.Int, _token common.Address, ti NBridgeTokenInfo) (*types.Transaction, error) {
	return _NBridge.contract.Transact(opts, "addSupportToken", _chainId, _token, ti)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0xebf3c976.
//
// Solidity: function addSupportToken(uint256 _chainId, address _token, (uint256,address,uint256,bool) ti) returns()
func (_NBridge *NBridgeSession) AddSupportToken(_chainId *big.Int, _token common.Address, ti NBridgeTokenInfo) (*types.Transaction, error) {
	return _NBridge.Contract.AddSupportToken(&_NBridge.TransactOpts, _chainId, _token, ti)
}

// AddSupportToken is a paid mutator transaction binding the contract method 0xebf3c976.
//
// Solidity: function addSupportToken(uint256 _chainId, address _token, (uint256,address,uint256,bool) ti) returns()
func (_NBridge *NBridgeTransactorSession) AddSupportToken(_chainId *big.Int, _token common.Address, ti NBridgeTokenInfo) (*types.Transaction, error) {
	return _NBridge.Contract.AddSupportToken(&_NBridge.TransactOpts, _chainId, _token, ti)
}

// CrossIn is a paid mutator transaction binding the contract method 0xab01c038.
//
// Solidity: function crossIn(address _originToken, uint256 _originChainId, uint256 fromChainId, uint256 toChainId, address from, address to, uint256 amount, string txid) returns()
func (_NBridge *NBridgeTransactor) CrossIn(opts *bind.TransactOpts, _originToken common.Address, _originChainId *big.Int, fromChainId *big.Int, toChainId *big.Int, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _NBridge.contract.Transact(opts, "crossIn", _originToken, _originChainId, fromChainId, toChainId, from, to, amount, txid)
}

// CrossIn is a paid mutator transaction binding the contract method 0xab01c038.
//
// Solidity: function crossIn(address _originToken, uint256 _originChainId, uint256 fromChainId, uint256 toChainId, address from, address to, uint256 amount, string txid) returns()
func (_NBridge *NBridgeSession) CrossIn(_originToken common.Address, _originChainId *big.Int, fromChainId *big.Int, toChainId *big.Int, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _NBridge.Contract.CrossIn(&_NBridge.TransactOpts, _originToken, _originChainId, fromChainId, toChainId, from, to, amount, txid)
}

// CrossIn is a paid mutator transaction binding the contract method 0xab01c038.
//
// Solidity: function crossIn(address _originToken, uint256 _originChainId, uint256 fromChainId, uint256 toChainId, address from, address to, uint256 amount, string txid) returns()
func (_NBridge *NBridgeTransactorSession) CrossIn(_originToken common.Address, _originChainId *big.Int, fromChainId *big.Int, toChainId *big.Int, from common.Address, to common.Address, amount *big.Int, txid string) (*types.Transaction, error) {
	return _NBridge.Contract.CrossIn(&_NBridge.TransactOpts, _originToken, _originChainId, fromChainId, toChainId, from, to, amount, txid)
}

// CrossOut is a paid mutator transaction binding the contract method 0xdf7e600a.
//
// Solidity: function crossOut(address _token, uint256 toChainID, address to, uint256 amount) returns()
func (_NBridge *NBridgeTransactor) CrossOut(opts *bind.TransactOpts, _token common.Address, toChainID *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NBridge.contract.Transact(opts, "crossOut", _token, toChainID, to, amount)
}

// CrossOut is a paid mutator transaction binding the contract method 0xdf7e600a.
//
// Solidity: function crossOut(address _token, uint256 toChainID, address to, uint256 amount) returns()
func (_NBridge *NBridgeSession) CrossOut(_token common.Address, toChainID *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NBridge.Contract.CrossOut(&_NBridge.TransactOpts, _token, toChainID, to, amount)
}

// CrossOut is a paid mutator transaction binding the contract method 0xdf7e600a.
//
// Solidity: function crossOut(address _token, uint256 toChainID, address to, uint256 amount) returns()
func (_NBridge *NBridgeTransactorSession) CrossOut(_token common.Address, toChainID *big.Int, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NBridge.Contract.CrossOut(&_NBridge.TransactOpts, _token, toChainID, to, amount)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_NBridge *NBridgeTransactor) SetThreshold(opts *bind.TransactOpts, token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _NBridge.contract.Transact(opts, "setThreshold", token0, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_NBridge *NBridgeSession) SetThreshold(token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _NBridge.Contract.SetThreshold(&_NBridge.TransactOpts, token0, _threshold)
}

// SetThreshold is a paid mutator transaction binding the contract method 0x9d879990.
//
// Solidity: function setThreshold(address token0, uint256 _threshold) returns()
func (_NBridge *NBridgeTransactorSession) SetThreshold(token0 common.Address, _threshold *big.Int) (*types.Transaction, error) {
	return _NBridge.Contract.SetThreshold(&_NBridge.TransactOpts, token0, _threshold)
}

// NBridgeCrossInIterator is returned from FilterCrossIn and is used to iterate over the raw logs and unpacked data for CrossIn events raised by the NBridge contract.
type NBridgeCrossInIterator struct {
	Event *NBridgeCrossIn // Event containing the contract specifics and raw log

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
func (it *NBridgeCrossInIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NBridgeCrossIn)
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
		it.Event = new(NBridgeCrossIn)
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
func (it *NBridgeCrossInIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NBridgeCrossInIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NBridgeCrossIn represents a CrossIn event raised by the NBridge contract.
type NBridgeCrossIn struct {
	OriginToken   common.Address
	OriginChainId *big.Int
	FromChainId   *big.Int
	ToChainId     *big.Int
	From          common.Address
	To            common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCrossIn is a free log retrieval operation binding the contract event 0xe69376bccafdb809f685ce8a90249211489ab0ec624c2faf6c91a22f8dddad5f.
//
// Solidity: event CrossIn(address originToken, uint256 originChainId, uint256 fromChainId, uint256 toChainId, address from, address to, uint256 amount)
func (_NBridge *NBridgeFilterer) FilterCrossIn(opts *bind.FilterOpts) (*NBridgeCrossInIterator, error) {

	logs, sub, err := _NBridge.contract.FilterLogs(opts, "CrossIn")
	if err != nil {
		return nil, err
	}
	return &NBridgeCrossInIterator{contract: _NBridge.contract, event: "CrossIn", logs: logs, sub: sub}, nil
}

// WatchCrossIn is a free log subscription operation binding the contract event 0xe69376bccafdb809f685ce8a90249211489ab0ec624c2faf6c91a22f8dddad5f.
//
// Solidity: event CrossIn(address originToken, uint256 originChainId, uint256 fromChainId, uint256 toChainId, address from, address to, uint256 amount)
func (_NBridge *NBridgeFilterer) WatchCrossIn(opts *bind.WatchOpts, sink chan<- *NBridgeCrossIn) (event.Subscription, error) {

	logs, sub, err := _NBridge.contract.WatchLogs(opts, "CrossIn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NBridgeCrossIn)
				if err := _NBridge.contract.UnpackLog(event, "CrossIn", log); err != nil {
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

// ParseCrossIn is a log parse operation binding the contract event 0xe69376bccafdb809f685ce8a90249211489ab0ec624c2faf6c91a22f8dddad5f.
//
// Solidity: event CrossIn(address originToken, uint256 originChainId, uint256 fromChainId, uint256 toChainId, address from, address to, uint256 amount)
func (_NBridge *NBridgeFilterer) ParseCrossIn(log types.Log) (*NBridgeCrossIn, error) {
	event := new(NBridgeCrossIn)
	if err := _NBridge.contract.UnpackLog(event, "CrossIn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NBridgeCrossOutIterator is returned from FilterCrossOut and is used to iterate over the raw logs and unpacked data for CrossOut events raised by the NBridge contract.
type NBridgeCrossOutIterator struct {
	Event *NBridgeCrossOut // Event containing the contract specifics and raw log

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
func (it *NBridgeCrossOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NBridgeCrossOut)
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
		it.Event = new(NBridgeCrossOut)
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
func (it *NBridgeCrossOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NBridgeCrossOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NBridgeCrossOut represents a CrossOut event raised by the NBridge contract.
type NBridgeCrossOut struct {
	OriginToken   common.Address
	OriginChainId *big.Int
	FromChainId   *big.Int
	ToChainId     *big.Int
	From          common.Address
	To            common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCrossOut is a free log retrieval operation binding the contract event 0x549f0ee31ecfdd93c4d6c1591746373d18f77138d790e2424e759e4f3b88798a.
//
// Solidity: event CrossOut(address originToken, uint256 originChainId, uint256 fromChainId, uint256 toChainId, address from, address to, uint256 amount)
func (_NBridge *NBridgeFilterer) FilterCrossOut(opts *bind.FilterOpts) (*NBridgeCrossOutIterator, error) {

	logs, sub, err := _NBridge.contract.FilterLogs(opts, "CrossOut")
	if err != nil {
		return nil, err
	}
	return &NBridgeCrossOutIterator{contract: _NBridge.contract, event: "CrossOut", logs: logs, sub: sub}, nil
}

// WatchCrossOut is a free log subscription operation binding the contract event 0x549f0ee31ecfdd93c4d6c1591746373d18f77138d790e2424e759e4f3b88798a.
//
// Solidity: event CrossOut(address originToken, uint256 originChainId, uint256 fromChainId, uint256 toChainId, address from, address to, uint256 amount)
func (_NBridge *NBridgeFilterer) WatchCrossOut(opts *bind.WatchOpts, sink chan<- *NBridgeCrossOut) (event.Subscription, error) {

	logs, sub, err := _NBridge.contract.WatchLogs(opts, "CrossOut")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NBridgeCrossOut)
				if err := _NBridge.contract.UnpackLog(event, "CrossOut", log); err != nil {
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

// ParseCrossOut is a log parse operation binding the contract event 0x549f0ee31ecfdd93c4d6c1591746373d18f77138d790e2424e759e4f3b88798a.
//
// Solidity: event CrossOut(address originToken, uint256 originChainId, uint256 fromChainId, uint256 toChainId, address from, address to, uint256 amount)
func (_NBridge *NBridgeFilterer) ParseCrossOut(log types.Log) (*NBridgeCrossOut, error) {
	event := new(NBridgeCrossOut)
	if err := _NBridge.contract.UnpackLog(event, "CrossOut", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
