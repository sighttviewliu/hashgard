# Changelog

## 0.7.0

### BREAKING CHANGES

+ Cosmos SDK
    + Upgrade cosmos-sdk from `v0.35.0` to `v0.35.5`

### FEATURES

+ Hashgard REST API (hashgardlcd)
    + [record] Add REST API for recording module
    + [account]  Add REST API for account-must-memo module
+ Hashgard CLI (hashgardcli)
    + [record] Add commands for recording module
    * [keys] Add commands for account-must-memo module

## 0.6.0

### BREAKING CHANGES

+ Tendermint
    + Upgrade tendermint from `v0.31.4` to `v0.31.5`

+ Cosmos SDK
    + Upgrade cosmos-sdk from `v0.34.4` to `v0.35.0`

### FEATURES

+ Hashgard REST API (hashgardlcd)
    + [lock] Add REST API for lock module
    + [issue]  Add REST API for deposit(HRC10) module
    + [deposit]  Add REST API for deposit(HRC12) module
    + [future]  Add REST API for future(HRC13) module
    * [issue] Refactor REST API for issue module
    * [exchange] Refactor REST API for exchange(HRC11) module
+ Hashgard CLI (hashgardcli)
    + [lock] Add commands for lock module
    + [deposit] Add commands for deposit(HRC12) module
    + [future] Add commands for future(HRC13) module
    - [box] Remove commands for box module
    + [issue] Refactor commands for issue(HRC10) module
    * [exchange] Refactor commands for exchange(HRC11)  module

+ Hashgard APP
    + [gov] Implementation  parameter change governance

### IMPROVEMENTS
+ Hashgard APP
    + [fee] Add new fee collection to support lock,deposit,future and issue module
      
### BUG FIXES

* Fix for the `x/staking` module security advisory for downstream consumers
* Fix gas consumption bug in `Undelegate` preventing the ability to sync from genesis.
* Unbonding from a validator is now only considered "complete" after the full
unbonding period has elapsed regardless of the validator's status.
    
## 0.5.0

### BREAKING CHANGES

+ Hashgard REST API (hashgardlcd)
    + Remove REST server's secure mode altogether

+ Hashgard Server (hashgard)
    + Validate genesis before running gentx

+ Hashgard CLI (hashgardcli)
    + Disable Keybase for Generate Only in CLI
    + Get fromAddress in BaseReq

+ Hashgard APP
    + Default don't check invarants blockly
    + Add vendor dir hash to version detail
    + Remove Shares Concept from Unbond/Redelegate
    + Refactor the tags structure
    + Modify the HD drive path from `44'/118'/` to `44'/322'/`

+ Tendermint
    + Upgrade tendermint from `v0.31.0-dev0-fix0` to `v0.31.4`

+ Cosmos SDK
    + Upgrade cosmos-sdk from `v0.33.2` to `v0.34.4`

### FEATURES

+ Hashgard REST API (hashgardlcd)
    + [mint] Add REST API for querying Mint/Inflation info
    + [distribution] Add new community-pool REST API
    + Add route for querying signing_info for all validators
    + [exchange] Add REST API for exchange module

+ Hashgard CLI (hashgardcli)
    + [mint] Add commands for querying Mint/Inflation info
    + [distribution] Add new community-pool command
    + Support native coin(gard) units in tx commands

+ Hashgard APP
      + [crisis] Add new module to do invarants check
      + [box] Add new module to support lock, share function

### IMPROVEMENTS

+ Hashgard REST API (hashgardlcd)
    + Update lcd swagger ui for /staking/validators query

+ Hashgard Server (hashgard)
    + Add optional flags for gentx command

+ Test
   + Add simulation test
   + Add cli test

### BUG FIXES

+ Hashgard REST API (hashgardlcd)
    + fix the incorrect field when marshal take order request


## 0.4.0

### BREAKING CHANGES

+ Hashgard REST API (hashgardlcd)
    + [base_req] Remove `generate_only` and `password` field, all functional APIs only support generate tx, cannot use keybase to sign tx any more.
    + [ICS1] Remove APIs of keys module.
    + [ICS20] Do not support `/tx/sign`, `/tx/broadcast` any more. Now use POST `/txs` (ICS0)to broadcast tx. Move encoding endpoint to `/txs/encode`.
    + [ICS24] Remove `/distribution/outstanding_rewards`, add new GET ``/distribution/validators/{validatorAddr}/outstanding_rewards` API to query the outstanding rewards of a specific validator.

+ Hashgard (hashgard)
    + Update validator creation flow:
        - Remove `NewMsgCreateValidatorOnBehalfOf` and corresponding business logic
        - Ensure the validator address equals the delegator address during `MsgCreateValidator#ValidateBasic`

