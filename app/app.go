package app

import (
	"fmt"
	"io"
	"os"
	"sort"

	bam "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	abci "github.com/tendermint/tendermint/abci/types"
	cmn "github.com/tendermint/tendermint/libs/common"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	acc "github.com/cosmos/cosmos-sdk/x/account"
	"github.com/hashgard/hashgard/x/box"
	"github.com/hashgard/hashgard/x/distribution"
	"github.com/hashgard/hashgard/x/exchange"
	"github.com/hashgard/hashgard/x/gov"
	"github.com/hashgard/hashgard/x/issue"
	"github.com/hashgard/hashgard/x/mint"
	"github.com/hashgard/hashgard/x/record"
)

const (
	appName = "HashgardApp"
	// DefaultKeyPass contains key password for genesis transactions
	DefaultKeyPass = "12345678"
)

// default home directories for expected binaries
var (
	DefaultLCDHome  = os.ExpandEnv("$HOME/.hashgardlcd")
	DefaultNodeHome = os.ExpandEnv("$HOME/.hashgard")
	DefaultCLIHome  = os.ExpandEnv("$HOME/.hashgardcli")
)

// Extended ABCI application
type HashgardApp struct {
	*bam.BaseApp
	cdc *codec.Codec

	invCheckPeriod uint

	// keys to access the multistore
	keyMain          *sdk.KVStoreKey
	keyAccMustMemo   *sdk.KVStoreKey
	keyAccount       *sdk.KVStoreKey
	keyStaking       *sdk.KVStoreKey
	tkeyStaking      *sdk.TransientStoreKey
	keySlashing      *sdk.KVStoreKey
	keyMint          *sdk.KVStoreKey
	keyDistribution  *sdk.KVStoreKey
	tkeyDistribution *sdk.TransientStoreKey
	keyGov           *sdk.KVStoreKey
	keyIssue         *sdk.KVStoreKey
	keyRecord        *sdk.KVStoreKey
	keyBox           *sdk.KVStoreKey
	keyFeeCollection *sdk.KVStoreKey
	keyExchange      *sdk.KVStoreKey
	keyParams        *sdk.KVStoreKey
	tkeyParams       *sdk.TransientStoreKey

	// manage getting and setting accounts
	accMustMemoKeeper   acc.Keeper
	accountKeeper       auth.AccountKeeper
	feeCollectionKeeper auth.FeeCollectionKeeper
	bankKeeper          bank.Keeper
	stakingKeeper       staking.Keeper
	slashingKeeper      slashing.Keeper
	mintKeeper          mint.Keeper
	distributionKeeper  distribution.Keeper
	govKeeper           gov.Keeper
	exchangeKeeper      exchange.Keeper
	crisisKeeper        crisis.Keeper
	paramsKeeper        params.Keeper
	issueKeeper         issue.Keeper
	recordKeeper        record.Keeper
	boxKeeper           box.Keeper
}

