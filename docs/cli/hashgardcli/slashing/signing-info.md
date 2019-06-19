# hashgardcli slashing signing-info

## Description

Use a validators' consensus public key to find the signing-info for that validator:

## Usage

```shell
hashgardcli slashing signing-info [validator-conspub] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example


```shell
hashgardcli slashing signing-info \
gardvalconspub1zcjduepqgsmuj0qallsw79hjj9qztcke6hj3ujdcpjv249uny9fvzp4eulms0tqvgs \
--trust-node
```

The result is as follows：

```txt
Start Height:          0
Index Offset:          80
Jailed Until:          1970-01-01 00:00:00 +0000 UTC
Tombstoned:            false
Missed Blocks Counter: 0
```
