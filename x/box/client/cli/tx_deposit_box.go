package cli

import (
	"github.com/hashgard/hashgard/x/box/params"

	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientutils "github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/types"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetCmdDepositBoxCreate implements create deposit box transaction command.
func GetCmdDepositBoxCreate(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create-deposit [name] [total-amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Create a new deposit box",
		Long:    "Create a new deposit box",
		Example: "$ hashgardcli box create-deposit foocoin 100000000coin174876e800 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			// parse coins trying to be sent
			coin, err := sdk.ParseCoin(args[1])
			if err != nil {
				return err
			}

			txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
			if err != nil {
				return err
			}

			decimal, err := clientutils.GetCoinDecimal(cdc, cliCtx, coin)
			if err != nil {
				return err
			}
			coin.Amount = boxutils.MulDecimals(coin, decimal)

			box := params.BoxDepositParams{}
			box.Name = args[0]
			box.BoxType = types.Deposit
			box.TotalAmount = types.BoxToken{Token: coin, Decimals: decimal}
			box.TransferDisabled = viper.GetBool(flagTransferDisabled)
			box.Deposit = types.DepositBox{
				StartTime:     viper.GetInt64(flagStartTime),
				EstablishTime: viper.GetInt64(flagEstablishTime),
				MaturityTime:  viper.GetInt64(flagMaturityTime)}

			num, ok := sdk.NewIntFromString(viper.GetString(flagBottomLine))
			if !ok {
				return errors.Errorf(errors.ErrAmountNotValid(flagBottomLine))
			}
			box.Deposit.BottomLine = num
			num, ok = sdk.NewIntFromString(viper.GetString(flagPrice))
			if !ok {
				return errors.Errorf(errors.ErrAmountNotValid(flagPrice))
			}
			box.Deposit.Price = num
			box.Deposit.Price = boxutils.MulDecimals(boxutils.ParseCoin(box.TotalAmount.Token.Denom, box.Deposit.Price), decimal)
			box.Deposit.BottomLine = boxutils.MulDecimals(boxutils.ParseCoin(box.TotalAmount.Token.Denom, box.Deposit.BottomLine), decimal)

			interest, err := sdk.ParseCoin(viper.GetString(flagInterest))
			if err != nil {
				return err
			}
			decimal, err = clientutils.GetCoinDecimal(cdc, cliCtx, interest)
			if err != nil {
				return err
			}

			interest.Amount = boxutils.MulDecimals(interest, decimal)
			box.Deposit.Interest = types.BoxToken{Token: interest, Decimals: decimal}

			box.Deposit.PerCoupon = boxutils.CalcInterestRate(box.TotalAmount.Token.Amount, box.Deposit.Price,
				box.Deposit.Interest.Token, box.Deposit.Interest.Decimals)

			msg := msgs.NewMsgDepositBox(account.GetAddress(), &box)
			if err := msg.ValidateService(); err != nil {
				return errors.Errorf(err)
			}
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	cmd.Flags().Bool(flagTransferDisabled, true, "Disable the box transfer")
	cmd.Flags().String(flagBottomLine, "", "Box bottom line")
	cmd.Flags().String(flagPrice, "", "Box unit price")
	cmd.Flags().String(flagInterest, "", "Box interest")
	cmd.Flags().Int64(flagStartTime, 0, "Box start time")
	cmd.Flags().Int64(flagEstablishTime, 0, "Box establish time")
	cmd.Flags().Int64(flagMaturityTime, 0, "Box maturity time")

	return cmd
}

// GetCmdDepositBoxInterestInjection implements interest injection a deposit box transaction command.
func GetCmdDepositBoxInterestInjection(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "interest-injection [box-id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Injection interest to the deposit box",
		Long:    "Injection interest to the deposit box",
		Example: "$ hashgardcli box interest-injection box174876e800 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return interest(cdc, args[0], args[1], types.Injection)
		},
	}
	return cmd
}

// GetCmdDepositBoxInterestFetch implements fetch interest from a deposit box transaction command.
func GetCmdDepositBoxInterestFetch(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "interest-fetch [box-id] [amount]",
		Args:    cobra.ExactArgs(2),
		Short:   "Fetch interest from a deposit box",
		Long:    "Fetch interest from a deposit box",
		Example: "$ hashgardcli box interest-fetch box174876e800 88888 --from foo",
		RunE: func(cmd *cobra.Command, args []string) error {
			return interest(cdc, args[0], args[1], types.Fetch)
		},
	}
	return cmd
}

func interest(cdc *codec.Codec, boxID string, amountStr string, operation string) error {

	txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
	if err != nil {
		return err
	}
	msg, err := clientutils.GetInterestMsg(cdc, cliCtx, account, boxID, amountStr, operation, true)
	if err != nil {
		return err
	}

	return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
}

func MarkCmdDepositBoxCreateFlagRequired(cmd *cobra.Command) {
	cmd.MarkFlagRequired(flagBottomLine)
	cmd.MarkFlagRequired(flagPrice)
	cmd.MarkFlagRequired(flagInterest)
	cmd.MarkFlagRequired(flagStartTime)
	cmd.MarkFlagRequired(flagEstablishTime)
	cmd.MarkFlagRequired(flagMaturityTime)
}
