# hashgardcli bank encode

## Description

Encode transactions created with the --generate-only flag and signed with the sign command.
Read a transaction from <file>, serialize it to the Amino wire protocol, and output it as base64.
If you supply a dash (-) argument in place of an input filename, the command reads from standard input.



## Usage

```shell
hashgardcli bank encode [file] [flags]
```


## Flags

**Global flags, query command flags** [hashgardcli](../README.md)



## Example

### Send token to the specified address

```shell
hashgardcli bank encode sign.json
```

After that, you can get remote node status as follows:

```shell
"2QLwYl3uCjwqLIf6ChSn6EEMEUpOn8aHmisX4xFXxWYbERIUKu/P18RXUuHr363+wt+UoWPHPGAaCgoEZ2FyZBICMTASBBDAmgwajgIKfiLB9+IIAhIm61rphyED30pWrLHFH6T+RX4kqgkSg8CvPDDkgSwwxpgMg2/CVB4SJuta6YchA5l0etexBHD8jaIC+QrpuVtxsRt5q1/3vx3ooQrZOOzCEibrWumHIQNRtYp0E4MQlDy4xrtq0zGNTCcGryjsh4yKOTIiThQP7RKLAQoFCAMSAWASQBJWW1uqYiw5nfvJhtVSz1WLkCva/+X4rbF7wzjbYmq1TxUs6n/A5G7MjwTgkDpn7jJIRbfktU6shclGbmhNuNoSQNo9kE5rVvHhLajjwJMnndI//e6vaYYN+ClfeYL36dMHe1dLpiqMo/xV/1k7w+4mDVktrLG8I6c7SLYIDnAk3gs="

```
