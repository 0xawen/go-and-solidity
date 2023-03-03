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

// BAYC1155MetaData contains all meta data concerning the BAYC1155 contract.
var BAYC1155MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"mintBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604080518082018252600880825267424159433131353560c01b60208084018290528451808601909552918452908301529060006200005283826200010f565b5060016200006182826200010f565b505050620001db565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806200009557607f821691505b602082108103620000b657634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200010a57600081815260208120601f850160051c81016020861015620000e55750805b601f850160051c820191505b818110156200010657828155600101620000f1565b5050505b505050565b81516001600160401b038111156200012b576200012b6200006a565b62000143816200013c845462000080565b84620000bc565b602080601f8311600181146200017b5760008415620001625750858301515b600019600386901b1c1916600185901b17855562000106565b600085815260208120601f198616915b82811015620001ac578886015182559484019460019091019084016200018b565b5085821015620001cb5787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6119b880620001eb6000396000f3fe608060405234801561001057600080fd5b50600436106100b35760003560e01c80634e1273f4116100715780634e1273f41461015157806395d89b4114610171578063a22cb46514610179578063d81d0a151461018c578063e985e9c51461019f578063f242432a146101b257600080fd5b8062fdd58e146100b857806301ffc9a7146100de57806306fdde03146101015780630e89341c14610116578063156e29f6146101295780632eb2c2d61461013e575b600080fd5b6100cb6100c6366004610fe0565b6101c5565b6040519081526020015b60405180910390f35b6100f16100ec366004611023565b610260565b60405190151581526020016100d5565b6101096102b1565b6040516100d59190611090565b6101096101243660046110a3565b61033f565b61013c6101373660046110bc565b61039d565b005b61013c61014c36600461123b565b6103fc565b61016461015f3660046112e5565b61061b565b6040516100d591906113eb565b610109610745565b61013c6101873660046113fe565b610752565b61013c61019a36600461143a565b610828565b6100f16101ad3660046114ae565b6108ba565b61013c6101c03660046114e1565b6108e8565b60006001600160a01b0383166102355760405162461bcd60e51b815260206004820152602a60248201527f455243313135353a2061646472657373207a65726f206973206e6f742061207660448201526930b634b21037bbb732b960b11b60648201526084015b60405180910390fd5b5060008181526002602090815260408083206001600160a01b03861684529091529020545b92915050565b60006001600160e01b03198216636cdb3d1360e11b148061029157506001600160e01b031982166303a24d0760e21b145b8061025a57506001600160e01b031982166301ffc9a760e01b1492915050565b600080546102be90611546565b80601f01602080910402602001604051908101604052809291908181526020018280546102ea90611546565b80156103375780601f1061030c57610100808354040283529160200191610337565b820191906000526020600020905b81548152906001019060200180831161031a57829003601f168201915b505050505081565b6060600061034b610a78565b9050600081511161036b5760405180602001604052806000815250610396565b8061037584610a98565b604051602001610386929190611580565b6040516020818303038152906040525b9392505050565b61271082106103dc5760405162461bcd60e51b815260206004820152600b60248201526a6964206f766572666c6f7760a81b604482015260640161022c565b6103f783838360405180602001604052806000815250610ba1565b505050565b336001600160a01b038616811480610419575061041986826108ba565b61047d5760405162461bcd60e51b815260206004820152602f60248201527f455243313135353a2063616c6c6572206973206e6f7420746f6b656e206f776e60448201526e195c881b9bdc88185c1c1c9bdd9959608a1b606482015260840161022c565b825184511461049e5760405162461bcd60e51b815260040161022c906115af565b6001600160a01b0385166104c45760405162461bcd60e51b815260040161022c906115f7565b60005b84518110156105ad5760008582815181106104e4576104e461163c565b6020026020010151905060008583815181106105025761050261163c565b60209081029190910181015160008481526002835260408082206001600160a01b038e1683529093529190912054909150818110156105535760405162461bcd60e51b815260040161022c90611652565b60008381526002602090815260408083206001600160a01b038e8116855292528083208585039055908b168252812080548492906105929084906116b2565b92505081905550505050806105a6906116c5565b90506104c7565b50846001600160a01b0316866001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb87876040516105fd9291906116de565b60405180910390a4610613818787878787610c62565b505050505050565b606081518351146106805760405162461bcd60e51b815260206004820152602960248201527f455243313135353a206163636f756e747320616e6420696473206c656e677468604482015268040dad2e6dac2e8c6d60bb1b606482015260840161022c565b6000835167ffffffffffffffff81111561069c5761069c6110ef565b6040519080825280602002602001820160405280156106c5578160200160208202803683370190505b50905060005b845181101561073d576107108582815181106106e9576106e961163c565b60200260200101518583815181106107035761070361163c565b60200260200101516101c5565b8282815181106107225761072261163c565b6020908102919091010152610736816116c5565b90506106cb565b509392505050565b600180546102be90611546565b6001600160a01b03821633036107bc5760405162461bcd60e51b815260206004820152602960248201527f455243313135353a2073657474696e6720617070726f76616c20737461747573604482015268103337b91039b2b63360b91b606482015260840161022c565b3360008181526003602090815260408083206001600160a01b03871680855290835292819020805460ff191686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b60005b825181101561089e576127108382815181106108495761084961163c565b60200260200101511061088c5760405162461bcd60e51b815260206004820152600b60248201526a6964206f766572666c6f7760a81b604482015260640161022c565b80610896816116c5565b91505061082b565b506103f783838360405180602001604052806000815250610dbd565b6001600160a01b03918216600090815260036020908152604080832093909416825291909152205460ff1690565b336001600160a01b038616811480610905575061090586826108ba565b6109675760405162461bcd60e51b815260206004820152602d60248201527f455243313135353a2063616c6c206973206e6f7420746f6b656e206f776e657260448201526c081b9bdd08185c1c1c9bdd9959609a1b606482015260840161022c565b6001600160a01b03851661098d5760405162461bcd60e51b815260040161022c906115f7565b60008481526002602090815260408083206001600160a01b038a168452909152902054838110156109d05760405162461bcd60e51b815260040161022c90611652565b60008581526002602090815260408083206001600160a01b038b8116855292528083208785039055908816825281208054869290610a0f9084906116b2565b909155505060408051868152602081018690526001600160a01b03808916928a821692918616917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4610a6f828888888888610f09565b50505050505050565b606060405180606001604052806036815260200161194d60369139905090565b606081600003610abf5750506040805180820190915260018152600360fc1b602082015290565b8160005b8115610ae95780610ad3816116c5565b9150610ae29050600a83611722565b9150610ac3565b60008167ffffffffffffffff811115610b0457610b046110ef565b6040519080825280601f01601f191660200182016040528015610b2e576020820181803683370190505b5090505b8415610b9957610b43600183611736565b9150610b50600a86611749565b610b5b9060306116b2565b60f81b818381518110610b7057610b7061163c565b60200101906001600160f81b031916908160001a905350610b92600a86611722565b9450610b32565b949350505050565b6001600160a01b038416610bc75760405162461bcd60e51b815260040161022c9061175d565b60008381526002602090815260408083206001600160a01b0388168452909152812080543392859291610bfb9084906116b2565b909155505060408051858152602081018590526001600160a01b0380881692600092918516917fc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62910160405180910390a4610c5b81600087878787610f09565b5050505050565b6001600160a01b0384163b156106135760405163bc197c8160e01b81526001600160a01b0385169063bc197c8190610ca6908990899088908890889060040161179e565b6020604051808303816000875af1925050508015610ce1575060408051601f3d908101601f19168201909252610cde918101906117fc565b60015b610d8d57610ced611819565b806308c379a003610d265750610d01611835565b80610d0c5750610d28565b8060405162461bcd60e51b815260040161022c9190611090565b505b60405162461bcd60e51b815260206004820152603460248201527f455243313135353a207472616e7366657220746f206e6f6e2d455243313135356044820152732932b1b2b4bb32b91034b6b83632b6b2b73a32b960611b606482015260840161022c565b6001600160e01b0319811663bc197c8160e01b14610a6f5760405162461bcd60e51b815260040161022c906118bf565b6001600160a01b038416610de35760405162461bcd60e51b815260040161022c9061175d565b8151835114610e045760405162461bcd60e51b815260040161022c906115af565b3360005b8451811015610ea157838181518110610e2357610e2361163c565b602002602001015160026000878481518110610e4157610e4161163c565b602002602001015181526020019081526020016000206000886001600160a01b03166001600160a01b031681526020019081526020016000206000828254610e8991906116b2565b90915550819050610e99816116c5565b915050610e08565b50846001600160a01b031660006001600160a01b0316826001600160a01b03167f4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb8787604051610ef29291906116de565b60405180910390a4610c5b81600087878787610c62565b6001600160a01b0384163b156106135760405163f23a6e6160e01b81526001600160a01b0385169063f23a6e6190610f4d9089908990889088908890600401611907565b6020604051808303816000875af1925050508015610f88575060408051601f3d908101601f19168201909252610f85918101906117fc565b60015b610f9457610ced611819565b6001600160e01b0319811663f23a6e6160e01b14610a6f5760405162461bcd60e51b815260040161022c906118bf565b80356001600160a01b0381168114610fdb57600080fd5b919050565b60008060408385031215610ff357600080fd5b610ffc83610fc4565b946020939093013593505050565b6001600160e01b03198116811461102057600080fd5b50565b60006020828403121561103557600080fd5b81356103968161100a565b60005b8381101561105b578181015183820152602001611043565b50506000910152565b6000815180845261107c816020860160208601611040565b601f01601f19169290920160200192915050565b6020815260006103966020830184611064565b6000602082840312156110b557600080fd5b5035919050565b6000806000606084860312156110d157600080fd5b6110da84610fc4565b95602085013595506040909401359392505050565b634e487b7160e01b600052604160045260246000fd5b601f8201601f1916810167ffffffffffffffff8111828210171561112b5761112b6110ef565b6040525050565b600067ffffffffffffffff82111561114c5761114c6110ef565b5060051b60200190565b600082601f83011261116757600080fd5b8135602061117482611132565b6040516111818282611105565b83815260059390931b85018201928281019150868411156111a157600080fd5b8286015b848110156111bc57803583529183019183016111a5565b509695505050505050565b600082601f8301126111d857600080fd5b813567ffffffffffffffff8111156111f2576111f26110ef565b604051611209601f8301601f191660200182611105565b81815284602083860101111561121e57600080fd5b816020850160208301376000918101602001919091529392505050565b600080600080600060a0868803121561125357600080fd5b61125c86610fc4565b945061126a60208701610fc4565b9350604086013567ffffffffffffffff8082111561128757600080fd5b61129389838a01611156565b945060608801359150808211156112a957600080fd5b6112b589838a01611156565b935060808801359150808211156112cb57600080fd5b506112d8888289016111c7565b9150509295509295909350565b600080604083850312156112f857600080fd5b823567ffffffffffffffff8082111561131057600080fd5b818501915085601f83011261132457600080fd5b8135602061133182611132565b60405161133e8282611105565b83815260059390931b850182019282810191508984111561135e57600080fd5b948201945b838610156113835761137486610fc4565b82529482019490820190611363565b9650508601359250508082111561139957600080fd5b506113a685828601611156565b9150509250929050565b600081518084526020808501945080840160005b838110156113e0578151875295820195908201906001016113c4565b509495945050505050565b60208152600061039660208301846113b0565b6000806040838503121561141157600080fd5b61141a83610fc4565b91506020830135801515811461142f57600080fd5b809150509250929050565b60008060006060848603121561144f57600080fd5b61145884610fc4565b9250602084013567ffffffffffffffff8082111561147557600080fd5b61148187838801611156565b9350604086013591508082111561149757600080fd5b506114a486828701611156565b9150509250925092565b600080604083850312156114c157600080fd5b6114ca83610fc4565b91506114d860208401610fc4565b90509250929050565b600080600080600060a086880312156114f957600080fd5b61150286610fc4565b945061151060208701610fc4565b93506040860135925060608601359150608086013567ffffffffffffffff81111561153a57600080fd5b6112d8888289016111c7565b600181811c9082168061155a57607f821691505b60208210810361157a57634e487b7160e01b600052602260045260246000fd5b50919050565b60008351611592818460208801611040565b8351908301906115a6818360208801611040565b01949350505050565b60208082526028908201527f455243313135353a2069647320616e6420616d6f756e7473206c656e677468206040820152670dad2e6dac2e8c6d60c31b606082015260800190565b60208082526025908201527f455243313135353a207472616e7366657220746f20746865207a65726f206164604082015264647265737360d81b606082015260800190565b634e487b7160e01b600052603260045260246000fd5b6020808252602a908201527f455243313135353a20696e73756666696369656e742062616c616e636520666f60408201526939103a3930b739b332b960b11b606082015260800190565b634e487b7160e01b600052601160045260246000fd5b8082018082111561025a5761025a61169c565b6000600182016116d7576116d761169c565b5060010190565b6040815260006116f160408301856113b0565b828103602084015261170381856113b0565b95945050505050565b634e487b7160e01b600052601260045260246000fd5b6000826117315761173161170c565b500490565b8181038181111561025a5761025a61169c565b6000826117585761175861170c565b500690565b60208082526021908201527f455243313135353a206d696e7420746f20746865207a65726f206164647265736040820152607360f81b606082015260800190565b6001600160a01b0386811682528516602082015260a0604082018190526000906117ca908301866113b0565b82810360608401526117dc81866113b0565b905082810360808401526117f08185611064565b98975050505050505050565b60006020828403121561180e57600080fd5b81516103968161100a565b600060033d11156118325760046000803e5060005160e01c5b90565b600060443d10156118435790565b6040516003193d81016004833e81513d67ffffffffffffffff816024840111818411171561187357505050505090565b828501915081518181111561188b5750505050505090565b843d87010160208285010111156118a55750505050505090565b6118b460208286010187611105565b509095945050505050565b60208082526028908201527f455243313135353a204552433131353552656365697665722072656a656374656040820152676420746f6b656e7360c01b606082015260800190565b6001600160a01b03868116825285166020820152604081018490526060810183905260a06080820181905260009061194190830184611064565b97965050505050505056fe697066733a2f2f516d65536a53696e4870506e6d586d73704d6a776958794e367a533445397a63636172694752336a7863615774712fa26469706673582212200a159da6b3fe658fcab8468362f7cc812ad42f9a4edbaf6b6f9f807313d0416564736f6c63430008130033",
}

