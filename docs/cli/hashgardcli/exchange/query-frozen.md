# hashgardcli exchange query-frozen

## Description

Search frozen order list of a specific address

## Usage

```shell
hashgardcli exchange query-frozen [address] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Query address freeze

```shell
hashgardcli exchange query-frozen gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa --chain-id hashgard -o=json --indent
```

The result is as follows：

```txt
[
 {
  "denom": "apple",
  "amount": "33"
 },
 {
  "denom": "gard",
  "amount": "100"
 }
]
```
