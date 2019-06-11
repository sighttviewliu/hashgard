package gov

import (
	"bytes"
	"log"
	"sort"
	"testing"

	"github.com/cosmos/cosmos-sdk/x/slashing"

	"github.com/hashgard/hashgard/x/mint"

	"github.com/hashgard/hashgard/x/distribution"

	"github.com/hashgard/hashgard/x/issue"

	keeper2 "github.com/hashgard/hashgard/x/distribution/keeper"

	boxtypes "github.com/hashgard/hashgard/x/box/types"
	issuetypes "github.com/hashgard/hashgard/x/issue/types"

	"github.com/hashgard/hashgard/x/box"
	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/mock"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

func getMockApp(t *testing.T, numGenAccs int, genState GenesisState, genAccs []auth.Account) (
	mapp *mock.App, keeper Keeper, sk staking.Keeper, addrs []sdk.AccAddress,
	pubKeys []crypto.PubKey, privKeys []crypto.PrivKey) {
	mapp, keeper, sk, _, _, addrs, pubKeys, privKeys = getMockAppParams(t, numGenAccs, genState, genAccs)
	return mapp, keeper, sk, addrs, pubKeys, privKeys
}

// initialize the mock application for this module
func getMockAppParams(t *testing.T, numGenAccs int, genState GenesisState, genAccs []auth.Account) (
	mapp *mock.App, keeper Keeper, sk staking.Keeper, boxKeeper box.Keeper, issueKeeper issue.Keeper, addrs []sdk.AccAddress,
	pubKeys []crypto.PubKey, privKeys []crypto.PrivKey) {

	mapp = mock.NewApp()

	staking.RegisterCodec(mapp.Cdc)
	RegisterCodec(mapp.Cdc)

	keyStaking := sdk.NewKVStoreKey(staking.StoreKey)
	tkeyStaking := sdk.NewTransientStoreKey(staking.TStoreKey)
	keyGov := sdk.NewKVStoreKey(StoreKey)
	keyDistribution := sdk.NewKVStoreKey(distribution.StoreKey)
	keyMint := sdk.NewKVStoreKey(mint.StoreKey)
	keySlashing := sdk.NewKVStoreKey(slashing.StoreKey)

	fck := keeper2.DummyFeeCollectionKeeper{}
	pk := mapp.ParamsKeeper
	ck := bank.NewBaseKeeper(mapp.AccountKeeper, mapp.ParamsKeeper.Subspace(bank.DefaultParamspace), bank.DefaultCodespace)
	sk = staking.NewKeeper(mapp.Cdc, keyStaking, tkeyStaking, ck, pk.Subspace(staking.DefaultParamspace), staking.DefaultCodespace)
	distributionKeeper := distribution.NewKeeper(
		mapp.Cdc,
		keyDistribution,
		mapp.ParamsKeeper.Subspace(distribution.DefaultParamspace),
		&ck,
		&sk,
		fck,
		distribution.DefaultCodespace,
	)
	mintKeeper := mint.NewKeeper(
		mapp.Cdc,
		keyMint,
		mapp.ParamsKeeper.Subspace(mint.DefaultParamspace),
		&sk,
		fck,
		ck,
	)
	slashingKeeper := slashing.NewKeeper(
		mapp.Cdc,
		keySlashing,
		&sk,
		mapp.ParamsKeeper.Subspace(slashing.DefaultParamspace),
		slashing.DefaultCodespace,
	)
	keeper = NewKeeper(
		mapp.Cdc,
		keyGov,
		pk,
		pk.Subspace("testgov"),
		ck,
		sk,
		DefaultCodespace,
		mapp.AccountKeeper,
		distributionKeeper,
		mintKeeper,
		slashingKeeper,
		sk,
	)

	keyBox := sdk.NewKVStoreKey(boxtypes.StoreKey)
	keyIssue := sdk.NewKVStoreKey(issuetypes.StoreKey)
	//fck := keeper2.DummyFeeCollectionKeeper{}

	issueKeeper = issue.NewKeeper(mapp.Cdc, keyIssue, mapp.ParamsKeeper, mapp.ParamsKeeper.Subspace(issuetypes.ModuleName), &ck,
		fck, issuetypes.DefaultCodespace)

	boxKeeper = box.NewKeeper(mapp.Cdc, keyBox, mapp.ParamsKeeper, mapp.ParamsKeeper.Subspace(boxtypes.ModuleName),
		&ck, issueKeeper, fck, boxtypes.DefaultCodespace)

	mapp.Router().AddRoute(RouterKey, NewHandler(keeper))
	mapp.QueryRouter().AddRoute(QuerierRoute, NewQuerier(keeper))

	mapp.SetEndBlocker(getEndBlocker(keeper))
	mapp.SetInitChainer(getInitChainer(mapp, keeper, issueKeeper, boxKeeper, mintKeeper, sk, slashingKeeper, genState))

	require.NoError(t, mapp.CompleteSetup(keyStaking, tkeyStaking, keyDistribution, keyMint, keySlashing, keyGov, keyIssue, keyBox))

	valTokens := sdk.TokensFromTendermintPower(42)
	if genAccs == nil || len(genAccs) == 0 {
		genAccs, addrs, pubKeys, privKeys = mock.CreateGenAccounts(numGenAccs,
			sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, valTokens)})
	}

	mock.SetGenesis(mapp, genAccs)

	return mapp, keeper, sk, boxKeeper, issueKeeper, addrs, pubKeys, privKeys
}

