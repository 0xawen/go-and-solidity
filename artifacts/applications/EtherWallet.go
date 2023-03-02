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

// EtherWalletMetaData contains all meta data concerning the EtherWallet contract.
var EtherWalletMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556101e3806100326000396000f3fe6080604052600436106100385760003560e01c806312065fe0146100445780632e1a7d4d146100645780638da5cb5b1461008657600080fd5b3661003f57005b600080fd5b34801561005057600080fd5b506040514781526020015b60405180910390f35b34801561007057600080fd5b5061008461007f366004610194565b6100be565b005b34801561009257600080fd5b506000546100a6906001600160a01b031681565b6040516001600160a01b03909116815260200161005b565b6000546001600160a01b031633146101135760405162461bcd60e51b815260206004820152601360248201527231b0b63632b91034b9903737ba1037bbb732b960691b60448201526064015b60405180910390fd5b804711156101635760405162461bcd60e51b815260206004820152601c60248201527f6d6f6e6579206e6f7420656e6f75676820746f20776974686472617700000000604482015260640161010a565b604051339082156108fc029083906000818181858888f19350505050158015610190573d6000803e3d6000fd5b5050565b6000602082840312156101a657600080fd5b503591905056fea26469706673582212204505ce3bdb2fadd17e9fdb201760e2bb647612411fa55d79244f77ed8951051264736f6c63430008130033",
}

// EtherWalletABI is the input ABI used to generate the binding from.
// Deprecated: Use EtherWalletMetaData.ABI instead.
var EtherWalletABI = EtherWalletMetaData.ABI

// EtherWalletBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EtherWalletMetaData.Bin instead.
var EtherWalletBin = EtherWalletMetaData.Bin

// DeployEtherWallet deploys a new Ethereum contract, binding an instance of EtherWallet to it.
func DeployEtherWallet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EtherWallet, error) {
	parsed, err := EtherWalletMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EtherWalletBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EtherWallet{EtherWalletCaller: EtherWalletCaller{contract: contract}, EtherWalletTransactor: EtherWalletTransactor{contract: contract}, EtherWalletFilterer: EtherWalletFilterer{contract: contract}}, nil
}

// EtherWallet is an auto generated Go binding around an Ethereum contract.
type EtherWallet struct {
	EtherWalletCaller     // Read-only binding to the contract
	EtherWalletTransactor // Write-only binding to the contract
	EtherWalletFilterer   // Log filterer for contract events
}

// EtherWalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type EtherWalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherWalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EtherWalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherWalletFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EtherWalletFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EtherWalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EtherWalletSession struct {
	Contract     *EtherWallet      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EtherWalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EtherWalletCallerSession struct {
	Contract *EtherWalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EtherWalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EtherWalletTransactorSession struct {
	Contract     *EtherWalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EtherWalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type EtherWalletRaw struct {
	Contract *EtherWallet // Generic contract binding to access the raw methods on
}

// EtherWalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EtherWalletCallerRaw struct {
	Contract *EtherWalletCaller // Generic read-only contract binding to access the raw methods on
}

// EtherWalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EtherWalletTransactorRaw struct {
	Contract *EtherWalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEtherWallet creates a new instance of EtherWallet, bound to a specific deployed contract.
func NewEtherWallet(address common.Address, backend bind.ContractBackend) (*EtherWallet, error) {
	contract, err := bindEtherWallet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EtherWallet{EtherWalletCaller: EtherWalletCaller{contract: contract}, EtherWalletTransactor: EtherWalletTransactor{contract: contract}, EtherWalletFilterer: EtherWalletFilterer{contract: contract}}, nil
}

// NewEtherWalletCaller creates a new read-only instance of EtherWallet, bound to a specific deployed contract.
func NewEtherWalletCaller(address common.Address, caller bind.ContractCaller) (*EtherWalletCaller, error) {
	contract, err := bindEtherWallet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EtherWalletCaller{contract: contract}, nil
}

// NewEtherWalletTransactor creates a new write-only instance of EtherWallet, bound to a specific deployed contract.
func NewEtherWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*EtherWalletTransactor, error) {
	contract, err := bindEtherWallet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EtherWalletTransactor{contract: contract}, nil
}

// NewEtherWalletFilterer creates a new log filterer instance of EtherWallet, bound to a specific deployed contract.
func NewEtherWalletFilterer(address common.Address, filterer bind.ContractFilterer) (*EtherWalletFilterer, error) {
	contract, err := bindEtherWallet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EtherWalletFilterer{contract: contract}, nil
}

// bindEtherWallet binds a generic wrapper to an already deployed contract.
func bindEtherWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EtherWalletMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherWallet *EtherWalletRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherWallet.Contract.EtherWalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherWallet *EtherWalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherWallet.Contract.EtherWalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherWallet *EtherWalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherWallet.Contract.EtherWalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EtherWallet *EtherWalletCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EtherWallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EtherWallet *EtherWalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherWallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EtherWallet *EtherWalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EtherWallet.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_EtherWallet *EtherWalletCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EtherWallet.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_EtherWallet *EtherWalletSession) GetBalance() (*big.Int, error) {
	return _EtherWallet.Contract.GetBalance(&_EtherWallet.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_EtherWallet *EtherWalletCallerSession) GetBalance() (*big.Int, error) {
	return _EtherWallet.Contract.GetBalance(&_EtherWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherWallet *EtherWalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EtherWallet.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherWallet *EtherWalletSession) Owner() (common.Address, error) {
	return _EtherWallet.Contract.Owner(&_EtherWallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EtherWallet *EtherWalletCallerSession) Owner() (common.Address, error) {
	return _EtherWallet.Contract.Owner(&_EtherWallet.CallOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_EtherWallet *EtherWalletTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _EtherWallet.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_EtherWallet *EtherWalletSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _EtherWallet.Contract.Withdraw(&_EtherWallet.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_EtherWallet *EtherWalletTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _EtherWallet.Contract.Withdraw(&_EtherWallet.TransactOpts, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherWallet *EtherWalletTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EtherWallet.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherWallet *EtherWalletSession) Receive() (*types.Transaction, error) {
	return _EtherWallet.Contract.Receive(&_EtherWallet.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_EtherWallet *EtherWalletTransactorSession) Receive() (*types.Transaction, error) {
	return _EtherWallet.Contract.Receive(&_EtherWallet.TransactOpts)
}
