# hashgardcli bank sign

## Description

Sign transactions created with the --generate-only flag

## Usage

```
hashgardcli bank sign [file] [flags]
```

## Flags

| Name   | Type   | Required  | Default        | Description                  |
| ---------------- | ------- | -------- | --------------------- | ------------------------------------------------------------ |
| --append | bool | false| true | Append the signature to the existing ones. If disabled, old signatures would be overwritten. Ignored if --multisig is on |
| --multisig | string | false| |  Address of the multisig account on behalf of which the transaction shall be signed |
| --from | string | false| |  Name or address of private key with which to sign|
| --offline | bool | false| false |  Offline mode; Do not query a full node|
| --output-document | string |  |  | The document will be written to the given file instead of STDOUT |
| --signature-only | bool | false| | Print only the generated signature, then exit|
| --validate-signatures | bool | false| false |  Print the addresses that must sign the transaction, those who have already signed it, and make sure that signatures are in the correct order|


**Global flags, query command flags** [hashgardcli](../README.md)

## Example

### Sign a send file

First you must **hashgardcli bank send** command with flag **--generate-only** to generate a send recorder. Just like this.

```  
hashgardcli bank send gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx 10gard --from=test --chain-id=hashgard --generate-only

{"type":"auth/StdTx","value":{"msg":[{"type":"cosmos-sdk/Send","value":{"inputs":[{"address":"gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx","coins":[{"denom":"gard","amount":"10000000000000000000"}]}],"outputs":[{"address":"gard9aamjx3xszzxgqhrh0yqd4hkurkea7f6d429yx","coins":[{"denom":"gard","amount":"10000000000000000000"}]}]}}],"fee":{"amount":[{"denom":"gard","amount":"4000000000000000"}],"gas":"200000"},"signatures":null,"memo":""}}
```

And then save the output in /root/node0/test_send_10hashgard.txt

Then you can sign the offline file.

```
hashgardcli bank sign /root/node0/test_send_10hashgard.txt --from=test  --offline=false --print-response --append=true
```

After that, you will get the detail info for the sign. Like the follow output you will see the signature:
**ci+5QuYUVcsARBQWyPGDgmTKYu/SRj6TpCGvrC7AE3REMVdqFGFK3hzlgIphzOocGmOIa/wicXGlMK2G89tPJg==**
