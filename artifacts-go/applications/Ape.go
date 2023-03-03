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

// ApeMetaData contains all meta data concerning the Ape contract.
var ApeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalFroAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_APES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526127106006553480156200001757600080fd5b50604051620014bc380380620014bc8339810160408190526200003a9162000129565b818160006200004a838262000222565b50600162000059828262000222565b5050505050620002ee565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200008c57600080fd5b81516001600160401b0380821115620000a957620000a962000064565b604051601f8301601f19908116603f01168101908282118183101715620000d457620000d462000064565b81604052838152602092508683858801011115620000f157600080fd5b600091505b83821015620001155785820183015181830184015290820190620000f6565b600093810190920192909252949350505050565b600080604083850312156200013d57600080fd5b82516001600160401b03808211156200015557600080fd5b62000163868387016200007a565b935060208501519150808211156200017a57600080fd5b5062000189858286016200007a565b9150509250929050565b600181811c90821680620001a857607f821691505b602082108103620001c957634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200021d57600081815260208120601f850160051c81016020861015620001f85750805b601f850160051c820191505b81811015620002195782815560010162000204565b5050505b505050565b81516001600160401b038111156200023e576200023e62000064565b62000256816200024f845462000193565b84620001cf565b602080601f8311600181146200028e5760008415620002755750858301515b600019600386901b1c1916600185901b17855562000219565b600085815260208120601f198616915b82811015620002bf578886015182559484019460019091019084016200029e565b5085821015620002de5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6111be80620002fe6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80636352211e11610097578063b88d4fde11610066578063b88d4fde146101ff578063bb8a16bd14610212578063c87b56dd1461021b578063e985e9c51461022e57600080fd5b80636352211e146101b057806370a08231146101c357806395d89b41146101e4578063a22cb465146101ec57600080fd5b8063095ea7b3116100d3578063095ea7b31461016257806323b872dd1461017757806340c10f191461018a57806342842e0e1461019d57600080fd5b806301ffc9a7146100fa57806306fdde0314610122578063081812fc14610137575b600080fd5b61010d610108366004610d54565b61026a565b60405190151581526020015b60405180910390f35b61012a6102bc565b6040516101199190610dc1565b61014a610145366004610dd4565b61034a565b6040516001600160a01b039091168152602001610119565b610175610170366004610e04565b6103c5565b005b610175610185366004610e2e565b61046a565b610175610198366004610e04565b6104d9565b6101756101ab366004610e2e565b61052f565b61014a6101be366004610dd4565b61054a565b6101d66101d1366004610e6a565b6105aa565b604051908152602001610119565b61012a610615565b6101756101fa366004610e85565b610622565b61017561020d366004610ed7565b61068e565b6101d660065481565b61012a610229366004610dd4565b6106ff565b61010d61023c366004610fb3565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b60006001600160e01b031982166380ac58cd60e01b148061029b57506001600160e01b031982166301ffc9a760e01b145b806102b657506001600160e01b03198216635b5e139f60e01b145b92915050565b600080546102c990610fe6565b80601f01602080910402602001604051908101604052809291908181526020018280546102f590610fe6565b80156103425780601f1061031757610100808354040283529160200191610342565b820191906000526020600020905b81548152906001019060200180831161032557829003601f168201915b505050505081565b6000818152600260205260408120546001600160a01b03166103a95760405162461bcd60e51b81526020600482015260136024820152721d1bdad95b88191bd95cdb89dd08195e1a5cdd606a1b60448201526064015b60405180910390fd5b506000908152600460205260409020546001600160a01b031690565b6000818152600260205260409020546001600160a01b03163381148061040e57506001600160a01b038116600090815260056020908152604080832033845290915290205460ff165b61045a5760405162461bcd60e51b815260206004820152601e60248201527f6e6f74206f776e6572206e6f7220617070726f76656420666f7220616c6c000060448201526064016103a0565b6104658184846107b4565b505050565b60006104758261054a565b9050610482813384610810565b6104c75760405162461bcd60e51b81526020600482015260166024820152751b9bdd081bdddb995c881b9bdc88185c1c1c9bdd995960521b60448201526064016103a0565b6104d381858585610883565b50505050565b60065481106105215760405162461bcd60e51b8152602060048201526014602482015273746f6b656e4964206f7574206f662072616e676560601b60448201526064016103a0565b61052b82826109eb565b5050565b6104658383836040518060200160405280600081525061068e565b6000818152600260205260409020546001600160a01b0316806105a55760405162461bcd60e51b81526020600482015260136024820152721d1bdad95b88191bd95cdb89dd08195e1a5cdd606a1b60448201526064016103a0565b919050565b60006001600160a01b0382166105f95760405162461bcd60e51b81526020600482015260146024820152736f776e6572203d207a65726f206164647265737360601b60448201526064016103a0565b506001600160a01b031660009081526003602052604090205490565b600180546102c990610fe6565b3360008181526005602090815260408083206001600160a01b03871680855290835292819020805460ff191686151590811790915590519081529192917fc7b99ff4ca33e4f27553208175f2098c0a9bff3bdf645a7549ff1d98b69ed843910160405180910390a35050565b60006106998361054a565b90506106a6813385610810565b6106eb5760405162461bcd60e51b81526020600482015260166024820152751b9bdd081bdddb995c881b9bdc88185c1c1c9bdd995960521b60448201526064016103a0565b6106f88186868686610b1b565b5050505050565b6000818152600260205260409020546060906001600160a01b03166107585760405162461bcd60e51b815260206004820152600f60248201526e151bdad95b88139bdd08115e1a5cdd608a1b60448201526064016103a0565b6000610762610b74565b9050600081511161078257604051806020016040528060008152506107ad565b8061078c84610b94565b60405160200161079d929190611020565b6040516020818303038152906040525b9392505050565b60008181526004602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92591a4505050565b6000836001600160a01b0316836001600160a01b0316148061084b57506000828152600460205260409020546001600160a01b038481169116145b8061087b57506001600160a01b0380851660009081526005602090815260408083209387168352929052205460ff165b949350505050565b836001600160a01b0316836001600160a01b0316146108d05760405162461bcd60e51b81526020600482015260096024820152683737ba1037bbb732b960b91b60448201526064016103a0565b6001600160a01b0382166109265760405162461bcd60e51b815260206004820152601c60248201527f7472616e7366657220746f20746865207a65726f20616464726573730000000060448201526064016103a0565b610932846000836107b4565b6001600160a01b038316600090815260036020526040812080546001929061095b908490611065565b90915550506001600160a01b0382166000908152600360205260408120805460019290610989908490611078565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a450505050565b6001600160a01b038216610a385760405162461bcd60e51b81526020600482015260146024820152736d696e7420746f207a65726f206164647265737360601b60448201526064016103a0565b6000818152600260205260409020546001600160a01b031615610a945760405162461bcd60e51b81526020600482015260146024820152731d1bdad95b88185b1c9958591e481b5a5b9d195960621b60448201526064016103a0565b6001600160a01b0382166000908152600360205260408120805460019290610abd908490611078565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b610b2785858585610883565b610b3384848484610c95565b6106f85760405162461bcd60e51b81526020600482015260126024820152713737ba1022a9219b9918a932b1b2b4bb32b960711b60448201526064016103a0565b606060405180606001604052806036815260200161115360369139905090565b606081600003610bbb5750506040805180820190915260018152600360fc1b602082015290565b8160005b8115610be55780610bcf8161108b565b9150610bde9050600a836110ba565b9150610bbf565b60008167ffffffffffffffff811115610c0057610c00610ec1565b6040519080825280601f01601f191660200182016040528015610c2a576020820181803683370190505b5090505b841561087b57610c3f600183611065565b9150610c4c600a866110ce565b610c57906030611078565b60f81b818381518110610c6c57610c6c6110e2565b60200101906001600160f81b031916908160001a905350610c8e600a866110ba565b9450610c2e565b60006001600160a01b0384163b15610d3057604051630a85bd0160e11b808252906001600160a01b0386169063150b7a0290610cdb9033908a90899089906004016110f8565b6020604051808303816000875af1158015610cfa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d1e9190611135565b6001600160e01b03191614905061087b565b506001949350505050565b6001600160e01b031981168114610d5157600080fd5b50565b600060208284031215610d6657600080fd5b81356107ad81610d3b565b60005b83811015610d8c578181015183820152602001610d74565b50506000910152565b60008151808452610dad816020860160208601610d71565b601f01601f19169290920160200192915050565b6020815260006107ad6020830184610d95565b600060208284031215610de657600080fd5b5035919050565b80356001600160a01b03811681146105a557600080fd5b60008060408385031215610e1757600080fd5b610e2083610ded565b946020939093013593505050565b600080600060608486031215610e4357600080fd5b610e4c84610ded565b9250610e5a60208501610ded565b9150604084013590509250925092565b600060208284031215610e7c57600080fd5b6107ad82610ded565b60008060408385031215610e9857600080fd5b610ea183610ded565b915060208301358015158114610eb657600080fd5b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b60008060008060808587031215610eed57600080fd5b610ef685610ded565b9350610f0460208601610ded565b925060408501359150606085013567ffffffffffffffff80821115610f2857600080fd5b818701915087601f830112610f3c57600080fd5b813581811115610f4e57610f4e610ec1565b604051601f8201601f19908116603f01168101908382118183101715610f7657610f76610ec1565b816040528281528a6020848701011115610f8f57600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b60008060408385031215610fc657600080fd5b610fcf83610ded565b9150610fdd60208401610ded565b90509250929050565b600181811c90821680610ffa57607f821691505b60208210810361101a57634e487b7160e01b600052602260045260246000fd5b50919050565b60008351611032818460208801610d71565b835190830190611046818360208801610d71565b01949350505050565b634e487b7160e01b600052601160045260246000fd5b818103818111156102b6576102b661104f565b808201808211156102b6576102b661104f565b60006001820161109d5761109d61104f565b5060010190565b634e487b7160e01b600052601260045260246000fd5b6000826110c9576110c96110a4565b500490565b6000826110dd576110dd6110a4565b500690565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061112b90830184610d95565b9695505050505050565b60006020828403121561114757600080fd5b81516107ad81610d3b56fe697066733a2f2f516d65536a53696e4870506e6d586d73704d6a776958794e367a533445397a63636172694752336a7863615774712fa2646970667358221220075cd5d8e2dc63593a44dc2d9b9013ad853102a4af6786507c577775183cbddb64736f6c63430008130033",
}

