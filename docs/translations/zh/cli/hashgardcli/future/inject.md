# hashgardcli  future inject

## 描述
对盒子进行通证注入。



## 用法
```shell
hashgardcli future inject [box-id] [amount][flags]
```



### 子命令

| 名称   | 类型   | 必需 | 默认值 | 描述         |
| ------ | ------ | -------- | ------ | ------------ |
| box-id | string | 是       |        | 盒子的 id |
| amount | int   | 是       |        | 存款的数量   |



## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 进行存款

```shell
hashgardcli future  inject boxac3jlxpt2ps 1800 --from $you_wallet_name
```


成功后，返回结果:

```txt
Response:
  Height: 1212
  TxHash: 3AE63B74450CEF0BB712B0788784310D5711A825BF60BC29821BD41EADC00FBF
  Data: 0F0E626F786163336A6C787074327073
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 134927
  Tags:
    - action = inject
    - category = future
    - id = boxac3jlxpt2ps
    - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
```
