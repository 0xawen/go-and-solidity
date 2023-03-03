package reentrancy_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"go-and-solidity/artifacts-go/security/reentrancy"
	"go-and-solidity/test/helper"
	"math/big"
	"testing"
)

func TestReentrancy(t *testing.T) {
	simBackend := helper.NewSimulateBackend()
	defer simBackend.Close()

	account0 := helper.GetAccountPrivateKey(0)
	deployAuth, err := bind.NewKeyedTransactorWithChainID(account0, big.NewInt(1337))
	if err != nil {
		t.Fatal("new deploy auth err:", err)
	}

	// 部署 Bank 有漏洞合约的
	bankContractAddr, _, bankInstance, err := reentrancy.DeployBank(deployAuth, simBackend)
	if err != nil {
		t.Fatal("deploy bank err:", err)
	}
	simBackend.Commit()

	// 向合约赚钱
	auth0, err := bind.NewKeyedTransactorWithChainID(account0, big.NewInt(1337))
	value, _ := new(big.Int).SetString("10000000000000000000", 10)
	auth0.Value = value
	_, err = bankInstance.Deposit(auth0)
	if err != nil {
		t.Fatal("deposit 10 eth")
	}
	simBackend.Commit()

	bankBalance, _ := bankInstance.GetBalance(nil)
	fmt.Println("bank balance should be 10 eth, =>", bankBalance)

	// 部署攻击合约
	_, _, AttackInstance, err := reentrancy.DeployAttack(deployAuth, simBackend, bankContractAddr)
	if err != nil {
		t.Fatal("deploy attack contract err:", err)
	}
	simBackend.Commit()

	// 攻击
	auth1, _ := bind.NewKeyedTransactorWithChainID(account0, big.NewInt(1337))
	value1, _ := new(big.Int).SetString("1000000000000000000", 10)
	auth1.Value = value1
	_, err = AttackInstance.Attack(auth1)
	if err != nil {
		t.Fatal("attack err:", err)
	}
	simBackend.Commit()

	// 查询攻击合约余额
	attackBalance, _ := AttackInstance.GetBalance(nil)
	fmt.Println("attack balance should be 11 eth, =>", attackBalance)
}
