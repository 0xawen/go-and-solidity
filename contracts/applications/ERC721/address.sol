// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


library Address {
    // 判断地址是否存在 bytecode 判断一个地址是否是合约地址
    function isContract(address account) internal view returns(bool) {
        uint size;
        assembly{
            size := extcodesize(account)
        }
        return size > 0;
    }
}
