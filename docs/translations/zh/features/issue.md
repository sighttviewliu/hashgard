# HRC10 同质化通证发行协议 （issue模块）

## 简介
issue 是Hashgard网络中通证发行模块。只要用户具有Hashgard主网代币GARD，都可以发行自己的资产。不需要懂得复杂的代码，即可迅速的发行安全的区块链资产。原生的数字资产具有很多的特性，包括资产描述，增发，冻结/解冻，转移代币管理权限。



## 使用场景


### 发行
```shell
hashgardcli issue create [name] [symbol] [total-supply] [flags]
```
以发行总数量为100亿的BTC

```shell
hashgardcli issue create bitcoin BTC 10000000000000 --from

```
详细信息请参阅 [issue create](../cli/hashgardcli/issue/create.md)

发行代币的返回信息中会有代币的issue-id，根据issue-id来做以下的操作。

### 为代币添加描述信息

```shell
hashgardcli issue describe [issue-id] [description-file] [flags]
```
desription-file
描述信息文件，格式要求是json，文件大小不能超过1024字节。可以按照官方模版也可以自己按格式书写。
详细信息请参阅[issue describe](../cli/hashgardcli/issue/describe.md)

### 查询账户
可以通过账户地址查询账户的余额、

```shell
hashgardcli bank account  $YouWalletAddress
```
输入命令和你的账户地址即可查询

详细信息请参阅[bank account](../cli/hashgardcli/bank/account.md)

### 交易

```shell
hashgardcli bank send [to_address] [amount] [flags]
```
输入命令和你要转出的地址和数量，调取[flags]中的`--from`来调取钱包

```shell
hashgardcli bank send gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0  20000coin174876e800 --from one
```
以本地one 钱包为例子给gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0 发送20000个coin74876e800
详细信息请参阅[bank send](../cli/hashgardcli/bank/send.md)

### 给代币进行增发

```shell
hashgardcli issue mint [issue-id] [amount][flags]
```
增发可以增发给自己也可以增发给指定的账户。
详细信息请参阅[issue mint](../cli/hashgardcli/issue/mint.md)
给代币进行增发的必要条件: 1.必须具有代币的owner权限。2.代币的增发功能没有被禁用。

### 冻结账号
作为代币的owner，为了维护整个经济系统的公正与繁荣以及惩罚某些生态中的作恶行为。系统默认是开启owner冻结功能。用户可以选择自行关闭。该功能在初次发行中设置。或者在后期使用 [issue disable](../cli/hashgardcli/issue/disable.md)进行冻结功能关闭。该功能关闭后不可逆，不影响解冻功能。
冻结功能分为1.冻结转入。 2.冻结转出 3.冻结转入和转出。

```shell
 hashgardcli issue freeze [freeze-type] [issue-id][acc-address][end-time] --from
```
以冻结 gardkenrwk5k4ng70e5s9zfsttxpnlesx5ps0gfdv7 Apple（coin74876e800）代币 账号转入和转出功能 至公元2100-1-1 01:01:01为例
```shell
 hashgardcli issue freeze in-out coin74876e800 gardkenrwk5k4ng70e5s9zfsttxpnlesx5ps0gfdv7 4102419661 --from WalletName
```
详细信息请参阅[issue freeze](../cli/hashgardcli/issue/freeze.md)
冻结中需要设置时间，时间格式为Unix格式 时间转换工具请参阅[Unix timestamp](./Unix-timestamp.md)



### 解冻账号
在被该代币的Owner冻结了转账功能后，owner解冻用户地址该代币的转账功能。

```shell
 hashgardcli issue unfreeze [unfreeze-type] [issue-id][address] [Flags]
```

### 燃烧代币
燃烧代币，会使代币总量减少

### 关闭增发

### 关闭冻结

### 关闭燃烧
