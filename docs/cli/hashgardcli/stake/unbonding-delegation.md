# hashgardcli stake unbonding-delegation

## Description

Query unbonding delegations for an individual delegator on an individual validator:

## Usage

```shell
hashgardcli stake unbonding-delegation [delegator-addr] [validator-addr] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example



```shell
hashgardcli stake unbonding-delegation gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx --chain-id=hashgard
```

The result is as follows：

```txt
Unbonding Delegation
Delegator: gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet
Validator: gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx
Creation height: 12610
Min time to unbond (unix): 2018-12-20 08:07:17.286706585 +0000 UTC
Expected balance: 9gard

```
