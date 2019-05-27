# hashgardcli gov tally

## 描述

查询指定提案的投票统计

## 用法

```shell
 hashgardcli gov tally [proposal-id] [flags]
```

## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)


## 例子

### 查询投票统计

```shell
hashgardcli gov tally 1 --trust-node
```

可以查询指定提案每个投票选项的投票统计。

```txt
{
  "yes": "89.1000000000",
  "abstain": "0.0000000000",
  "no": "0.0000000000",
  "no_with_veto": "0.0000000000"
}
```
