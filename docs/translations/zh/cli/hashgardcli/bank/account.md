# hashgardcli bank account

## 描述

查询账户信息

## 用法

```
hashgardcli bank account [address] [flags]
```

## Flags

 **全局 flags、查询命令 flags** 参考：[hashgardcli](../README.md)

## 例子

### 查询账户信息

```
 hashgardcli bank account gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx -o json --trust-node --indent
```

执行完命令后，获得账户的详细信息如下：

```
{
 "type": "auth/Account",
 "value": {
  "address": "gard10tfnpxvxjh6tm6gxq978ssg4qlk7x6j9aeypzn",
  "coins": [
   {
    "denom": "gard",
    "amount": "1900000000"
   }
  ],
  "public_key": {
   "type": "tendermint/PubKeySecp256k1",
   "value": "AvM1uBBEml3ZtXP5GZD6vr7UIcit6GMjS0ZUdxuejShH"
  },
  "account_number": "0",
  "sequence": "1"
 }
}
```
如果你查询一个错误的地址，将会返回如下信息:
```
hashgardcli bank account gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429zz
ERROR: decoding bech32 failed: checksum failed. Expected d429yx, got d429zz.
```
如果查询一个空地址，将会返回如下信息:
```
hashgardcli bank account gardkenrwk5k4ng70e5s9zfsttxpnlesx5ps0gfdv7
ERROR: No account with address gardkenrwk5k4ng70e5s9zfsttxpnlesx5ps0gfdv7 was found in the state.
Are you sure there has been a transaction involving it?
```
