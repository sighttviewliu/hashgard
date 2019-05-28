# hashgardcli stake delegate

## 介绍

发送委托交易

## 用法

```shell
hashgardcli stake delegate [validator-addr] [amount] [flags]
```

## Flags

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

```shell
hashgardcli stake delegate \
gardvaloper1rkqp5w062sdqu68qsufn3safwz0e5m9frmy4dm \
80gard \
--from=hashgard \
--fees=2gard \
--chain-id=hashgard \
--output=json
--indent
```

然后你将会得到如下消息：

```json
{
    "height": "19574",
    "txhash": "B221ABC5E47685281DE6EE8A62EF286A654A85CAD12BBA13F961932129C4A271",
    "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
    "gas_wanted": "200000",
    "gas_used": "95865",
    "tags": [
        {
            "key": "action",
            "value": "delegate"
        },
        {
            "key": "delegator",
            "value": "gard19ddgmeywcg004aq92vhgrzav4mdvnntuz567yj"
        },
        {
            "key": "destination-validator",
            "value": "gardvaloper1rkqp5w062sdqu68qsufn3safwz0e5m9frmy4dm"
        }
    ]
}
```
