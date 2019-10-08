# Deploy Your Own Hashgard Testnet

## Single-node, Local, Manual Testnet

This guide helps you create a single validator node that runs a network locally for testing and other development related uses.

### Requirements

- [Install hashgard](./installation.md)

### Create Genesis File and Start the Network

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

This setup puts all the data for `hashgard` in `~/.hashgard`. You can examine the genesis file you created at `~/.hashgard/config/genesis.json`. With this configuration `hashgardcli` is also ready to use and has an account with tokens (both staking and custom).

### Keys & Accounts

To interact with `hashgardcli` and start querying state or creating txs, you use the
`hashgardcli` directory of any given node as your `home`, for example:

```shell
hashgardcli keys list --home ./build/node0/hashgardcli
```

Now that accounts exists, you may create new accounts and send those accounts
funds!

::: tip
**Note**: Each node's seed is located at `./build/nodeN/hashgardcli/key_seed.json` and can be restored to the CLI using the `hashgardcli keys add --restore` command
:::
