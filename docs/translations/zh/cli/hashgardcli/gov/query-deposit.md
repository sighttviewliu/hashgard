# hashgardcli gov query-deposit

## 描述

查询保证金的充值明细

## 用法

```
 hashgardcli gov query-deposit [proposal-id] [depositer-addr] [flags]
```

## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

### 查询充值保证金

```shell
hashgardcli gov query-deposit 1 gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet --trust-node -o=json --indent
```

通过指定提案、指定存款人查询保证金充值详情，得到结果如下：

```txt
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
```
