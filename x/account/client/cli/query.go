package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/account"
	"github.com/hashgard/hashgard/x/account/client/queriers"
	"github.com/spf13/cobra"
)

// GetCmdQueryRecord implements the query record command.
func GetCmdQueryMustMemo(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "query [address]",
		Args:    cobra.ExactArgs(1),
		Short:   "Query must-memo status of an address",
		Long:    "Query must-memo status of an address",
		Example: "$ hashgardcli must-memo query gard16e36z0k53wvefmtel0ghazjy23gt5lpg758ndu",
		RunE: func(cmd *cobra.Command, args []string) error {
			return processQuery(cdc, args)
		},
	}
}

func processQuery(cdc *codec.Codec, args []string) error {
	cliCtx := context.NewCLIContext().WithCodec(cdc)

	addr, err := sdk.AccAddressFromBech32(args[0])
	if err != nil {
		return err
	}
	// Query the record
	res, err := queriers.QueryMustMemoAddress(addr.String(), cliCtx)
	if err != nil {
		return err
	}
	var acc account.AccountInfo
	cdc.MustUnmarshalJSON(res, &acc)
	return cliCtx.PrintOutput(acc)
}

func GetCmdQueryMustMemoList(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Query must-memo address list",
		Long:    "Query must-memo address list",
		Example: "$ hashgardcli must-memo list",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			// Query the record
			res, err := queriers.QueryMustMemoAddressList(cliCtx)
			if err != nil {
				return err
			}

			var ls account.Accounts
			cdc.MustUnmarshalJSON(res, &ls)
			if len(ls) == 0 {
				fmt.Println("No records")
				return nil
			}
			return cliCtx.PrintOutput(ls)
		},
	}
	return cmd
}
