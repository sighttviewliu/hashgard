# hashgardcli stake

## 介绍

Stake 模块提供了一系列查询 staking 状态和发送 staking 交易的命令

## 用法

```shell
hashgardcli stake [subcommand]
```

## 子命令

| 名称                                                     | 描述       |
| ----------- | --------------------------------- |
| [validator](validator.md)                                   | 查询某个验证人        |
| [validators](validators.md)                                 | 查询所有的验证人          |
| [delegation](delegation.md)                                 | 基于委托人地址和验证人地址的委托查询 |
| [delegations](delegations.md)                               | 基于委托人地址的所有委托查询   |
| [delegations-to](delegations-to.md)                         | 查询对一个验证人进行的所有委托    |
| [unbonding-delegation](unbonding-delegation.md)             | 基于委托人地址和验证人地址的 unbonding-delegation 记录查询|
| [unbonding-delegations](unbonding-delegations.md)           | 基于委托人地址的所有 unbonding-delegation 记录查询  |
| [unbonding-delegations-from](unbonding-delegations-from.md) | 基于验证人地址的所有 unbonding-delegation 记录查询  |
| [redelegation](redelegation.md)                             | 基于委托人地址、原验证人地址、目标验证人地址查询重新委托记录 |
| [redelegations](redelegations.md)                           | 基于委托人地址的所有重新委托记录查询     |
| [redelegations-from](redelegations-from.md)                 | 基于某一验证人地址的所有重新委托查询      |
| [pool](pool.md)                                             | 查询最新的权益池       |
| [params](params.md)                                         | 查询最新的权益参数信息         |
| [create-validator](create-validator.md)                     | 以自委托的方式创建一个新的验证人         |
| [edit-validator](edit-validator.md)                         | 编辑已存在的验证人信息           |
| [delegate](delegate.md)                                     | 委托一定代币到某个验证人           |
| [unbond](unbond.md)                                         | 从指定的验证人解绑一定的 token  |
| [redelegate](redelegate.md)                                 | 从一个验证人重新委托一定的 token 到另一个验证人 |
