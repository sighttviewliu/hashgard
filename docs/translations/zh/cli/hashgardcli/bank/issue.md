# hashgardcli bank issue

## 描述
根据发行 issue 的 ID 进行查询
## 用法
```shell
hashgardcli  bank issue [issueid]
```
## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 查询
```shell
hashgardcli bank issue  coin174876e800
```

成功后，返回结果:
```shell
Issue:
  IssueId:          			coin174876e800
  Issuer:           			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Owner:           				gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Name:             			bitcoin
  Symbol:    	    			BTC
  TotalSupply:      			1000000000000000000000000000000000000
  Decimals:         			18
  IssueTime:					1559557191
  Description:
  BurnOwnerDisabled:  			false
  BurnHolderDisabled:  			false
  BurnFromDisabled:  			false
  FreezeDisabled:  				false
  MintingFinished:  			false 
```
