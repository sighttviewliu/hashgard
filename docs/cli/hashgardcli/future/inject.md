# hashgardcli  future inject

## Description
Inject coin into the future box



## Usage
```shell
hashgardcli future inject [box-id] [amount]  --from
```



## Subcommands

| Name| Type  | Required | Default   | Description        |
| ------ | ------ | -------- | ------ | ------------ |
| box-id | string | true       |        | 盒子的 id |
| amount | int   | true       |        | 存款的数量   |



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### Inject coin

```shell
hashgardcli future  inject boxac3jlxpt2ps 1800 --from
```


The result is as follows：

```txt
Response:
  Height: 1212
  TxHash: 3AE63B74450CEF0BB712B0788784310D5711A825BF60BC29821BD41EADC00FBF
  Data: 0F0E626F786163336A6C787074327073
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 134927
  Tags:
    - action = inject
    - category = future
    - id = boxac3jlxpt2ps
    - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
```
