# hashgardcli distribution rewards

## Description

Query all rewards earned by a delegator, optionally restrict to rewards from a single validator:

## Usage

```shell
hashgardcli distribution rewards [delegator-addr] [<validator-addr>] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example


```shell
hashgardcli distribution rewards gard1hr7vm7t7paeyg33ggd6efek2w58mu2hutvjrms -o=json --trust-node
hashgardcli distribution rewards gard1hr7vm7t7paeyg33ggd6efek2w58mu2hutvjrms gardvaloper1hr7vm7t7paeyg33ggd6efek2w58mu2huewltta -o=json --trust-node
```

Example response:

```txt
[
 {
  "denom": "gard",
  "amount": "0.131833867963517125"
 }
]
```
