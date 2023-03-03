// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

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

// ERC20MintableMetaData contains all meta data concerning the ERC20Mintable contract.
var ERC20MintableMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5060405162000e5f38038062000e5f833981016040819052620000349162000193565b818160036200004483826200028c565b5060046200005382826200028c565b505050620000706200006a6200007860201b60201c565b6200007c565b505062000358565b3390565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620000f657600080fd5b81516001600160401b0380821115620001135762000113620000ce565b604051601f8301601f19908116603f011681019082821181831017156200013e576200013e620000ce565b816040528381526020925086838588010111156200015b57600080fd5b600091505b838210156200017f578582018301518183018401529082019062000160565b600093810190920192909252949350505050565b60008060408385031215620001a757600080fd5b82516001600160401b0380821115620001bf57600080fd5b620001cd86838701620000e4565b93506020850151915080821115620001e457600080fd5b50620001f385828601620000e4565b9150509250929050565b600181811c908216806200021257607f821691505b6020821081036200023357634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200028757600081815260208120601f850160051c81016020861015620002625750805b601f850160051c820191505b8181101562000283578281556001016200026e565b5050505b505050565b81516001600160401b03811115620002a857620002a8620000ce565b620002c081620002b98454620001fd565b8462000239565b602080601f831160018114620002f85760008415620002df5750858301515b600019600386901b1c1916600185901b17855562000283565b600085815260208120601f198616915b82811015620003295788860151825594840194600190910190840162000308565b5085821015620003485787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b610af780620003686000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c806370a0823111610097578063a457c2d711610066578063a457c2d7146101eb578063a9059cbb146101fe578063dd62ed3e14610211578063f2fde38b1461022457600080fd5b806370a0823114610197578063715018a6146101c05780638da5cb5b146101c857806395d89b41146101e357600080fd5b806323b872dd116100d357806323b872dd1461014d578063313ce56714610160578063395093511461016f57806340c10f191461018257600080fd5b806306fdde03146100fa578063095ea7b31461011857806318160ddd1461013b575b600080fd5b610102610237565b60405161010f9190610941565b60405180910390f35b61012b6101263660046109ab565b6102c9565b604051901515815260200161010f565b6002545b60405190815260200161010f565b61012b61015b3660046109d5565b6102e3565b6040516012815260200161010f565b61012b61017d3660046109ab565b610307565b6101956101903660046109ab565b610329565b005b61013f6101a5366004610a11565b6001600160a01b031660009081526020819052604090205490565b61019561033f565b6005546040516001600160a01b03909116815260200161010f565b610102610353565b61012b6101f93660046109ab565b610362565b61012b61020c3660046109ab565b6103e2565b61013f61021f366004610a33565b6103f0565b610195610232366004610a11565b61041b565b60606003805461024690610a66565b80601f016020809104026020016040519081016040528092919081815260200182805461027290610a66565b80156102bf5780601f10610294576101008083540402835291602001916102bf565b820191906000526020600020905b8154815290600101906020018083116102a257829003601f168201915b5050505050905090565b6000336102d7818585610494565b60019150505b92915050565b6000336102f18582856105b8565b6102fc858585610632565b506001949350505050565b6000336102d781858561031a83836103f0565b6103249190610aa0565b610494565b6103316107d6565b61033b8282610830565b5050565b6103476107d6565b61035160006108ef565b565b60606004805461024690610a66565b6000338161037082866103f0565b9050838110156103d55760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b6102fc8286868403610494565b6000336102d7818585610632565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b6104236107d6565b6001600160a01b0381166104885760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016103cc565b610491816108ef565b50565b6001600160a01b0383166104f65760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b60648201526084016103cc565b6001600160a01b0382166105575760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b60648201526084016103cc565b6001600160a01b0383811660008181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b60006105c484846103f0565b9050600019811461062c578181101561061f5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e636500000060448201526064016103cc565b61062c8484848403610494565b50505050565b6001600160a01b0383166106965760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084016103cc565b6001600160a01b0382166106f85760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b60648201526084016103cc565b6001600160a01b038316600090815260208190526040902054818110156107705760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b60648201526084016103cc565b6001600160a01b03848116600081815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a361062c565b6005546001600160a01b031633146103515760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016103cc565b6001600160a01b0382166108865760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064016103cc565b80600260008282546108989190610aa0565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b600580546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600060208083528351808285015260005b8181101561096e57858101830151858201604001528201610952565b506000604082860101526040601f19601f8301168501019250505092915050565b80356001600160a01b03811681146109a657600080fd5b919050565b600080604083850312156109be57600080fd5b6109c78361098f565b946020939093013593505050565b6000806000606084860312156109ea57600080fd5b6109f38461098f565b9250610a016020850161098f565b9150604084013590509250925092565b600060208284031215610a2357600080fd5b610a2c8261098f565b9392505050565b60008060408385031215610a4657600080fd5b610a4f8361098f565b9150610a5d6020840161098f565b90509250929050565b600181811c90821680610a7a57607f821691505b602082108103610a9a57634e487b7160e01b600052602260045260246000fd5b50919050565b808201808211156102dd57634e487b7160e01b600052601160045260246000fdfea2646970667358221220278a517a2b18daa502ce7369605b5e0517303625f7cc67e609055cddafc28ca764736f6c63430008130033",
}

