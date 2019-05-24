# hashgardcli box list-box

## 描述
查询指定类型盒子列表

## 用法
```shell
hashgardcli box list-box [flag]
```

### flag

| 名称    | 描述  |
| ------- | -------- |
| lock    | 锁仓盒子 |
| deposit | 存款盒子 |
| future  | 远期支付盒子  |"type"
 ### 全局flags 参考：[hashgardcli](../README.md)


## 例子
### 查询盒子信息

```shell
hashgardcli box list-box future
```

返回盒子信息

```txt
[
    {
        "box_id":"boxac3jlxpt2pt",
        "box_status":"depositing",
        "owner":"gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7",
        "name":"PayBox",
        "type":"future",
        "created_time":"1558182333",
        "total_amount":{
            "token":{
                "denom":"agard",
                "amount":"1800000000000000000000"
            },
            "decimals":"1"
        },
        "description":"{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"新一代金融公有链"}",
        "trade_disabled":true,
        "future":{
            "mini_multiple":"1",
            "deposits":null,
            "time":[
                "1657912000",
                "1657912001",
                "1657912002"
            ],
            "receivers":[
                [
                    "gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj",
                    "100000000000000000000",
                    "200000000000000000000",
                    "300000000000000000000"
                ],
                [
                    "gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0",
                    "100000000000000000000",
                    "200000000000000000000",
                    "300000000000000000000"
                ],
                [
                    "gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz",
                    "100000000000000000000",
                    "200000000000000000000",
                    "300000000000000000000"
                ]
            ],
            "distributed":null
        }
    },
    {
        "box_id":"boxac3jlxpt2ps",
        "box_status":"actived",
        "owner":"gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7",
        "name":"pay",
        "type":"future",
        "created_time":"1558090817",
        "total_amount":{
            "token":{
                "denom":"agard",
                "amount":"1800000000000000000000"
            },
            "decimals":"1"
        },
        "description":"",
        "trade_disabled":true,
        "future":{
            "mini_multiple":"1",
            "deposits":[
                {
                    "address":"gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7",
                    "amount":"1800000000000000000000"
                }
            ],
            "time":[
                "1657912000",
                "1657912001",
                "1657912002"
            ],
            "receivers":[
                [
                    "gard1cyxhqanlxc3u9025ngz5awzzex2jys6xc96ktj",
                    "100000000000000000000",
                    "200000000000000000000",
                    "300000000000000000000"
                ],
                [
                    "gard14wgcav3k99yz309vn7j6n3m50j32vkg426ktt0",
                    "100000000000000000000",
                    "200000000000000000000",
                    "300000000000000000000"
                ],
                [
                    "gard1hncel873ermm9e9009sthrys7ttdv6mtudfluz",
                    "100000000000000000000",
                    "200000000000000000000",
                    "300000000000000000000"
                ]
            ],
            "distributed":null
        }
    }
]


```
