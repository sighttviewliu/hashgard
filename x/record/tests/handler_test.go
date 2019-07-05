package tests

import (
	"testing"

	"github.com/hashgard/hashgard/x/record"
	"github.com/hashgard/hashgard/x/record/msgs"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestHandlerNewMsgRecord(t *testing.T) {
	mapp, keeper, _, _, _ := getMockApp(t, record.GenesisState{}, nil)
	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx := mapp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})

	handler := record.NewHandler(keeper)

	res := handler(ctx, msgs.NewMsgRecord(SenderAccAddr, &RecordParams))
	require.True(t, res.IsOK())

	var recordID string
	keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(res.Data, &recordID)
	require.NotNil(t, recordID)
}
