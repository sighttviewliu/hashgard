# hashgardcli issue describe

## Description
Owner Describes the issue token，Must be json file no larger than 1024 bytes.
## Usage
```
 hashgardcli issue describe [issue-id] [description-file] [flags]
```
## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### Set a description for the token
```shell
hashgardcli issue describe coin174876e802 /description.json --from
```
#### Template
```
{
    "organization":"Hashgard",
    "website":"https://www.hashgard.com",
    "logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png",
    "intro":"This is a good project"
}
```
The result is as follows：
```txt
{
 Height: 3069
  TxHash: 02ED02AF5CD9C140C05D6C120BD7D57D196C27C9B3C794E6133DE912FD8243C1
  Data: 0F0E636F696E31373438373665383032
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 27465
  Tags:
    - action = issue_description
    - category = issue
    - issue-id = coin174876e802
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
}
```
### Query issue information
```shell
hashgardcli issue query-issue coin174876e802
```
The result is as follows：
```
{
Issue:
  IssueId:          			coin174876e802
  Issuer:           			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Owner:           				gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Name:             			issuename
  Symbol:    	    			AAA
  TotalSupply:      			9999999991024
  Decimals:         			18
  IssueTime:					1558179518
  Description:	    			{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"新一代金融公有链"}
  BurnOwnerDisabled:  			false
  BurnHolderDisabled:  			false
  BurnFromDisabled:  			false
  FreezeDisabled:  				false
  MintingFinished:  			false
}
```
