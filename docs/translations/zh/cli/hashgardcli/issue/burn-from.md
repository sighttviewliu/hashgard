# hashgardcli issue burn-from

## 描述
某个代币的Owner在没有关闭持币者自己可以销毁自己持有该代币前提下，持币者对自己持有的该代币进行销毁。
## 用法
```
 hashgardcli issue burn-from [issue-id] [acc-address][amount] --from
```
## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子
### 销毁代币
```shell
hashgardcli issue burn-from coin174876e801 gard1lgs73mwr56u2f4z4yz36w8mf7ym50e7myrqn65 88 --from
```
输入正确的密码之后，你就销毁了其他人账户里的代币。
```txt
{
  Height: 2991
  TxHash: 09E2591037100326AC7730E3E8C53103D72C1C38BF4DF82600338DD6DF38CC4B
  Data: 0F0E636F696E31373438373665383032
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 29892
  Tags: 
    - action = issue_burn_from
    - category = issue
    - issue-id = coin174876e802
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
}
```
