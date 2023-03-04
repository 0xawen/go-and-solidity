# solidity基础知识点

代码： contracts/basis (todo)

作用：用于快速的过一遍

参考：https://www.tutorialspoint.com/solidity/index.htm

## 类型
bool; address; uint8-256;int8-256;
## 变量
- 状态变量： 记录在链上的
- 本地变量： 函数内部定义的变量
- 全局变量： 链信息,比如区块的高度,交易发起人
## 以太单位
1 ether = 10^9 Gwei
## 时间单位
## 字符串 string
## 数组 array
## 枚举 enums (没有怎么使用过)
## 结构体 structs
## 映射 mapping
## 数据位置
- storage
- memory
- calldata
## 函数 function
- function xxx(params...) ... returns(xxx);
- private|internal|public|external
- view|pure
- virtual|override
- modifier
- constructors 构造函数
- fallback 函数
- receive 函数
- 内置错误处理函数
  - assert(bool condition)
  - require(bool condition)
  - require(bool condition, string memory message)
  - revert()
  - revert(string memory reason)
- 内置的数据函数
  - addmod(uint x, uint y, uint k) returns (uint) − computes (x + y) % k 
  - mulmod(uint x, uint y, uint k) returns (uint) − computes (x * y) % k
- 内置的加密函数
  - keccak256(bytes memory) returns (bytes32) 
  - ripemd160(bytes memory) returns (bytes20)
  - sha256(bytes memory) returns (bytes32)
  - ecrecover(bytes32 hash, uint8 v, bytes32 r, bytes32 s) returns (address)

## 事件 event
```solidity
event Deposit(address indexed _from, bytes32 indexed _id, uint _value);
```
## 合约继承
- virtual: 父合约中的函数，如果希望子合约重写，需要加上virtual关键字。
- override：子合约重写了父合约中的函数，需要加上override关键字。
## 抽象合约
至少有一个函数没有实现，即某个函数缺少主体{}中的内容，则必须将该合约标为abstract。
## 合约接口
## 库合约
## 导入其他合约
