# hashgardcli lock search

## Description
Search Lock Box by Name
## Usage
```shell
hashgardcli lock search [name]
```
## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### search
```shell
hashgardcli lock search love
```
The result is as follows：
```shell
[{
	"id": "boxaa3jlxpt2ps",
	"type": "lock",
	"status": "locked",
	"owner": "gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60",
	"name": "love",
	"created_time": "1559546607",
	"total_amount": {
		"token": {
			"denom": "agard",
			"amount": "8989000000000000000000"
		},
		"decimals": "1"
	},
	"description": "",
	"transfer_disabled": true,
	"lock": {
		"end_time": "1564816979"
	}
}]
```
