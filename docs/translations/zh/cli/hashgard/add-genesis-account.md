# hashgard add-genesis-account

## 描述

将 genesis 帐户添加到 /path/to/.hashgard/config/genesis.json 文件中

## 用法

```shell
hashgard add-genesis-account [address_or_key_name] [coin][,[coin]] [flags]
```

## 参数

| 名称，缩写            | 类型   | 默认值 | 描述                       | 必需 |
| --------------------- | ------ | ------ | -------------------------- | -------- |
| [address_or_key_name] | string |        | 被添加的账户名或地址       | `是`     |
| [coin]                | string |        | 为该账户分配的 coin 及数量 | `是`     |

## Flags

| 名称，缩写           | 类型   | 默认值         | 描述                             | 必需 |
| -------------------- | ------ | -------------- | -------------------------------- | -------- |
| -h, --help           |        |                | add-genesis-account 模块帮助文档 | 否       |
| --home-client        | string | ~/.hashgardcli | 客户端根目录                     | 否       |
| --vesting-amount     | string |                | 冻结数量                         | 否       |
| --vesting-end-time   | int    |                | 冻结结束时间 (unix epoch)        | 否       |
| --vesting-start-time | int    |                | 冻结起始时间 (unix epoch)        | 否       |
| --home               | string | ~/.hashgard    | 配置和数据的目录                 | 否       |
| --trace              | bool   |                | 在出错时打印完整的调用栈         | 否       |

## 例子

```bash
hashgardcli keys add root
hashgard add-genesis-account root 100000000gard
```
