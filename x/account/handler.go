package account

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/account/tags"
)

//Handle Msg
func HandleMsgSetMustMemo(ctx sdk.Context, keeper Keeper, msg MsgSetMustMemo) sdk.Result {
	acc := AccountInfo{
		Sender:             msg.Sender,
		MustMemo:           msg.MustMemo,
	}

	keeper.SetMustMemoAddress(ctx, &acc)

	return sdk.Result{
		Tags: sdk.NewTags(
			tags.Category, tags.TxCategory,
			tags.Sender, acc.Sender.String(),
			tags.MustMemo, strconv.FormatBool(acc.MustMemo),
		),
	}
}

// Handle all "account" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgSetMustMemo:
			return HandleMsgSetMustMemo(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized account msg type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
