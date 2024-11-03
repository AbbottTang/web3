// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract SimpleStorage {
    //你需要发送一个事务去读写状态变量
    //另一方面你可以自由的读取状态变量而无需任何费用

    //状态变量存储一个数字
    uint256 public num;

    function set(uint _num) public {
        num = _num;
    }

    function get() public view returns (uint256) {
        return num;
    }
}