// NewHashgardApp returns a reference to an initialized HashgardApp.
func NewHashgardApp(logger log.Logger, db dbm.DB, traceStore io.Writer,
	loadLatest bool,
	invCheckPeriod uint,
	baseAppOptions ...func(*bam.BaseApp)) *HashgardApp {

	cdc := MakeCodec()

	bApp := bam.NewBaseApp(appName, logger, db, auth.DefaultTxDecoder(cdc), baseAppOptions...)
	bApp.SetCommitMultiStoreTracer(traceStore)

	// create your application type
	var app = &HashgardApp{
		BaseApp:          bApp,
		cdc:              cdc,
		invCheckPeriod:   invCheckPeriod,
		keyMain:          sdk.NewKVStoreKey(bam.MainStoreKey),
		keyAccMustMemo:   sdk.NewKVStoreKey(acc.StoreKey),
		keyAccount:       sdk.NewKVStoreKey(auth.StoreKey),
		keyStaking:       sdk.NewKVStoreKey(staking.StoreKey),
		tkeyStaking:      sdk.NewTransientStoreKey(staking.TStoreKey),
		keyMint:          sdk.NewKVStoreKey(mint.StoreKey),
		keyDistribution:  sdk.NewKVStoreKey(distribution.StoreKey),
		tkeyDistribution: sdk.NewTransientStoreKey(distribution.TStoreKey),
		keySlashing:      sdk.NewKVStoreKey(slashing.StoreKey),
		keyGov:           sdk.NewKVStoreKey(gov.StoreKey),
		keyIssue:         sdk.NewKVStoreKey(issue.StoreKey),
		keyRecord:        sdk.NewKVStoreKey(record.StoreKey),
		keyBox:           sdk.NewKVStoreKey(box.StoreKey),
		keyFeeCollection: sdk.NewKVStoreKey(auth.FeeStoreKey),
		keyExchange:      sdk.NewKVStoreKey(exchange.StoreKey),
		keyParams:        sdk.NewKVStoreKey(params.StoreKey),
		tkeyParams:       sdk.NewTransientStoreKey(params.TStoreKey),
	}

	app.paramsKeeper = params.NewKeeper(
		app.cdc,
		app.keyParams,
		app.tkeyParams,
	)

	app.accMustMemoKeeper = acc.NewKeeper(
		app.cdc,
		app.keyAccMustMemo,
		app.paramsKeeper.Subspace(acc.DefaultParamspace),
		acc.DefaultCodespace,
	)
	// define the accountKeeper
	app.accountKeeper = auth.NewAccountKeeper(
		app.cdc,
		app.keyAccount, // target store
		app.paramsKeeper.Subspace(auth.DefaultParamspace),
		auth.ProtoBaseAccount, // prototype
	)

	app.feeCollectionKeeper = auth.NewFeeCollectionKeeper(
		app.cdc,
		app.keyFeeCollection,
	)

	bankKeeper := bank.NewBaseKeeper(
		app.accountKeeper,
		app.paramsKeeper.Subspace(bank.DefaultParamspace),
		bank.DefaultCodespace,
	)

	stakingKeeper := staking.NewKeeper(
		app.cdc,
		app.keyStaking,
		app.tkeyStaking,
		&bankKeeper,
		app.paramsKeeper.Subspace(staking.DefaultParamspace),
		staking.DefaultCodespace,
	)

	app.mintKeeper = mint.NewKeeper(
		app.cdc,
		app.keyMint,
		app.paramsKeeper.Subspace(mint.DefaultParamspace),
		&stakingKeeper,
		app.feeCollectionKeeper,
		&bankKeeper,
	)

	app.distributionKeeper = distribution.NewKeeper(
		app.cdc,
		app.keyDistribution,
		app.paramsKeeper.Subspace(distribution.DefaultParamspace),
		&bankKeeper,
		&stakingKeeper,
		app.feeCollectionKeeper,
		distribution.DefaultCodespace,
	)

	app.slashingKeeper = slashing.NewKeeper(
		app.cdc,
		app.keySlashing,
		&stakingKeeper,
		app.paramsKeeper.Subspace(slashing.DefaultParamspace),
		slashing.DefaultCodespace,
	)

	app.issueKeeper = issue.NewKeeper(
		app.cdc,
		app.keyIssue,
		app.paramsKeeper,
		app.paramsKeeper.Subspace(issue.DefaultParamspace),
		&bankKeeper,
		app.feeCollectionKeeper,
		issue.DefaultCodespace)

	app.recordKeeper = record.NewKeeper(
		app.cdc,
		app.keyRecord,
		app.paramsKeeper,
		app.paramsKeeper.Subspace(record.DefaultParamspace),
		record.DefaultCodespace)

	app.boxKeeper = box.NewKeeper(
		app.cdc,
		app.keyBox,
		app.paramsKeeper,
		app.paramsKeeper.Subspace(box.DefaultParamspace),
		&bankKeeper,
		app.issueKeeper,
		app.feeCollectionKeeper,
		box.DefaultCodespace)

	app.exchangeKeeper = exchange.NewKeeper(
		app.cdc,
		app.keyExchange,
		app.paramsKeeper,
		app.paramsKeeper.Subspace(exchange.DefaultParamspace),
		&bankKeeper,
		exchange.DefaultCodespace,
	)

	app.crisisKeeper = crisis.NewKeeper(
		app.paramsKeeper.Subspace(crisis.DefaultParamspace),
		app.distributionKeeper,
		&bankKeeper,
		app.feeCollectionKeeper,
	)

	app.govKeeper = gov.NewKeeper(
		app.cdc,
		app.keyGov,
		app.paramsKeeper,
		app.paramsKeeper.Subspace(gov.DefaultParamspace),
		&bankKeeper,
		&stakingKeeper,
		gov.DefaultCodespace,
		app.accountKeeper,
		app.distributionKeeper,
		app.mintKeeper,
		app.slashingKeeper,
		&stakingKeeper,
	)

	// register the staking hooks
	// NOTE: stakeKeeper above are passed by reference,
	// so that it can be modified like below:
	app.stakingKeeper = *stakingKeeper.SetHooks(
		NewStakingHooks(app.distributionKeeper.Hooks(), app.slashingKeeper.Hooks()),
	)

	app.bankKeeper = *bankKeeper.SetHooks(
		NewBankHooks(app.boxKeeper.Hooks(), app.issueKeeper.Hooks(), app.accMustMemoKeeper.Hooks()),
	)

	// register the crisis routes
	bank.RegisterInvariants(&app.crisisKeeper, app.accountKeeper)
	distribution.RegisterInvariants(&app.crisisKeeper, app.distributionKeeper, app.stakingKeeper)
	staking.RegisterInvariants(&app.crisisKeeper, app.stakingKeeper, app.feeCollectionKeeper, app.distributionKeeper, app.accountKeeper)

	// register message routes
	app.Router().
		AddRoute(acc.RouterKey, acc.NewHandler(app.accMustMemoKeeper)).
		AddRoute(bank.RouterKey, bank.NewHandler(app.bankKeeper)).
		AddRoute(staking.RouterKey, staking.NewHandler(app.stakingKeeper)).
		AddRoute(distribution.RouterKey, distribution.NewHandler(app.distributionKeeper)).
		AddRoute(slashing.RouterKey, slashing.NewHandler(app.slashingKeeper)).
		AddRoute(gov.RouterKey, gov.NewHandler(app.govKeeper)).
		AddRoute(exchange.RouterKey, exchange.NewHandler(app.exchangeKeeper)).
		AddRoute(issue.RouterKey, issue.NewHandler(app.issueKeeper)).
		AddRoute(record.RouterKey, record.NewHandler(app.recordKeeper)).
		AddRoute(box.RouterKey, box.NewHandler(app.boxKeeper)).
		AddRoute(crisis.RouterKey, crisis.NewHandler(app.crisisKeeper))

	app.QueryRouter().
		AddRoute(acc.QuerierRoute, acc.NewQuerier(app.accMustMemoKeeper)).
		AddRoute(auth.QuerierRoute, auth.NewQuerier(app.accountKeeper)).
		AddRoute(staking.QuerierRoute, staking.NewQuerier(app.stakingKeeper, app.cdc)).
		AddRoute(slashing.QuerierRoute, slashing.NewQuerier(app.slashingKeeper, app.cdc)).
		AddRoute(gov.QuerierRoute, gov.NewQuerier(app.govKeeper)).
		AddRoute(distribution.QuerierRoute, distribution.NewQuerier(app.distributionKeeper)).
		AddRoute(exchange.QuerierRoute, exchange.NewQuerier(app.exchangeKeeper, app.cdc)).
		AddRoute(issue.QuerierRoute, issue.NewQuerier(app.issueKeeper)).
		AddRoute(record.QuerierRoute, record.NewQuerier(app.recordKeeper)).
		AddRoute(box.QuerierRoute, box.NewQuerier(app.boxKeeper)).
		AddRoute(mint.QuerierRoute, mint.NewQuerier(app.mintKeeper))

	// initialize BaseApp
	app.MountStores(
		app.keyMain,
		app.keyAccMustMemo,
		app.keyAccount,
		app.keyStaking,
		app.keyMint,
		app.keyDistribution,
		app.keySlashing,
		app.keyGov,
		app.keyIssue,
		app.keyRecord,
		app.keyBox,
		app.keyFeeCollection,
		app.keyExchange,
		app.keyParams,
		app.tkeyParams,
		app.tkeyStaking,
		app.tkeyDistribution,
	)
	app.SetInitChainer(app.initChainer)
	app.SetBeginBlocker(app.BeginBlocker)
	app.SetAnteHandler(auth.NewAnteHandler(app.accountKeeper, app.feeCollectionKeeper))
	app.SetEndBlocker(app.EndBlocker)

	if loadLatest {
		err := app.LoadLatestVersion(app.keyMain)
		if err != nil {
			cmn.Exit(err.Error())
		}
	}

	return app
}

