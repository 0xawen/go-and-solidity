package applications_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"go-and-solidity/artifacts-go/applications"
	"go-and-solidity/test/go-test/helper"
	"math/big"
	"testing"
)

// 测试以太坊钱包
func TestEtherWallet(t *testing.T) {
	// get a simulate backend
	simBackend := helper.NewSimulateBackend()
	// get a account private key
	amount0 := helper.GetAccountPrivateKey(0)
	// new deploy auth
	deployAuth, err := bind.NewKeyedTransactorWithChainID(amount0, big.NewInt(1337))
	if err != nil {
		t.Fatal("new deploy auth err:", err)
	}

	contractAddr, tx, ethWalletInstance, err := applications.DeployEtherWallet(deployAuth, simBackend)
	if err != nil {
		t.Fatal("deploy err:", err)
	}
	// mine
	simBackend.Commit()
	fmt.Println("tx hash:", tx.Hash().Hex())
	fmt.Printf("EtherWallet Address:%s\n", contractAddr.Hex())

	// test contract function
	balance, err := ethWalletInstance.GetBalance(nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("balance:", balance)
	err = simBackend.Close()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(helper.GetAccountAddress(0).Hex())
}
