# 钻石标准

参考资料：

- https://learnblockchain.cn/article/1398
- https://dev.to/mudgen/understanding-diamonds-on-ethereum-1fb
- https://github.com/mudgen/Diamond 实现标准

钻石
钻石切面
钻石存储 **
钻石cut **
钻石loupe
用在哪里,真实的案例

## diamond-3
1. 编译
```bash
./bin/solc-windows-0.8.19.exe --abi --bin --optimize --overwrite -o ./build/diamond/diamond-3/ ./contracts/diamond/diamond-3/Diamond.sol
```
2. go编译绑定
```bash
./bin/abigen.exe --abi=./build/diamond/diamond-3/Diamond.abi --bin=./build/diamond/diamond-3/Diamond.bin --pkg=diamond_3 --type Diamond3 --out=./artifacts-go/diamond/diamond-3/diamond3.go
```
