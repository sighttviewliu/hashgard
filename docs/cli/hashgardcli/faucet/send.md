# hashgardcli faucet send

## Description

get some test coins from faucet account, this function just be available in testnet

## Usage

```shell
hashgardcli faucet send [address] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

```shell
 hashgardcli faucet send gard18vdg4r46qtfkwyghsg67dpn9p7vassw30z0f46 --chain-id=hashgard -o=json --indent
```

After successful execution, you will get 50gard, 50apple.

```txt
{
 "height": "6846",
 "txhash": "DBE1C8E78F91B3FBA1E000B92D751651DFAF3894C660DA0A7A19A11BA2CE7A56",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "50000",
 "gas_used": "23899",
 "tags": [
  {
   "key": "action",
   "value": "send"
  },
  {
   "key": "sender",
   "value": "gard1vka9svfaf5vg5cjaaju3d9g08ydk80z59ts0la"
  },
  {
   "key": "recipient",
   "value": "gard18vdg4r46qtfkwyghsg67dpn9p7vassw30z0f46"
  }
 ]
}
```
