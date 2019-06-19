# hashgardcli stake delegations

## 描述

查询某个委托人发起的所有委托记录

## 用法

```shell
hashgardcli stake delegations [delegator-address] [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

查询某个委托人发起的所有委托记录

```shell
hashgardcli stake delegations gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet --trust-node
```

运行成功以后，返回的结果如下：

```json
[
    {
        "delegator_addr": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
        "validator_addr": "gardvaloper1xn4kvq867rap8vkrwfnp5n2entvpq2avtd0ytq",
        "shares": "11.0000000000"
    },
    {
        "delegator_addr": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
        "validator_addr": "gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx",
        "shares": "99.0000000000"
    }
]
```
