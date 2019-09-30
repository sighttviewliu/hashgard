# 加入公共测试网

::: 提示 当前测试网
请查看[testnet repo](https://github.com/hashgard/testnets)获取最新的公共测试网信息，包含了所使用的Cosmos-SDK的正确版本和genesis文件。
:::

::: 警告
你需要先完成[安装`hashgard`](./installation.md)
:::

## 创建一个新节点

> 注意：如果你在之前的测试网中运行过一个全节点，请跳至[升级之前的Testnet](#升级之前的Testnet)。

这些指令适用于从头开始设置一个全节点。

首先，初始化节点并创建必要的配置文件：

```bash
hashgard init <your_custom_moniker>
```

::: 注意
moniker只能包含ASCII字符。使用Unicode字符会使得你的节点不可访问
:::

你可以稍后在`~/.hashgard/config/config.toml`文件中编辑`moniker`:

```toml
# A custom human readable name for this node
moniker = "<your_custom_moniker>"
```

你可以编辑`~/.hashgard/config/config.toml`文件来开启垃圾交易过滤机制以拒绝收到的手续费过低的交易：

```
# This is a TOML config file.
# For more information, see https://github.com/toml-lang/toml

##### main base config options #####

# The minimum gas prices a validator is willing to accept for processing a
# transaction. A transaction's fees must meet the minimum of any denomination
# specified in this config (e.g. 10000000000.0agard).

minimum-gas-prices = ""
```

你的全节点已经初始化成功！

## Genesis & Seeds

### 复制genesis文件

将主网的`genesis.json`文件放置在`hashgard`的配置文件夹中

```bash
curl https://raw.githubusercontent.com/hashgard/testnets/master/sif/sif-7000/config/genesis.json > $HOME/.hashgard/config/genesis.json
```

注意我们使用了[testnet repo](https://github.com/hashgard/testnets)中的`sif-7000`文件夹，该文件夹包含了最新版本测试网的详细信息。

运行命令验证配置的正确性:

```bash
hashgard start
```

### 添加种子节点

你的节点需要知道如何寻找伙伴节点。你需要添加有用的种子节点到`$HOME/.hashgard/config/config.toml`文件中。

[testnet repo](https://github.com/hashgard/testnets) 提供的`config.toml`文件中包含了种子节点的配置。

```bash
curl https://raw.githubusercontent.com/hashgard/testnets/master/sif/sif-7000/config/config.toml > $HOME/.hashgard/config/config.toml
```

::: 警告
在Cosmos Hub主网中，可接受的币种是`agard`,`1agard= 0.000000000000000001gard`
:::

Hashgard 网络中的交易手续费计算公式如下：

```
fees = gas * gasPrices
```

`gas`由交易本身决定。不同的交易需要不同数量的`gas`。一笔交易的`gas`数量在它被执行时计算，但有一种方式可以提前估算，那就是把标识`gas`
的值设置为`auto`。当然，这只是给出一个预估值。如果你想要确保为交易提供足够的gas，你可以使用`--gas-adjustment`标识来调整预估值(默认是`1.0`)。

`gasPrice`是每个单位`gas`的单价。每个验证人节点可以设置`min-gas-price`，只会把那些`gasPrice`高于自己设置的`min-gas-price`的交易打包。

交易的`fees`是`gas`与`gasPrice`的结果。作为一个用户，你必须输入三者中的两者。更高的`gasPrice`/`fees`，将提高你的交易被打包的机会。

## 运行一个全节点

通过这条命令开始运行全节点：

```bash
hashgard start
```

检查一切是否平稳运行中:

```bash
hashgardcli status
```

## 升级之前的Testnet

这些指令用以把运行过以前测试网络的全节点升级至最新的测试网络。

### 重置数据

首先，移除过期的文件并重置数据：

```bash
rm $HOME/.hashgard/config/addrbook.json $HOME/.hashgard/config/genesis.json
hashgard unsafe-reset-all
```

你的节点现在处于原始状态并保留了最初的`priv_validator.json`文件和`config.toml`文件。如果之前你还有其他的哨兵节点或者全节点，你的节点仍然会连接他们，但是会失败，因为他们还没有升级。

::: 警告
确保每个节点有一个独一无二的`priv_validator.json`文件。不要从一个旧节点拷贝`priv_validator.json`到多个新的节点。运行两个有着相同`priv_validator.json`文件的节点会导致双签。
:::

### 升级软件

现在升级软件：

```bash
cd $GOPATH/src/github.com/cosmos/cosmos-sdk
git fetch --all && git checkout master
make update_tools install
```

::: 提示
*注意*：如果在这一步出现问题，请检查是否安装了最新稳定版本的Go。
:::

或者直接从[Github Releases](https://github.com/hashgard/hashgard/releases)下载可执行程序，覆盖旧版本，并重新启动：

```bash
hashgard start
```

你的全节点已经升级成功！

## 升级成为验证人节点

你现在有了一个运行状态的全节点。接下来，你可以升级你的全节点，成为一个 Hashgard 验证人。排名前100的验证人节点可以向网络提交新的区块。

请查看[创建验证人节点](./validators/validator-setup.md)。