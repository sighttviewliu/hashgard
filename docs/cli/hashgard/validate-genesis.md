# hashgard validate-genesis

## Description

validates the genesis file at the default location or at the location passed as an arg

## Usage

```shell
hashgard validate-genesis [file] [flags]
```

## Available Commands

| Name, shorthand|Type  | Default                         | Description        | Required  |
| ---------- | ------ | ------------------------------- | ---------------- | -------- |
| [file]     | string | ~/.hashgard/config/genesis.json | genesis 文件位置 | false  |

## Flags

| Name, shorthand|Type  | Default     | Description                        | Required  |
| ---------- | ------ | ----------- | -------------------------------- | -------- |
| -h, --help |        |             | help for validate-genesis | false  |
| --home     | string | ~/.hashgard | directory for config and data                | false  |
| --trace    | bool   |             | print out full stack trace on errors         | false  |

## Example

```bash
hashgard validate-genesis
```
