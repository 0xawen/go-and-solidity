# solidity 知识点

## 1. 基础语法知识点
(掌握语法基础的可以跳过这一部分)

## 2. 进阶知识点

### 函数重载

相同函数名字，不同参数类型的函数可以重载。但是，相同函数名字，相似参数类型的函数不能重载。
```solidity
// 可以重载
function test(uint a) public returns (uint) {
    return a;
}

function test(uint a, uint b) public returns (uint) {
    return a + b;
}
```

```solidity
// 不能重载
function test(uint8 a) public returns (uint) {
    return a;
}

function test(uint256 b) public returns (uint) {
    return b;
}
```

### 库合约

### 接收以太币 (receive/fallback)

### 发送以太币(transfer/send/call)

### call 直接调用

### delegateCall 委托调用

### create 从合约中创建合约

### create2 从的合约中创建合约

## 代理合约

## 可升级代理(selector clash)

## 透明代理(transparent proxy)
