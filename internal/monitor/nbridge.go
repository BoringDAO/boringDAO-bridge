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

// NCrossInParams is an auto generated low-level Go binding around an user-defined struct.
type NCrossInParams struct {
	OriginToken   common.Address
	OriginChainId *big.Int
	FromChainId   *big.Int
	ToChainId     *big.Int
	From          common.Address
	To            common.Address
	Amount        *big.Int
	Txid          string
}

// NBridgeABI is the input ABI used to generate the binding from.
const NBridgeABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CrossIn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"originToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"originChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CrossOut\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_chainIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mirrorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mirrorChainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"internalType\":\"structNBridge.TokenInfo[]\",\"name\":\"tis\",\"type\":\"tuple[]\"}],\"name\":\"addMultiSupportTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mirrorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mirrorChainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"internalType\":\"structNBridge.TokenInfo\",\"name\":\"ti\",\"type\":\"tuple\"}],\"name\":\"addSupportToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_originToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_originChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fromChainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toChainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"txid\",\"type\":\"string\"}],\"internalType\":\"structNCrossInParams\",\"name\":\"p\",\"type\":\"tuple\"}],\"name\":\"crossIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"toChainID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"crossOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"}],\"name\":\"getRoleKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"minCrossAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_feeRatio\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"setFeeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setMinCrossAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_threshold\",\"type\":\"uint256\"}],\"name\":\"setThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supportedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"mirrorAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mirrorChainId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSupported\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"txHandled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// NBridgeFuncSigs maps the 4-byte function signature to its string representation.
var NBridgeFuncSigs = map[string]string{
	"11dc39d3": "addMultiSupportTokens(uint256[],address[],(uint256,address,uint256,bool)[])",
	"ebf3c976": "addSupportToken(uint256,address,(uint256,address,uint256,bool))",
	"9a8a0592": "chainId()",
	"4871a760": "crossIn((address,uint256,uint256,uint256,address,address,uint256,string))",
	"df7e600a": "crossOut(address,uint256,address,uint256)",
	"8881a1aa": "getRoleKey(address,uint256)",
	"c3470cb0": "minCrossAmount(address,uint256)",
	"e55156b5": "setFee(address,uint256)",
	"f46901ed": "setFeeTo(address)",
	"9393ffde": "setMinCrossAmount(address,uint256,uint256)",
	"9d879990": "setThreshold(address,uint256)",
	"15431c48": "supportedTokens(uint256,address)",
	"824df449": "txHandled(string)",
}

