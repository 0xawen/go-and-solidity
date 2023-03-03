# openzeppelin 合约库使用示例


## token
1. 编译
```bash 
./bin/solc-windows-0.8.19.exe --abi --bin --optimize --overwrite -o ./build/openzeppelin-examples/token ./contracts/openzeppelin-examples/token/ERC20Mintable.sol
```
2. go绑定
```bash 
./bin/abigen.exe --abi=./build/openzeppelin-examples/token/ERC20Mintable.abi --bin=./build/openzeppelin-examples/token/ERC20Mintable.bin --pkg=token --type ERC20Mintable --out=./artifacts-go/openzeppelin-examples/token/ERC20Mintable.go
```
