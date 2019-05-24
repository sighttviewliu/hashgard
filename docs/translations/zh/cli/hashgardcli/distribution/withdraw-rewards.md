# hashgardcli distribution withdraw-rewards

## 介绍

从给定的委托地址中取回收益，同时可选择取回验证人佣金收益。

## 用法

```
hashgardcli distribution withdraw-rewards [validator-addr] [flags]
```

## Flags

| 名称                | 类型   | 是否必填 | 默认值  | 功能描述        |
| --------------------- | -----  | -------- | -------- | ------------ |
| --commission | bool | 否 | false  | 取回验证人佣金收益 |

## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

1. 取回委托产生的收益
    ```
    hashgardcli distribution withdraw-rewards gard34mhjjyyc7mehvaay0f3d4hj8qx3ee3w3eq5nq --from mykey --chain-id=hashgard
    ```
2. 取回委托产生的收益以及验证人的佣金收益:
    ```
    hashgardcli distribution withdraw-rewards --commission=true from mykey  --chain-id=sif-1000
    ```
