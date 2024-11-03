// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

//数组可以具有编译时固定大小或动态大小。
contract Array {
    //初始化数组的几种方法
    uint256[] public arr;
    uint256[] public arr2 = [1, 2, 3];
    //固定大小的数组，所有元素初始化为0
    uint256[10] public myFixedSizeArr;

    function get(uint256 i) public view returns (uint256) {
        return arr[i];
    }

    // Solidity可以返回整个数组。
    //但是应该避免使用这个函数
    //长度可以无限增长的数组。
    function getArr() public view returns (uint256[] memory) {
        return arr;
    }

    function pop() public {
        //从数组中移除最后一个元素
        //这将使数组长度减少1
        arr.pop();
    }

    function getLength() public view returns (uint256) {
        return arr.length;
    }

    function remove(uint256 index) public {
        //删除不改变数组的长度。
        //将索引处的值重置为默认值；
        //在本例中为0
        delete arr[index];
    }

    function examples() external {
        //在内存中创建数组，只能创建固定大小的数组
        uint256[] memory a = new uint256[](5);
    }
}
