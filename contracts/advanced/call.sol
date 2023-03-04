// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// call 是address类型的低级成员函数，它用来与其他合约交互。它的返回值为(bool, data)，分别对应call是否成功以及目标函数的返回值。
// - call是solidity官方推荐的通过触发fallback或receive函数发送ETH的方法。
// - 不推荐用call来调用另一个合约，因为当你调用不安全合约的函数时，你就把主动权交给了它。推荐的方法仍是声明合约变量后调用函数
// - 当我们不知道对方合约的源代码或ABI，就没法生成合约变量；这时，我们仍可以通过call调用对方合约的函数。


// 使用
// 目标合约地址.call(二进制编码)
// 目标合约地址.call{value:发送数额, gas:gas数额}(二进制编码);



contract Call {
    function CallSetX(address payable _addr, uint256 x) public payable {
        // call setX 同时发送 eth
        // call不是调用合约的推荐方法，因为不安全。但他能让我们在不知道源代码和ABI的情况下调用目标合约，很有用
        (bool success1, bytes memory data1) = _addr.call{value: msg.value}(abi.encodeWithSignature("setX(uint256)", x));
        (bool success2, bytes memory data2) = _addr.call(abi.encodeWithSignature("getX()"));
    }
}
