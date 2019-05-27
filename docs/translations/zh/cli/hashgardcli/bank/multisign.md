# hashgardcli bank multisign

## 描述

签名多个签名生成的离线传输文件。该文件由 --generate-only 标志及其他签名者生成的多个签名文件。

## 用法

```
hashgardcli bank multisign [file] [name] [[signature]...] [flags]
```

## Flags

| 名称       | 类型    | 必需 | 默认值                | 描述                                                         |
| ---------------- | ------- | -------- | --------------------- | ------------------------------------------------------------ |
| --offline | bool | 否 | false | 链下模式，不查询全节点 |
| --output-document | string |  |  | 该文档将写入给定文件而不是STDOUT |
| --signature-only | bool | 否 | | 仅打印生成的签名，然后退出 |


 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

#### 创建多签账户

1. 先通过hashgardcli keys add创建三个账号：

```
- hashgardcli keys add a1

- hashgardcli keys add a2

- hashgardcli keys add a3
```

2. 创建一个多签账号：

```
 # 用账户a1&a2&a3创建一个名为a123的多签账户，
 # 只需要其中a1&a2&a3其中任意两个账户就能使用a123这个多签账户
  hashgardcli keys add a123 --multisig=a1,a2,a3 --multisig-threshold=2
```
> 注：该命令表示多签账号名为：a123，只需a1，a2，a3中其中两个任意签名即可。

3. 查询创建的多签账号：

```
hashgardcli keys show a123 -o=text
```
4. 执行完命令后，获得账户的详细信息如下：

```
NAME:   Type:   ADDRESS:                                                PUBKEY:
a123  offline gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n     gardpub1ytql0csgqgfzd666axrjzq7lfft2evw9r7j0u3t7yj4qjy5rczhncv8ysykrp35cpjpklsj5rcfzd666axrjzquew3ad0vgywr7gmgszly9wnw2mwxc3k7dttlmm780g5y9djw8vcgfzd666axrjzq63kk98gyurzz2rewxxhd4dxvvdfsnsdtegajrcez3exg3yu9q0a5kpkkj3
```
#### 给多签账号转账：

```
# 使用hashagrd这个账户转10gard至
# a123（gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n）这个账户
hashgardcli bank send gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n 10gard --from=hashgard
```

#### 使用多签账号给其他账号转账：

1. 发出多签账户a123的转账请求，获得文件

```
# 使用a123这个多签账户给地址为
# gard19thul47y2afwr67l4hlv9hu5593uw0rqhashgjdm7jj 的钱包转账10元
hashgardcli bank send gard19thul47y2afwr67l4hlv9hu5593uw0rqhashgjdm7jj 10gard --from a123 --generate-only >unsignedTx.json
```
>注：该步骤只生成签名文件：unsignedTx.json，请检查系统输出的文件编码是否为 utf-8，如果不是表使用其他文本编辑器将其修改为 utf-8 格式。

2. 使用账户a1签名：

```
hashgardcli bank sign unsignedTx.json --multisig=gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n  --from=a1 --output-document=a1sign.json
```
3. 使用账户a2签名：

```
hashgardcli bank sign unsignedTx.json --multisig=gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n  --from=a2 --output-document=a2sign.json
```
4. 使用多签账号进行多签：

```
hashgardcli bank multisign  unsignedTx.json a123 a1sign.json a2sign.json --output-document=sign.json
```
5. 广播交易：

```
hashgardcli bank broadcast sign.json -o=json --indent
```
6. 执行完命令后，返回详细信息如下：

```
{
 "height": "63108",
 "txhash": "6A66C370834F097CA36F60FE9B4E8ABEEEF3549D089071FDB5EE33277B615035",
 "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
 "gas_wanted": "200000",
 "gas_used": "31450",
 "tags": [
  {
   "key": "action",
   "value": "send"
  },
  {
   "key": "sender",
   "value": "gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n"
  },
  {
   "key": "recipient",
   "value": "gard19thul47y2afwr67l4hlv9hu5593uw0rqjdm7jj"
  }
 ]
}
```
7. 交易查询:

```
hashgardcli tendermint tx 6A66C370834F097CA36F60FE9B4E8ABEEEF3549D089071FDB5EE33277B615035 -o=json --indent
```
8. 执行完命令后，返回详细信息如下：

```
{
 "height": "63108",
 "txhash": "6A66C370834F097CA36F60FE9B4E8ABEEEF3549D089071FDB5EE33277B615035",
 "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
 "gas_wanted": "200000",
 "gas_used": "31450",
 "tags": [
  {
   "key": "action",
   "value": "send"
  },
  {
   "key": "sender",
   "value": "gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n"
  },
  {
   "key": "recipient",
   "value": "gard19thul47y2afwr67l4hlv9hu5593uw0rqjdm7jj"
  }
 ],
 "tx": {
  "type": "auth/StdTx",
  "value": {
   "msg": [
    {
     "type": "cosmos-sdk/Send",
     "value": {
      "from_address": "gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n",
      "to_address": "gard19thul47y2afwr67l4hlv9hu5593uw0rqjdm7jj",
      "amount": [
       {
        "denom": "gard",
        "amount": "10"
       }
      ]
     }
    }
   ],
   "fee": {
    "amount": null,
    "gas": "200000"
   },
   "signatures": [
    {
     "pub_key": {//多签交易
      "type": "tendermint/PubKeyMultisigThreshold",
      "value": {
       "threshold": "2", //两个签名
       "pubkeys": [
        {
         "type": "tendermint/PubKeySecp256k1",
         "value": "A99KVqyxxR+k/kV+JKoJEoPArzww5IEsMMaYDINvwlQe"
        },
        {
         "type": "tendermint/PubKeySecp256k1",
         "value": "A5l0etexBHD8jaIC+QrpuVtxsRt5q1/3vx3ooQrZOOzC"
        },
        {
         "type": "tendermint/PubKeySecp256k1",
         "value": "A1G1inQTgxCUPLjGu2rTMY1MJwavKOyHjIo5MiJOFA/t"
        }
       ]
      }
     },
     "signature": "CgUIAxIBYBJAElZbW6piLDmd+8mG1VLPVYuQK9r/5fitsXvDONtiarVPFSzqf8DkbsyPBOCQOmfuMkhFt+S1TqyFyUZuaE242hJA2j2QTmtW8eEtqOPAkyed0j/97q9phg34KV95gvfp0wd7V0umKoyj/FX/WTvD7iYNWS2ssbwjpztItggOcCTeCw=="
    }
   ],
   "memo": ""
  }
 }
}
```
