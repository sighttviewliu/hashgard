# hashgardcli gov

## 描述

该模块提供如下所述的基本功能：

* 讨论方案的链上提案
* 参数变化的链上治理
* 软件升级的链上治理

## 用法

```shell
hashgardcli gov [command]
```

打印子命令和参数

```
hashgardcli gov --help
```
## 相关命令

| 名称                                  | 描述                                                             |
| ------------------------------------- | --------------------------------------------------------------- |
| [proposal](proposal.md)   | 查询单个提案的详细信息                                             |
| [proposals](proposals.md) | 通过可选过滤器查询提案                                             |
| [query-vote](query-vote.md)           | 查询投票信息                                                      |
| [votes](votes.md)         | 查询提案的投票情况                                                 |
| [query-deposit](query-deposit.md)     | 查询保证金详情                                                    |
| [deposits](deposits.md)   | 查询提案的保证金                                                  |
| [tally](tally.md)         | 查询提案投票的统计                                                 |
| [param](param.md)       | 查询参数提案的配置                                                 |                                            |
| [submit-proposal](submit-proposal.md) | 创建新密钥，或者通过助记词导入恢复                                   |
| [deposit](deposit.md)                 | 充值保证金代币以激活提案                                            |
| [vote](vote.md)                       | 为有效的提案投票，选项：Yes/No/NoWithVeto/Abstain                   |

## 补充描述

* 任何用户都可以存入一些令牌来发起提案。存款达到一定值min_deposit后，进入投票期，否则将保留存款期。其他人可以在存款期内存入提案。一旦存款总额达到min_deposit，输入投票期。但是，如果冻结时间超过存款期间的max_deposit_period，则提案将被关闭。
* 进入投票期的提案只能由验证人和委托人投票。未投票的代理人的投票将与其验证人的投票相同，并且投票的代理人的投票将保留。到达“voting_period”时，票数将被计算在内。
