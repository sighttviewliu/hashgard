# hashgardcli gov tally

## Description

Query tally of votes on a proposal.

## Usage

```shell
 hashgardcli gov tally [proposal-id] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)


## Example

### Query tally

```shell
hashgardcli gov tally 1 --trust-node
```

You could query the statistics of each voting option.

```txt
{
  "yes": "89.1000000000",
  "abstain": "0.0000000000",
  "no": "0.0000000000",
  "no_with_veto": "0.0000000000"
}
```
