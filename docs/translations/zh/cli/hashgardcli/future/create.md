# hashgardcli future create

## 描述
创建一个远期支付盒子，可以设定多个时间点对不同账户的账户进行远期支付。

## 用法
```shell
hashgardcli future create [name] [total-amount][distribute-file]  --from
```



### 命令解释

| 名称          | 类型   | 必需 | 默认值 | 描述                   |
| ------------- | ------ | -------- | ------ | ---------------------- |
| name          | string | 是       |        | 盒子的名称         |
| total-amount  | string | 是       |        | 支付的种类和数量       |
| Mini-multiple | int    | 是       | 1      | 待收款凭证交易最小单位 |



#### distribute-file

```shell
{
   "time":[1657912000,1657912001,1657912002],//不同批次的支付时间
   "receivers":[
     ["gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj","100","200","300"],//支付的账户和数量
     ["gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0","100","200","300"],
     ["gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz","100","200","300"]
    ]
}
```


## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 创建远期支付盒子
```shell
hashgardcli future create  pay-one 1800gard  .path/future.json --from
```
输入正确的密码后，远期支付盒子创建完成。
```txt
Response:
  Height: 1169
  TxHash: 886D10C0E1B5ACD755D3333E08AB36F91E6D5DDCFD3ED073E34DCCFAE5637D3B
  Data: 0F0E626F786163336A6C787074327073
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 65595
  Tags:
    - action = create
    - category = future
    - id = boxac3jlxpt2ps
    - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
    - fee = 1000000000000000000000agard
```

为盒子存入需要支付的存款

```shell
hashgardcli future inject boxac3jlxpt2ps 1800  --from
```

成功后，返回结果:

```txt
Height: 1212
TxHash: 3AE63B74450CEF0BB712B0788784310D5711A825BF60BC29821BD41EADC00FBF
Data: 0F0E626F786163336A6C787074327073
Raw Log: [{"msg_index":"0","success":true,"log":""}]
Logs: [{"msg_index":0,"success":true,"log":""}]
GasWanted: 200000
GasUsed: 134927
Tags:
 - action = inject
 - category = future
 - id = boxac3jlxpt2ps
 - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
```

查询盒子信息

```shell
hashgardcli future query boxac3jlxpt2ps
```

返回盒子信息

```txt
BoxInfo:
  Id:			boxac3jlxpt2ps
  Status:			actived
  Owner:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Name:				pay-one
  TotalAmount:
  Token:			1800000000000000000000agard
  Decimals:			1
  CreatedTime:			1559550097
  Description:
  TransferDisabled:		true
FutureInfo:
  Injects:			[
  Address:			gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
  Amount:			1800000000000000000000]
  TimeLine:			[1593762909 1599119709 1606982109]
  Receivers:			[[gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj 100000000000000000000 200000000000000000000 300000000000000000000] [gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0 100000000000000000000 200000000000000000000 300000000000000000000] [gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz 100000000000000000000 200000000000000000000 300000000000000000000]]
  TotalWithdrawal:			0
```
