//go:build ignore

package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"go-and-solidity/artifacts-go/storage"
	"math/big"
	"time"
)

var (
	Address    = "0x3c2Ea2A42529C55e914420f75bB3c25136E2b876"
	PrivateKey = "e55409fb049aaf42edad8b92347c112e3081d9019a20dfdbb5c8d4bf32bfecce"
	Rpc        = "https://matic-mumbai.chainstacklabs.com"
)

// 部署polygon 测试网
func main() {
	// Deploy the storage contract
	client, err := ethclient.Dial(Rpc)
	if err != nil {
		fmt.Println("new rpc client err:", err)
		return
	}
	privateKey, err := crypto.HexToECDSA(PrivateKey)
	if err != nil {
		fmt.Println("private key err:", err)
		return
	}
	nonce, err := client.NonceAt(context.TODO(), common.HexToAddress(Address), nil)
	if err != nil {
		fmt.Println("get nonce err:", err)
		return
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(80001))
	if err != nil {
		fmt.Println("new auth err:", err)
		return
	}
	auth.GasPrice, _ = client.SuggestGasPrice(context.TODO())
	auth.Nonce = big.NewInt(int64(nonce))

	// 部署
	contractAddr, tx, _, err := storage.DeployStorage(auth, client)
	if err != nil {
		fmt.Println("deploy err:", err)
		return
	}

	// wait mine
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	receipt, err := bind.WaitDeployed(ctx, client, tx)
	if err != nil {
		fmt.Println("mine err:", err)
		return
	}
	fmt.Println("=== deploy success ===")
	fmt.Println("tx hash:", receipt.Hash().Hex())
	fmt.Println("tx type:", tx.Type())
	fmt.Println("contract address", contractAddr.Hex())
}
