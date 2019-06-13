package tags

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Exchange tags
var (
	TxCategory = "exchange"

	Action         = sdk.TagAction
	Category       = sdk.TagCategory
	Sender         = sdk.TagSender
	OrderId        = "id"
	OrderStatus    = "status"
	SupplyDenom    = "supply_denom"
	TargetDenom    = "target_denom"
	SupplyAmount   = "supply_amount"
	TargetAmount   = "target_amount"
	SupplyTurnover = "supply_turnover"
	TargetTurnover = "target_turnover"
)
