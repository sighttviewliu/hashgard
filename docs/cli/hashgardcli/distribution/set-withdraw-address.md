# hashgardcli distribution set-withdraw-addr

## Description

Set the withdraw address for rewards associated with a delegator address:

## Usage

```shell
hashgardcli distribution set-withdraw-addr [withdraw-address] [flags]
```
## Flags

**Global flags, query command flags** [hashgardcli](../README.md)


## Example
```shell
hashgardcli distribution set-withdraw-addr gard1c9vrvvz08hd4entr0y5kfrt43v6malv60qtjfl --from $you_wallet_name 
```

The result is as follows：

```txt
{
 "height": "33500",
 "txhash": "58AB9D329A043FC86DCE2B66E91BEDC1D13DD4000DF22E290041214C56DB04B8",
 "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
 "gas_wanted": "200000",
 "gas_used": "12018",
 "tags": [
  {
   "key": "action",
   "value": "set_withdraw_address"
  },
  {
   "key": "delegator",
   "value": "gard1c9vrvvz08hd4entr0y5kfrt43v6malv60qtjfl"
  }
 ]
}
```
