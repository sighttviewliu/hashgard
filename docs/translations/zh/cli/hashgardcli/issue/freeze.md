# hashgardcli issue freeze

## 描述
在没有关闭通证冻结的前提下，owenr 冻结用户的转入转出功能。
## 用法
```shell
 hashgardcli issue freeze [freeze-type] [issue-id][acc-address][end-time] --from
```
### freeze-type

| 名称 | 描述                 |
| ------ | -------------------- |
| in     | 该账号本通证转入功能 |
| out    | 该账号本通证转出功能 |
| In-out | 该账号转入和转出功能 |



## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子
### 冻结某账户转入功能
```shell
hashgardcli issue freeze in coin174876e800 gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n\ 253382641454 --from
```
输入正确的密码后，你就冻结了该地址的该通证的转入功能。
```txt
{
Height: 2570
  TxHash: DA8EEDB42B3177E281B462A88AB77D04E398286A4215D5BA0898ABA98F0270AA
  Data: 0F0E636F696E31373438373665383030
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 16459
  Tags:
    - action = issue_freeze
    - category = issue
    - issue-id = coin174876e800
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
    - freeze-type = in

}
```
