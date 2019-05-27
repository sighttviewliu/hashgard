# hashgardcli keys show

## Description

Return public details of one local key.

#

## Usage

```shell
hashgardcli keys show [name] [flags]
```

## Flags

| Name, shorthand | Type      | Required    |Default  | Description   |
| --------------- | ---------- | ---------- | -------- | -------- |
| -a, --address | string | false| "" | output the address only             |
| --bech               | string         | false           | acc    | The Bech32 prefix encoding for a key (acc/val/cons)|
| --multisig-threshold | int   | false      | 1   | [uint] K out of N required signatures                          |
| --pubkey     | bool | false | false  | Output only public key|

**Global flags, query command flags** [hashgardcli](../README.md)

## # Examples

### Show a given key

```shell
hashgardcli keys show MyKey
```

You'll get the local public keys with 'address' and 'pubkey' element of a given key.

```txt
NAME:   Type:   ADDRESS:                                    PUBKEY:
MyKey   local   gardkkm4w5pvmcw0e3vjcxqtfxwqpm3k0zakl7lxn5  gardaddwnpepq0gsl90v9dgac3r9hzgz53ul5ml5ynq89ax9x8qs5jgv5z5vyssskww57lw
```
