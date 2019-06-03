# hashgardcli gov proposals

## Description

Query for a all proposals. You can filter the returns with the following flags

## Usage

```shell
hashgardcli gov proposals [flags]
```
## Flags

| Name     | Type         | Required      | Default   | Description  |
| --------------- | -------------------------- | -- | -------- | ------ |
| --depositor     | string | false|| filter by proposals deposited on by depositor                                       |
| --limit         | string | false|| limit to latest[number] proposals. Defaults to all proposals       |
| --status        | string | false|| filter proposals by proposal status               |
| --voter         | string | false|| filter by proposals voted on by voted      |

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Query proposals

```shell
hashgardcli gov proposals --trust-node
```

You could query all the proposals by default.

```txt
ID - (Status) [Type] Title
1 - (DepositPeriod) [Text] Test Proposal
2 - (DepositPeriod) [Text] Test Proposal
3 - (DepositPeriod) [Text] Test Proposal
4 - (VotingPeriod) [Text] Test Proposal
```

Also you can query proposal by filters, such as:

```shell
gov proposals --chain-id=hashgard --depositor=gard4q5rf9sl2dqd2uxrxykafxq3nu3lj2fp9l7pgd
```

Finally, here shows the proposal who's depositor address is  gard4q5rf9sl2dqd2uxrxykafxq3nu3lj2fp9l7pgd
```txt
  2 - new proposal
```
