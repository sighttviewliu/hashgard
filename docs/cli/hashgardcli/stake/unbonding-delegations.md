# hashgardcli stake unbonding-delegations

## Description

Query unbonding delegations for an individual delegator:

## Usage

```shell
hashgardcli stake unbonding-delegations [delegator-address] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example


```shell
hashgardcli stake unbonding-delegations faa13lcwnxpyn2ea3skzmek64vvnp97jsk8qmhl6vx
```

The result is as follows：

```json
[
    {
        "delegator_addr": "faa13lcwnxpyn2ea3skzmek64vvnp97jsk8qmhl6vx",
        "validator_addr": "fva15grv3xg3ekxh9xrf79zd0w077krgv5xf6d6thd",
        "creation_height": "1310",
        "min_time": "2018-11-15T06:24:22.754703377Z",
        "initial_balance": "0.02hashgard",
        "balance": "0.02hashgard"
    }
]
```
