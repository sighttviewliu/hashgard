# hashgardcli tendermint tendermint-validator-set

## 描述

获取指定高度的所有验证人集合

## 用法

```shell
  hashgardcli tendermint tendermint-validator-set [height] [flags]
```

## Flags

| 名称, 速记    | 默认值                | 描述                                 | 必需                               |
| ------------- | --------------------- | ------------------------------------ | ---------------------------------- |
| --chain-id    | 无                    | [string] tendermint 节点的链 ID      | 否，当`--trust-node=false`时，为是 |
| --node string | tcp://localhost:26657 | 要连接的节点                         | 否                                 |
| --trust-node  | true                  | 信任连接的完整节点，关闭响应结果校验 | 否                                 |

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

```shell
hashgardcli tendermint tendermint-validator-set 114360 --trust-node
```

返回结果如下：

```txt
block height: 123

  Address:          gardvalcons13ja77lpt0deamvuwz5eugy9kwkutxukjwjwwf3
  Pubkey:           gardvalconspub1zcjduepqgsmuj0qallsw79hjj9qztcke6hj3ujdcpjv249uny9fvzp4eulms0tqvgs
  ProposerPriority: 0
  VotingPower:      1000

```
