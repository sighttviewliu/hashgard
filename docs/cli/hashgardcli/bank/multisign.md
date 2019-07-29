# hashgardcli bank multisign

## Description

Multisig transactions require signatures of multiple private keys. Thus, generating and signing a transaction from a multisig account involve cooperation among the parties involved. A multisig transaction can be initiated by any of the key holders, and at least one of them would need to import other parties' public keys into their Keybase and generate a multisig public key in order to finalize and broadcast the transaction.

## Usage

```shell
hashgardcli bank multisign [file] [name] [[signature]...] [flags]
```

## Flags

| Name   | Type   | Required | Default | Description        |
| ----------------- | ------ | -------- | ------- | --------------- |
| --offline         | bool   | false    | false   | Off-chain mode, do not query the full node    |
| --output-document | string | false    | false   | The document will be written to the given file instead of STDOUT |
| --signature-only  | bool   | false    | false   | Print only the generated signature and then exit             |

**Global flags, query command flags** [hashgardcli](../README.md)



## Example

given a multisig key comprising the keys a1, a2, and a3, each of which is held by a distinct party, the user holding a1 would require to import both a2 and a3 in order to generate the multisig account public key:

```shell
hashgardcli keys add a1

hashgardcli keys add a2

hashgardcli keys add a3

hashgardcli keys add a123 \
    --multisig=a1,a2,a3 \
    --multisig-threshold=2
```

A new multisig public key a123 has been stored, and its address will be used as signer of multisig transactions:

```shell
hashgardcli keys show --address a123
```

```shell
NAME:   Type:   ADDRESS:                                                PUBKEY:
a123  offline gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n     gardpub1ytql0csgqgfzd666axrjzq7lfft2evw9r7j0u3t7yj4qjy5rczhncv8ysykrp35cpjpklsj5rcfzd666axrjzquew3ad0vgywr7gmgszly9wnw2mwxc3k7dttlmm780g5y9djw8vcgfzd666axrjzq63kk98gyurzz2rewxxhd4dxvvdfsnsdtegajrcez3exg3yu9q0a5kpkkj3
```

The first step to create a multisig transaction is to initiate it on behalf of the multisig address created above:

```shell
hashgardcli bank send gard19thul47y2afwr67l4hlv9hu5593uw0rqhashgjdm7jj 10gard \
    --from a123 \
    --generate-only >unsignedTx.json
```

The file unsignedTx.json contains the unsigned transaction encoded in JSON. a1 can now sign the transaction with its own private key:

```shell
hashgardcli bank sign unsignedTx.json \
    --multisig=gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n \
    --from a1 \
    --output-document=a1sign.json
```

Once the signature is generated, a1 transmits both unsignedTx.json and a1sign.json to a2 or a3, which in turn will generate their respective signature:

```shell
hashgardcli bank sign unsignedTx.json \
    --multisig=gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n \
    --from a2 \
    --output-document=a2sign.json
```

a123 is a 2-of-3 multisig key, therefore one additional signature is sufficient. Any the key holders can now generate the multisig transaction by combining the required signature files:

```shell
hashgardcli bank multisign  unsignedTx.json a123 a1sign.json a2sign.json \
    --output-document=signedTx.json
```

The transaction can now be sent to the node:

```bash
hashgardcli tx broadcast signedTx.json
```

Tx query:

```shell
hashgardcli tendermint tx 6A66C370834F097CA36F60FE9B4E8ABEEEF3549D089071FDB5EE33277B615035
```

```json
{
 "height": "63108",
 "txhash": "6A66C370834F097CA36F60FE9B4E8ABEEEF3549D089071FDB5EE33277B615035",
 "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
 "gas_wanted": "200000",
 "gas_used": "31450",
 "tags": [
  {
   "key": "action",
   "value": "send"
  },
  {
   "key": "sender",
   "value": "gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n"
  },
  {
   "key": "recipient",
   "value": "gard19thul47y2afwr67l4hlv9hu5593uw0rqjdm7jj"
  }
 ],
 "tx": {
  "type": "auth/StdTx",
  "value": {
   "msg": [
    {
     "type": "cosmos-sdk/Send",
     "value": {
      "from_address": "gard15l5yzrq3ff8fl358ng430cc32lzkvxc30n405n",
      "to_address": "gard19thul47y2afwr67l4hlv9hu5593uw0rqjdm7jj",
      "amount": [
       {
        "denom": "gard",
        "amount": "10"
       }
      ]
     }
    }
   ],
   "fee": {
    "amount": null,
    "gas": "200000"
   },
   "signatures": [
    {
     "pub_key": {// multi-sign transaction
      "type": "tendermint/PubKeyMultisigThreshold",
      "value": {
       "threshold": "2", // two signatures
       "pubkeys": [
        {
         "type": "tendermint/PubKeySecp256k1",
         "value": "A99KVqyxxR+k/kV+JKoJEoPArzww5IEsMMaYDINvwlQe"
        },
        {
         "type": "tendermint/PubKeySecp256k1",
         "value": "A5l0etexBHD8jaIC+QrpuVtxsRt5q1/3vx3ooQrZOOzC"
        },
        {
         "type": "tendermint/PubKeySecp256k1",
         "value": "A1G1inQTgxCUPLjGu2rTMY1MJwavKOyHjIo5MiJOFA/t"
        }
       ]
      }
     },
     "signature": "CgUIAxIBYBJAElZbW6piLDmd+8mG1VLPVYuQK9r/5fitsXvDONtiarVPFSzqf8DkbsyPBOCQOmfuMkhFt+S1TqyFyUZuaE242hJA2j2QTmtW8eEtqOPAkyed0j/97q9phg34KV95gvfp0wd7V0umKoyj/FX/WTvD7iYNWS2ssbwjpztItggOcCTeCw=="
    }
   ],
   "memo": ""
  }
 }
}

```
