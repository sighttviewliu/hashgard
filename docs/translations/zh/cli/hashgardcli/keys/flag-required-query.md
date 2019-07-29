# hashgardcli keys flag-required-query

## 描述

查询指定地址设置交易强制备注信息状态

## 用法

```shell
hashgardcli keys flag-required-query [flag] [address] [flags]
```

## 子命令

| 名称, 速记       | 类型   | 必需 | 默认值 | 描述                                                              |
| --------------- | --------- | -------------------------- | ----------------------- | -------------------------- |
| memo       | strings | 是 | 0 | 备注                                       |



## Flag

| 名称, 速记       | 类型   | 必需 | 默认值 | 描述                                                              |
| --------------- | --------- | -------------------------- | ----------------------- | -------------------------- |
| address  | strings|   |   | 查询的地址是否开启强制转入信息填写功能 |



**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

### 按账户地址进行查询

```shell
hashgardcli keys flag-required-query memo gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
```

成功后，返回结果:

```txt
{
  "sender": "gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60",
  "memo_required": true
}
```
