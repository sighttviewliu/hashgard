package account

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	// ModuleKey is the name of the module
	ModuleName = "account"
	// StoreKey is the store key string for record
	StoreKey = ModuleName
	// RouterKey is the message route for record
	RouterKey = ModuleName
	// QuerierRoute is the querier route for record
	QuerierRoute = ModuleName
	// Parameter store default namestore
	DefaultParamspace = ModuleName

	TypeMsgSetMustMemo = ModuleName
)

const (
	DefaultCodespace sdk.CodespaceType = ModuleName
)

const (
	Custom   				= "custom"
	QueryMustMemoAddresses  = "list"
	QueryMustMemoAddress    = "query"
)