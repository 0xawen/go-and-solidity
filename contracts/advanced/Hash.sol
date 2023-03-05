// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract Hash {
    bytes32 _msg = keccak256(abi.encodePacked("0xboxi"));

    function hash(uint256 _num, string memory _string, address _address) public pure returns(bytes32) {
        return keccak256(abi.encodePacked(_num, _string, _address));
    }

}
