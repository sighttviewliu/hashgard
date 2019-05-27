# hashgardcli gov param

## Description

Query the all the parameters for the governance process

## Usage

```shell
 hashgardcli gov param [param-Type] [flags]
```
## Flags

**Global flags, query command flags** [hashgardcli](../README.md)


## Example

### Check by voting

```shell
hashgardcli gov param voting --trust-node -o=json --indent
```

The result is as follows：

```txt
{
  "voting_period": "172800000000000"
}
```

### Check by deposit

```shell
hashgardcli gov param deposit --trust-node -o=json --indent
```

The result is as follows：

```txt
{
  "min_deposit": [
    {
      "denom": "gard",
      "amount": "10"
    }
  ],
  "max_deposit_period": "172800000000000"
}
```


### Check by tallying
```shell
hashgardcli gov param tallying --trust-node -o=json --indent
```

The result is as follows：
```txt
{
  "quorum": "0.3340000000",
  "threshold": "0.5000000000",
  "veto": "0.3340000000",
  "governance_penalty": "0.0100000000"
}
```