// ApeABI is the input ABI used to generate the binding from.
// Deprecated: Use ApeMetaData.ABI instead.
var ApeABI = ApeMetaData.ABI

// ApeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApeMetaData.Bin instead.
var ApeBin = ApeMetaData.Bin

// DeployApe deploys a new Ethereum contract, binding an instance of Ape to it.
func DeployApe(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *Ape, error) {
	parsed, err := ApeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApeBin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ape{ApeCaller: ApeCaller{contract: contract}, ApeTransactor: ApeTransactor{contract: contract}, ApeFilterer: ApeFilterer{contract: contract}}, nil
}

// Ape is an auto generated Go binding around an Ethereum contract.
type Ape struct {
	ApeCaller     // Read-only binding to the contract
	ApeTransactor // Write-only binding to the contract
	ApeFilterer   // Log filterer for contract events
}

// ApeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApeSession struct {
	Contract     *Ape              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApeCallerSession struct {
	Contract *ApeCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApeTransactorSession struct {
	Contract     *ApeTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApeRaw struct {
	Contract *Ape // Generic contract binding to access the raw methods on
}

// ApeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApeCallerRaw struct {
	Contract *ApeCaller // Generic read-only contract binding to access the raw methods on
}

// ApeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApeTransactorRaw struct {
	Contract *ApeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApe creates a new instance of Ape, bound to a specific deployed contract.
func NewApe(address common.Address, backend bind.ContractBackend) (*Ape, error) {
	contract, err := bindApe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ape{ApeCaller: ApeCaller{contract: contract}, ApeTransactor: ApeTransactor{contract: contract}, ApeFilterer: ApeFilterer{contract: contract}}, nil
}

// NewApeCaller creates a new read-only instance of Ape, bound to a specific deployed contract.
func NewApeCaller(address common.Address, caller bind.ContractCaller) (*ApeCaller, error) {
	contract, err := bindApe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApeCaller{contract: contract}, nil
}

// NewApeTransactor creates a new write-only instance of Ape, bound to a specific deployed contract.
func NewApeTransactor(address common.Address, transactor bind.ContractTransactor) (*ApeTransactor, error) {
	contract, err := bindApe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApeTransactor{contract: contract}, nil
}

// NewApeFilterer creates a new log filterer instance of Ape, bound to a specific deployed contract.
func NewApeFilterer(address common.Address, filterer bind.ContractFilterer) (*ApeFilterer, error) {
	contract, err := bindApe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApeFilterer{contract: contract}, nil
}

// bindApe binds a generic wrapper to an already deployed contract.
func bindApe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ape *ApeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ape.Contract.ApeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ape *ApeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ape.Contract.ApeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ape *ApeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ape.Contract.ApeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ape *ApeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ape.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ape *ApeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ape.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ape *ApeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ape.Contract.contract.Transact(opts, method, params...)
}

