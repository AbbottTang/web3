// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

contract EtherUnits {
    /**
     * 交易用以太币支付。
类似于1美元等于100美分，1以太等于10^18威。
     */
    /*
    由于以太坊上很多交易数额都比较小，因此以太币有几种面额，可以作为较小的记账单位。 在这些面额中，Wei 与 Gwei 特别重要。

Wei 是最小的以太币面额，因此在以太坊黄皮书(opens in a new tab)

等众多技术实现中，都以 Wei 为单位进行计算。

Gwei（giga-wei 的缩写），常用于描述以太坊上的燃料费用。
1wei = 1e^-9 gwei
1wei = 1e^-18 Ether

1gwei = 10^9wei
1gwei = 1e^-9Ether

1Ether = 10^18wei
1Ether = 10^9gwei

 */
    uint public oneWei = 1 wei;

    //1wei等于1
    bool public isOnewei = (oneWei == 1);

    uint256 public oneGwei = 1e-9 gwei;
    //1gwei等于10^ 9wei
    bool public isOneGwei = (oneGwei == 1e9 );

    uint256 public oneEther = 1 ether;
    // 1 ether 等于10^18 wei
    bool public isOneEther = (oneEther == 1e18);

}