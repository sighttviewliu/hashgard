# hashgardcli issue query

## 描述
查询指定issue-id值的发行的信息。

## 用法
```
hashgardcli issue query-issue [issue-id] [flags]
```

## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

### 查询发行信息

```shell
hashgardcli issue query-issue coin174876e800
```
```txt
{
Issue:
  IssueId:          			coin174876e802
  Issuer:           			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Owner:           				gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Name:             			issuename
  Symbol:    	    			AAA
  TotalSupply:      			10000000001023
  Decimals:         			18
  IssueTime:					1558179518
  Description:	    			{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"新一代金融公有链"}
  BurnOwnerDisabled:  			false
  BurnHolderDisabled:  			false
  BurnFromDisabled:  			false
  FreezeDisabled:  				false
  MintingFinished:  			false
 }
}
```
