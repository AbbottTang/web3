// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract EtherWallet {
    /**
     * 任何人都可以发送金额到合约
只有 owner 可以取款
3 种取钱方式
     */
    address payable public immutable owner;
    event Log(string funName, address from, uint256 value, bytes data);

    constructor() {
        owner = payable(msg.sender);
    }

    receive() external payable {
        emit Log("receive", msg.sender, msg.value, "");
    }

    function withdraw1() external {
        require(msg.sender == owner, "Not owner");
        //owner.transfer 相比 msg.sender 更消耗gas
        //owner.transfer(address(this).balance);
        //在Solidity中，payable 是一个修饰符，用于标记一个函数可以接受以太币转账。它不是一个可以调用的函数或方法。
        //你不能直接将 payable 作为函数来调用，如 payable(msg.sender) 是不正确的。
        payable(msg.sender).transfer(100);
    }

    function withdraw2() external {
        require(msg.sender == owner, "Not owner");
        bool success = payable(msg.sender).send(200);
        require(success, "Send Failed");
    }

    function withdraw3() external {
        require(msg.sender == owner, "Not owner");
        (bool success, ) = msg.sender.call{value: address(this).balance}("");
        require(success, "Call Failed");
    }

    //查看余额
    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }
}
