// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/**
 * 修饰符是可以在函数调用之前和/或之后运行的代码。
修饰语可用于：
限制访问
验证输入
防止重入攻击
 */
contract FunctionModifier {
    ////我们将使用这些变量来演示如何使用
    //修饰符。
    address public owner;
    uint256 public x = 10;
    bool public locked;

    constructor() {
        //将交易发送方设置为合约的所有者。
        owner = msg.sender;
    }

    ////检查调用者是否是owner的修饰符
    modifier onlyOwner() {
        require(msg.sender == owner, "Not owner");
        //下划线是只在内部使用的特殊字符
        //一个函数修饰符，它告诉Solidity
        //执行剩下的代码。
        _;
    }
    ////修饰符可以接受输入。这个修饰符检查
    //传入的地址不是零地址。
    modifier validAddress(address _addr) {
        require(_addr != address(0), "Not valid address");
        _;
    }
    //修饰符可以在函数之前和/或之后调用。
    //这个修饰符防止函数被while调用
    //它仍在执行
    modifier noReentrancy() {
        require(!locked, "No reentrancy");
        locked = true;
        _;
        locked = false;
    }

    function decrement(uint256 i) public noReentrancy {
        x -= i;
        if (i > 1) {
            decrement(i - 1);
        }
    }
}
