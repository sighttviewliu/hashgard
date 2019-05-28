# hashgardcli box describe

## 描述
Owner 可以对自己发行的盒子进行补充描述，描述文件使用不超过 1024 字节的 json 格式。可以自定义各种属性，也可以使用官方推荐的模板。
## 用法
```shell
 hashgardcli box describe [box-id] [description-file] --from
```
## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子
### 给盒子设置描述
```shell
hashgardcli box describe boxac3jlxpt2pt ./description.json  --from
```
#### 模板
```shell
{
    "org":"Hashgard",
    "website":"https://www.hashgard.com",
    "logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png",
    "intro":"盒子的简单描述"
}
```
输入正确的密码之后，你的该代币的描述就设置成功了。
```txt
{
 Height: 3536
  TxHash: 026E871E6D7356ECA0A3DAF5A4B1EC563951256B502DC92959424CBF484099BE
  Data: 0F0E626F786163336A6C787074327074
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 41724
  Tags:
    - action = box_description
    - category = box
    - box-id = boxac3jlxpt2pt
    - box-type = future
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
}
```
### 查询盒子信息
```shell
hashgardcli box query-box boxac3jlxpt2pt
```
最新的描述信息就生效了
```shell
{
BoxInfo:
  BoxId:			boxac3jlxpt2pt
  BoxStatus:			depositing
  Owner:			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Name:				PayBox
  BoxType:			future
  TotalAmount:
  Token:			1800000000000000000000agard
  Decimals:			1
  CreatedTime:			1558182333
  Description:			{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"新一代金融公有链"}
  TradeDisabled:		true
FutureInfo:
  MiniMultiple:			1
  Deposit:			[]
  TimeLine:			[]
  Distributed:			[1657912000 1657912001 1657912002]
  Receivers:			[[gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj 100000000000000000000 200000000000000000000 300000000000000000000] [gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0 100000000000000000000 200000000000000000000 300000000000000000000] [gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz 100000000000000000000 200000000000000000000 300000000000000000000]]
}
```
