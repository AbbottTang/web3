// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract CrowdFunding {
    //受益人-immutable-类似java中的final-常量不可变，在构造函数中初始化赋值
    address public immutable beneficiary;
    //筹资目标数量
    uint256 public immutable fundingGoal;
    //当前金额
    uint256 public fundingAmount;
    //mapping入资者地址-金额记录
    mapping(address => uint256) public funders;
    //可迭代映射
    mapping(address => bool) private funderInserted;
    //address-记录捐赠者地址
    address[] public fundersKey;
    //不用自销毁方法，使用变量来控制
    //状态
    bool public AVAILABLED = true;

    //构造函数-部署的时候，写入受益人+筹资目标数量
    constructor(address beneficiary_, uint256 goal_) {
        beneficiary = beneficiary_;
        fundingGoal = goal_;
    }

    /**
     * 资助方法
     * 状态可用时可以调用入资捐助
     * 合约关闭后就不能再操作了
     */
    function contribute() external payable {
        require(AVAILABLED, "CrowdFunding is close");
        //检查捐赠金额是否会超过目标金额 当前金额+捐赠金额
        uint256 potentialFundingAmount = fundingAmount + msg.value;
        //返还金额
        uint256 refundAmount = 0;
        //判断逻辑，大于目标金额，计算剩余所需捐赠金额 = 目标金额 - 当前金额
        //并返还金额= 捐赠金额 -计算剩余所需捐赠金额
        //小于目标金额 就等于捐赠金额
        if (potentialFundingAmount > fundingGoal) {
            refundAmount = potentialFundingAmount - fundingGoal;
            funders[msg.sender] += (msg.value - refundAmount);
            fundingAmount += (msg.value - refundAmount);
        } else {
            funders[msg.sender] += msg.value;
            fundingAmount += msg.value;
        }
        //更新捐赠者信息
        if (!funderInserted[msg.sender]) {
            funderInserted[msg.sender] = true;
            fundersKey.push(msg.sender);
        }
        //退还多余的金额
        if (refundAmount > 0) {
            payable(msg.sender).transfer(refundAmount);
        }
    }

    //关闭捐赠
    function close() external returns (bool) {
        //1.检查金额
        if (fundingAmount < fundingGoal) {
            return false;
        }
        uint256 amount = fundingAmount;
        //2.归零融资金额-修改状态
        fundingAmount = 0;
        AVAILABLED = false;
        //3.操作转账
        payable(beneficiary).transfer(amount);
        return true;
    }

    //查看捐赠人数
    function fundersLength() public view returns (uint256) {
        return fundersKey.length;
    }
}
