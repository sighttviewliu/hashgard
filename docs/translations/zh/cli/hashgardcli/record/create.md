# hashgardcli record create

## 描述
用户存储存证
## 用法
```shell
hashgardcli record create [name] [hash] [flags]
```
### 子命令

| 名称         | 类型   | 必需 | 默认值 | 描述                 |
| ------------ | ------ | -------- | ------ | -------------------- |
| name         | string | 是       |        | 存证的名称     |
| hash  | string | 是       |        | 需要存储的 hash 值 长度小于等于 64 位且唯一|

## Flags

| 名称, 速记       | 类型   | 必需 | 默认值 | 描述                                                              |
| --------------- | --------- | -------------------------- | ----------------------- | -------------------------- |
| --author      | string  | 否 |  |     存证归属的作者   |
| --record-type   | string  |  否  |   | 存证归属的类型  |
| --description    | string   | 否  |   | 存证归属的描述  |

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 创建
```shell
hashgardcli record create myRecord CC38CAEE32149BEF4CCFAEAB518EC9A5FBC85AE6AC8D5A9F6CD710FAF5E4A22b --from one --author Hashgard --record-type Transaction --description "hashgard:Good Projects，GARD:Good currency" -y
```
返回结果
```shell
{
  "height": "104",
  "txhash": "19581EB61CD91B797C7009410329FA6BC423DAE650D62A28B8FCADC18430E303",
  "raw_log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
  "logs": [
    {
      "msg_index": "0",
      "success": true,
      "log": ""
    }
  ],
  "gas_wanted": "200000",
  "gas_used": "26002",
  "tags": [
    {
      "key": "action",
      "value": "record"
    },
    {
      "key": "category",
      "value": "record"
    },
    {
      "key": "sender",
      "value": "gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60"
    },
    {
      "key": "id",
      "value": "rec174876e800"
    },
    {
      "key": "hash",
      "value": "CC38CAEE32149BEF4CCFAEAB518EC9A5FBC85AE6AC8D5A9F6CD710FAF5E4A22b"
    }
  ]
}
```