+ Tendermint
    + Upgrade tendermint from `v0.30.0` to `v0.31.0-dev0-fix0`

+ Cosmos SDK
    + Upgrade cosmos-sdk from `v0.32.0` to `v0.33.2`

+ Dependency Management Tool
    + Now use `go mod` (go 1.11+) instead of `glide`

### FEATURES

+ Hashgard CLI (hashgardcli)
    + [faucet] Integrade `faucet send` function in CLI

+ Hashgard (hashgard)
    + [issue] Add new module to issue and manage token asserts.
    + [exchange] Add new module to support atomic-exchange of asserts.

### IMPROVEMENTS

+ Hashgard REST API (hashgardlcd)
    + Update the `TxResponse` type allowing for the `Logs` result to be JSON decoded automatically.

+ Hashgard CLI (hashgardcli)
    + Prompt user confirmation prior to signing and broadcasting a transaction.
    + Update `tx sign` to use `--from` instead of the deprecated `--name` CLI flag.
    + Querying account related information using custom querier in auth module.

### BUG FIXES

+ Hashgard CLI (hashgardcli)
    + `keys add --interactive` bip32 passphrase regression fix.

+ Hashgard (hashgard)
    + Fix distribution delegation for zero height export bug.
    + Properly return errors from a couple of struct Unmarshal functions.


## 0.3.1

### BREAKING CHANGES

+ Hashgard REST API (hashgardlcd)
    + Error responses are now JSON objects
    + [distribution] endpoint changed "all_delegation_rewards" -> "delegator_total_rewards"
    + `hashgardlcd` switched back to insecure mode by default
    + use `--tls` flag to enable secure layer
    + `GET /tx/{hash}` now returns `404` instead of `500` if the transaction is not found

+ Hashgard CLI (hashgardcli)
    + Add `hashgard validate-genesis` command to facilitate checking of genesis files
    + `version` prints out short info by default. Add `--long` flag. Proper handling of `--format` flag introduced.
    + now returns transactions in plain text including tags
    + Change validator address Bech32 encoding to consensus address in `tendermint-validator-set`.

+ Hashgard (hashgard)
    + Added Validator Minimum Self Delegation

+ Tendermint
    + upgrade tendermint from v0.29.0 to v0.30.0

+ Cosmos SDK
    + upgrade cosmos-sdk from v0.30.0 to v0.32.0

### FEATURES

+ Hashgard REST API (hashgardlcd)
    + Add distribution module REST API in LCD

+ Hashgard CLI (hashgardcli)
    + Support querying for all delegator distribution rewards.

+ Hashgard (hashgard)
    + Add support vesting accounts to the `add-genesis-account` command.

### IMPROVEMENTS

+ Hashgard REST API (hashgardlcd)
    + REST service to support the following:
        + Automatic account number and sequence population when fields are omitted
        + Generate only functionality no longer requires access to a local Keybase
        + `from` field in the `base_req` body can be a Keybase name or account address
    + Allow simulation(auto gas) to work with generate only.
    + Added `/tx/encode` endpoint to serialize a JSON tx to base64-encoded Amino

+ Hashgard CLI (hashgardcli)
    + Add new `withdraw-all-rewards` command to withdraw all delegations rewards for delegators.
    + `hashgard gentx` supports --ip and --node-id flags to override defaults.
    + Add `bank encode` command to serialize a JSON tx to base64-encoded Amino.

+ Hashgard (hashgard)
    + Add `--jail-whitelist` to `hashgard export` to enable testing of complex exports

### BUG FIXES

+ Hashgard REST API (hashgardlcd)
    + LCD didn't respect persistent flags such as `--chain-id` and `--trust-node` if they were passed on the command line.

+ Hashgard CLI (hashgardcli)
    + Fix `slashing signing-info` panic by ensuring safety of user input and properly returning not found error
    + Fix `distribution slashes` panic
    +

+ Hashgard (hashgard)
    + Return an empty `TxResponse` when Tendermint returns an empty `ResultBroadcastTx`.
--------------------------

## 0.3.0

### BREAKING CHANGES

+ Hashgard REST API (hashgardlcd)
	+ now default mode is insecure, use `--tls` flag to use https.
	+ `tx/sign` endpoint now expects `BaseReq` fields as nested object.
	+ all endpoints renamed from `/stake` -> `/staking`
	+ rename:
		+ `LooseTokens` -> `NotBondedTokens`
		+ `Validator.UnbondingMinTime` -> `Validator.UnbondingCompletionTime`
		+ `Delegation` -> `Value` in `MsgCreateValidator` and `MsgDelegate`
		+ `MsgBeginUnbonding` -> `MsgUndelegate`

