# hashgardcli gov query-votes

## Description

Query vote details for a single proposal by its identifier

## Usage

```shell
hashgardcli gov query-votes [proposal-id] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Query votes

```shell
hashgardcli gov query-votes 1 --trust-node -o=json --indent
```

You could query the voting of all the voters by specifying the proposal id.

```txt
[
  {
    "voter": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
    "proposal_id": "1",
    "option": "Yes"
  }
]
```
