# hashgardcli lock describe

## 描述
Owner 可以对自己发行的盒子进行补充描述，描述文件使用不超过 1024 字节的 json 格式。可以自定义各种属性，也可以使用官方推荐的模板。
## 用法
```shell
 hashgardcli lock describe [box-id] [description-file] [flags]
```
## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
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
### 查询盒子信息
```shell
hashgardcli lock query boxaa3jlxpt2ps
```
成功后，返回结果:
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
