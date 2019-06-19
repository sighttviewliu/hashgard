# hashgardcli distribution

## 介绍

这里主要介绍 distribution 模块提供的命令行接口

## 用法

```shell
hashgardcli distribution [subcommand]
```

打印子命令和参数

```shell
hashgardcli distribution --help
```

## 子命令

| 名称                            | 功能    |
| --------------------------------| ------------------------|
| [params](params.md)  | 查询分配参数信息 |
| [outstanding-rewards](outstanding-rewards.md)  | 查询优先（未提取）奖励 |
| [commission](commission.md)  | 查询验证人和委托人的佣金奖励 |
| [slashes](slashes.md)  | 查询给定块范围的验证人的所有处罚 |
| [rewards](rewards.md)  | 查询委托人获得的所有奖励 |
| [set-withdraw-addr](set-withdraw-address.md)  | 设置收益取回地址 |
| [withdraw-rewards](withdraw-rewards.md) | 发起取回收益的交易 |
| [withdraw-all-rewards](withdraw-all-rewards.md) | 取回所有的收益 |
| [community-pool](community-pool.md)  | 查询 community-pool 奖励
