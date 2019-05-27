# hashgardcli gov submit-proposal

## Description

Submit a proposal along with an initial deposit. Proposal Type：Text/ParameterChange/SoftwareUpgrade。

## Usage

```shell
hashgardcli gov submit-proposal [flags]
```
## Flags

| name       | Type               | Required      | Required                   | Description      |
| ---------------- | -------------------------- | ------------ | -------------- | --------------- |
| --deposit        | string | false| "" | deposit of proposal                                                                                                     |
| --description    | string | true| "" | description of proposal                                                                                   |
| --proposal | string | false| "" | proposal file path (if this path is given, other proposal flags are ignored)                 |
| --title          | string | true| "" | title of proposal                                                                                                         |
| --Type           | string | true| "" | proposalType of proposal, Types: text/parameter_change/software_upgrade    |

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Submit a 'text' Type proposal

```shell
hashgardcli gov submit-proposal \
    --title="notice proposal" \
    --Type="Text" \
    --description="a new text proposal" \
    --from=foo
```

输入正确的密码之后，你就完成提交了一个提案，需要注意的是要记下你的提案 ID，这是可以检索你的提案的唯一要素。

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
### Submit a 'Text' Type proposal
```shell
hashgardcli gov submit-proposal \
    --proposal="path/to/proposal.json" \
    --from=foo
```
提案文件内容如下：
```json
{
  "title": "Test Proposal",
  "description": "My awesome proposal",
  "type": "Text",
  "deposit": "10gard"
}
```

输入正确的密码之后，你就完成提交了一个提案，需要注意的是要记下你的提案 ID，这是可以检索你的提案的唯一要素。
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
### Submit a 'SoftwareUpgrade' Type proposal

```shell
hashgardcli gov submit-proposal \
    --title="hashgard" \
    --Type="SoftwareUpgrade" \
    --description="a new software upgrade proposal" \
    --from=hashgard
```

在这种场景下，提案的 --title、--Type 和 --description 参数必不可少，另外你也应该保留好提案 ID，这是检索所提交提案的唯一方法。


How to query proposal

[proposal](proposal.md)

[proposals](proposals.md)
