# hashgardcli issue

## 描述

基于 Hashgard 公链发行自己的 HRC10 代币，我们提供增发、销毁、冻结、转移所有者，代币查询等链上原生服务。

```
注：我们不限制你在公链账号里持有多少GARD币来发行自己的代币，只需足够支付每步操作的gas费用即可。
```

## 用法

```shell
hashgardcli issue [command]
```

打印子命令和参数

```
hashgardcli issue --help
```

## 相关命令

| 名称                                       | 描述                                 |
| ------------------------------------------- | ------------------------------------ |
| [create](create.md)                         | 发行一个新的代币。                   |
| [describe](describe.md)                     | 设置代币的描述信息。                 |
| [transfer-ownership](transfer-ownership.md) | 转移所有者。                         |
| [freeze](freeze.md)                         | 冻结用户的转入和转出。               |
| [unfreeze](unfreeze.md)                     | 解冻用户转账功能                 |
| [mint](mint.md)                             | 增发                                 |
| [burn](burn.md)                             | 持有者销毁自身持有的代币             |
| [burn-from](burn-from.md)                   | Owner销毁任意持币者的代币            |
| [disable](disable.md)                       | 通证高级特性开关。                   |
| [query-issue](query-issue.md)               | 查询指定issue-id值的发行的币的信息。 |
| [search](search.md)                         | 根据代币符号搜索                     |
