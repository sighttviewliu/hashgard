# hashgardcli stake redelegations

## Description

Query all redelegation records for an individual delegator

## Usage

```shell
hashgardcli stake redelegations [delegator-address] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example


```shell
hashgardcli stake redelegations gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet --trust-node
```

The result is as follows：

```json
[
    {
        "delegator_addr": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
        "validator_src_addr": "gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx",
        "validator_dst_addr": "gardvaloper1xn4kvq867rap8vkrwfnp5n2entvpq2avtd0ytq",
        "creation_height": "24800",
        "min_time": "2018-12-21T02:49:44.731658304Z",
        "initial_balance": {
            "denom": "gard",
            "amount": "8"
        },
        "balance": {
            "denom": "gard",
            "amount": "8"
        },
        "shares_src": "8.9100000000",
        "shares_dst": "8.0000000000"
    }
]
```
