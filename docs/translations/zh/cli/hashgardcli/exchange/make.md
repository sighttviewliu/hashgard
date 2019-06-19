# hashgardcli exchange make

## 描述

创建订单, 指定提供的币种和数量，以及目标币种及数量。创建成功后，订单处于有效状态，提供的币会处于冻结状态。任何人都能同该订单进行交易。

## 用法

```shell
hashgardcli exchange create [flags]
```

## Flags

| 名称       | 类型                  | 必需        | 默认值            | 描述      |
| --------------- | ---------------- | -------------- | --------- | ------------- |
| --supply     | string | 是 || 订单用于交易的币种及数量             |
| --target        | string | 是 || 订单的目标币种及数量           |

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

### 创建订单

```shell
hashgardcli exchange make --supply 100gard --target 800apple --from mykey --chain-id hashgard --indent -o=json
```

输入正确的密码后，创建了一笔 100gard 交换 800apple 的订单。

```txt
{
 "height": "6907",
 "txhash": "12FE1C3ECDF960AFB6A5E2D1A0DCE678EDB9A35812137AFA2CFA0DF7340C8F12",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "200000",
 "gas_used": "35757",
 "tags": [
  {
   "key": "action",
   "value": "make"
  },
  {
   "key": "order_id",
   "value": "1"
  },
  {
   "key": "seller",
   "value": "gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa"
  },
  {
   "key": "supply_token",
   "value": "gard"
  },
  {
   "key": "target_token",
   "value": "apple"
  }
 ]
}
```

如何查询订单，查看某地址的所有订单，查看某地址冻结的资金及撤销订单？

请点击下述链接：

[query](query.md)
[list](list.md)
[query-frozen](query-frozen.md)
