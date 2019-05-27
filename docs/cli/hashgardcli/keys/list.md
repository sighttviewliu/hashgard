# hashgardcli keys list

## Description

Return a list of all public keys stored by this key manager along with their associated name and address.

## Usage

```shell
hashgardcli keys list [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example


```shell
hashgardcli keys list
```

You'll get all the local public keys with 'address' and 'pubkey' element.

```txt
NAME:	Type:	ADDRESS:						            PUBKEY:
abc  	local	gardva2eu9qhwn5fx58kvl87x05ee4qrgh44yd8teh	gardpub1addwnpepqvu549hgyhnxlveqmtdn2xywygxpgzcsqefxur47zkz4e0e9x67hvjr6r6p
```
