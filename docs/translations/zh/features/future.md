# Future HRC13 远期支付模块

## 介绍

## 现实场景与问题
Alice 和 Bobo 进行金融交易，Bobo 给 Alice 给出一部分现实中的资产，Alice 承诺在未来的一年内给予 Bobo 证券资产 1000god，每个季度末尾付出 250god。那么 Alice 创建了一个 futurebox，并起名为 Alice 与 Bobo 资产互换一期，对 futurebox 存入 1000 个 god 并设定当年每 3 月 28、6 月 31 日、9 月 30 日、12 月 30 各付出 250god 给 Bobo。远期支付协议的优势，事先存入，一旦约定，将按预定的条件进行执行。不担心任何中途的其他人为意愿或者经济形势的变化而发生改变。充分的使用了区块链技术中不可逆的价值属性。
#### 支付盒子发行期

- 发行方创建盒子，注入利息，可以设置多个时间段支付信息。设定必要参数后，支付账单的总额度=支付盒子注入额度，支付盒子激活并为用户发放带时间的收款凭证。
- 盒子发行期可以注入或者取回通证。

#### 支付期间

- 按照发行者设定的支付时间为将用户的收款凭证兑换为设定的通证。
- 用户获取收款凭证可进行交易。

## 使用

### 发行
```shell
hashgardcli future create [name] [total-amount][distribute-file] [flags]
```

### 注入代币
```shell
hashgardcli future inject [box-id] [amount][flags]
```

### 取消注入
```shell
hashgardcli future cancel [box-id] [amount] [flags]
```

### 查询
```shell
hashgardcli future query [box-id]
```

### 查询列表
```shell
hashgardcli future list
```

### 搜索
```shell
hashgardcli deposit search [name]
```

对于其他查询 Future 状态的命令，请参考[Future](../cli/hashgardcli/future/README.md)
