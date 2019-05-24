# Installation

You will install two executables according to this guide:

1. hashgard. This is the main program, it runs as a node of Hashgard blockchain.
2. hashgardcli. This is the client of Hashgard, it's used to execute most commands like creating wallet and sending transactions.

## Prepare Your VPS

It is recommended to run the Hashgard Validator Node on VPS.

If you run the node on your local machine, it will be jailed when your computer is hibernating or shutting down.

**Recommend Configurations**

- CPU: 2 Cores
- RAM: 4GB
- Storage: 60GB SSD
- OS: Ubuntu 16.04 LTS
- Security Config: allow any incoming TCP connections of port 26656-26657

## Method 1: Install from Sourcecode

It is recommended to use `Method 2` to install hashgard if you're not familiar with Linux commands.

### 1.1 Install Golang

Hashgard is based on [consmos-sdk](https://github.com/cosmos/cosmos-sdk) which is built in Golang.

So it requires [Go 1.11.5+](https://golang.org/dl).

Please install Golang according to the links below:

1. [https://golang.org/doc/install](https://golang.org/doc/install)
2. [https://github.com/golang/go/wiki/Ubuntu](https://github.com/golang/go/wiki/Ubuntu)

In addition, you shold configure the `$GOPATH`, `$GOBIN` and `$PATH` ENV for Golang:

```bash
mkdir -p $HOME/go/bin
echo "export GOPATH=$HOME/go" >> ~/.bash_profile
source ~/.bash_profile
echo "export GOBIN=$GOPATH/bin" >> ~/.bash_profile
source ~/.bash_profile
echo "export PATH=$PATH:$GOBIN" >> ~/.bash_profile
source ~/.bash_profile
```

### 1.2 Clone Hashgard Sourcecode

Make sure you have installed `git`：

```
apt-get install git -y
```

clone Hashgard Repo to GOPATH：

```bash
mkdir -p $GOPATH/src/github.com/hashgard
cd $GOPATH/src/github.com/hashgard
git clone https://github.com/hashgard/hashgard
```

### 1.3 Build Hashgard

Switch to branch master, and build：

```bash
cd hashgard && git checkout master
make get_tools && make install
```

### 1.4 Test the Executable Installation

Test with command `help`:

```bash
hashgard help
hashgardcli help
```

## Method 2: Download Executables

Download packages from [Github Releases](https://github.com/hashgard/hashgard/releases) according to your OS,

then extract the excutables (hashgard and hashgardcli) to the specified directory:

- Linux / MacOS: /usr/local/bin
- Windows: C:\windows\system32\

Run the commands below to test the executable installation:

```bash
hashgard help
hashgardcli help
```

If they are installed successfully, you'll see the output:

```
Hashgard Daemon (server)

Usage:
  hashgard [command]

Available Commands:
  start               Run the full node
  init                Initialize genesis config, priv-validator file, p2p-node file, and application configuration files
  collect-gentxs      Collect genesis txs and output a genesis.json file
  testnet             Initialize files for a Hashgard testnet
  gentx               Generate a genesis tx carrying a self delegation
  add-genesis-account Add genesis account to genesis.json
  validate-genesis    validates the genesis file at the default location or at the location passed as an arg
  unsafe-reset-all    Resets the blockchain database, removes address book files, and resets priv_validator.json to the genesis state

  tendermint          Tendermint subcommands

  export              Export state to JSON

  version             Print the app version
  help                Help about any command

······

Use "hashgardcli [command] --help" for more information about a command.
```
