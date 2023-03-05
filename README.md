# go and solidity

目标：使用go开发与测试solidity合约。

### 起因

hardhat 是目前最流行的solidity开发测试框架。但是hardhat需要使用到 ts/js来测试合约。

而我一直对 ts/js 不感冒，虽然可以使用起来也没有问题，但是感觉使用go来开发测试合约更加的顺手。

于是，就有了这个项目`go-and-solidity`。 同时，也巩固一下go知识和solidity知识。

视情况,后面增加foundry工具支持。

### 项目结构
- `artifacts-go`：存放abigen生成的go文件
- `bin`: 存放solidity编译工具与abi工具
- `build`：存放solidity编译后的abi与bin文件
- `contracts`：存放solidity合约
- `docs`：存放文档
- `lib`: 第三方合约库,使用git submodule管理
- `script`：存放部署合约的go脚本或者foundry测试脚本
- `test`：存放solidity测试文件
  - `foundry-test` foundry 测试文件
  - `go-test` go 测试文件


更多相关文档说明存在`docs`目录下, 请移步查看。

### 更新计划

内容有点多,更新比较慢。

- [solidity基础知识点](./docs/solidity基础知识点.md)
- [solidity的重点](./docs/solidity重点.md)
- [solidity的难点](./docs/solidity难点_代理合约.md)
- [智能合约的常见应用](./docs/常见应用.md)
- [合约安全专题](./docs/合约安全.md)  待更新
- [defi专题](./docs/defi.md) 待更新
- [nft专题](./docs/nft.md) 待更新
- [opnezeppelin合约库](./docs/openzeppelin合约库.md)  待更新
- [HQ20合约库](./docs/HQ20合约库.md)  待更新
- [构建大型智能合约系统之钻石标准](./docs/diamond.md) 待更新

## 使用说明

1. 编译工具安装,查看[solc/abigen](./docs/solidity编译工具与abigen工具.md)
2. 使用go去测试合约,详细过程查看:[实践说明文档](./docs/go-and-solidity的实践说明.md)

## Foundry 工具的支持

foundry.toml 的 `src` 设置为 contracts。

forge install foundry-rs/forge-std --no-commit

foundry 文档： https://github.com/foundry-rs/forge-std/tree/master

## 其他

本仓库参考的教程有：
1. https://solidity-by-example.org/
2. https://github.com/AmazingAng/WTF-Solidity
