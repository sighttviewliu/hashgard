# hashgardcli box query-box

## 描述
查询指定盒子进行信息查询

## 用法
```shell
hashgardcli box query-box [box-id]
```

### 子命令

| 名称   | 类型   | 必需 | 默认值 | 描述         |
| ------ | ------ | -------- | ------ | ------------ |
| box-id | string | 是       |        | 盒子的id |



## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子
### 查询盒子信息

```
hashgardcli box query-box boxac3jlxpt2ps
```

返回盒子信息

```txt
BoxInfo:
  BoxId:			boxac3jlxpt2ps
  BoxStatus:			actived
  Owner:			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Name:				pay
  BoxType:			future
  TotalAmount:			
  Token:			1800000000000000000000agard
  Decimals:			1
  CreatedTime:			1558090817
  Description:			
  TradeDisabled:		true
FutureInfo:
  MiniMultiple:			1
  Deposit:			[
  Address:			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Amount:			1800000000000000000000]			
  TimeLine:			[]
  Distributed:			[1657912000 1657912001 1657912002]
  Receivers:			[[gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj 100000000000000000000 200000000000000000000 300000000000000000000] [gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0 100000000000000000000 200000000000000000000 300000000000000000000] [gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz 100000000000000000000 200000000000000000000 300000000000000000000]]

```
