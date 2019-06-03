# hashgardcli issue mint

## 描述
在可以增发的前提下，Owner 可以对自己发行的代币进行增发，默认增发到自己的账户，也可以增发到指定的账号。

## 用法
```shell
 hashgardcli issue mint [issue-id] [amount] [flags]
```
## 子命令
| 名称                | 类型     | 必需                 | 默认值                      | 描述                |
| -----------------  | -------------------------- | ----------------- | --------------- | ------------------ |
| --to                  | string | 否 || （可选）增发到指定的账号地址                   |

## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

### 增发到指定的地址
```shell
hashgardcli issue mint coin174876e802 9999 --to=gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7 --from
```
输入正确的密码之后，你的该代币的便完成了增发。
```txt
{
  Height: 3138
  TxHash: 110F99B71B2F206E29EDA2A5EC9DB1E372045693C06EDB9C32B9C9767AB92F93
  Data: 0F0E636F696E31373438373665383032
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 40402
  Tags:
    - action = issue_mint
    - category = issue
    - issue-id = coin174876e802
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
}
```
