# hashgardcli record create

## Description

create  record
## Usage
```shell
hashgardcli record create [name] [hash] [flags]
```
### subcommand

| Name          | Type| Required  | Default| Description                               |
| ------------ | ------ | -------- | ------ | -------------------- |
| name         | string | true      |        | name of record   |
| hash  | string | true      |        | hash length≤64 and unique |

## Flags

| Name          | Type| Required  | Default| Description                               |
| --------------- | --------- | -------------------------- | ----------------------- | -------------------------- |
| --author      | string  | false |  |     author of record  |
| --record-type   | string  |  false |   | type of record |
| --description    | string   | false  |   | description of record |

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### Create
```shell
hashgardcli record create myRecord CC38CAEE32149BEF4CCFAEAB518EC9A5FBC85AE6AC8D5A9F6CD710FAF5E4A22b --from one --author Hashgard --record-type Transaction --description "hashgard:Good Projects，GARD:Good currency" -y
```
The result is as follows：
```shell
{
  "height": "104",
  "txhash": "19581EB61CD91B797C7009410329FA6BC423DAE650D62A28B8FCADC18430E303",
  "raw_log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
  "logs": [
    {
      "msg_index": "0",
      "success": true,
      "log": ""
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "26002",
  "tags": [
    {
      "key": "action",
      "value": "record"
    },
    {
      "key": "category",
      "value": "record"
    },
    {
      "key": "sender",
      "value": "gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60"
    },
    {
      "key": "id",
      "value": "rec174876e800"
    },
    {
      "key": "hash",
      "value": "CC38CAEE32149BEF4CCFAEAB518EC9A5FBC85AE6AC8D5A9F6CD710FAF5E4A22b"
    }
  ]
}
```
