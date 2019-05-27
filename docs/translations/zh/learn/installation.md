# Hashgard 安装文档

Hashgard 公链采用 Go 语言编写，它可以在任何能够编译并运行 Go 语言程序的平台上工作。

根据本文档，您将安装两个可执行程序：

1. hashgard，这是 Hashgard 节点的主程序，它将作为系统服务运行并接入 Hashgard 网络。
2. hashgardcli，这是 Hashgard 命令行客户端，用来执行大部分命令，如创建钱包、发送交易等。

## 配置您的服务器

Hashgard 公链基于 Cosmos-SDK 开发，Cosmos SDK 是使用 Go 语言开发的区块链应用程序的框架。

建议在 Linux 服务器中运行验证人节点，如果您在本地电脑上运行，那么当您的电脑休眠或关机时，您的验证人节点将会进入离线 jailed 状态。

**推荐配置:**

- CPU：2Core
- 内存：4GB
- 磁盘：60GB SSD
- 操作系统：Ubuntu 16.04 LTS
- 安全配置：允许来自 TCP 26656-26657 端口的所有传入连接

## 方法 1：源码编译安装

如果您对 Linux 命令不熟悉，我们推荐您使用 `方法二` 安装。

### 1.1 安装 Golang

需要保证 Go 的版本在 1.12 以上，下载 [Go 1.12.3+](https://golang.org/dl)

Golang 安装文档：

1. [https://golang.org/doc/install](https://golang.org/doc/install)
2. [https://github.com/golang/go/wiki/Ubuntu](https://github.com/golang/go/wiki/Ubuntu)

此外，你需要指定运行 Go 所需的 `$GOPATH`、`$GOBIN` 和 `$PATH` 变量, 例如:

```shell
mkdir -p $HOME/go/bin
echo "export GOPATH=$HOME/go" >> ~/.bash_profile
echo "export GOBIN=\$GOPATH/bin" >> ~/.bash_profile
echo "export PATH=\$PATH:\$GOBIN" >> ~/.bash_profile
source ~/.bash_profile
```

### 1.2 克隆 Hashgard 源码

确保您的服务器上安装了 `git`：

```shell
apt-get install git -y
```

将 Hashgard 项目放在指定目录：

```shell
mkdir -p $GOPATH/src/github.com/hashgard
cd $GOPATH/src/github.com/hashgard
git clone https://github.com/hashgard/hashgard
```

### 1.3 编译安装 Hashgard

切换至 master 分支，进行安装：

```shell
cd hashgard && git checkout master
make get_tools && make install
```

### 1.4 确认是否安装成功

使用 help 命令检查是否安装成功：

```shell
hashgard help
hashgardcli help
```

## 方法 2: 直接下载可执行文件

从 github [下载](https://github.com/hashgard/hashgard/releases) 您的操作系统所对应的版本，
并将可执行程序 hashgard、hashgardcli 解压到对应的目录下：

- Linux / MacOS: /usr/local/bin
- Windows: C:\windows\system32\

当完成解压之后，可在 Terminal / CMD 中检查是否安装成功：

```shell
hashgard help
hashgardcli help
```

如果出现

```shell
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

则表示安装成功
