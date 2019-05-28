# hashgardcli box deposit-to

## Description
Deposit the box



## Usage
```shell
hashgardcli box deposit-to [box-id] [amount]  --from
```



### Subcommands

| Name | Type  | Required  | Default| Description    |
| ------ | ------ | -------- | ------ | ------------ |
| box-id | string | true |        | box id |
| amount | int   | true |        | amount deposited  |



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### deposit

```shell
hashgardcli box deposit-to boxab3jlxpt2pw 300 --from
```

Deposit address only



The result is as follows：

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
    - box-type = deposit
    - sender = gard1lgs73mwr56u2f4z4yz36w8mf7ym50e7myrqn65
    - operation = deposit-to
}
```



### Available Commands

| Name                           | Description                |
| --------------------------------- | ------------------------ |
| [deposit-to](deposit-to.md)       | Deposit the box |
| [deposit-fetch](deposit-fetch.md) | Withdrawal of the box |
