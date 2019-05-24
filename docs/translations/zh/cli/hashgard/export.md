# hashgard export

## 描述

将状态导出为 json 文件

## 用法

```
hashgard export [flags]
```

## Flags

| 名称，缩写        | 类型   | 默认值 | 描述           | 必需 |
| ----------------- | ------ | ------ | ---------- | -------- |
| -h, --help        |        |        | testnet 模块帮助                            | 否       |
| --for-zero-height |        |        | 导出的状态将用于一个新网络 (执行预处理)     | 否       |
| --height          | int    | -1     | 从指定块高度导出状态                        | 否       |
| --jail-whitelist  | string |        | List of validators to not jail state export | 否       |

## 例子

`hashgard export`
