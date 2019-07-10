package cli

import (
	"strings"

	"github.com/hashgard/hashgard/x/record/params"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hashgard/hashgard/x/record/errors"
	"github.com/hashgard/hashgard/x/record/msgs"
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

// GetCmdIssue implements record a coin transaction command.
func GetCmdRecordCreate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create [name] [hash]",
		Args:    cobra.ExactArgs(2),
		Short:   "Create a new record",
		Long:    "Create a new record",
		Example: "$ hashgardcli record create myRecord BC38CAEE32149BEF4CCFAEAB518EC9A5FBC85AE6AC8D5A9F6CD710FAF5E4A2B8 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr, cliCtx, account, err := GetCliContext(cdc)
			if err != nil {
				return err
			}

			para := params.RecordParams{
				Name:               args[0],
				Hash:             	strings.ToUpper(args[1]),
				Type:  				viper.GetString(flagType),
				Author:  			viper.GetString(flagAuthor),
				RecordNo:  			viper.GetString(flagRecordNo),
				Description:  		viper.GetString(flagDescription),
			}
			msg := msgs.NewMsgRecord(account.GetAddress(), &para)

			validateErr := msg.ValidateBasic()

			if validateErr != nil {
				return errors.Errorf(validateErr)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().String(flagType, "", "customized record-type")
	cmd.Flags().String(flagAuthor, "", "author of the record data")
	cmd.Flags().String(flagRecordNo, "", "customized record-number")
	cmd.Flags().String(flagDescription, "", "customized description of the record")

	return cmd
}