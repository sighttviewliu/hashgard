# hashgardcli keys update

## Description

Change the password used to protect private key

## Usage

```shell
hashgardcli keys update <name> [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Change the password of a given key

```shell
hashgardcli keys update MyKey
```
You'll be asked to enter the current password for your key.

```txt
Enter the current passphrase:
```

Then you'll be asked to enter a new password and repeat it.

```txt
Enter the new passphrase:
Repeat the new passphrase:
```

It will be done if you enter a new password that meets the criteria.

```txt
Password successfully updated!
```
