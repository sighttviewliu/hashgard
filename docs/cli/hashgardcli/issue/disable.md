# hashgardcli issue disable

## Description

Non-reversible function setting of token

## Usage

```
hashgardcli issue disable [issue-id][flags]--from
```

## Flags

| Name          | Type| Required  | Default| Description                               |
| ------------- | ---- | -------- | ------ | --------------------------------------- |
| --burn-owner  | Bool | false   | false  | Disable owner to burn your own token|
| --burn-holder | bool | false  | false  | Disable Non-owner users burn their own tokens|
| --burn-from   | bool | false   | false  | Disable owner burning other user tokens|
| --minting     | bool | false   | false  | Disable the mint              |
| --freeze      | bool | false | false  | Disable freeze          |

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### disable mint of coin

```shell
hashgardcli issue disable coin174876e800 minting --from=
```

After entering the correct password，Disable token minting

```txt
{
 Height: 2255
  TxHash: EA08ACDF6ED5C15D2353B60001B3E4BB3BECC2293B3602AEED09492DE2659E50
  Data: 0F0E636F696E31373438373665383030
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 23013
  Tags:
    - action = issue_disable_feature
    - category = issue
    - issue-id = coin174876e800
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
    - feature = minting
}
```

Query the token

```shell
hashgardcli issue query-issue coin174876e800
```

The result is as follows：

```
{
 Issue:
  IssueId:          			coin174876e800
  Issuer:           			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Owner:           				gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Name:             			dedede
  Symbol:    	    			DDD
  TotalSupply:      			1000000
  Decimals:         			18
  IssueTime:					1558163118
  Description:	    			{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"新一代金融公有链"}
  BurnOwnerDisabled:  			false
  BurnHolderDisabled:  			false
  BurnFromDisabled:  			false
  FreezeDisabled:  				false
  MintingFinished:  			true

}
```