+ Hashgard CLI (hashgardcli)
	+ Rename chain_id and trust_node to chain-id and trust-node respectively in config file.
	+ Remove unimplemented `init` command.
	+ `version` command output extra latest commit and build machine info.

+ Hashgard (hashgard)
	+ use Storekeys of each module instead of literals.
	+ the `--gas` flag now takes auto instead of simulate .
	+ `version` command output extra latest commit and build machine info.
	+ `tendermint`'s show-validator and `show-address` `--json` flags removed in favor of `--output-format=json`.

+ Tendermint
	+ upgrade tendermint from v0.27.3 to v0.29.0

+ Cosmos SDK
	+ upgrade cosmos-sdk from v0.29.0 to v0.30.0
	+ rename module `stake` -> `staking`.
	+ rename `LooseTokens` -> `NotBondedTokens`
	+ [staking] Validator power type from `Dec` -> `Int`.
	+ [gov] Remove redundant action tag
	+ Sanitize sdk.Coin denom. Coins denoms are now case insensitive, i.e. 100fooToken equals to 100FOOTOKEN.
	+ Fix infinite gas meter utilization during aborted ante handler executions.
	+ Increase decimal precision to 18


### FEATURES

+ Hashgard REST API (hashgardlcd)
	+ add support for fees on transactions.
	+ add a custom memo on transactions.
	+ implement `/gov/proposals/{proposalID}/proposer` to query for a proposal's proposer.

+ Hashgard CLI (hashgardcli)
	+ new `keys add --multisig` flag to store multisig keys locally.
	+ new `bank sign --multisig` flag to enable multisig mode.
	+ add new `bank multisign` command to generate multisig signatures for transactions generated offline
	+ add `distribution params`, `distribution outstanding-rewards`, `distribution commission`, `distribution slashes`, `distributionrewards` commands to query more info.
	+ add new `slashing params` command to query the current slashing parameters.
	+ add new `gov param`, `gov proposer` commands to query more relative info.

+ Hashgard (hashgard)
	+ added queriers for querying a single redelegation
	+ queriers for all distribution state worth querying
	+ Add support for vesting accounts at genesis.
	+ Add multisig transactions support
	+ `add-genesis-account` can take both account addresses and key names

+ Cosmos SDK
	+ Vesting account implementation.
	+ show bech32-ify accounts address in error message


### IMPROVEMENTS

+ Hashgard REST API (hashgardlcd)
	+ Validate tx/sign endpoint POST body.

+ Hashgard CLI (hashgardcli)
	+ Support adding offline public keys to the keystore

+ Hashgard (hashgard)
	+ add validation for slashing genesis
	+ support minimum fees in a local testnet
	+ Refactor tx fee:
		+ Validators specify minimum gas prices instead of minimum fees
		+ Clients may provide either fees or gas prices directly
		+ The gas prices of a tx must meet a validator's minimum
		+ `hashgard start` and config file take `--minimum-gas-prices` flag and `minimum-gas-price` field respectively.
	+ now `hashgard gentx` command printout of account's addresses, i.e. user bech32 instead of hex.

+ Cosmos SDK
	+ Add tag documentation for each module along with cleaning up a few existing tags in the governance, slashing, and staking modules.


### BUG FIXES

+ Hashgard CLI (hashgardcli)
	+ Fix the bug in GetAccount when `len(res) == 0` or `err == nil`

+ Hashgard (hashgard)
	+ Correctly reset total accum update height and jailed-validator bond height / unbonding height on export-for-zero-height
	+ Fix unset governance proposal queues when importing state from old chain
	+ Fix `hashgard export` by resetting each validator's slashing period.


--------------------------

## 0.2.1

BREAKING CHANGES
* Cosmos-sdk
  * [cosmos-sdk] Now using cosmos-sdk 0.29.1
* Change the name of stake coin


--------------------------

## 0.2.0

BREAKING CHANGES
* Cosmos-sdk
  * [cosmos-sdk] Now using cosmos-sdk 0.29.0
* Tendermint
  * [tendermint] Now using tendermint 0.27.3

FEATURES
* Hashgard REST API(hashgardlcd)
  * [hashgardlcd] Split the LCD service from hashgardcli
* Other
  * [logjack] Introduced the logjack tool for saving logs w/ rotation

BUG FIXES
* Fixed the bug that did not recognize the msg of the distribution command
