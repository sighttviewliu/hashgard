# hashgardcli gov proposal

## Description

Query details for a proposal. You can find the proposal-id by running hashgardcli gov query-proposals:

## Usage

```shell
hashgardcli gov proposal [proposal-id] [flags]
```
## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Query specified proposal

```shell
hashgardcli gov proposal 1 --trust-node -o=json --indent
```

The result is as follows：

```txt
{
  "type": "gov/TextProposal",
  "value": {
    "proposal_id": "1",
    "title": "notice proposal",
    "description": "a new text proposal",
    "proposal_Type": "Text",
    "proposal_status": "DepositPeriod",
    "tally_result": {
      "yes": "0.0000000000",
      "abstain": "0.0000000000",
      "no": "0.0000000000",
      "no_with_veto": "0.0000000000"
    },
    "submit_time": "2018-12-20T11:40:43.123286817Z",
    "deposit_end_time": "2018-12-22T11:40:43.123286817Z",
    "total_deposit": null,
    "voting_start_time": "0001-01-01T00:00:00Z",
    "voting_end_time": "0001-01-01T00:00:00Z"
  }
}
```
