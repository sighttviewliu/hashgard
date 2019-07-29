# hashgardcli keys flag-required

## 描述

设置账户转入必须填写备注信息

## 用法

```shell
hashgardcli keys flag-required [flag] [is_required] [flags]
```

## 子命令

| 名称, 速记       | 类型   | 必需 | 默认值 | 描述                                                              |
| --------------- | --------- | -------------------------- | ----------------------- | -------------------------- |
| is_required   | Boolean |   |   | 设置是否必须 true/false  |

## Flag

| 名称, 速记       | 类型   | 必需 | 默认值 | 描述                                                              |
| --------------- | --------- | -------------------------- | ----------------------- | -------------------------- |
| memo       | strings| 是 | 0 | 备注                                       |

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

### 设置账号转入必须填写备注

```shell
hashgardcli keys flag-required memo true --from $your_wallet_name
```

成功后，返回结果:

```txt
{
  "height": "260",
  "txhash": "5F5AFED2F5BA558D5110BDC74EC34AA5D969D39D79EBE46AFD981578A0243523",
  "raw_log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
  "logs": [
    {
      "msg_index": "0",
      "success": true,
      "log": ""
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "11421",
  "tags": [
    {
      "key": "action",
      "value": "account"
    },
    {
      "key": "category",
      "value": "account"
    },
    {
      "key": "sender",
      "value": "gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60"
    },
    {
      "key": "must-memo",
      "value": "true"
    }
  ]
}

```
