package tags

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// tags
var (
	TxCategory 		= "account"

	Action          = sdk.TagAction
	Category        = sdk.TagCategory
	Sender          = sdk.TagSender
	MustMemo    	= "must-memo"
)
