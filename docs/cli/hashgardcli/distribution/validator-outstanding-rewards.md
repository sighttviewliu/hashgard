# hashgardcli distribution validator-outstanding-rewards

## Description

Query distribution outstanding (un-withdrawn) rewards for a validator and all their delegations

## Usage

```shell
hashgardcli distribution outstanding-rewards [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example


```shell
hashgardcli distribution outstanding-rewards -o=json --trust-node
```

The result is as follows：

```txt
[
 {
  "denom": "gard",
  "amount": "14.716656508693175779"
 }
]
```
