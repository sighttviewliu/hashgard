# hashgardcli distribution rewards

## 描述

查询委托人获得的所有奖励，可选择限制来自单个验证者的奖励

## 用法

```
hashgardcli distribution rewards [delegator-addr] [<validator-addr>] [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

查询参数信息

```
hashgardcli distribution rewards gard1hr7vm7t7paeyg33ggd6efek2w58mu2hutvjrms -o=json --trust-node
hashgardcli distribution rewards gard1hr7vm7t7paeyg33ggd6efek2w58mu2hutvjrms gardvaloper1hr7vm7t7paeyg33ggd6efek2w58mu2huewltta -o=json --trust-node
```

运行成功以后，返回的结果如下：

```
[
 {
  "denom": "gard",
  "amount": "0.131833867963517125"
 }
]
```
