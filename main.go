package main

import "fmt"

// 编译合约
//go:generate ./bin/solc-windows-0.8.19.exe --abi --bin --optimize --overwrite -o ./build ./contracts/Storage.sol
// 绑定合约
//go:generate ./bin/abigen.exe --abi=./build/Storage.abi --bin=./build/Storage.bin --pkg=storage --type Storage --out=./artifacts/storage/storage.go

func main() {
	fmt.Println("用于编译合约和生成go绑定合约程序")
}
