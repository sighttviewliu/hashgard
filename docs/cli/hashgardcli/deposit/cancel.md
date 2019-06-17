# hashgardcli deposit cancel

## Description

Retrieve the deposit in the box


## Usage
```shell
hashgardcli deposit cancel [box-id] [amount] [flags]
```



## Subcommands

| Name，shorthand | Type  | Required|Default| Description   |
| ------ | ------ | -------- | ------ | -------------- |
| box-id | string | true       |        | box id   |
| amount | int    | true       |        | Retrieve quantity |



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### Retrieve the deposit in the box

```shell
hashgardcli deposit cancel boxab3jlxpt2pt 10000 --from
```



The result is as follows：

```txt
{
  Response:
    Height: 3505
    TxHash: 5FBED5D99BBF2AA0E5A41F48CA4C1F2CD1BDBFC9B148F8B811615786E7710DCA
    Data: 0F0E626F786162336A6C787074327074
    Raw Log: [{"msg_index":"0","success":true,"log":""}]
    Logs: [{"msg_index":0,"success":true,"log":""}]
    GasWanted: 200000
    GasUsed: 54919
    Tags:
      - action = cancel
      - category = deposit
      - id = boxab3jlxpt2pt
      - sender = gard180m2gg7fsm98gd7lrzakgy7qxlx6ke7p6zywe6
}
```
