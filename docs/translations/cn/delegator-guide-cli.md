# 委托人指南 (CLI)


本文介绍了如何使用 Hashgard 的命令行交互（CLI）程序实现通证委托的相关知识和操作步骤。 

同时，本文也介绍了如何管理账户，如何从筹款人那里恢复账户，以及如何使用一个硬件钱包的相关知识。 

> 风险提示
>
> **重要提示**：请务必按照下面的操作步骤谨慎操作，过程中发生任何错误都有可能导致您永远失去所拥有的通证。
因此，请在开始操作之前先仔细阅读全文，如果有任何问题可以联系我们获得支持。 
> 
> 请务必谨慎行事！

## 目录
    
- [安装 `hashgardcli`](#安装-hashgardcli)
- [Hashgard 账户](#Hashgard账户)
    + [创建一个账户](#创建一个账户)
- [访问 Hashgard 网络](#访问Hashgard网络)
    + [运行您自己的全节点](#运行您自己的全节点)
    + [连接到一个远程全节点](#连接到一个远程全节点)
- [设置`hashgardcli`](#设置-hashgardcli)
- [状态查询](#状态查询)
- [发起交易](#发起交易)
    + [关于gas费和手续费](#关于gas费和手续费)
    + [抵押 GARD 通证 & 提取奖励](#抵押GARD通证--提取奖励)
    + [参与链上治理](#参与链上治理)
    + [从一台离线电脑上签署交易](#从一台离线电脑上签署交易)

## 安装 `hashgardcli` 

`hashgardcli`: 与`hashgard`全节点交互的命令行用户界面。 

> 安全提示
>
> **请检查并且确认你下载的`hashgardcli`是可获得的最新稳定版本**


[**下载已编译代码**]暂不提供


[**安装 Hashgard**](./installation.md)

::: tip 提示 

`hashgardcli` 需要通过操作系统的终端窗口使用，打开步骤如下所示： 
- **Windows**: `开始` > `所有程序` > `附件` > `终端`
- **MacOS**: `访达` > `应用程序` > `实用工具` > `终端`
- **Linux**: `Ctrl` + `Alt` + `T`
:::

## Hashgard账户

每个 Hashgard 账户的核心基础是一个包含24个词的助记词组，通过这个助记词可以生成无数个 Hashgard 账户，
例如，一组私钥/公钥对。这被称为一个硬件钱包（跟多硬件钱包相关说明请参见[BIP32](https://github.com/bitcoin/bips/blob/master/bip-0032.mediawiki)）  

```
        账户 0                            账户 1                              账户 2

+------------------+              +------------------+               +------------------+
|                  |              |                  |               |                  |
|       地址 0      |              |      地址 1      |               |       地址 2      |
|        ^         |              |        ^         |               |        ^         |
|        |         |              |        |         |               |        |         |
|        |         |              |        |         |               |        |         |
|        |         |              |        |         |               |        |         |
|        +         |              |        +         |               |        +         |
|       公钥 0      |              |      公钥 1      |               |       公钥 2      |
|        ^         |              |        ^         |               |        ^         |
|        |         |              |        |         |               |        |         |
|        |         |              |        |         |               |        |         |
|        |         |              |        |         |               |        |         |
|        +         |              |        +         |               |        +         |
|       私钥 0      |              |      私钥 1      |               |       私钥 2      |
|        ^         |              |        ^         |               |        ^         |
+------------------+              +------------------+               +------------------+
         |                                 |                                  |
         |                                 |                                  |
         |                                 |                                  |
         +--------------------------------------------------------------------+
                                           |
                                           |
                                 +---------+---------+
                                 |                   |
                                 |  助记词 (Seed)     |
                                 |                   |
                                 +-------------------+
```


私钥是控制一个账户中所存资产的钥匙。私钥是通过助记词单向产生的。如果您不小心丢失了私钥，你可以通过助记词恢复。
然而，如果你丢失了助记词，那么你就有可能失去对由这个助记词产生的所有私钥的控制。
同样，如果有人获得了你的助记词，他们就可以操作所有相关账户。  

::: 警告

**谨慎保管并不要告诉他人你的助记词。 为了防止资产被盗或者丢失，您最好多备份几份助记词，
并且把它们存放在只有您知道的安全地方，这样做将有助于保障您的私钥以及相关账户的安全。**
:::


Hashgard 地址是一个以可读词开头的字符串(比如`gard10snjt8dmpr5my0h76xj48ty80uzwhraqalu4eg`) 
如果有人想给你转账通证，他们就往这个地址转账。根据给定地址来推算私钥是不可行的。

### 创建一个账户

前，您需要先安装`hashgardcli`，同时，您需要知道你希望在哪里保存和使用您的私钥。
最好的办法是把他们保存在一台没有上网的电脑或者一个硬件钱包设备里面。
将私钥保存在一台联网的电脑里面比较危险，任何人通过网络攻击都有可能获取您的私钥，进而盗取您的资产。 

#### 使用电脑进行操作

::: 警告

**注意：在一台没有联网的电脑上操作会更加安全**
:::

然后，通过以下命令创建账户：

```bash
hashgardcli keys add <yourKeyName>
```

这个命令会产生一个24个词的助记词组，并且同时保存账户 `0` 的私钥和公钥。
另外，您还需要输入一个密码来对您硬盘上账户`0`的私钥进行加密。
每次您发出一笔交易时都将需要输入这个密码。如果您丢失了密码，您可以通过助记词来恢复您的私钥。 

::: 危险提示

**千万不要丢失或者告诉其他人你的12个词的助记词组。
为了防止资产被盗或者丢失，您最好多备份几份助记词，并且把它们存放在只有您知道的安全地方，
如果有人取得您的助记词，那么他也就取得了您的私钥并且可以控制相关账户。**
::: 

::: 警告
在确认已经安全保存好您的助记词以后（至少3遍！），你可以用如下命令清除终端窗口中的命令历史记录，以防有人通过历史记录获得您的助记词。 

```bash
history -c
rm ~/.bash_history
```
:::

你可以用以下命令使用助记词生成多个账户： 

```bash
hashgardcli keys add <yourKeyName> --recover --account 1
```

- - `<yourKeyName>` 是账户名称，用来指代用助记词生成私钥/公钥对的 Hashgard 账户。
在您发起交易时，这个账户名称用来识别您的账户。 
  
- 您可以通过增加 `--account` 标志来指定您账户生成的路径 (`0`, `1`, `2`, ...)， `0` 是默认值。


这条命令需要您输入一个密码。改变账号就代表生成了一个新的账户。  

## 访问Hashgard网络


为了查询网络状态和发起交易，你需要通过自建一个全节点，或者连接到其他人的全节点访问 Hashgard 网络

::: 警告

**注意： 请不要与任何人分享您的助记词，您是唯一需要知道这些助记词的人。
如果任何人通过邮件或者其他社交媒体向您要求您提供您的助记词，那就需要警惕了。
请记住，Hashgard 团队，Hashgard 基金会永远不会通过邮件要求您提供您的账户密码或助记词。**
::: 

### 运行您自己的全节点


这是最安全的途径，但需要相当多的资源。您需要有较大的带宽和至少1TB的磁盘容量来运行一个全节点。  


 `hashgard`的安装教程在[这里](./installation.md)， 如何运行一个全节点指导在[这里](./join-testnet.md)

### 连接到一个远程全节点


如果您不想或没有能力运行一个全节点，您也可以连接到其他人的全节点。
您需要谨慎的选择一个可信的运营商，因为恶意的运营商往往会向您反馈错误的查询结果，或者对您的交易进行监控。
然而，他们永远也无法盗取您的资产，因为您的私钥仅保持在您的本地计算机或者钱包设备中。
验证人，钱包供应商或者交易所是可以提供全节点的运营商。 


连接到其他人提供的全节点，你需要一个全节点地址，如`https://77.87.106.33:26657`。
这个地址是您的供应商提供的一个可信的接入地址。你会在下一节中用到这个地址。 

## 设置 `hashgardcli`

::: 提示

**在开始设置 `hashgardcli`前，请确认你已经可以[访问Hashgard网络](#访问Hashgard网络)**
:::

::: 警告

**请确认您使用的`hashgardcli`是最新的稳定版本**
:::

无论您是否在自己运行一个节点，`hashgardcli` 都可以帮您实现与 Hashgard 网络的交互。让我们来完成对它的配置。 


您需要用下面的命令行完成对`hashgardcli`的配置：

```bash
hashgardcli config <flag> <value>
```


此命名允许您为每个参数设置默认值。 


首先，设置你想要访问的全节点的地址：

```bash
hashgardcli config node <host>:<port

// 样例: hashgardcli config node https://77.87.106.33:26657
```


如果你是自己运行全节点，可以使用 `tcp://localhost:26657` 作为地址。 


然后，让我们设置 `--trust-node` 指标的默认值。 

```bash
hashgardcli config trust-node false

// Set to true if you run a light-client node, false otherwise
```

最后，让我们设置需要访问区块链的 `chain-id` 

```bash
hashgardcli config chain-id sif-7000
```

## 状态查询

::: 提示
** 准备抵押 GARD 通证和取回奖励前，需要先完成 [`hashgardcli` 配置](#设置-hashgardcli)**
:::


`hashgardcli` 可以帮助您获得所有区块链的相关信息， 比如账户余额，抵押通证数量，奖励，治理提案以及其他信息。
下面是一组用于委托操作的命令

```bash

// 查询账户余额或者其他账户相关信息
hashgardcli query account


// 查询验证人列表
hashgardcli query staking validators


// 查询指定地址的验证人的信息(e.g. gardvaloper1n5pepvmgsfd3p2tqqgvt505jvymmstf6s9gw27)
hashgardcli query staking validator <validatorAddress>


//查询指定地址的验证人相关的所有委托信息 (e.g. gard10snjt8dmpr5my0h76xj48ty80uzwhraqalu4eg)
hashgardcli query staking delegations <delegatorAddress>

// 查询从一个指定地址的委托人(e.g. gard10snjt8dmpr5my0h76xj48ty80uzwhraqalu4eg)和一个指定地址的验证人(e.g. gardvaloper1n5pepvmgsfd3p2tqqgvt505jvymmstf6s9gw27) 之间的委托交易
hashgardcli query staking delegation <delegatorAddress> <validatorAddress>

// 查询一个指定地址的委托人(e.g. gard10snjt8dmpr5my0h76xj48ty80uzwhraqalu4eg)所能获得的奖励情况
hashgardcli query distribution rewards <delegatorAddress> 

// 查询所有现在正等待抵押的提案
hashgardcli query gov proposals --status deposit_period

//查询所有现在正等待投票的填
hashgardcli query gov proposals --status voting_period

// 查询一个指定propsalID的提案信息
hashgardcli query gov proposal <proposalID>
```

需要了解跟多的命令，只需要用：

```bash
hashgardcli query
```

对于每条命令，您都可以使用`-h` 或 `--help` 参数来获得更多的信息。 

## 发起交易

### 关于gas费和手续费


Hashgard 网络上的交易在被执行时需要支付手续费。这个手续费是用于支付执行交易所需的gas。计算公式如下： 
```
fees = gas * gasPrices
```

`gas` 的多少取决于交易类型，不同的交易类型会收取不同的 `gas` 。每个交易的 `gas` 费是在实际执行过程中计算的，
但我们可以通过设置 `gas` 参数中的 `auto` 值实现在交易前对 `gas` 费的估算，但这只是一个粗略的估计，
你可以通过 `--gas-adjustment` (缺省值 `1.0`) 对预估的`gas` 值进行调节，以确保为交易支付足够的`gas` 。 

`gasPrice` 用于设置单位 `gas` 的价格。每个验证人会设置一个最低gas价`min-gas-price`, 
并只会将`gasPrice`大于`min-gas-price`的交易打包。 

交易的`fees` 是`gas` 和 `gasPrice`的乘积。作为一个用户，你需要输入3个参数中至少2个，
`gasPrice`/`fees`的值越大，您的交易就越有机会被打包执行。 

### 抵押Atom通证 & 提取奖励

::: 提示
**在您抵押通证或取回奖励之前，您需要完成[`hashgardcli` 设置](#设置-hashgardcli) 
和 [创建账户](#创建一个账户)**
:::

::: 警告
**注意：执行以下命令需要在一台联网的计算机。用一个硬件钱包设备执行这些命令会更安全。
关于离线交易过程请看[这里](#从一台离线电脑上签署交易).**
::: 

```bash
// 向指定验证人绑定一定数量的Atom通证
// 参数设定样例: <validatorAddress>=gardvaloper18thamkhnj9wz8pa4nhnp9rldprgant57pk2m8s, <amountToBound>=10000000000gard, <gasPrice>=1000gard

hashgardcli tx staking delegate <validatorAddress> <amountToBond> --from <delegatorKeyName> --gas auto --gas-prices <gasPrice>


// 提取所有的奖励
// 参数设定样例: <gasPrice>=1000gard

hashgardcli tx distribution withdraw-all-rewards --from <delegatorKeyName> --gas auto --gas-prices <gasPrice>


// 向指定验证人申请解绑一定数量的Atom通证
// 解绑的通证需要3周后才能完全解绑并可以交易，
// 参数设定样例: <validatorAddress>=gardvaloper18thamkhnj9wz8pa4nhnp9rldprgant57pk2m8s, <amountToUnbound>=10000000000gard, <gasPrice>=1000gard

hashgardcli tx staking unbond <validatorAddress> <amountToUnbond> --from <delegatorKeyName> --gas auto --gas-prices <gasPrice>
```

::: 提示
::: 
如果您是使用一个联网的钱包设备，在交易被广播到网络前您需要在设备上确认交易。 

确认您的交易已经发出，可以用以下查询：

```bash
// 您的账户余额在您抵押Atom通证或者取回奖励后会发生变化
hashgardcli query account

// 您在抵押后应该能查到委托交易
hashgardcli query staking delegations <delegatorAddress>

// 如果交易已经被打包，将会返回交易记录（tx）
// 在以下查询命令中可以使用显示的交易哈希值作为参数
hashgardcli query tx <txHash>

```

如果您是连接到一个可信全节点的话，您可以通过一个区块链浏览器做检查。 

## 参与链上治理

#### 链上治理入门

Hashgard 有一个内建的治理系统，该系统允许抵押通证的持有人参与提案投票。系统现在支持3种提案类型：

- `Text Proposals`: 这是最基本的一种提案类型，通常用于获得大家对某个网络治理意见的观点。 
- `Parameter Proposals`: 这种提案通常用于改变网络参数的设定。 
- `Software Upgrade Proposal`: 这个提案用于升级 Hashgard 的软件。 

任何 GARD 通证的持有人都能够提交一个提案。为了让一个提案获准公开投票，提议人必须要抵押一定量的通证 `deposit`，
且抵押值必须大于 `minDeposit` 参数设定值. 提案的抵押不需要提案人一次全部交付。
如果早期提案人交付的 `deposit` 不足，那么提案进入 `deposit_period` 状态。 
此后，任何通证持有人可以通过 `depositTx` 交易增加对提案的抵押。 

当`deposit` 达到 `minDeposit`，提案进入2周的 `voting_period` 。
任何**抵押了通证**的持有人都可以参与对这个提案的投票。
投票的选项有`Yes`, `No`, `NoWithVeto` 和 `Abstain`。
投票的权重取决于投票人所抵押的通证数量。如果通证持有人不投票，那么委托人将继承其委托的验证人的投票选项。
当然，委托人也可以自己投出与所委托验证人不同的票。  


当投票期结束后，获得50%（不包括投`Abstain `票）以上 `Yes` 投票权重且少于33.33% 的
`NoWithVeto`（不包括投`Abstain `票）提案将被接受。

#### 实践练习

::: 提示
**在您能够抵押通证或者提取奖励以前，您需要了解[通证抵押](#抵押GARD通证--提取奖励)**
:::

::: 警告

**注意：执行以下命令需要一台联网的计算机。用一个硬件钱包设备执行这些命令会更安全。
关于离线交易过程请看[这里](#从一台离线电脑上签署交易).**
::: 

```bash
// 提交一个提案
// <type>=text/parameter_change/software_upgrade
// ex value for flag: <gasPrice>=100gard

hashgardcli tx gov submit-proposal --title "Test Proposal" --description "My awesome proposal" --type <type> --deposit=10000000gard --gas auto --gas-prices <gasPrice> --from <delegatorKeyName>

// 增加对提案的抵押
// Retrieve proposalID from $hashgardcli query gov proposals --status deposit_period
// 通过 $hashgardcli query gov proposals --status deposit_period 命令获得 `proposalID` 
// 参数设定样例: <deposit>=1000000gard

hashgardcli tx gov deposit <proposalID> <deposit> --gas auto --gas-prices <gasPrice> --from <delegatorKeyName>

// 对一个提案投票
// Retrieve proposalID from $hashgardcli query gov proposals --status voting_period 
通过 $hashgardcli query gov proposals --status deposit_period 命令获得 `proposalID` 
// <option>=yes/no/no_with_veto/abstain

hashgardcli tx gov vote <proposalID> <option> --gas auto --gas-prices <gasPrice> --from <delegatorKeyName>
```

### 从一台离线电脑上签署交易


如果你没有数字钱包设备，而且希望和一台存有私钥的没有联网的电脑进行交互，你可以使用如下过程。
首先，在**联网的电脑上**生成一个没有签名的交易，然后通过下列命令操作（下面以抵押交易为例）：

```bash
// 抵押Atom通证
// 参数设定样例: <amountToBound>=10000000000gard, <bech32AddressOfValidator>=gardvaloper18thamkhnj9wz8pa4nhnp9rldprgant57pk2m8s, <gasPrice>=1000gard

hashgardcli tx staking delegate <validatorAddress> <amountToBond> --from <delegatorKeyName> --gas auto --gas-prices <gasPrice> --generate-only > unsignedTX.json
```

然后，复制 `unsignedTx.json` 并且帮它转移到没有联网的电脑上（比如通过U盘）。
如果没有在没联网的电脑上建立账户，可先[在没有联网的电脑上建立账户](#使用电脑设备进行操作)。
为了进一步保障安全，你可以在签署交易前用以下命令对参数进行检查。 

```bash
cat unsignedTx.json
```

现在，通过以下命令对交易签名：

```bash
hashgardcli tx sign unsignedTx.json --from-addr <delegatorAddr>> signedTx.json
```

复制 `signedTx.json` 并转移回联网的那台电脑。最后，用以下命令向网络广播交易。 

```bash
hashgardcli tx broadcast signedTx.json
```
