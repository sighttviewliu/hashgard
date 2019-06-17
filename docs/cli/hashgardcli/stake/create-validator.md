# hashgardcli stake create-validator

## Description

create new validator initialized with a self-delegation to it

## Usage

```shell
hashgardcli stake create-validator [flags]
```

## Flags

| Name                         | Type  | Required| Default| Description      |
| ---------------------------- | ------ | -------- | ------ | --------------------------- |
| --address-delegator          | string | false    |    | delegator address                          |
| --amount                     | string | true     |    | Amount of coins to bond                    |
| --commission-max-change-rate | float  | false    | 0.0    | The maximum commission change rate percentage (per day)|
| --commission-max-rate        | float  | false    | 0.0    | The maximum commission rate percentage        |
| --commission-rate            | float  | false    | 0.0    | The initial commission rate percentage     |
| --details                    | string | false    |    | TExport the transaction in gen-tx format; it implies --generate-only |
| --identity                   | string | false    |    | Optional identity signature (ex. UPort or Keybase)|
| --ip                         | string | false    |    | The node's public IP. It takes effect only when used in combination with --generate-only|
| --min-self-delegation        | string | true     |    | The minimum self delegation required on the validator（Unit agard）|
| --moniker                    | string | true     |    | The validator's name              |
| --node-id                    | string | false    |    | The node's ID                      |
| --pubkey                     | string | true     |    | The Bech32 encoded PubKey of the validator   |
| --website                    | string | false    |    | The validator's (optional) website   |

**Global flags, query command flags** [hashgardcli](../README.md)

## Example

```shell
hashgardcli stake create-validator \
--from=gard1rkqp5w062sdqu68qsufn3safwz0e5m9f3efaak \
--pubkey=gardvalconspub1zcjduepqm58zvp4fmssekze8m04ppyf55uwfhecnpe5lfs3znxtes2mhz8esrvvtqv \
--commission-max-change-rate=0.01 \
--commission-rate=0.1 \
--commission-max-rate=0.2 \
--amount=40gard \
--moniker=joelove \
--min-self-delegation=10 \
--chain-id=chain-HwMeL0 \
--fees=2gard \
--output=json \
--indent
```

The result is as follows：

```json
{
    "height": "19195",
    "txhash": "0DCBF4DB2F64A625FD13166D323A5033D15D4CCFB07217F08E0CFDFF8FC29998",
    "log": "[{\"msg_index\":\"0\",\"success\":true,\"log\":\"\"}]",
    "gas_wanted": "200000",
    "gas_used": "99976",
    "tags": [
        {
            "key": "action",
            "value": "create_validator"
        },
        {
            "key": "destination-validator",
            "value": "gardvaloper1rkqp5w062sdqu68qsufn3safwz0e5m9frmy4dm"
        },
        {
            "key": "moniker",
            "value": "joelove"
        },
        {
            "key": "identity"
        }
    ]
}
```
