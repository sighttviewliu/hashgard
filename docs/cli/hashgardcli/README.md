# Command Line Client

## Flags of commands

| Name, shorthand| Type  | Required| Default               | Description    |
| ------------ | ------ | ----- | -------------------- | ------------ |
| --chain-id   | string | false |                   | Chain ID of tendermint node                                |
| --encoding   | string | false | "hex"                 | Encoding format(hex, b64, btc)                             |
| --home       | string | false | "\$HOME/.hashgardcli" | directory for config and data                              |
| --help, -h   | string | false |                       | Print help information                                     |
| --indent     | bool   | false | false                 | Add indent to JSON response                                |
| --ledger     | bool   | false | false                 | Use a connected Ledger device                              |
| --node       | string | false | tcp://localhost:26657 | <host>:<port> to tendermint rpc interface for this chain   |
| --output     | string | false | text                  | output format(text, json)                                  |
| --trace      | bool   | false |                       | print out full stack trace on errors   |
| --trust-node | bool   | false | true                  | Don't verify proofs for responses      |

## Flags of query

| Name,shorthand | Type| Required| Default| Description          |
| ---------- | ---- | ----- | ------ | --------------------------|
| --height   | int  | false | 0      | Block height to query, omit to get most recent provable block |

All query commands has these global flags. Their unique flags will be introduced later.

##  Flags of commands to send transactions

|Name, shorthand| Type  | Required| Default| Description            |
| ---------------- | ------ | ----- | ------ | ---------------- |
| --account-number | int    | false | 0      | AccountNumber to sign the tx                |
| --async          | bool   | false | false  | broadcast transactions asynchronously(only works with commit = false)|
| --dry-run        | bool   | false | false  | Ignore the --gas flag and perform a simulation of a transaction, but don't broadcast it    |
| --fees           | string | false |    | Fee to pay along with transaction E.g. 10gard, 1atom        |
| --from           | string | false |    | Name of private key with which to sign   |
| --gas            | string | false | 200000 | Gas limit to set per-transaction; set to "simulate" to calculate required gas automatically|
| --gas-adjustment | float  | false | 1      | Adjustment factor to be multiplied against the estimate returned by the tx simulation; if the gas limit is set|
| --gas-prices     | string | false |    | Decide on the transaction (`fees`)  gas prices，e.g.`0.001gard`;                |
| --generate-only  | bool   | false | false  | Build an unsigned transaction and write it to STDOUT|
| --memo           | string | false |    | Memo to send along with transaction|
| --print-response | bool   | false | true   | return tx response (only works with async = false)|
| --sequence int   | int    | false | 0      | Sequence number to sign the tx|

All commands which can be used to send transactions have these global flags. Their unique flags will be introduced later.

## Module command list

1. [bank command](./bank/README.md)
2. [distribution command](./distribution/README.md)
3. [gov command](./gov/README.md)
4. [keys command](./keys/README.md)
5. [stake command](./stake/README.md)
6. [status command](./status.md)
7. [tendermint command](./tendermint/README.md)
8. [slashing command](./slashing/README.md)
9. [issue command](./issue/README.md)
10. [box command](./box/README.md)
11. [faucet command](./faucet/send.md)

## hashgardcli Config command

`config` command interactively configures some default parameters, such as chain-id, home, fees, and node.

`hashgardcli config <key> [value] [flags]`

## hashgardcli init command

`hashgardcli init` Command can initialize the client
