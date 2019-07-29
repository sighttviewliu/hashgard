package app

import (
	"os"
	"testing"

	"github.com/hashgard/hashgard/x/issue"

	"github.com/hashgard/hashgard/x/box"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"

	abci "github.com/tendermint/tendermint/abci/types"

	acc "github.com/cosmos/cosmos-sdk/x/account"
	"github.com/hashgard/hashgard/x/distribution"
	"github.com/hashgard/hashgard/x/exchange"
	"github.com/hashgard/hashgard/x/gov"
	"github.com/hashgard/hashgard/x/mint"
)

func setGenesis(happ *HashgardApp, accs ...*auth.BaseAccount) error {
	genaccs := make([]GenesisAccount, len(accs))
	for i, acc := range accs {
		genaccs[i] = NewGenesisAccount(acc)
	}

	genesisState := GenesisState{
		Accounts:         genaccs,
		AccMustMemoData:  acc.DefaultGenesisState(),
		AuthData:         auth.DefaultGenesisState(),
		BankData:         bank.DefaultGenesisState(),
		StakingData:      staking.DefaultGenesisState(),
		MintData:         mint.DefaultGenesisState(),
		DistributionData: distribution.DefaultGenesisState(),
		SlashingData:     slashing.DefaultGenesisState(),
		GovData:          gov.DefaultGenesisState(),
		ExchangeData:     exchange.DefaultGenesisState(),
		BoxData:          box.DefaultGenesisState(),
		IssueData:        issue.DefaultGenesisState(),
		CrisisData:       crisis.DefaultGenesisState(),
	}

	stateBytes, err := codec.MarshalJSONIndent(happ.cdc, genesisState)
	if err != nil {
		return err
	}

	// Initialize the chain
	vals := []abci.ValidatorUpdate{}
	happ.InitChain(abci.RequestInitChain{Validators: vals, AppStateBytes: stateBytes})
	happ.Commit()

	return nil
}

func TestHashgardExport(t *testing.T) {
	db := db.NewMemDB()
	gapp := NewHashgardApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0)
	err := setGenesis(gapp)
	require.NoError(t, err, "setGenesis should not have an error")

	// Making a new app object with the db, so that initchain hasn't been called
	newHapp := NewHashgardApp(log.NewTMLogger(log.NewSyncWriter(os.Stdout)), db, nil, true, 0)
	_, _, err = newHapp.ExportAppStateAndValidators(false, []string{})
	require.NoError(t, err, "ExportAppStateAndValidators should not have an error")
}
