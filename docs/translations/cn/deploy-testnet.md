# 部署你自己的测试网

## 单节点，本地的，手动的测试网

本教程可帮助你创建一个在本地运行网络的验证人节点，以进行测试和其他相关的用途。

### 需要
+ [安装 Hashgard](./installation.md)

### 创建genesis文件并启动网络

```bash
# You can run all of these commands from your home directory
cd $HOME

# Initialize the genesis.json file that will help you to bootstrap the network
hashgard init --chain-id=testing testing

# Create a key to hold your validator account
hashgardcli keys add validator

# Add that key into the genesis.app_state.accounts array in the genesis file
# NOTE: this command lets you set the number of coins. Make sure this account has some coins
# with the genesis.app_state.staking.params.bond_denom denom, the default is staking
hashgard add-genesis-account $(hashgardcli keys show validator -a) 1000000000stake,1000000000validatortoken

# Generate the transaction that creates your validator
hashgard gentx --name validator

# Add the generated bonding transaction to the genesis file
hashgard collect-gentxs

# Now its safe to start `hashgard`
hashgard start
```

启动将会把`hashgard`相关的所有数据放在`~/.hashgard`目录。
你可以检查所创建的genesis文件——`~/.hashgard/config/genesis.json`。
同时`hashgardcli`也已经配置完成并且有了一个拥有token的账户(stake和自定义的代币)。

### 密钥&账户

你需要使用指定节点的`hashgardcli`目录作为你的`home`来同`hashgardcli`交互，并执行查询或者创建交易:

```bash
hashgardcli keys list --home ./build/node0/hashgardcli
```

现在账户已经存在了，你可以创建新的账户并向其发送资金！

::: 提示
注意：每个节点的密钥种子放在`./build/nodeN/hashgardcli/key_seed.json`中，
可以通过`hashgardcli keys add --recover`命令来回复。
:::