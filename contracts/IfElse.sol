// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/**
 * Solidity支持条件语句if、else if和else。
 */
contract IfElse {
    function foo(uint256 x) public pure returns (uint256) {
        if (x < 10) {
            return 0;
        } else if (x < 20) {
            return 1;
        } else {
            return 2;
        }
    }

    //三元表达式
    function ternary(uint256 _x) public pure returns (uint256) {
        // if(_x<10){
        //     return 1;
        // }
        // return 2;
        // if / else语句的简写方式
        //“？”操作符称为三元操作符
        return _x < 10 ? 1 : 2;
    }
}
