package cli

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	recordqueriers "github.com/hashgard/hashgard/x/record/client/queriers"
	"github.com/hashgard/hashgard/x/record/errors"
	"github.com/hashgard/hashgard/x/record/params"
	recordutils "github.com/hashgard/hashgard/x/record/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GetCmdQueryRecord implements the query record command.
func GetCmdQueryRecord(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "query [hash]",
		Args:    cobra.ExactArgs(1),
		Short:   "Query a single record",
		Long:    "Query detail of a record by record hash",
		Example: "$ hashgardcli record query BC38CAEE32149BEF4CCFAEAB518EC9A5FBC85AE6AC8D5A9F6CD710FAF5E4A2B8",
		RunE: func(cmd *cobra.Command, args []string) error {
			return processQuery(cdc, args)
		},
	}
}

func processQuery(cdc *codec.Codec, args []string) error {
	cliCtx := context.NewCLIContext().WithCodec(cdc)
	hash := args[0]
	if err := recordutils.CheckRecordHash(hash); err != nil {
		return errors.Errorf(err)
	}
	// Query the record
	res, err := recordqueriers.QueryRecord(hash, cliCtx)
	if err != nil {
		return err
	}
	//var recordInfo types.Record
	//cdc.MustUnmarshalJSON(res, &recordInfo)
	//return cliCtx.PrintOutput(recordInfo)
	_, err = cliCtx.Output.Write(res)
	if err != nil {
		return err
	}
	return nil
}

// GetCmdQueryList implements the query records command.
func GetCmdQueryList(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Query record list by sender and author",
		Long:    "Query record list by sender and author, flag sender or author must be provided, default limit is 30",
		Example: "$ hashgardcli record list --sender gard1s6auwlcevspesynw44vx23e3zhuz7as9ulz56l --author authorName",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			sender, err := sdk.AccAddressFromBech32(viper.GetString(flagSender))
			if err != nil {
				return err
			}
			recordQueryParams := params.RecordQueryParams{
				StartRecordId: 	viper.GetString(flagStartRecordId),
				Sender:        	sender,
				Author:        	viper.GetString(flagAuthor),
				Limit:        	viper.GetInt(flagLimit),
			}
			// Query the record
			res, err := recordqueriers.QueryRecords(recordQueryParams, cdc, cliCtx)
			if err != nil {
				return err
			}

			//var tokenRecords types.CoinRecords
			//cdc.MustUnmarshalJSON(res, &tokenRecords)
			//if len(tokenRecords) == 0 {
			//	fmt.Println("No records")
			//	return nil
			//}
			//return cliCtx.PrintOutput(tokenRecords)
			_, err = cliCtx.Output.Write(res)
			if err != nil {
				return err
			}
			return nil
		},
	}

	cmd.Flags().String(flagSender, "", "sender address")
	cmd.Flags().String(flagAuthor, "", "auther name")
	cmd.Flags().String(flagStartRecordId, "", "Start recordId of records")
	cmd.Flags().Int32(flagLimit, 30, "Query number of record results per page returned")
	return cmd
}
