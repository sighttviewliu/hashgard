# hashgardcli issue unfreeze

## 描述
解冻用户的转入和转出功能

## 用法
```shell
 hashgardcli issue unfreeze [unfreeze-type] [issue-id][address] --from
```
### unfreeze-type

| 名称   | 描述                 |
| ------ | -------------------- |
| in     | 该账号本通证转入功能 |
| out    | 该账号本通证转出功能 |
| In-out | 该账号转入和转出功能 |


## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

### 解冻某地址的转入功能
```shell
hashgardcli issue unfreeze in coin174876e800 gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n --from
```
输入正确的密码之后，你就解冻该账户的解冻功能
```txt
{
  Height: 2758
  TxHash: C6CE11D458B0F64C164E91CF2FF692A65D1EA9C0B1C2A2B228A7C1699C6423FE
  Data: 0F0E636F696E31373438373665383030
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 16203
  Tags:
    - action = issue_unfreeze
    - category = issue
    - issue-id = coin174876e800
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
    - freeze-type = in
}
```
