# Join the Public Testnet 

::: tip Current Testnet
See the [testnet repo](https://github.com/hashgard/testnets) for
information on the latest testnet, including the correct version
of Hashgard to use and details about the genesis file.
:::

::: warning
**You need to [install hashgard](./installation.md) before you go further**
:::

## Starting a New Node

First, remove the outdated files and reset the data.

```bash
rm $HOME/.hashgard/config/addrbook.json $HOME/.hashgard/config/genesis.json
hashgard unsafe-reset-all
```

Your node is now in a pristine state while keeping the original `priv_validator.json` and `config.toml`. If you had any sentry nodes or full nodes setup before,
your node will still try to connect to them, but may fail if they haven't also
been upgraded.

::: danger Warning
Make sure that every node has a unique `priv_validator.json`. Do not copy the `priv_validator.json` from an old node to multiple new nodes. Running two nodes with the same `priv_validator.json` will cause you to double sign.
:::

### Software Upgrade

Now it is time to upgrade the software:

```bash
cd $GOPATH/src/gitlab.hashgard.com/hashgard/hashgard
git fetch --all && git checkout master
make update_tools install
```

::: tip
*NOTE*: If you have issues at this step, please check that you have the latest stable version of GO installed.
:::

Note we use `master` here since it contains the latest stable release.
See the [testnet repo](https://github.com/hashgard/testnets) for details on which version is needed for which testnet, and the [Hashgard release page](https://gitlab.hashgard.com/hashgard/hashgard/releases) for details on each release.

Your full node has been cleanly upgraded!
