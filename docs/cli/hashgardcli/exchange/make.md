# hashgardcli exchange make

## Description

Create a new order

## Usage

```shell
hashgardcli exchange create [flags]
```

## Flags

| Name     | Type                 | Required                 | Default        | Description   |
| -------- | --------- | ------------- | ---------------------- | -------- |
| --supply     | string | true|| The coin type and amount of the order   |
| --target        | string | true|| The coin type and amount of the target        |

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Create order

```shell
hashgardcli exchange make --supply 100gard --target 800apple --from mykey --chain-id hashgard --indent -o=json
```

Created a 100gard exchange 800apple order

```txt
{
 "height": "6907",
 "txhash": "12FE1C3ECDF960AFB6A5E2D1A0DCE678EDB9A35812137AFA2CFA0DF7340C8F12",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "200000",
 "gas_used": "35757",
 "tags": [
  {
   "key": "action",
   "value": "make"
  },
  {
   "key": "order_id",
   "value": "1"
  },
  {
   "key": "seller",
   "value": "gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa"
  },
  {
   "key": "supply_token",
   "value": "gard"
  },
  {
   "key": "target_token",
   "value": "apple"
  }
 ]
}
```

How to check orders, view all orders at an address, view funds frozen at an address, and cancel orders?

Please click on the link below:

[query](query.md)
[list](list.md)
[query-frozen](query-frozen.md)
