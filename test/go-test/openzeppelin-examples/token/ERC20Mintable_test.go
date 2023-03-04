package token_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"go-and-solidity/artifacts-go/openzeppelin-examples/token"
	"go-and-solidity/test/go-test/helper"
	"math/big"
	"testing"
)

func TestERC20Mintable(t *testing.T) {
	simBackend := helper.NewSimulateBackend()
	defer simBackend.Close()

	account0 := helper.GetAccountPrivateKey(0)
	deployAuth, err := bind.NewKeyedTransactorWithChainID(account0, big.NewInt(1337))
	if err != nil {
		t.Fatal("new deploy auth err:", err)
	}

	contractAddr, tx, ERC20Mintable, err := token.DeployERC20Mintable(deployAuth, simBackend, "erc20-test", "TEST")
	if err != nil {
		t.Fatal("deploy err:", err)
	}
	// mine
	simBackend.Commit()
	fmt.Println("tx hash:", tx.Hash().Hex())
	fmt.Printf("contract Address:%s\n", contractAddr.Hex())

	// test owner mint
	t.Run("owner mint", func(t *testing.T) {
		// account0 is owner
		auth, _ := bind.NewKeyedTransactorWithChainID(helper.GetAccountPrivateKey(0), big.NewInt(1337))
		mintTX, err := ERC20Mintable.Mint(auth, helper.GetAccountAddress(0), big.NewInt(1000))
		if err != nil {
			t.Error("mint error:", err)
		}
		t.Log("mint tx hash:", mintTX.Hash().Hex())
		t.Log("tx type", mintTX.Type())
	})
	// test not owner mint
	t.Run("not owner mint", func(t *testing.T) {
		// account1 is not owner
		auth, _ := bind.NewKeyedTransactorWithChainID(helper.GetAccountPrivateKey(1), big.NewInt(1337))
		mintTX, err := ERC20Mintable.Mint(auth, helper.GetAccountAddress(1), big.NewInt(1000))
		if err != nil {
			t.Log("mint should error:", err)
			return
		}
		t.Log("mint tx hash:", mintTX.Hash().Hex())
		t.Log("tx type", mintTX.Type())
	})
}
