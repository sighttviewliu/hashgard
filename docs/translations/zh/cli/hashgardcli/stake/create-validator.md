# hashgardcli stake create-validator

## 介绍

发送交易申请成为验证人，并在在此验证人上委托一定数额的 token

## 用法

```shell
hashgardcli stake create-validator [flags]
```

## Flags

| 名称                         | 类型   | 是否必填 | 默认值 | 功能描述                                                |
| ---------------------------- | ------ | -------- | ------ | ------------------------------------------------------- |
| --address-delegator          | string | false    | ""     | 委托人地址                                              |
| --amount                     | string | true     | ""     | 委托 token 的数量                                       |
| --commission-max-change-rate | float  | false    | 0.0    | 佣金比率每天变化的最大数量                              |
| --commission-max-rate        | float  | false    | 0.0    | 最大佣金比例                                            |
| --commission-rate            | float  | false    | 0.0    | 初始佣金比例                                            |
| --details                    | string | false    | ""     | 验证人节点的详细信息                                    |
| --genesis-format             | bool   | false    | false  | 是否已 genesis transaction 的方式倒出                   |
| --identity                   | string | false    | ""     | 身份信息的签名                                          |
| --ip                         | string | false    | ""     | 验证人节点的 IP, 与`--generate-only`flag 同时使用时生效 |
| --min-self-delegation        | string | true     | ""     | 验证人节点要求的自委托最小股权数量                      |
| --moniker                    | string | true     | ""     | 验证人节点名称                                          |
| --node-id                    | string | false    | ""     | 节点 ID                                                 |
| --pubkey                     | string | true     | ""     | Amino 编码的验证人公钥                                  |
| --website                    | string | false    | ""     | 验证人节点的网址                                        |

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

```shell
hashgardcli stake create-validator \
--from=gard1rkqp5w062sdqu68qsufn3safwz0e5m9f3efaak \
--pubkey=gardvalconspub1zcjduepqm58zvp4fmssekze8m04ppyf55uwfhecnpe5lfs3znxtes2mhz8esrvvtqv \
--commission-max-change-rate=0.01 \
--commission-rate=0.1 \
--commission-max-rate=0.2 \
--amount=40gard \
--moniker=joelove \
--min-self-delegation=10 \
--chain-id=chain-HwMeL0 \
--fees=2gard \
--output=json \
--indent
```

成功后返回结果如下：

```json
{
    "height": "19195",
    "txhash": "0DCBF4DB2F64A625FD13166D323A5033D15D4CCFB07217F08E0CFDFF8FC29998",
    "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
    "gas_wanted": "200000",
    "gas_used": "99976",
    "tags": [
        {
            "key": "action",
            "value": "create_validator"
        },
        {
            "key": "destination-validator",
            "value": "gardvaloper1rkqp5w062sdqu68qsufn3safwz0e5m9frmy4dm"
        },
        {
            "key": "moniker",
            "value": "joelove"
        },
        {
            "key": "identity"
        }
    ]
}
```
