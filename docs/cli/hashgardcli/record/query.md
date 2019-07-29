# hashgardcli record query

## Description
query record based on specified hash
## Usage
```shell
hashgardcli record query [hash] [flags]
```


### subcommand

| Name          | Type| Required  | Default| Description                               |
| ------------ | ------ | -------- | ------ | -------------------- |
| hash        | string | true       |        | hash  |


**Global flags, query command flags** [hashgardcli](../README.md)

## Example

```shell
hashgardcli record query CC38CAEE32149BEF4CCFAEAB518EC9A5FBC85AE6AC8D5A9F6CD710FAF5E4A22b
```
The result is as follows：
```shell
{
  "type": "record/RecordInfo",
  "value": {
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
}
```
