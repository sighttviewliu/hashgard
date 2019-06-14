# hashgardcli deposit

## 描述
用户发行存款盒子

## 用法

```shell
hashgardcli deposit [subcommand]
```

打印子命令和参数

```shell
hashgardcli lock --help
```
## 子命令
| 名称                        | 功能     |
| -------------------------- | ------------ |
| [params](params.md)        | 参数及费用    |
| [create](create.md)        | 创建存款盒子    |
| [describe](describe.md)   |  对存款进行描述 |
| [inject](inject.md)   |  进行存款 |
| [cancel](cancel.md)   |  取消存款 |
| [interest-inject](interest-inject.md)   |  对盒子注入利息 |
| [interest-cancel](interest-cancel.md)   |  对盒子取回注入利息|
| [query](query.md)  |  查询指定 id 存款盒子 |
| [list](list.md)   | 查询所有存款列表  |
| [search](search.md)  | 按存款盒子名字进行查询  |


## 相关命令
| 名称                        | 功能     |
| -------------------------- | ------------ |
| [bank withdraw](../bank/withdraw.md) | 到期后进行提取 token   |
| [bank box](../bank/box.md)        | 查询指定 box id 信息|
| [bank issue](../bank/issue.md)        | 查询指定发行 token 信息|
