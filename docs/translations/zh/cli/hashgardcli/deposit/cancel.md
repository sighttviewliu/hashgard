# hashgardcli deposit cancel

## 描述
用户对盒子注入的通证进行取回。



## 用法
```shell
hashgardcli deposit cancel [box-id] [amount] [flags]
```



### 子命令

| 名称   | 类型   | 必需 | 默认值 | 描述           |
| ------ | ------ | -------- | ------ | -------------- |
| box-id | string | 是       |        | 盒子的 id   |
| amount | int    | 是       |        | 取回存款的数量 |



## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 进行存款取回

```shell
hashgardcli deposit cancel boxab3jlxpt2pt 10000 --from $you_wallet_name
```



成功后，返回结果:

```txt
{
  Response:
    Height: 3505
    TxHash: 5FBED5D99BBF2AA0E5A41F48CA4C1F2CD1BDBFC9B148F8B811615786E7710DCA
    Data: 0F0E626F786162336A6C787074327074
    Raw Log: [{"msg_index":"0","success":true,"log":""}]
    Logs: [{"msg_index":0,"success":true,"log":""}]
    GasWanted: 200000
    GasUsed: 54919
    Tags:
      - action = cancel
      - category = deposit
      - id = boxab3jlxpt2pt
      - sender = gard180m2gg7fsm98gd7lrzakgy7qxlx6ke7p6zywe6
}
```
