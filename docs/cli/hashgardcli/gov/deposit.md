# hashgardcli gov deposit

## Description

Deposit tokens for active proposal


## Usage

```shell
hashgardcli gov deposit [proposal-id] [deposit] [flags]
```
## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Deposit

```shell
 hashgardcli gov deposit  1 50gard --from $you_wallet_name

You need to deposit 50 gard to activate the proposal

```txt
{
 "height": "106707",
 "txhash": "1D048A63AB37015700F22C5C90DA79127E0FFDBC8A9F5D2418B00D1916389B74",
 "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
 "gas_wanted": "200000",
 "gas_used": "32733",
 "tags": [
  {
   "key": "action",
   "value": "deposit"
  },
  {
   "key": "depositor",
   "value": "gard10tfnpxvxjh6tm6gxq978ssg4qlk7x6j9aeypzn"
  },
  {
   "key": "proposal-id",
   "value": "3"
  }
 ]
}
```

How to query deposit


[query-deposit](query-deposit.md)

[query-deposits](query-deposits.md)
