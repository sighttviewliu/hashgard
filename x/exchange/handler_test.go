package exchange

import (
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestHandlerNewMsgMake(t *testing.T) {
	mapp, keeper, _, addrs, _, _ := getMockApp(t, 1, GenesisState{}, nil)
	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx := mapp.NewContext(false, abci.Header{})

	handler := NewHandler(keeper)

	res := handler(ctx, NewMsgMake(addrs[0], sdk.NewInt64Coin(sdk.DefaultBondDenom, 1000), sdk.NewInt64Coin("foocoin", 100)))
	require.True(t, res.IsOK())
}
