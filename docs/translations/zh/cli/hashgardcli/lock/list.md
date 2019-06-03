# hashgardcli lock list

## 描述
查询指定类型盒子列表

## 用法
```shell
hashgardcli lock list
```

### flag

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)


## 例子
### 查询盒子信息

```shell
hashgardcli lock list
```

成功后，返回结果:

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
