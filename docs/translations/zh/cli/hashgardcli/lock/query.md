# hashgardcli lock query

## 描述
查询指定盒子进行信息查询

## 用法
```shell
hashgardcli lock query [box-id]
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
hashgardcli box query-box boxac3jlxpt2ps
```

成功后，返回结果:

```txt
BoxInfo:
  Id:			boxaa3jlxpt2ps
  Status:			locked
  Owner:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Name:				love
  TotalAmount:
  Token:			8989000000000000000000agard
  Decimals:			1
  CreatedTime:			1559546607
  Description:
  TransferDisabled:		true
LockInfo:
  EndTime:			1564816979
```
