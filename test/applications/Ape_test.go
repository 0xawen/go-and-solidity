package applications_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"go-and-solidity/artifacts/applications"
	"go-and-solidity/test/helper"
	"math/big"
	"testing"
)

// Ape 是一个erc721 合约
func TestApe(t *testing.T) {
	simBackend := helper.NewSimulateBackend()
	defer simBackend.Close()

	account0 := helper.GetAccountPrivateKey(0)

	deployAuth, err := bind.NewKeyedTransactorWithChainID(account0, big.NewInt(1337))
	if err != nil {
		t.Fatal("new deploy auth err:", err)
	}
	contractAddr, tx, ApeInstance, err := applications.DeployApe(deployAuth, simBackend, "ApeTest", "APE")
	if err != nil {
		t.Fatal("deploy err:", err)
	}
	// mine
	simBackend.Commit()
	fmt.Println("tx hash:", tx.Hash().Hex())
	fmt.Printf("contract Address:%s\n", contractAddr.Hex())

	// --------------------------------
	t.Run("test mint", func(t *testing.T) {
		// account0 is owner
		auth, _ := bind.NewKeyedTransactorWithChainID(helper.GetAccountPrivateKey(0), big.NewInt(1337))
		mintTx, err := ApeInstance.Mint(auth, helper.GetAccountAddress(1), big.NewInt(1))
		if err != nil {
			t.Error("mint token err:", err)
		}
		t.Log("tx hash:", mintTx.Hash().Hex())

	})

	// todo test other function
}
