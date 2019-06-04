# hashgardcli lock describe

## Description
Owner describes the box。The description file must be in josn format and no more than 1024 bytes.
## Usage
```shell
 hashgardcli lock describe [box-id] [description-file] [flags]
```
## Flags

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### 给盒子设置描述
```shell
hashgardcli lock describe boxaa3jlxpt2ps .path/Desktop/description.json --from
```
#### 模板
```shell
{
    "org":"Hashgard",
    "website":"https://www.hashgard.com",
    "logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png",
    "intro":"Foundation lock"
}
```
输入正确的密码之后，你的该代币的描述就设置成功了。
```txt
{
  Height: 4271
 TxHash: 58535592CE82533E14AEE1FFF1D64115B8BCBBFC563FBF4A40000D691E993CCE
 Data: 0F0E626F786161336A6C787074327073
 Raw Log: [{"msg_index":"0","success":true,"log":""}]
 Logs: [{"msg_index":0,"success":true,"log":""}]
 GasWanted: 200000
 GasUsed: 45327
 Tags:
   - action = describe
   - category = lock
   - id = boxaa3jlxpt2ps
   - sender = gard1gzs53ktf6u3yaplpqmx4wpq7n8s0gw75n73x3t
}
```
## Query
```shell
hashgardcli lock query boxaa3jlxpt2ps
```
The result is as follows：
```shell
{
  BoxInfo:
    Id:			boxaa3jlxpt2ps
    Status:			locked
    Owner:			gard1gzs53ktf6u3yaplpqmx4wpq7n8s0gw75n73x3t
    Name:				lockbox
    TotalAmount:
    Token:			8987000000000000000000agard
    Decimals:			1
    CreatedTime:			1559531133
    Description:			{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"Foundation lock"}
    TransferDisabled:		true
  LockInfo:
    EndTime:			1562123084

}
```
