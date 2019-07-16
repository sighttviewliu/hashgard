package main

import (
	"fmt"
	"os"
	"path"

	"github.com/hashgard/hashgard/x/exchange"

	"github.com/hashgard/hashgard/x/box"

	"github.com/cosmos/cosmos-sdk/client/keys"
	acc "github.com/cosmos/cosmos-sdk/x/account/client/cli"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/hashgard/hashgard/x/deposit"
	"github.com/hashgard/hashgard/x/future"
	"github.com/hashgard/hashgard/x/lock"

	"github.com/cosmos/cosmos-sdk/client"
	_ "github.com/cosmos/cosmos-sdk/client/lcd/statik"
	"github.com/cosmos/cosmos-sdk/client/rpc"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authcmd "github.com/cosmos/cosmos-sdk/x/auth/client/cli"
	crisiscmd "github.com/cosmos/cosmos-sdk/x/crisis/client/cli"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	slashingcmd "github.com/cosmos/cosmos-sdk/x/slashing/client/cli"
	"github.com/cosmos/cosmos-sdk/x/staking"
	stakecmd "github.com/cosmos/cosmos-sdk/x/staking/client/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/hashgard/hashgard/app"
	hashgardInit "github.com/hashgard/hashgard/init"
	"github.com/hashgard/hashgard/version"
	"github.com/hashgard/hashgard/x/distribution"
	distributioncmd "github.com/hashgard/hashgard/x/distribution/client/cli"
	faucetcmd "github.com/hashgard/hashgard/x/faucet/client/cli"
	"github.com/hashgard/hashgard/x/gov"
	govcmd "github.com/hashgard/hashgard/x/gov/client/cli"
	"github.com/hashgard/hashgard/x/issue"
	mintcmd "github.com/hashgard/hashgard/x/mint/client/cli"
	"github.com/hashgard/hashgard/x/record"
)

// rootCmd is the entry point for this binary
var (
	rootCmd = &cobra.Command{
		Use:   "hashgardcli",
		Short: "Command line interface for interacting with hashgard",
	}
)

func main() {
	// get the codec
	cdc := app.MakeCodec()

	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(hashgardInit.Bech32PrefixAccAddr, hashgardInit.Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(hashgardInit.Bech32PrefixValAddr, hashgardInit.Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(hashgardInit.Bech32PrefixConsAddr, hashgardInit.Bech32PrefixConsPub)
	config.Seal()

	// disable sorting
	cobra.EnableCommandSorting = false

	// TODO: setup keybase, viper object, etc. to be passed into
	// the below functions and eliminate global vars, like we do
	// with the cdc

	// Add --chain-id to persistent flags and mark it required
	rootCmd.PersistentFlags().String(client.FlagChainID, "", "Chain ID of tendermint node")
	rootCmd.PersistentPreRunE = func(_ *cobra.Command, _ []string) error {
		return initConfig(rootCmd)
	}
	rootCmd.AddCommand(
		client.ConfigCmd(app.DefaultCLIHome),
		rpc.StatusCommand(),
		client.LineBreak,
		keys.Commands(),
	)
	// Add tendermint subcommands
	addTendermintCmd(cdc, rootCmd)
	//add x moudle
	rootCmd.AddCommand(client.LineBreak)
	// Add must-memo subcommands
	addMustMemoCmd(cdc, rootCmd)
	// Add bank subcommands
	addBankCmd(cdc, rootCmd)
	// Add distribution subcommands
	addDistributionCmd(cdc, rootCmd)
	// Add exchange subcommands
	addExchangeCmd(cdc, rootCmd)
	// Add gov subcommands
	addGovCmd(cdc, rootCmd)
	// Add issue subcommands
	addIssueCmd(cdc, rootCmd)
	// Add lock subcommands
	addLockCmd(cdc, rootCmd)
	// Add deposit subcommands
	addDepositCmd(cdc, rootCmd)
	// Add future subcommands
	addFutureCmd(cdc, rootCmd)
	// Add record subcommands
	addRecordCmd(cdc, rootCmd)
	// Add slashing subcommands
	addSlashingCmd(cdc, rootCmd)
	// Add stake subcommands
	addStakeCmd(cdc, rootCmd)
	// Add faucet subcommands
	addFaucetCmd(cdc, rootCmd)
	// Add crisis subcommands
	addCrisisCmd(cdc, rootCmd)
	// Add mint subcommands
	addMintCmd(cdc, rootCmd)

	rootCmd.AddCommand(
		client.LineBreak,
		version.VersionCmd,
	)

	// prepare and add flags
	executor := cli.PrepareMainCmd(rootCmd, "BC", app.DefaultCLIHome)
	err := initConfig(rootCmd)
	if err != nil {
		panic(err)
	}

	err = executor.Execute()
	if err != nil {
		fmt.Printf("Failed executing CLI command: %s, exiting...\n", err)
		os.Exit(1)
	}
}

// Add tendermint subcommands
func addTendermintCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	tendermintCmd := &cobra.Command{
		Use:   "tendermint",
		Short: "Tendermint state querying subcommands",
	}
	tendermintCmd.AddCommand(
		rpc.BlockCommand(),
		rpc.ValidatorCommand(cdc),
		tx.SearchTxCmd(cdc),
		tx.QueryTxCmd(cdc),
	)
	rootCmd.AddCommand(tendermintCmd)
}

// Add must-memo subcommands
func addMustMemoCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	mustMemoCmd := &cobra.Command{
		Use:   "must-memo",
		Short: "must-memo subcommands",
	}
	mustMemoCmd.AddCommand(
		acc.GetCmdSetMustMemo(cdc),
		client.LineBreak,
		acc.GetCmdQueryMustMemo(cdc),
		acc.GetCmdQueryMustMemoList(cdc),
		client.LineBreak,
	)
	rootCmd.AddCommand(mustMemoCmd)
}

