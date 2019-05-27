# hashgardcli stake unbonding-delegation

## 描述

基于委托人地址和验证人地址查询 unbonding-delegation 记录

## 用法

```shell
hashgardcli stake unbonding-delegation [delegator-addr] [validator-addr] [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

查询委托人和验证人的 unbonding-delegation

```shell
hashgardcli stake unbonding-delegation gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx --chain-id=hashgard
```

运行成功以后，返回的结果如下：

```txt
Unbonding Delegation
Delegator: gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet
Validator: gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx
Creation height: 12610
Min time to unbond (unix): 2018-12-20 08:07:17.286706585 +0000 UTC
Expected balance: 9gard

```
