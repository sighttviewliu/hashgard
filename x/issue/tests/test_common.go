package tests

import (
	"testing"

	"github.com/hashgard/hashgard/x/issue/params"

	keeper2 "github.com/cosmos/cosmos-sdk/x/distribution/keeper"

	"github.com/cosmos/cosmos-sdk/x/staking"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/mock"

	"github.com/hashgard/hashgard/x/issue"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/types"

	"github.com/hashgard/hashgard/x/issue/keeper"
)

var (
	ReceiverCoinsAccAddr = sdk.AccAddress(crypto.AddressHash([]byte("receiverCoins")))
	TransferAccAddr      sdk.AccAddress
	SenderAccAddr        sdk.AccAddress

	IssueParams = params.IssueParams{
		Name:               "testCoin",
		Symbol:             "TEST",
		TotalSupply:        sdk.NewInt(10000),
		Decimals:           types.CoinDecimalsMaxValue,
		BurnOwnerDisabled:  false,
		BurnHolderDisabled: false,
		BurnFromDisabled:   false,
		MintingFinished:    false}

	CoinIssueInfo = types.CoinIssueInfo{
		Owner:              SenderAccAddr,
		Issuer:             SenderAccAddr,
		Name:               "testCoin",
		Symbol:             "TEST",
		TotalSupply:        sdk.NewInt(10000),
		Decimals:           types.CoinDecimalsMaxValue,
		BurnOwnerDisabled:  false,
		BurnHolderDisabled: false,
		BurnFromDisabled:   false,
		MintingFinished:    false}
)

// initialize the mock application for this module
func getMockApp(t *testing.T, genState issue.GenesisState, genAccs []auth.Account) (
	mapp *mock.App, keeper keeper.Keeper, sk staking.Keeper, addrs []sdk.AccAddress,
	pubKeys []crypto.PubKey, privKeys []crypto.PrivKey) {
	mapp = mock.NewApp()
	msgs.RegisterCodec(mapp.Cdc)
	keyIssue := sdk.NewKVStoreKey(types.StoreKey)

	keyStaking := sdk.NewKVStoreKey(staking.StoreKey)
	tkeyStaking := sdk.NewTransientStoreKey(staking.TStoreKey)

	pk := mapp.ParamsKeeper
	ck := bank.NewBaseKeeper(mapp.AccountKeeper, mapp.ParamsKeeper.Subspace(bank.DefaultParamspace), bank.DefaultCodespace)
	fck := keeper2.DummyFeeCollectionKeeper{}

	sk = staking.NewKeeper(mapp.Cdc, keyStaking, tkeyStaking, ck, pk.Subspace(staking.DefaultParamspace), staking.DefaultCodespace)
	keeper = issue.NewKeeper(mapp.Cdc, keyIssue, pk, pk.Subspace("testissue"), &ck, fck, types.DefaultCodespace)
	ck.SetHooks(keeper.Hooks())

	mapp.Router().AddRoute(types.RouterKey, issue.NewHandler(keeper))
	mapp.QueryRouter().AddRoute(types.QuerierRoute, issue.NewQuerier(keeper))
	//mapp.SetEndBlocker(getEndBlocker(keeper))
	mapp.SetInitChainer(getInitChainer(mapp, keeper, sk, genState))

	require.NoError(t, mapp.CompleteSetup(keyIssue))

	valTokens := sdk.TokensFromTendermintPower(1000000000000)
	if len(genAccs) == 0 {
		genAccs, addrs, pubKeys, privKeys = mock.CreateGenAccounts(2,
			sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, valTokens)))
	}
	SenderAccAddr = genAccs[0].GetAddress()
	TransferAccAddr = genAccs[1].GetAddress()

	CoinIssueInfo.Owner = SenderAccAddr
	CoinIssueInfo.Issuer = SenderAccAddr

	mock.SetGenesis(mapp, genAccs)

	return mapp, keeper, sk, addrs, pubKeys, privKeys
}
func getInitChainer(mapp *mock.App, keeper keeper.Keeper, stakingKeeper staking.Keeper, genState issue.GenesisState) sdk.InitChainer {

	return func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {

		mapp.InitChainer(ctx, req)

		stakingGenesis := staking.DefaultGenesisState()
		tokens := sdk.TokensFromTendermintPower(100000)
		stakingGenesis.Pool.NotBondedTokens = tokens

		//validators, err := staking.InitGenesis(ctx, stakingKeeper, stakingGenesis)
		//if err != nil {
		//	panic(err)
		//}
		if genState.IsEmpty() {
			issue.InitGenesis(ctx, keeper, issue.DefaultGenesisState())
		} else {
			issue.InitGenesis(ctx, keeper, genState)
		}
		return abci.ResponseInitChain{
			//Validators: validators,
		}
	}
}
