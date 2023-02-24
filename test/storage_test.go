package test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"go-and-solidity/artifacts/storage"
	"math/big"
	"testing"
)

// Test deploy and interact
func TestDeployStorage(t *testing.T) {
	// generate a random account and a funded simulator
	key, _ := crypto.GenerateKey()
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	if err != nil {
		t.Fatal(err)
	}
	blockGasLimit := uint64(4712388)
	sim := backends.NewSimulatedBackend(core.GenesisAlloc{
		auth.From: {
			Balance: big.NewInt(1000000000000000000),
		},
	}, blockGasLimit)

	// deploy contract
	contractAddr, tx, storageInstance, err := storage.DeployStorage(auth, sim)
	// mine
	sim.Commit()

	fmt.Println("contract address:", contractAddr.Hex())
	fmt.Println("tx:", tx.Hash())

	number, err := storageInstance.Value(nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("before value is:", number.String())

	storeTx, err := storageInstance.Store(auth, big.NewInt(666))
	// mine
	sim.Commit()
	fmt.Println("store tx:", storeTx.Hash())

	number, err = storageInstance.Value(nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("after value is:", number.String())
}
