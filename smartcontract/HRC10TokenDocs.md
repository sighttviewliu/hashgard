# HRC10 Token 发行文档

HRC10 Token Contract 分为两个部分：

1. standard contract
2. premium contract

## standard contract

#### 查询方法：

1. name
2. symbol
3. owner
4. decimals
5. totalSupply
6. balanceOf
7. allowance
8. description
9. contractKeys

#### 交易方法：

1. init
2. transfer
3. approve
4. increaseApproval
5. decreaseApproval
6. transferFrom
7. describe


## premium contract

高级合约包含标准合约中的所有方法，同时拥有额外的高级操作方法：

#### 查询方法：

1. mintDisabled
2. burnDisabled
3. freezeDisabled

#### 交易方法：

1. disableMint
2. disableBurn
3. disableFreeze
4. mint
5. burn
6. freezeIn
7. freezeOut
8. changeOwner