# hashgardcli box create-future

## 描述
创建一个远期支付盒子，可以设定多个时间点对不同账户的账户进行远期支付。

## 用法
```shell
hashgardcli box create-future [name] [total-amount][mini-multiple] [distribute-file]  --from
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

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子
### 创建远期支付盒子
```shell
hashgardcli box create-future pay 1800gard  2 ./future.json --from
```
输入正确的密码后，远期支付盒子创建完成。
```txt
  {
 Height: 263
  TxHash: A34024F7C36A345A7C42519890F59D93B05D2FFE4EE33C0994E7D1981A3A1EA5
  Data: 0F0E626F786163336A6C787074327073
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 43797
  Tags:
    - action = box_create_future
    - category = box
    - box-id = boxac3jlxpt2ps
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7

    }
```

为盒子存入需要支付的存款

```shell
hashgardcli box deposit-to boxac3jlxpt2ps 1800  --from
```

返回信息

```txt
 {
  Height: 275
  TxHash: E96FBC4F9C2B3EB3B0C04B091DAAEF45E72E19C24E079879432460B077E137DF
  Data: 0F0E626F786163336A6C787074327073
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 140217
  Tags:
    - action = box_deposit
    - category = box
    - box-id = boxac3jlxpt2ps
    - box-type = future
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
    - operation = deposit-to
}
```

查询盒子信息

```shell
hashgardcli box query-box boxac3jlxpt2ps
```

返回盒子信息

```txt
BoxInfo:
  BoxId:			boxac3jlxpt2ps
  BoxStatus:			actived
  Owner:			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Name:				pay
  BoxType:			future
  TotalAmount:
  Token:			1800000000000000000000agard
  Decimals:			1
  CreatedTime:			1558090817
  Description:
  TradeDisabled:		true
FutureInfo:
  MiniMultiple:			1
  Deposit:			[
  Address:			gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Amount:			1800000000000000000000]
  TimeLine:			[]
  Distributed:			[1657912000 1657912001 1657912002]
  Receivers:			[[gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj 100000000000000000000 200000000000000000000 300000000000000000000] [gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0 100000000000000000000 200000000000000000000 300000000000000000000] [gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz 100000000000000000000 200000000000000000000 300000000000000000000]]

```



### 相关命令

| 名称                        | 描述                   |
| --------------------------- | ---------------------- |
| [deposit-to](deposit-to.md) | 对存款盒子进行分红存入 |
| [query-box](query-box.md)   | 对指定盒子进行信息查询 |
| [list-box](list-box.md)    | 罗列指定类型盒子列表   |
