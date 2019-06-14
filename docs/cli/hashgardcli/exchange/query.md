# hashgardcli exchange query

## Description

Query order details for the specified id

## Usage

```shell
hashgardcli exchange query [order_id] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### checking order

```shell
hashgardcli exchange query 1 --chain-id hashgard --indent -o=json
```

The result is as follows：

```txt
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
}
```



How to trade with orders?

Please click on the link below:

[take](take.md)
