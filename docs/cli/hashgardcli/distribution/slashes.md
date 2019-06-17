# hashgardcli distribution slashes

## Description

Query all slashes of a validator for a given block range:

## Usage

```shell
hashgardcli distribution slashes [validator] [start-height] [end-height] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example


```shell
hashgardcli distribution slashes gardvaloper1hr7vm7t7paeyg33ggd6efek2w58mu2huewltta 0 999999 -o=json --trust-node
```

The result is as follows：

```txt
[
 {
  "validator_period": "4",
  "fraction": "0.010000000000000000"
 },
 {
  "validator_period": "5",
  "fraction": "0.010000000000000000"
 },
 {
  "validator_period": "17",
  "fraction": "0.010000000000000000"
 }
]
```
