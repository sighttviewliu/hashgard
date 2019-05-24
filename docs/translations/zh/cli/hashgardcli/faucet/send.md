# hashgardcli faucet send

## 描述

该功能仅在测试网中开放，用于领取测试网代币

## 用法

```
hashgardcli faucet send [address] [flags]
```

## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

```shell
 hashgardcli faucet send gard18vdg4r46qtfkwyghsg67dpn9p7vassw30z0f46 --chain-id=hashgard -o=json --indent
```

执行成功后，你将会从水龙头地址获得50gard,50apple的转账。

```txt
{
 "height": "6846",
 "txhash": "DBE1C8E78F91B3FBA1E000B92D751651DFAF3894C660DA0A7A19A11BA2CE7A56",
 "logs": [
  {
   "msg_index": "0",
   "success": true,
   "log": ""
  }
 ],
 "gas_wanted": "50000",
 "gas_used": "23899",
 "tags": [
  {
   "key": "action",
   "value": "send"
  },
  {
   "key": "sender",
   "value": "gard1vka9svfaf5vg5cjaaju3d9g08ydk80z59ts0la"
  },
  {
   "key": "recipient",
   "value": "gard18vdg4r46qtfkwyghsg67dpn9p7vassw30z0f46"
  }
 ]
}
```
