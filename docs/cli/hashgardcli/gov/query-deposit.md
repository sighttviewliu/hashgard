# hashgardcli gov query-deposit

## Description

Query details for a single proposal deposit on a proposal by its identifier

## Usage

```shell
 hashgardcli gov query-deposit [proposal-id] [depositer-addr] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Query deposit

```shell
hashgardcli gov query-deposit 1 gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet --trust-node -o=json --indent
```

You could query the deposited tokens on a specific proposal by `proposal-id` and `depositor`.

```txt
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
```
