// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

/**
 * 在solid中有3种类型的变量
当地的
在函数内部声明
没有存储在区块链上
状态
在函数外部声明
存储在区块链上
全局（提供关于区块链的信息）
 */

contract Variables {
    //状态变量存储在区块链上
    string public text = "Hello";
    uint256 public num = 1234;

    function doSomething() public view {
        //当地变量没有存储在区块链上
        uint256 i = 45677;

        //全局变量
        uint256 timestamp = block.timestamp; //当前块时间戳
        address sender = msg.sender; //访客地址
    }
}
