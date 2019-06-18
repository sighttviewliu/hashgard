package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/hashgard/hashgard/x/exchange/queriers"
	"github.com/hashgard/hashgard/x/exchange/types"
)

func GetQueryCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "query [id]",
		Args:    cobra.ExactArgs(1),
		Short:   "Search detailed info of an order with the specific id ",
		Example: "$ hashgardcli exchange query 1",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// validate that the proposal id is a uint
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("id %s not a valid uint, please input a valid id", args[0])
			}

			params := queriers.NewQueryOrderParams(id)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/order", types.QuerierRoute), bz)
			if err != nil {
				return err
			}

			var order types.Order
			cdc.MustUnmarshalJSON(res, &order)
			return cliCtx.PrintOutput(order)
		},
	}
}

func GetListCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "list [address]",
		Args:    cobra.ExactArgs(1),
		Short:   "Search all orders of a specific address",
		Example: "$ hashgardcli exchange list gard1hf4n743fujvxrwx8af7u35anjqpdd2cx8p6cdd",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			seller, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			params := queriers.NewQueryOrdersParams(seller)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/orders", types.QuerierRoute), bz)
			if err != nil {
				return err
			}

			var orders types.Orders
			cdc.MustUnmarshalJSON(res, &orders)
			return cliCtx.PrintOutput(orders)
		},
	}
}

func GetFrozenFundCmd(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "query-frozen [address]",
		Args:    cobra.ExactArgs(1),
		Short:   "Search frozen order list of a specific address ",
		Example: "$ hashgardcli exchange query-frozen gard1hf4n743fujvxrwx8af7u35anjqpdd2cx8p6cdd",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			seller, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			params := queriers.NewQueryFrozenFundParams(seller)
			bz, err := cdc.MarshalJSON(params)
			if err != nil {
				return err
			}

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/frozen", types.QuerierRoute), bz)
			if err != nil {
				return err
			}

			var coins sdk.Coins
			cdc.MustUnmarshalJSON(res, &coins)
			return cliCtx.PrintOutput(coins)
		},
	}
}
