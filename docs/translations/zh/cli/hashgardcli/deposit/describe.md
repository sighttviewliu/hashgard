# hashgardcli describe

## 描述
Owner 可以对自己发行的盒子进行补充描述，描述文件使用不超过 1024 字节的 json 格式。可以自定义各种属性，也可以使用官方推荐的模板。
## 用法
```shell
 hashgardcli deposit describe [box-id] [description-file] [flags]
```
## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 给盒子设置描述
```shell
hashgardcli deposit describe boxab3jlxpt2pt .path/description.json --from one -y
```
#### 模板
```shell
{
    "org":"Hashgard",
    "website":"https://www.hashgard.com",
    "logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png",
    "intro":"This is a description of the project"
}
```
成功后，返回结果:
```txt
{
  Response:
    Height: 3556
    TxHash: 031F0CE15F30D68142BD22B23A4555E428506FBDFA4D46490D182356BCCA97DB
    Data: 0F0E626F786162336A6C787074327074
    Raw Log: [{"msg_index":"0","success":true,"log":""}]
    Logs: [{"msg_index":0,"success":true,"log":""}]
    GasWanted: 200000
    GasUsed: 55284
    Tags:
      - action = describe
      - category = deposit
      - id = boxab3jlxpt2pt
      - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
}
```
### 查询盒子信息
```shell
hashgardcli deposit query boxab3jlxpt2pt
```
成功后，返回结果:
```shell
BoxInfo:
  Id:			boxab3jlxpt2pt
  Status:			interest
  Owner:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Name:				deposit88
  TotalAmount:
  Token:			10000000000000000000000agard
  Decimals:			1
  CreatedTime:			1559616171
  Description:			{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"This is a description of the project"}
  TransferDisabled:		true
DepositInfo:
  StartTime:			1559616600
  EstablishTime:		1559617200
  MaturityTime:			1559703600
  BottomLine:			0
  Interest:
  Token:			9090000000000000000000coin174876e800
  Decimals:			18
  Price:			2000000000000000000
  PerCoupon:			1.818000000000000000
  Share:			5000
  TotalInject:			10000000000000000000000
  WithdrawalShare:			0,
  WithdrawalInterest:			0,
  InterestInject:			[
  Address:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Amount:			9090000000000000000000]
```
