// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract Immutable {
    /**
     * 不可变变量-常量与常量相似，
     * 不可变变量-常量的值可以在构造函数内部设置，但之后不能修改
     */
    //常量与变量命名都是大写
    address public immutable MY_ADDRESS;
    uint256 public immutable MY_UINT;

    constructor(uint256 _myUint) {
        MY_ADDRESS = msg.sender;
        MY_UINT = _myUint;
    }
}
