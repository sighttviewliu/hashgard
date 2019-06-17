package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/utils"
)

//Handle MsgBoxInterestInject
func HandleMsgBoxInterestInject(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgBoxInterestInject) sdk.Result {
	boxInfo, err := keeper.InjectDepositBoxInterest(ctx, msg.Id, msg.Sender, msg.Amount)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.Id),
		Tags: utils.GetBoxTags(msg.Id, boxInfo.BoxType, msg.Sender),
	}
}

//Handle MsgBoxInterestCancel
func HandleMsgBoxInterestCancel(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgBoxInterestCancel) sdk.Result {
	boxInfo, err := keeper.CancelInterestFromDepositBox(ctx, msg.Id, msg.Sender, msg.Amount)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.Id),
		Tags: utils.GetBoxTags(msg.Id, boxInfo.BoxType, msg.Sender),
	}
}