// MakeCodec creates a new codec and registers all the necessary types with the codec.
func MakeCodec() *codec.Codec {
	cdc := codec.New()

	acc.RegisterCodec(cdc)
	auth.RegisterCodec(cdc)
	bank.RegisterCodec(cdc)
	staking.RegisterCodec(cdc)
	distribution.RegisterCodec(cdc)
	slashing.RegisterCodec(cdc)
	gov.RegisterCodec(cdc)
	exchange.RegisterCodec(cdc)
	issue.RegisterCodec(cdc)
	record.RegisterCodec(cdc)
	box.RegisterCodec(cdc)
	crisis.RegisterCodec(cdc)
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)

	return cdc
}

// application updates every end block
func (app *HashgardApp) BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock {

	// mint new tokens for this new block
	mint.BeginBlocker(ctx, app.mintKeeper)

	// distribute rewards from previous block
	distribution.BeginBlocker(ctx, req, app.distributionKeeper)

	// slash anyone who double signed.
	// NOTE: This should happen after distr.BeginBlocker so that
	// there is nothing left over in the validator fee pool,
	// so as to keep the CanWithdrawInvariant invariant.
	// TODO: This should really happen at EndBlocker.
	tags := slashing.BeginBlocker(ctx, req, app.slashingKeeper)

	return abci.ResponseBeginBlock{
		Tags: tags.ToKVPairs(),
	}
}

