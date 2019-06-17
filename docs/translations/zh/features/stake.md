# Stake 用户手册

## 介绍

本文简要介绍了 stake 模块的功能以及常见用户接口。

## 核心概念

### 投票权重
投票权重是一个共识层面的概念。hub 是一个拜占庭容错的 POS 区块链网络。在共识过程中，一个验证人集将对提案区块进行投票。如果验证人认为提案块有效，它将投赞成票，否则，它将投反对票。来自不同验证人的投票所占权重不同。投票的权重称为相应验证人的投票权重。

### 验证人数量
验证人的数量不是无限的，验证人数量的增加会导致各个节点成本的增加，降低共识效率、增加带宽等。从而降低了整个区块链系统的效率。因此拜占庭容错的 POS 区块链都是有数量上线，根据目前网络的情况，前期确定为 21 个，后期根据社区的治理结果来进行调整增加，至多不超过 100 个验证人节点。

### 活跃验证人节点
验证人节点是全节点且委托的 share 的数量的排位必需是验证人节点数量内。否则只是候选验证人或者普通验证人。

### 候选验证人
被委托但是 share 数量排名低于当前验证人规定数量的排名内。

### 普通节点
仅进行区块链数据记录没有升级成为验证人的节点。

### 委托人
任何人都可以成为委托人，只要将自己持有的 token 委托给验证人就是委托人。委托人可以获取减去验证人设定的佣金后获取相应的份额的 token 奖励。

**验证节点所有者自压的 token 切勿小于最小自抵押数量 `--min-self-delegation`。一旦小于最小自抵押数量 `--min-self-delegation` 该验证人节点将被处于 jailed 状态，该验证人将收不到任何出块和交易奖励， 在该节点上委托代币的投资人的利益也会受到相应的损失。**

### 绑定，解绑
验证人节点的所有者必须将他们自己流通的 token 绑定到自己的验证人节点。且需要达到系统设定的最小自押数。才能成为
验证人节点投票权重与绑定的 token 数量成正比，包括所有者自己绑定的 token 和来自其他委托人的 token。验证人节点的所有者可以通过发送解绑交易来降低他们自己绑定的 token。委托人同样可以通过发送解绑交易来降低绑定的 token。

### 解绑期
委托人对委托的 share 进行解锁，需要等待三周，一旦解绑期结束，token 将返回用户的地址，并可以流通。解绑期 token 没有收益且不能做其他任何操作。解绑期内对 POS 区块链的安全很重要。此外，验证人节点所有者自押的数量不能小于最小自抵押数量，否则会被踢出验证人集。



###  作恶证据和惩罚

拜占庭容错 POS 区块链网络假设拜占庭节点拥有不到总投票权重的 1/3，而且要惩罚这些作恶节点。因此有必要收集作恶行为的证据。根据收集到的证据，stake 模块将从相应的验证人和委托人中拿走一定数量的 token。被拿走的 token 会被销毁。此外，作恶验证人将会被踢出验证人集，并被标记为关押(jailed)状态，而且他们的投票权将立刻变为零。在关押期间，这些节点也不是候选验证人。当关押期结束，他们可以发送 unjail 交易来解除关押状态并再次成为候选验证人。

###  收益
- 作为委托人，向验证人抵押 token 的份额越多，获得的收益就越多。
- 对于验证人节点的所有者，它将有额外的收益：验证人佣金。奖励来自 token 通胀和交易费和服务费。

## 用户操作
### 安装 Hashgard
请参考[安装文档](../learn/installation.md)进行安装。

###  运行全节点
**如果尚未创建钱包，请参考[keys add](../cli/hashgardcli/keys/add.md)进行钱包创建。**

设置客户端进行链接默认节点

```shell
hashgardcli config chain-id ${chain-id}
hashgardcli config trust-node true
```
### 初始化节点并设置节点名称
```shell
hashgard init --chain-id=${chain-id} --moniker=${your_node_name}
```
> 注意：name 仅支持 ASCII 字符。使用 Unicode 字符将使您的节点无法访问。

###  申请成为验证人
#### 方法 1. genisis 阶段成为创世节点

- 向 genesis.json 添加账户信息和资产信息

```shell
hashgard add-genesis-account ${your_wallet} 100000000gard
```
-  向 genesis.json 添加节点信息

```shell
hashgard gentx [flags]
```
请参看[gentx](../cli/hashgad/gentx.md)来设置节点的配置的佣金、和自我委托的创世交易。


#### 方法 2. 发送交易成为验证人

```shell
ashgardcli stake create-validator [flags]
```
请参考[文档](../cli/hashgardcli/stake/create-validator.md)来申请成为验证人。

### 运行节点

```shell
hashgard start
```

### 检查节点

```shell
hashgardcli status
```


###  查询自己的验证人节点
您可以通过运行以下命令找到验证人 pubkey：
```shell
hashgard tendermint show-validator
```

### 编辑节点

```shell
hashgardcli stake edit-validator [flags]
```
对节点的佣金比例、信息等进行编辑

### 查看所有验证人

```shell
hashgardcli stake validators [flags]
```

### 查看验证人信息

```shell
hashgardcli stake validator ${validator-address}
```


### 对验证人进行委托

```shell
hashgardcli stake delegate [validator-addr] [amount] [flags]
```


### 解绑委托
按 share 数量解绑。
```shell
hashgardcli stake unbond [validator-addr] [amount] [flags]
```



### 转委托
按 share 数量转委托至另外节点。委托人可以将其抵押的 token 从一个验证人转移到另一个验证人。

```shell
hashgardcli stake redelegate [src-validator-addr] [dst-validator-addr] [amount] [flags]
```

对于其他查询 stake 状态的命令，请参考[stake](../cli/hashgardcli/stake/README.md)
