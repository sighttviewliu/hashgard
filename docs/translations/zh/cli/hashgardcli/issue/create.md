# hashgardcli issue create

## 描述

发行一个新的代币

## 用法

```shell
hashgardcli issue create [name] [symbol] [total-supply] [flags] --from
```

## Flags

| 名称          | 类型 | 必需 | 默认值 | 描述                                    |
| ------------- | ---- | -------- | ------ | --------------------------------------- |
| --decimals    | int  | 否       | 18     | （可选）代币精度，默认 18 位，最大 18 位    |
| --burn-owner  | Bool | 否       | false  | 关闭代币所有者销毁自己持有的代币功能    |
| --burn-holder | bool | 否       | false  | 关闭普通账号销毁该自己持有的代币功能    |
| --burn-from   | bool | 否       | false  | 关闭 Owner 销毁非管理员账户持有的代币功能 |
| --minting     | bool | 否       | false  | 是否不再增发功能                        |
| --freeze      | bool | 否       | false  | 关闭冻结用户转入转出功能                |

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

### 发行一个新币

```shell
hashgardcli issue create issuename AAA 10000000000000 --from
```

输入正确的密码之后，你就完成发行了一个代币，需要注意的是要记下你的 issue-id 值，这是可以检索及操作你的代币的唯一要素。

```txt
{
   Height: 2967
  TxHash: 84B19F831958A6334C4806967E66E6C8640F0A2E7958A5E99A1DF3B6B6E6378C
  Data: 0F0E636F696E31373438373665383032
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 43428
  Tags:
    - action = issue
    - category = issue
    - issue-id = coin174876e802
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
}
```

查询自己的账号

```shell
hashgardcli bank account gard1f203m5q7hr4tkf0vredrn4wpxkx7zngn4pntye
```

你将会看到你的持币列表里多了一个形如“币名（issue-id）”特殊名称的币。后续对该币的操作请使用 issue-id 的值来进行，包括进行转账操作，要转的币也请使用该 issue-id。

```txt
{
 Account:
  Address:       gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Pubkey:        gardpub1addwnpepqfpd8mkl3jg43fw7y02fe99cgaxutf5npv9y9gx9dvrrcdwl36shv694apw
  Coins:         9999999990001issuename(coin174876e802)
  AccountNumber: 0
  Sequence:      16
}
```
