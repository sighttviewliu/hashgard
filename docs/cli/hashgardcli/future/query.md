# hashgardcli future query

## Description
Query by future box ID

## Usage
```shell
hashgardcli future query [box-id]
```

## Subcommands

| Name| Type  | Required | Default   | Description        |
| ------ | ------ | -------- | ------ | ------------ |
| box-id | string | true       |        | box ID|



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
## Query

```shell
hashgardcli future query boxac3jlxpt2pt
```

The result is as follows：

```txt
BoxInfo:
  Id:			boxac3jlxpt2pt
  Status:			injecting
  Owner:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Name:				pay-two
  TotalAmount:
  Token:			1800000000000000000000agard
  Decimals:			1
  CreatedTime:			1559552267
  Description:
  TransferDisabled:		true
FutureInfo:
  Injects:			[
  Address:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Amount:			100000000000000000000]
  TimeLine:			[1593762909 1599119709 1606982109]
  Receivers:			[[gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj 100000000000000000000 200000000000000000000 300000000000000000000] [gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0 100000000000000000000 200000000000000000000 300000000000000000000] [gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz 100000000000000000000 200000000000000000000 300000000000000000000]]
  TotalWithdrawal:			0
```
