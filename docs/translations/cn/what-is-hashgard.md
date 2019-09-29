# Hashgard是什么

`Hashgard`是 Cosmos 跨链生态中的金融公链。它有两个主要的入口：

+ `hashgard` : Hashgard 的服务进程，运行着`hashgard`程序的全节点。
+ `hashgardcli` : Hashgard 的命令行界面，用于同一个 Hashgard 的全节点交互。

`hashgard`基于Cosmos SDK构建，使用了如下模块:

+ `x/auth` : 账户和签名
+ `x/bank` : token转账
+ `x/staking` : 抵押逻辑
+ `x/mint` : 增发通胀逻辑
+ `x/distribution` : 费用分配逻辑
+ `x/slashing` : 处罚逻辑
+ `x/gov` : 治理逻辑
+ `x/ibc` : 跨链交易
+ `x/params` : 处理应用级别的参数
+ `x/contract` : 智能合约

接着，学习如何[安装Hashgard](./installation.md)