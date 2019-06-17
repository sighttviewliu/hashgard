# hashgardcli stake params

## 描述

查询最新的权益参数信息

## 用法

```shell
hashgardcli stake params [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

查询最新的权益参数信息

```shell
hashgardcli stake params --trust-node
```

运行成功以后，返回的结果如下：

```txt
Params:
  Unbonding Time:    10m0s
  Max Validators:    100
  Max Entries:       0
  Bonded Coin Denom: gard
```
