# hashgardcli  deposit inject

## 描述
对盒子进行通证注入。



## 用法
```shell
hashgardcli deposit inject [box-id] [amount] [flags]
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
 hashgardcli deposit inject boxab3jlxpt2pt 10000 --from $you_wallet_name
```


成功后，返回结果:

```txt
{
  Response:
    Height: 3501
    TxHash: 4E5B7A297D99C2E7460E94F2D16B49F9A2CE54A6DDFBD681CBC83FAFD5B13ED3
    Data: 0F0E626F786162336A6C787074327074
    Raw Log: [{"msg_index":"0","success":true,"log":""}]
    Logs: [{"msg_index":0,"success":true,"log":""}]
    GasWanted: 200000
    GasUsed: 56194
    Tags:
      - action = inject
      - category = deposit
      - id = boxab3jlxpt2pt
      - sender = gard180m2gg7fsm98gd7lrzakgy7qxlx6ke7p6zywe6
}
```
