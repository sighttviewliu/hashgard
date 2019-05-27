# hashgardcli gov query-vote

## Description

Query details for a single vote on a proposal given its identifier

## Usage

```shell
hashgardcli gov query-vote [proposal-id] [voter-addr] [flags]
```

## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

### Query vote

```shell
hashgardcli gov query-vote 1 gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet --trust-node -o=json --indent
```

You could query the voting by specifying the proposal id and the voter.

```txt
{
  "voter": "gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet",
  "proposal_id": "1",
  "option": "Yes"
}

```
