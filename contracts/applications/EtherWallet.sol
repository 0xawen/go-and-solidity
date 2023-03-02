// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

// from: https://solidity-by-example.org/app/ether-wallet/
//
// 功能：
//   所有人都可以发送eth, 但是只有拥有者可以提款
contract EtherWallet {
        address payable public owner;
        constructor() {
            owner = payable(msg.sender);
        }
        // 接受eth
        receive() external payable{}

        // 提款函数
        function withdraw(uint256 _amount) external {
            require(msg.sender == owner, "caller is not owner");
            require(address(this).balance <= _amount, "money not enough to withdraw");
            payable(msg.sender).transfer(_amount);
        }
        // 余额
        function getBalance() external view returns (uint256) {
            return address(this).balance;
        }
}
