# hashgardcli box deposit-to

## 描述
对盒子进行通证注入。



## 用法
```shell
hashgardcli box deposit-to [box-id] [amount]  --from
```



### 子命令

| 名称   | 类型   | 必需 | 默认值 | 描述         |
| ------ | ------ | -------- | ------ | ------------ |
| box-id | string | 是       |        | 盒子的id |
| amount | int   | 是       |        | 存款的数量   |



## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子
### 进行存款

```
hashgardcli box deposit-to boxab3jlxpt2pw 300 --from
```


得到的结果是

```txt
{
  Height: 5657
  TxHash: 29C0A2CCFFDB38A64FB2187D8F7BA8AA86367F86C4FF10D131CEF6E9D5770235
  Data: 0F0E626F786162336A6C787074327077
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 44419
  Tags:
    - action = box_deposit
    - category = box
    - box-id = boxab3jlxpt2pw
    - box-Type = deposit
    - sender = gard1lgs73mwr56u2f4z4yz36w8mf7ym50e7myrqn65
    - operation = deposit-to
}
```



### 相关命令

| 名称                              | 描述                     |
| --------------------------------- | ------------------------ |
| [deposit-to](deposit-to.md)       | 用户对盒子进行存款       |
| [deposit-fetch](deposit-fetch.md) | 用户对盒子的存款进行取回 |
