# hashgardcli exchange query-orders

## 描述

查看指定地址的所有有效订单

## 用法

```shell
hashgardcli exchange query-orders [address] [flags]
```

## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

### 查询地址所有的有效订单

```shell
hashgardcli exchange query-orders gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa --chain-id hashgard -o=json --indent
```

下面是地址 gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa 所有有效的订单

```txt
[
 {
  "order_id": "1",
  "seller": "gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa",
  "supply": {
   "denom": "gard",
   "amount": "100"
  },
  "target": {
   "denom": "apple",
   "amount": "800"
  },
  "remains": {
   "denom": "gard",
   "amount": "100"
  },
  "create_time": "2019-04-16T07:12:39.254071Z"
 },
 {
  "order_id": "2",
  "seller": "gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa",
  "supply": {
   "denom": "apple",
   "amount": "33"
  },
  "target": {
   "denom": "horn",
   "amount": "10000"
  },
  "remains": {
   "denom": "apple",
   "amount": "33"
  },
  "create_time": "2019-04-16T07:27:31.379469Z"
 }
]
```
