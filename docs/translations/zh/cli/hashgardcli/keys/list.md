# hashgardcli keys list

## 描述

返回此密钥管理器存储的所有公钥的列表以及他们的相关名称和地址。

## 用法

```shell
hashgardcli keys list [flags]
```
## Flags
 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)


## 例子

### 列出所有的密钥

```shell
hashgardcli keys list
```

执行命令后，你会得到所有存于本地客户端存储的密钥，包括它们的地址和公钥信息。

```txt
NAME:	Type:	ADDRESS:						            PUBKEY:
abc  	local	gardva2eu9qhwn5fx58kvl87x05ee4qrgh44yd8teh	gardpub1addwnpepqvu549hgyhnxlveqmtdn2xywygxpgzcsqefxur47zkz4e0e9x67hvjr6r6p
```

需要注意的是，如果本地有多个.hashgardcli 存储，需要通过--home 参数来定位查询源。
