# hashgardcli stake delegations-to

## Description

Query delegations on an individual validator:

## Usage

```shell
hashgardcli stake delegations-to [validator-addr] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)


## Example

```shell
hashgardcli stake delegations-to gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx --trust-node
```

The result is as follows：

```txt
[
  {
    "delegator_addr": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
    "validator_addr": "gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx",
    "shares": "99.0000000000"
  }
]
```
