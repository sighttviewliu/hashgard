# hashgardcli deposit interest-cancel

## 描述

存款盒子发行期对于发行的盒子注入的利息取回。



## 用法

```shell
hashgardcli deposit interest-cancel [box-id] [amount] [flags]
```



### 子命令

| 名称   | 类型   | 必需 | 默认值 | 描述         |
| ------ | ------ | -------- | ------ | ------------ |
| box-id | string | 是       |        | 盒子的 id |
| amount | int    | 是       |        | 存款的数量   |



## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 取回注入的利息

```shell
hashgardcli deposit interest-cancel boxab3jlxpt2pt 200 --from $you_wallet_name
```

仅限注入地址取回注入的利息。



成功后，返回结果:
```txt
{
  Height: 3377
   TxHash: AAFECFC993E0B920C71ABEE8E21920726872018502468B71A22A03C12745412E
   Data: 0F0E626F786162336A6C787074327074
   Raw Log: [{"msg_index":"0","success":true,"log":""}]
   Logs: [{"msg_index":0,"success":true,"log":""}]
   GasWanted: 200000
   GasUsed: 44292
   Tags:
     - action = interest_cancel
     - category = deposit
     - id = boxab3jlxpt2pt
     - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
}
```
