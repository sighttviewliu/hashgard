# hashgardcli box create-deposit

## Description

Create a new deposit box


## Usage

```
 hashgardcli box create-deposit [name][total-amount][flags] --from
```

### Subcommands


| Name         | Type  | Required  | Default| Description            |
| ------------ | ------ | -------- | ------ | -------------------- |
| name         | string | true      |        | depositbox name     |
| total-amount | string | true      |        | Total amount and coin Type of deposit accepted |

### Flags

| Name            | Type  | Required  | Default| Description                      |
| ---------------- | ------ | -------- | ------ | ------------------------------ |
| --bottom-line    | Int    | true     | ""     | depositBox bottom line        |
| --price          | int    | true     | ""     |  depositBox unit price |
| --start-time     | int    | true      | ""     | depositBox start time               |
| --establish-time | int    | true     | ""     | Box establish time              |
| --maturity-time  | int    | true     | ""     | Box maturity time                   |
| --interest       | string | true     | ""     | Add indent to JSON response           |

**Global flags, query command flags** [hashgardcli](../README.md)

## Example
### Great deposit box
```shell
hashgardcli box create-deposit mingone 10000coin174876e800  --bottom-line=0 --price=2  --start-time=1558079700  --establish-time=1558080300 --maturity-time=1558080900 --interest=9898coin174876e800  --from
```
After the password is confirmed，The result is as follows：
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



### Available Commands

| Name                                  | Description                    |
| ------------------------------------------- | ---------------------------- |
| [interest-injection](interest-injection.md) | Inject interest into the box |
| [interest-fetch](interest-fetch.md)         | Withdrawal interest on the box |
| [deposit-to](deposit-to.md)                 | Deposit the box |
| [deposit-fetch](deposit-fetch.md)           | Withdrawal of the box |
| [query-box](query-box.md)                   | Query box information   |
| [list-box](list-box.md)                     | Query box list       |
