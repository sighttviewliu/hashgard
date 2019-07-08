package tests

import (
	"testing"

	"github.com/hashgard/hashgard/x/record"

	"github.com/hashgard/hashgard/x/record/msgs"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestImportExportQueues(t *testing.T) {
	mapp, keeper, _, _, _ := getMockApp(t, record.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	handler := record.NewHandler(keeper)

	res := handler(ctx, msgs.NewMsgRecord(SenderAccAddr, &RecordParams))
	require.True(t, res.IsOK())

	var recordHash1 string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &recordHash1)
	require.NotNil(t, recordHash1)
	require.Equal(t, recordHash1, RecordParams.Hash)

	RecordParams.Hash = "BC38CAEE32149BEF4CCFAEAB518EC9A5FBC85AE6AC8D5A9F6CD710FAF5E4A2B9"
	res = handler(ctx, msgs.NewMsgRecord(SenderAccAddr, &RecordParams))
	require.True(t, res.IsOK())

	var recordHash2 string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &recordHash2)
	require.NotNil(t, recordHash2)
	require.Equal(t, recordHash2, RecordParams.Hash)

	genAccs := mapp.AccountKeeper.GetAllAccounts(ctx)

	// Export the state and import it into a new Mock App
	genState := record.ExportGenesis(ctx, keeper)
	mapp2, keeper2, _, _, _ := getMockApp(t, genState, genAccs)

	header = abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp2.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx2 := mapp2.BaseApp.NewContext(false, abci.Header{})

	recordInfo1 := keeper2.GetRecord(ctx2, recordHash1)
	require.NotNil(t, recordInfo1)
	recordInfo2 := keeper2.GetRecord(ctx2, recordHash2)
	require.NotNil(t, recordInfo2)
}
