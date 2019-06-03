# hashgardcli keys add

## 描述

创建一个新密钥，或通过助记词导入已有密钥

## 用法

```shell
hashgardcli keys add <name> [flags]
```

## Flags

| 名称, 速记       | 类型   | 必需 | 默认值 | 描述                                                              |
| --------------- | --------- | -------------------------- | ----------------------- | -------------------------- |
| --account       | int | 否 | 0 | [uint32] HD 推导的账号                                              |
| --index         | int | 否 | 0 | [uint32] HD 推导的索引号                                            |
| --interactive | string | 否 || 以交互方式提示用户输入 BIP39 密码和助记符 |
| --multisig | strings | 否 || 构造并存储 multisig 公钥（当使用了 --pubkey） |
| --multisig-threshold | int | 否 | 1 | N 个中的 K 个需要签名。与 --multisig 一起使用 |
| --no-backup     | bool | 否 | false | 不输出助记词（如果其他人正在看着操作终端）                              |
| --nosort | bool | 否 | false | 传递给 --multisig 的密钥按照它们提供的顺序获取 |
| --pubkey | string | 否 || 解析 bech32 格式的公钥并将其保存到磁盘 |
| --recover       | string | 否 || 提供助记词以恢复现有密钥而不是新建                                     |

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

### 创建密钥

```shell
hashgardcli keys add MyKey
```

执行命令后，系统会要求你输入密钥密码，注意：密码必须至少为 8 个字符。

```txt
Enter a passphrase for your key:
Repeat the passphrase:
```

之后，你已经完成了创建新密钥的工作，但请记住备份你的助记词短语，如果你不慎忘记密码或丢失了密钥，这是唯一能恢复帐户的方法。

```txt
NAME:	Type:	ADDRESS:						PUBKEY:
MyKey	local	gard1m3m4l6g5774qe5jj8cwlyasue22yh32jf4wwet	gardpub1addwnpepqvu549hgyhnxlveqmtdn2xywygxpgzcsqefxur47zkz4e0e9x67hvjr6r6p
**Important** write this seed phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

oval green shrug term already arena pilot spirit jump gain useful symbol hover grid item concert kiss zero bleak farm capable peanut snack basket
```

上面 24 个单词只是助记词的例子，**不要**在生产环境中使用。

### 通过助记词恢复密钥

如果你忘记了密码或丢失了密钥，或者你想在其他地方使用密钥，则可以通过助记词短语来恢复。

```txt
hashgardcli keys add MyKey --recover
```

系统会要求你输入并确认密钥的新密码，然后输入助记词。这样就能恢复你的密钥。

```txt
Enter a passphrase for your key:
Repeat the passphrase:
Enter your recovery seed phrase:
```
