# hashgard unsafe-reset-all

## 描述

重置区块链数据库，删除地址簿文件，并将 priv_validator.json 重置为 genesis 状态

## 用法

```
hashgard unsafe-reset-all [flags]
```

## Flags

| 名称，缩写 | 默认值      | 描述                      | 必需 |
| ---------- | ----------- | ------------------------- | -------- |
| -h, --help |             | unsafe-reset-all 模块帮助 | 否       |
| --home     | ~/.hashgard | 配置和数据的目录          | 否       |

## 例子

``` shell
hashgard unsafe-reset-all
```