// application updates every end block
// nolint: unparam
func (app *HashgardApp) EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {

	tags := gov.EndBlocker(ctx, app.govKeeper)
	validatorUpdates, endBlockerTags := staking.EndBlocker(ctx, app.stakingKeeper)
	tags = append(tags, endBlockerTags...)
	boxTags := box.EndBlocker(ctx, app.boxKeeper)
	tags = append(tags, boxTags...)

	if app.invCheckPeriod != 0 && ctx.BlockHeight()%int64(app.invCheckPeriod) == 0 {
		app.assertRuntimeInvariants()
	}

	return abci.ResponseEndBlock{
		ValidatorUpdates: validatorUpdates,
		Tags:             tags,
	}
}

// initialize store from a genesis state
func (app *HashgardApp) initFromGenesisState(ctx sdk.Context, genesisState GenesisState) []abci.ValidatorUpdate {
	genesisState.Sanitize()

	// load the accounts
	for _, gacc := range genesisState.Accounts {
		acc := gacc.ToAccount()
		acc = app.accountKeeper.NewAccount(ctx, acc) // set account number
		app.accountKeeper.SetAccount(ctx, acc)
	}

	// initialize distribution (must happen before staking)
	distribution.InitGenesis(ctx, app.distributionKeeper, genesisState.DistributionData)

	// load the initial staking information
	validators, err := staking.InitGenesis(ctx, app.stakingKeeper, genesisState.StakingData)
	if err != nil {
		panic(err) // TODO find a way to do this w/o panics
	}

	// initialize module-specific stores
	acc.InitGenesis(ctx, app.accMustMemoKeeper, genesisState.AccMustMemoData)
	auth.InitGenesis(ctx, app.accountKeeper, app.feeCollectionKeeper, genesisState.AuthData)
	bank.InitGenesis(ctx, app.bankKeeper, genesisState.BankData)
	slashing.InitGenesis(ctx, app.slashingKeeper, genesisState.SlashingData, genesisState.StakingData.Validators.ToSDKValidators())
	gov.InitGenesis(ctx, app.govKeeper, genesisState.GovData)
	mint.InitGenesis(ctx, app.mintKeeper, genesisState.MintData)
	issue.InitGenesis(ctx, app.issueKeeper, genesisState.IssueData)
	record.InitGenesis(ctx, app.recordKeeper, genesisState.RecordData)
	box.InitGenesis(ctx, app.boxKeeper, genesisState.BoxData)
	exchange.InitGenesis(ctx, app.exchangeKeeper, genesisState.ExchangeData)
	crisis.InitGenesis(ctx, app.crisisKeeper, genesisState.CrisisData)

	// validate genesis state
	if err := HashgardValidateGenesisState(genesisState); err != nil {
		panic(err) // TODO find a way to do this w/o panics
	}

	if len(genesisState.GenTxs) > 0 {
		for _, genTx := range genesisState.GenTxs {
			var tx auth.StdTx
			err = app.cdc.UnmarshalJSON(genTx, &tx)
			if err != nil {
				panic(err)
			}
			bz := app.cdc.MustMarshalBinaryLengthPrefixed(tx)
			res := app.BaseApp.DeliverTx(bz)
			if !res.IsOK() {
				panic(res.Log)
			}
		}

		validators = app.stakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	}

	return validators
}

