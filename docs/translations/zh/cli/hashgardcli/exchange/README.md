# hashgardcli exchange

## 介绍

这里主要介绍 exchange 原子交易模块提供的命令行接口

## 用法

```shell
hashgardcli exchange [subcommand]
```

打印子命令和参数

```shell
hashgardcli exchange --help
```

## 子命令

| 名称                            | 功能    |
| --------------------------------| ------------------------|
| [create](create.md)  | 创建卖单 |
| [take](take.md)  | 创建买单 |
| [cancel](cancel.md)  | 取消挂单 |
| [query-frozen](query-frozen.md)  | 查询指定地址冻结的资金|
| [query](query.md)  | 查询指定订单 |
| [list](list.md)  | 查询所有交易订单 |
