# hashgardcli stake edit-validator

## Description

edit an existing validator account

## Usage

```shell
hashgardcli stake edit-validator [flags]
```

## Flags

| Name             | Type  | Required| Default           | Description           |
| ----------------- | ------ | -------- | ----------------- | -------------------- |
| --commission-rate | float  | false    | 0.0               | The new commission rate percentage|
| --details         | string | false    || The validator's (optional) details |
| --identity        | string | false    || The (optional) identity signature |
| --moniker         | string | false    || The validator's name |
| --website         | string | false    || The validator's (optional) website|
**After setting the information, please try not to modify it.**
**Global flags, query command flags** [hashgardcli](../README.md)

## Example

```shell
hashgardcli stake edit-validator \
--from=hashgard \
--chain-id=hashgard \
--website=http://hashgard.com \
--details=hashgard_is_great
```
