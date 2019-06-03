# hashgardcli box list-box

## Description
Query box list

## Usage
```shell
hashgardcli box list-box [flag]
```

### flag

| Name    | Description   |
| ------- | -------- |
| lock    | lock box |
| deposit | deposit box|
| future  | future box  |


**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### Query box list

```shell
hashgardcli box list-box future
```

The result is as follows：

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
        "description":"{"org":"Hashgard","website":"https://www.hashgard.com","logo":"https://cdn.hashgard.com/static/logo.2d949f3d.png","intro":"This is a description of the project"}",
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
