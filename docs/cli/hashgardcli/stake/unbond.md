# hashgardcli stake unbond

## Description

Unbond an amount of bonded shares from a validator:

## Usage

```shell
hashgardcli stake unbond [validator-addr] [amount] [flags]
```

## Available Commands（Subcommands）

|     Name      | Type  | Required| Default| description         |
| -------------- | ------ | -------- | ------ | ------------------- |
| validator-addr | string | true     |        | Bech address of the validator |
| amount         | string   | true     |        | Amount of source-shares to either unbond or redelegate as a positive integer or decimal（Unit gard）|

## Flags

| Name   | Type  | Required| Default| description          |
| ------ | ------ | -------- | ------ | -------------------- |
| --from | string | true     |    | Delegators account name or address|

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

```shell
hashgardcli stake unbond \
gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx \
10gard \
--from=hashgard \
--chain-id=hashgard
```
