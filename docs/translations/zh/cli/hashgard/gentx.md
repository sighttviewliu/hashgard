# hashgard gentx

## 描述

此命令是 'hashgard tx create-validator' 命令的别名。
它会创建一个带有自我委托的创世交易，
同时带有以下委托和佣金默认参数：

```
授权金额：100gard
佣金率：0.1
佣金最高税率：0.2
佣金最高变动率：0.01
最小自委托股权：1股 (100gard)
```

## 用法

```
 hashgard gentx [flags]
```

## Flags

| 名称，缩写                   | 类型   | 默认值         | 描述                    | 必需 |
| ---------------------------- | ------ | -------------- | --------------------------- | -------- |
| --amount                     | string |                | 债券的金额                                          | `是`     |
| --commission-max-change-rate | string |                | 最高佣金变动率百分比（每天）                        | 否       |
| --commission-max-rate        | string |                | 最高佣金率百分比                                    | 否       |
| --commission-rate            | string |                | 初始佣金率百分比                                    | 否       |
| --home-client                | string | ~/.hashgardcli | 客户端的主目录                                      | 否       |
| --ip                         | string | 本机 IP 地址   | 节点公网 IP 地址                                    | 否       |
| --min-self-delegation        | string |                | 验证人最小委托股份                                  | `是`     |
| --name                       | string |                | 用于签署 gentx 的私钥的名称                         | 否       |
| --node-id                    | string |                | 节点 ID                                             | 否       |
| --output-document            | string |                | 将 genesis 事务 JSON 文档写入给定文件而不是默认位置 | 否       |
| --pubkey                     | string |                | 验证人的 Bech32 编码的 PubKey                       | 否       |
| -h, --help                   |        |                | gentx 模块帮助文档                                  | 否       |
| --home                       | string | ~/.hashgard    | 配置和数据的目录                                    | 否       |
| --trace                      | bool   |                | 在出错时打印完整的调用栈                            | 否       |

## 例子

`hashgard gentx --name=root --amount=10000gard --ip=${validator_ip}`
