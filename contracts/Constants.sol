// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/**
 * 常量是不可修改的变量
 * 常量值是硬编码，使用常量可以节省gas花费
 */
contract Constants {
    //常量通常以大写约定
    address public constant MY_ADDRESS =
        0x777788889999AaAAbBbbCcccddDdeeeEfFFfCcCc;

    uint256 public constant MY_UINT = 123;
}
