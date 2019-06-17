package queriers

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/hashgard/hashgard/x/box/params"
	"github.com/hashgard/hashgard/x/box/types"
)

func GetQueryBoxPath(id string) string {
	return fmt.Sprintf("%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryBox, id)
}
func GetQueryBoxParamsPath() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryParams)
}
func GetQueryBoxSearchPath(boxType string, name string) string {
	return fmt.Sprintf("%s/%s/%s/%s/%s", types.Custom, types.QuerierRoute, types.QuerySearch, boxType, name)
}
func GetQueryBoxsPath() string {
	return fmt.Sprintf("%s/%s/%s", types.Custom, types.QuerierRoute, types.QueryList)
}
func QueryBoxParams(cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryBoxParamsPath(), nil)
}
func QueryBoxByName(boxType string, name string, cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryBoxSearchPath(boxType, name), nil)
}

func QueryBoxByID(id string, cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryBoxPath(id), nil)
}

func QueryBoxsList(params params.BoxQueryParams, cdc *codec.Codec, cliCtx context.CLIContext) ([]byte, error) {
	bz, err := cdc.MarshalJSON(params)
	if err != nil {
		return nil, err
	}
	return cliCtx.QueryWithData(GetQueryBoxsPath(), bz)
}
