# hashgardcli future describe

## Description
Owner describes the box。The description file must be in josn format and no more than 1024 bytes.
## Usage
```shell
hashgardcli future describe [box-id] [description-file] [flags]
```
## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### describes the box
```shell
hashgardcli future describe boxaa3jlxpt2ps .path/Desktop/description.json --from $you_wallet_name
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
Response:
  Height: 1375
  TxHash: 4BD7D42A83D0A797F612F4FC1EED36208628AEE533CD82796EAE5AB51D6F6DCD
  Data: 0F0E626F786163336A6C787074327073
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 59373
  Tags:
    - action = describe
    - category = future
    - id = boxac3jlxpt2ps
    - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
```
## Query
```shell
hashgardcli future query boxac3jlxpt2ps
```
The result is as follows：
```shell
BoxInfo:
  Id:			boxac3jlxpt2ps
  Status:			actived
  Owner:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Name:				pay-one
  TotalAmount:
  Token:			1800000000000000000000agard
  Decimals:			1
  CreatedTime:			1559550097
  Description:			{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"This is a description of the project"}
  TransferDisabled:		true
FutureInfo:
  Injects:			[
  Address:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Amount:			1800000000000000000000]
  TimeLine:			[1593762909 1599119709 1606982109]
  Receivers:			[[gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj 100000000000000000000 200000000000000000000 300000000000000000000] [gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0 100000000000000000000 200000000000000000000 300000000000000000000] [gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz 100000000000000000000 200000000000000000000 300000000000000000000]]
  TotalWithdrawal:			0
```
