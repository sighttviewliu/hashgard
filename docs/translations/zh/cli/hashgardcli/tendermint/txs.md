# hashgardcli tendermint txs

## 描述

查询与指定标签匹配的所有交易

## 用法

```shell
hashgardcli tendermint txs [flags]
```

## Flags

| 名称, 缩写 | 类型   | 必需 | 默认值 | 描述                      |
| ---------- | ------ | ---- | ------ | ------------------------- |
| --limit    | int    | 否   | 32     | 返回结果每页数量          |
| --page     | int    | 否   | 1      | 指定返回第几页            |
| --tags     | string | 是   |        | 需要匹配的 tag:value 列表 |

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

```shell
 hashgardcli tendermint txs --tags '<tag1>:<value1>&<tag2>:<value2>' \
 --page 1 --limit 30 --trust-node
```
