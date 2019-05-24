# hashgardcli issue query

## Description
Query information about the release of the specified issue-id .

## Usage
```
hashgardcli issue query-issue [issue-id] [flags]
```
## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### Query issue information
```shell
hashgardcli issue query-issue coin174876e800
```
```txt
{
Issue:
  IssueId:          			coin174876e802
  Issuer:           			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Owner:           				gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Name:             			issuename
  Symbol:    	    			AAA
  TotalSupply:      			10000000001023
  Decimals:         			18
  IssueTime:					1558179518
  Description:	    			{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"a good project"}
  BurnOwnerDisabled:  			false
  BurnHolderDisabled:  			false
  BurnFromDisabled:  			false
  FreezeDisabled:  				false
  MintingFinished:  			false
 }
}
```
