# hashgardcli record query

## 描述
按照存证中的 hash 进行查询
## 用法
```shell
hashgardcli record query [hash] [flags]
```


### 子命令

| 名称         | 类型   | 必需 | 默认值 | 描述                 |
| ------------ | ------ | -------- | ------ | -------------------- |
| hash        | string | 是       |        | 需要查询的 hash   |


**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

```shell
hashgardcli record query CC38CAEE32149BEF4CCFAEAB518EC9A5FBC85AE6AC8D5A9F6CD710FAF5E4A22b
```
返回结果
```shell
{
  "type": "record/RecordInfo",
  "value": {
    "id": "rec174876e800",
    "hash": "CC38CAEE32149BEF4CCFAEAB518EC9A5FBC85AE6AC8D5A9F6CD710FAF5E4A22b",
    "record_number": "",
    "sender": "gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60",
    "record_time": "1564135537",
    "name": "myRecord",
    "author": "Hashgard",
    "record_type": "Transaction",
    "description": ""
  }
}
```
