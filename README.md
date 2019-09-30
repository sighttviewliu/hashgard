# Hashgard

[![version](https://img.shields.io/github/tag/hashgard/hashgard.svg)](https://github.com/hashgard/hashgard/releases/latest)
[![Go](https://img.shields.io/badge/golang-%3E%3D1.13-green.svg?style=flat-square")](https://golang.org)
[![license](https://img.shields.io/github/license/hashgard/hashgard.svg)](https://github.com/hashgard/hashgard/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/hashgard/hashgard)](https://goreportcard.com/report/github.com/hashgard/hashgard)
[![CircleCI](https://circleci.com/gh/hashgard/hashgard/tree/master.svg?style=shield)](https://circleci.com/gh/hashgard/hashgard/tree/master)

Hashgard is new generation digital finance public chain, based on [cosmos-sdk](https://github.com/cosmos/cosmos-sdk) development.

## Required

[Go 1.13+](https://golang.org/dl/)

## Install

Please make sure have already installed `Go` correctly, and set environment variables : `$GOPATH`, `$GOBIN`, `$PATH`.

Put the Hashgard project in the specific path，switch to `master` branch，download related dependencies, then make install:

```
mkdir -p $HOME/go/bin
echo "export GOPATH=$HOME/go" >> ~/.bash_profile
echo "export GOBIN=\$GOPATH/bin" >> ~/.bash_profile
echo "export PATH=\$PATH:\$GOBIN" >> ~/.bash_profile
source ~/.bash_profile
```

Check if the installation is successful:

```
$hashgard --help
$hashgardcli --help
```

## Explorer

[hashgard explorer](https://github.com/hashgard/gardplorer)

## Testnets

[hashgard testnets](https://github.com/hashgard/testnets)
