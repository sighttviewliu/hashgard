package tags

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Box tags
var (
	TxCategory = "box"

	Action    = sdk.TagAction
	Category  = sdk.TagCategory
	Sender    = sdk.TagSender
	Feature   = "feature"
	Fee       = "fee"
	Owner     = "owner"
	Operation = "operation"
	Interest  = "interest"
	BoxID     = "id"
	Status    = "status"
	Seq       = "seq"
)
