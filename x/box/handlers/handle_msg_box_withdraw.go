package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/tags"
	"github.com/hashgard/hashgard/x/box/types"
	"github.com/hashgard/hashgard/x/box/utils"
)

//Handle MsgBoxWithdraw
func HandleMsgBoxWithdraw(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgBoxWithdraw) sdk.Result {
	interest, boxInfo, err := keeper.ProcessBoxWithdraw(ctx, msg.Id, msg.Sender)
	if err != nil {
		return err.Result()
	}

	resTags := utils.GetBoxTags(msg.Id, boxInfo.BoxType, msg.Sender)
	if types.Deposit == boxInfo.BoxType {
		resTags = resTags.AppendTag(tags.Interest, interest.String())
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.Id),
		Tags: resTags,
	}
}
