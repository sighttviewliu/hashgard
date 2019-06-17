# hashgard start

## 描述

启动 hashgard 服务

## 用法

```shell
hashgard start [flags]
```

## Flags

| 名称，缩写                      | 默认值                | 描述          | 必需 |
| ------------------------------- | --------------------- | -------------------------- | -------- |
| --abci                          | socket                | 指定 abci 的传输方式，socket 或 grpc                  | 否       |
| --address                       | tcp://0.0.0.0:26658   | 监听地址                                              | 否       |
| --consensus.create_empty_blocks | true                  | 设置为 false 以仅在有 txs 或 AppHash 更改时生成块     | 否       |
| --fast_sync                     | true                  | 快速区块链同步                                        | 否       |
| --minimum_fees                  |                       | 验证人接受交易的最低费用                              | 否       |
| --moniker                       | instance-c5m0fg87     | 节点名称                                              | 否       |
| --p2p.laddr                     | tcp://0.0.0.0:26656   | 节点监听地址。 （0.0.0.0:0 表示任何接口，任何端口）   | 否       |
| --p2p.persistent_peers          |                       | 以逗号分隔的 ID @ host：端口持久对等体                | 否       |
| --p2p.pex                       | true                  | 启用/禁用 Peer-Exchange                               | 否       |
| --p2p.private_peer_ids          |                       | 以逗号分隔的私有对等 ID                               | 否       |
| --p2p.seed_mode                 |                       | 启用/禁用种子模式                                     | 否       |
| --p2p.seeds                     |                       | 逗号分隔的 ID @ host：端口种子节点                    | 否       |
| --p2p.upnp                      |                       | 启用/禁用 UPNP 端口转发                               | 否       |
| --priv_validator_laddr          |                       | 侦听来自外部 priv_validator 进程的连接的套接字地址    | 否       |
| --proxy_app                     | tcp://127.0.0.1:26658 | 代理应用程序地址，或“nilapp”或“kvstore”用于本地测试。 | 否       |
| --pruning                       | syncable              | pruning 策略：syncable, nothing, everything           | 否       |
| --replay                        |                       | 重播最后一个块                                        | 否       |
| --rpc.grpc_laddr                |                       | GRPC 监听地址（仅限 BroadcastTx）。需要端口           | 否       |
| --rpc.laddr                     | tcp://0.0.0.0:26657   | RPC 监听地址。需要端口                                | 否       |
| --rpc.unsafe                    |                       | 启用不安全的 rpc 方法                                 | 否       |
| --trace-store                   |                       | 启用 KVStore 跟踪到输出文件                           | 否       |
| --with-tendermint               | true                  | 用 tendermint 运行嵌入进程的 abci app                 | 否       |
| -h, --help                      |                       | start 模块的帮助文档                                  | 否       |

## 例子

```shell
hashgard start --home=/root/.hashgard
```
