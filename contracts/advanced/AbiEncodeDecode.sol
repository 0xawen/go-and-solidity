// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IERC20 {
    function transfer(address, uint) external;
}

contract Token {
    function transfer(address, uint)external {}
}


contract AbiEncode {
    function test(address _contract, bytes calldata data) external{
        (bool ok, ) = _contract.call(data);
        require(ok, "call failed");
    }

    function encodeWithSignature(address to, uint amount) external pure returns (bytes memory) {
        return abi.encodeWithSignature("transfer(address,uint256)",to, amount);
    }

    function encodeWithSelector(address to, uint amount) external pure returns (bytes memory) {
        return abi.encodeWithSelector(IERC20.Transfer.selector, to, amount);
    }
}

contract AbiDecode {
    struct MyStruct {
        string name;
        uint[2] nums;
    }

    function encode(
        uint x,
        address addr,
        uint[] calldata arr,
        MyStruct calldata myStruct
    ) external pure returns (bytes memory) {
        return abi.encode(x, addr, arr, myStruct);
    }

    function decode(
        bytes calldata data
    )
    external
    pure
    returns (uint x, address addr, uint[] memory arr, MyStruct memory myStruct)
    {
        // (uint x, address addr, uint[] memory arr, MyStruct myStruct) = ...
        (x, addr, arr, myStruct) = abi.decode(data, (uint, address, uint[], MyStruct));
    }
}
