package exchange

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/stretchr/testify/require"
)

func TestEqualOrderID(t *testing.T) {
	state1 := GenesisState{}
	state2 := GenesisState{}
	require.Equal(t, state1, state2)

	// Orders
	state1.StartingOrderId = 1
	require.NotEqual(t, state1, state2)
	require.False(t, state1.Equal(state2))

	state2.StartingOrderId = 1
	require.Equal(t, state1, state2)
	require.True(t, state1.Equal(state2))
}

func TestImportExportQueues(t *testing.T) {
	mapp, keeper, _, addrs, _, _ := getMockApp(t, 1, GenesisState{}, nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	order, err := keeper.Make(ctx, addrs[0], sdk.NewInt64Coin(sdk.DefaultBondDenom, 100), sdk.NewInt64Coin("foocoin", 200))
	require.NoError(t, err)
	id := order.OrderId

	genAccs := mapp.AccountKeeper.GetAllAccounts(ctx)

	// Export the state and import it into a new Mock App
	genState := ExportGenesis(ctx, keeper)
	mapp2, keeper2, _, _, _, _ := getMockApp(t, 1, genState, genAccs)

	header = abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp2.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx2 := mapp2.BaseApp.NewContext(false, abci.Header{})

	_, err = keeper2.GetOrder(ctx2, id)
	require.NoError(t, err)

	orders, err := keeper2.GetOrdersByAddr(ctx2, addrs[0])
	require.NoError(t, err)
	require.Len(t, orders, 1)
}
