package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/types"
)

// MsgBoxInterestInject
type MsgBoxInterestInject struct {
	Id     string         `json:"id"`
	Sender sdk.AccAddress `json:"sender"`
	Amount sdk.Coin       `json:"amount"`
}

//New MsgBoxInterestInject Instance
func NewMsgBoxInterestInject(boxId string, sender sdk.AccAddress, interest sdk.Coin) MsgBoxInterestInject {
	return MsgBoxInterestInject{boxId, sender, interest}
}

// Route Implements Msg.
func (msg MsgBoxInterestInject) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgBoxInterestInject) Type() string { return types.TypeMsgBoxInterestInject }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgBoxInterestInject) ValidateBasic() sdk.Error {
	if len(msg.Id) == 0 {
		return errors.ErrUnknownBox("")
	}
	if msg.Amount.IsZero() || msg.Amount.IsNegative() {
		return errors.ErrAmountNotValid(msg.Amount.Denom)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBoxInterestInject) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBoxInterestInject) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgBoxInterestInject) String() string {
	return fmt.Sprintf("MsgBoxInterestInject{%s}", msg.Id)
}
