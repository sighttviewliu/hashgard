package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/hashgard/hashgard/x/exchange/msgs"
)

func GetMakeCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "make",
		Short: "Create a new order",
		Long: "The supply must have specific amount and coin name, that's what you want to sell." +
			"So make sure your address have sufficient balance.The target is what you want to get by this order.",
		Example: "$ hashgardcli exchange make --supply=100gard --target=800apple --from mykey",
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)

			// Get from address
			from := cliCtx.GetFromAddress()

			// Pull associated account
			account, err := cliCtx.GetAccount(from)
			if err != nil {
				return err
			}

			// Find supply
			supplyStr := viper.GetString(FlagSupply)
			supply, err := sdk.ParseCoin(supplyStr)
			if err != nil {
				return err
			}

			// ensure account has enough coins
			if !account.GetCoins().IsAllGTE([]sdk.Coin{supply}) {
				return fmt.Errorf("address %s doesn't have enough coins to create this order", from)
			}

			// Find target
			targetStr := viper.GetString(FlagTarget)
			target, err := sdk.ParseCoin(targetStr)
			if err != nil {
				return err
			}

			msg := msgs.NewMsgMake(from, supply, target)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().String(FlagSupply, "", "coin of supply, eg. 100000000agard")
	cmd.Flags().String(FlagTarget, "", "coin of target, eg. 825000apple")

	return cmd
}
func GetTakeCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "take [id]",
		Args:  cobra.ExactArgs(1),
		Short: "Take an order",
		Long: "Make sure the --amount is match the order's target coin." +
			"if send more than the remains of the order target, the order will be filled," +
			"The extra part will be returned to you. If your amount is more than the exchange threshold," +
			"but less than the remains target, the order will be executed partially, " +
			"you can get supply coins of corresponding amount.",
		Example: "$ hashgardcli exchange take 3 --amount=800apple --from mykey",
		RunE: func(cnd *cobra.Command, args []string) error {
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)

			// Get from address
			from := cliCtx.GetFromAddress()

			// Pull associated account
			account, err := cliCtx.GetAccount(from)
			if err != nil {
				return err
			}

			// Get id
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("id %s not a valid int, please input a valid id", args[0])
			}

			// Todo: check to see if the order is in the store

			// Get amount
			amountStr := viper.GetString(FlagAmount)
			amount, err := sdk.ParseCoin(amountStr)
			if err != nil {
				return err
			}

			// ensure account has enough coins
			if !account.GetCoins().IsAllGTE([]sdk.Coin{amount}) {
				return fmt.Errorf("address %s doesn't have enough coins to take order with specific amount", from)
			}

			msg := msgs.NewMsgTake(id, from, amount)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().String(FlagAmount, "", "coin to take order")

	return cmd
}

func GetCancelCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel [id]",
		Args:  cobra.ExactArgs(1),
		Short: "Cancel an existing order",
		Long: "If the order's owner is not you, will output error." +
			"When withdrawal an order successfully, the remains coin of the order will return to the owner address.'",
		Example: "$ hashgardcli exchange cancel 3 --from mykey",
		RunE: func(cnd *cobra.Command, args []string) error {
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)

			// Get from address
			from := cliCtx.GetFromAddress()

			// Get id
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("id %s not a valid int, please input a valid id", args[0])
			}

			// Todo: check to see if the order is in the store

			msg := msgs.NewMsgCancel(id, from)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	return cmd
}
