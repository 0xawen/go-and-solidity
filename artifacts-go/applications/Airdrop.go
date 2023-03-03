// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package applications

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

// AirdropMetaData contains all meta data concerning the Airdrop contract.
var AirdropMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_arr\",\"type\":\"uint256[]\"}],\"name\":\"getSum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sum\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"multiTransferETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"_addresses\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"multiTransferToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506106cc806100206000396000f3fe6080604052600436106100345760003560e01c806341ed24a21461003957806377988cf81461005b578063ccb8c1e01461006e575b600080fd5b34801561004557600080fd5b506100596100543660046104a8565b6100a0565b005b61005961006936600461052b565b6102a5565b34801561007a57600080fd5b5061008e610089366004610597565b6103f9565b60405190815260200160405180910390f35b8281146100e75760405162461bcd60e51b815260206004820152601060248201526f1b195b99dd1a081b9bdd08195c5d585b60821b60448201526064015b60405180910390fd5b8460006100f484846103f9565b604051636eb1769f60e11b815233600482015230602482015290915081906001600160a01b0384169063dd62ed3e90604401602060405180830381865afa158015610143573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061016791906105d9565b116101b45760405162461bcd60e51b815260206004820152601860248201527f4e65656420617070726f766520657263323020746f6b656e000000000000000060448201526064016100de565b60005b8581101561029b57826001600160a01b03166323b872dd338989858181106101e1576101e16105f2565b90506020020160208101906101f69190610608565b888886818110610208576102086105f2565b6040516001600160e01b031960e088901b1681526001600160a01b039586166004820152949093166024850152506020909102013560448201526064016020604051808303816000875af1158015610264573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610288919061062c565b508061029381610664565b9150506101b7565b5050505050505050565b8281146103075760405162461bcd60e51b815260206004820152602a60248201527f4c656e67746873206f662041646472657373657320616e6420416d6f756e7473604482015269081393d508115455505360b21b60648201526084016100de565b600061031383836103f9565b905080341461035c5760405162461bcd60e51b81526020600482015260156024820152742a3930b739b332b91030b6b7bab73a1032b93937b960591b60448201526064016100de565b60005b848110156103f157858582818110610379576103796105f2565b905060200201602081019061038e9190610608565b6001600160a01b03166108fc8585848181106103ac576103ac6105f2565b905060200201359081150290604051600060405180830381858888f193505050501580156103de573d6000803e3d6000fd5b50806103e981610664565b91505061035f565b505050505050565b6000805b8281101561043d57838382818110610417576104176105f2565b9050602002013582610429919061067d565b91508061043581610664565b9150506103fd565b5092915050565b6001600160a01b038116811461045957600080fd5b50565b60008083601f84011261046e57600080fd5b50813567ffffffffffffffff81111561048657600080fd5b6020830191508360208260051b85010111156104a157600080fd5b9250929050565b6000806000806000606086880312156104c057600080fd5b85356104cb81610444565b9450602086013567ffffffffffffffff808211156104e857600080fd5b6104f489838a0161045c565b9096509450604088013591508082111561050d57600080fd5b5061051a8882890161045c565b969995985093965092949392505050565b6000806000806040858703121561054157600080fd5b843567ffffffffffffffff8082111561055957600080fd5b6105658883890161045c565b9096509450602087013591508082111561057e57600080fd5b5061058b8782880161045c565b95989497509550505050565b600080602083850312156105aa57600080fd5b823567ffffffffffffffff8111156105c157600080fd5b6105cd8582860161045c565b90969095509350505050565b6000602082840312156105eb57600080fd5b5051919050565b634e487b7160e01b600052603260045260246000fd5b60006020828403121561061a57600080fd5b813561062581610444565b9392505050565b60006020828403121561063e57600080fd5b8151801515811461062557600080fd5b634e487b7160e01b600052601160045260246000fd5b6000600182016106765761067661064e565b5060010190565b808201808211156106905761069061064e565b9291505056fea26469706673582212208ae9651788718bae7b99fba1d9337bced8423065c55459476736b53b9876b79464736f6c63430008130033",
}

// AirdropABI is the input ABI used to generate the binding from.
// Deprecated: Use AirdropMetaData.ABI instead.
var AirdropABI = AirdropMetaData.ABI

// AirdropBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AirdropMetaData.Bin instead.
var AirdropBin = AirdropMetaData.Bin

// DeployAirdrop deploys a new Ethereum contract, binding an instance of Airdrop to it.
func DeployAirdrop(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Airdrop, error) {
	parsed, err := AirdropMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AirdropBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Airdrop{AirdropCaller: AirdropCaller{contract: contract}, AirdropTransactor: AirdropTransactor{contract: contract}, AirdropFilterer: AirdropFilterer{contract: contract}}, nil
}

// Airdrop is an auto generated Go binding around an Ethereum contract.
type Airdrop struct {
	AirdropCaller     // Read-only binding to the contract
	AirdropTransactor // Write-only binding to the contract
	AirdropFilterer   // Log filterer for contract events
}

