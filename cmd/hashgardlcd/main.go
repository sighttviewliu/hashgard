package main

import (
	"net/http"
	"os"
	"path"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"
	at "github.com/cosmos/cosmos-sdk/x/auth"
	auth "github.com/cosmos/cosmos-sdk/x/auth/client/rest"
	bank "github.com/cosmos/cosmos-sdk/x/bank/client/rest"
	slashing "github.com/cosmos/cosmos-sdk/x/slashing/client/rest"
	staking "github.com/cosmos/cosmos-sdk/x/staking/client/rest"
	"github.com/rakyll/statik/fs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"

	deposit "github.com/hashgard/hashgard/x/deposit/client/rest"
	exchange "github.com/hashgard/hashgard/x/exchange/client/rest"
	future "github.com/hashgard/hashgard/x/future/client/rest"
	issue "github.com/hashgard/hashgard/x/issue/client/rest"
	lock "github.com/hashgard/hashgard/x/lock/client/rest"
	record "github.com/hashgard/hashgard/x/record/client/rest"

	distributioncmd "github.com/cosmos/cosmos-sdk/x/distribution"

	"github.com/hashgard/hashgard/app"
	"github.com/hashgard/hashgard/client/lcd"
	_ "github.com/hashgard/hashgard/client/lcd/statik"
	hashgardInit "github.com/hashgard/hashgard/init"
	"github.com/hashgard/hashgard/version"
	distribution "github.com/hashgard/hashgard/x/distribution/client/rest"
	gov "github.com/hashgard/hashgard/x/gov/client/rest"
	mint "github.com/hashgard/hashgard/x/mint/client/rest"
)

// rootCmd is the entry point for this binary
var (
	rootCmd = &cobra.Command{
		Use:   "hashgardlcd",
		Short: "hashgard lcd server",
	}
)

func main() {

	// Add --chain-id to persistent flags and mark it required
	rootCmd.PersistentFlags().String(client.FlagChainID, "", "Chain ID of tendermint node")
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}

	hashgardInit.InitBech32Prefix()

	cobra.EnableCommandSorting = false
	cdc := app.MakeCodec()

	rootCmd.AddCommand(
		lcd.ServeCommand(cdc, registerRoutes),
		version.VersionCmd,
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(rootCmd, "HASHGARDLCD", app.DefaultLCDHome)
	err := executor.Execute()
	if err != nil {
		// handle with #870
		panic(err)
	}
}

// registerRoutes registers the routes from the different modules for the LCD.
// NOTE: details on the routes added for each module are in the module documentation
// NOTE: If making updates here you also need to update the test helper in client/lcd/test_helper.go
func registerRoutes(rs *lcd.RestServer) {
	registerSwaggerUI(rs)
	rpc.RegisterRoutes(rs.CliCtx, rs.Mux)
	tx.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	auth.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, at.StoreKey)
	bank.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	distribution.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, distributioncmd.StoreKey)
	staking.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	slashing.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc, rs.KeyBase)
	gov.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	issue.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	record.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	lock.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	deposit.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	future.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	exchange.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
	mint.RegisterRoutes(rs.CliCtx, rs.Mux, rs.Cdc)
}

func registerSwaggerUI(rs *lcd.RestServer) {
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	staticServer := http.FileServer(statikFS)
	rs.Mux.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", staticServer))
}

func initConfig(cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(cli.HomeFlag)
	if err != nil {
		return err
	}

	cfgFile := path.Join(home, "config", "config.toml")
	if _, err := os.Stat(cfgFile); err == nil {
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}
	if err := viper.BindPFlag(client.FlagChainID, cmd.PersistentFlags().Lookup(client.FlagChainID)); err != nil {
		return err
	}
	if err := viper.BindPFlag(cli.EncodingFlag, cmd.PersistentFlags().Lookup(cli.EncodingFlag)); err != nil {
		return err
	}
	return viper.BindPFlag(cli.OutputFlag, cmd.PersistentFlags().Lookup(cli.OutputFlag))
}
