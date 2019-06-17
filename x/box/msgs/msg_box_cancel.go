package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/types"
)

// MsgBoxInjectCancel
type MsgBoxInjectCancel struct {
	Id     string         `json:"id"`
	Sender sdk.AccAddress `json:"sender"`
	Amount sdk.Coin       `json:"amount"`
}

//New MsgBoxInjectCancel Instance
func NewMsgBoxInjectCancel(boxId string, sender sdk.AccAddress, amount sdk.Coin) MsgBoxInjectCancel {
	return MsgBoxInjectCancel{boxId, sender, amount}
}

// Route Implements Msg.
func (msg MsgBoxInjectCancel) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgBoxInjectCancel) Type() string { return types.TypeMsgBoxCancel }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgBoxInjectCancel) ValidateBasic() sdk.Error {
	if len(msg.Id) == 0 {
		return errors.ErrUnknownBox("")
	}
	if msg.Amount.IsZero() || msg.Amount.IsNegative() {
		return errors.ErrAmountNotValid(msg.Amount.Denom)
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBoxInjectCancel) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBoxInjectCancel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgBoxInjectCancel) String() string {
	return fmt.Sprintf("MsgBoxInjectCancel{%s}", msg.Id)
}
