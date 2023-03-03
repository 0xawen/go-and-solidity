// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

// 导入erc20
import "../../openzeppelin-contracts/contracts/token/ERC20/ERC20.sol";
// 导入权限控制合约
import "../../openzeppelin-contracts/contracts/access/Ownable.sol";

contract ERC20Mintable is ERC20, Ownable {

    constructor(string memory name_, string memory symbol_) ERC20(name_, symbol_){
    }

    function mint(address account, uint256 amount) public virtual onlyOwner {
        _mint(account, amount);
    }
}
