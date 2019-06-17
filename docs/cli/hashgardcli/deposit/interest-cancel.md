# hashgardcli deposit interest-cancel

## Description

Retrieve the interest in the box.



## Usage

```shell
hashgardcli deposit interest-cancel [box-id] [amount] [flags]
```



## Subcommands

| Name| Type  | Required | Default   | Description        |
| ------ | ------ | -------- | ------ | ------------ |
| box-id | string | true       |        | 盒子的 id |
| amount | int    | true       |        | 存款的数量   |



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### 取回注入的利息

```shell
hashgardcli deposit interest-cancel boxab3jlxpt2pt 200 --from
```

仅限注入地址取回注入的利息。



The result is as follows：

```txt
{
  Height: 3377
   TxHash: AAFECFC993E0B920C71ABEE8E21920726872018502468B71A22A03C12745412E
   Data: 0F0E626F786162336A6C787074327074
   Raw Log: [{"msg_index":"0","success":true,"log":""}]
   Logs: [{"msg_index":0,"success":true,"log":""}]
   GasWanted: 200000
   GasUsed: 44292
   Tags:
     - action = interest_cancel
     - category = deposit
     - id = boxab3jlxpt2pt
     - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
}
```
