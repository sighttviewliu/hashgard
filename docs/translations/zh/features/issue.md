# Fungible Token Issuance Protocol HRC10 同质化通证发行协议

## 简介

issue 是 Hashgard 网络中通证发行模块。只要用户具有 Hashgard 主网代币 GARD，都可以发行自己的资产。不需要懂得复杂的代码，即可迅速的发行安全的区块链资产。原生的数字资产具有很多的高级特性，包括资产描述，增发，冻结/解冻，转移代币管理权限。

## 使用场景
为用户建立完整资产发行及管理模块。

### 发行

```shell
hashgardcli issue create [name] [symbol] [total-supply] [flags]
```

以发行总数量为 100 亿的 BTC

```shell
hashgardcli issue create bitcoin BTC 10000000000000 --from
```

详细信息请参阅 [issue create](../cli/hashgardcli/issue/create.md)

发行代币的返回信息中会有代币的 issue-id，根据 issue-id 来做以下的操作。

### 为代币添加描述信息

```shell
hashgardcli issue describe [issue-id] [description-file] [flags]
```

desription-file
描述信息文件，格式要求是 json，文件大小不能超过 1024 字节。可以按照官方模版也可以自己按格式书写。
详细信息请参阅[issue describe](../cli/hashgardcli/issue/describe.md)

### 查询账户

可以通过账户地址查询账户的余额、

```shell
hashgardcli bank account  $YouWalletAddress
```

输入命令和你的账户地址即可查询

详细信息请参阅[bank account](../cli/hashgardcli/bank/account.md)

### 交易

```shell
hashgardcli bank send [to_address] [amount] [flags]
```

输入命令和你要转出的地址和数量，调取[flags]中的`--from`来调取钱包

```shell
hashgardcli bank send gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0  20000coin174876e800 --from one
```

以本地 one 钱包为例子给 gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0 发送 20000 个 coin74876e800
详细信息请参阅[bank send](../cli/hashgardcli/bank/send.md)

### 给代币进行增发

```shell
hashgardcli issue mint [issue-id] [amount][flags]
```

增发可以增发给自己也可以增发给指定的账户。
详细信息请参阅[issue mint](../cli/hashgardcli/issue/mint.md)
给代币进行增发的必要条件: 1.必须具有代币的 owner 权限。2.代币的增发功能没有被禁用。

### 冻结账号

作为代币的 owner，为了维护整个经济系统的公正与繁荣以及惩罚某些生态中的作恶行为。系统默认是开启 owner 冻结功能。用户可以选择自行关闭。该功能在初次发行中设置。或者在后期使用 [issue disable](../cli/hashgardcli/issue/disable.md)进行冻结功能关闭。该功能关闭后不可逆，不影响解冻功能。
冻结功能分为
- 冻结转入。
- 冻结转出。
- 冻结转入和转出。

```shell
 hashgardcli issue freeze [freeze-type] [issue-id][acc-address][end-time] [flags]
```

以冻结 gardkenrwk5k4ng70e5s9zfsttxpnlesx5ps0gfdv7 Apple（coin74876e800）代币 账号转入和转出功能 至公元 2100-1-1 01:01:01 为例

```shell
 hashgardcli issue freeze in-out coin74876e800 gardkenrwk5k4ng70e5s9zfsttxpnlesx5ps0gfdv7 4102419661 --from WalletName
```

详细信息请参阅[issue freeze](../cli/hashgardcli/issue/freeze.md)
冻结中需要设置时间，时间格式为 Unix 格式 时间转换工具请参阅[Unix timestamp](./Unix-timestamp.md)



### 解冻账号

在被该代币的 Owner 冻结了转账功能后，owner 解冻用户地址该代币的转账功能。

```shell
 hashgardcli issue unfreeze [unfreeze-type] [issue-id][address] [Flags]
```

### 燃烧代币

燃烧代币，会使代币总量减少。
```shell
hashgardcli issue burn [issue-id] [amount] [flags]
```


### 关闭增发

```shell
hashgardcli issue disable [issue-id] --minting [flags]
```

### 关闭冻结
```shell
hashgardcli issue disable [issue-id] --freeze [flags]
```

### 关闭销毁代币
关闭代币销毁分为 3 类:
1. Owner 销毁自己持有的可用代币。

```shell
hashgardcli issue disable [issue-id] --burn-owner [flags]
```
2. 非 owner 销毁自己持有的可用代币。

```shell
hashgardcli issue disable [issue-id] --burn-holder [flags]
```
3. Owner 销毁非 Owner 持有的可用代币。

```shell
hashgardcli issue disable [issue-id] --burn-from [flags]
```

对于其他查询 issue 状态的命令，请参考[issue](../cli/hashgardcli/issue/README.md)
