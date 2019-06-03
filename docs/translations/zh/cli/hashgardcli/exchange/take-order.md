# hashgardcli exchange create-order

## 描述

创建订单, 指定提供的币种和数量，以及目标币种及数量。创建成功后，订单处于有效状态，提供的币会处于冻结状态。任何人都能同该订单进行交易。

## 用法

```shell
hashgardcli exchange take-order [order_id] [flags]
```

## Flags

| 名称       | 类型                  | 必需          | 默认值            | 描述         |
| --------------- | -------------- | ------------------ | --------------------- | --------------- |
| --amount     | string | 是 || 买方用于交易的币种及数量                                                                                    |                                                       |
**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)
## 计算

A(supply) 为 supply 的数量，A(target) 为 target 的数量， A(remains) 是 remains 的数量（订单创建时，A(supply) = A(target) ）

divisor 是 A(supply) 与 A(target) 的最大公约数

上述的数量全都是整数，如果把订单的 supply 拆分成每份相等的整数份额，最多能拆成的份数即为 divisor,

每份的价格分别以 supply，target 计为：

Price(supply) = A(supply) / divisor
Price(target) = A(target) / divisor

当前订单可供交易的份数为 ：

Shares(remains) = A(remains) / Price(supply)


当对一个挂单进行交易时，发送的币(币种要和 targe 一致)数量为 Amount

1. 首先判断 Amount 是否超过 Price(target)，如果低于，则不够交易门槛。
2. 计算成交份额，Shares(afford) = Amount / Price(target)
3. 计算实际成交份额，如果 Shares(afford) < Shares(remains), 则 Shares(actual) = Shares(afford) , 如果 >= , 则 Shares(actual) = Shares(remains)
4. 然后，计算实际支付的 target 和从订单 remains 中扣除的 supply，Sum(target) = Shares(actual) * Price(target)，
这部分金额会从购买者的 Amount 中扣除，并加到 seller 的账户里，如果 Amount > Sum(target), 多余的部分会返还给购买者。
Sum(supply) = Shares(actual) * Price(supply), 这部分的金额会从订单的 A(remains) 中扣除， 并加到买家的账户里。
5. 如果最后 A(remains) 没有剩余，这该挂单被吃光，删除该挂单。



## 例子

### 进行交易

已知 order_id 为 3 的订单的 supply 是 100gard， target 是 800apple，remains 是 100gard。
使用 18apple 与之交易， 实际扣除 16 apple， 会获得 2 gard。

```shell
hashgardcli exchange take-order 3 --amount 18apple --from mykey --chain-id hashgard -o=json --indent
```

输入正确的密码后，同 order_id 为 3 的订单交易。

```txt
{
 "height": "145",
 "txhash": "9560252F7F887A2DFFA30B5FC7C35BCC6B93608877590207F127595E5CFE7897",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "200000",
 "gas_used": "42606",
 "tags": [
  {
   "key": "action",
   "value": "take_order"
  },
  {
   "key": "order_id",
   "value": "3"
  },
  {
   "key": "buyer",
   "value": "gard19at42yyn62hr5e2xdze90al3g693932ha0kz38"
  },
  {
   "key": "order_status",
   "value": "active"
  }
 ]
}
```

查询订单现在的状态：

买家账户扣除了 16apple， 获得 2gard，order_id 为 3 的订单 remains 更新为 98gard
