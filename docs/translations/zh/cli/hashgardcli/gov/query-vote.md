# hashgardcli gov query-vote

## 描述

查询指定提案、指定投票者的投票情况

## 用法

```
hashgardcli gov query-vote [proposal-id] [voter-addr] [flags]
```

## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)
 
### 查询投票

```shell
hashgardcli gov query-vote 1 gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet --trust-node -o=json --indent
```

通过指定提案、指定投票者查询投票情况。

```txt
{
  "voter": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
  "proposal_id": "1",
  "option": "Yes"
}

```
