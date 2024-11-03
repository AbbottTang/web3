// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/**
 * 错误将撤销在事务期间对状态所做的所有更改。
可以通过调用require、revert或assert抛出错误。
Require用于在执行之前验证输入和条件。
Revert类似于require。有关详细信息，请参阅下面的代码。
Assert用于检查不应该为false的代码。断言失败可能意味着存在错误。
使用自定义错误来节省gas。
 */
contract Error {
    function testRequire(uint256 _i) public pure {
        require(_i > 10, "Input must be greater than 10");
    }

    function testRevert(uint256 _i) public pure {
        //当要检查的条件很复杂时，Revert是有用的。
        //这段代码所做的事情与上面的例子完全相同
        if (_i <= 10) {
            revert("Input must be greater than 10");
        }
    }

    uint256 public num;

    function testAssert() public view {
        // Assert应该只用于测试内部错误
        //检查不变量。
        //这里我们断言num总是等于0
        //因为无法更新num的值
        assert(num == 0);
    }

    //自定义错误
    error InsufficientBalance(uint256 balance, uint256 withdrawAmount);

    function testCustomError(uint256 _withdrawAmount) public view {
        uint256 bal = address(this).balance;
        if (bal < _withdrawAmount) {
            revert InsufficientBalance({
                balance: bal,
                withdrawAmount: _withdrawAmount
            });
        }
    }
}

contract Account {
    uint256 public balance;
    uint256 public constant MAX_UINT = 2 ** 256 - 1;

    function deposit(uint256 _amount) public {
        uint256 oldBalance = balance;
        uint256 newBalance = balance + _amount;
        require(newBalance >= oldBalance, "Overflow");
        balance = newBalance;
        assert(balance >= oldBalance);
    }

    function withdraw(uint256 _amount) public {
        uint256 oldBalance = balance;
        require(balance >= _amount, "Underflow");
        if (balance < _amount) {
            revert("Underflow");
        }
        balance -= _amount;
        assert(balance <= oldBalance);
    }
}
