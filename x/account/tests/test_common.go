package tests

import (
	"github.com/hashgard/hashgard/x/account"
	"testing"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/stretchr/testify/require"

	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/mock"
)

var (
	SenderAccAddr sdk.AccAddress

	AccountInfo = account.AccountInfo{
		Sender:   SenderAccAddr,
		MustMemo: true}

	QueryParams = account.QueryParams{
		Sender: SenderAccAddr}
)

// initialize the mock application for this module
func getMockApp(t *testing.T, genState account.GenesisState, genAccs []auth.Account) (
	mapp *mock.App, keeper account.Keeper, addrs []sdk.AccAddress,
	pubKeys []crypto.PubKey, privKeys []crypto.PrivKey) {
	mapp = mock.NewApp()
	account.RegisterCodec(mapp.Cdc)
	keyMustMemoAddr := sdk.NewKVStoreKey(account.StoreKey)

	pk := mapp.ParamsKeeper

	keeper = account.NewKeeper(mapp.Cdc, keyMustMemoAddr, pk.Subspace("testMustMemo"), account.DefaultCodespace)

	mapp.Router().AddRoute(account.RouterKey, account.NewHandler(keeper))
	mapp.QueryRouter().AddRoute(account.QuerierRoute, account.NewQuerier(keeper))
	//mapp.SetEndBlocker(getEndBlocker(keeper))
	mapp.SetInitChainer(getInitChainer(mapp, keeper, genState))

	require.NoError(t, mapp.CompleteSetup(keyMustMemoAddr))

	valTokens := sdk.TokensFromTendermintPower(1000000000000)
	if len(genAccs) == 0 {
		genAccs, addrs, pubKeys, privKeys = mock.CreateGenAccounts(2,
			sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, valTokens)))
	}
	SenderAccAddr = genAccs[0].GetAddress()

	AccountInfo.Sender = SenderAccAddr

	mock.SetGenesis(mapp, genAccs)

	return mapp, keeper, addrs, pubKeys, privKeys
}
func getInitChainer(mapp *mock.App, keeper account.Keeper, genState account.GenesisState) sdk.InitChainer {

	return func(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain {

		mapp.InitChainer(ctx, req)

		if genState.IsEmpty() {
			account.InitGenesis(ctx, keeper, account.DefaultGenesisState())
		} else {
			account.InitGenesis(ctx, keeper, genState)
		}
		return abci.ResponseInitChain{}
	}
}
