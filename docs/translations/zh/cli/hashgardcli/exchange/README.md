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
| [create-order](create-order.md)  | 创建卖单 |
| [take-order](take-order.md)  | 创建买单 |
| [withdrawal-order](withdrawal-order.md)  | 取消挂单 |
| [query-frozen](query-frozen.md)  | 查询指定地址冻结的资金|
| [query-order](query-order.md)  | 查询指定订单 |
| [query-orders](query-orders.md)  | 查询所有交易订单 |
