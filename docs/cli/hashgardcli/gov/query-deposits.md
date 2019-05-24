# hashgardcli gov query-deposits

## Description

Query details for all deposits on a proposal. You can find the proposal-id by running hashgardcli gov query-proposals

## Usage

```
hashgardcli gov query-deposits [proposal-id] [flags]
```
## Flags
**Global flags, query command flags** [hashgardcli](../README.md)

## Example

###  Query deposits

```shell
hashgardcli gov query-deposits 1 --trust-node -o=json --indent

```

You could query all the deposited tokens on a specific proposal, includes deposit details for each depositor.

```txt
[
  {
    "depositor": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
    "proposal_id": "1",
    "amount": [
      {
        "denom": "gard",
        "amount": "50"
      }
    ]
  }
]

```