// ERC20MintableABI is the input ABI used to generate the binding from.
// Deprecated: Use ERC20MintableMetaData.ABI instead.
var ERC20MintableABI = ERC20MintableMetaData.ABI

// ERC20MintableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERC20MintableMetaData.Bin instead.
var ERC20MintableBin = ERC20MintableMetaData.Bin

// DeployERC20Mintable deploys a new Ethereum contract, binding an instance of ERC20Mintable to it.
func DeployERC20Mintable(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *ERC20Mintable, error) {
	parsed, err := ERC20MintableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERC20MintableBin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Mintable{ERC20MintableCaller: ERC20MintableCaller{contract: contract}, ERC20MintableTransactor: ERC20MintableTransactor{contract: contract}, ERC20MintableFilterer: ERC20MintableFilterer{contract: contract}}, nil
}

// ERC20Mintable is an auto generated Go binding around an Ethereum contract.
type ERC20Mintable struct {
	ERC20MintableCaller     // Read-only binding to the contract
	ERC20MintableTransactor // Write-only binding to the contract
	ERC20MintableFilterer   // Log filterer for contract events
}

// ERC20MintableCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20MintableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20MintableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20MintableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20MintableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20MintableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20MintableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20MintableSession struct {
	Contract     *ERC20Mintable    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20MintableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20MintableCallerSession struct {
	Contract *ERC20MintableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ERC20MintableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20MintableTransactorSession struct {
	Contract     *ERC20MintableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ERC20MintableRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20MintableRaw struct {
	Contract *ERC20Mintable // Generic contract binding to access the raw methods on
}

// ERC20MintableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20MintableCallerRaw struct {
	Contract *ERC20MintableCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20MintableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20MintableTransactorRaw struct {
	Contract *ERC20MintableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Mintable creates a new instance of ERC20Mintable, bound to a specific deployed contract.
func NewERC20Mintable(address common.Address, backend bind.ContractBackend) (*ERC20Mintable, error) {
	contract, err := bindERC20Mintable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Mintable{ERC20MintableCaller: ERC20MintableCaller{contract: contract}, ERC20MintableTransactor: ERC20MintableTransactor{contract: contract}, ERC20MintableFilterer: ERC20MintableFilterer{contract: contract}}, nil
}

// NewERC20MintableCaller creates a new read-only instance of ERC20Mintable, bound to a specific deployed contract.
func NewERC20MintableCaller(address common.Address, caller bind.ContractCaller) (*ERC20MintableCaller, error) {
	contract, err := bindERC20Mintable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20MintableCaller{contract: contract}, nil
}

// NewERC20MintableTransactor creates a new write-only instance of ERC20Mintable, bound to a specific deployed contract.
func NewERC20MintableTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20MintableTransactor, error) {
	contract, err := bindERC20Mintable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20MintableTransactor{contract: contract}, nil
}

// NewERC20MintableFilterer creates a new log filterer instance of ERC20Mintable, bound to a specific deployed contract.
func NewERC20MintableFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20MintableFilterer, error) {
	contract, err := bindERC20Mintable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20MintableFilterer{contract: contract}, nil
}

// bindERC20Mintable binds a generic wrapper to an already deployed contract.
func bindERC20Mintable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ERC20MintableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Mintable *ERC20MintableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Mintable.Contract.ERC20MintableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Mintable *ERC20MintableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.ERC20MintableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Mintable *ERC20MintableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.ERC20MintableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Mintable *ERC20MintableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20Mintable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Mintable *ERC20MintableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Mintable *ERC20MintableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Mintable *ERC20MintableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Mintable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Mintable *ERC20MintableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Mintable.Contract.Allowance(&_ERC20Mintable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20Mintable *ERC20MintableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20Mintable.Contract.Allowance(&_ERC20Mintable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Mintable *ERC20MintableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Mintable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Mintable *ERC20MintableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Mintable.Contract.BalanceOf(&_ERC20Mintable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20Mintable *ERC20MintableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20Mintable.Contract.BalanceOf(&_ERC20Mintable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Mintable *ERC20MintableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20Mintable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Mintable *ERC20MintableSession) Decimals() (uint8, error) {
	return _ERC20Mintable.Contract.Decimals(&_ERC20Mintable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20Mintable *ERC20MintableCallerSession) Decimals() (uint8, error) {
	return _ERC20Mintable.Contract.Decimals(&_ERC20Mintable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Mintable *ERC20MintableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Mintable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Mintable *ERC20MintableSession) Name() (string, error) {
	return _ERC20Mintable.Contract.Name(&_ERC20Mintable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20Mintable *ERC20MintableCallerSession) Name() (string, error) {
	return _ERC20Mintable.Contract.Name(&_ERC20Mintable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Mintable *ERC20MintableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERC20Mintable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Mintable *ERC20MintableSession) Owner() (common.Address, error) {
	return _ERC20Mintable.Contract.Owner(&_ERC20Mintable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ERC20Mintable *ERC20MintableCallerSession) Owner() (common.Address, error) {
	return _ERC20Mintable.Contract.Owner(&_ERC20Mintable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Mintable *ERC20MintableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20Mintable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Mintable *ERC20MintableSession) Symbol() (string, error) {
	return _ERC20Mintable.Contract.Symbol(&_ERC20Mintable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20Mintable *ERC20MintableCallerSession) Symbol() (string, error) {
	return _ERC20Mintable.Contract.Symbol(&_ERC20Mintable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Mintable *ERC20MintableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20Mintable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Mintable *ERC20MintableSession) TotalSupply() (*big.Int, error) {
	return _ERC20Mintable.Contract.TotalSupply(&_ERC20Mintable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Mintable *ERC20MintableCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Mintable.Contract.TotalSupply(&_ERC20Mintable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.Approve(&_ERC20Mintable.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.Approve(&_ERC20Mintable.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Mintable *ERC20MintableSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.DecreaseAllowance(&_ERC20Mintable.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.DecreaseAllowance(&_ERC20Mintable.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Mintable *ERC20MintableSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.IncreaseAllowance(&_ERC20Mintable.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.IncreaseAllowance(&_ERC20Mintable.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ERC20Mintable *ERC20MintableTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ERC20Mintable *ERC20MintableSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.Mint(&_ERC20Mintable.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ERC20Mintable *ERC20MintableTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.Mint(&_ERC20Mintable.TransactOpts, account, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Mintable *ERC20MintableTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Mintable *ERC20MintableSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20Mintable.Contract.RenounceOwnership(&_ERC20Mintable.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ERC20Mintable *ERC20MintableTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ERC20Mintable.Contract.RenounceOwnership(&_ERC20Mintable.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.Transfer(&_ERC20Mintable.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.Transfer(&_ERC20Mintable.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.TransferFrom(&_ERC20Mintable.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ERC20Mintable *ERC20MintableTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.TransferFrom(&_ERC20Mintable.TransactOpts, from, to, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Mintable *ERC20MintableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Mintable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Mintable *ERC20MintableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.TransferOwnership(&_ERC20Mintable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ERC20Mintable *ERC20MintableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ERC20Mintable.Contract.TransferOwnership(&_ERC20Mintable.TransactOpts, newOwner)
}

// ERC20MintableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20Mintable contract.
type ERC20MintableApprovalIterator struct {
	Event *ERC20MintableApproval // Event containing the contract specifics and raw log

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
func (it *ERC20MintableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MintableApproval)
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
		it.Event = new(ERC20MintableApproval)
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
func (it *ERC20MintableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MintableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MintableApproval represents a Approval event raised by the ERC20Mintable contract.
type ERC20MintableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Mintable *ERC20MintableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20MintableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Mintable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MintableApprovalIterator{contract: _ERC20Mintable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20Mintable *ERC20MintableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20MintableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20Mintable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MintableApproval)
				if err := _ERC20Mintable.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_ERC20Mintable *ERC20MintableFilterer) ParseApproval(log types.Log) (*ERC20MintableApproval, error) {
	event := new(ERC20MintableApproval)
	if err := _ERC20Mintable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20MintableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ERC20Mintable contract.
type ERC20MintableOwnershipTransferredIterator struct {
	Event *ERC20MintableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ERC20MintableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MintableOwnershipTransferred)
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
		it.Event = new(ERC20MintableOwnershipTransferred)
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
func (it *ERC20MintableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MintableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MintableOwnershipTransferred represents a OwnershipTransferred event raised by the ERC20Mintable contract.
type ERC20MintableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20Mintable *ERC20MintableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ERC20MintableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20Mintable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MintableOwnershipTransferredIterator{contract: _ERC20Mintable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ERC20Mintable *ERC20MintableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ERC20MintableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ERC20Mintable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MintableOwnershipTransferred)
				if err := _ERC20Mintable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ERC20Mintable *ERC20MintableFilterer) ParseOwnershipTransferred(log types.Log) (*ERC20MintableOwnershipTransferred, error) {
	event := new(ERC20MintableOwnershipTransferred)
	if err := _ERC20Mintable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC20MintableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Mintable contract.
type ERC20MintableTransferIterator struct {
	Event *ERC20MintableTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20MintableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20MintableTransfer)
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
		it.Event = new(ERC20MintableTransfer)
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
func (it *ERC20MintableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20MintableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20MintableTransfer represents a Transfer event raised by the ERC20Mintable contract.
type ERC20MintableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Mintable *ERC20MintableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20MintableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Mintable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20MintableTransferIterator{contract: _ERC20Mintable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Mintable *ERC20MintableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20MintableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Mintable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20MintableTransfer)
				if err := _ERC20Mintable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_ERC20Mintable *ERC20MintableFilterer) ParseTransfer(log types.Log) (*ERC20MintableTransfer, error) {
	event := new(ERC20MintableTransfer)
	if err := _ERC20Mintable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