// gov and staking endblocker
func getEndBlocker(keeper Keeper) sdk.EndBlocker {
	return func(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock {
		tags := EndBlocker(ctx, keeper)
		return abci.ResponseEndBlock{
			Tags: tags,
		}
	}
}

// gov and staking initchainer
func getInitChainer(mapp *mock.App, keeper Keeper, issueKeeper issue.Keeper, boxKeeper box.Keeper, mintKeeper mint.Keeper,
	stakingKeeper staking.Keeper, slashingKeeper slashing.Keeper, genState GenesisState) sdk.InitChainer {
	return func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {
		mapp.InitChainer(ctx, req)

		stakingGenesis := staking.DefaultGenesisState()
		tokens := sdk.TokensFromTendermintPower(100000)
		stakingGenesis.Pool.NotBondedTokens = tokens

		validators, err := staking.InitGenesis(ctx, stakingKeeper, stakingGenesis)
		if err != nil {
			panic(err)
		}
		if genState.IsEmpty() {
			InitGenesis(ctx, keeper, DefaultGenesisState())
			box.InitGenesis(ctx, boxKeeper, box.DefaultGenesisState())
			issue.InitGenesis(ctx, issueKeeper, issue.DefaultGenesisState())
			mint.InitGenesis(ctx, mintKeeper, mint.DefaultGenesisState())
			slashing.InitGenesis(ctx, slashingKeeper, slashing.DefaultGenesisState(), nil)
		} else {
			InitGenesis(ctx, keeper, genState)
		}
		return abci.ResponseInitChain{
			Validators: validators,
		}
	}
}

// TODO: Remove once address interface has been implemented (ref: #2186)
func SortValAddresses(addrs []sdk.ValAddress) {
	var byteAddrs [][]byte
	for _, addr := range addrs {
		byteAddrs = append(byteAddrs, addr.Bytes())
	}

	SortByteArrays(byteAddrs)

	for i, byteAddr := range byteAddrs {
		addrs[i] = byteAddr
	}
}

// Sorts Addresses
func SortAddresses(addrs []sdk.AccAddress) {
	var byteAddrs [][]byte
	for _, addr := range addrs {
		byteAddrs = append(byteAddrs, addr.Bytes())
	}
	SortByteArrays(byteAddrs)
	for i, byteAddr := range byteAddrs {
		addrs[i] = byteAddr
	}
}

// implement `Interface` in sort package.
type sortByteArrays [][]byte

func (b sortByteArrays) Len() int {
	return len(b)
}

func (b sortByteArrays) Less(i, j int) bool {
	// bytes package already implements Comparable for []byte.
	switch bytes.Compare(b[i], b[j]) {
	case -1:
		return true
	case 0, 1:
		return false
	default:
		log.Panic("not fail-able with `bytes.Comparable` bounded [-1, 1].")
		return false
	}
}

func (b sortByteArrays) Swap(i, j int) {
	b[j], b[i] = b[i], b[j]
}

// Public
func SortByteArrays(src [][]byte) [][]byte {
	sorted := sortByteArrays(src)
	sort.Sort(sorted)
	return sorted
}

func testProposal() TextProposal {
	return NewTextProposal("Test", "description")
}

// checks if two proposals are equal (note: slow, for tests only)
func ProposalEqual(proposalA Proposal, proposalB Proposal) bool {
	return bytes.Equal(msgCdc.MustMarshalBinaryBare(proposalA), msgCdc.MustMarshalBinaryBare(proposalB))
}
