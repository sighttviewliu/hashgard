# hashgardcli distribution withdraw-rewards

## 介绍

取回收益

## 用法

```
hashgardcli distribution withdraw-rewards [flags]
```

打印帮助信息:

```
hashgardcli distribution withdraw-rewards --help
```

## 特有的flags

| 名称                | 类型   | 是否必填 | 默认值  | 功能描述        |
| --------------------- | -----  | -------- | -------- | ------------------------------------------------------------------- |
| --only-from-validator | string | false    | ""       | 验证人地址，仅取回在这个验证人上的委托收益 |
| --is-validator        | bool   | false    | false    | 取回验证人佣金收益 |

不能同时使用两个flags。

## 示例

1. 仅取回某一个委托产生的收益
    ```
    hashgardcli distribution withdraw-rewards --only-from-validator gard34mhjjyyc7mehvaay0f3d4hj8qx3ee3w3eq5nq --from mykey  --chain-id=sif-1000
    ```
2. 取回所有委托产生的收益，不包含验证人的佣金收益:
    ```
    hashgardcli distribution withdraw-rewards --from mykey  --chain-id=sif-1000
    ```
3. 取回所有委托产生的收益以及验证人的佣金收益:
    ```
    hashgardcli distribution withdraw-rewards --is-validator=true --from mykey  --chain-id=sif-1000
    ```