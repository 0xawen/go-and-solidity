// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// 发送eth

// 1.transfer
// `接收者地址`.transfer(amount)
// transfer()的gas限制是2300，足够用于转账，但对方合约的fallback()或receive()函数不能实现太复杂的逻辑
// transfer()如果转账失败，会自动revert（回滚交易）

// 2.send
// `接收者地址`.send(amount)
// send()的gas限制是2300，足够用于转账，但对方合约的fallback()或receive()函数不能实现太复杂的逻辑。
// send()如果转账失败，不会revert
// send()的返回值是bool，代表着转账成功或失败，需要额外代码处理一下

// 3.call
// `接收方地址`.call{value: amount}("")
// call()没有gas限制，可以支持对方合约fallback()或receive()函数实现复杂逻辑
// call()如果转账失败，不会revert
// call()的返回值是(bool, data)，其中bool代表着转账成功或失败，需要额外代码处理一下

// 总结
// call没有gas限制，最为灵活，是最提倡的方法；
// transfer有2300 gas限制，但是发送失败会自动revert交易，是次优选择；
// send有2300 gas限制，而且发送失败不会自动revert交易，几乎没有人用它。

error SendFailed(); // 用send发送ETH失败error
error CallFailed(); // 用call发送ETH失败error

contract SendEth {
    // 构造函数， payable 使得部署的时候可以转eth进去
    constructor() payable {}
    // receive() 函数 接受eth
    receive() external payable{}

    // 使用 transfer 发送 eth
    function transferEth(address payable _to, uint256 _amount) external payable {
        _to.transfer(_amount);
    }

    // 使用 send 发送eth
    function sendEth(address payable _to, uint256 _amount) external payable {
        // send 有返回值
        bool success = _to.send(_amount);
        if (!success) {
            revert SendFailed();
        }
    }

    // 使用 call 发送 eth
    function callEth(address payable _to, uint256 _amount) external payable {
        (bool success,) = _to.call{value: _amount}("");
        if (!success) {
            revert CallFailed();
        }
    }

}
