package main

import (
	"encoding/json"
	"io"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/libs/cli"
	dbm "github.com/tendermint/tendermint/libs/db"
	"github.com/tendermint/tendermint/libs/log"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/hashgard/hashgard/app"
	hashgardInit "github.com/hashgard/hashgard/init"
	"github.com/hashgard/hashgard/version"
)

const flagInvCheckPeriod = "inv-check-period"

var invCheckPeriod uint

func main() {
	cdc := app.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(hashgardInit.Bech32PrefixAccAddr, hashgardInit.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(hashgardInit.Bech32PrefixValAddr, hashgardInit.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(hashgardInit.Bech32PrefixConsAddr, hashgardInit.Bech32PrefixConsPub)
	config.Seal()

	ctx := server.NewDefaultContext()
	cobra.EnableCommandSorting = false
	rootCmd := &cobra.Command{
		Use:               "hashgard",
		Short:             "Hashgard Daemon (server)",
		PersistentPreRunE: server.PersistentPreRunEFn(ctx),
	}

	tendermintCmd := &cobra.Command{
		Use:   "tendermint",
		Short: "Tendermint subcommands",
	}

	tendermintCmd.AddCommand(
		server.ShowNodeIDCmd(ctx),
		server.ShowValidatorCmd(ctx),
		server.ShowAddressCmd(ctx),
	)

	rootCmd.AddCommand(
		server.StartCmd(ctx, newApp),
		hashgardInit.InitCmd(ctx, cdc),
		hashgardInit.CollectGenTxsCmd(ctx, cdc),
		hashgardInit.TestnetFilesCmd(ctx, cdc),
		hashgardInit.GenTxCmd(ctx, cdc),
		hashgardInit.AddGenesisAccountCmd(ctx, cdc),
		hashgardInit.ValidateGenesisCmd(ctx, cdc),
		server.UnsafeResetAllCmd(ctx),
		client.LineBreak,
		tendermintCmd,
		client.LineBreak,
		server.ExportCmd(ctx, cdc, exportAppStateAndTMValidators),
		client.LineBreak,
		version.VersionCmd,
	)

	// prepare and add flags
	executor := cli.PrepareBaseCmd(rootCmd, "BC", app.DefaultNodeHome)
	rootCmd.PersistentFlags().UintVar(&invCheckPeriod, flagInvCheckPeriod,
		0, "Assert registered invariants every N blocks")
	err := executor.Execute()
	if err != nil {
		// handle with #870
		panic(err)
	}
}

func newApp(logger log.Logger, db dbm.DB, traceStore io.Writer) abci.Application {
	return app.NewHashgardApp(
		logger,
		db,
		traceStore,
		true,
		invCheckPeriod,
		baseapp.SetPruning(store.NewPruningOptionsFromString(viper.GetString("pruning"))),
		baseapp.SetMinGasPrices(viper.GetString(server.FlagMinGasPrices)),
	)
}

func exportAppStateAndTMValidators(
	logger log.Logger, db dbm.DB, traceStore io.Writer, height int64, forZeroHeight bool,
	jailWhiteList []string) (json.RawMessage, []tmtypes.GenesisValidator, error) {
	if height != -1 {
		hApp := app.NewHashgardApp(logger, db, traceStore, false, uint(1))
		err := hApp.LoadHeight(height)
		if err != nil {
			return nil, nil, err
		}

		return hApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
	}

	hApp := app.NewHashgardApp(logger, db, traceStore, true, uint(1))
	return hApp.ExportAppStateAndValidators(forZeroHeight, jailWhiteList)
}
