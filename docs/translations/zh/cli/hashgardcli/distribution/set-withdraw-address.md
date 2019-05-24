# hashgardcli distribution set-withdraw-addr

## 介绍

设置取回收益时的收款地址

## 用法

```
hashgardcli distribution set-withdraw-addr [withdraw-address] [flags]
```
## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)


## 例子
```
hashgardcli distribution set-withdraw-addr gard1c9vrvvz08hd4entr0y5kfrt43v6malv60qtjfl --from joehe1 --chain-id=hashgard --output=json --indent
```

执行结果例子：

```
{
 "height": "33500",
 "txhash": "58AB9D329A043FC86DCE2B66E91BEDC1D13DD4000DF22E290041214C56DB04B8",
 "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
 "gas_wanted": "200000",
 "gas_used": "12018",
 "tags": [
  {
   "key": "action",
   "value": "set_withdraw_address"
  },
  {
   "key": "delegator",
   "value": "gard1c9vrvvz08hd4entr0y5kfrt43v6malv60qtjfl"
  }
 ]
}
```
