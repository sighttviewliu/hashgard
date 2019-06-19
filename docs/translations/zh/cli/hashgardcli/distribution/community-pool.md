# hashgardcli distribution community-pool

## 描述

查询 community-pool

## 用法

```shell
hashgardcli distribution community-pool [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

查询参数信息

```shell
hashgardcli distribution community-pool -o=json --trust-node
```

运行成功以后，返回的结果如下：

```txt
[
 {
  "denom": "gard",
  "amount": "1000.337966901187138531"
 }
]
```
