# hashgardcli lock search

## 描述
根据发行盒子的名字进行查询
## 用法
```shell
hashgardcli lock search [name]
```
## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 搜索
```shell
hashgardcli lock search love
```
成功后，返回结果:
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
