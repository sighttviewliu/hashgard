# hashgardcli deposit query

## Description
Query by deposit box ID

## Usage
```shell
hashgardcli deposit query [box-id]
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
hashgardcli deposit query boxab3jlxpt2pt
```

The result is as follows：

```txt
BoxInfo:
  Id:			boxab3jlxpt2pt
  Status:			interest
  Owner:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Name:				deposit88
  TotalAmount:
  Token:			10000000000000000000000agard
  Decimals:			1
  CreatedTime:			1559616171
  Description:			{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"This is a description of the project"}
  TransferDisabled:		true
DepositInfo:
  StartTime:			1559616600
  EstablishTime:		1559617200
  MaturityTime:			1559703600
  BottomLine:			0
  Interest:
  Token:			9090000000000000000000coin174876e800
  Decimals:			18
  Price:			2000000000000000000
  PerCoupon:			1.818000000000000000
  Share:			5000
  TotalInject:			10000000000000000000000
  WithdrawalShare:			0,
  WithdrawalInterest:			0,
  InterestInject:			[
  Address:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Amount:			9090000000000000000000]

```
