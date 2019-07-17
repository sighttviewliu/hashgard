package tests

import (
	"github.com/hashgard/hashgard/x/account"
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
)

func TestHandlerNewMsg(t *testing.T) {
	mapp, keeper, _, _, _ := getMockApp(t, account.GenesisState{}, nil)
	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx := mapp.NewContext(false, abci.Header{})
	mapp.InitChainer(ctx, abci.RequestInitChain{})

	handler := account.NewHandler(keeper)

	msg := account.NewMsgSetMustMemo(SenderAccAddr, true)
	err := msg.ValidateBasic()
	require.Nil(t, err)

	res := handler(ctx, msg)
	require.True(t, res.IsOK())

	sender := string(res.Tags[1].Value)
	require.NotNil(t, sender)
	require.Equal(t, SenderAccAddr.String(), sender)
}
