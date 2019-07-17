package tests

import (
	"github.com/hashgard/hashgard/x/account"
	"testing"

	queriers2 "github.com/hashgard/hashgard/x/account/client/queriers"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestQueryAddr(t *testing.T) {
	mapp, keeper, _, _, _ := getMockApp(t, account.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.NewContext(false, abci.Header{})

	querier := account.NewQuerier(keeper)
	handler := account.NewHandler(keeper)

	// set data by handler
	res := handler(ctx, account.NewMsgSetMustMemo(SenderAccAddr, AccountInfo.MustMemo))
	sender := string(res.Tags[1].Value)
	require.NotNil(t, sender)
	require.Equal(t, AccountInfo.Sender.String(), sender)

	// get data by querier
	bz := getQueried(t, ctx, querier, queriers2.GetQueryMustMemoAddressPath(sender), account.QueryMustMemoAddress, sender, nil)
	var acc account.AccountInfo
	keeper.Getcdc().MustUnmarshalJSON(bz, &acc)

	require.Equal(t, AccountInfo.Sender, acc.GetSender())
	require.Equal(t, AccountInfo.MustMemo, acc.GetMustMemo())
}

func TestQueryRecords(t *testing.T) {
	mapp, keeper, addrs, _, _ := getMockApp(t, account.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.NewContext(false, abci.Header{})
	//querier := record.NewQuerier(keeper)
	handler := account.NewHandler(keeper)
	for i := 0; i < len(addrs); i++ {
		AccountInfo.Sender = addrs[i]
		handler(ctx, account.NewMsgSetMustMemo(AccountInfo.Sender, AccountInfo.MustMemo))
	}

	// query all
	bz := getQueried(t, ctx, account.NewQuerier(keeper), queriers2.GetQueryMustMemoAddressListPath(), account.QueryMustMemoAddresses, "", nil)
	var records []*account.AccountInfo
	keeper.Getcdc().MustUnmarshalJSON(bz, &records)
	require.Len(t, records, len(addrs))
}

func getQueried(t *testing.T, ctx sdk.Context, querier sdk.Querier, path string, querierRoute string, queryPathParam string, queryParam []byte) (res []byte) {
	query := abci.RequestQuery{
		Path: path,
		Data: queryParam,
	}
	bz, err := querier(ctx, []string{querierRoute, queryPathParam}, query)
	require.Nil(t, err)
	require.NotNil(t, bz)

	return bz
}
