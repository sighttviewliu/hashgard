# hashgardcli distribution params

## 描述

查询给定块范围的验证人的所有处罚

## 用法

```shell
hashgardcli distribution slashes [validator] [start-height] [end-height] [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

查询参数信息

```shell
hashgardcli distribution slashes gardvaloper1hr7vm7t7paeyg33ggd6efek2w58mu2huewltta 0 999999 -o=json --trust-node
```

运行成功以后，返回的结果如下：

```txt
[
 {
  "validator_period": "4",
  "fraction": "0.010000000000000000"
 },
 {
  "validator_period": "5",
  "fraction": "0.010000000000000000"
 },
 {
  "validator_period": "17",
  "fraction": "0.010000000000000000"
 }
]
```
