# hashgard testnet

## 描述

在单机上生成多节点的 Hashgard 公链网络配置文件

## 用法

```shell
hashgard testnet [flags]
```

## Flags

| 名称，缩写            | 类型   | 默认值       | 描述                                                 | 必需 |
| --------------------- | ------ | ------------ | ---------------------------------------------- | ----- |
| -h, --help            |        |              | testnet 模块帮助                                     | 否       |
| --chain-id            | string |              | genesis file chain-id，如果留空则将被随机创建        | `是`     |
| --minimum-gas-prices  | string | 0.000006gard | 交易中要求的 gas price 最小值                        | `是`     |
| --node-cli-home       | string | hashgardcli  | 节点的 cli 配置的主目录                              | 否       |
| --node-daemon-home    | string | hashgard     | 节点守护程序配置的主目录                             | 否       |
| --node-dir-prefix     | string | falsede         | 使用（ node0，node1，...）为每个节点目录添加前缀名称 | 否       |
| --output-dir          | string | ./mytestnet  | 用于存储 testnet 初始化数据的目录                    | 否       |
| --starting-ip-address | string | 192.168.0.1  | 起始 IP 地址                                         | 否       |
| --v                   | int    | 4            | 初始化 testnet 的验证器数量                          | 否       |

## 例子

`hashgard testnet--chain-id=${chain-id}`
