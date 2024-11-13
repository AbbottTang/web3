// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

/**
 * 多签钱包的功能: 合约有多个 owner，一笔交易发出后，需要多个 owner 确认，确认数达到最低要求数之后，才可以真正的执行。
 * 部署时候传入地址参数和需要的签名数
多个 owner 地址
发起交易的最低签名数
有接受 ETH 主币的方法，
除了存款外，其他所有方法都需要 owner 地址才可以触发
发送前需要检测是否获得了足够的签名数
使用发出的交易数量值作为签名的凭据 ID（类似上么）
每次修改状态变量都需要抛出事件
允许批准的交易，在没有真正执行前取消。
足够数量的 approve 后，才允许真正执行。
 */
contract MultiSigWallet {
    //状态变量
    address[] public owners;
    mapping(address => bool) public isOwner;
    uint256 public required;
    struct Transaction {
        address to;
        uint256 value;
        bytes data;
        bool exected;
    }
    Transaction[] public transactions;
    mapping(uint256 => mapping(address => bool)) public approved;
    //事件
    event Deposit(address indexed sender, uint256 amount);
    event Submit(uint256 indexed txId);
    event Approved(address indexed owner, uint256 indexed txId);
    event Revoke(address indexed owner, uint256 indexed txId);
    event Execute(uint256 indexed txId);

    //receive
    receive() external payable {
        emit Deposit(msg.sender, msg.value);
    }

    //定义一个自定义修饰符，用于检查调用者是否是合约的所有者
    modifier onlyOwner() {
        require(isOwner[msg.sender], "Not the owner");
        _; // 这是被修饰函数的代码执行的位置
    }
    //定义一个自定义修饰符，用于要求转账申请下标(id)小于转账列表长度
    //转账存在判断
    modifier txExists(uint256 _txId) {
        require(_txId < transactions.length, "tx doesn't exist");
        _;
    }
    //未被通过
    modifier notApproved(uint256 _txId) {
        require(!approved[_txId][msg.sender], "tx already approved");
        _;
    }
    //未被执行
    modifier notExecuted(uint256 _txId) {
        require(!transactions[_txId].exected, "tx is exected");
        _;
    }

    //构造函数
    constructor(address[] memory _owners, uint256 _required) {
        require(_owners.length > 0, "owner required");
        require(
            _required < owners.length && _required > 0,
            "invalid required number of owners"
        );
        for (uint256 index = 0; index < _owners.length; index++) {
            address owner = _owners[index];
            //它检查 owner 这个变量是否不等于 address(0)。在Solidity中，address(0) 是一个特殊的地址，通常被认为是“零地址”或“无效地址”。它常用于表示一个尚未初始化或不存在的地址。
            require(owner != address(0), "invalid owner");
            //地址重复检查
            require(!isOwner[owner], "owner is not unique");
            //初始赋值为true
            isOwner[owner] = true;
            //放入拥有者
            owners.push(owner);
        }
        //赋值同意人数
        required = _required;
    }

    //获取钱包余额
    function getBalance() external view returns (uint256) {
        return address(this).balance;
    }

    //提交一个转账任务
    function submit(
        address _to,
        uint256 _value,
        bytes calldata _data
    ) external onlyOwner returns (uint256) {
        //钱包拥有者才可调用
        transactions.push(
            Transaction({to: _to, value: _value, data: _data, exected: false})
        );
        emit Submit(transactions.length - 1);
        //返回提交转账任务下标
        return transactions.length - 1;
    }

    //同意某笔转账
    function approv(
        uint256 _txId
    ) external onlyOwner txExists(_txId) notApproved(_txId) notExecuted(_txId) {
        approved[_txId][msg.sender] = true;
        emit Approved(msg.sender, _txId);
    }

    function execute(
        uint256 _txId
    ) external onlyOwner txExists(_txId) notExecuted(_txId) {
        require(getApprovalCount(_txId) >= required, "approvals<required");
        Transaction storage transaction = transactions[_txId];
        transaction.exected = true;
        //通过to、value和data属性来执行一个低级别的函数调用交易发送（包括价值和数据）
        (bool success, ) = transaction.to.call{value: transaction.value}(
            transaction.data
        );
        require(success, "tx failed");
        emit Execute(_txId);
    }

    //获取赞成数量
    function getApprovalCount(
        uint256 _txId
    ) public view returns (uint256 count) {
        for (uint256 index = 0; index < owners.length; index++) {
            if (approved[_txId][owners[index]]) {
                count += 1;
            }
        }
    }

    //不同意
    function revoke(
        uint256 _txId
    ) external onlyOwner txExists(_txId) notExecuted(_txId) {
        require(approved[_txId][msg.sender], "tx not approved");
        approved[_txId][msg.sender] = false;
        emit Revoke(msg.sender, _txId);
    }
}