// custom logic for hashgard initialization
func (app *HashgardApp) initChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
	stateJSON := req.AppStateBytes
	// TODO is this now the whole genesis file?

	var genesisState GenesisState
	err := app.cdc.UnmarshalJSON(stateJSON, &genesisState)
	if err != nil {
		panic(err) // TODO https://github.com/cosmos/cosmos-sdk/issues/468
		// return sdk.ErrGenesisParse("").TraceCause(err, "")
	}

	validators := app.initFromGenesisState(ctx, genesisState)

	// sanity check
	if len(req.Validators) > 0 {
		if len(req.Validators) != len(validators) {
			panic(fmt.Errorf("len(RequestInitChain.Validators) != len(validators) (%d != %d)",
				len(req.Validators), len(validators)))
		}
		sort.Sort(abci.ValidatorUpdates(req.Validators))
		sort.Sort(abci.ValidatorUpdates(validators))
		for i, val := range validators {
			if !val.Equal(req.Validators[i]) {
				panic(fmt.Errorf("validators[%d] != req.Validators[%d] ", i, i))
			}
		}
	}

	// assert runtime invariants
	app.assertRuntimeInvariants()

	return abci.ResponseInitChain{
		Validators: validators,
	}
}

// load a particular height
func (app *HashgardApp) LoadHeight(height int64) error {
	return app.LoadVersion(height, app.keyMain)
}

// ______________________________________________________________________________________________

var _ sdk.StakingHooks = StakingHooks{}

// StakingHooks contains combined distribution and slashing hooks needed for the
// staking module.
type StakingHooks struct {
	dh distribution.Hooks
	sh slashing.Hooks
}

func NewStakingHooks(dh distribution.Hooks, sh slashing.Hooks) StakingHooks {
	return StakingHooks{dh, sh}
}

