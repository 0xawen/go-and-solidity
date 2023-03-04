package helper

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
)

/*
	help for test
*/

var AccountMap = map[int]struct {
	Address    string
	PrivateKey string
}{
	0: {
		Address:    "0x3c2Ea2A42529C55e914420f75bB3c25136E2b876",
		PrivateKey: "e55409fb049aaf42edad8b92347c112e3081d9019a20dfdbb5c8d4bf32bfecce",
	},
	1: {
		Address:    "0x5b9C6b9F1fE79f8493FeC1bDebF76b1DECCcfcc0",
		PrivateKey: "eced2cae8cea8954b0c05ea6f8cea1dd8ef5c73e2b369cdcb7ebd988aadf8b72",
	},
	2: {
		Address:    "0x446Ae890e924722B2bdA91C47515E9499F589E21",
		PrivateKey: "ca4cbeb7c4aa4bb1dd6250d1fb5dbf31874e07348d9fbf2ae9e9a4cd2b84811d",
	},
	3: {
		Address:    "0x8Ea69adB4d5dd42260de5cba377e98A7492C730C",
		PrivateKey: "9bed65dce5f6add5a501226390fe6bf1e38689e432b3c0dfb99302cc44708487",
	},
	4: {
		Address:    "0x9fDC25EDb243CE12a8eCA6B4A6E64035B8Ad6F5e",
		PrivateKey: "ea5fed65aa5f39f668c0a50c210c3f861e222d11b35db15b1d9aa21749214545",
	},
	5: {
		Address:    "0xc44aa94e16605b2c1CC23293e80ADa043edA3386",
		PrivateKey: "3af058aa7bf2aa55f19aeab8517e5dd1aea03b9e02a1b3a97c5bb7d5fa1fdbfe",
	},
	6: {
		Address:    "0x0439d9DdCdFf849Dd7767fFe6f6DFbD026c8c322",
		PrivateKey: "c47f0ad63ccc26cbdf8d0b7a738ab0049e3860a6db235870b14a89a7a81b6812",
	},
	7: {
		Address:    "0x3D27F99550405955c7A436fC9C9774b216d9Ba45",
		PrivateKey: "6d1b1b1da6ea355bf8092494595f3c5726d3e3c5fe1cd273dc7210a7653ee658",
	},
	8: {
		Address:    "0x6138FAB08e6F008337428bB99A5f5eE8E4f45595",
		PrivateKey: "e538acdfdc74bdf78f19cb67815d234253fa6808ebc2357f2566b820da80cfc5",
	},
	9: {
		Address:    "0x16BD7aaA24670b69077b3Fddd845b3b17FA156BC",
		PrivateKey: "e1028c777d8883c0bfde36479272d32e1ada2705f8067d86932ff2795e2476c4",
	},
}

// GetAccountAddress 根据index返回一个账号的地址
// index [0,9]
func GetAccountAddress(index int) common.Address {
	if index < 0 || index > 9 {
		panic("index out of range")
	}
	return common.HexToAddress(AccountMap[index].Address)
}

// GetAccountPrivateKey 根据index返回一个账号的私钥
// index [0,9]
func GetAccountPrivateKey(index int) *ecdsa.PrivateKey {
	if index < 0 || index > 9 {
		panic("index out of range")
	}
	pK, err := crypto.HexToECDSA(AccountMap[index].PrivateKey)
	if err != nil {
		panic(err)
	}

	return pK
}

// NewSimulateBackend 返回一个模拟的链后端,内置了10个账号，每一个账号有100eth
func NewSimulateBackend() *backends.SimulatedBackend {
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{
		common.HexToAddress("0x3c2Ea2A42529C55e914420f75bB3c25136E2b876"): {
			Balance: getBigint(),
		},
		common.HexToAddress("0x5b9C6b9F1fE79f8493FeC1bDebF76b1DECCcfcc0"): {
			Balance: getBigint(),
		},
		common.HexToAddress("0x446Ae890e924722B2bdA91C47515E9499F589E21"): {
			Balance: getBigint(),
		},
		common.HexToAddress("0x8Ea69adB4d5dd42260de5cba377e98A7492C730C"): {
			Balance: getBigint(),
		},
		common.HexToAddress("0x9fDC25EDb243CE12a8eCA6B4A6E64035B8Ad6F5e"): {
			Balance: getBigint(),
		},
		common.HexToAddress("0xc44aa94e16605b2c1CC23293e80ADa043edA3386"): {
			Balance: getBigint(),
		},
		common.HexToAddress("0x0439d9DdCdFf849Dd7767fFe6f6DFbD026c8c322"): {
			Balance: getBigint(),
		},
		common.HexToAddress("0x3D27F99550405955c7A436fC9C9774b216d9Ba45"): {
			Balance: getBigint(),
		},
		common.HexToAddress("0x6138FAB08e6F008337428bB99A5f5eE8E4f45595"): {
			Balance: getBigint(),
		},
		common.HexToAddress("0x16BD7aaA24670b69077b3Fddd845b3b17FA156BC"): {
			Balance: getBigint(),
		},
	}, 4712388)

	return sim
}

func getBigint() *big.Int {
	a, _ := new(big.Int).SetString("100000000000000000000", 10)
	return a
}
