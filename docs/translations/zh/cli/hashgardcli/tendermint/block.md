# hashgardcli tendermint block

## 描述

获取指定高度的块的数据

## 用法

```shell
  hashgardcli tendermint block [height] [flags]
```

## Flags

| 名称, 速记   | 默认值                | 描述                                                | 必需                               |
| ------------ | --------------------- | --------------------------------------------------- | ---------------------------------- |
| --chain-id   | 无                    | [string] tendermint 节点的链 ID                     | 否，当`--trust-node=false`时，为是 |
| --node       | tcp://localhost:26657 | [string] 要连接的节点                               | 否                                 |
| --trust-node | true                  | 信任连接的完整节点，值为 false 时对响应结果进行校验 | 否                                 |

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

```shell
hashgardcli tendermint block 114263  --trust-node
```
