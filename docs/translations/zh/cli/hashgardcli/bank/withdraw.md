# hashgardcli bank withdraw

## 描述
使用凭证对到期的盒子进行 Token 兑换



## 用法
```shell
hashgardcli bank withdraw [box-id] [flags]
```



### 子命令

| 名称   | 类型   | 必需 | 默认值 | 描述           |
| ------ | ------ | -------- | ------ | -------------- |
| box-id | string | 是       |        | 盒子的 id   |



## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 对到期盒子中的通证进行取回

```shell
hashgardcli bank withdraw boxac3jlxpt2pw01 --from
```



得到的结果是

```txt
{
  Response:
    Height: 2736
    TxHash: 5CB61311D144D1BDA0DBEE3B2A79B1166B881CD131E0F307E3D3FD118415F87B
    Data: 1110626F786163336A6C7870743270773031
    Raw Log: [{"msg_index":"0","success":true,"log":""}]
    Logs: [{"msg_index":0,"success":true,"log":""}]
    GasWanted: 200000
    GasUsed: 64153
    Tags:
      - action = withdraw
      - category = future
      - id = boxac3jlxpt2pw01
      - sender = gard180m2gg7fsm98gd7lrzakgy7qxlx6ke7p6zywe6
}
```
