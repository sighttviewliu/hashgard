# hashgardcli keys delete

## Description

Delete a key from the store.

## Usage

```shell
hashgardcli keys delete <name> [flags]
```

## Flags

| Name, shorthand | Type| Required  | Default    | Description    |
| ------------- | ------ | --------- | --------- | ------------- |
| -f, --force | bool | false| false | Remove the key unconditionally without asking for the passphrase     |
| -y, --yes | bool | false| false | Skip confirmation prompt when deleting offline or ledger key references|

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Delete a given key

```shell
hashgardcli keys delete MyKey
```
You'll receive a danger warning and be asked to enter a password for your key.


```txt
DANGER - enter password to permanently delete key:
```

After you enter the correct password, you're done with deleting your key.

```txt
Password deleted forever (uh oh!)
```
