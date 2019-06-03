# hashgardcli future search

## 描述
根据发行盒子的名字进行查询
## 用法
```shell
hashgardcli future search [name]
```
## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 搜索
```shell
hashgardcli future search pay
```

成功后，返回信息：

```shell
BoxID            |Owner                                       |Name            |TotalAmount                             |CreatedTime
boxac3jlxpt2ps   |gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60 |pay-one         |1800000000000000000000agard             |1559550097
boxac3jlxpt2pt   |gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60 |pay-two         |1800000000000000000000agard             |1559552267

```
