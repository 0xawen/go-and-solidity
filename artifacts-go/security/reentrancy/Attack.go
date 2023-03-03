// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package reentrancy

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

// AttackMetaData contains all meta data concerning the Attack contract.
var AttackMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractBank\",\"name\":\"_bank\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"attack\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bank\",\"outputs\":[{\"internalType\":\"contractBank\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161030538038061030583398101604081905261002f91610054565b600080546001600160a01b0319166001600160a01b0392909216919091179055610084565b60006020828403121561006657600080fd5b81516001600160a01b038116811461007d57600080fd5b9392505050565b610272806100936000396000f3fe6080604052600436106100385760003560e01c806312065fe0146100bb57806376cdb03b146100db5780639e5faafc1461011357600080fd5b366100b657600054670de0b6b3a76400006001600160a01b0390911631106100b4576000805460408051633ccfd60b60e01b815290516001600160a01b0390921692633ccfd60b9260048084019382900301818387803b15801561009b57600080fd5b505af11580156100af573d6000803e3d6000fd5b505050505b005b600080fd5b3480156100c757600080fd5b506040514781526020015b60405180910390f35b3480156100e757600080fd5b506000546100fb906001600160a01b031681565b6040516001600160a01b0390911681526020016100d2565b6100b434670de0b6b3a7640000146101715760405162461bcd60e51b815260206004820152601960248201527f52657175697265203120457468657220746f2061747461636b00000000000000604482015260640160405180910390fd5b60008054906101000a90046001600160a01b03166001600160a01b031663d0e30db0670de0b6b3a76400006040518263ffffffff1660e01b81526004016000604051808303818588803b1580156101c757600080fd5b505af11580156101db573d6000803e3d6000fd5b50506000805460408051633ccfd60b60e01b815290516001600160a01b039092169550633ccfd60b9450600480820194509082900301818387803b15801561022257600080fd5b505af1158015610236573d6000803e3d6000fd5b5050505056fea2646970667358221220c2c2163f1749b1aa813217db358a5dc1720b79e436c8b9c897f9639f76c5154e64736f6c63430008130033",
}

// AttackABI is the input ABI used to generate the binding from.
// Deprecated: Use AttackMetaData.ABI instead.
var AttackABI = AttackMetaData.ABI

// AttackBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AttackMetaData.Bin instead.
var AttackBin = AttackMetaData.Bin

// DeployAttack deploys a new Ethereum contract, binding an instance of Attack to it.
func DeployAttack(auth *bind.TransactOpts, backend bind.ContractBackend, _bank common.Address) (common.Address, *types.Transaction, *Attack, error) {
	parsed, err := AttackMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AttackBin), backend, _bank)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Attack{AttackCaller: AttackCaller{contract: contract}, AttackTransactor: AttackTransactor{contract: contract}, AttackFilterer: AttackFilterer{contract: contract}}, nil
}

// Attack is an auto generated Go binding around an Ethereum contract.
type Attack struct {
	AttackCaller     // Read-only binding to the contract
	AttackTransactor // Write-only binding to the contract
	AttackFilterer   // Log filterer for contract events
}

// AttackCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttackCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttackTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttackTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttackFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttackFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttackSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttackSession struct {
	Contract     *Attack           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttackCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttackCallerSession struct {
	Contract *AttackCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AttackTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttackTransactorSession struct {
	Contract     *AttackTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttackRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttackRaw struct {
	Contract *Attack // Generic contract binding to access the raw methods on
}

// AttackCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttackCallerRaw struct {
	Contract *AttackCaller // Generic read-only contract binding to access the raw methods on
}

// AttackTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttackTransactorRaw struct {
	Contract *AttackTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttack creates a new instance of Attack, bound to a specific deployed contract.
func NewAttack(address common.Address, backend bind.ContractBackend) (*Attack, error) {
	contract, err := bindAttack(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Attack{AttackCaller: AttackCaller{contract: contract}, AttackTransactor: AttackTransactor{contract: contract}, AttackFilterer: AttackFilterer{contract: contract}}, nil
}

// NewAttackCaller creates a new read-only instance of Attack, bound to a specific deployed contract.
func NewAttackCaller(address common.Address, caller bind.ContractCaller) (*AttackCaller, error) {
	contract, err := bindAttack(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttackCaller{contract: contract}, nil
}

// NewAttackTransactor creates a new write-only instance of Attack, bound to a specific deployed contract.
func NewAttackTransactor(address common.Address, transactor bind.ContractTransactor) (*AttackTransactor, error) {
	contract, err := bindAttack(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttackTransactor{contract: contract}, nil
}

// NewAttackFilterer creates a new log filterer instance of Attack, bound to a specific deployed contract.
func NewAttackFilterer(address common.Address, filterer bind.ContractFilterer) (*AttackFilterer, error) {
	contract, err := bindAttack(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttackFilterer{contract: contract}, nil
}

// bindAttack binds a generic wrapper to an already deployed contract.
func bindAttack(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AttackMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attack *AttackRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attack.Contract.AttackCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attack *AttackRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attack.Contract.AttackTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attack *AttackRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attack.Contract.AttackTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attack *AttackCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attack.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attack *AttackTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attack.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attack *AttackTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attack.Contract.contract.Transact(opts, method, params...)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Attack *AttackCaller) Bank(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Attack.contract.Call(opts, &out, "bank")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Attack *AttackSession) Bank() (common.Address, error) {
	return _Attack.Contract.Bank(&_Attack.CallOpts)
}

// Bank is a free data retrieval call binding the contract method 0x76cdb03b.
//
// Solidity: function bank() view returns(address)
func (_Attack *AttackCallerSession) Bank() (common.Address, error) {
	return _Attack.Contract.Bank(&_Attack.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Attack *AttackCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Attack.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Attack *AttackSession) GetBalance() (*big.Int, error) {
	return _Attack.Contract.GetBalance(&_Attack.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Attack *AttackCallerSession) GetBalance() (*big.Int, error) {
	return _Attack.Contract.GetBalance(&_Attack.CallOpts)
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() payable returns()
func (_Attack *AttackTransactor) Attack(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attack.contract.Transact(opts, "attack")
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() payable returns()
func (_Attack *AttackSession) Attack() (*types.Transaction, error) {
	return _Attack.Contract.Attack(&_Attack.TransactOpts)
}

// Attack is a paid mutator transaction binding the contract method 0x9e5faafc.
//
// Solidity: function attack() payable returns()
func (_Attack *AttackTransactorSession) Attack() (*types.Transaction, error) {
	return _Attack.Contract.Attack(&_Attack.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Attack *AttackTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attack.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Attack *AttackSession) Receive() (*types.Transaction, error) {
	return _Attack.Contract.Receive(&_Attack.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Attack *AttackTransactorSession) Receive() (*types.Transaction, error) {
	return _Attack.Contract.Receive(&_Attack.TransactOpts)
}