// AirdropCaller is an auto generated read-only Go binding around an Ethereum contract.
type AirdropCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AirdropTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AirdropFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AirdropSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AirdropSession struct {
	Contract     *Airdrop          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AirdropCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AirdropCallerSession struct {
	Contract *AirdropCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AirdropTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AirdropTransactorSession struct {
	Contract     *AirdropTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AirdropRaw is an auto generated low-level Go binding around an Ethereum contract.
type AirdropRaw struct {
	Contract *Airdrop // Generic contract binding to access the raw methods on
}

// AirdropCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AirdropCallerRaw struct {
	Contract *AirdropCaller // Generic read-only contract binding to access the raw methods on
}

// AirdropTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AirdropTransactorRaw struct {
	Contract *AirdropTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAirdrop creates a new instance of Airdrop, bound to a specific deployed contract.
func NewAirdrop(address common.Address, backend bind.ContractBackend) (*Airdrop, error) {
	contract, err := bindAirdrop(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Airdrop{AirdropCaller: AirdropCaller{contract: contract}, AirdropTransactor: AirdropTransactor{contract: contract}, AirdropFilterer: AirdropFilterer{contract: contract}}, nil
}

// NewAirdropCaller creates a new read-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropCaller(address common.Address, caller bind.ContractCaller) (*AirdropCaller, error) {
	contract, err := bindAirdrop(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropCaller{contract: contract}, nil
}

// NewAirdropTransactor creates a new write-only instance of Airdrop, bound to a specific deployed contract.
func NewAirdropTransactor(address common.Address, transactor bind.ContractTransactor) (*AirdropTransactor, error) {
	contract, err := bindAirdrop(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AirdropTransactor{contract: contract}, nil
}

// NewAirdropFilterer creates a new log filterer instance of Airdrop, bound to a specific deployed contract.
func NewAirdropFilterer(address common.Address, filterer bind.ContractFilterer) (*AirdropFilterer, error) {
	contract, err := bindAirdrop(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AirdropFilterer{contract: contract}, nil
}

// bindAirdrop binds a generic wrapper to an already deployed contract.
func bindAirdrop(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AirdropMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.AirdropCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.AirdropTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Airdrop *AirdropCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Airdrop.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Airdrop *AirdropTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Airdrop *AirdropTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Airdrop.Contract.contract.Transact(opts, method, params...)
}

// GetSum is a free data retrieval call binding the contract method 0xccb8c1e0.
//
// Solidity: function getSum(uint256[] _arr) pure returns(uint256 sum)
func (_Airdrop *AirdropCaller) GetSum(opts *bind.CallOpts, _arr []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Airdrop.contract.Call(opts, &out, "getSum", _arr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSum is a free data retrieval call binding the contract method 0xccb8c1e0.
//
// Solidity: function getSum(uint256[] _arr) pure returns(uint256 sum)
func (_Airdrop *AirdropSession) GetSum(_arr []*big.Int) (*big.Int, error) {
	return _Airdrop.Contract.GetSum(&_Airdrop.CallOpts, _arr)
}

// GetSum is a free data retrieval call binding the contract method 0xccb8c1e0.
//
// Solidity: function getSum(uint256[] _arr) pure returns(uint256 sum)
func (_Airdrop *AirdropCallerSession) GetSum(_arr []*big.Int) (*big.Int, error) {
	return _Airdrop.Contract.GetSum(&_Airdrop.CallOpts, _arr)
}

// MultiTransferETH is a paid mutator transaction binding the contract method 0x77988cf8.
//
// Solidity: function multiTransferETH(address[] _addresses, uint256[] _amounts) payable returns()
func (_Airdrop *AirdropTransactor) MultiTransferETH(opts *bind.TransactOpts, _addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "multiTransferETH", _addresses, _amounts)
}

// MultiTransferETH is a paid mutator transaction binding the contract method 0x77988cf8.
//
// Solidity: function multiTransferETH(address[] _addresses, uint256[] _amounts) payable returns()
func (_Airdrop *AirdropSession) MultiTransferETH(_addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.MultiTransferETH(&_Airdrop.TransactOpts, _addresses, _amounts)
}

// MultiTransferETH is a paid mutator transaction binding the contract method 0x77988cf8.
//
// Solidity: function multiTransferETH(address[] _addresses, uint256[] _amounts) payable returns()
func (_Airdrop *AirdropTransactorSession) MultiTransferETH(_addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.MultiTransferETH(&_Airdrop.TransactOpts, _addresses, _amounts)
}

// MultiTransferToken is a paid mutator transaction binding the contract method 0x41ed24a2.
//
// Solidity: function multiTransferToken(address _token, address[] _addresses, uint256[] _amounts) returns()
func (_Airdrop *AirdropTransactor) MultiTransferToken(opts *bind.TransactOpts, _token common.Address, _addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Airdrop.contract.Transact(opts, "multiTransferToken", _token, _addresses, _amounts)
}

// MultiTransferToken is a paid mutator transaction binding the contract method 0x41ed24a2.
//
// Solidity: function multiTransferToken(address _token, address[] _addresses, uint256[] _amounts) returns()
func (_Airdrop *AirdropSession) MultiTransferToken(_token common.Address, _addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.MultiTransferToken(&_Airdrop.TransactOpts, _token, _addresses, _amounts)
}

// MultiTransferToken is a paid mutator transaction binding the contract method 0x41ed24a2.
//
// Solidity: function multiTransferToken(address _token, address[] _addresses, uint256[] _amounts) returns()
func (_Airdrop *AirdropTransactorSession) MultiTransferToken(_token common.Address, _addresses []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Airdrop.Contract.MultiTransferToken(&_Airdrop.TransactOpts, _token, _addresses, _amounts)
}
