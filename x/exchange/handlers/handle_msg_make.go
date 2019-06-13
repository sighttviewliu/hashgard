package handlers

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
	"github.com/hashgard/hashgard/x/exchange/msgs"
	"github.com/hashgard/hashgard/x/exchange/tags"
)

func HandleMsgMake(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgMake) sdk.Result {
	order, err := keeper.Make(ctx, msg.Seller, msg.Supply, msg.Target)
	if err != nil {
		return err.Result()
	}
	resTags := sdk.NewTags(
		tags.Category, tags.TxCategory,
		tags.OrderId, fmt.Sprintf("%d", order.OrderId),
		tags.Sender, order.Seller.String(),
		tags.SupplyDenom, order.Supply.Denom,
		tags.SupplyAmount, order.Supply.Amount.String(),
		tags.TargetDenom, order.Target.Denom,
		tags.TargetAmount, order.Target.Amount.String(),
	)
	return sdk.Result{
		Tags: resTags,
	}
}
