# hashgardcli box

## 描述

基于 Hashgard 公链发行的各种原生盒子，来进行各种形态的金融需求。

- 锁仓盒子
- 存款盒子
- 远期支付盒子

## 用法

```shell
hashgardcli box [command]
```

打印子命令和参数

```shell
hashgardcli box --help
```

## 相关命令

| 名称                                   | 描述                                   |
| ------------------------------------------- | -------------------------------------- |
| [create-lock](create-lock.md)               | 创建一个锁定存款的盒子。               |
| [create-deposit](create-deposit.md)         | 创建一个存款盒子。                  |
| [create-future](create-future.md)           | 创建远期支付盒子。                     |
| [interest-injection](interest-injection.md) | 为盒子注入利息。                   |
| [interest-fetch](interest-fetch.md)         | 取回盒子中的利息。                 |
| [deposit-to](deposit-to.md)                 | 用户对盒子进行存款。                   |
| [deposit-fetch](deposit-fetch.md)           | 用户在存款吸纳期对已存入存款进行取回。 |
| [describe](describe.md)                     | 对于盒子的描述与介绍。                 |
| [list-box](list-box.md)                     | 查询不同类型盒子列表                   |
| [query-box](query-box.md)                   | 指定查询盒子信息内容                   |
