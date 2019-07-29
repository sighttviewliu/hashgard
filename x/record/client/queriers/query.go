package queriers

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/hashgard/hashgard/x/record/params"
	"github.com/hashgard/hashgard/x/record/types"
)

func GetQueryRecordPath(recordHash string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryRecord, recordHash)
}
func GetQueryParamsPath() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryParams)
}
func GetQueryRecordsPath() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryRecords)
}

func QueryParams(cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryParamsPath(), nil)
}
func QueryRecord(hash string, cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryRecordPath(hash), nil)
}
func QueryRecords(params params.RecordQueryParams, cdc *codec.Codec, cliCtx context.CLIContext) ([]byte, error) {
	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	return cliCtx.QueryWithData(GetQueryRecordsPath(), bz)
}
