// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/**
 * 视图和纯函数
Getter函数可以声明为view或pure。
View函数声明状态不会被改变。
纯函数声明不改变或读取状态变量。
 */
contract ViewAndPure {
    uint256 public x = 1;

    //承诺不修改状态。
    function addToX(uint256 y) public view returns (uint256) {
        return x + y;
    }

    ////保证不修改或读取状态。
    function add(uint256 i, uint256 j) public pure returns (uint256) {
        return i + j;
    }
}
