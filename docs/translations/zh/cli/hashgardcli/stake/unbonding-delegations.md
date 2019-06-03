# hashgardcli stake unbonding-delegations

## 描述

基于委托人地址查询所有 unbonding-delegation 记录

## 用法

```shell
hashgardcli stake unbonding-delegations [delegator-address] [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

查询委托人的 unbonding-delegation

```shell
hashgardcli stake unbonding-delegations faa13lcwnxpyn2ea3skzmek64vvnp97jsk8qmhl6vx
```

运行成功以后，返回的结果如下：

```json
[
    {
        "delegator_addr": "faa13lcwnxpyn2ea3skzmek64vvnp97jsk8qmhl6vx",
        "validator_addr": "fva15grv3xg3ekxh9xrf79zd0w077krgv5xf6d6thd",
        "creation_height": "1310",
        "min_time": "2018-11-15T06:24:22.754703377Z",
        "initial_balance": "0.02hashgard",
        "balance": "0.02hashgard"
    }
]
```
