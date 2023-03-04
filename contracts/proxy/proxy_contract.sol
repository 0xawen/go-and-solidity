// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// 代理合约
//
// 操作步骤：
//      首先部署逻辑合约Logic。
//      创建代理合约Proxy，状态变量implementation记录Logic合约地址。
//      Proxy合约利用回调函数fallback，将所有调用委托给Logic合约
//      最后部署调用示例Caller合约，调用Proxy合约。
//      - Logic合约和Proxy合约的状态变量存储结构相同，不然delegatecall会产生意想不到的行为，有安全隐患
//
// 调用链
//
//  Caller合约 -----call----> Proxy合约 -----delegate ----> Logic合约

// 逻辑合约 logic
contract Logic {
    address public implementation;
    uint public x = 99;
    event CallSuccess();

    function increment() external returns(uint){
        emit CallSuccess();
        return x;
    }
}

// 代理合约 proxy
contract Proxy {
    // 逻辑合约地址。implementation合约同一个位置的状态变量类型必须和Proxy合约的相同，不然会报错。
    address public implementation;
    constructor(address implementation_){
        implementation = implementation_;
    }
    fallback() external payable {
        address _implementation = implementation;
        assembly {
            // msg.data 拷贝到内存上
            // calldatacopy: 内存起始位置， calldata起始位置， calldata长度
            calldatacopy(0,0,calldatasize())

            // 利用 delegatecall 调用 implementation 合约
            let result := delegatecall(gas(), _implementation,0, calldatasize(), 0,0)

            // 将return data拷贝到内存
            // returndata操作码的参数：内存起始位置，returndata起始位置，returndata长度
            returndatacopy(0, 0, returndatasize())

            switch result
            // 失败revert
            case 0 {
                revert(0,returndatasize())
            }
            default {
                return(0, returndatasize())
            }
        }
    }
}

// 调用者 Caller
contract Caller {
    address public proxy; // 代理合约地址
    constructor(address proxy_) {
        proxy = proxy_;
    }

    function increment() external returns(uint) {
        (,bytes memory data) = proxy.call(abi.encodeWithSignature("increment()"));
        return abi.decode(data,(uint));
    }
}
