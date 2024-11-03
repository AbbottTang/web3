// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

contract Primitives {
    bool public boo = true;
    /**
     * ‌uint‌：32位无符号整数，范围从0到4294967295，
     * 用于存储非负整数，避免整数溢出问题，
     * 在某些操作中比有符号整数更快。
     * different sizes are available
        uint8   ranges from 0 to 2 ** 8 - 1
        uint16  ranges from 0 to 2 ** 16 - 1
        ...
        uint256 ranges from 0 to 2 ** 256 - 1
     */
    uint8 public u8 = 1;
    uint256 public u256 = 456;
    uint256 public u = 123; //uint 是 uint256 的别名

    /**
     * ‌int‌：32位有符号整数，
     * 范围从-2147483648到2147483647，用于存储负数和正数
     *   int256 ranges from -2 ** 255 to 2 ** 255 - 1
    int128 ranges from -2 ** 127 to 2 ** 127 - 1
     */
    int8 public i8 = -1;
    int256 public i256 = 456;
    int256 public i = -123; //Int与int256相同

    //int的最小值与最大值
    int256 public minInt = type(int256).min;
    int256 public maxInt = type(int256).max;

    address public addr = 0xCA35b7d915458EF540aDe6068dFe2F44E8fa733c;

    /**
     * 在Solidity中，数据类型byte表示一个字节序列。
        Solidity提供了两种类型的字节类型：
        -固定大小的字节数组
        -动态大小的字节数组。
        Solidity中的术语bytes表示一个动态的字节数组。
        它是byte[]的简写。
     */
    bytes1 a = 0xb5; //  [10110101]
    bytes1 b = 0x56; //  [01010110]

    //默认值
    //未赋值变量有默认值
    bool public defaultBoo; //false
    uint256 public defaulyUint; //0
    int256 public defaultInnt; //0
    address public defaultAddr; //0x0000000000000000000000000000000000000000
}
