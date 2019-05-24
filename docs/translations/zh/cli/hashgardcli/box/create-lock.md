# hashgardcli box create-lock

## 描述
用户将自己的通证进行限定期限的锁定。
## 用法
```
 hashgardcli box create-lock [name] [total-amount] [end-time] --from
```
### 子命令

| 名称         | 类型   | 必需 | 默认值 | 描述                 |
| ------------ | ------ | -------- | ------ | -------------------- |
| name         | string | 是       |        | 盒子的名称       |
| total-amount | string | 是       |        | 锁定通证的种类和数量 |
| end-time     | int    | 是       |        | 锁定到期的时间       |



## Flags

 ### 全局flags 参考：[hashgardcli](../README.md)

## 例子
### 创建锁定盒子
```shell
hashgardcli box create-lock ff 1000coin174876e800 1558066440 --from
```
输入正确的密码后，锁定完成。
```txt
  {Height: 1936
  TxHash: B32D14F7F9D208733EB522CA80B4AB1CA6667271862DE2182E8501CF645E763D
  Data: 0F0E626F786161336A6C787074327074
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 70033
  Tags:
    - action = box_create
    - category = box
    - box-id = boxaa3jlxpt2pt
    - box-Type = lock
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
    }
```

接着我们对锁定的账户进行查询

```
hashgardcli bank account gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
```

得到的结果是

```txt
{
  Account:
  Address:       gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Pubkey:        gardpub1addwnpepqfpd8mkl3jg43fw7y02fe99cgaxutf5npv9y9gx9dvrrcdwl36shv694apw
  Coins:         1000000000000000000000boxaa3jlxpt2pt,9999999907005070apple(coin174876e800)
  AccountNumber: 0
  Sequence:      7
}
```



### 相关命令

| 名称                      | 描述                   |
| ------------------------- | ---------------------- |
| [query-box](query-box.md) | 对指定盒子进行信息查询 |
| [list-box](list-box.md)  | 罗列指定类型盒子列表   |
