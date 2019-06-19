# hashgardcli lock list

## Description
查询指定类型盒子列表

## Usage
```shell
hashgardcli lock list
```

### flag

**Global flags, query command flags** [hashgardcli](../README.md)


## Example
## Query

```shell
hashgardcli lock list
```

The result is as follows：

```txt
[{
	"id": "boxaa3jlxpt2ps",
	"type": "lock",
	"status": "locked",
	"owner": "gard1gzs53ktf6u3yaplpqmx4wpq7n8s0gw75n73x3t",
	"name": "lockbox",
	"created_time": "1559531133",
	"total_amount": {
		"token": {
			"denom": "agard",
			"amount": "8987000000000000000000"
		},
		"decimals": "1"
	},
	"description": "{\"org\":\"Hashgard\",\"website\":\"https://www.hashgard.com\",\"logo\":\"https://cdn.hashgard.com/static/logo.2d949f3d.png\",\"intro\":\"Foundation lock\"}",
	"transfer_disabled": true,
	"lock": {
		"end_time": "1562123084"
	}
}]


```
