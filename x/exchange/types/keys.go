package types

const (
	// ModuleName is the name of the exchange module
	ModuleName = "exchange"

	// StoreKey is the string store representation
	StoreKey = ModuleName

	// RouteKey is the msg router key for the exchange module
	RouterKey = ModuleName

	// QuerierRoute is the querier route for the staking module
	QuerierRoute = ModuleName

	// Parameter store default namestore
	DefaultParamspace = ModuleName
)
const (
	TypeMsgMake   = "make"
	TypeMsgTake   = "take"
	TypeMsgCancel = "cancel"
)
const (
	FinishedStatus = "finished"
)

const (
	DefaultStartingOrderId uint64 = 1
)
