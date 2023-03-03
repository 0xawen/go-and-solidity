# go-and-solidity 实践说明

## 编译合约
```
./bin/solc-windows-0.8.19.exe --abi --bin --optimize --overwrite -o ./build ./contracts/Storage.sol
```

```
./bin/solc-static-linux-0.8.19 --abi --bin --optimize --overwrite -o ./build ./contracts/Storage.sol
```

## go绑定合约

```
mkdir ./artifacts/storage
```

```
./bin/abigen.exe --abi=./build/Storage.abi --bin=./build/Storage.bin --pkg=storage --type Storage --out=./artifacts/storage/storage.go
```

## go测试合约

go test -v ./test/...
```
go test -v ./test/storage_test.go
```
## go部署合约

go run scripts/deploy_storage.go

```
同一个package下,有多个main。在文件开头添加`//go:build ignore`。
```

## 第三方库的使用

直接使用 git submodule 来管理第三方库。例如oppenzeppelin使用如下：

```
git submodule add https://github.com/OpenZeppelin/openzeppelin-contracts.git contracts/openzeppelin-contracts
```

```
git submodule update // 更新本地项目子模块到最新版本
git submodule update --remote // 更新远程项目子模块到最新版本
```

## Todo

- 安全分析: Mythril-Slither-Manticore-MythX-Echidna-Oyente
