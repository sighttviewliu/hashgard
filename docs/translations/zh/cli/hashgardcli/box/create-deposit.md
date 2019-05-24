# hashgardcli box create-deposit

## 描述
创建一个存款盒子，设定必要的参数。用来吸纳存款。盒子分为三个时期，1.发行期。2存款吸纳期。3.存款期。



## 用法
```
 hashgardcli box create-deposit [name][total-amount][flags] --from
```



### 命令解释

| 名称         | 类型   | 必需 | 默认值 | 描述                 |
| ------------ | ------ | -------- | ------ | -------------------- |
| name         | string | 是       |        | 盒子的名称      |
| total-amount | string | 是       |        | 接受存款的总量和种类 |



### Flags

| 名称             | 类型   | 必需 | 默认值 | 描述                           |
| ---------------- | ------ | -------- | ------ | ------------------------------ |
| --bottom-line    | int    | 是       | ""     | 达成存款计息的存款数量         |
| --price          | int    | 是       | ""     | 存款最小倍数且能被存款总量整除 |
| --start-time     | int    | 是       | ""     | 吸纳存款开始时间               |
| --establish-time | int    | 是       | ""     | 吸纳存款结束时间               |
| --maturity-time  | int    | 是       | ""     | 存款交割时间                   |
| --interest       | string | 是       | ""     | 注入利息的数量和种类           |

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子
### 创建存款盒子
```shell
hashgardcli box create-deposit mingone 10000coin174876e800  --bottom-line=0 --price=2  --start-time=1558079700  --establish-time=1558080300 --maturity-time=1558080900 --interest=9898coin174876e800  --from
```
输入正确的密码后，存款盒子创建完成。
```txt
  {
  Height: 4141
  TxHash: 9CDC3111A4FF78DB5F53CB5C6518025DB2B8DDB038BC2CB1C2E52FE9F2B1BD91
  Data: 0F0E626F786162336A6C787074327073
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 41233
  Tags:
    - action = box_create
    - category = box
    - box-id = boxab3jlxpt2ps
    - box-Type = deposit
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
    }
```



### 相关命令

| 名称                                        | 描述                         |
| ------------------------------------------- | ---------------------------- |
| [interest-injection](interest-injection.md) | 对存款盒子进行利息注入       |
| [interest-fetch](interest-fetch.md)         | 对存款盒子利息进行取回       |
| [deposit-to](deposit-to.md)                 | 用户对存款盒子进行存款       |
| [deposit-fetch](deposit-fetch.md)           | 用户在存款吸纳期进行取回存款 |
| [query-box](query-box.md)                   | 对指定盒子进行信息查询       |
| [list-box](list-box.md)                     | 罗列指定类型盒子列表         |
