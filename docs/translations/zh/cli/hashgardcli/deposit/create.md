# hashgardcli  deposit create

## 描述
创建一个存款盒子，设定必要的参数。用来吸纳存款。盒子分为三个时期，1. 发行期。2. 存款吸纳期。3. 存款期。4.交割期



## 用法
```shell
hashgardcli  deposit create [name] [total-amount] [flags]
```



### 命令解释

| 名称         | 类型   | 必需 | 默认值 | 描述                 |
| ------------ | ------ | -------- | ------ | -------------------- |
| Name         | string | 是       |        | 盒子的名称      |
| total-amount | string | 是       |        | 接受存款的总量和种类 |



### Flags

| 名称             | 类型   | 必需 | 默认值 | 描述                           |
| ---------------- | ------ | -------- | ------ | ---------------------- |
| --bottom-line    | int    | 是       |    | 达成存款计息的存款数量         |
| --price          | int    | 是       |    | 存款最小倍数且能被存款总量整除 |
| --start-time     | int    | 是       |    | 吸纳存款开始时间               |
| --establish-time | int    | 是       |    | 吸纳存款结束时间               |
| --maturity-time  | int    | 是       |    | 存款及利息交割时间                   |
| --interest       | string | 是       |    | 注入利息的数量和种类           |

**全局 flags、查询命令 flags** 参考:[hashgardcli](../README.md)

## 例子
### 创建存款盒子
```shell
hashgardcli deposit create  mingone 10000coin174876e800  --bottom-line=0 --price=2  --start-time=1558079700  --establish-time=1558080300 --maturity-time=1558080900 --interest=9898coin174876e800  --from
```
成功后，返回结果:
```txt
  {
    Response:
      Height: 3353
      TxHash: D14BFA58729BEEB1FFE113108404B465757F1C8F7B4041DD91EAEA021870873A
      Data: 0F0E626F786162336A6C787074327074
      Raw Log: [{"msg_index":"0","success":true,"log":""}]
      Logs: [{"msg_index":0,"success":true,"log":""}]
      GasWanted: 200000
      GasUsed: 60857
      Tags:
        - action = create
        - category = deposit
        - id = boxab3jlxpt2pt
        - sender = gard1prflhd5h66l498vdyy95hyh898r0tjxvv6vc60
        - fee = 1000000000000000000000agard
    }
```



### 相关命令

| 名称                                        | 描述                         |
| ------------------------------------------- | ---------------------------- |
| [interest-inject](interest-inject.md) | 对存款盒子进行利息注入       |
| [interest-cancel](interest-cancel.md)         | 对存款盒子利息进行取回       |
| [inject](inject.md)                 | 用户对存款盒子进行存款       |
| [cancel](cancel.md)           | 用户在存款吸纳期进行取回存款 |
| [params](params.md)                   |  参数及费用  |
| [query](query.md)                     |   查询指定 id 盒子   |
| [list](list.md)                     |   查询全部盒子列表  |
| [search](search.md)                     |   按盒子名称进行搜索    |
