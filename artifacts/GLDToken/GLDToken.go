// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GLDToken

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

// GLDToeknMetaData contains all meta data concerning the GLDToekn contract.
var GLDToeknMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialSupply\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000b8838038062000b88833981016040819052620000349162000174565b6040518060400160405280600481526020016311dbdb1960e21b8152506040518060400160405280600381526020016211d31160ea1b81525081600390816200007e919062000232565b5060046200008d828262000232565b505050620000a23382620000a960201b60201c565b5062000326565b6001600160a01b038216620001045760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604482015260640160405180910390fd5b8060026000828254620001189190620002fe565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b505050565b6000602082840312156200018757600080fd5b5051919050565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680620001b957607f821691505b602082108103620001da57634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200016f57600081815260208120601f850160051c81016020861015620002095750805b601f850160051c820191505b818110156200022a5782815560010162000215565b505050505050565b81516001600160401b038111156200024e576200024e6200018e565b62000266816200025f8454620001a4565b84620001e0565b602080601f8311600181146200029e5760008415620002855750858301515b600019600386901b1c1916600185901b1785556200022a565b600085815260208120601f198616915b82811015620002cf57888601518255948401946001909101908401620002ae565b5085821015620002ee5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b808201808211156200032057634e487b7160e01b600052601160045260246000fd5b92915050565b61085280620003366000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012357806370a082311461013657806395d89b411461015f578063a457c2d714610167578063a9059cbb1461017a578063dd62ed3e1461018d57600080fd5b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ef57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b66101a0565b6040516100c3919061069c565b60405180910390f35b6100df6100da366004610706565b610232565b60405190151581526020016100c3565b6002545b6040519081526020016100c3565b6100df61010f366004610730565b61024c565b604051601281526020016100c3565b6100df610131366004610706565b610270565b6100f361014436600461076c565b6001600160a01b031660009081526020819052604090205490565b6100b6610292565b6100df610175366004610706565b6102a1565b6100df610188366004610706565b610321565b6100f361019b36600461078e565b61032f565b6060600380546101af906107c1565b80601f01602080910402602001604051908101604052809291908181526020018280546101db906107c1565b80156102285780601f106101fd57610100808354040283529160200191610228565b820191906000526020600020905b81548152906001019060200180831161020b57829003601f168201915b5050505050905090565b60003361024081858561035a565b60019150505b92915050565b60003361025a85828561047e565b6102658585856104f8565b506001949350505050565b600033610240818585610283838361032f565b61028d91906107fb565b61035a565b6060600480546101af906107c1565b600033816102af828661032f565b9050838110156103145760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b610265828686840361035a565b6000336102408185856104f8565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166103bc5760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b606482015260840161030b565b6001600160a01b03821661041d5760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b606482015260840161030b565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b600061048a848461032f565b905060001981146104f257818110156104e55760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000604482015260640161030b565b6104f2848484840361035a565b50505050565b6001600160a01b03831661055c5760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b606482015260840161030b565b6001600160a01b0382166105be5760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b606482015260840161030b565b6001600160a01b038316600090815260208190526040902054818110156106365760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b606482015260840161030b565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a36104f2565b600060208083528351808285015260005b818110156106c9578581018301518582016040015282016106ad565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b038116811461070157600080fd5b919050565b6000806040838503121561071957600080fd5b610722836106ea565b946020939093013593505050565b60008060006060848603121561074557600080fd5b61074e846106ea565b925061075c602085016106ea565b9150604084013590509250925092565b60006020828403121561077e57600080fd5b610787826106ea565b9392505050565b600080604083850312156107a157600080fd5b6107aa836106ea565b91506107b8602084016106ea565b90509250929050565b600181811c908216806107d557607f821691505b6020821081036107f557634e487b7160e01b600052602260045260246000fd5b50919050565b8082018082111561024657634e487b7160e01b600052601160045260246000fdfea2646970667358221220846f41ac4d1a8a2b8dac9c913f24952ba3d5f010fe5a9d861a1a65f67836193964736f6c63430008130033",
}

// GLDToeknABI is the input ABI used to generate the binding from.
// Deprecated: Use GLDToeknMetaData.ABI instead.
var GLDToeknABI = GLDToeknMetaData.ABI

// GLDToeknBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GLDToeknMetaData.Bin instead.
var GLDToeknBin = GLDToeknMetaData.Bin

