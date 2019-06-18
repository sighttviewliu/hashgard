# hashgardcli distribution

## Description

This document describes how to use the the command line interfaces of distribution module.


## Usage

```shell
hashgardcli distribution [subcommand]
```

Print all supported subcommands and flags:

```shell
hashgardcli distribution --help
```

## Available Subcommands

| Name                          | Description                                                 |
| --------------------------------| --------------------------------------------------------------|
| [params](params.md)  |Query distribution params |
| [validator-outstanding-rewards](validator-outstanding-rewards.md)  |Query distribution outstanding (un-withdrawn) rewards for a validator and all their delegations|
| [commission](commission.md)  |Query distribution validator commission|
| [slashes](slashes.md)  |Query distribution validator slashes|
| [rewards](rewards.md)  |Query all distribution delegator rewards or rewards from a particular validator|
| [set-withdraw-addr](set-withdraw-address.md)  |change the default withdraw address for rewards associated with an address|
| [withdraw-rewards](withdraw-rewards.md) |witdraw rewards from a given delegation address, and optionally withdraw validator commission if the delegation address given is a validator operator|
| [withdraw-all-rewards](withdraw-rewards.md) | withdraw all delegations rewards for a delegator|
| [community-pool](community-pool.md)  | Query community-pool information|
