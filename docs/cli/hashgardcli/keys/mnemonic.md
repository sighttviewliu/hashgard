# hashgardcli keys mnemonic

## Description

Create a bip39 mnemonic, sometimes called a seed phrase, by reading from the system entropy. To pass your own entropy, use --unsafe-entropy

## Usage

```shell
hashgardcli keys mnemonic [flags]
```

## Flags

| Name, shorthand       | Default    | Description       | Required |
| ---------------- | --------- | --------- | -------- |
| --unsafe-entropy |           | Prompt the user to supply their own entropy, instead of relying on the system   |          |

**Global flags, query command flags** [hashgardcli](../README.md)

## Example


```shell
hashgardcli keys mnemonic MyKey
```

You'll get a bip39 mnemonic with 24 words.For security reasons, please pay attention to save, such as handwritten paper and save the paper properly.
```txt
police possible oval milk network indicate usual blossom spring wasp taste canal announce purpose rib mind river pet brown web response sting remain airport
```
