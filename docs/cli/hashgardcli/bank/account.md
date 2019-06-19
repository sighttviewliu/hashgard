# hashgardcli bank account

## Description

This command is used for querying balance information of certain address.


## Usage

```shell
hashgardcli bank account [address] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Query your account in trust-mode

```shell
hashgardcli bank account gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx -o json --trust-node --indent
```

After that, you will get the detail info for the account.

```shell
{
 "type": "auth/Account",
 "value": {
  "address": "gard10tfnpxvxjh6tm6gxq978ssg4qlk7x6j9aeypzn",
  "coins": [
   {
    "denom": "gard",
    "amount": "1900000000"
   }
  ],
  "public_key": {
   "type": "tendermint/PubKeySecp256k1",
   "value": "AvM1uBBEml3ZtXP5GZD6vr7UIcit6GMjS0ZUdxuejShH"
  },
  "account_number": "0",
  "sequence": "1"
 }
}
```
If you query an wrong account, you will get the follow information.
```shell
hashgardcli bank account gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429zz
ERROR: decoding bech32 failed: checksum failed. Expected d429yx, got d429zz.
```
If you query an account with no transactions on the chain, you will get the follow error.
```shell
hashgardcli bank account gardkenrwk5k4ng70e5s9zfsttxpnlesx5ps0gfdv7
ERROR: No account with address gardkenrwk5k4ng70e5s9zfsttxpnlesx5ps0gfdv7 was found in the state.
Are you sure there has been a transaction involving it?
```
