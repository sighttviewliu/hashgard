package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientutils "github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/spf13/cobra"
)

// WithdrawCmd implements withdraw a box transaction command.
func WithdrawCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "withdraw [box-id]",
		Args:    cobra.ExactArgs(1),
		Short:   "Withdraw a box from the account coins",
		Long:    "Box holder withdraw a deposit box or future box from the account coins when the box can be withdraw",
		Example: "$ hashgardcli bank withdraw boxab3jlxpt2ps --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return processBoxWithdrawCmd(cdc, args[0])
		},
	}
	cmd = client.PostCommands(cmd)[0]
	_ = cmd.MarkFlagRequired(client.FlagFrom)
	return cmd
}

// ProcessBoxWithdrawCmd implements withdraw a box transaction command.
func processBoxWithdrawCmd(cdc *codec.Codec, id string) error {
	txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
	if err != nil {
		return err
	}
	msg, err := clientutils.GetWithdrawMsg(cdc, cliCtx, account, id)
	if err != nil {
		return err
	}
	return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)

}
