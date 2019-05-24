# hashgardcli stake validators

## 描述

查询所有验证人

## 用法

```
hashgardcli stake validators [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

查询验证人

```
hashgardcli stake validators --trust-node
```

运行成功以后，返回的结果如下：

```txt
Validator
Operator Address: gardvaloper1xn4kvq867rap8vkrwfnp5n2entvpq2avtd0ytq
Validator Consensus Pubkey: gardvalconspub1zcjduepq0gsvt058udqf46xwnjmjqd2d9357ec0h3329stlprzky4cu343dsywc93a
Jailed: false
Status: Bonded
Tokens: 111.0000000000
Delegator Shares: 111.0000000000
Description: {hashgard_03   }
Bond Height: 0
Unbonding Height: 0
Minimum Unbonding Time: 1970-01-01 00:00:00 +0000 UTC
Commission: {rate: 0.1000000000, maxRate: 0.2000000000, maxChangeRate: 0.0100000000, updateTime: 0001-01-01 00:00:00 +0000 UTC}

Validator
Operator Address: gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx
Validator Consensus Pubkey: gardvalconspub1zcjduepq7h0hv847a27ck3vmn4ednw5qrsjeykhdg7gnuj28ls5snsallt3svmlckm
Jailed: false
Status: Bonded
Tokens: 89.1000000000
Delegator Shares: 89.1000000000
Description: {instance-c5m0fg87  http://hgdev.com hashgard}
Bond Height: 0
Unbonding Height: 0
Minimum Unbonding Time: 1970-01-01 00:00:00 +0000 UTC
Commission: {rate: 0.1000000000, maxRate: 0.2000000000, maxChangeRate: 0.0100000000, updateTime: 0001-01-01 00:00:00 +0000 UTC}

Validator
Operator Address: gardvaloper1uhauythtet90ewtuy40v4hrymlqf5n45u6tsk4
Validator Consensus Pubkey: gardvalconspub1zcjduepqf5jxe70fqxmzgxvrqmz3zpl8txsgr7fd7g044z3gnq70mhleheds56gtjz
Jailed: false
Status: Bonded
Tokens: 100.0000000000
Delegator Shares: 100.0000000000
Description: {hashgard_01   }
Bond Height: 0
Unbonding Height: 0
Minimum Unbonding Time: 1970-01-01 00:00:00 +0000 UTC
Commission: {rate: 0.1000000000, maxRate: 0.2000000000, maxChangeRate: 0.0100000000, updateTime: 0001-01-01 00:00:00 +0000 UTC}
```
