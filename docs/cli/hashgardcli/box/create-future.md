# hashgardcli box create-future

## Description
Create a new future box

## Usage
```shell
hashgardcli box create-future [name] [total-amount][mini-multiple] [distribute-file] --from
```

### Subcommands

| Name | Type  | Required  | Default| Description              |
| ------------- | ------ | -------- | ------ | ---------------------- |
| name          | string | true   |        | The name of the payment box |
| total-amount  | string | true   |        | Coin Type and quantity of payment |
| Mini-multiple | int    | true    | 1      | Minimum trading unit |

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

#### distribute-file

```json
{
   "time":[1657912000,1657912001,1657912002],
   "receivers":[
     ["gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj","100","200","300"],
     ["gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0","100","200","300"],
     ["gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz","100","200","300"]
    ]
}
```



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### Create future box
```shell
hashgardcli box create-future pay 1800coin174876e800  2 /Users/ming/Desktop/future.json --from
```
The result is as follows：
```txt
  {
 Height: 263
  TxHash: A34024F7C36A345A7C42519890F59D93B05D2FFE4EE33C0994E7D1981A3A1EA5
  Data: 0F0E626F786163336A6C787074327073
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 43797
  Tags:
    - action = box_create_future
    - category = box
    - box-id = boxac3jlxpt2ps
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7

    }
```

Deposit for the box for payment

```shell
hashgardcli box deposit-to boxac3jlxpt2ps 1800  --from
```

The result is as follows：

```shell
 {
  Height: 275
  TxHash: E96FBC4F9C2B3EB3B0C04B091DAAEF45E72E19C24E079879432460B077E137DF
  Data: 0F0E626F786163336A6C787074327073
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 140217
  Tags:
    - action = box_deposit
    - category = box
    - box-id = boxac3jlxpt2ps
    - box-Type = future
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
    - operation = deposit-to
}
```
Query box information

```shell
hashgardcli box query-box boxac3jlxpt2ps
```

The result is as follows：

```
BoxInfo:
  BoxId:			boxac3jlxpt2ps
  BoxStatus:			actived
  Owner:			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Name:				pay
  BoxType:			future
  TotalAmount:
  Token:			1800000000000000000000agard
  Decimals:			1
  CreatedTime:			1558090817
  Description:
  TradeDisabled:		true
FutureInfo:
  MiniMultiple:			1
  Deposit:			[
  Address:			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Amount:			1800000000000000000000]
  TimeLine:			[]
  Distributed:			[1657912000 1657912001 1657912002]
  Receivers:			[[gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj 100000000000000000000 200000000000000000000 300000000000000000000] [gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0 100000000000000000000 200000000000000000000 300000000000000000000] [gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz 100000000000000000000 200000000000000000000 300000000000000000000]]

```



### Available Commands

| Name          | Description              |
| --------------------------- | ---------------------- |
| [deposit-to](deposit-to.md) |Deposit the box    |
| [query-box](query-box.md)   | Query box information |
| [list-box](list-box.md)    | Query box list     |
