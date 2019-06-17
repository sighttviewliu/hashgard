# hashgardcli future create

## Description
Create future box.

## Usage
```shell
hashgardcli future create [name] [total-amount][distribute-file] [flags]
```



## Subcommands

| Name，shorthand | Type  | Required|Default| Description   |
| ------------- | ------ | -------- | ------ | ---------------------- |
| name          | string | true       |        | 盒子的名称         |
| total-amount  | string | true       |        | 支付的种类和数量       |
| Mini-multiple | int    | true       | 1      | 待收款凭证交易最小单位 |



#### distribute-file

```shell
{
   "time":[1657912000,1657912001,1657912002],//不同批次的支付时间
   "receivers":[
     ["gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj","100","200","300"],//支付的账户和数量
     ["gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0","100","200","300"],
     ["gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz","100","200","300"]
    ]
}
```


## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### create future box
```shell
hashgardcli future create  pay-one 1800gard  .path/future.json --from
```
The result is as follows：
```txt
Response:
  Height: 1169
  TxHash: 886D10C0E1B5ACD755D3333E08AB36F91E6D5DDCFD3ED073E34DCCFAE5637D3B
  Data: 0F0E626F786163336A6C787074327073
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 65595
  Tags:
    - action = create
    - category = future
    - id = boxac3jlxpt2ps
    - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
    - fee = 1000000000000000000000agard
```

为盒子存入需要支付的存款

```shell
hashgardcli future inject boxac3jlxpt2ps 1800  --from
```

The result is as follows：

```txt
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

Query by box id

```shell
hashgardcli future query boxac3jlxpt2ps
```

The result is as follows：

```txt
BoxInfo:
  Id:			boxac3jlxpt2ps
  Status:			actived
  Owner:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Name:				pay-one
  TotalAmount:
  Token:			1800000000000000000000agard
  Decimals:			1
  CreatedTime:			1559550097
  Description:
  TransferDisabled:		true
FutureInfo:
  Injects:			[
  Address:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Amount:			1800000000000000000000]
  TimeLine:			[1593762909 1599119709 1606982109]
  Receivers:			[[gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj 100000000000000000000 200000000000000000000 300000000000000000000] [gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0 100000000000000000000 200000000000000000000 300000000000000000000] [gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz 100000000000000000000 200000000000000000000 300000000000000000000]]
  TotalWithdrawal:			0
```
