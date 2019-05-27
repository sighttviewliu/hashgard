# hashgardcli stake redelegation

## Description

Query a redelegation record  for an individual delegator between a source and destination validator:

## Usage

```shell
hashgardcli stake redelegation [delegator-addr] [src-validator-addr] [dst-validator-addr] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example


```shell
hashgardcli stake redelegation gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx gardvaloper1xn4kvq867rap8vkrwfnp5n2entvpq2avtd0ytq --trust-node
```

The result is as follows：

```txt
Redelegation
Delegator: gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet
Source Validator: gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx
Destination Validator: gardvaloper1xn4kvq867rap8vkrwfnp5n2entvpq2avtd0ytq
Creation height: 1130
Min time to unbond (unix): 2018-11-16 07:22:48.740311064 +0000 UTC
Source shares: 0.1000000000
Destination shares: 0.1000000000
```
