# hashgardcli deposit search

## 描述
根据发行盒子的名字进行查询
## 用法
```shell
hashgardcli deposit search [name]
```
## Flags

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 搜索
```shell
hashgardcli deposit search  deposit
```

成功后，返回结果:
```shell
BoxID            |Owner                                       |Name            |TotalAmount                             |CreatedTime
boxab3jlxpt2pt   |gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60 |deposit88       |10000000000000000000000agard            |1559616171
```
