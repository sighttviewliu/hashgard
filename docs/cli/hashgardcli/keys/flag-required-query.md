# hashgardcli keys flag-required-query

## Description
Query Specific address of memo status

## Usage

```shell
hashgardcli keys flag-required-query [flag] [address] [flags]
```


## subcommand

| Name          | Type| Required  | Default| Description                               |
| --------------- | --------- | -------------------------- | ----------------------- | -------------------------- |
| address  | strings|   |   | address  |

## Flag

| Name          | Type| Required  | Default| Description                               |
| --------------- | --------- | -------------------------- | ----------------------- | -------------------------- |
| memo       | strings | Yes | 0 | memo           |


**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Query Specific address of memo status 

```shell
hashgardcli keys flag-required-query memo gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
```

The result is as follows：

```txt
{
  "sender": "gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60",
  "memo_required": true
}
```
