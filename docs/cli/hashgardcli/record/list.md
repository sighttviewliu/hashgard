# hashgardcli record list

## Description
query all records
## Usage
```shell
hashgardcli record list [flags]
```


## Flags

| Name          | Type| Required  | Default| Description                               |
| --------------- | --------- | -------------------------- | ----------------------- | -------------------------- |
| --sender       |   | false |  |          query based on sender address         |


**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### query record List
```shell
hashgardcli record list
```
The result is as follows：
```shell
[
  {
    "id": "rec17lin4876e802",
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
The result is as follows：
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
