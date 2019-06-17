# hashgardcli issue unfreeze

## Description
Token owner unFreeze the transfer from a address
## Usage
```shell
 hashgardcli issue unfreeze [unfreeze-Type] [issue-id][address] --from
```
### unfreeze-Type

| Name | Description            |
| ------ | -------------------- |
| in     | Transfer in|
| out    | Transfer out|
| In-out | Transfer in and Transfer out|

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### Unfreeze the transfer  of an address
```shell
hashgardcli issue unfreeze in coin174876e800 gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n --from
```
The result is as follows：
```txt
{
  Height: 2758
  TxHash: C6CE11D458B0F64C164E91CF2FF692A65D1EA9C0B1C2A2B228A7C1699C6423FE
  Data: 0F0E636F696E31373438373665383030
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 16203
  Tags:
    - action = issue_unfreeze
    - category = issue
    - issue-id = coin174876e800
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
    - freeze-Type = in
}
```
