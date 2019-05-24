# hashgardcli slashing params

## 描述

查询当前的惩罚参数信息

## 用法

```
hashgardcli slashing params [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

查询当前的惩罚参数信息

```
hashgardcli slashing params --trust-node
```

运行成功以后，返回的结果如下：

```txt
Slashing Params:
  MaxEvidenceAge:          2m0s
  SignedBlocksWindow:      100
  MinSignedPerWindow:      0.500000000000000000
  DowntimeJailDuration:    10m0s
  SlashFractionDoubleSign: 0.050000000000000000
  SlashFractionDowntime:   0.010000000000000000
```
