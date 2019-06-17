package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/hashgard/hashgard/x/exchange/client/cli"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"
)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	cdc *amino.Codec
}

//New ModuleClient Instance
func NewModuleClient(cdc *amino.Codec) ModuleClient {
	return ModuleClient{cdc}
}

// GetLockCmd returns the box commands for this module
func (mc ModuleClient) GetCmd() *cobra.Command {
	exchangeCmd := &cobra.Command{
		Use:   "exchange",
		Short: "Exchange (HRC11) subcommands",
	}
	exchangeCmd.AddCommand(
		client.GetCommands(
			cli.GetQueryCmd(mc.cdc),
			cli.GetListCmd(mc.cdc),
			cli.GetFrozenFundCmd(mc.cdc),
		)...)
	exchangeCmd.AddCommand(client.LineBreak)

	txCmd := client.PostCommands(
		cli.GetMakeCmd(mc.cdc),
		cli.GetTakeCmd(mc.cdc),
		cli.GetCancelCmd(mc.cdc),
	)

	for _, cmd := range txCmd {
		_ = cmd.MarkFlagRequired(client.FlagFrom)
		exchangeCmd.AddCommand(cmd)
	}

	return exchangeCmd
}
