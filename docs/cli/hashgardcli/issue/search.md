# hashgardcli issue search

## Description
Search issues based on symbol

## Usage
```
hashgardcli issue search [symbol] [flags]
```
## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Search
```shell
hashgardcli issue search AAA
```
```txt
 [
    {
        "issue_id":"coin174876e802",
        "issuer":"gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7",
        "owner":"gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7",
        "issue_time":"1558179518",
        "name":"issuename",
        "symbol":"AAA",
        "total_supply":"10000000001023",
        "decimals":"18",
        "description":"{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"新一代金融公有链"}",
        "burn_owner_disabled":false,
        "burn_holder_disabled":false,
        "burn_from_disabled":false,
        "freeze_disabled":false,
        "minting_finished":false
    }
]

```
