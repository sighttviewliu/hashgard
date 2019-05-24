# hashgardcli issue burn

## 描述
在代币的销毁属性没有关闭的情况下，可以对自己发行的代币进行销毁。
## 用法
```
 hashgardcli issue burn [issue-id] [amount] [flags]
```
## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子
### 销毁代币
```shell
hashgardcli issue burn coin174876e800 88888 --from
```
输入正确的密码之后，你的该代币的便完成了销毁。
```txt
{
   Height: 3020
  TxHash: 9C74FB0071940687E026EDEAB3666F8E3C0624C8541ABCF61C6BBFBFBA533F97
  Data: 0F0E636F696E31373438373665383032
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 27544
  Tags: 
    - action = issue_burn_holder
    - category = issue
    - issue-id = coin174876e802
    - sender = gard1lgs73mwr56u2f4z4yz36w8mf7ym50e7myrqn65
}
```