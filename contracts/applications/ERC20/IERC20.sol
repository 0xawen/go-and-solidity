// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IERC20 {
    // @dev 当 value 单位货币从账号 from 转移到另一个账号 to时候
    event Transfer(address indexed from, address indexed to, uint256 value);

    // @dev 释放条件：当 `value` 单位的货币从账户 (`owner`) 授权给另一账户 (`spender`)时.
    event Approval(address indexed owner, address indexed spender, uint256 value);

    // @dev 返回代币总供应量
    function totalSupply() external view returns(uint256);

    // @dev 返回账号所持有的代币数量
    function balanceOf(address account) external view returns(uint256);

    // @dev 转账 `amount` 单位代币，从调用者账户到另一账户 `to`
    // 如果成功返回 true, 释放 transfer 事件
    function transfer(address to, uint256 amount) external returns (bool);

    // @dev 返回`owner`账户授权给`spender`账户的额度，默认是0
    // 当 {approve} 或 {transferFrom} 被调用时，`allowance`会改变.
    function allowance(address owner, address spender) external view returns (uint256);

    // @dev 调用者给 spender 账户授权amount数量代币
    // 成功返回true, 释放事件
    function approve(address spender, uint256 amount) external returns (bool);

    // @dev 通过授权机制，从 from 账户向 to 账户转账 amount代币。转账的部分会从调用者的`allowance`中扣除
    // 成功返回true, 释放 transfer 事件
    function transferFrom(address from, address to, uint256 amount) external returns (bool);
}
