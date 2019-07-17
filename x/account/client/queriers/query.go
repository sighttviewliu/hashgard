package queriers

import (
	"fmt"
	"github.com/hashgard/hashgard/x/account"

	"github.com/cosmos/cosmos-sdk/client/context"
)

func GetQueryMustMemoAddressPath(addr string) string {
	return fmt.Sprintf("%s/%s/%s/%s", account.Custom, account.QuerierRoute, account.QueryMustMemoAddress, addr)
}
func GetQueryMustMemoAddressListPath() string {
	return fmt.Sprintf("%s/%s/%s", account.Custom, account.QuerierRoute, account.QueryMustMemoAddresses)
}

func QueryMustMemoAddress(addr string, cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryMustMemoAddressPath(addr), nil)
}
func QueryMustMemoAddressList(cliCtx context.CLIContext) ([]byte, error) {
	return cliCtx.QueryWithData(GetQueryMustMemoAddressListPath(), nil)
}
