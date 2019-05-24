# hashgardcli gov query-deposits

## 描述

查询指定提案的保证金详细情况

## 用法

```
hashgardcli gov query-deposits [proposal-id] [flags]
```
## Flags

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

### 查询所有保证金

```shell
hashgardcli gov query-deposits 1 --trust-node -o=json --indent

```

你可以查询到指定提案的所有保证金代币，包括每个存款人的充值详情。

```txt
[
  {
    "depositor": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
    "proposal_id": "1",
    "amount": [
      {
        "denom": "gard",
        "amount": "50"
      }
    ]
  }
]

```
