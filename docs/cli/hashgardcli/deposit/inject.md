# hashgardcli  deposit inject

## Description
Inject coin into the box



## Usage
```shell
hashgardcli deposit inject [box-id] [amount]  [flags]
```



## Subcommands

| Name| Type  | Required | Default   | Description        |
| ------ | ------ | -------- | ------ | ------------ |
| box-id | string | true       |        | 盒子的 id |
| amount | int   | true       |        | 存款的数量   |



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### 进行存款

```shell
 hashgardcli deposit inject boxab3jlxpt2pt 10000 --from
```


The result is as follows：

```txt
{
  Response:
    Height: 3501
    TxHash: 4E5B7A297D99C2E7460E94F2D16B49F9A2CE54A6DDFBD681CBC83FAFD5B13ED3
    Data: 0F0E626F786162336A6C787074327074
    Raw Log: [{"msg_index":"0","success":true,"log":""}]
    Logs: [{"msg_index":0,"success":true,"log":""}]
    GasWanted: 200000
    GasUsed: 56194
    Tags:
      - action = inject
      - category = deposit
      - id = boxab3jlxpt2pt
      - sender = gard180m2gg7fsm98gd7lrzakgy7qxlx6ke7p6zywe6
}
```