// nolint
func (h StakingHooks) AfterValidatorCreated(ctx sdk.Context, valAddr sdk.ValAddress) {
	h.dh.AfterValidatorCreated(ctx, valAddr)
	h.sh.AfterValidatorCreated(ctx, valAddr)
}
func (h StakingHooks) BeforeValidatorModified(ctx sdk.Context, valAddr sdk.ValAddress) {
	h.dh.BeforeValidatorModified(ctx, valAddr)
	h.sh.BeforeValidatorModified(ctx, valAddr)
}
func (h StakingHooks) AfterValidatorRemoved(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	h.dh.AfterValidatorRemoved(ctx, consAddr, valAddr)
	h.sh.AfterValidatorRemoved(ctx, consAddr, valAddr)
}
func (h StakingHooks) AfterValidatorBonded(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	h.dh.AfterValidatorBonded(ctx, consAddr, valAddr)
	h.sh.AfterValidatorBonded(ctx, consAddr, valAddr)
}
func (h StakingHooks) AfterValidatorBeginUnbonding(ctx sdk.Context, consAddr sdk.ConsAddress, valAddr sdk.ValAddress) {
	h.dh.AfterValidatorBeginUnbonding(ctx, consAddr, valAddr)
	h.sh.AfterValidatorBeginUnbonding(ctx, consAddr, valAddr)
}
func (h StakingHooks) BeforeDelegationCreated(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.dh.BeforeDelegationCreated(ctx, delAddr, valAddr)
	h.sh.BeforeDelegationCreated(ctx, delAddr, valAddr)
}
func (h StakingHooks) BeforeDelegationSharesModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.dh.BeforeDelegationSharesModified(ctx, delAddr, valAddr)
	h.sh.BeforeDelegationSharesModified(ctx, delAddr, valAddr)
}
func (h StakingHooks) BeforeDelegationRemoved(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.dh.BeforeDelegationRemoved(ctx, delAddr, valAddr)
	h.sh.BeforeDelegationRemoved(ctx, delAddr, valAddr)
}
func (h StakingHooks) AfterDelegationModified(ctx sdk.Context, delAddr sdk.AccAddress, valAddr sdk.ValAddress) {
	h.dh.AfterDelegationModified(ctx, delAddr, valAddr)
	h.sh.AfterDelegationModified(ctx, delAddr, valAddr)
}
func (h StakingHooks) BeforeValidatorSlashed(ctx sdk.Context, valAddr sdk.ValAddress, fraction sdk.Dec) {
	h.dh.BeforeValidatorSlashed(ctx, valAddr, fraction)
	h.sh.BeforeValidatorSlashed(ctx, valAddr, fraction)
}

// ______________________________________________________________________________________________

var _ bank.BankHooks = BankHooks{}

type BankHooks struct {
	boxHooks   box.Hooks
	issueHooks issue.Hooks
	accHooks acc.Hooks

}

func NewBankHooks(boxHooks box.Hooks, issueHooks issue.Hooks, accHooks acc.Hooks) BankHooks {
	return BankHooks{
		boxHooks:   boxHooks,
		issueHooks: issueHooks,
		accHooks: accHooks,
	}
}

// nolint
func (bankHooks BankHooks) CanSend(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) (bool, sdk.Error) {
	_, err := bankHooks.boxHooks.CanSend(ctx, fromAddr, toAddr, amt)
	if err != nil {
		return false, err
	}

	_, err = bankHooks.issueHooks.CanSend(ctx, fromAddr, toAddr, amt)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (bankHooks BankHooks) CheckMustMemoAddress(ctx sdk.Context, toAddr sdk.AccAddress, memo string) (bool, sdk.Error) {
	_, err := bankHooks.accHooks.CheckMustMemoAddress(ctx, toAddr, memo)
	if err != nil {
		return true, err
	}
	return false, nil
}
