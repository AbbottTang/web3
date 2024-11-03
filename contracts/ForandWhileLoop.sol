// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/**
 * Solidity支持for、while和do while循环。
不要编写无界的循环，因为这会达到gas限制，导致事务失败。
由于上述原因，while和do while循环很少使用。
 */

contract Loop {
    function loop() public {
        //for loop
        for (uint256 i = 0; i < 10; i++) {
            if (i == 3) {
                //continue跳转到下一个迭代
                continue;
            }
            if (i == 5) {
                //带中断的退出循环
                break;
            }
        }
        // while loop
        uint256 j;
        while (j < 10) {
            j++;
        }
    }
}
