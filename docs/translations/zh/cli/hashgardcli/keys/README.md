# hashgardcli keys

## 描述

Keys 模块用于管理本地密钥库。

## 用法

```shell
hashgardcli keys [command]
```

## 相关命令

| 名称                    | 描述                                                 |
| ----------------------- | ------------------------------------------- |
| [mnemonic](mnemonic.md) | 通过读取系统熵来创建 bip39 助记词，也可以称为种子短语。         |
| [add](add.md)           | 创建新密钥，或从助记词导入已有密钥                         |
| [list](list.md)         | 列出所有密钥                                           |
| [show](show.md)         | 显示指定名称的关键信息                                     |
| [delete](delete.md)     | 删除指定的密钥                                         |
| [update](update.md)     | 更改用于保护私钥的密码                                    |

## Flags
 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)


## 补充说明

这些密钥可以是 go-crypto 支持的任何格式，并且可以由轻客户端，完整节点或需要使用私钥签名的任何其他应用程序使用。
