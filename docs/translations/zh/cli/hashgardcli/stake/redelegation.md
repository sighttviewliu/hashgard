# hashgardcli stake redelegation

## 描述

基于委托人地址、原验证人地址、目标验证人地址查询重新委托记录

## 用法

```
hashgardcli stake redelegation [delegator-addr] [src-validator-addr] [dst-validator-addr] [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

查询重新委托记录

```
hashgardcli stake redelegation gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx gardvaloper1xn4kvq867rap8vkrwfnp5n2entvpq2avtd0ytq --trust-node
```

运行成功以后，返回的结果如下：

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
