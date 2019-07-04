package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/record/tags"
)

func GetRecordTags(id string, hash string, sender sdk.AccAddress) sdk.Tags {
	return sdk.NewTags(
		tags.Category, tags.TxCategory,
		tags.Sender, sender.String(),
		tags.ID, id,
		tags.Hash, hash,
	)

}
