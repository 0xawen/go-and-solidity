// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// 在合约代码中调用已部署合约

// 前提： 已知部署合约的代码或者是接口



contract OtherContract {
    uint256 private _x = 0; // state val

    event Log(uint amount, uint gas);

    function getBalance() view public returns(uint) {
        return address(this).balance;
    }

    function setX(uint x) external payable {
        _x = x;
        if (msg.value > 0) {
            emit log(msg.value, gasleft());
        }
    }

    function getX() external view returns(uint x) {
        x = _x;
    }
}


contract CallOtherContract {
    // 方式1
    function callSetX(address _Address, uint256 x) external {
        OtherContract(_Address).setX(x);
    }

    // 方式2
    function callGetX2(address _Address) external view returns(uint x) {
        OtherContract oc =OtherContract(_Address);
        x = oc.getX();
    }
}