// Add bank subcommands
func addBankCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	bankCmd := &cobra.Command{
		Use:   "bank",
		Short: "Bank subcommands",
	}
	bankCmd.AddCommand(
		authcmd.GetAccountCmd(auth.StoreKey, cdc),
		//issue.GetAccountCmd(cdc),
		client.LineBreak,
	)
	bankCmd.AddCommand(
		issue.QueryCmd(cdc),
		box.QueryCmd(cdc),
		box.WithdrawCmd(cdc),
		client.LineBreak,
		box.SendTxCmd(cdc),
		authcmd.GetSignCommand(cdc),
		authcmd.GetMultiSignCommand(cdc),
		tx.GetBroadcastCommand(cdc),
		tx.GetEncodeCommand(cdc),
	)
	rootCmd.AddCommand(bankCmd)
}

// Add issue subcommands
func addIssueCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	moduleClient := issue.NewModuleClient(cdc)
	rootCmd.AddCommand(moduleClient.GetCmd())
}

// Add future subcommands
func addFutureCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	moduleClient := future.NewModuleClient(cdc)
	rootCmd.AddCommand(moduleClient.GetCmd())
}

// Add lock subcommands
func addLockCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	moduleClient := lock.NewModuleClient(cdc)
	rootCmd.AddCommand(moduleClient.GetCmd())
}

// Add deposit subcommands
func addDepositCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	moduleClient := deposit.NewModuleClient(cdc)
	rootCmd.AddCommand(moduleClient.GetCmd())
}

// Add record subcommands
func addRecordCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	moduleClient := record.NewModuleClient(cdc)
	rootCmd.AddCommand(moduleClient.GetCmd())
}

// Add exchange subcommands
func addExchangeCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	moduleClient := exchange.NewModuleClient(cdc)
	rootCmd.AddCommand(moduleClient.GetCmd())
}

// Add gov subcommands
func addGovCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	govCmd := &cobra.Command{
		Use:   "gov",
		Short: "Governance subcommands",
	}
	govCmd.AddCommand(
		client.GetCommands(
			govcmd.GetCmdQueryProposal(gov.StoreKey, cdc),
			govcmd.GetCmdQueryProposals(gov.StoreKey, cdc),
			govcmd.GetCmdQueryVote(gov.StoreKey, cdc),
			govcmd.GetCmdQueryVotes(gov.StoreKey, cdc),
			govcmd.GetCmdQueryParam(gov.StoreKey, cdc),
			govcmd.GetCmdQueryParams(gov.StoreKey, cdc),
			govcmd.GetCmdQueryProposer(gov.StoreKey, cdc),
			govcmd.GetCmdQueryDeposit(gov.StoreKey, cdc),
			govcmd.GetCmdQueryDeposits(gov.StoreKey, cdc),
			govcmd.GetCmdQueryTally(gov.StoreKey, cdc),
		)...)
	govCmd.AddCommand(client.LineBreak)
	govCmd.AddCommand(
		client.PostCommands(
			govcmd.GetCmdDeposit(gov.StoreKey, cdc),
			govcmd.GetCmdVote(gov.StoreKey, cdc),
			govcmd.GetCmdSubmitProposal(cdc),
		)...)
	rootCmd.AddCommand(govCmd)
}

