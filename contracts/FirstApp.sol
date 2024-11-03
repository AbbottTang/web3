// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract Counter {
    uint256 public count;

    //获取当前公共变量count
    function get() public view returns (uint256) {
        return count;
    }

    //count+1
    function inc() public {
        count += 1;
    }

    //count-1,如果count=0，则失败-原因：‌uint‌：32位无符号整数，范围从0到4294967295，用于存储非负整数，避免整数溢出问题，在某些操作中比有符号整数更快。
    function dec() public {
        count -= 1;
    }
}
