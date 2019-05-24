# hashgardcli status

## 描述

查询远程节点状态

## 用法

```shell
hashgardcli status [flags]
```

## Flags

| 名称, 缩写 | 默认值                | 描述                | 是否必需 |
| ---------- | --------------------- | ------------------- | ---- |
| --node, -n | tcp://localhost:26657 | [string] 连接的节点 | 否   |

**全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

查询远程节点状态

```shell
hashgardcli status --node=tcp://localhost:26657 --indent
```

查询到的结果如下：

```json
{
    "node_info": {
        "protocol_version": { "p2p": "5", "block": "8", "app": "0" },
        "id": "b783ac2b41da096ddc3a26c2a39e3b0c1ea49d9e",
        "listen_addr": "47.244.32.147:26656",
        "network": "hashgard",
        "version": "0.27.0",
        "channels": "4020212223303800",
        "moniker": "hashgard_test",
        "other": { "tx_index": "on", "rpc_address": "tcp://0.0.0.0:26657" }
    },
    "sync_info": {
        "latest_block_hash": "C91F7F46109F09E50B399893D666DF635944F0743F532F1D7D47B8D6CAD63926",
        "latest_app_hash": "AF897AEDDF4C24A4DEC38EDC76F2CC32D9B419C3BF75D922F95A48B35332D108",
        "latest_block_height": "13379",
        "latest_block_time": "2018-12-18T09:03:09.164197624Z",
        "catching_up": false
    },
    "validator_info": {
        "address": "5959DF3D28F601407A98D0B038174E288ED45BD7",
        "pub_key": { "type": "tendermint/PubKeyEd25519", "value": "wYrxKp7cw14eQiqzfGBggEV474ZA2lc35AieJM5SM6Y=" },
        "voting_power": "0"
    }
}
```
