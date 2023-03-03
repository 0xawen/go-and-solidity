package applications_test

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"go-and-solidity/artifacts-go/applications"
	"go-and-solidity/artifacts-go/openzeppelin-examples/token"
	"go-and-solidity/test/helper"
	"math/big"
	"testing"
)

// 测试空投合约
func TestAirdrop(t *testing.T) {
	simBackend := helper.NewSimulateBackend()
	defer simBackend.Close()

	account0 := helper.GetAccountPrivateKey(0)
	deployAuth, err := bind.NewKeyedTransactorWithChainID(account0, big.NewInt(1337))
	if err != nil {
		t.Fatal("new deploy auth err:", err)
	}

	// 部署erc20代币合约
	tokenAddr, tx, ERC20Mintable, err := token.DeployERC20Mintable(deployAuth, simBackend, "erc20-test", "TEST")
	if err != nil {
		t.Fatal("deploy err:", err)
	}
	simBackend.Commit()
	fmt.Println("tx hash:", tx.Hash().Hex())
	fmt.Printf("contract Address:%s\n", tokenAddr.Hex())

	// 部署空投合约
	airDripAddr, aTx, AirdropInstant, err := applications.DeployAirdrop(deployAuth, simBackend)
	if err != nil {
		t.Fatal("deploy airdrop contract err:", err)
	}
	fmt.Println("airdrop tx hash:", aTx.Hash().Hex())
	fmt.Printf("airdrop contract Address:%s\n", airDripAddr.Hex())

	//  -----------------------
	auth0, _ := bind.NewKeyedTransactorWithChainID(helper.GetAccountPrivateKey(0), big.NewInt(1337))
	_, err = ERC20Mintable.Mint(auth0, helper.GetAccountAddress(0), big.NewInt(1000))
	if err != nil {
		t.Fatal("mint erc20 token error:", err)
	}
	simBackend.Commit()
	beforeBalance, err := ERC20Mintable.BalanceOf(nil, helper.GetAccountAddress(0))
	if err != nil {
		t.Fatal("get balance error:", err)
	}
	fmt.Println("banlance shoud equal 1000 =>", beforeBalance)

	// 授权给airdrop合约
	_, err = ERC20Mintable.Approve(auth0, airDripAddr, big.NewInt(1000))
	if err != nil {
		t.Fatal("approve error:", err)
	}
	simBackend.Commit()

	addresses := []common.Address{
		helper.GetAccountAddress(1),
		helper.GetAccountAddress(2),
	}
	amount := []*big.Int{
		big.NewInt(100),
		big.NewInt(100),
	}
	// 执行空投合约
	airdropTx, err := AirdropInstant.MultiTransferToken(auth0, tokenAddr, addresses, amount)
	if err != nil {
		t.Fatal("airdrop err:", err)
	}
	simBackend.Commit()
	fmt.Println("airdropTx:", airdropTx.Hash().Hex())
	afterBalance, _ := ERC20Mintable.BalanceOf(nil, helper.GetAccountAddress(0))
	fmt.Println("after balance equal 800 =>", afterBalance)

	account1Balance, _ := ERC20Mintable.BalanceOf(nil, helper.GetAccountAddress(1))
	account2Balance, _ := ERC20Mintable.BalanceOf(nil, helper.GetAccountAddress(2))
	fmt.Println("account1 balance equal 100 =>", account1Balance)
	fmt.Println("account2 balance equal 100 =>", account2Balance)
}
