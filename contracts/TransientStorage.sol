// SPDX-License-Identifier: MIT
pragma solidity ^0.8.27;

/**存储在暂态存储器中的数据在事务完成后被清除 */

/**
 * //确保EVM版本和虚拟机设置为“Cancun
//Storage -数据存储在区块链上
//Memory -函数调用后清空内存数据
//Transient storage-暂态存储——事务完成后清空数据
 */

interface ITest {
    function val() external view returns (uint256);

    function test() external;
}

contract Callback {
    uint256 public val;

    fallback() external {
        val = ITest(msg.sender).val();
    }

    function test(address target) external {
        ITest(target).test();
    }
}

contract TestStorage {
    uint256 public val;

    function test() public {
        val = 123;
        bytes memory b = "";
        msg.sender.call(b);
    }
}

contract TestTransientStorage {
    bytes32 constant SLOT = 0;

    function test() public {
        assembly {
            tstore(SLOT, 321)
        }
        bytes memory b = "";
        msg.sender.call(b);
    }

    function val() public view returns (uint256 v) {
        assembly {
            v := tload(SLOT)
        }
    }
}

contract ReentrancyGuard {
    bool private locked;
    modifier lock() {
        require(!locked);
        locked = true;
        _;
        locked = false;
    }

    function test() public lock {
        bytes memory b = "";
        msg.sender.call(b);
    }
}

/**“tstore”指令只适用于坎昆兼容的虚拟机（您目前正在编译“paris”）。 */
contract ReentrancyGuardTransient {
    bytes32 constant SLOT = 0;

    modifier lock() {
        assembly {
            if tload(SLOT) {
                revert(0, 0)
            }
            tstore(SLOT, 1)
        }
        _;
        assembly {
            tstore(SLOT, 0)
        }
    }

    // 21887 gas
    function test() external lock {
        // Ignore call error
        bytes memory b = "";
        msg.sender.call(b);
    }
}
