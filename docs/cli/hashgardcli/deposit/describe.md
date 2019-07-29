# hashgardcli describe

## Description
Owner describes the box。The description file must be in josn format and no more than 1024 bytes.
## Usage
```shell
 hashgardcli deposit describe [box-id] [description-file] [flags]
```
## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### Set a description for the box
```shell
hashgardcli deposit describe boxab3jlxpt2pt .path/description.json --from $you_wallet_name
```
### Template
```shell
{
    "org":"Hashgard",
    "website":"https://www.hashgard.com",
    "logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png",
    "intro":"This is a description of the project"
}
```
The result is as follows：
```txt
{
  Response:
    Height: 3556
    TxHash: 031F0CE15F30D68142BD22B23A4555E428506FBDFA4D46490D182356BCCA97DB
    Data: 0F0E626F786162336A6C787074327074
    Raw Log: [{"msg_index":"0","success":true,"log":""}]
    Logs: [{"msg_index":0,"success":true,"log":""}]
    GasWanted: 200000
    GasUsed: 55284
    Tags:
      - action = describe
      - category = deposit
      - id = boxab3jlxpt2pt
      - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
}
```
## Query
```shell
hashgardcli deposit query boxab3jlxpt2pt
```
The result is as follows：
```shell
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