// NBridgeBin is the compiled bytecode used for deploying new contracts.
var NBridgeBin = "0x608060405234801561001057600080fd5b50604051610ee1380380610ee183398101604081905261002f91610037565b600055610050565b60006020828403121561004957600080fd5b5051919050565b610e828061005f6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80639a8a05921161008c578063df7e600a11610066578063df7e600a1461024f578063e55156b514610212578063ebf3c97614610262578063f46901ed1461027557600080fd5b80639a8a0592146102095780639d87999014610212578063c3470cb01461022457600080fd5b806311dc39d3146100d457806315431c48146100e95780634871a76014610163578063824df449146101765780638881a1aa146101b45780639393ffde146101d5575b600080fd5b6100e76100e2366004610aa8565b610286565b005b6101346100f7366004610c92565b6001602081815260009384526040808520909152918352912080549181015460028201546003909201546001600160a01b03909116919060ff1684565b604080519485526001600160a01b03909316602085015291830152151560608201526080015b60405180910390f35b6100e7610171366004610bc9565b61034a565b6101a4610184366004610b8c565b805160208183018101805160028252928201919093012091525460ff1681565b604051901515815260200161015a565b6101c76101c2366004610a07565b6104af565b60405190815260200161015a565b6100e76101e3366004610a75565b6001600160a01b0390921660009081526003602090815260408083209383529290522055565b6101c760005481565b6100e7610220366004610a07565b5050565b6101c7610232366004610a07565b600360209081526000928352604080842090915290825290205481565b6100e761025d366004610a31565b6104f3565b6100e7610270366004610cbe565b6106ff565b6100e76102833660046109e5565b50565b8151815181146102d05760405162461bcd60e51b815260206004820152601060248201526f0d8cadccee8d040dcdee840dac2e8c6d60831b60448201526064015b60405180910390fd5b60005b81811015610343576103318582815181106102f0576102f0610e20565b602002602001015185838151811061030a5761030a610e20565b602002602001015185848151811061032457610324610e20565b60200260200101516106ff565b8061033b81610df7565b9150506102d3565b5050505050565b80516020820151600061035d83836104af565b90508360e001516002816040516103749190610cfb565b9081526040519081900360200190205460ff16156103c15760405162461bcd60e51b815260206004820152600a6024820152691d1e081a185b991b195960b21b60448201526064016102c7565b6000548560600151146104065760405162461bcd60e51b815260206004820152600d60248201526c31b430b4b724b21032b93937b960991b60448201526064016102c7565b60208086015160009081526001808352604080832089516001600160a01b0390811685529085529281902081516080810183528154815292810154909316938201939093526002820154928101929092526003015460ff161515606082018190526104a75760405162461bcd60e51b81526020600482015260116024820152703737ba1039bab83837b93a103a37b5b2b760791b60448201526064016102c7565b505050505050565b6040805160609390931b6bffffffffffffffffffffffff19166020808501919091526034808501939093528151808503909301835260549093019052805191012090565b600081116105345760405162461bcd60e51b815260206004820152600e60248201526d063726f737320616d6f756e7420360941b60448201526064016102c7565b6000805481526001602081815260408084206001600160a01b03808a16865290835293819020815160808101835281548152938101549094169183019190915260028301549082015260039091015460ff161515606082018190526105cf5760405162461bcd60e51b81526020600482015260116024820152703737ba1039bab83837b93a103a37b5b2b760791b60448201526064016102c7565b6020808201516001600160a01b031660009081526003825260408082208783529092522054801561064a578083101561064a5760405162461bcd60e51b815260206004820152601a60248201527f616d6f756e74206c657373207468616e206d696e416d6f756e7400000000000060448201526064016102c7565b8151600114156106a0577f549f0ee31ecfdd93c4d6c1591746373d18f77138d790e2424e759e4f3b88798a86600054600054883389896040516106939796959493929190610d36565b60405180910390a16104a7565b8151600214156104a7577f549f0ee31ecfdd93c4d6c1591746373d18f77138d790e2424e759e4f3b88798a82602001518360400151600054883389896040516106ef9796959493929190610d36565b60405180910390a1505050505050565b6001600160a01b03821661074b5760405162461bcd60e51b81526020600482015260136024820152726572726f7220746f6b656e206164647265737360681b60448201526064016102c7565b606081015115156001146107905760405162461bcd60e51b815260206004820152600c60248201526b3830b930b6b99032b93937b960a11b60448201526064016102c7565b60009283526001602081815260408086206001600160a01b039586168752825294859020835181559083015191810180546001600160a01b03191692909416919091179092559182015160028201556060909101516003909101805460ff1916911515919091179055565b80356001600160a01b038116811461081257600080fd5b919050565b600082601f83011261082857600080fd5b8135602061083d61083883610dd3565b610da2565b80838252828201915082860187848660051b890101111561085d57600080fd5b60005b8581101561088357610871826107fb565b84529284019290840190600101610860565b5090979650505050505050565b600082601f8301126108a157600080fd5b813560206108b161083883610dd3565b80838252828201915082860187848660071b89010111156108d157600080fd5b60005b85811015610883576108e6898361096b565b845292840192608091909101906001016108d4565b600082601f83011261090c57600080fd5b813567ffffffffffffffff81111561092657610926610e36565b610939601f8201601f1916602001610da2565b81815284602083860101111561094e57600080fd5b816020850160208301376000918101602001919091529392505050565b60006080828403121561097d57600080fd5b6040516080810181811067ffffffffffffffff821117156109a0576109a0610e36565b604052823581529050806109b6602084016107fb565b602082015260408301356040820152606083013580151581146109d857600080fd5b6060919091015292915050565b6000602082840312156109f757600080fd5b610a00826107fb565b9392505050565b60008060408385031215610a1a57600080fd5b610a23836107fb565b946020939093013593505050565b60008060008060808587031215610a4757600080fd5b610a50856107fb565b935060208501359250610a65604086016107fb565b9396929550929360600135925050565b600080600060608486031215610a8a57600080fd5b610a93846107fb565b95602085013595506040909401359392505050565b600080600060608486031215610abd57600080fd5b833567ffffffffffffffff80821115610ad557600080fd5b818601915086601f830112610ae957600080fd5b81356020610af961083883610dd3565b8083825282820191508286018b848660051b8901011115610b1957600080fd5b600096505b84871015610b3c578035835260019690960195918301918301610b1e565b5097505087013592505080821115610b5357600080fd5b610b5f87838801610817565b93506040860135915080821115610b7557600080fd5b50610b8286828701610890565b9150509250925092565b600060208284031215610b9e57600080fd5b813567ffffffffffffffff811115610bb557600080fd5b610bc1848285016108fb565b949350505050565b600060208284031215610bdb57600080fd5b813567ffffffffffffffff80821115610bf357600080fd5b908301906101008286031215610c0857600080fd5b610c10610d78565b610c19836107fb565b8152602083013560208201526040830135604082015260608301356060820152610c45608084016107fb565b6080820152610c5660a084016107fb565b60a082015260c083013560c082015260e083013582811115610c7757600080fd5b610c83878286016108fb565b60e08301525095945050505050565b60008060408385031215610ca557600080fd5b82359150610cb5602084016107fb565b90509250929050565b600080600060c08486031215610cd357600080fd5b83359250610ce3602085016107fb565b9150610cf2856040860161096b565b90509250925092565b6000825160005b81811015610d1c5760208186018101518583015201610d02565b81811115610d2b576000828501525b509190910192915050565b6001600160a01b0397881681526020810196909652604086019490945260608501929092528416608084015290921660a082015260c081019190915260e00190565b604051610100810167ffffffffffffffff81118282101715610d9c57610d9c610e36565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610dcb57610dcb610e36565b604052919050565b600067ffffffffffffffff821115610ded57610ded610e36565b5060051b60200190565b6000600019821415610e1957634e487b7160e01b600052601160045260246000fd5b5060010190565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220b0749a7fc489318479fb70691eaf56b8ab9d61c209eb0077ceffffb6fa5cfbe664736f6c63430008060033"

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

