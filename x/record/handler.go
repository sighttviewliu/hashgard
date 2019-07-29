package record

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/record/handlers"
	"github.com/hashgard/hashgard/x/record/keeper"
	"github.com/hashgard/hashgard/x/record/msgs"
)

// Handle all "record" type messages.
func NewHandler(keeper keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case msgs.MsgRecord:
			return handlers.HandleMsgRecord(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized record msg type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}
