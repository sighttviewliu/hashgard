# hashgardcli stake

## Description

Stake and validation subcommands

## Usage

```shell
hashgardcli stake [subcommand]
```

## Subcommands

| Name      | description            |
| ----------- | ------- |
| [validator](validator.md)                                   | Query a validator                         |
| [validators](validators.md)                                 | Query for all validators                   |
| [delegation](delegation.md)                                 | Query a delegation based on address and validator address |
| [delegations](delegations.md)                               | Query all delegations made by one delegator  |
| [delegations-to](delegations-to.md)                         | Query all delegations made to one validator|
| [unbonding-delegation](unbonding-delegation.md)             | Query an unbonding-delegation record based on delegator and validator address|
| [unbonding-delegations](unbonding-delegations.md)           | Query all unbonding-delegations records for one delegator|
| [unbonding-delegations-from](unbonding-delegations-from.md) | Query all unbonding delegatations from a validator |
| [redelegation](redelegation.md)                             | Query a redelegation record based on delegator and a source and destination validator address|
| [redelegations](redelegations.md)                           | Query all redelegations records for one delegator|
| [redelegations-from](redelegations-from.md)                 | Query all unbonding delegatations from a validator |
| [pool](pool.md)                                             | Query the current staking pool values        |
| [params](params.md)                                         | Query the current staking parameters information   |
| [create-validator](create-validator.md)                     | create new validator initialized with a self-delegation to it |
| [edit-validator](edit-validator.md)                         | edit an existing validator account      |
| [delegate](delegate.md)                                     | delegate liquid tokens to a validator             |
| [unbond](unbond.md)                                         | unbond shares from a validator         |
| [redelegate](redelegate.md)                                 | redelegate illiquid tokens from one validator to another |
