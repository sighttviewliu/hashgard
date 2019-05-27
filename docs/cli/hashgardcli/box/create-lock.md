# hashgardcli box create-lock



## Description

Create a new lock box



## Usage

```shell
hashgardcli box create-lock [name] [total-amount] [end-time] --from
```
### Subcommands

| Name         | Type   | Required | Default | Description          |
| ------------ | ------ | -------- | ------- | -------------------- |
| name         | string | true     |         | Name of the lock box     |
| total-amount | string | true     |         | Lock the coin Type and quantity of the box|
| end-time     | int    | true     |         | Lock expiration time |



## Flags

**Global flags, query command flags** [hashgardcli](../README.md)



## Example

### Create a locked box
```shell
hashgardcli box create-lock ff 1000coin174876e800 1558066440 --from
```
Enter password to return
```txt
  {Height: 1936
  TxHash: B32D14F7F9D208733EB522CA80B4AB1CA6667271862DE2182E8501CF645E763D
  Data: 0F0E626F786161336A6C787074327074
  Raw Log: [{"msg_index":"0","success":true,"log":""}]
  Logs: [{"msg_index":0,"success":true,"log":""}]
  GasWanted: 200000
  GasUsed: 70033
  Tags:
    - action = box_create
    - category = box
    - box-id = boxaa3jlxpt2pt
    - box-Type = lock
    - sender = gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
    }
```

Check our account

```shell
hashgardcli bank account gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
```

The result is as follows：

```txt
{
  Account:
  Address:       gard1f76ncl7d9aeq2thj98pyveee8twplfqy3q4yv7
  Pubkey:        gardpub1addwnpepqfpd8mkl3jg43fw7y02fe99cgaxutf5npv9y9gx9dvrrcdwl36shv694apw
  Coins:         1000000000000000000000boxaa3jlxpt2pt,9999999907005070apple(coin174876e800)
  AccountNumber: 0
  Sequence:      7
}
```



### Available Commands

| Name                   | Description        |
| ------------------------- | ---------------------- |
| [query-box](query-box.md) | Query the specified box |
| [list-box](list-box.md)  | Query box list |
