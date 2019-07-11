package tags

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Record tags
var (
	TxCategory = "record"

	Action          = sdk.TagAction
	Category        = sdk.TagCategory
	Sender          = sdk.TagSender
	ID         		= "id"
	Name            = "name"
	Hash            = "hash"
	Author          = "author"
	RecordType      = "record-type"
	RecordNo        = "record-number"
)
