# hashgardcli distribution withdraw-rewards

## Description

witdraw rewards from a given delegation address, and optionally withdraw validator commission if the delegation address given is a validator operator

## Usage

```shell
hashgardcli distribution withdraw-rewards [validator-addr] [flags]
```

## Flags

| Name               | Type  | Required| Default | description        |
| --------------------- | -----  | -------- | -------- | --- |
| --commission | bool | false| false  | also withdraw validator's commission |

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

1. Retrieve the proceeds from the delegator
    ```shell
    hashgardcli distribution withdraw-rewards gard34mhjjyyc7mehvaay0f3d4hj8qx3ee3w3eq5nq --from mykey --chain-id=hashgard
    ```
2. If the delegator is a owner of a validator, withdraw all delegation rewards and validator commission rewards:
    ```shell
    hashgardcli distribution withdraw-rewards --commission=true from mykey  --chain-id=sif-1000
    ```
