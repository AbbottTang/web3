// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

contract Enum {
    //表示shipping 状态的枚举
    enum Status {
        Pending,
        Shipped,
        Accepted,
        Rejected,
        Canceled
    }
    //默认值是列表中的第一个元素
    //类型的定义，在本例中为“Pending”
    Status public status;

    // Returns uint
    // Pending  - 0
    // Shipped  - 1
    // Accepted - 2
    // Rejected - 3
    // Canceled - 4
    function get() public view returns (Status) {
        return status;
    }

    //通过向input传递int来更新状态
    function set(Status _status) public {
        status = _status;
    }

    //更新status到特定的枚举
    function cancel() public {
        status = Status.Canceled;
    }

    //Delete将枚举重置为它的第一个值0
    function reset() public {
        delete status;
    }
}
