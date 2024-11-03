// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/**
 * 你支付gas花费* gas价格的以太，其中
gas是一种计算单位
gas消耗是指在一笔交易中使用的总gas量
gas价格是你愿意为每次gas支付多少钱
gas价格高的交易有更高的优先级被包含在区块中。
未用完的gas将退还。
 */

/**
 * gas 上限
 * 你可以花费的gas量有两个上限
gas限额（您愿意为您的交易使用的最大gas量，由您设置）
区块gas限制（一个区块允许的最大gas量，由网络设定）
 */
contract Gas {
    uint256 public i = 0;

    //耗尽你发送的所有gas导致交易失败。
    //状态更改被撤消。
    //gas不退还。
    function forever() public {
        //这里我们运行一个循环，直到所有的气体被消耗
        //事务失败

        while (true) {
            i += 1;
        }
    }
}
