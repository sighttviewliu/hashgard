# hashgardcli record list

## 描述
用户存储存证
## 用法
```shell
hashgardcli record list [flags]
```


## Flags

| 名称, 速记       | 类型   | 必需 | 默认值 | 描述                                                              |
| --------------- | --------- | -------------------------- | ----------------------- | -------------------------- |
| --sender       |   | 否 |  |          按发送存证地址进行查询                |


**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 查询存证列表
```shell
hashgardcli record list
```
返回结果
```shell
[
  {
    "id": "rec174876e802",
    "hash": "CC38CAEE32149BEF4CCFAEAB518EC9A5FBC85AE6AC8D5A9F6CD710FAF5E4A2Bs",
    "record_number": "",
    "sender": "gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60",
    "record_time": "1564126691",
    "name": "存证名称",
    "author": "bobo",
    "record_type": "docs",
    "description": ""
  },
  {
    "id": "rec174876e801",
    "hash": "CC38CAEE32149BEF4CCFAEAB518EC9A5FBC85AE6AC8D5A9F6CD710FAF5E4A2Bb",
    "record_number": "",
    "sender": "gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60",
    "record_time": "1564126576",
    "name": "存证名称",
    "author": "",
    "record_type": "",
    "description": ""
  },
  {
    "id": "rec174876e800",
    "hash": "CC38CAEE32149BEF4CCFAEAB518EC9A5FBC85AE6AC8D5A9F6CD710FAF5E4A2B3",
    "record_number": "",
    "sender": "gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60",
    "record_time": "1564126555",
    "name": "存证名称",
    "author": "",
    "record_type": "",
    "description": ""
  }
]
```
### 
```shell
hashgardcli record list --sender gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
```
返回结果
```shell
[
  {
    "id": "rec174876e800",
    "hash": "CC38CAEE32149BEF4CCFAEAB518EC9A5FBC85AE6AC8D5A9F6CD710FAF5E4A22b",
    "record_number": "",
    "sender": "gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60",
    "record_time": "1564135537",
    "name": "myRecord",
    "author": "Hashgard",
    "record_type": "Transaction",
    "description": ""
  }
]
```
