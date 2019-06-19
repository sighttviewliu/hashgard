# hashgardcli issue describe

## 描述
Owner 可以对自己代币进行补充描述，描述文件使用不超过 1024 字节的 json 格式。可以自定义各种属性，也可以使用官方推荐的模板。
## 用法
```shell
 hashgardcli issue describe [issue-id] [description-file] [flags]
```
## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 给代币设置描述
```shell
hashgardcli issue describe coin174876e802 /description.json --from
```
#### 模板
```json
{
    "organization":"Hashgard",
    "website":"https://www.hashgard.com",
    "logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png",
    "intro":"这是一个牛逼的项目"
}
```
输入正确的密码之后，你的该代币的描述就设置成功了。
```txt
{
 Height: 3069
  TxHash: 02ED02AF5CD9C140C05D6C120BD7D57D196C27C9B3C794E6133DE912FD8243C1
  Data: 0F0E636F696E31373438373665383032
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 27465
  Tags:
    - action = issue_description
    - category = issue
    - issue-id = coin174876e802
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
}
```
### 查询发行信息
```shell
hashgardcli issue query-issue coin174876e802
```
最新的描述信息就生效了
```txt
{
Issue:
  IssueId:          			coin174876e802
  Issuer:           			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Owner:           				gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Name:             			issuename
  Symbol:    	    			AAA
  TotalSupply:      			9999999991024
  Decimals:         			18
  IssueTime:					1558179518
  Description:	    			{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"This is a description of the project"}
  BurnOwnerDisabled:  			false
  BurnHolderDisabled:  			false
  BurnFromDisabled:  			false
  FreezeDisabled:  				false
  MintingFinished:  			false
}
```
