# hashgardcli bank withdraw

## Description
Retrieve tokens in the closed box

## Usage
```shell
hashgardcli bank withdraw [box-id]   --from
```



## Subcommands

| Name，shorthand | Type  | Required|Default| Description   |
| ------ | ------ | -------- | ------ | -------------- |
| box-id | string | true       |        | box id   |



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### retrieve tokens in the expired box

```shell
hashgardcli bank withdraw boxac3jlxpt2pw01 --from
```



The result is as follows：

```txt
{
  Response:
    Height: 2736
    TxHash: 5CB61311D144D1BDA0DBEE3B2A79B1166B881CD131E0F307E3D3FD118415F87B
    Data: 1110626F786163336A6C7870743270773031
    Raw Log: [{"msg_index":"0","success":true,"log":""}]
    Logs: [{"msg_index":0,"success":true,"log":""}]
    GasWanted: 200000
    GasUsed: 64153
    Tags:
      - action = withdraw
      - category = future
      - id = boxac3jlxpt2pw01
      - sender = gard180m2gg7fsm98gd7lrzakgy7qxlx6ke7p6zywe6
}
```
