# hashgard add-genesis-account

## Description
add genesis account to /path/to/.hashgard/config/genesis.json


## Usage
```shell
hashgard add-genesis-account [address_or_key_name] [coin][,[coin]] [flags]
```


## Subcommands
| Name         | Type  | Default| description                | Required |
| --------------------- | ------ | ------ | ------------------- | -------- |
| [address_or_key_name] | string |        | Added account name or address    | true    |
| [coin]                | string |        | coin Type and amount | true    |


## Flags
| Name，shorthand         | Type  | Default        | Description                      | Required |
| -------------------- | ------ | -------------- | -------------------------------- | -------- |
| -h, --help           |        |                | help for add-genesis-account  | false  |
| --home-client        | string | ~/.hashgardcli | client's home directory       | false   |
| --vesting-amount     | string |                | amount of coins for vesting accounts  | false    |
| --vesting-end-time   | int    |                | schedule end time (unix epoch) for vesting accounts| false    |
| --vesting-start-time | int    |                | schedule start time (unix epoch) for vesting accounts| false    |
| --home               | string | ~/.hashgard    | directory for config and data| false    |
| --trace              | bool   |                | print out full stack trace on errors| false   |


## Example
```shell
hashgardcli keys add root
hashgard add-genesis-account root 100000000gard
```
