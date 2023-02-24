# go and solidity

- 目的：使用go开发与测试solidity合约。

### 起因

hardhat 是目前最流行的solidity开发测试框架。但是hardhat需要使用到 ts/js来测试合约。

而我一直对 ts/js 不感冒，虽然可以使用起来也没有问题，但是感觉使用go来开发测试合约更加的顺手。

于是，就有了这个项目`go-and-solidity`。 同时，也巩固一下go知识和solidity知识。


### 项目结构
- `artifacts`：存放abigen生成的go文件
- `bin`: 存放solidity编译工具与abi工具
- `build`：存放solidity编译后的abi与bin文件
- `contracts`：存放solidity合约
- `docs`：存放文档
- `deploy`：存放部署合约的go脚本
- `scripts`：存放部署合约的go脚本
- `test`：存放solidity测试go脚本


更多相关文档说明存在`docs`目录下, 请移步查看。


## 编译合约

./bin/solc-windows-0.8.19.exe --abi --bin --optimize --overwrite -o ./build ./contracts/Storage.sol

./bin/solc-static-linux-0.8.19 --abi --bin --optimize --overwrite -o ./build ./contracts/Storage.sol

## go绑定合约

mkdir ./artifacts/storage
./bin/abigen.exe --abi=./build/Storage.abi --bin=./build/Storage.bin --pkg=storage --type Storage --out=./artifacts/storage/storage.go

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