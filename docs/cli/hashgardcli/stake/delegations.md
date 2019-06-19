# hashgardcli stake delegations

## Description

Query all delegations delegated from one delegator


## Usage

```shell
hashgardcli stake delegations [delegator-address] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

```shell
hashgardcli stake delegations gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet --trust-node
```

The result is as follows：

```json
[
    {
        "delegator_addr": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
        "validator_addr": "gardvaloper1xn4kvq867rap8vkrwfnp5n2entvpq2avtd0ytq",
        "shares": "11.0000000000"
    },
    {
        "delegator_addr": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
        "validator_addr": "gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx",
        "shares": "99.0000000000"
    }
]
```
