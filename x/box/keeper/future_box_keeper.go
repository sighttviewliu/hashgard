package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"
	"github.com/hashgard/hashgard/x/box/utils"
)

//Process Future box

func (keeper Keeper) ProcessFutureBoxCreate(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	box.Status = types.BoxInjecting
	box.Future.TotalWithdrawal = sdk.ZeroInt()
	//keeper.InsertActiveBoxQueue(ctx, box.Future.TimeLine[0], keeper.GetFutureBoxSeqString(box, 0))
	keeper.InsertActiveBoxQueue(ctx, box.Future.TimeLine[0], box.Id)
	return nil
}
func (keeper Keeper) processFutureBoxInject(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, amount sdk.Coin, operation string) sdk.Error {
	switch operation {
	case types.Inject:
		return keeper.injectFutureBox(ctx, box, sender, amount)
	case types.Cancel:
		return keeper.cancelDepositFromFutureBox(ctx, box, sender, amount)
	default:
		return errors.ErrUnknownOperation()
	}
}
func (keeper Keeper) injectFutureBox(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, amount sdk.Coin) sdk.Error {
	if box.Future.TimeLine[0] < ctx.BlockHeader().Time.Unix() {
		return errors.ErrNotSupportOperation()
	}
	if box.TotalAmount.Token.Denom != amount.Denom {
		return errors.ErrAmountNotValid(amount.Denom)
	}
	totalDeposit := sdk.ZeroInt()
	if box.Future.Injects == nil {
		box.Future.Injects = []types.AddressInject{{Address: sender, Amount: amount.Amount}}
	} else {
		exist := false
		for i, v := range box.Future.Injects {
			totalDeposit = totalDeposit.Add(v.Amount)
			if v.Address.Equals(sender) {
				box.Future.Injects[i].Amount = box.Future.Injects[i].Amount.Add(amount.Amount)
				exist = true
			}
		}
		if !exist {
			box.Future.Injects = append(box.Future.Injects, types.NewAddressInject(sender, amount.Amount))
		}
	}
	totalDeposit = totalDeposit.Add(amount.Amount)
	if totalDeposit.GT(box.TotalAmount.Token.Amount) {
		return errors.ErrNotEnoughAmount()
	}
	if err := keeper.SendDepositedCoin(ctx, sender, sdk.Coins{amount}, box.Id); err != nil {
		return err
	}
	if totalDeposit.Equal(box.TotalAmount.Token.Amount) {
		if err := keeper.processFutureBoxDistribute(ctx, box); err != nil {
			return err
		}
		box.Status = types.BoxActived
		//keeper.RemoveFromActiveBoxQueue(ctx, box.Future.TimeLine[0], keeper.GetFutureBoxSeqString(box, 0))
		keeper.RemoveFromActiveBoxQueue(ctx, box.Future.TimeLine[0], box.Id)
	}
	keeper.setBox(ctx, box)
	return nil
}
func (keeper Keeper) cancelDepositFromFutureBox(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, amount sdk.Coin) sdk.Error {
	if box.Status == types.BoxActived {
		return errors.ErrNotAllowedOperation(box.Status)
	}
	if box.TotalAmount.Token.Denom != amount.Denom {
		return errors.ErrAmountNotValid(amount.Denom)
	}
	if box.Future.Injects == nil {
		return errors.ErrNotEnoughAmount()
	}
	exist := false
	for i, v := range box.Future.Injects {
		if v.Address.Equals(sender) {
			if box.Future.Injects[i].Amount.LT(amount.Amount) {
				return errors.ErrNotEnoughAmount()
			}
			box.Future.Injects[i].Amount = box.Future.Injects[i].Amount.Sub(amount.Amount)
			exist = true
			break
		}
	}
	if !exist {
		return errors.ErrNotEnoughAmount()
	}

	if err := keeper.CancelDepositedCoin(ctx, sender, sdk.NewCoins(amount), box.Id); err != nil {
		return err
	}
	keeper.setBox(ctx, box)
	return nil
}
func (keeper Keeper) processFutureBoxDistribute(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	var address sdk.AccAddress
	var total = sdk.ZeroInt()
	for _, items := range box.Future.Receivers {
		for j, rec := range items {
			if j == 0 {
				addr, err := sdk.AccAddressFromBech32(rec)
				if err != nil {
					return sdk.ErrInvalidAddress(rec)
				}
				address = addr
				continue
			}
			amount, ok := sdk.NewIntFromString(rec)
			if !ok {
				return errors.ErrAmountNotValid(rec)
			}
			boxDenom := utils.GetCoinDenomByFutureBoxSeq(box.Id, j)
			_, err := keeper.GetBankKeeper().AddCoins(ctx, address, sdk.NewCoins(sdk.NewCoin(boxDenom, amount)))
			if err != nil {
				return err
			}
			total = total.Add(amount)
		}
	}
	if !total.Equal(box.TotalAmount.Token.Amount) {
		return errors.ErrAmountNotValid("Receivers")
	}
	//times := len(box.Future.TimeLine)
	//keeper.InsertActiveBoxQueue(ctx, box.Future.TimeLine[times-1], keeper.GetFutureBoxSeqString(box, times))
	keeper.InsertActiveBoxQueue(ctx, box.Future.TimeLine[len(box.Future.TimeLine)-1], box.Id)
	return nil
}

