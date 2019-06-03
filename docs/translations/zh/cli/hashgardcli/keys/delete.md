# hashgardcli keys delete

## 描述

删除指定的密钥

## 用法
```shell
hashgardcli keys delete <name> [flags]
```

## Flags

| 速记,名称   | 类型 | 必需 | 默认值     | 描述          |
| -------- | ----- | ------ | ------- | --------------- |
| -f, --force | bool | 否 | false | 无条件删除密钥而不要求密码      |
| -y, --yes | bool | 否 | false | 跳过确认提示 |

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子

### 删除指定的密钥

```shell
hashgardcli keys delete MyKey
```

执行命令后，你会收到一个危险警告，并且要求你输入密码用于执行删除指令。

```txt
DANGER - enter password to permanently delete key:
```

输入正确的密码之后，你就完成了删除操作。

```txt
Password deleted forever (uh oh!)
```
