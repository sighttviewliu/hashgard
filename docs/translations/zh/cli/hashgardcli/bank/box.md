# hashgardcli bank box

## 描述
根据发行盒子的 ID 进行查询
## 用法
```shell
hashgardcli  bank box [boxid]
```
## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 查询
```shell
hashgardcli bank box  boxac3jlxpt2pw
```

成功后，返回结果:
```shell
BoxInfo:
  Id:			boxac3jlxpt2pw
  Status:			actived
  Owner:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Name:				pay-two
  TotalAmount:
  Token:			1800000000000000000000agard
  Decimals:			1
  CreatedTime:			1559555762
  Description:
  TransferDisabled:		true
FutureInfo:
  Injects:			[
  Address:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Amount:			1800000000000000000000]
  TimeLine:			[1559556000 1599119709 1606982109]
  Receivers:			[[gard180m2gg7fsm98gd7lrzakgy7qxlx6ke7p6zywe6 100000000000000000000 200000000000000000000 300000000000000000000] [gard180m2gg7fsm98gd7lrzakgy7qxlx6ke7p6zywe6 100000000000000000000 200000000000000000000 300000000000000000000] [gard180m2gg7fsm98gd7lrzakgy7qxlx6ke7p6zywe6 100000000000000000000 200000000000000000000 300000000000000000000]]
  TotalWithdrawal:			0
```
