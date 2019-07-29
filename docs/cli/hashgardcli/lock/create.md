# hashgardcli lock create

## Description
用户将自己的通证进行限定期限的锁定。
## Usage
```shell
 hashgardcli lock create [name] [total-amount] [end-time] [flags]
```
## Subcommands

| Name      | Type  | Required | Default   | Description                |
| ------------ | ------ | -------- | ------ | -------------------- |
| Name         | string | true       |        | 盒子的名称       |
| total-amount | string | true       |        | 锁定通证的种类和数量 |
| end-time     | int    | true       |        | 锁定到期的时间       |



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)
 在线时间格式转换工具[Unix timestamp](../features/tool/Unix-timestamp.md)

## Example
### 创建锁定盒子
```shell
hashgardcli lock create lockbox 8987gard 1562123084 --from $you_wallet_name
```
输入正确的密码后，锁定完成。
```txt
Height: 4121
 TxHash: 9E7A793C5F69261DB55155498C2FFC723953F8273D84461C2EF3B435B3418020
 Data: 0F0E626F786161336A6C787074327073
 Raw Log: [{"msg_index":"0","success":true,"log":""}]
 Logs: [{"msg_index":0,"success":true,"log":""}]
 GasWanted: 200000
 GasUsed: 75410
 Tags:
   - action = create
   - category = lock
   - id = boxaa3jlxpt2ps
   - sender = gard1gzs53ktf6u3yaplpqmx4wpq7n8s0gw75n73x3t
   - fee = 100000000000000000000agard
```

接着我们对锁定的账户进行查询

```shell
hashgardcli bank account gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
```

得到的结果是

```txt
Account:
  Address:       gard1gzs53ktf6u3yaplpqmx4wpq7n8s0gw75n73x3t
  Pubkey:        gardpub1addwnpepqw8t7jqulvwqj77pl73h8cexx9lvcmjcrfm0l82h5txec2rjm2zmgwlf7m0
  Coins:         1000000000000000000000000118259956340999315537238agard,8987000000000000000000boxaa3jlxpt2ps
  AccountNumber: 0
  Sequence:      8

```
