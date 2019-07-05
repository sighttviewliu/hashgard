package tests

import (
	"github.com/hashgard/hashgard/x/record/params"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	"strconv"
	"testing"

	"github.com/hashgard/hashgard/x/record"
)

func TestCreateRecord(t *testing.T) {
	mapp, keeper, _, _, _ := getMockApp(t, record.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	err := keeper.CreateRecord(ctx, &RecordInfo)
	require.Nil(t, err)
	recordRes := keeper.GetRecord(ctx, RecordInfo.Hash)
	require.Equal(t, recordRes.Hash, RecordInfo.Hash)
	require.Equal(t, recordRes.Name, RecordInfo.Name)
	require.Equal(t, recordRes.Type, RecordInfo.Type)
	require.Equal(t, recordRes.Author, RecordInfo.Author)
	require.Equal(t, recordRes.RecordNo, RecordInfo.RecordNo)
}

func TestCreateRecordDuplicated(t *testing.T) {
	mapp, keeper, _, _, _ := getMockApp(t, record.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	err := keeper.CreateRecord(ctx, &RecordInfo)
	require.Nil(t, err)
	err2 := keeper.CreateRecord(ctx, &RecordInfo)
	require.NotNil(t, err2)
}

func TestGetRecords(t *testing.T) {
	mapp, keeper, _, _, _ := getMockApp(t, record.GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	cap := 10
	for i := 0; i < cap; i++ {
		RecordInfo.Hash = RecordInfo.Hash[0:len(RecordInfo.Hash)-2] + strconv.Itoa(i)
		err := keeper.CreateRecord(ctx, &RecordInfo)
		require.Nil(t, err)
	}
	records := keeper.List(ctx, params.RecordQueryParams{
		Sender: RecordInfo.Sender,
		Author: RecordInfo.Author,
	})

	require.Len(t, records, cap)
}