// GetRoleKey is a free data retrieval call binding the contract method 0x8881a1aa.
//
// Solidity: function getRoleKey(address token, uint256 _chainId) pure returns(bytes32)
func (_NBridge *NBridgeCaller) GetRoleKey(opts *bind.CallOpts, token common.Address, _chainId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _NBridge.contract.Call(opts, &out, "getRoleKey", token, _chainId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleKey is a free data retrieval call binding the contract method 0x8881a1aa.
//
// Solidity: function getRoleKey(address token, uint256 _chainId) pure returns(bytes32)
func (_NBridge *NBridgeSession) GetRoleKey(token common.Address, _chainId *big.Int) ([32]byte, error) {
	return _NBridge.Contract.GetRoleKey(&_NBridge.CallOpts, token, _chainId)
}

// GetRoleKey is a free data retrieval call binding the contract method 0x8881a1aa.
//
// Solidity: function getRoleKey(address token, uint256 _chainId) pure returns(bytes32)
func (_NBridge *NBridgeCallerSession) GetRoleKey(token common.Address, _chainId *big.Int) ([32]byte, error) {
	return _NBridge.Contract.GetRoleKey(&_NBridge.CallOpts, token, _chainId)
}

// MinCrossAmount is a free data retrieval call binding the contract method 0xc3470cb0.
//
// Solidity: function minCrossAmount(address , uint256 ) view returns(uint256)
func (_NBridge *NBridgeCaller) MinCrossAmount(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NBridge.contract.Call(opts, &out, "minCrossAmount", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinCrossAmount is a free data retrieval call binding the contract method 0xc3470cb0.
//
// Solidity: function minCrossAmount(address , uint256 ) view returns(uint256)
func (_NBridge *NBridgeSession) MinCrossAmount(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NBridge.Contract.MinCrossAmount(&_NBridge.CallOpts, arg0, arg1)
}

// MinCrossAmount is a free data retrieval call binding the contract method 0xc3470cb0.
//
// Solidity: function minCrossAmount(address , uint256 ) view returns(uint256)
func (_NBridge *NBridgeCallerSession) MinCrossAmount(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _NBridge.Contract.MinCrossAmount(&_NBridge.CallOpts, arg0, arg1)
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

// CrossIn is a paid mutator transaction binding the contract method 0x4871a760.
//
// Solidity: function crossIn((address,uint256,uint256,uint256,address,address,uint256,string) p) returns()
func (_NBridge *NBridgeTransactor) CrossIn(opts *bind.TransactOpts, p NCrossInParams) (*types.Transaction, error) {
	return _NBridge.contract.Transact(opts, "crossIn", p)
}

// CrossIn is a paid mutator transaction binding the contract method 0x4871a760.
//
// Solidity: function crossIn((address,uint256,uint256,uint256,address,address,uint256,string) p) returns()
func (_NBridge *NBridgeSession) CrossIn(p NCrossInParams) (*types.Transaction, error) {
	return _NBridge.Contract.CrossIn(&_NBridge.TransactOpts, p)
}

// CrossIn is a paid mutator transaction binding the contract method 0x4871a760.
//
// Solidity: function crossIn((address,uint256,uint256,uint256,address,address,uint256,string) p) returns()
func (_NBridge *NBridgeTransactorSession) CrossIn(p NCrossInParams) (*types.Transaction, error) {
	return _NBridge.Contract.CrossIn(&_NBridge.TransactOpts, p)
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

// SetFee is a paid mutator transaction binding the contract method 0xe55156b5.
//
// Solidity: function setFee(address token, uint256 _feeRatio) returns()
func (_NBridge *NBridgeTransactor) SetFee(opts *bind.TransactOpts, token common.Address, _feeRatio *big.Int) (*types.Transaction, error) {
	return _NBridge.contract.Transact(opts, "setFee", token, _feeRatio)
}

// SetFee is a paid mutator transaction binding the contract method 0xe55156b5.
//
// Solidity: function setFee(address token, uint256 _feeRatio) returns()
func (_NBridge *NBridgeSession) SetFee(token common.Address, _feeRatio *big.Int) (*types.Transaction, error) {
	return _NBridge.Contract.SetFee(&_NBridge.TransactOpts, token, _feeRatio)
}

// SetFee is a paid mutator transaction binding the contract method 0xe55156b5.
//
// Solidity: function setFee(address token, uint256 _feeRatio) returns()
func (_NBridge *NBridgeTransactorSession) SetFee(token common.Address, _feeRatio *big.Int) (*types.Transaction, error) {
	return _NBridge.Contract.SetFee(&_NBridge.TransactOpts, token, _feeRatio)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address account) returns()
func (_NBridge *NBridgeTransactor) SetFeeTo(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _NBridge.contract.Transact(opts, "setFeeTo", account)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address account) returns()
func (_NBridge *NBridgeSession) SetFeeTo(account common.Address) (*types.Transaction, error) {
	return _NBridge.Contract.SetFeeTo(&_NBridge.TransactOpts, account)
}

// SetFeeTo is a paid mutator transaction binding the contract method 0xf46901ed.
//
// Solidity: function setFeeTo(address account) returns()
func (_NBridge *NBridgeTransactorSession) SetFeeTo(account common.Address) (*types.Transaction, error) {
	return _NBridge.Contract.SetFeeTo(&_NBridge.TransactOpts, account)
}

// SetMinCrossAmount is a paid mutator transaction binding the contract method 0x9393ffde.
//
// Solidity: function setMinCrossAmount(address token, uint256 _chainId, uint256 amount) returns()
func (_NBridge *NBridgeTransactor) SetMinCrossAmount(opts *bind.TransactOpts, token common.Address, _chainId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _NBridge.contract.Transact(opts, "setMinCrossAmount", token, _chainId, amount)
}

// SetMinCrossAmount is a paid mutator transaction binding the contract method 0x9393ffde.
//
// Solidity: function setMinCrossAmount(address token, uint256 _chainId, uint256 amount) returns()
func (_NBridge *NBridgeSession) SetMinCrossAmount(token common.Address, _chainId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _NBridge.Contract.SetMinCrossAmount(&_NBridge.TransactOpts, token, _chainId, amount)
}

// SetMinCrossAmount is a paid mutator transaction binding the contract method 0x9393ffde.
//
// Solidity: function setMinCrossAmount(address token, uint256 _chainId, uint256 amount) returns()
func (_NBridge *NBridgeTransactorSession) SetMinCrossAmount(token common.Address, _chainId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _NBridge.Contract.SetMinCrossAmount(&_NBridge.TransactOpts, token, _chainId, amount)
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
