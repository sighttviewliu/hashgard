# hashgard unsafe-reset-all

## Description

Resets the blockchain database, removes address book files, and resets priv_validator.json to the genesis state

## Usage

```shell
hashgard unsafe-reset-all [flags]
```

## Flags

| Name, shorthand|Default     | description               | Required  |
| ---------- | ----------- | ------------------------- | -------- |
| -h, --help |             | help for unsafe-reset-all| false  |
| --home     | ~/.hashgard | directory for config and data  | false    |

## Example

``` shell
hashgard unsafe-reset-all
```
