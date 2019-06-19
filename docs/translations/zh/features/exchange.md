# Atomic Swap Protocol HRC11 原子交易模块 (Exchange)

## 介绍

Hashgard 为用户提供原生的原子交换系统。可以在点对点的基础上实现加密货币的交换，借助整个 Hashgard 网络来进行交换，无需受信任的第三方，也不存在交易一方在交易中违约的风险。

## 现实场景与问题

现实生活中，譬如最早期的游戏道具交易中，Alice 和 Bobo 两位同学在游戏中进行交易。最早的游戏的交易规则是双方都点击确认交易按钮后交易成功。 Alice 手上有 500 个金币，Bobo 手上有一把法杖 A。Alice 和 Bobo 商议 500 金币可以购买 Bobo 手中的法杖 A。在交易窗弹出的时候 Alice 输入 500 金币，Bobo 把法杖 Alice 放入到了交易栏中。Bobo 看到 Alice 输入的是 500 金币 于是点击确认。这时候 Alice 将金币输入为 1 并点击确认。最后 Alice 以 1 枚金币的价格购买到了 法杖 A。后来游戏公司为了避免这种骗局的发生，增加交易的流程来避免这种欺诈问题。


原子交换协议将不存在这种问题。简单举个例子

2 个互不相信的人想要完成一桩交易，比如红宝石换金币，他们通过 2 个保险箱达成交易。Bobo 在 1 号保险箱中放入给 Alice 的红宝石，设定保险箱只有 Alice 才能打开，并且设定保险箱密码为他选定的值，这个值 Alice 并不知道。Alice 在 2 号保险箱中放入给 Bobo 的金币，设定保险箱只有 Bobo 才能打开，并且在不了解保险箱密码锁的情况下克隆了 1 号保险箱的密码锁。
Bobo 因为知道保险箱的密码，所以 Bobo 可以打开 2 号保险箱，取走金币。2 号保险箱的密码锁打开后就留下了密码。
Alice 现在知道了密码锁的密码，他就可以去 1 号保险箱取走红宝石。
如果 Bobo 拒绝打开 2 号保险箱，Alice 虽然无法拿到红宝石，但 Bobo 自己也拿不到金币。



## 使用

### 用户创建卖单
```shell
hashgardcli exchange make[flags]
```

### 用户创建买单
```shell
hashgardcli exchange take [order_id] [flags]
```
### 用户取消订单
```shell
hashgardcli exchange cancel [order_id] [flags]
```
### 查询指定订单
```shell
hashgardcli exchange query [order_id] [flags]
```
### 查询订单列表
```shell
hashgardcli exchange list [address] [flags]
```
### 查询指定地址冻结资金
```shell
hashgardcli exchange query-frozen [address] [flags]
```
对于其他查询 exchange 状态的命令，请参考[exchange](../cli/hashgardcli/exchange/README.md)
