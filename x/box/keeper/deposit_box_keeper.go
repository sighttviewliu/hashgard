package keeper

import (
	"github.com/hashgard/hashgard/x/box/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"
)

//Process deposit box

func (keeper Keeper) ProcessDepositBoxCreate(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	decimal, err := keeper.GetCoinDecimals(ctx, box.Deposit.Interest.Token)
	if err != nil {
		return err
	}
	if box.Deposit.Interest.Decimals != decimal {
		return errors.ErrDecimalsNotValid(box.Deposit.Interest.Decimals)
	}
	box.Status = types.BoxCreated
	box.Deposit.TotalInject = sdk.ZeroInt()
	box.Deposit.WithdrawalInterest = sdk.ZeroInt()
	box.Deposit.WithdrawalShare = sdk.ZeroInt()
	//box.Deposit.TotalInterestInject = sdk.ZeroInt()
	box.Deposit.Share = sdk.ZeroInt()
	keeper.InsertActiveBoxQueue(ctx, box.Deposit.StartTime, box.Id)
	return nil
}

func (keeper Keeper) InjectDepositBoxInterest(ctx sdk.Context, id string, sender sdk.AccAddress, interest sdk.Coin) (*types.BoxInfo, sdk.Error) {
	box := keeper.GetBox(ctx, id)
	if box == nil {
		return nil, errors.ErrUnknownBox(id)
	}
	if types.BoxCreated != box.Status {
		return nil, errors.ErrNotAllowedOperation(box.Status)
	}
	if box.Deposit.Interest.Token.Denom != interest.Denom {
		return nil, errors.ErrInterestInjectNotValid(interest)
	}
	if len(box.Deposit.InterestInjects) >= types.BoxMaxInjectInterest {
		return nil, errors.ErrInterestInjectNotValid(interest)
	}
	totalInterest := sdk.ZeroInt()
	if box.Deposit.InterestInjects == nil {
		box.Deposit.InterestInjects = []types.AddressInject{{Address: sender, Amount: interest.Amount}}
	} else {
		exist := false
		for i, v := range box.Deposit.InterestInjects {
			totalInterest = totalInterest.Add(v.Amount)
			if v.Address.Equals(sender) {
				box.Deposit.InterestInjects[i].Amount = box.Deposit.InterestInjects[i].Amount.Add(interest.Amount)
				exist = true
			}
		}
		if !exist {
			box.Deposit.InterestInjects = append(box.Deposit.InterestInjects, types.NewAddressInject(sender, interest.Amount))
		}
	}
	totalInterest = totalInterest.Add(interest.Amount)
	if totalInterest.GT(box.Deposit.Interest.Token.Amount) {
		return nil, errors.ErrInterestInjectNotValid(interest)
	}
	if err := keeper.SendDepositedCoin(ctx, sender, sdk.Coins{interest}, id); err != nil {
		return nil, err
	}
	keeper.setBox(ctx, box)
	return box, nil
}
func (keeper Keeper) CancelInterestFromDepositBox(ctx sdk.Context, id string, sender sdk.AccAddress, interest sdk.Coin) (*types.BoxInfo, sdk.Error) {
	box := keeper.GetBox(ctx, id)
	if box == nil {
		return nil, errors.ErrUnknownBox(id)
	}
	if types.BoxCreated != box.Status {
		return nil, errors.ErrNotAllowedOperation(box.Status)
	}
	if box.Deposit.Interest.Token.Denom != interest.Denom {
		return nil, errors.ErrInterestInjectNotValid(interest)
	}
	if box.Deposit.InterestInjects == nil {
		return nil, errors.ErrInterestCancelNotValid(interest)
	} else {
		remove := -1
		for i, v := range box.Deposit.InterestInjects {
			if v.Address.Equals(sender) {
				if box.Deposit.InterestInjects[i].Amount.LT(interest.Amount) {
					return nil, errors.ErrNotEnoughAmount()
				}
				box.Deposit.InterestInjects[i].Amount = box.Deposit.InterestInjects[i].Amount.Sub(interest.Amount)
				if box.Deposit.InterestInjects[i].Amount.IsZero() {
					remove = i
				}
			}
		}
		if remove != -1 {
			box.Deposit.InterestInjects = append(box.Deposit.InterestInjects[:remove], box.Deposit.InterestInjects[remove+1:]...)
		}
	}
	if err := keeper.CancelDepositedCoin(ctx, sender, sdk.Coins{interest}, id); err != nil {
		return nil, err
	}
	keeper.setBox(ctx, box)
	return box, nil
}
func (keeper Keeper) processDepositBoxInject(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, amount sdk.Coin, operation string) sdk.Error {
	switch operation {
	case types.Inject:
		return keeper.injectDepositBox(ctx, box, sender, amount)
	case types.Cancel:
		return keeper.cancelDepositFromDepositBox(ctx, box, sender, amount)
	}
	return errors.ErrUnknownOperation()
}
func (keeper Keeper) injectDepositBox(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, amount sdk.Coin) sdk.Error {
	if !amount.Amount.Mod(box.Deposit.Price).IsZero() {
		return errors.ErrAmountNotValid(amount.Denom)
	}
	if box.TotalAmount.Token.Denom != amount.Denom {
		return errors.ErrAmountNotValid(amount.Denom)
	}
	box.Deposit.TotalInject = box.Deposit.TotalInject.Add(amount.Amount)
	if box.Deposit.TotalInject.GT(box.TotalAmount.Token.Amount) {
		return errors.ErrAmountNotValid(amount.Denom)
	}
	if err := keeper.SendDepositedCoin(ctx, sender, sdk.Coins{amount}, box.Id); err != nil {
		return err
	}
	share := amount.Amount.Quo(box.Deposit.Price)
	_, err := keeper.ck.AddCoins(ctx, sender, sdk.NewCoins(sdk.NewCoin(box.Id, share)))
	if err != nil {
		return err
	}
	box.Deposit.Share = box.Deposit.Share.Add(share)
	keeper.setBox(ctx, box)
	return nil
}
func (keeper Keeper) cancelDepositFromDepositBox(ctx sdk.Context, box *types.BoxInfo, sender sdk.AccAddress, amount sdk.Coin) sdk.Error {
	if !amount.Amount.Mod(box.Deposit.Price).IsZero() {
		return errors.ErrAmountNotValid(amount.Amount.String())
	}
	if box.TotalAmount.Token.Denom != amount.Denom {
		return errors.ErrAmountNotValid(amount.Denom)
	}
	share := amount.Amount.Quo(box.Deposit.Price)
	_, err := keeper.GetBankKeeper().SubtractCoins(ctx, sender, sdk.NewCoins(sdk.NewCoin(box.Id, share)))
	if err != nil {
		return err
	}
	if err := keeper.CancelDepositedCoin(ctx, sender, sdk.NewCoins(amount), box.Id); err != nil {
		return err
	}
	box.Deposit.Share = box.Deposit.Share.Sub(share)
	box.Deposit.TotalInject = box.Deposit.TotalInject.Sub(amount.Amount)
	//if types.BoxClosed == box.Status && box.Deposit.Share.IsZero() {
	//	keeper.RemoveBox(ctx, box)
	//} else {
	//	keeper.setBox(ctx, box)
	//}
	keeper.setBox(ctx, box)
	return nil
}
func (keeper Keeper) ProcessDepositBoxByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	switch box.Status {
	case types.BoxCreated:
		return keeper.processBoxCreatedByEndBlocker(ctx, box)
	case types.BoxInjecting:
		return keeper.processDepositBoxInjectByEndBlocker(ctx, box)
	case types.DepositBoxInterest:
		return keeper.processDepositBoxInterestByEndBlocker(ctx, box)
	default:
		return errors.ErrNotAllowedOperation(box.Status)
	}
}
func (keeper Keeper) backBoxInterestInjects(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	for _, v := range box.Deposit.InterestInjects {
		keeper.CancelDepositedCoin(ctx, v.Address,
			sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, v.Amount)), box.Id)
	}
	return nil
}
func (keeper Keeper) backBoxUnUsedInterestInjects(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	totalCoupon := box.TotalAmount.Token.Amount.Quo(box.Deposit.Price)
	if totalCoupon.Equal(box.Deposit.Share) {
		return nil
	}
	unused := utils.CalcInterest(box.Deposit.PerCoupon, totalCoupon.Sub(box.Deposit.Share), box.Deposit.Interest)
	interestInjectsLen := len(box.Deposit.InterestInjects)
	if interestInjectsLen == 0 {
		if err := keeper.CancelDepositedCoin(ctx, box.Deposit.InterestInjects[0].Address,
			sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, unused)), box.Id); err != nil {
			return err
		}
		box.Deposit.InterestInjects[0].Amount = box.Deposit.InterestInjects[0].Amount.Sub(unused)
	} else {
		total := sdk.ZeroInt()
		for i, v := range box.Deposit.InterestInjects {
			var amount sdk.Int
			if i == interestInjectsLen-1 {
				amount = unused.Sub(total)
			} else {
				amount = sdk.NewDecFromInt(v.Amount).QuoInt(box.Deposit.Interest.Token.Amount).MulInt(unused).TruncateInt()
			}
			if err := keeper.CancelDepositedCoin(ctx, v.Address,
				sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, amount)), box.Id); err != nil {
				return err
			}
			box.Deposit.InterestInjects[i].Amount = box.Deposit.InterestInjects[i].Amount.Sub(amount)
			total = total.Add(amount)
		}
	}
	box.Deposit.Interest.Token.Amount = box.Deposit.Interest.Token.Amount.Sub(unused)
	return nil
}
func (keeper Keeper) backBoxAllDeposit(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	for _, v := range box.Deposit.InterestInjects {
		if err := keeper.CancelDepositedCoin(ctx, v.Address,
			sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, v.Amount)), box.Id); err != nil {
			return err
		}
	}
	return nil
}
func (keeper Keeper) processBoxCreatedByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.Status != types.BoxCreated {
		return nil
	}
	totalInterest := sdk.ZeroInt()
	if box.Deposit.InterestInjects != nil {
		for _, v := range box.Deposit.InterestInjects {
			totalInterest = totalInterest.Add(v.Amount)
		}
	}
	keeper.RemoveFromActiveBoxQueue(ctx, box.Deposit.StartTime, box.Id)
	if box.Deposit.Interest.Token.Amount.Equal(totalInterest) {
		box.Status = types.BoxInjecting
		keeper.InsertActiveBoxQueue(ctx, box.Deposit.EstablishTime, box.Id)
		keeper.setBox(ctx, box)
	} else {
		box.Status = types.BoxClosed
		if err := keeper.backBoxInterestInjects(ctx, box); err != nil {
			return err
		}
	}
	keeper.setBox(ctx, box)
	return nil
}
func (keeper Keeper) processDepositBoxInjectByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.Status != types.BoxInjecting {
		return nil
	}

	keeper.RemoveFromActiveBoxQueue(ctx, box.Deposit.EstablishTime, box.Id)

	box.Status = types.BoxClosed

	if box.Deposit.TotalInject.IsZero() || box.Deposit.TotalInject.LT(box.Deposit.BottomLine) {
		keeper.backBoxInterestInjects(ctx, box)
	} else {
		keeper.backBoxUnUsedInterestInjects(ctx, box)
		box.Status = types.DepositBoxInterest
		keeper.InsertActiveBoxQueue(ctx, box.Deposit.MaturityTime, box.Id)
	}
	keeper.setBox(ctx, box)
	return nil
}
func (keeper Keeper) processDepositBoxWithdraw(ctx sdk.Context, id string, sender sdk.AccAddress) (sdk.Int, *types.BoxInfo, sdk.Error) {
	box := keeper.GetBox(ctx, id)
	if box == nil {
		return sdk.ZeroInt(), nil, errors.ErrUnknownBox(id)
	}
	if types.Deposit != box.BoxType {
		return sdk.ZeroInt(), nil, errors.ErrNotSupportOperation()
	}
	if types.BoxFinished != box.Status {
		return sdk.ZeroInt(), nil, errors.ErrNotAllowedOperation(box.Status)
	}
	share := keeper.GetBankKeeper().GetCoins(ctx, sender).AmountOf(id)
	box.Deposit.WithdrawalShare = share.Add(box.Deposit.WithdrawalShare)
	_, err := keeper.GetBankKeeper().SubtractCoins(ctx, sender, sdk.NewCoins(sdk.NewCoin(id, share)))
	if err != nil {
		return sdk.ZeroInt(), nil, err
	}

	amount := share.Mul(box.Deposit.Price)
	if err := keeper.CancelDepositedCoin(ctx, sender, sdk.NewCoins(sdk.NewCoin(box.TotalAmount.Token.Denom, amount)), box.Id); err != nil {
		return sdk.ZeroInt(), nil, err
	}

	interest := sdk.ZeroInt()
	if box.Deposit.WithdrawalShare == box.Deposit.Share {
		totalInterest := sdk.ZeroInt()
		for _, v := range box.Deposit.InterestInjects {
			totalInterest = totalInterest.Add(v.Amount)
		}
		interest = totalInterest.Sub(box.Deposit.WithdrawalInterest)
	} else {
		interest = utils.CalcInterest(box.Deposit.PerCoupon, share, box.Deposit.Interest)
	}

	if err = keeper.CancelDepositedCoin(ctx, sender,
		sdk.NewCoins(sdk.NewCoin(box.Deposit.Interest.Token.Denom, interest)), box.Id); err != nil {
		return sdk.ZeroInt(), nil, err
	}

	box.Deposit.WithdrawalInterest = interest.Add(box.Deposit.WithdrawalInterest)
	keeper.setBox(ctx, box)
	return interest, box, nil

}
func (keeper Keeper) processDepositBoxInterestByEndBlocker(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	if box.Status != types.DepositBoxInterest {
		return nil
	}
	keeper.RemoveFromActiveBoxQueue(ctx, box.Deposit.MaturityTime, box.Id)
	box.Status = types.BoxFinished
	keeper.setBox(ctx, box)
	return nil
}
