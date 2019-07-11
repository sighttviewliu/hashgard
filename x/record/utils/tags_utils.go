package utils

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/record/types"

	"github.com/hashgard/hashgard/x/record/tags"
)

func GetRecordTags(info *types.RecordInfo) sdk.Tags {
	res := sdk.NewTags(
		tags.Category, tags.TxCategory,
		tags.Sender, info.Sender.String(),
		tags.ID, info.ID,
		tags.Hash, info.Hash,
	)
	if len(info.RecordType) > 0 {
		res.AppendTag(tags.RecordType, info.RecordType)
	}
	if len(info.RecordNo) > 0 {
		res.AppendTag(tags.RecordNo, info.RecordNo)
	}
	if len(info.Author) > 0 {
		res.AppendTag(tags.Author, info.Author)
	}
	return res
}