// DeployGLDToekn deploys a new Ethereum contract, binding an instance of GLDToekn to it.
func DeployGLDToekn(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int) (common.Address, *types.Transaction, *GLDToekn, error) {
	parsed, err := GLDToeknMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GLDToeknBin), backend, initialSupply)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GLDToekn{GLDToeknCaller: GLDToeknCaller{contract: contract}, GLDToeknTransactor: GLDToeknTransactor{contract: contract}, GLDToeknFilterer: GLDToeknFilterer{contract: contract}}, nil
}

// GLDToekn is an auto generated Go binding around an Ethereum contract.
type GLDToekn struct {
	GLDToeknCaller     // Read-only binding to the contract
	GLDToeknTransactor // Write-only binding to the contract
	GLDToeknFilterer   // Log filterer for contract events
}

// GLDToeknCaller is an auto generated read-only Go binding around an Ethereum contract.
type GLDToeknCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GLDToeknTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GLDToeknTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GLDToeknFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GLDToeknFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GLDToeknSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GLDToeknSession struct {
	Contract     *GLDToekn         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GLDToeknCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GLDToeknCallerSession struct {
	Contract *GLDToeknCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// GLDToeknTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GLDToeknTransactorSession struct {
	Contract     *GLDToeknTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// GLDToeknRaw is an auto generated low-level Go binding around an Ethereum contract.
type GLDToeknRaw struct {
	Contract *GLDToekn // Generic contract binding to access the raw methods on
}

// GLDToeknCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GLDToeknCallerRaw struct {
	Contract *GLDToeknCaller // Generic read-only contract binding to access the raw methods on
}

// GLDToeknTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GLDToeknTransactorRaw struct {
	Contract *GLDToeknTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGLDToekn creates a new instance of GLDToekn, bound to a specific deployed contract.
func NewGLDToekn(address common.Address, backend bind.ContractBackend) (*GLDToekn, error) {
	contract, err := bindGLDToekn(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GLDToekn{GLDToeknCaller: GLDToeknCaller{contract: contract}, GLDToeknTransactor: GLDToeknTransactor{contract: contract}, GLDToeknFilterer: GLDToeknFilterer{contract: contract}}, nil
}

// NewGLDToeknCaller creates a new read-only instance of GLDToekn, bound to a specific deployed contract.
func NewGLDToeknCaller(address common.Address, caller bind.ContractCaller) (*GLDToeknCaller, error) {
	contract, err := bindGLDToekn(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GLDToeknCaller{contract: contract}, nil
}

// NewGLDToeknTransactor creates a new write-only instance of GLDToekn, bound to a specific deployed contract.
func NewGLDToeknTransactor(address common.Address, transactor bind.ContractTransactor) (*GLDToeknTransactor, error) {
	contract, err := bindGLDToekn(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GLDToeknTransactor{contract: contract}, nil
}

// NewGLDToeknFilterer creates a new log filterer instance of GLDToekn, bound to a specific deployed contract.
func NewGLDToeknFilterer(address common.Address, filterer bind.ContractFilterer) (*GLDToeknFilterer, error) {
	contract, err := bindGLDToekn(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GLDToeknFilterer{contract: contract}, nil
}

// bindGLDToekn binds a generic wrapper to an already deployed contract.
func bindGLDToekn(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GLDToeknMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GLDToekn *GLDToeknRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GLDToekn.Contract.GLDToeknCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GLDToekn *GLDToeknRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GLDToekn.Contract.GLDToeknTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GLDToekn *GLDToeknRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GLDToekn.Contract.GLDToeknTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GLDToekn *GLDToeknCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GLDToekn.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GLDToekn *GLDToeknTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GLDToekn.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GLDToekn *GLDToeknTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GLDToekn.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_GLDToekn *GLDToeknCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GLDToekn.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_GLDToekn *GLDToeknSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _GLDToekn.Contract.Allowance(&_GLDToekn.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_GLDToekn *GLDToeknCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _GLDToekn.Contract.Allowance(&_GLDToekn.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GLDToekn *GLDToeknCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GLDToekn.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GLDToekn *GLDToeknSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _GLDToekn.Contract.BalanceOf(&_GLDToekn.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_GLDToekn *GLDToeknCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _GLDToekn.Contract.BalanceOf(&_GLDToekn.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GLDToekn *GLDToeknCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _GLDToekn.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GLDToekn *GLDToeknSession) Decimals() (uint8, error) {
	return _GLDToekn.Contract.Decimals(&_GLDToekn.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_GLDToekn *GLDToeknCallerSession) Decimals() (uint8, error) {
	return _GLDToekn.Contract.Decimals(&_GLDToekn.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GLDToekn *GLDToeknCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GLDToekn.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GLDToekn *GLDToeknSession) Name() (string, error) {
	return _GLDToekn.Contract.Name(&_GLDToekn.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GLDToekn *GLDToeknCallerSession) Name() (string, error) {
	return _GLDToekn.Contract.Name(&_GLDToekn.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GLDToekn *GLDToeknCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GLDToekn.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GLDToekn *GLDToeknSession) Symbol() (string, error) {
	return _GLDToekn.Contract.Symbol(&_GLDToekn.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_GLDToekn *GLDToeknCallerSession) Symbol() (string, error) {
	return _GLDToekn.Contract.Symbol(&_GLDToekn.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GLDToekn *GLDToeknCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GLDToekn.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GLDToekn *GLDToeknSession) TotalSupply() (*big.Int, error) {
	return _GLDToekn.Contract.TotalSupply(&_GLDToekn.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_GLDToekn *GLDToeknCallerSession) TotalSupply() (*big.Int, error) {
	return _GLDToekn.Contract.TotalSupply(&_GLDToekn.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_GLDToekn *GLDToeknTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GLDToekn.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_GLDToekn *GLDToeknSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GLDToekn.Contract.Approve(&_GLDToekn.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_GLDToekn *GLDToeknTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GLDToekn.Contract.Approve(&_GLDToekn.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_GLDToekn *GLDToeknTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _GLDToekn.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_GLDToekn *GLDToeknSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _GLDToekn.Contract.DecreaseAllowance(&_GLDToekn.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_GLDToekn *GLDToeknTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _GLDToekn.Contract.DecreaseAllowance(&_GLDToekn.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_GLDToekn *GLDToeknTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _GLDToekn.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_GLDToekn *GLDToeknSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _GLDToekn.Contract.IncreaseAllowance(&_GLDToekn.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_GLDToekn *GLDToeknTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _GLDToekn.Contract.IncreaseAllowance(&_GLDToekn.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_GLDToekn *GLDToeknTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GLDToekn.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_GLDToekn *GLDToeknSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GLDToekn.Contract.Transfer(&_GLDToekn.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_GLDToekn *GLDToeknTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GLDToekn.Contract.Transfer(&_GLDToekn.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_GLDToekn *GLDToeknTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GLDToekn.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_GLDToekn *GLDToeknSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GLDToekn.Contract.TransferFrom(&_GLDToekn.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_GLDToekn *GLDToeknTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _GLDToekn.Contract.TransferFrom(&_GLDToekn.TransactOpts, from, to, amount)
}

// GLDToeknApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the GLDToekn contract.
type GLDToeknApprovalIterator struct {
	Event *GLDToeknApproval // Event containing the contract specifics and raw log

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
func (it *GLDToeknApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GLDToeknApproval)
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
		it.Event = new(GLDToeknApproval)
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
func (it *GLDToeknApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GLDToeknApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GLDToeknApproval represents a Approval event raised by the GLDToekn contract.
type GLDToeknApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GLDToekn *GLDToeknFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GLDToeknApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GLDToekn.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GLDToeknApprovalIterator{contract: _GLDToekn.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_GLDToekn *GLDToeknFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GLDToeknApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _GLDToekn.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GLDToeknApproval)
				if err := _GLDToekn.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_GLDToekn *GLDToeknFilterer) ParseApproval(log types.Log) (*GLDToeknApproval, error) {
	event := new(GLDToeknApproval)
	if err := _GLDToekn.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GLDToeknTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the GLDToekn contract.
type GLDToeknTransferIterator struct {
	Event *GLDToeknTransfer // Event containing the contract specifics and raw log

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
func (it *GLDToeknTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GLDToeknTransfer)
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
		it.Event = new(GLDToeknTransfer)
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
func (it *GLDToeknTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GLDToeknTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GLDToeknTransfer represents a Transfer event raised by the GLDToekn contract.
type GLDToeknTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GLDToekn *GLDToeknFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GLDToeknTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GLDToekn.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GLDToeknTransferIterator{contract: _GLDToekn.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_GLDToekn *GLDToeknFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GLDToeknTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _GLDToekn.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GLDToeknTransfer)
				if err := _GLDToekn.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_GLDToekn *GLDToeknFilterer) ParseTransfer(log types.Log) (*GLDToeknTransfer, error) {
	event := new(GLDToeknTransfer)
	if err := _GLDToekn.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
