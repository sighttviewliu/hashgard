# hashgardcli  deposit create

## Description
创建一个存款盒子，设定必要的参数。用来吸纳存款。盒子分为三个时期，1. 发行期。2. 存款吸纳期。3. 存款期。4.交割期



## Usage
```shell
 hashgardcli  deposit create [name] [total-amount] [flags]
```



## Subcommands

| Name      | Type  | Required | Default   | Description                |
| ------------ | ------ | -------- | ------ | -------------------- |
| name         | string | true       |        | The name of the box    |
| total-amount | string | true       |        | Deposit limit and coin type |



### Flags

| Name          | Type  | Required | Default   | Description             |
| ---------------- | ------ | -------- | ------ | ---------------------- |
| --bottom-line    | int    | true       |    | The number of boxes established |
| --price          | int    | true       |    | Minimum unit of deposit and can be divisible by the total amount |
| --start-time     | int    | true       |    | Start time of deposit |
| --establish-time | int    | true       |    | End time of deposit |
| --maturity-time  | int    | true       |    | Time to retrieve deposits and interest  |
| --interest       | string | true       |    | Inject interest amount and coin type |

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### create deposit box
```shell
hashgardcli deposit create  mingone 10000coin174876e800  --bottom-line=0 --price=2  --start-time=1558079700  --establish-time=1558080300 --maturity-time=1558080900 --interest=9898 coin174876e800  --from
```
The result is as follows：
```txt
  {
    Response:
      Height: 3353
      TxHash: D14BFA58729BEEB1FFE113108404B465757F1C8F7B4041DD91EAEA021870873A
      Data: 0F0E626F786162336A6C787074327074
      Raw Log: [{"msg_index":"0","success":true,"log":""}]
      Logs: [{"msg_index":0,"success":true,"log":""}]
      GasWanted: 200000
      GasUsed: 60857
      Tags:
        - action = create
        - category = deposit
        - id = boxab3jlxpt2pt
        - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
        - fee = 1000000000000000000000agard
    }
```



### Related commands

| Name                                     | Description                        |
| ------------------------------------------- | ---------------------------- |
| [interest-inject](interest-inject.md) | Inject interest into the deposit box     |
| [interest-cancel](interest-cancel.md)         | Retrieve the deposit box interest      |
| [inject](inject.md)                 | Deposit the box  |
| [cancel](cancel.md)           | Retrieve the deposit in the box |
| [params](params.md)                   |  Parameters and fees |
| [query](query.md)                     |   Query the specified id box  |
| [list](list.md)                     |   Query all deposit box lists |
| [search](search.md)                     |  Search by deposit box name  |
