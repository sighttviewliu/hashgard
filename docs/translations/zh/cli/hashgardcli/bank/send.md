# hashgardcli bank send

## 描述

发送通证到指定地址

## 用法

```shell
hashgardcli bank send [to_address] [amount] [flags]
```
## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

### 发送通证到指定地址

```shell
hashgardcli bank send gard1c9vrvvz08hd4entr0y5kfrt43v6malv60qtjfl 10gard --from=hashgard --chain-id=hashgard --indent -o json
```

成功后，返回结果:

```txt
{
 "height": "21667",
 "txhash": "58110E97BD93CFA123B43B7C893386BA26F238570E1131A7B6E1E6ED5B7DA605",
 "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
 "gas_wanted": "200000",
 "gas_used": "22344",
 "tags": [
  {
   "key": "action",
   "value": "send"
  },
  {
   "key": "sender",
   "value": "gard10tfnpxvxjh6tm6gxq978ssg4qlk7x6j9aeypzn"
  },
  {
   "key": "recipient",
   "value": "gard1c9vrvvz08hd4entr0y5kfrt43v6malv60qtjfl"
  }
 ]
}
PS

```