// MAXAPES is a free data retrieval call binding the contract method 0xbb8a16bd.
//
// Solidity: function MAX_APES() view returns(uint256)
func (_Ape *ApeCaller) MAXAPES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ape.contract.Call(opts, &out, "MAX_APES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXAPES is a free data retrieval call binding the contract method 0xbb8a16bd.
//
// Solidity: function MAX_APES() view returns(uint256)
func (_Ape *ApeSession) MAXAPES() (*big.Int, error) {
	return _Ape.Contract.MAXAPES(&_Ape.CallOpts)
}

// MAXAPES is a free data retrieval call binding the contract method 0xbb8a16bd.
//
// Solidity: function MAX_APES() view returns(uint256)
func (_Ape *ApeCallerSession) MAXAPES() (*big.Int, error) {
	return _Ape.Contract.MAXAPES(&_Ape.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ape *ApeCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Ape.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ape *ApeSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Ape.Contract.BalanceOf(&_Ape.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Ape *ApeCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Ape.Contract.BalanceOf(&_Ape.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ape *ApeCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ape.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ape *ApeSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Ape.Contract.GetApproved(&_Ape.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Ape *ApeCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Ape.Contract.GetApproved(&_Ape.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ape *ApeCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Ape.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ape *ApeSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Ape.Contract.IsApprovedForAll(&_Ape.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Ape *ApeCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Ape.Contract.IsApprovedForAll(&_Ape.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ape *ApeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ape.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ape *ApeSession) Name() (string, error) {
	return _Ape.Contract.Name(&_Ape.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Ape *ApeCallerSession) Name() (string, error) {
	return _Ape.Contract.Name(&_Ape.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Ape *ApeCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ape.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Ape *ApeSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Ape.Contract.OwnerOf(&_Ape.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address owner)
func (_Ape *ApeCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Ape.Contract.OwnerOf(&_Ape.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Ape *ApeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Ape.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Ape *ApeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ape.Contract.SupportsInterface(&_Ape.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Ape *ApeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ape.Contract.SupportsInterface(&_Ape.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ape *ApeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Ape.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ape *ApeSession) Symbol() (string, error) {
	return _Ape.Contract.Symbol(&_Ape.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Ape *ApeCallerSession) Symbol() (string, error) {
	return _Ape.Contract.Symbol(&_Ape.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Ape *ApeCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Ape.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Ape *ApeSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Ape.Contract.TokenURI(&_Ape.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Ape *ApeCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Ape.Contract.TokenURI(&_Ape.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ape *ApeTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ape.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ape *ApeSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ape.Contract.Approve(&_Ape.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Ape *ApeTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ape.Contract.Approve(&_Ape.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_Ape *ApeTransactor) Mint(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ape.contract.Transact(opts, "mint", to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_Ape *ApeSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ape.Contract.Mint(&_Ape.TransactOpts, to, tokenId)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 tokenId) returns()
func (_Ape *ApeTransactorSession) Mint(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ape.Contract.Mint(&_Ape.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ape *ApeTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ape.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ape *ApeSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ape.Contract.SafeTransferFrom(&_Ape.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Ape *ApeTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ape.Contract.SafeTransferFrom(&_Ape.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Ape *ApeTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ape.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Ape *ApeSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ape.Contract.SafeTransferFrom0(&_Ape.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Ape *ApeTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Ape.Contract.SafeTransferFrom0(&_Ape.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ape *ApeTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ape.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ape *ApeSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ape.Contract.SetApprovalForAll(&_Ape.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Ape *ApeTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Ape.Contract.SetApprovalForAll(&_Ape.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ape *ApeTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ape.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ape *ApeSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ape.Contract.TransferFrom(&_Ape.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Ape *ApeTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Ape.Contract.TransferFrom(&_Ape.TransactOpts, from, to, tokenId)
}

// ApeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Ape contract.
type ApeApprovalIterator struct {
	Event *ApeApproval // Event containing the contract specifics and raw log

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
func (it *ApeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApeApproval)
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
		it.Event = new(ApeApproval)
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
func (it *ApeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApeApproval represents a Approval event raised by the Ape contract.
type ApeApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ape *ApeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ApeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Ape.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ApeApprovalIterator{contract: _Ape.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ape *ApeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ApeApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Ape.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApeApproval)
				if err := _Ape.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Ape *ApeFilterer) ParseApproval(log types.Log) (*ApeApproval, error) {
	event := new(ApeApproval)
	if err := _Ape.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApeApprovalFroAllIterator is returned from FilterApprovalFroAll and is used to iterate over the raw logs and unpacked data for ApprovalFroAll events raised by the Ape contract.
type ApeApprovalFroAllIterator struct {
	Event *ApeApprovalFroAll // Event containing the contract specifics and raw log

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
func (it *ApeApprovalFroAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApeApprovalFroAll)
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
		it.Event = new(ApeApprovalFroAll)
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
func (it *ApeApprovalFroAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApeApprovalFroAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApeApprovalFroAll represents a ApprovalFroAll event raised by the Ape contract.
type ApeApprovalFroAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalFroAll is a free log retrieval operation binding the contract event 0xc7b99ff4ca33e4f27553208175f2098c0a9bff3bdf645a7549ff1d98b69ed843.
//
// Solidity: event ApprovalFroAll(address indexed owner, address indexed operator, bool approved)
func (_Ape *ApeFilterer) FilterApprovalFroAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ApeApprovalFroAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Ape.contract.FilterLogs(opts, "ApprovalFroAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ApeApprovalFroAllIterator{contract: _Ape.contract, event: "ApprovalFroAll", logs: logs, sub: sub}, nil
}

// WatchApprovalFroAll is a free log subscription operation binding the contract event 0xc7b99ff4ca33e4f27553208175f2098c0a9bff3bdf645a7549ff1d98b69ed843.
//
// Solidity: event ApprovalFroAll(address indexed owner, address indexed operator, bool approved)
func (_Ape *ApeFilterer) WatchApprovalFroAll(opts *bind.WatchOpts, sink chan<- *ApeApprovalFroAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Ape.contract.WatchLogs(opts, "ApprovalFroAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApeApprovalFroAll)
				if err := _Ape.contract.UnpackLog(event, "ApprovalFroAll", log); err != nil {
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

// ParseApprovalFroAll is a log parse operation binding the contract event 0xc7b99ff4ca33e4f27553208175f2098c0a9bff3bdf645a7549ff1d98b69ed843.
//
// Solidity: event ApprovalFroAll(address indexed owner, address indexed operator, bool approved)
func (_Ape *ApeFilterer) ParseApprovalFroAll(log types.Log) (*ApeApprovalFroAll, error) {
	event := new(ApeApprovalFroAll)
	if err := _Ape.contract.UnpackLog(event, "ApprovalFroAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Ape contract.
type ApeTransferIterator struct {
	Event *ApeTransfer // Event containing the contract specifics and raw log

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
func (it *ApeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApeTransfer)
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
		it.Event = new(ApeTransfer)
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
func (it *ApeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApeTransfer represents a Transfer event raised by the Ape contract.
type ApeTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ape *ApeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ApeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Ape.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ApeTransferIterator{contract: _Ape.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ape *ApeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ApeTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Ape.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApeTransfer)
				if err := _Ape.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Ape *ApeFilterer) ParseTransfer(log types.Log) (*ApeTransfer, error) {
	event := new(ApeTransfer)
	if err := _Ape.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
