// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/**
 * Events 允许记录到以太坊区块链。事件的一些用例是：
监听事件和更新用户界面
一种廉价的存储方式
 */
contract Event {
    //事件声明
    //最多可以索引3个参数。
    //通过索引参数对日志进行过滤
    event Log(address indexed sender, string message);
    event AnotherLog();

    function test() public {
        emit Log(msg.sender, "Hello World!");
        emit Log(msg.sender, "Hello EVM!");
        emit AnotherLog();
    }
}
