package msgs

import (
	"fmt"
	"time"

	"github.com/hashgard/hashgard/x/box/utils"

	"github.com/hashgard/hashgard/x/box/params"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"
)

// MsgDepositBox to allow a registered boxr
// to box new coins.
type MsgDepositBox struct {
	Sender                   sdk.AccAddress `json:"sender"`
	*params.BoxDepositParams `json:"params"`
}

func NewMsgDepositBox(sender sdk.AccAddress, params *params.BoxDepositParams) MsgDepositBox {
	return MsgDepositBox{sender, params}
}

// Route Implements Msg.
func (msg MsgDepositBox) Route() string { return types.RouterKey }

// Type Implements Msg.789
func (msg MsgDepositBox) Type() string { return types.TypeMsgBoxCreate }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgDepositBox) ValidateBasic() sdk.Error {
	if len(msg.Sender) == 0 {
		return sdk.ErrInvalidAddress("Sender address cannot be empty")
	}
	if msg.TotalAmount.Token.IsZero() || msg.TotalAmount.Token.Amount.IsNegative() {
		return errors.ErrAmountNotValid("Token amount")
	}
	if len(msg.Name) > types.BoxNameMaxLength {
		return errors.ErrBoxNameNotValid()
	}
	if len(msg.Description) > types.BoxDescriptionMaxLength {
		return errors.ErrBoxDescriptionMaxLengthNotValid()
	}
	return nil
}
func (msg MsgDepositBox) ValidateService() sdk.Error {
	if err := msg.ValidateBasic(); err != nil {
		return err
	}
	zero := sdk.ZeroInt()
	if msg.Deposit.StartTime <= time.Now().Unix() {
		return errors.ErrTimeNotValid("StartTime")
	}
	if msg.Deposit.EstablishTime <= msg.Deposit.StartTime {
		return errors.ErrTimeNotValid("EstablishTime")
	}
	if msg.Deposit.MaturityTime <= msg.Deposit.EstablishTime {
		return errors.ErrTimeNotValid("MaturityTime")
	}
	if msg.Deposit.BottomLine.LT(zero) || msg.Deposit.BottomLine.GT(msg.TotalAmount.Token.Amount) {
		return errors.ErrAmountNotValid("BottomLine")
	}
	if msg.Deposit.Interest.Token.Amount.LT(zero) {
		return errors.ErrAmountNotValid("Interest")
	}
	if msg.Deposit.Price.LTE(zero) || !msg.TotalAmount.Token.Amount.Mod(msg.Deposit.Price).IsZero() {
		return errors.ErrAmountNotValid("Price")
	}
	if !msg.Deposit.PerCoupon.Equal(utils.CalcInterestRate(msg.TotalAmount.Token.Amount, msg.Deposit.Price,
		msg.Deposit.Interest.Token, msg.Deposit.Interest.Decimals)) {
		return errors.ErrAmountNotValid("PerCoupon")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgDepositBox) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgDepositBox) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgDepositBox) String() string {
	return fmt.Sprintf("MsgDepositBox{%s - %s}", "", msg.Sender.String())
}
