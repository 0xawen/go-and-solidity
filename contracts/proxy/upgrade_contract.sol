// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// 可升级合约
//
//  Call ---call-> Proxy --delegate-->  old logic
//                               |
//                               |--->  new logic
// 它是一个可以改变逻辑合约的代理合约，给不可更改的智能合约增加了升级功能。
// 但是，这个合约有选择器冲突的问题，存在安全隐患。之后我们会介绍解决这一隐患的可升级合约标准：透明代理和UUPS

contract SimpleUpgrade {
    address public implementation; // logic address
    address public admin; // admin 地址
    string public words; // 通过逻辑合约的函数改变

    constructor(address _implementation) {
        admin = msg.sender;
        implementation = _implementation;
    }

    // 将调用委托给逻辑合约
    fallback() external payable{
        (bool success,bytes memory data) = implementation.delegatecall(msg.data);
    }

    // 升级函数，改变逻辑合约地址，只能由admin调用
    function upgrade(address newImplementation) external {
        require(msg.sender == admin);
        implementation = newImplementation;
    }
}

// 旧逻辑合约

contract Logic1 {
    // 状态变量和proxy合约一致，防止插槽冲突
    address public implementation;
    address public admin;
    string public words; // 字符串，可以通过逻辑合约的函数改变

    // 改变proxy中状态变量，选择器： 0xc2985578
    function foo() public{
        words = "old";
    }
}

// 新逻辑合约

contract Logic2 {
    // 状态变量和proxy合约一致，防止插槽冲突
    address public implementation;
    address public admin;
    string public words; // 字符串，可以通过逻辑合约的函数改变

    // 改变proxy中状态变量，选择器：0xc2985578
    function foo() public{
        words = "new";
    }
}
