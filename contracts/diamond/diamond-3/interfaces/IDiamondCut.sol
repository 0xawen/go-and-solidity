// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IDiamondCut {
    // Add=0; Replace=1; Remove=2
    enum FacetCutAction {Add, Replace, Remove}

    struct FacetCut {
        address facetAddress;
        FacetCutAction action;
        bytes4[] functionSelectors;
    }

    // @notice 添加/替换/移除 可执行委托调用的函数
    // @param _diamondCut 具体到某个facet 的功能函数
    // @param _init 合约地址或者 执行_calldata的facet
    // @param _calldata _init 执行所需要的参数
    function diamondCut(FacetCut[] calldata _diamondCut, address _init, bytes calldata _calldata) external;

    event DiamondCut(FacetCut[] _diamondCut, address _init, bytes _calldata);
}
