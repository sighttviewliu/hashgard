# hashgardcli stake edit-validator

## 介绍

编辑已有验证人账户，包括佣金比率，验证人节点名称以及其他描述信息等

## 用法

```shell
hashgardcli stake edit-validator [flags]
```

## Flags

| 名称              | 类型   | 是否必填 | 默认值            | 功能描述             |
| ----------------- | ------ | -------- | ----------------- | -------------------- |
| --commission-rate | float  | false    | 0.0               | 佣金比率             |
| --details         | string | false    | | 验证人节点的详细信息 |
| --identity        | string | false    | | 身份签名             |
| --moniker         | string | false    | | 验证人名称           |
| --website         | string | false    | | 网址                 |
**验证人信息设置后,请尽量不要修改。**

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

```shell
hashgardcli stake edit-validator \
--from=hashgard \
--chain-id=hashgard \
--website=http://hashgard.com \
--details=hashgard_is_great
```
