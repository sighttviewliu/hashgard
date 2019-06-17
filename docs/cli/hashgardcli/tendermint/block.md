# hashgardcli tendermint block

## Description

Get verified data for a the block at given height

## Usage

```shell
  hashgardcli tendermint block [height] [flags]
```

## Flags

| Name, shorthand  | Default               | Description            | Required               |
| ------------ | --------------------- | -------------------------- | ---------------------- |
| --chain-id   | false                    | Chain ID of Tendermint node                     | false|
| --node       | tcp://localhost:26657 | falsede to connect to                      | false        |
| --trust-node | true        | Trust connected full node (don't verify proofs for responses)| false |

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

```shell
hashgardcli tendermint block 114263  --trust-node
```
