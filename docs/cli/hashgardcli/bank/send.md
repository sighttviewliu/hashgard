# hashgardcli bank send

## Description

Create and sign a send tx

## Usage

```shell
hashgardcli bank send [to_address] [amount] [flags]
```
## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Create and sign a send tx

```shell
hashgardcli bank send gard1c9vrvvz08hd4entr0y5kfrt43v6malv60qtjfl 10gard --from=hashgard --chain-id=hashgard --indent -o json
```

After that, you can get remote node status as follows:

```shell
{
 "height": "21667",
 "txhash": "58110E97BD93CFA123B43B7C893386BA26F238570E1131A7B6E1E6ED5B7DA605",
 "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
 "gas_wanted": "200000",
 "gas_used": "22344",
 "tags": [
  {
   "key": "action",
   "value": "send"
  },
  {
   "key": "sender",
   "value": "gard10tfnpxvxjh6tm6gxq978ssg4qlk7x6j9aeypzn"
  },
  {
   "key": "recipient",
   "value": "gard1c9vrvvz08hd4entr0y5kfrt43v6malv60qtjfl"
  }
 ]
}
PS

```
