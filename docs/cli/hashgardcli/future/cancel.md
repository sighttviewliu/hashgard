# hashgardcli future cancel

## Description
Retrieve the coin in the future box



## Usage
```shell
hashgardcli future cancel [box-id] [amount]  --from
```



## Subcommands

| Name| Type  | Required | Default   | Description          |
| ------ | ------ | -------- | ------ | -------------- |
| box-id | string | true       |        | 盒子的 id   |
| amount | int    | true       |        | 取回 token 的数量 |



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### Retrieve coin

```shell
hashgardcli future cancel boxac3jlxpt2pt 100 --from
```



The result is as follows：

```txt
Response:
  Height: 1636
  TxHash: DC68FBCEA9F1C3B37781D05CA20FB455822C65047326D8B7CC40F7FE3B162E90
  Data: 0F0E626F786163336A6C787074327074
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 49564
  Tags:
    - action = cancel
    - category = future
    - id = boxac3jlxpt2pt
    - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
```