// Add stake subcommands
func addStakeCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	stakeCmd := &cobra.Command{
		Use:   "stake",
		Short: "Stake and validation subcommands",
	}
	stakeCmd.AddCommand(
		client.GetCommands(
			stakecmd.GetCmdQueryDelegation(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryDelegations(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryUnbondingDelegation(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryUnbondingDelegations(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryRedelegation(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryRedelegations(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryValidator(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryValidators(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryValidatorDelegations(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryValidatorUnbondingDelegations(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryValidatorRedelegations(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryParams(staking.StoreKey, cdc),
			stakecmd.GetCmdQueryPool(staking.StoreKey, cdc),
		)...)
	stakeCmd.AddCommand(client.LineBreak)
	stakeCmd.AddCommand(
		client.PostCommands(
			stakecmd.GetCmdCreateValidator(cdc),
			stakecmd.GetCmdEditValidator(cdc),
			stakecmd.GetCmdDelegate(cdc),
			stakecmd.GetCmdRedelegate(staking.StoreKey, cdc),
			stakecmd.GetCmdUnbond(staking.StoreKey, cdc),
		)...)
	rootCmd.AddCommand(stakeCmd)
}

// Add slashing subcommands
func addSlashingCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	slashingCmd := &cobra.Command{
		Use:   "slashing",
		Short: "Slashing subcommands",
	}

	slashingCmd.AddCommand(
		client.GetCommands(
			slashingcmd.GetCmdQuerySigningInfo(slashing.StoreKey, cdc),
			slashingcmd.GetCmdQueryParams(cdc),
		)...)
	slashingCmd.AddCommand(client.LineBreak)
	slashingCmd.AddCommand(
		client.PostCommands(
			slashingcmd.GetCmdUnjail(cdc),
		)...)
	rootCmd.AddCommand(slashingCmd)
}

// Add distribution subcommands
func addDistributionCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	distributionCmd := &cobra.Command{
		Use:   "distribution",
		Short: "Distribution subcommands",
	}

	distributionCmd.AddCommand(
		client.GetCommands(
			distributioncmd.GetCmdQueryParams(distribution.StoreKey, cdc),
			distributioncmd.GetCmdQueryValidatorOutstandingRewards(distribution.StoreKey, cdc),
			distributioncmd.GetCmdQueryValidatorCommission(distribution.StoreKey, cdc),
			distributioncmd.GetCmdQueryValidatorSlashes(distribution.StoreKey, cdc),
			distributioncmd.GetCmdQueryDelegatorRewards(distribution.StoreKey, cdc),
			distributioncmd.GetCmdQueryCommunityPool(distribution.StoreKey, cdc),
		)...)
	distributionCmd.AddCommand(
		client.PostCommands(
			distributioncmd.GetCmdWithdrawRewards(cdc),
			distributioncmd.GetCmdSetWithdrawAddr(cdc),
			distributioncmd.GetCmdWithdrawAllRewards(cdc, distribution.StoreKey),
		)...)
	rootCmd.AddCommand(distributionCmd)
}

// Add faucet subcommands
func addFaucetCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	faucetCmd := &cobra.Command{
		Use:   "faucet",
		Short: "faucet subcommands",
	}
	faucetCmd.AddCommand(
		client.PostCommands(
			faucetcmd.GetCmdFaucetSend(cdc),
		)...)
	rootCmd.AddCommand(faucetCmd)
}

// Add crisis subcommands
func addCrisisCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	crisisCmd := &cobra.Command{
		Use:   "crisis",
		Short: "crisis subcommands",
	}
	crisisCmd.AddCommand(
		client.PostCommands(
			crisiscmd.GetCmdInvariantBroken(cdc),
		)...)
	rootCmd.AddCommand(crisisCmd)
}

// Add mint subcommands
func addMintCmd(cdc *codec.Codec, rootCmd *cobra.Command) {
	mintCmd := &cobra.Command{
		Use:   "mint",
		Short: "commands for the minting module",
	}
	mintCmd.AddCommand(
		client.GetCommands(
			mintcmd.GetCmdQueryParams(cdc),
			mintcmd.GetCmdQueryMinter(cdc),
		)...)
	rootCmd.AddCommand(mintCmd)
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

	if err := viper.BindPFlag(cli.EncodingFlag, cmd.PersistentFlags().Lookup(cli.EncodingFlag)); err != nil {
		return err
	}
	return viper.BindPFlag(cli.OutputFlag, cmd.PersistentFlags().Lookup(cli.OutputFlag))
}
