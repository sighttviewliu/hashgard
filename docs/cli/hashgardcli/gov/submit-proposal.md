# hashgardcli gov submit-proposal

## Description

Submit a proposal along with an initial deposit. Proposal type：Text/ParameterChange/SoftwareUpgrade。

## Usage

```shell
hashgardcli gov submit-proposal [flags]
```
## Flags

| Name       | Type               | Required      | Required                   | Description      |
| ---------------- | -------------------------- | ------------ | -------------- | --------------- |
| --deposit        | string | false|| deposit of proposal                                                                                                     |
| --description    | string | true|| description of proposal                                                                                   |
| --proposal | string | false|| proposal file path (if this path is given, other proposal flags are ignored)                 |
| --title          | string | true|| title of proposal                                                                                                         |
| --type           | string | true|| proposaltype of proposal, types: text/parameter_change/software_upgrade    |

**Global flags, query command flags** [hashgardcli](../README.md)

### ParameterChange

| Name      | Type   | Required | Description        |
| ---------- | ------------- |---------------------- | ------- |
| distribution/community_tax       | float  | 0.02 | Community tax rate  |
| int/foundation_address  | address       | gard1j2znq44kdk2t8kxznlppl0x0j940y62yadeua8 | Foundation address   |
| mint/inflation                   | float       | 0.08                                        | annual inflation               |
| mint/inflation_base              | int         | 100000000000000000000000000000              | Inflation base               |
| gov/min_deposit                  | string        | 100000000000000000000agard                  | Proposed minimum deposit         |
| slashing/signed_blocks_window    | int         | 100                                         | Signed Blocks Window    |
| slashing/min_signed_per_window   | float        | 0.5                                         | Min Signed Per Window     |
| slashing/slash_fraction_downtime | float        | 0.02                                        | Slash Fraction Downtime   |
| slashing/downtime_jail_duration  | time.Duration | 12h or 720m or 3600s (Supported by： h,m,s)       | downtime jail duration|
| staking/max_validators           | int       | 21                                          | Max validators         |
| staking/unbonding_time           | time.Duration | 12h or 720m or 3600s (Supported by： h,m,s) | unbonding time|
| issue/issue_fee                  | string        | 20000000000000000000000agard                | Issuance fee           |
| issue/mint_fee                   | string        | 10000000000000000000000agard                |  mint free             |
| issue/burn_fee                   | string        | 10000000000000000000000agard                | Owner burns his own token fee|
| issue/burn_from_fee              | string        | 10000000000000000000000agard                | Owner burns other tokens fee |
| issue/transfer_owner_fee         | string        | 20000000000000000000000agard                | Token owner transfer fee  |
| issue/describe_fee               | string        | 4000000000000000000000agard                 |  Modify descrption fee        |
| issue/freeze_fee                 | string        | 20000000000000000000000agard                | freeze address fee   |
| issue/unfreeze_fee               | string        | 20000000000000000000000agard                | unfreeze address fee       |
| box/lock_create_fee              | string        | 1000000000000000000000agard                 | Create a lock box fee        |
| box/deposit_box_create_fee       | string        | 10000000000000000000000agard                | Create a deposit box fee         |
| box/future_box_create_fee        | string        | 10000000000000000000000agard                | Create a future box fee     |
| box/disable_feature_fee          | string        | 10000000000000000000000agard                | Use Disabled Function Fee        |
| box/describe_fee                 | string        | 10000000000000000000000agard                | Modifying Description Fee        |

## Example

### Submit a 'text' type proposal

```shell
hashgardcli gov submit-proposal \
    --title="notice proposal" \
    --type="Text" \
    --description="a new text proposal" \
    --from=foo
```

The result is as follows：

```json
{
 "height": "85719",
 "txhash": "8D65804B7259957971AA69515A71AFC1F423080C9484F35ACC6ECD3FBC8EDDDD",
 "data": "AQM=",
 "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
 "gas_wanted": "200000",
 "gas_used": "44583",
 "tags": [
  {
   "key": "action",
   "value": "submit_proposal"
  },
  {
   "key": "proposer",
   "value": "gard10tfnpxvxjh6tm6gxq978ssg4qlk7x6j9aeypzn"
  },
  {
   "key": "proposal-id",
   "value": "3"
  }
 ]
}
```
### Submit a 'Text' type proposal
```shell
hashgardcli gov submit-proposal \
    --proposal="path/to/proposal.json" \
    --from=foo
```
File template
```json
{
  "title": "Test Proposal",
  "description": "My awesome proposal",
  "type": "Text",
  "deposit": "10gard"
}
```
The result is as follows：
```json
{
 "height": "85903",
 "txhash": "9680C11E6631D4EA4B6CE06775D7AC1DAFDA5BD64A98F68E940990CF3E6142D0",
 "data": "AQQ=",
 "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
 "gas_wanted": "200000",
 "gas_used": "55848",
 "tags": [
  {
   "key": "action",
   "value": "submit_proposal"
  },
  {
   "key": "proposer",
   "value": "gard10tfnpxvxjh6tm6gxq978ssg4qlk7x6j9aeypzn"
  },
  {
   "key": "proposal-id",
   "value": "4"
  },
  {
   "key": "voting-period-start",
   "value": "4"
  }
 ]
}
```

### Submit a 'ParameterChange' type proposal
```shell
hashgardcli gov submit-proposal --title="Test-Proposal" --description="My awesome proposal" --type="ParameterChange" --deposit="10gard" --param="box/lock_create_fee=10gard,mint/inflation=1" --from
```
```text
Height: 14596
TxHash: D2CAE79F0278CE39A0B17789E0875F997765F3FEE14A58420E74E31C2BF1DC52
Data: 0103
Raw Log: [{"msg_index":"0","success":true,"log":""}]
Logs: [{"msg_index":0,"success":true,"log":""}]
GasWanted: 200000
GasUsed: 66960
Tags:
  - action = submit_proposal
  - proposal-id = 3
  - category = governance
  - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  - proposal-type = ParameterChange
  - voting-period-start = 3

```

### Submit a 'SoftwareUpgrade' type proposal

```shell
hashgardcli gov submit-proposal \
    --title="hashgard" \
    --type="SoftwareUpgrade" \
    --description="a new software upgrade proposal" \
    --from=hashgard
```



How to query proposal

[proposal](proposal.md)

[proposals](proposals.md)
