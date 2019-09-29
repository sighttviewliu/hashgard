# 安装 Hashgard

Hashgard 公链采用 Go 语言编写，它可以在任何能够编译并运行 Go 语言程序的平台上工作。

根据本文档，您将安装两个可执行程序：

`hashgard`，这是 Hashgard 节点的主程序，它将作为系统服务运行并接入 Hashgard 网络。
`hashgardcli`，这是 Hashgard 命令行客户端，用来执行大部分命令，如创建钱包、发送交易等。

安装后，你可以作为[全节点](./join-testnet.md)或是[验证人节点](./validators/validator-setup.md)加入到测试网。

### 配置您的服务器

Hashgard 公链基于 Cosmos-SDK 开发，Cosmos SDK 是使用 Go 语言开发的区块链应用程序的框架。

建议在 Linux 服务器中运行验证人节点，如果您在本地电脑上运行，那么当您的电脑休眠或关机时，您的验证人节点将会进入离线 jailed 状态。

#### 推荐配置

```
CPU：2Core
内存：4GB
磁盘：60GB SSD
操作系统：Ubuntu 16.04 LTS
安全配置：允许来自 TCP 26656-26657 端口的所有传入连接
```

## 方法 1: 直接下载可执行文件

从 github [下载](https://github.com/hashgard/hashgard/releases) 您的操作系统所对应的版本， 并将可执行程序 hashgard、hashgardcli 解压到对应的目录下：

```
Linux / MacOS: /usr/local/bin
Windows: C:\windows\system32\
```

当完成解压之后，可在 Terminal / CMD 中检查是否安装成功：

```
hashgard help
hashgardcli help
```

## 方法 2：源码编译安装

### 安装 Go

按照[官方文档](https://golang.org/doc/install)安装`go`。记得设置环境变量`$GOPATH`,`$GOBIN`和`$PATH`:

```bash
mkdir -p $HOME/go/bin
echo "export GOPATH=$HOME/go" >> ~/.bash_profile
echo "export PATH=\$PATH:\$GOPATH/bin" >> ~/.bash_profile
echo "export GO111MODULE=on" >> ~/.bash_profile
source ~/.bash_profile
```

::: tip
Cosmos SDK 需要安装**Go 1.13+**
:::

### 安装二进制执行程序

接下来，安装最新版本的 Hashgard。需要确认您 `git checkout 了正确的[发布版本](https://github.com/hashgard/hashgard/releases)。

```bash
git clone -b <latest-release-tag> https://gitlab.hashgard.com/hashgard/hashgard
cd gaia && make install
```

> _注意_: 如果在这一步中出现问题，请检查你是否安装的是 Go 的最新稳定版本。

等`hashgard`和`hashgardcli`可执行程序安装完之后，请检查:

```bash
$ hashgard version --long
$ hashgardcli version --long
```

`hashgardcli`的返回应该类似于：

```
cosmos-sdk: 0.33.0
git commit: 7b4104aced52aa5b59a96c28b5ebeea7877fc4f0
go.sum hash: d156153bd5e128fec3868eca9a1397a63a864edb5cfa0ac486db1b574b8eecfe
build tags: netgo ledger
go version go1.13+ linux/amd64
```

##### Build Tags

build tags 指定了可执行程序具有的特殊特性。

| Build Tag | Description                           |
| --------- | ------------------------------------- |
| netgo     | Name resolution will use pure Go code |
| ledger    | 支持 Ledger 设备(硬件钱包)            |

### 接下来

然后你可以选择 [加入测试网](./join-testnet.md) 或是 [创建私有测试网](./deploy-testnet.md)。
