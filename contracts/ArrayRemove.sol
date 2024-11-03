// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

contract ArrayRemoveByShifting {
    // [1, 2, 3] -- remove(1) --> [1, 3, 3] --> [1, 3]
    // [1, 2, 3, 4, 5, 6] -- remove(2) --> [1, 2, 4, 5, 6, 6] --> [1, 2, 4, 5, 6]
    // [1, 2, 3, 4, 5, 6] -- remove(0) --> [2, 3, 4, 5, 6, 6] --> [2, 3, 4, 5, 6]
    // [1] -- remove(0) --> [1] --> []
    uint256[] public arr;

    //通过从右向左移动元素来移除数组元素
    function remove(uint256 _index) public {
        require(_index < arr.length, "index out of bound");
        //从数组中移除最后一个元素
        //这将使数组长度减少1
        //下标位替换为后一位，整体向前一位；再删除最后一位
        for (uint256 i = _index; i < arr.length - 1; i++) {
            arr[i] = arr[i + 1];
        }
        arr.pop();
    }

    function test() external {
        arr = [1, 2, 3, 4, 5];
        remove(2);
        //【1，2，4，5】
        assert(arr[0] == 1);
        assert(arr[1] == 2);
        assert(arr[2] == 4);
        assert(arr[3] == 5);
        assert(arr.length == 4);

        arr = [1];
        remove(0);
        assert(arr.length == 0);
    }
}

//通过将最后一个元素复制到要删除的位置来删除数组元素
contract ArrayReplaceFromEnd {
    uint256[] public arr;

    /**
     * //删除一个元素会在数组中创建一个空白。
//保持数组紧凑的一个技巧是
//将最后一个元素移到要删除的位置
     */
    function remove(uint256 index) public {
        //将最后一个元素移动到要删除的位置
        arr[index] = arr[arr.length - 1];
        //移除最后一个元素
        arr.pop;
    }

    function test() public {
        arr = [1, 2, 3, 4];
        remove(1);
        //[1,4,3]
        assert(arr.length == 3);
        assert(arr[0] == 1);
        assert(arr[1] == 4);
        assert(arr[2] == 3);

        remove(2);
        //[1,4]
        assert(arr.length == 2);
        assert(arr[0] == 1);
        assert(arr[1] == 4);
    }
}
