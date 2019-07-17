package tests

import (
	"github.com/hashgard/hashgard/x/account"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"testing"
)

func TestAddMustMemoAddr(t *testing.T) {
	mapp, keeper, _, _, _ := getMockApp(t, account.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	// set true
	keeper.SetMustMemoAddress(ctx, &AccountInfo)

	res, err := keeper.GetMustMemoAddress(ctx, AccountInfo.Sender.String())
	require.Equal(t, res.Sender, AccountInfo.Sender)
	require.Equal(t, res.MustMemo, AccountInfo.MustMemo)
	require.Nil(t, err)

	// set false
	AccountInfo.MustMemo = false
	keeper.SetMustMemoAddress(ctx, &AccountInfo)

	res2, err2 := keeper.GetMustMemoAddress(ctx, AccountInfo.Sender.String())
	require.Equal(t, res2.Sender, AccountInfo.Sender)
	require.Equal(t, res2.MustMemo, AccountInfo.MustMemo)
	require.Nil(t, err2)
}

func TestGetAll(t *testing.T) {
	mapp, keeper, addrs, _, _ := getMockApp(t, account.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	for i := 0; i < len(addrs); i++ {
		AccountInfo.Sender = addrs[i]
		keeper.SetMustMemoAddress(ctx, &AccountInfo)
	}
	ls := keeper.GetMustMemoAddressList(ctx)
	require.Len(t, ls, len(addrs))

	// one addr set false
	AccountInfo.Sender = addrs[0]
	AccountInfo.MustMemo = false
	keeper.SetMustMemoAddress(ctx, &AccountInfo)
	ls2 := keeper.GetMustMemoAddressList(ctx)
	require.Len(t, ls2, len(addrs) - 1)
}