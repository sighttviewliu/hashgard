# hashgardcli deposit interest-inject

## Description
Inject interest into the box.



## Usage
```shell
hashgardcli deposit interest-inject [box-id] [amount]  --from
```



## Subcommands

| Name| Type  | Required | Default   | Description                  |
| ------ | ------ | -------- | ------ | ---------------------- |
| box-id | string | true       |        | box id           |
| amount | int    | true       |        | The amount of interest injected into the box |



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### Inject interest

```shell
hashgardcli deposit interest-inject boxab3jlxpt2ps 9898  --from
```

PS：`interest-inject` 注入的数量是指按最大值和时间来计算的。譬如发行一个 10000gard 的存款盒子，周期是 10 天，达成存款数量为 2000，利息总量是 500apple。那么日利率为 500/10/10000=0.5%。在 establish-time 的时候，如果只有 5000gard 存入，那么系统会自动退回 500*5000/10000=250gard 至利息注入的账户。



The result is as follows：

```txt
{
  Response:
  Height: 3381
  TxHash: 5736728B9F0EED5CD2EB9BBBA8B996938F38EC851FD39021657F57FC9B1D5AEB
  Data: 0F0E626F786162336A6C787074327074
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 46651
  Tags:
    - action = interest_inject
    - category = deposit
    - id = boxab3jlxpt2pt
    - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
}
```

注入利息 = 设定的利息数量时候开始，存款盒子激活，等待至 start-time 开始存款吸纳。

注入利息 < 设定利息数量，且到达 start-time 后，盒子失败，返还利息至注入账户
