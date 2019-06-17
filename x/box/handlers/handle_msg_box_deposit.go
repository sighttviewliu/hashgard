package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/types"

	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/utils"
)

//Handle MsgBoxInject
func HandleMsgBoxInject(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgBoxInject) sdk.Result {
	boxInfo, err := keeper.ProcessInjectBox(ctx, msg.Id, msg.Sender, msg.Amount, types.Inject)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.Id),
		Tags: utils.GetBoxTags(msg.Id, boxInfo.BoxType, msg.Sender),
	}
}

//Handle MsgBoxInject
func HandleMsgBoxInjectCancel(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgBoxInjectCancel) sdk.Result {
	boxInfo, err := keeper.ProcessInjectBox(ctx, msg.Id, msg.Sender, msg.Amount, types.Cancel)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(msg.Id),
		Tags: utils.GetBoxTags(msg.Id, boxInfo.BoxType, msg.Sender),
	}
}
