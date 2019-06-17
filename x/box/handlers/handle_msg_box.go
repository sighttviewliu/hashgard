package handlers

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/tags"
	"github.com/hashgard/hashgard/x/box/types"
)

//Handle MsgLockBox
func HandleMsgLockBox(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgLockBox) sdk.Result {
	fee := keeper.GetParams(ctx).LockCreateFee
	if err := keeper.Fee(ctx, msg.Sender, fee); err != nil {
		return err.Result()
	}

	box := &types.BoxInfo{
		Owner:            msg.Sender,
		Name:             msg.Name,
		BoxType:          types.Lock,
		TotalAmount:      msg.TotalAmount,
		Description:      msg.Description,
		TransferDisabled: true,
		Lock:             msg.Lock,
	}
	return createBox(ctx, keeper, box, fee)
}

//Handle MsgDepositBox
func HandleMsgDepositBox(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgDepositBox) sdk.Result {
	fee := keeper.GetParams(ctx).DepositBoxCreateFee
	if err := keeper.Fee(ctx, msg.Sender, fee); err != nil {
		return err.Result()
	}

	box := &types.BoxInfo{
		Owner:            msg.Sender,
		Name:             msg.Name,
		BoxType:          types.Deposit,
		TotalAmount:      msg.TotalAmount,
		Description:      msg.Description,
		TransferDisabled: msg.TransferDisabled,
		Deposit:          msg.Deposit,
	}
	return createBox(ctx, keeper, box, fee)
}

//Handle MsgFutureBox
func HandleMsgFutureBox(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgFutureBox) sdk.Result {
	fee := keeper.GetParams(ctx).FutureBoxCreateFee
	if err := keeper.Fee(ctx, msg.Sender, fee); err != nil {
		return err.Result()
	}
	box := &types.BoxInfo{
		Owner:            msg.Sender,
		Name:             msg.Name,
		BoxType:          types.Future,
		TotalAmount:      msg.TotalAmount,
		Description:      msg.Description,
		TransferDisabled: msg.TransferDisabled,
		Future:           msg.Future,
	}
	return createBox(ctx, keeper, box, fee)
}
func createBox(ctx sdk.Context, keeper keeper.Keeper, box *types.BoxInfo, fee sdk.Coin) sdk.Result {
	err := keeper.CreateBox(ctx, box)
	if err != nil {
		return err.Result()
	}
	return sdk.Result{
		Data: keeper.Getcdc().MustMarshalBinaryLengthPrefixed(box.Id),
		Tags: sdk.NewTags(
			tags.Category, box.BoxType,
			tags.BoxID, box.Id,
			tags.Sender, box.Owner.String(),
			tags.Fee, fee.String(),
		),
	}
}
