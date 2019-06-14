# hashgardcli exchange cancel

## 描述

某笔有效订单的 seller 可以撤销该笔订单，冻结的资金将返回其账户

## 用法

```shell
hashgardcli exchange cancel [order_id] [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

### 撤销订单

```shell
hashgardcli exchange cancel 2 --from mykey --chain-id hashgard -o=json --indent
```

必须是订单的 seller 账户操作，输入正确的密码后，order_id 为 2 的订单已经撤销。

```txt
{
 "height": "7162",
 "txhash": "7FA99BF4C271E8145A8DB695B9B08883A58A46F1AA369D8A7F6002684FDBC06A",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "200000",
 "gas_used": "27608",
 "tags": [
  {
   "key": "action",
   "value": "withdrawal_order"
  },
  {
   "key": "order_id",
   "value": "2"
  },
  {
   "key": "seller",
   "value": "gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa"
  },
  {
   "key": "order_status",
   "value": "inactive"
  }
 ]
}
```
