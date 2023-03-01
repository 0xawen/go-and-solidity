// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


// receive()
// receive()只用于处理接收ETH。一个合约最多有一个receive()函数，声明方式与一般函数不一样，不需要function关键字：receive() external payable { ... }。
// receive()函数不能有任何的参数，不能返回任何值，必须包含external和payable。

// fallback()
// 在调用合约不存在的函数时被触发,可以用于接受eth，也可以用于代理合约。
// fallback()声明时不需要function关键字，必须由external修饰，一般也会用payable修饰


// 区别
/* 触发fallback() 还是 receive()?
   接收ETH
      |
 msg.data是空？
    /  \
  是    否
  /      \
receive()存在?   fallback()
/ \
是  否
/     \
receive()  fallback
*/

// receive()和payable fallback()均不存在的时候，向合约直接发送ETH将会报错，但是仍可以通过带有payable的函数向合约发送ETH

contract Fallback {
    // event
    event receivedCalled(address Sender, uint Value);
    event fallbackCalled(address Sender, uint Value, bytes Data);

    receive() external payable {
        // 不适宜写太多的逻辑,防止 out of gas;
        emit receivedCalled(msg.sender, msg.value);
    }

    // fallback
    fallback() external payable {
        emit fallbackCalled(msg.sender, msg.value, msg.data);
    }
}
