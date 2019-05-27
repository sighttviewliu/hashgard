# hashgard validate-genesis

## 描述

验证 genesis 文件内容的有效性

## 用法

```shell
hashgard validate-genesis [file] [flags]
```

## 参数

| 名称，缩写 | 类型   | 默认值                          | 描述             | 必需 |
| ---------- | ------ | ------------------------------- | ---------------- | -------- |
| [file]     | string | ~/.hashgard/config/genesis.json | genesis 文件位置 | 否       |

## Flags

| 名称，缩写 | 类型   | 默认值      | 描述                             | 必需 |
| ---------- | ------ | ----------- | -------------------------------- | -------- |
| -h, --help |        |             | add-genesis-account 模块帮助文档 | 否       |
| --home     | string | ~/.hashgard | 配置和数据的目录                 | 否       |
| --trace    | bool   |             | 在出错时打印完整的调用栈         | 否       |

## 例子

```shell
hashgard validate-genesis
```
