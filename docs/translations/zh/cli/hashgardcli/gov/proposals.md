# hashgardcli gov proposals

## 描述

通过可选项过滤查询满足条件的提案

## 用法

```shell
hashgardcli gov proposals [flags]
```
## Flags

| 名称       | 类型                  | 必需                 | 默认值                      | 描述          |
| --------------- | -------------- | --------------------- | --------------------------- | ------------------- |
| --depositor     | string | 否 || （可选）按存款人过滤                         |
| --limit         | string | 否 || （可选）限制最新[数量]提案。 默认为所有提案                  |
| --status        | string | 否 || （可选）按提案状态过滤提案                  |
| --voter         | string | 否 || （可选）按投票人过滤                        |

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

### 查询提案

```shell
hashgardcli gov proposals --trust-node
```

默认查询所有的提案。

```txt
ID - (Status) [Type] Title
1 - (DepositPeriod) [Text] Test Proposal
2 - (DepositPeriod) [Text] Test Proposal
3 - (DepositPeriod) [Text] Test Proposal
4 - (VotingPeriod) [Text] Test Proposal
```

当然这里可以查询指定条件的提案，例如：

```shell
gov proposals --chain-id=hashgard --depositor=gard4q5rf9sl2dqd2uxrxykafxq3nu3lj2fp9l7pgd
```

可以得到存款人是 gard4q5rf9sl2dqd2uxrxykafxq3nu3lj2fp9l7pgd 地址的提案。
```txt
  2 - new proposal
```
