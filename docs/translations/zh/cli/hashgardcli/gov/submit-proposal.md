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
| --deposit        | string | 否 || 提案的保证金                                                                 |
| --description    | string | 是 || 提案的描述                             |
| --proposal | string | 否 || 提案文件路径（如果设置了此路径则忽略其他提案参数）         |
| --title          | string | 是 || 提案标题                                                           |
| --type           | string | 是 || 提案类型,例如: Text/ParameterChange/SoftwareUpgrade               |


**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

### ParameterChange

| 名称，速记        | 类型  | 默认值 | 描述          |
| ---------- | ------ | -------- | ------- | --- |
| distribution/community_tax  |  浮点型   |   0.02  |  community 抽成的税率  |
| int/foundation_address | address  |  gard1j2znq44kdk2t8kxznlppl0x0j940y62yadeua8   |  用于扣款的基金会地址  |
| mint/inflation    | 浮点型    |   0.08  |  年通胀率  |
| mint/inflation_base    |  整型  | 100000000000000000000000000000  |   通胀基数  |
| gov/min_deposit    |   string   |   100000000000000000000agard   | 提案的最低存款  |
| slashing/signed_blocks_window    |  整型  | 100 | 掉线处罚的监测窗口  |
| slashing/min_signed_per_window    | 浮点型   |  0.5 |   窗口的签名率  |
| slashing/slash_fraction_downtime     | 浮点型   |  0.02    | 掉线的罚金率   |
| slashing/downtime_jail_duration|time.Duration| 12h or 720m or 3600s (支持单位 h,m,s)   |  掉线监禁时长  |
| staking/max_validators    |   整型   | 21  |  活跃验证人数量  |
| staking/unbonding_time     |  time.Duration   |  12h or 720m or 3600s (支持单位 h,m,s) | 股权解绑时间|
| issue/issue_fee   |   string  | 20000000000000000000000agard   |  发行的费用 |
| issue/mint_fee   |   string  | 10000000000000000000000agard   |  增发的费用 |
| issue/burn_fee   |  string   |  10000000000000000000000agard  | 销毁自身持有代币的费用  |
| issue/burn_from_fee    |  string  |  10000000000000000000000agard |  owner 消毁持有者的费用  |
| issue/transfer_owner_fee    |  string  | 20000000000000000000000agard  | 转移 owner 权限的费用  |
| issue/describe_fee   |   string  |  4000000000000000000000agard |修改描述的费用 |
| issue/freeze_fee   |   string  | 20000000000000000000000agard  |  冻结账户地址的费用  |
| issue/unfreeze_fee     |   string | 20000000000000000000000agard  |  解冻地址的费用 |
| box/lock_create_fee    |   string | 1000000000000000000000agard  |  创建锁仓的费用  |
| box/deposit_box_create_fee     |   string | 10000000000000000000000agard  | 创建存款的费用   |
| box/future_box_create_fee    |  string  |10000000000000000000000agard   |   创建远期支付的费用 |
| box/disable_feature_fee   |  string  | 10000000000000000000000agard  |  禁用特性的费用 |
| box/describe_fee    |  string  |10000000000000000000000agard| 修改描述的费用   |



## 例子

### 提交一个'Text'类型的提案

```shell
hashgardcli gov submit-proposal --title="notice proposal" --type="Text" --description="a new text proposal" --from=hashgard --chain-id=hashgard -o json --indent
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


### 提交一个 ‘ParameterChange’ 类型的提案
```shell
hashgardcli gov submit-proposal --title="Test-Proposal" --description="My awesome proposal" --type="ParameterChange" --deposit="10gard" --param="box/lock_create_fee=10gard,mint/inflation=1" --from
```
返回结果
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
### 提交一个 'SoftwareUpgrade' 类型的提案


```shell
hashgardcli gov submit-proposal --title="hashgard" --type="SoftwareUpgrade" --description="a new software upgrade proposal" --from
```

在这种场景下，提案的 --title、--type 和 --description 参数必不可少，另外你也应该保留好提案 ID，这是检索所提交提案的唯一方法。


如何查询提案详情？

请点击下述链接：

[proposal](proposal.md)

[proposals](proposal.md)
