# hashgardcli future query

## 描述
查询指定盒子进行信息查询

## 用法
```shell
hashgardcli future query [box-id]
```

### 子命令

| 名称   | 类型   | 必需 | 默认值 | 描述         |
| ------ | ------ | -------- | ------ | ------------ |
| box-id | string | 是       |        | 盒子的 id |



## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 查询盒子信息

```shell
hashgardcli future query boxac3jlxpt2pt
```

成功后，返回结果:

```txt
BoxInfo:
  Id:			boxac3jlxpt2pt
  Status:			injecting
  Owner:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Name:				pay-two
  TotalAmount:
  Token:			1800000000000000000000agard
  Decimals:			1
  CreatedTime:			1559552267
  Description:
  TransferDisabled:		true
FutureInfo:
  Injects:			[
  Address:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Amount:			100000000000000000000]
  TimeLine:			[1593762909 1599119709 1606982109]
  Receivers:			[[gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj 100000000000000000000 200000000000000000000 300000000000000000000] [gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0 100000000000000000000 200000000000000000000 300000000000000000000] [gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz 100000000000000000000 200000000000000000000 300000000000000000000]]
  TotalWithdrawal:			0
```
