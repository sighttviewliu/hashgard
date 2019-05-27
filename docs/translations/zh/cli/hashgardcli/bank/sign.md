# hashgardcli bank sign

## 描述

签名生成的离线传输文件。该文件由 --generate-only 标志生成。

## 用法

```shell
hashgardcli bank sign [file] [flags]
```

## Flags

| 名称       | 类型    | 必需 | 默认值                | 描述                                                         |
| ---------------- | ------- | -------- | --------------------- | ------------------------------------------------------------ |
| --append | bool | 否 | true | 将签名附加到现有签名。如果禁用，旧签名将被覆盖。如果 --multisig 打开则忽略（默认为 true ） |
| --multisig | string | 否 | | 代表交易签署的 multisig 帐户的地址 |
| --name | string | 否 | | 与之签名的私钥的名称 |
| --offline | bool | 否 | false | 链下模式，不查询全节点 |
| --output-document | string |  |  | 该文档将写入给定文件而不是 STDOUT  |
| --signature-only | bool | 否 | | 仅打印生成的签名，然后退出 |
| --validate-signatures | bool | 否 | false | 打印必须签署交易的地址，已签名的地址，并确保签名的顺序正确 |

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

### 对一个离线发送文件签名

首先你必须使用 **hashgardcli bank send**  命令和标志 **--generate-only** 来生成一个发送记录，如下：

```shell
hashgardcli bank send gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx 10gard --from=test --chain-id=hashgard --generate-only

{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/Send","value":{"inputs":[{"address":"gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx","coins":[{"denom":"gard","amount":"10000000000000000000"}]}],"outputs":[{"address":"gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx","coins":[{"denom":"gard","amount":"10000000000000000000"}]}]}}],"fee":{"amount":[{"denom":"gard","amount":"4000000000000000"}],"gas":"200000"},"signatures":null,"memo":""}}
```

保存输出到文件中，如：/root/node0/test_send_10hashgard.txt

接着来签名这个离线文件。

```shell
hashgardcli bank sign /root/node0/test_send_10hashgard.txt --name=test  --offline=false --print-response --append=true
```

随后得到签名详细信息，在输出中你会看到签名信息:
**ci+5QuYUVcsARBQWyPGDgmTKYu/SRj6TpCGvrC7AE3REMVdqFGFK3hzlgIphzOocGmOIa/wicXGlMK2G89tPJg==**
