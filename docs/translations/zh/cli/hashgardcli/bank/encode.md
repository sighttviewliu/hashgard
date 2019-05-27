# hashgardcli bank encode

## 描述

对使用--generate-only标志创建的事务进行编码，并使用sign命令进行签名。
从<file>读取事务，将其序列化为Amino wire协议，并将其输出为base64。
如果您提供短划线（ - ）参数来代替输入文件名，则该命令将从标准输入读取。
## 用法

```
hashgardcli bank encode [file] [flags]
```
## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

### 发送通证到指定地址

```
hashgardcli bank encode sign.json
```

命令执行完成后，返回执行的细节信息

```
"2QLwYl3uCjwqLIf6ChSn6EEMEUpOn8aHmisX4xFXxWYbERIUKu/P18RXUuHr363+wt+UoWPHPGAaCgoEZ2FyZBICMTASBBDAmgwajgIKfiLB9+IIAhIm61rphyED30pWrLHFH6T+RX4kqgkSg8CvPDDkgSwwxpgMg2/CVB4SJuta6YchA5l0etexBHD8jaIC+QrpuVtxsRt5q1/3vx3ooQrZOOzCEibrWumHIQNRtYp0E4MQlDy4xrtq0zGNTCcGryjsh4yKOTIiThQP7RKLAQoFCAMSAWASQBJWW1uqYiw5nfvJhtVSz1WLkCva/+X4rbF7wzjbYmq1TxUs6n/A5G7MjwTgkDpn7jJIRbfktU6shclGbmhNuNoSQNo9kE5rVvHhLajjwJMnndI//e6vaYYN+ClfeYL36dMHe1dLpiqMo/xV/1k7w+4mDVktrLG8I6c7SLYIDnAk3gs="

```
