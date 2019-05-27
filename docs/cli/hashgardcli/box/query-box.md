# hashgardcli box query-box

## Description
Query specified box information

## Usage
```shell
hashgardcli box query-box [box-id]
```

### Subcommands

| Name  | Type  | Required  | Default| Description    |
| ------ | ------ | -------- | ------ | ------------ |
| box-id | string | true      |        | box id|



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### Query box

```shell
hashgardcli box query-box boxac3jlxpt2ps
```

The result is as follows：

```txt
BoxInfo:
  BoxId:			boxac3jlxpt2ps
  BoxStatus:			actived
  Owner:			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Name:				pay
  BoxType:			future
  TotalAmount:
  Token:			1800000000000000000000agard
  Decimals:			1
  CreatedTime:			1558090817
  Description:
  TradeDisabled:		true
FutureInfo:
  MiniMultiple:			1
  Deposit:			[
  Address:			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Amount:			1800000000000000000000]
  TimeLine:			[]
  Distributed:			[1657912000 1657912001 1657912002]
  Receivers:			[[gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj 100000000000000000000 200000000000000000000 300000000000000000000] [gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0 100000000000000000000 200000000000000000000 300000000000000000000] [gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz 100000000000000000000 200000000000000000000 300000000000000000000]]

```
