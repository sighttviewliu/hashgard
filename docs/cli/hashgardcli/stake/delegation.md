# hashgardcli stake delegation

## Description

Query delegations for an individual delegator on an individual validator:


## Usage

```shell
hashgardcli stake delegation [delegator-addr] [validator-addr] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example


```shell
hashgardcli stake delegation gard13nyheuxft7nylrmxmtzewdrs8ukh9r6ejhwvdu gardvaloper13nyheuxft7nylrmxmtzewdrs8ukh9r6eq4rya3 --trust-node
```

The result is as follows：

```txt
Delegation:
  Delegator: gard13nyheuxft7nylrmxmtzewdrs8ukh9r6ejhwvdu
  Validator: gardvaloper13nyheuxft7nylrmxmtzewdrs8ukh9r6eq4rya3
  Shares:    100000000.000000000000000000
```
