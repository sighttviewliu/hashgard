package cli

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/spf13/cobra"

	"github.com/hashgard/hashgard/x/mint"
)

// GetCmdQueryParams implements a command to return the current minting parameter(inflation and inflation base)
func GetCmdQueryParams(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "params",
		Short: "Query the current minting parameters",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", mint.QuerierRoute, mint.QueryParameters)
			res, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var params mint.Params
			if err := cdc.UnmarshalJSON(res, &params); err != nil {
				return err
			}

			return cliCtx.PrintOutput(params)
		},
	}
}

// GetCmdQueryMinter implements a command to return the minter info(denom and blocks per year)
func GetCmdQueryMinter(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "minter",
		Short: "Query the current minter info",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			route := fmt.Sprintf("custom/%s/%s", mint.QuerierRoute, mint.QueryMinter)
			res, err := cliCtx.QueryWithData(route, nil)
			if err != nil {
				return err
			}

			var minter mint.Minter
			if err := cdc.UnmarshalJSON(res, &minter); err != nil {
				return err
			}

			return cliCtx.PrintOutput(minter)
		},
	}
}
