# hashgardcli box interest-fetch

## 描述

存款盒子发行期对于发行的盒子注入的利息取回。



## 用法

```shell
hashgardcli box interest-fetch [box-id] [amount]  --from
```



### 子命令

| 名称   | 类型   | 必需 | 默认值 | 描述         |
| ------ | ------ | -------- | ------ | ------------ |
| box-id | string | 是       |        | 盒子的 id |
| amount | int    | 是       |        | 存款的数量   |



## Flags

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子


```shell
hashgardcli box interest-fetch boxab3jlxpt2pt 200 --from
```

仅限注入地址取回注入的利息。



得到的结果是

```txt
{
   Height: 5037
  TxHash: E3743F7EF405600B23C2987C4689FC49F64BEF6DC3CA8A5A75A975B084FCCEE5
  Data: 0F0E626F786162336A6C787074327074
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 48149
  Tags:
    - action = box_interest
    - category = box
    - box-id = boxab3jlxpt2pt
    - box-type = deposit
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
    - operation = fetch

}
```



### 相关命令

| 名称                                        | 描述               |
| ------------------------------------------- | ------------------ |
| [interest-injection](interest-injection.md) | 用户对盒子利息注入 |