// BAYC1155ABI is the input ABI used to generate the binding from.
// Deprecated: Use BAYC1155MetaData.ABI instead.
var BAYC1155ABI = BAYC1155MetaData.ABI

// BAYC1155Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BAYC1155MetaData.Bin instead.
var BAYC1155Bin = BAYC1155MetaData.Bin

// DeployBAYC1155 deploys a new Ethereum contract, binding an instance of BAYC1155 to it.
func DeployBAYC1155(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BAYC1155, error) {
	parsed, err := BAYC1155MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BAYC1155Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BAYC1155{BAYC1155Caller: BAYC1155Caller{contract: contract}, BAYC1155Transactor: BAYC1155Transactor{contract: contract}, BAYC1155Filterer: BAYC1155Filterer{contract: contract}}, nil
}

// BAYC1155 is an auto generated Go binding around an Ethereum contract.
type BAYC1155 struct {
	BAYC1155Caller     // Read-only binding to the contract
	BAYC1155Transactor // Write-only binding to the contract
	BAYC1155Filterer   // Log filterer for contract events
}

// BAYC1155Caller is an auto generated read-only Go binding around an Ethereum contract.
type BAYC1155Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BAYC1155Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BAYC1155Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BAYC1155Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BAYC1155Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BAYC1155Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BAYC1155Session struct {
	Contract     *BAYC1155         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BAYC1155CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BAYC1155CallerSession struct {
	Contract *BAYC1155Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BAYC1155TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BAYC1155TransactorSession struct {
	Contract     *BAYC1155Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BAYC1155Raw is an auto generated low-level Go binding around an Ethereum contract.
type BAYC1155Raw struct {
	Contract *BAYC1155 // Generic contract binding to access the raw methods on
}

// BAYC1155CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BAYC1155CallerRaw struct {
	Contract *BAYC1155Caller // Generic read-only contract binding to access the raw methods on
}

// BAYC1155TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BAYC1155TransactorRaw struct {
	Contract *BAYC1155Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBAYC1155 creates a new instance of BAYC1155, bound to a specific deployed contract.
func NewBAYC1155(address common.Address, backend bind.ContractBackend) (*BAYC1155, error) {
	contract, err := bindBAYC1155(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BAYC1155{BAYC1155Caller: BAYC1155Caller{contract: contract}, BAYC1155Transactor: BAYC1155Transactor{contract: contract}, BAYC1155Filterer: BAYC1155Filterer{contract: contract}}, nil
}

// NewBAYC1155Caller creates a new read-only instance of BAYC1155, bound to a specific deployed contract.
func NewBAYC1155Caller(address common.Address, caller bind.ContractCaller) (*BAYC1155Caller, error) {
	contract, err := bindBAYC1155(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BAYC1155Caller{contract: contract}, nil
}

// NewBAYC1155Transactor creates a new write-only instance of BAYC1155, bound to a specific deployed contract.
func NewBAYC1155Transactor(address common.Address, transactor bind.ContractTransactor) (*BAYC1155Transactor, error) {
	contract, err := bindBAYC1155(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BAYC1155Transactor{contract: contract}, nil
}

// NewBAYC1155Filterer creates a new log filterer instance of BAYC1155, bound to a specific deployed contract.
func NewBAYC1155Filterer(address common.Address, filterer bind.ContractFilterer) (*BAYC1155Filterer, error) {
	contract, err := bindBAYC1155(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BAYC1155Filterer{contract: contract}, nil
}

// bindBAYC1155 binds a generic wrapper to an already deployed contract.
func bindBAYC1155(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BAYC1155MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BAYC1155 *BAYC1155Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BAYC1155.Contract.BAYC1155Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BAYC1155 *BAYC1155Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BAYC1155.Contract.BAYC1155Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BAYC1155 *BAYC1155Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BAYC1155.Contract.BAYC1155Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BAYC1155 *BAYC1155CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BAYC1155.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BAYC1155 *BAYC1155TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BAYC1155.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BAYC1155 *BAYC1155TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BAYC1155.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_BAYC1155 *BAYC1155Caller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BAYC1155.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_BAYC1155 *BAYC1155Session) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _BAYC1155.Contract.BalanceOf(&_BAYC1155.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_BAYC1155 *BAYC1155CallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _BAYC1155.Contract.BalanceOf(&_BAYC1155.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_BAYC1155 *BAYC1155Caller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _BAYC1155.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_BAYC1155 *BAYC1155Session) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _BAYC1155.Contract.BalanceOfBatch(&_BAYC1155.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_BAYC1155 *BAYC1155CallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _BAYC1155.Contract.BalanceOfBatch(&_BAYC1155.CallOpts, accounts, ids)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_BAYC1155 *BAYC1155Caller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _BAYC1155.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_BAYC1155 *BAYC1155Session) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _BAYC1155.Contract.IsApprovedForAll(&_BAYC1155.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_BAYC1155 *BAYC1155CallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _BAYC1155.Contract.IsApprovedForAll(&_BAYC1155.CallOpts, account, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BAYC1155 *BAYC1155Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BAYC1155.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BAYC1155 *BAYC1155Session) Name() (string, error) {
	return _BAYC1155.Contract.Name(&_BAYC1155.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BAYC1155 *BAYC1155CallerSession) Name() (string, error) {
	return _BAYC1155.Contract.Name(&_BAYC1155.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BAYC1155 *BAYC1155Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BAYC1155.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BAYC1155 *BAYC1155Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BAYC1155.Contract.SupportsInterface(&_BAYC1155.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BAYC1155 *BAYC1155CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BAYC1155.Contract.SupportsInterface(&_BAYC1155.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BAYC1155 *BAYC1155Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BAYC1155.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BAYC1155 *BAYC1155Session) Symbol() (string, error) {
	return _BAYC1155.Contract.Symbol(&_BAYC1155.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BAYC1155 *BAYC1155CallerSession) Symbol() (string, error) {
	return _BAYC1155.Contract.Symbol(&_BAYC1155.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_BAYC1155 *BAYC1155Caller) Uri(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _BAYC1155.contract.Call(opts, &out, "uri", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_BAYC1155 *BAYC1155Session) Uri(id *big.Int) (string, error) {
	return _BAYC1155.Contract.Uri(&_BAYC1155.CallOpts, id)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 id) view returns(string)
func (_BAYC1155 *BAYC1155CallerSession) Uri(id *big.Int) (string, error) {
	return _BAYC1155.Contract.Uri(&_BAYC1155.CallOpts, id)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 id, uint256 amount) returns()
func (_BAYC1155 *BAYC1155Transactor) Mint(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BAYC1155.contract.Transact(opts, "mint", to, id, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 id, uint256 amount) returns()
func (_BAYC1155 *BAYC1155Session) Mint(to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BAYC1155.Contract.Mint(&_BAYC1155.TransactOpts, to, id, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 id, uint256 amount) returns()
func (_BAYC1155 *BAYC1155TransactorSession) Mint(to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _BAYC1155.Contract.Mint(&_BAYC1155.TransactOpts, to, id, amount)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd81d0a15.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts) returns()
func (_BAYC1155 *BAYC1155Transactor) MintBatch(opts *bind.TransactOpts, to common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _BAYC1155.contract.Transact(opts, "mintBatch", to, ids, amounts)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd81d0a15.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts) returns()
func (_BAYC1155 *BAYC1155Session) MintBatch(to common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _BAYC1155.Contract.MintBatch(&_BAYC1155.TransactOpts, to, ids, amounts)
}

// MintBatch is a paid mutator transaction binding the contract method 0xd81d0a15.
//
// Solidity: function mintBatch(address to, uint256[] ids, uint256[] amounts) returns()
func (_BAYC1155 *BAYC1155TransactorSession) MintBatch(to common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _BAYC1155.Contract.MintBatch(&_BAYC1155.TransactOpts, to, ids, amounts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_BAYC1155 *BAYC1155Transactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _BAYC1155.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_BAYC1155 *BAYC1155Session) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _BAYC1155.Contract.SafeBatchTransferFrom(&_BAYC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_BAYC1155 *BAYC1155TransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _BAYC1155.Contract.SafeBatchTransferFrom(&_BAYC1155.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_BAYC1155 *BAYC1155Transactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BAYC1155.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_BAYC1155 *BAYC1155Session) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BAYC1155.Contract.SafeTransferFrom(&_BAYC1155.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_BAYC1155 *BAYC1155TransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _BAYC1155.Contract.SafeTransferFrom(&_BAYC1155.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BAYC1155 *BAYC1155Transactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _BAYC1155.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BAYC1155 *BAYC1155Session) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BAYC1155.Contract.SetApprovalForAll(&_BAYC1155.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BAYC1155 *BAYC1155TransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BAYC1155.Contract.SetApprovalForAll(&_BAYC1155.TransactOpts, operator, approved)
}

// BAYC1155ApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the BAYC1155 contract.
type BAYC1155ApprovalForAllIterator struct {
	Event *BAYC1155ApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BAYC1155ApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BAYC1155ApprovalForAll)
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
		it.Event = new(BAYC1155ApprovalForAll)
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
func (it *BAYC1155ApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BAYC1155ApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BAYC1155ApprovalForAll represents a ApprovalForAll event raised by the BAYC1155 contract.
type BAYC1155ApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_BAYC1155 *BAYC1155Filterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*BAYC1155ApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BAYC1155.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BAYC1155ApprovalForAllIterator{contract: _BAYC1155.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_BAYC1155 *BAYC1155Filterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BAYC1155ApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BAYC1155.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BAYC1155ApprovalForAll)
				if err := _BAYC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_BAYC1155 *BAYC1155Filterer) ParseApprovalForAll(log types.Log) (*BAYC1155ApprovalForAll, error) {
	event := new(BAYC1155ApprovalForAll)
	if err := _BAYC1155.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BAYC1155TransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the BAYC1155 contract.
type BAYC1155TransferBatchIterator struct {
	Event *BAYC1155TransferBatch // Event containing the contract specifics and raw log

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
func (it *BAYC1155TransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BAYC1155TransferBatch)
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
		it.Event = new(BAYC1155TransferBatch)
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
func (it *BAYC1155TransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BAYC1155TransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BAYC1155TransferBatch represents a TransferBatch event raised by the BAYC1155 contract.
type BAYC1155TransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_BAYC1155 *BAYC1155Filterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*BAYC1155TransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BAYC1155.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BAYC1155TransferBatchIterator{contract: _BAYC1155.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_BAYC1155 *BAYC1155Filterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *BAYC1155TransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BAYC1155.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BAYC1155TransferBatch)
				if err := _BAYC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_BAYC1155 *BAYC1155Filterer) ParseTransferBatch(log types.Log) (*BAYC1155TransferBatch, error) {
	event := new(BAYC1155TransferBatch)
	if err := _BAYC1155.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BAYC1155TransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the BAYC1155 contract.
type BAYC1155TransferSingleIterator struct {
	Event *BAYC1155TransferSingle // Event containing the contract specifics and raw log

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
func (it *BAYC1155TransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BAYC1155TransferSingle)
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
		it.Event = new(BAYC1155TransferSingle)
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
func (it *BAYC1155TransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BAYC1155TransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BAYC1155TransferSingle represents a TransferSingle event raised by the BAYC1155 contract.
type BAYC1155TransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_BAYC1155 *BAYC1155Filterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*BAYC1155TransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BAYC1155.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BAYC1155TransferSingleIterator{contract: _BAYC1155.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_BAYC1155 *BAYC1155Filterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *BAYC1155TransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _BAYC1155.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BAYC1155TransferSingle)
				if err := _BAYC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_BAYC1155 *BAYC1155Filterer) ParseTransferSingle(log types.Log) (*BAYC1155TransferSingle, error) {
	event := new(BAYC1155TransferSingle)
	if err := _BAYC1155.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BAYC1155URIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the BAYC1155 contract.
type BAYC1155URIIterator struct {
	Event *BAYC1155URI // Event containing the contract specifics and raw log

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
func (it *BAYC1155URIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BAYC1155URI)
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
		it.Event = new(BAYC1155URI)
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
func (it *BAYC1155URIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BAYC1155URIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BAYC1155URI represents a URI event raised by the BAYC1155 contract.
type BAYC1155URI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_BAYC1155 *BAYC1155Filterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*BAYC1155URIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BAYC1155.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &BAYC1155URIIterator{contract: _BAYC1155.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_BAYC1155 *BAYC1155Filterer) WatchURI(opts *bind.WatchOpts, sink chan<- *BAYC1155URI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _BAYC1155.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BAYC1155URI)
				if err := _BAYC1155.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_BAYC1155 *BAYC1155Filterer) ParseURI(log types.Log) (*BAYC1155URI, error) {
	event := new(BAYC1155URI)
	if err := _BAYC1155.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
