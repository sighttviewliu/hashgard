package handlers

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
	"github.com/hashgard/hashgard/x/exchange/msgs"
	"github.com/hashgard/hashgard/x/exchange/tags"
)

func HandleMsgCancel(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgCancel) sdk.Result {
	_, err := keeper.Cancel(ctx, msg.OrderId, msg.Seller)
	if err != nil {
		return err.Result()
	}
	resTags := sdk.NewTags(
		tags.Category, tags.TxCategory,
		tags.OrderId, fmt.Sprintf("%d", msg.OrderId),
		tags.Sender, msg.Seller.String(),
	)
	return sdk.Result{
		Tags: resTags,
	}
}
