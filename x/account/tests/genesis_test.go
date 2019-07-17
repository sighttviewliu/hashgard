package tests

import (
	"github.com/hashgard/hashgard/x/account"
	"testing"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestImportExportQueues(t *testing.T) {
	mapp, keeper, addrs, _, _ := getMockApp(t, account.DefaultGenesisState(), nil)

	header := abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp.BeginBlock(abci.RequestBeginBlock{Header: header})
	ctx := mapp.BaseApp.NewContext(false, abci.Header{})

	handler := account.NewHandler(keeper)

	// init test data
	for i := 0; i < len(addrs); i++ {
		AccountInfo.Sender = addrs[i]
		handler(ctx, account.NewMsgSetMustMemo(AccountInfo.Sender, AccountInfo.MustMemo))
	}

	genAccs := mapp.AccountKeeper.GetAllAccounts(ctx)

	// Export the state and import it into a new Mock App
	genState := account.ExportGenesis(ctx, keeper)
	mapp2, keeper2, _, _, _ := getMockApp(t, genState, genAccs)

	header = abci.Header{Height: mapp.LastBlockHeight() + 1}
	mapp2.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx2 := mapp2.BaseApp.NewContext(false, abci.Header{})

	for i := 0; i < len(addrs); i++ {
		data, err := keeper2.GetMustMemoAddress(ctx2, addrs[i].String())
		require.Nil(t, err)
		require.NotNil(t, data)
	}
	ls := keeper2.GetMustMemoAddressList(ctx2)
	require.NotNil(t, ls)
	require.Equal(t, len(addrs), len(ls))
}
