# 命令行客户端

## 全局可用的 flags

| 名称, 缩写   | 类型   | 必需  | 默认值                | 描述                                                               |
| ------------ | ------ | ----- | --------------------- | ------------------------------------------------------------------ |
| --chain-id   | string | false | ""                    | Tendermint 节点的 Chain ID                                         |
| --encoding   | string | false | "hex"                 | 字节编码 (hex, b64, btc)                                           |
| --home       | string | false | "\$HOME/.hashgardcli" | 配置文件和区块数据存放目录                                         |
| --help, -h   | string | false |                       | 打印帮助信息                                                       |
| --indent     | bool   | false | false                 | 格式化 json 字符串                                                 |
| --ledger     | bool   | false | false                 | 是否使用 ledger 硬件钱包                                           |
| --node       | string | false | tcp://localhost:26657 | tendermint 节点的 rpc 地址                                         |
| --output     | string | false | text                  | 输出格式 (text, json)                                              |
| --trace      | bool   | false |                       | 错误输出完整栈信息                                                 |
| --trust-node | bool   | false | true                  | 是否信任全节点返回的数据，如果不信任，客户端会验证查询结果的正确性 |

## 查询命令的 flags

| 名称, 缩写 | 类型 | 必需  | 默认值 | 描述                                                       |
| ---------- | ---- | ----- | ------ | ---------------------------------------------------------- |
| --height   | int  | false | 0      | 查询某个高度的区块链数据，如果是 0，这返回最新的区块链数据 |

每个区块链状态查询命令都包含上表中的 flags，同时不同查询命令还可能会有自己独有的 flags。

## 发送交易命令的 flags

| 名称, 缩写       | 类型   | 必需  | 默认值 | 描述                                                                                       |
| ---------------- | ------ | ----- | ------ | ------------------------------------------------------------------------------------------ |
| --account-number | int    | false | 0      | 发起交易的账户的编号                                                                       |
| --async          | bool   | false | false  | 是否异步广播交易                                                                           |
| --dry-run        | bool   | false | false  | 模拟执行交易，并返回消耗的`gas`。`--gas`指定的值会被忽略                                   |
| --fees           | string | false | ""     | 交易费，例如: 10gard, 1atom                                                                |
| --from           | string | false | ""     | 发送交易的账户名称或地址                                                                   |
| --gas            | string | false | 200000 | 交易的 gas 上限; 设置为"auto"将自动计算相应的阈值                                          |
| --gas-adjustment | float  | false | 1      | gas 调整因子，这个值将乘以模拟执行估算的`gas`; 如果`--gas`的值不是`auto`，这个标志将被忽略 |
| --gas-prices     | string | false | ""     | 决定交易费(`fees`)的 gas 价格，例如：`0.001gard`;                                          |
| --generate-only  | bool   | false | false  | 是否仅仅构建一个未签名的交易便返回                                                         |
| --memo           | string | false | ""     | 指定交易的 memo 字段                                                                       |
| --print-response | bool   | false | true   | 是否打印交易返回结果，仅在`---async=false`的情况下有效                                     |
| --sequence int   | int    | false | 0      | 发起交易的账户的 sequence                                                                  |

每个发送交易的命令都包含上表中的 flags，同时不同交易的命令还可能会有自己独有的 flags。

## 模块列表

1. [bank command](./bank/README.md)
2. [distribution command](./distribution/README.md)
3. [gov command](./gov/README.md)
4. [keys command](./keys/README.md)
5. [stake command](./stake/README.md)
6. [status command](./status.md)
7. [tendermint command](./tendermint/README.md)
8. [slashing command](./slashing/README.md)
9. [issue command](./issue/README.md)
10. [box command](./box/README.md)
11. [faucet command](./faucet/send.md)

## hashgardcli config 命令

`config` 命令可以交互式地配置一些默认参数，例如 chain-id，home，fees 和 node。

`hashgardcli config <key> [value] [flags]`

## hashgardcli init 命令

`hashgardcli init` 命令可以初始化客户端
