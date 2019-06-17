package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/tags"
)

func GetBoxTags(id string, boxType string, sender sdk.AccAddress) sdk.Tags {
	return sdk.NewTags(
		tags.Category, boxType,
		tags.BoxID, id,
		tags.Sender, sender.String(),
	)
}
