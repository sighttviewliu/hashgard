# hashgardcli exchange list

## Description

Query all valid orders for the specified address

## Usage

```shell
hashgardcli exchange list [address] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

 ## Example

### Query all valid orders for the specified address

```shell
hashgardcli exchange list gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa --chain-id hashgard -o=json --indent
```

The result is as follows：

```txt
[
 {
  "order_id": "1",
  "seller": "gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa",
  "supply": {
   "denom": "gard",
   "amount": "100"
  },
  "target": {
   "denom": "apple",
   "amount": "800"
  },
  "remains": {
   "denom": "gard",
   "amount": "100"
  },
  "create_time": "2019-04-16T07:12:39.254071Z"
 },
 {
  "order_id": "2",
  "seller": "gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa",
  "supply": {
   "denom": "apple",
   "amount": "33"
  },
  "target": {
   "denom": "horn",
   "amount": "10000"
  },
  "remains": {
   "denom": "apple",
   "amount": "33"
  },
  "create_time": "2019-04-16T07:27:31.379469Z"
 }
]
```
