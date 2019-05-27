# hashgardcli exchange withdrawal-order

## Description

Cancel order and withdrawal token

## Usage

```shell
hashgardcli exchange withdrawal-order [order_id] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Cancel order

```shell
hashgardcli exchange withdrawal-order 2 --from mykey --chain-id hashgard -o=json --indent
```

The result is as follows：

```txt
{
 "height": "7162",
 "txhash": "7FA99BF4C271E8145A8DB695B9B08883A58A46F1AA369D8A7F6002684FDBC06A",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "200000",
 "gas_used": "27608",
 "tags": [
  {
   "key": "action",
   "value": "withdrawal_order"
  },
  {
   "key": "order_id",
   "value": "2"
  },
  {
   "key": "seller",
   "value": "gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa"
  },
  {
   "key": "order_status",
   "value": "inactive"
  }
 ]
}
```
