# hashgardcli lock query

## Description
Query by box ID

## Usage
```shell
hashgardcli lock query [box-id]
```

## Subcommands

| Name| Type  | Required | Default   | Description        |
| ------ | ------ | -------- | ------ | ------------ |
| box-id | string | true       |        | box id |



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
## Query

```shell
hashgardcli box query-box boxac3jlxpt2ps
```

The result is as follows：

```txt
BoxInfo:
  Id:			boxaa3jlxpt2ps
  Status:			locked
  Owner:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Name:				love
  TotalAmount:
  Token:			8989000000000000000000agard
  Decimals:			1
  CreatedTime:			1559546607
  Description:
  TransferDisabled:		true
LockInfo:
  EndTime:			1564816979
```