//func (keeper Keeper) GetFutureBoxSeqString(box *types.BoxInfo, seq int) string {
//	return fmt.Sprintf("%s:%d", box.Id, seq)
//}

func (keeper Keeper) ProcessFutureBoxByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	switch box.Status {
	case types.BoxInjecting:
		return keeper.processFutureBoxInjectByEndBlocker(ctx, box)
	case types.BoxActived:
		return keeper.processFutureBoxActiveByEndBlocker(ctx, box)
	default:
		return errors.ErrNotAllowedOperation(box.Status)
	}
}
func (keeper Keeper) processFutureBoxInjectByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if types.BoxInjecting != box.Status {
		return errors.ErrNotAllowedOperation(box.Status)
	}
	if box.Future.Injects != nil {
		for _, v := range box.Future.Injects {
			keeper.CancelDepositedCoin(ctx, v.Address, sdk.NewCoins(sdk.NewCoin(box.TotalAmount.Token.Denom, v.Amount)), box.Id)
		}
	}
	box.Status = types.BoxClosed
	//keeper.RemoveFromActiveBoxQueue(ctx, box.Future.TimeLine[0], keeper.GetFutureBoxSeqString(box, seq))
	keeper.RemoveFromActiveBoxQueue(ctx, box.Future.TimeLine[0], box.Id)
	//keeper.RemoveBox(ctx, box)
	keeper.setBox(ctx, box)
	return nil
}

func (keeper Keeper) processFutureBoxWithdraw(ctx sdk.Context, idSeq string, sender sdk.AccAddress) (*types.BoxInfo, sdk.Error) {
	box := keeper.GetBox(ctx, idSeq)
	if box == nil {
		return nil, errors.ErrUnknownBox(idSeq)
	}
	if types.Future != box.BoxType {
		return nil, errors.ErrNotSupportOperation()
	}
	if types.BoxCreated == box.Status {
		return nil, errors.ErrNotAllowedOperation(box.Status)
	}
	seq := utils.GetSeqFromFutureBoxSeq(idSeq)
	if box.Future.TimeLine[seq-1] > ctx.BlockHeader().Time.Unix() {
		return nil, errors.ErrNotAllowedOperation(types.BoxUndue)
	}
	amount := keeper.GetBankKeeper().GetCoins(ctx, sender).AmountOf(idSeq)
	_, err := keeper.GetBankKeeper().SubtractCoins(ctx, sender, sdk.NewCoins(sdk.NewCoin(idSeq, amount)))
	if err != nil {
		return nil, err
	}
	if err := keeper.CancelDepositedCoin(ctx, sender, sdk.NewCoins(sdk.NewCoin(box.TotalAmount.Token.Denom, amount)), box.Id); err != nil {
		return nil, err
	}
	box.Future.TotalWithdrawal = amount.Add(box.Future.TotalWithdrawal)
	keeper.setBox(ctx, box)
	return box, nil
}

func (keeper Keeper) processFutureBoxActiveByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if types.BoxActived != box.Status {
		return errors.ErrNotAllowedOperation(box.Status)
	}
	box.Status = types.BoxFinished
	//keeper.RemoveFromActiveBoxQueue(ctx, box.Future.TimeLine[seq-1], keeper.GetFutureBoxSeqString(box, seq))
	keeper.RemoveFromActiveBoxQueue(ctx, box.Future.TimeLine[len(box.Future.TimeLine)-1], box.Id)
	keeper.setBox(ctx, box)
	return nil
}
