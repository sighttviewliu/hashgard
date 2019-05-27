# hashgard tendermint

## 描述

查询 tendermint 节点状态的子命令

## 用法

```shell
hashgard tendermint [subcommand] [flags]
```

## 子命令

| 名称             | 描述                                 |
| ---------------- | ------------------------------------ |
| --show-node-id   | 查询节点 ID                          |
| --show-validator | 查询 tendermint 节点的验证人 pubkey  |
| --show-address   | 查询 tendermint 节点的验证人共识地址 |

## Flags

| 名称，缩写 | 类型   | 默认值      | 描述                     | 必需 |
| ---------- | ------ | ----------- | ------------------------ | -------- |
| -h, --help |        |             | testnet 模块帮助         | 否       |
| --home     | string | ~/.hashgard | 配置和数据的目录         | 否       |
| --trace    | bool   |             | 在出错时打印完整的调用栈 | 否       |

## 例子

```shell
hashgard tendermint show-node-id
hashgard tendermint show-validator
hashgard tendermint show-address
```
