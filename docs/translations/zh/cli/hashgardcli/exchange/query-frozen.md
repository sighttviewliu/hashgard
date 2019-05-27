# hashgardcli exchange query-frozen

## 描述

查看指定地址的冻结资金

## 用法

```shell
hashgardcli exchange query-frozen [address] [flags]
```

## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

### 查询地址的冻结资金

```shell
hashgardcli exchange query-frozen gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa --chain-id hashgard -o=json --indent
```

下面是地址 gard1p48xfe62mwewxzuqpwkcdjyge42fck6xzc7xpa 的冻结资金

```txt
[
 {
  "denom": "apple",
  "amount": "33"
 },
 {
  "denom": "gard",
  "amount": "100"
 }
]
```
