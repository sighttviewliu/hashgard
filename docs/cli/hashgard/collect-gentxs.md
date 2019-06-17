# hashgard collect-gentxs


## Description
Collect genesis txs and output a genesis.json file

## Usage
```
hashgard collect-gentxs [flags]
```

## Flags
| Name，shorthand| Type  | Default                   | description                   | Required |
| ----------- | ------ | ------------------------- | ------------------------------ | -------- |
| --gentx-dir | string | ~/.hashgard/config/gentx/ |  override default "gentx" directory from which collect and execute genesis transactions| false  |
| -h, --help  |        |                           |  help for collect-gentxs                    | false  |
| --home      | string | ~/.hashgard               |  directory for config and data              | false  |
| --trace     | bool   |                           | print out full stack trace on erro          | false  |

## Example
`hashgard collect-gentxs`