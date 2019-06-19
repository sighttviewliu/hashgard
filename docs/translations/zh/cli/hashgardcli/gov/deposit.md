# hashgardcli gov deposit

## 描述

充值保证金以激活提案

## 用法

``` shell
hashgardcli gov deposit [proposal-id] [deposit] [flags]
```
## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

### 充值保证金

```shell
 hashgardcli gov deposit  1 50gard --from=hashgard --chain-id=hashgard -o=json --indent
```

输入正确的密码后，你就充值了 50 个 gard 用以激活提案的投票状态。

```txt
{
 "height": "106707",
 "txhash": "1D048A63AB37015700F22C5C90DA79127E0FFDBC8A9F5D2418B00D1916389B74",
 "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
 "gas_wanted": "200000",
 "gas_used": "32733",
 "tags": [
  {
   "key": "action",
   "value": "deposit"
  },
  {
   "key": "depositor",
   "value": "gard10tfnpxvxjh6tm6gxq978ssg4qlk7x6j9aeypzn"
  },
  {
   "key": "proposal-id",
   "value": "3"
  }
 ]
}
```

如何查询保证金充值明细？

请点击下述链接：

[query-deposit](query-deposit.md)

[query-deposits](query-deposits.md)
