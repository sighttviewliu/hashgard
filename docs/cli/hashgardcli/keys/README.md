# hashgardcli keys

## Description

Keys allows you to manage your local keystore for tendermint.

## Usage

```shell
hashgardcli keys [command]
```

## Available Commands

| Name               | Description             |
| --------- | ----------------------------- |
| [mnemonic](mnemonic.md) | Compute the bip39 mnemonic for some input entropy                   |
| [add](add.md)           | Add an encrypted private key (either newly generated or recovered), encrypt it, and save to disk    |
| [list](list.md)         | List all keys                                           |
| [show](show.md)         | Show key info for the given name                          |
| [delete](delete.md)     | Delete the given key                                         |
| [update](update.md)     | Change the password used to protect private key                                            |

## Flags

| Name, shorthand      | Default  | Description     | Required |
| --------------- | ------- | ------------- | -------- |
| --help, -h      |         | help for keys |          |

## Flags

| Name, shorthand      | Default         | Description                              | Required |
| --------------- | -------------- | -------------------------------------- | -------- |
| --encoding, -e  | hex            | [string] Binary encoding (hex/b64）|
| --home          | $HOME/.hashgard | [string] directory for config and data |          |
| --output, -o    | text           | [string] Output format (text/json) |     |
| --trace         |                | print out full stack trace on errors   |          |

## Extended description

These keys may be in any format supported by go-crypto and can be used by light-clients, full nodes, or any other application that needs to sign with a private key.
