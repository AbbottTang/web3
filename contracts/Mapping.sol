// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/**
 * 使用语法映射（keyType => valueType）创建映射。
keyType可以是任何内置值类型、字节、字符串或任何合约。
valueType可以是任何类型，包括另一个映射或数组。
映射是不可迭代的。
 */
contract Mapping {
    //Mapping from address to uint
    mapping(address => uint256) public myMap;

    function get(address _addr) public view returns (uint256) {
        //Mapping 总是返回一个值
        //如果value获取不到，则返回默认值
        return myMap[_addr];
    }

    function set(address _addr, uint256 _i) public {
        //更新地址值
        myMap[_addr] = _i;
    }

    function remove(address _addr) public {
        //重置为默认值
        delete myMap[_addr];
    }
}

contract NestedMapping {
    //嵌套mapping（maping从一个地址到另一个mapping）
    mapping(address => mapping(uint256 => bool)) public nested;

    function get(address _addr1, uint256 _i) public view returns (bool) {
        //你可以从嵌套映射中获取值
        //即使没有初始化
        return nested[_addr1][_i];
    }

    function set(address _addr1, uint256 _i, bool _boo) public {
        nested[_addr1][_i] = _boo;
    }

    function remove(address _addr1, uint256 _i) public {
        delete nested[_addr1][_i];
    }
}
