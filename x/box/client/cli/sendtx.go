package cli

import (
	"fmt"
	"github.com/spf13/viper"

	"github.com/hashgard/hashgard/x/box/errors"

	acc "github.com/hashgard/hashgard/x/account"
	"github.com/hashgard/hashgard/x/box/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/cosmos/cosmos-sdk/x/bank"
	boxclientutils "github.com/hashgard/hashgard/x/box/client/utils"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
	issueclientutils "github.com/hashgard/hashgard/x/issue/client/utils"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
	"github.com/spf13/cobra"
)

// SendTxCmd will create a send tx and sign it with the given key.
func SendTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send [to_address] [amount]",
		Short: "Create and sign a send tx",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			to, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			// check if address_to must memo
			memo := viper.GetString(client.FlagMemo)
			if len(memo) == 0 {
				bz, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", acc.QuerierRoute, acc.QueryMustMemoAddress, to), nil)
				if err != nil {
					return err
				}
				var accStatus acc.AccountInfo
				cdc.MustUnmarshalJSON(bz, &accStatus)
				if accStatus.MustMemo {
					return fmt.Errorf("Memo is required to transfer to address %s ", to)
				}
			}

			// parse coins trying to be sent
			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}
			from := cliCtx.GetFromAddress()

			for i, coin := range coins {
				if err = processBoxSend(cdc, cliCtx, &coin); err != nil {
					return err
				}
				if err = processIssueSend(cdc, cliCtx, &coin, from, to); err != nil {
					return err
				}
				coins[i] = coin
			}
			account, err := cliCtx.GetAccount(from)
			if err != nil {
				return err
			}
			// ensure account has enough coins
			if !account.GetCoins().IsAllGTE(coins) {
				return fmt.Errorf("address %s doesn't have enough coins to pay for this transaction", from)
			}

			// build and sign the transaction, then broadcast to Tendermint
			msg := bank.NewMsgSend(from, to, coins)
			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
	cmd = client.PostCommands(cmd)[0]
	_ = cmd.MarkFlagRequired(client.FlagFrom)
	return cmd
}

func processBoxSend(cdc *codec.Codec, cliCtx context.CLIContext, coin *sdk.Coin) error {
	if !boxutils.IsId(coin.Denom) {
		return nil
	}
	boxInfo, err := boxclientutils.GetBoxByID(cdc, cliCtx, coin.Denom)
	if err != nil {
		return err
	}
	if boxInfo.IsTransferDisabled() {
		return errors.Errorf(errors.ErrCanNotTransfer(coin.Denom))
	}
	if boxInfo.GetBoxType() == types.Future {
		coin.Amount = issueutils.MulDecimals(coin.Amount, boxInfo.GetTotalAmount().Decimals)
	}
	return nil
}
func processIssueSend(cdc *codec.Codec, cliCtx context.CLIContext, coin *sdk.Coin, from sdk.AccAddress, to sdk.AccAddress) error {
	if !issueutils.IsIssueId(coin.Denom) {
		return nil
	}
	issueInfo, err := issueclientutils.GetIssueByID(cdc, cliCtx, coin.Denom)
	if err != nil {
		return err
	}
	coin.Amount = issueutils.MulDecimals(coin.Amount, issueInfo.GetDecimals())
	if err = issueclientutils.CheckFreeze(cdc, cliCtx, issueInfo.GetIssueId(), from, to); err != nil {
		return err
	}
	return nil
}
