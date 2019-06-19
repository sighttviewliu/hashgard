# hashgardcli stake redelegations

## 描述

基于委托人地址, 查询所有重新委托记录

## 用法

```shell
hashgardcli stake redelegations [delegator-address] [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

基于委托者地址的所有重新委托记录查询

```shell
hashgardcli stake redelegations gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet --trust-node
```

运行成功以后，返回的结果如下：

```json
[
    {
        "delegator_addr": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
        "validator_src_addr": "gardvaloper1m3m4l6g5774qe5jj8cwlyasue22yh32jmhrxfx",
        "validator_dst_addr": "gardvaloper1xn4kvq867rap8vkrwfnp5n2entvpq2avtd0ytq",
        "creation_height": "24800",
        "min_time": "2018-12-21T02:49:44.731658304Z",
        "initial_balance": {
            "denom": "gard",
            "amount": "8"
        },
        "balance": {
            "denom": "gard",
            "amount": "8"
        },
        "shares_src": "8.9100000000",
        "shares_dst": "8.0000000000"
    }
]
```
