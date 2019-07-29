# hashgardcli keys flag-required

## Description

Set the account address to transfer into memo status

## Usage

```shell
hashgardcli keys flag-required [flag] [is_required] [flags]
```
## subcommand

| Name          | Type| Required  | Default| Description                               |
| --------------- | --------- | -------------------------- | ----------------------- | -------------------------- |
| is_required   | Boolean |   |   | true/false  |



## Flag

| Name          | Type| Required  | Default| Description                               |
| --------------- | --------- | -------------------------- | ----------------------- | -------------------------- |
| memo       | strings| YES |  |memo                                    |


**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Set the account address to transfer into memo status

```shell
hashgardcli keys flag-required memo true --from $your_wallet_name
```

The result is as follows：

```txt
{
  "height": "260",
  "txhash": "5F5AFED2F5BA558D5110BDC74EC34AA5D969D39D79EBE46AFD981578A0243523",
  "raw_log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
  "logs": [
    {
      "msg_index": "0",
      "success": true,
      "log": ""
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "11421",
  "tags": [
    {
      "key": "action",
      "value": "account"
    },
    {
      "key": "category",
      "value": "account"
    },
    {
      "key": "sender",
      "value": "gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60"
    },
    {
      "key": "must-memo",
      "value": "true"
    }
  ]
}

```
