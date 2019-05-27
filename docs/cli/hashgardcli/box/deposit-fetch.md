# hashgardcli box deposit-fetch

## Description
Fetch deposit from a deposit box



## Usage
```shell
hashgardcli box deposit-fetch [box-id] [amount]  --from
```



## Subcommands

| Name  | Type  | Required  | Default| Description      |
| ------ | ------ | -------- | ------ | -------------- |
| box-id | string | true   |        | Box id  |
| amount | Int    | true      |        | Number of retrieved |



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)



## Example

### Retrieve deposit

```shell
hashgardcli box interest-fetch boxab3jlxpt2pt 200 --from one
```



The result is as follows：

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
    - box-Type = deposit
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
    - operation = fetch
}
```



### Available Commands

| Name                                | Description                            |
| ----------------------------------- | -------------------------------------- |
| [interest-fetch](interest-fetch.md) | Retrieve interest|
