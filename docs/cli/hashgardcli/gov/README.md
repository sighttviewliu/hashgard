# hashgardcli gov

## Description

This module provides the basic functions as described below:

* On-chain governance proposals on idea
* On-chain governance proposals on parameter change
* On-chain governance proposals on software upgrade

## Usage

```shell
hashgardcli gov [command]
```

Print all supported subcommands and flags:

```shell
hashgardcli gov --help
```
## Available Commands

|  Commands                           | Description                                                        |
| ------------------------------------- | --------------------------------------------------------------- |
| [proposal](proposal.md)   | Query details of a single proposal           |
| [proposals](proposals.md) | Query proposals with optional filters        |
| [query-vote](query-vote.md)           | Query details of a single vote                     |
| [query-votes](query-votes.md) | Query votes on a proposal                      |
| [query-deposit](query-deposit.md)     |  Query details of a deposit  |
| [deposits](deposits.md)   |Deposit tokens for activing proposal|
| [tally](tally.md)         | Get the tally of a proposal vote |
| [param](param.md)       | Query the parameters (voting|tallying|deposit) of the governance process|
| [submit-proposal](submit-proposal.md) | Submit a proposal along with an initial deposit|
| [deposit](deposit.md)     | Deposit tokens for activing proposal |
| [vote](vote.md)   | Vote for an active proposal, options: yes/no/no_with_veto/abstain|

## Extended description

* Any user can deposit some tokens to submit a proposal.Once the deposit reaches a certain value min_deposit, the proposal enter voting period, otherwise it will remain in the deposit period. Others can deposit the proposals in the deposit period.  Once the sum of the deposit reaches min_deposit, but，if the block-time exceeds max_deposit_period in the deposit period, the proposal will be closed.
* 进入投票期的提案只能由验证人和委托人投票。未投票的代理人的投票将与其验证人的投票相同，并且投票的代理人的投票将保留。到达“voting_period”时，票数将被计算在内。
