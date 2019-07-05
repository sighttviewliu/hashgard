package tests

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/record"
	queriers2 "github.com/hashgard/hashgard/x/record/client/queriers"
	"github.com/hashgard/hashgard/x/record/msgs"
	"github.com/hashgard/hashgard/x/record/types"
)

func TestQueryRecord(t *testing.T) {
	mapp, keeper, _, _, _ := getMockApp(t, record.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.NewContext(false, abci.Header{})

	querier := record.NewQuerier(keeper)
	handler := record.NewHandler(keeper)

	res := handler(ctx, msgs.NewMsgRecord(SenderAccAddr, &RecordParams))
	var recordHash string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &recordHash)
	bz := getQueried(t, ctx, querier, queriers2.GetQueryRecordPath(recordHash), types.QueryRecord, recordHash, nil)
	var recordInfo types.RecordInfo
	keeper.Getcdc().MustUnmarshalJSON(bz, &recordInfo)

	require.Equal(t, recordInfo.Hash, recordHash)
	require.Equal(t, recordInfo.GetName(), RecordInfo.GetName())

}

func TestQueryRecords(t *testing.T) {
	mapp, keeper, _, _, _ := getMockApp(t, record.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.NewContext(false, abci.Header{})
	//querier := record.NewQuerier(keeper)
	handler := record.NewHandler(keeper)
	cap := 10
	for i := 0; i < cap; i++ {
		RecordParams.Hash = RecordParams.Hash[0:len(RecordParams.Hash)-1] + strconv.Itoa(i)
		handler(ctx, msgs.NewMsgRecord(SenderAccAddr, &RecordParams))
	}

	queryParams := keeper.Getcdc().MustMarshalJSON(RecordQueryParams)
	bz := getQueried(t, ctx, record.NewQuerier(keeper), queriers2.GetQueryRecordsPath(), types.QueryRecords, "", queryParams)
	var records []*types.RecordInfo
	keeper.Getcdc().MustUnmarshalJSON(bz, &records)

	require.Len(t, records, cap)
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
