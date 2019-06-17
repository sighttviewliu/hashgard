package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/tags"
	"github.com/hashgard/hashgard/x/issue/utils"
)

//Handle MsgIssueDisableFeature
func HandleMsgIssueDisableFeature(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueDisableFeature) sdk.Result {
	if err := keeper.DisableFeature(ctx, msg.Sender, msg.IssueId, msg.Feature); err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.GetIssueTags(msg.IssueId, msg.Sender).AppendTag(tags.Feature, msg.GetFeature()),
	}
}
