# hashgardcli gov submit-proposal

## 描述

提交区块链治理提案以及发起提案所涉及的初始保证金，其中提案的类型包括 Text/ParameterChange/SoftwareUpgrade 这三种类型。

## 用法

```shell
hashgardcli gov submit-proposal [flags]
```
## Flags

| 名称        | 类型                | 必需                 | 默认值                      | 描述               |
| ---------------- | -------------------------- | ----------------- | -------------- | ------------- |
| --deposit        | string | 否 | "" | 提案的保证金                                                                 |
| --description    | string | 是 | "" | 提案的描述                             |
| --proposal | string | 否 | "" | 提案文件路径（如果设置了此路径则忽略其他提案参数）         |
| --title          | string | 是 | "" | 提案标题                                                           |
| --Type           | string | 是 | "" | 提案类型,例如: Text/ParameterChange/SoftwareUpgrade               |


 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

### 提交一个'Text'类型的提案

```shell
hashgardcli gov submit-proposal --title="notice proposal" --Type="Text" --description="a new text proposal" --from=hashgard --chain-id=hashgard -o json --indent
```

输入正确的密码之后，你就完成提交了一个提案，需要注意的是要记下你的提案 ID，这是可以检索你的提案的唯一要素。

```txt
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
### 以提案文件的方式提交一个 'Text' 类型的提案
```shell
hashgardcli gov submit-proposal --proposal="path/to/proposal.json" --from=hashgard --chain-id=hashgard ---o json --indent
```
提案文件内容如下：
```shell
{
  "title": "Test Proposal",
  "description": "My awesome proposal",
  "type": "Text",
  "deposit": "10gard"
}
```

输入正确的密码之后，你就完成提交了一个提案，需要注意的是要记下你的提案 ID，这是可以检索你的提案的唯一要素。
```txt
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
### 提交一个 'SoftwareUpgrade' 类型的提案

```shell
hashgardcli gov submit-proposal --title="hashgard" --Type="SoftwareUpgrade" --description="a new software upgrade proposal" --from=hashgard --chain-id=hashgard
```

在这种场景下，提案的 --title、--Type 和 --description 参数必不可少，另外你也应该保留好提案 ID，这是检索所提交提案的唯一方法。


如何查询提案详情？

请点击下述链接：

[proposal](proposal.md)

[proposals](proposal.md)
