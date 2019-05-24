# hashgardcli issue search

## 描述
根据代币符号来搜索发行的代币信息
## 用法
```
hashgardcli issue search [symbol] [flags]
```
## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子
### 搜索
```shell
hashgardcli issue search AAA
```
```txt
 [
    {
        "issue_id":"coin174876e802",
        "issuer":"gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7",
        "owner":"gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7",
        "issue_time":"1558179518",
        "name":"issuename",
        "symbol":"AAA",
        "total_supply":"10000000001023",
        "decimals":"18",
        "description":"{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"新一代金融公有链"}",
        "burn_owner_disabled":false,
        "burn_holder_disabled":false,
        "burn_from_disabled":false,
        "freeze_disabled":false,
        "minting_finished":false
    }
]

```
