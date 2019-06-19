# hashgardcli bank issue

## Description
Query based on issueID
## Usage
```shell
hashgardcli  bank issue [issueid]
```
## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
## Query
```shell
hashgardcli bank issue  coin174876e800
```

The result is as follows：
```shell
Issue:
  IssueId:          			coin174876e800
  Issuer:           			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Owner:           				gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Name:             			bitcoin
  Symbol:    	    			BTC
  TotalSupply:      			1000000000000000000000000000000000000
  Decimals:         			18
  IssueTime:					1559557191
  Description:
  BurnOwnerDisabled:  			false
  BurnHolderDisabled:  			false
  BurnFromDisabled:  			false
  FreezeDisabled:  				false
  MintingFinished:  			false
```
