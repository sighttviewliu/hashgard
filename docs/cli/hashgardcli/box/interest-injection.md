# hashgardcli box interest-injection



## Description

Any address Inject interest into the box 。



## Usage
```shell
hashgardcli box interest-injection [box-id] [amount]  --from
```



### Subcommands

| Name | Type  | Required  | Default| Description              |
| ------ | ------ | -------- | ------ | ---------------------- |
| box-id | string | true     |        | box id           |
| amount | int    | true    |         | amount of interest |



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)



## Example

### Inject interest

```
hashgardcli box interest-injection boxab3jlxpt2ps 9898  --from
```

PS：interest-injection注入的数量是指按最大值和时间来计算的。譬如发行一个10000gard的存款盒子，周期是10天，达成存款数量为2000，利息总量是500apple。那么日利率为500/10/10000=0.5%。在establish-time的时候，如果只有5000gard存入，那么系统会自动退回500*5000/10000=250gard 至利息注入的账户。


The result is as follows：

```txt
{
   Height: 4169
  TxHash: 488BC63DBB898DF493B1C82E891971559B591CD1B4F9E41D2E1215F0BF42E024
  Data: 0F0E626F786162336A6C787074327073
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 50800
  Tags:
    - action = box_interest
    - category = box
    - box-id = boxab3jlxpt2ps
    - box-Type = deposit
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
    - operation = injection
}
```

注入利息=设定的利息数量时候开始，存款盒子激活，等待至start-time开始存款吸纳。

注入利息<设定利息数量，且到达start-time后，盒子失败，返还利息至注入账户

### Available Commands

| Name                            | Description                |
| ----------------------------------- | ------------------------ |
| [interest-fetch](interest-fetch.md) | Withdrawal interest of the box  |
