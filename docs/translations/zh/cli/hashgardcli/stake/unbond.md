# hashgardcli stake unbond

## 介绍

从一个验证人解绑委托

## 用法

```shell
hashgardcli stake unbond [validator-addr] [amount] [flags]
```

## 子命令

| 名称           | 类型   | 必需 | 默认值 | 描述            |
| -------------- | ------ | -------- | ------ | ------------------- |
| validator-addr | string |是    |        | 验证人地址          |
| amount         | string    | 是     |        | 要解绑的 stake 数量和种类（单位 gard）|

## Flags

| 名称   | 类型   | 必需 | 默认值 | 描述             |
| ------ | ------ | -------- | ------ | -------------------- |
| --from | string | 是     |    | 委托人的账户名或地址 |

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

```shell
hashgardcli stake unbond \
gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx \
10gard \
--from $you_wallet_name \
```
