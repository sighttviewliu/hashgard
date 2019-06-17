package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/types"
)

// MsgBoxInterestCancel
type MsgBoxInterestCancel struct {
	Id     string         `json:"id"`
	Sender sdk.AccAddress `json:"sender"`
	Amount sdk.Coin       `json:"amount"`
}

//New MsgBoxInterestCancel Instance
func NewMsgBoxInterestCancel(boxId string, sender sdk.AccAddress, interest sdk.Coin) MsgBoxInterestCancel {
	return MsgBoxInterestCancel{boxId, sender, interest}
}

// Route Implements Msg.
func (msg MsgBoxInterestCancel) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgBoxInterestCancel) Type() string { return types.TypeMsgBoxInterestCancel }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgBoxInterestCancel) ValidateBasic() sdk.Error {
	if len(msg.Id) == 0 {
		return errors.ErrUnknownBox("")
	}
	if msg.Amount.IsZero() || msg.Amount.IsNegative() {
		return errors.ErrAmountNotValid(msg.Amount.Denom)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBoxInterestCancel) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBoxInterestCancel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgBoxInterestCancel) String() string {
	return fmt.Sprintf("MsgBoxInterestCancel{%s}", msg.Id)
}
