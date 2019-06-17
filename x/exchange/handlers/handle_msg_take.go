package handlers

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
	"github.com/hashgard/hashgard/x/exchange/msgs"
	"github.com/hashgard/hashgard/x/exchange/tags"
	"github.com/hashgard/hashgard/x/exchange/types"
)

func HandleMsgTake(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgTake) sdk.Result {
	supplyTurnover, targetTurnover, soldOut, err := keeper.Take(ctx, msg.OrderId, msg.Buyer, msg.Value)
	if err != nil {
		return err.Result()
	}
	resTags := sdk.NewTags(
		tags.Category, tags.TxCategory,
		tags.OrderId, fmt.Sprintf("%d", msg.OrderId),
		tags.Sender, msg.Buyer.String(),
		tags.SupplyTurnover, supplyTurnover.String(),
		tags.TargetTurnover, targetTurnover.String(),
	)
	if soldOut {
		resTags = resTags.AppendTag(tags.OrderStatus, types.FinishedStatus)
	}
	return sdk.Result{
		Tags: resTags,
	}
}
