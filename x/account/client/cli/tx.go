package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	acc "github.com/hashgard/hashgard/x/account"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
)

func GetCliContext(cdc *codec.Codec) (authtxb.TxBuilder, context.CLIContext, auth.Account, error) {
	txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
	cliCtx := context.NewCLIContext().
		WithCodec(cdc).
		WithAccountDecoder(cdc)
	from := cliCtx.GetFromAddress()
	account, err := cliCtx.GetAccount(from)
	return txBldr, cliCtx, account, err
}

func GetCmdSetMustMemo(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set [is_must_memo] --flags",
		Args:    cobra.ExactArgs(1),
		Short:   "Set tx transfer in must have memo",
		Long:    "Set a tag for your address to ensure any tx transfer in must have memo",
		Example: "$ hashgardcli must-memo set true --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr, cliCtx, account, err := GetCliContext(cdc)
			if err != nil {
				return err
			}

			mustMemo, err := strconv.ParseBool(args[0])
			if err != nil {
				return err
			}

			msg := acc.NewMsgSetMustMemo(account.GetAddress(), mustMemo)

			validateErr := msg.ValidateBasic()
			if validateErr != nil {
				return validateErr
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd = client.PostCommands(cmd)[0]
	cmd.MarkFlagRequired(client.FlagFrom)

	return cmd
}
