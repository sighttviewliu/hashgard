# hashgardcli distribution outstanding-rewards

## 描述

查询优先（未提取）奖励

## 用法

```shell
hashgardcli distribution outstanding-rewards [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

查询参数信息

```shell
hashgardcli distribution outstanding-rewards -o=json --trust-node
```

运行成功以后，返回的结果如下：

```txt
[
 {
  "denom": "gard",
  "amount": "14.716656508693175779"
 }
]
```
