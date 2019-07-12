package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	recordqueriers "github.com/hashgard/hashgard/x/record/client/queriers"
	"github.com/hashgard/hashgard/x/record/errors"
	"github.com/hashgard/hashgard/x/record/params"
	"github.com/hashgard/hashgard/x/record/types"
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
	var recordInfo types.Record
	cdc.MustUnmarshalJSON(res, &recordInfo)
	return cliCtx.PrintOutput(recordInfo)
	//_, err = cliCtx.Output.Write(res)
	//if err != nil {
	//	return err
	//}
	//return nil
}

// GetCmdQueryList implements the query records command.
func GetCmdQueryList(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Short:   "Query record list",
		Long:    "Query record list, flag sender is optional, default limit is 30",
		Example: "$ hashgardcli record list --sender gard1s6auwlcevspesynw44vx23e3zhuz7as9ulz56l",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			sender, err := sdk.AccAddressFromBech32(viper.GetString(flagSender))
			if err != nil {
				return err
			}
			startId := viper.GetString(flagStartRecordId)
			if len(startId) > 0 {
				if  err := recordutils.CheckRecordId(startId); err != nil {
					return errors.Errorf(err)
				}
			}
			recordQueryParams := params.RecordQueryParams{
				StartRecordId: 	startId,
				Sender:        	sender,
				Limit:        	30,
			}
			limit := viper.GetInt(flagLimit)
			if limit > 0 {
				recordQueryParams.Limit = limit
			}
			// Query the record
			res, err := recordqueriers.QueryRecords(recordQueryParams, cdc, cliCtx)
			if err != nil {
				return err
			}

			var ls types.Records
			cdc.MustUnmarshalJSON(res, &ls)
			if len(ls) == 0 {
				fmt.Println("No records")
				return nil
			}
			return cliCtx.PrintOutput(ls)
			//_, err = cliCtx.Output.Write(res)
			//if err != nil {
			//	return err
			//}
			//return nil
		},
	}

	cmd.Flags().String(flagSender, "", "sender address")
	cmd.Flags().String(flagStartRecordId, "", "Start recordId of records")
	cmd.Flags().Int32(flagLimit, 30, "Query number of record results per page returned")
	return cmd
}
