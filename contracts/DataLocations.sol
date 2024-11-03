// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/**
 * 数据位置-storage，memory和Calldata
变量声明为storage、memory或calldata，以显式指定数据的位置。
storage变量是一个状态变量（存储在区块链上）
memory变量在内存中，它在函数被调用时存在
Calldata包含函数参数的特殊数据位置
 */
contract DataLocations {
    uint256[] public arr;
    mapping(uint256 => address) map;

    struct MyStruct {
        uint256 foo;
    }

    mapping(uint256 => MyStruct) myStructs;

    function f() public {
        //用状态变量调用_f
        _f(arr, map, myStructs[1]);

        //得到一个struct从mapping
        MyStruct storage myStruct = myStructs[1];

        //在内存中创建一个struct
        MyStruct memory myMenStruct = MyStruct(0);
    }

    function _f(
        uint256[] storage _arr,
        mapping(uint256 => address) storage _map,
        MyStruct storage _myStruct
    ) internal {
        //用storage变量做点啥
    }

    //返回一个内存变量
    function g(uint256[] memory _arr) public returns (uint256[] memory) {}

    function h(uint256[] calldata _arr) external {}
}
