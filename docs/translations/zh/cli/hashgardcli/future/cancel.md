# hashgardcli future cancel

## 描述
用户对盒子注入的通证进行取回。



## 用法
```shell
hashgardcli future cancel [box-id] [amount] [flags]
```



### 子命令

| 名称   | 类型   | 必需 | 默认值 | 描述           |
| ------ | ------ | -------- | ------ | -------------- |
| box-id | string | 是       |        | 盒子的 id   |
| amount | int    | 是       |        | 取回 token 的数量 |



## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 取回 token

```shell
hashgardcli future cancel boxac3jlxpt2pt 100 --from $you_wallet_name
```



成功后，返回结果:

```txt
Response:
  Height: 1636
  TxHash: DC68FBCEA9F1C3B37781D05CA20FB455822C65047326D8B7CC40F7FE3B162E90
  Data: 0F0E626F786163336A6C787074327074
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 49564
  Tags:
    - action = cancel
    - category = future
    - id = boxac3jlxpt2pt
    - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
```
