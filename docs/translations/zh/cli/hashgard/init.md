# hashgard init

## 描述

生成验证人和全节点所需要的配置文件和数据文件。

## 用法

```shell
hashgard init [flags]
```

## Flags

| 名称，缩写  | 类型   | 默认值      | 描述                            | 必需 |
| ----------- | ------ | ----------- | ------------------------------------- | -------- |
| -h, --help  |        |             | init 模块帮助                                         | 否       |
| --chain-id  | string |             | 公链 ID，如果留空则将被随机创建                          | 否       |
| --moniker   | string |             | 设置节点名称，将在浏览器的[验证人节点](https://www.gardplorer.io/validator)页面显示 | `是`     |
| --overwrite | bool   |             | 覆盖 genesis.json 文件          | 否       |
| --home      | string | ~/.hashgard | 配置和数据存放目录               | 否       |
| --trace     | bool   |             | 在出错时打印完整的调用栈                                                            | 否       |

## 例子

`hashgard init --chain-id=testnet-1000 --moniker=hashgard`
