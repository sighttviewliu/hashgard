s# hashgardcli stake delegation

## 描述

基于委托人和验证人地址查询委托交易

## 用法

```shell
hashgardcli stake delegation [delegator-addr] [validator-addr] [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

查询委托人和验证人的委托交易

```shell
hashgardcli stake delegation gard13nyheuxft7nylrmxmtzewdrs8ukh9r6ejhwvdu gardvaloper13nyheuxft7nylrmxmtzewdrs8ukh9r6eq4rya3 --trust-node
```

运行成功以后，返回结果如下：

```txt
Delegation:
  Delegator: gard13nyheuxft7nylrmxmtzewdrs8ukh9r6ejhwvdu
  Validator: gardvaloper13nyheuxft7nylrmxmtzewdrs8ukh9r6eq4rya3
  Shares:    100000000.000000000000000000
```
