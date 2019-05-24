# hashgardcli slashing signing-info

## 描述

使用验证人的共识公钥查询验证人的签名信息

## 用法

```
hashgardcli slashing signing-info [validator-conspub] [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

查询验证人的签名信息

```
hashgardcli slashing signing-info \
gardvalconspub1zcjduepqgsmuj0qallsw79hjj9qztcke6hj3ujdcpjv249uny9fvzp4eulms0tqvgs \
--trust-node
```

运行成功以后，返回的结果如下：

```txt
Start Height:          0
Index Offset:          80
Jailed Until:          1970-01-01 00:00:00 +0000 UTC
Tombstoned:            false
Missed Blocks Counter: 0
```
