# hashgardcli issue disable

## 描述

对代币的高级功能进行关闭，且该关闭不可逆。

## 用法

```
hashgardcli issue disable [issue-id][flags] --from
```

## Flags

| 名称          | 类型 | 必需 | 默认值 | 描述                                    |
| ------------- | ---- | -------- | ------ | --------------------------------------- |
| --burn-owner  | bool | 否       | false  | 关闭代币所有者销毁自己持有的代币功能    |
| --burn-holder | bool | 否       | false  | 关闭普通账号销毁该自己持有的代币功能    |
| --burn-from   | bool | 否       | false  | 关闭Owner销毁非管理员账户持有的代币功能 |
| --minting     | bool | 否       | false  | 关闭增发功能                        |
| --freeze      | bool | 否       | false  | 关闭冻结用户转入转出功能                |

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

### 禁用已经发行的代币的增发功能

```shell
hashgardcli issue disable coin174876e800 minting --from=
```

输入正确的密码之后，你就将你的通证的增发功能关闭了。

```txt
{
 Height: 2255
  TxHash: EA08ACDF6ED5C15D2353B60001B3E4BB3BECC2293B3602AEED09492DE2659E50
  Data: 0F0E636F696E31373438373665383030
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 23013
  Tags:
    - action = issue_disable_feature
    - category = issue
    - issue-id = coin174876e800
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
    - feature = minting
}
```

查询该通证情况

```shell
hashgardcli issue query-issue coin174876e800
```

你会看到增发代币的功能已经关闭。

```
{
 Issue:
  IssueId:          			coin174876e800
  Issuer:           			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Owner:           				gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Name:             			dedede
  Symbol:    	    			DDD
  TotalSupply:      			1000000
  Decimals:         			18
  IssueTime:					1558163118
  Description:	    			{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"新一代金融公有链"}
  BurnOwnerDisabled:  			false
  BurnHolderDisabled:  			false
  BurnFromDisabled:  			false
  FreezeDisabled:  				false
  MintingFinished:  			true

}
```
