package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/issue/utils"

	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
)

//Handle MsgIssueBurnHolder
func HandleMsgIssueBurnHolder(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgIssueBurnHolder) sdk.Result {
	_, err := keeper.BurnHolder(ctx, msg.IssueId, msg.Amount, msg.Sender)

	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.IssueId),
		Tags: utils.GetIssueTags(msg.IssueId, msg.Sender),
	}
}
