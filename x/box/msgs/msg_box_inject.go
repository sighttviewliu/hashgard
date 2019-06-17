package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/types"
)

// MsgBoxInject
type MsgBoxInject struct {
	Id     string         `json:"id"`
	Sender sdk.AccAddress `json:"sender"`
	Amount sdk.Coin       `json:"amount"`
}

//New MsgBoxInject Instance
func NewMsgBoxInject(boxId string, sender sdk.AccAddress, amount sdk.Coin) MsgBoxInject {
	return MsgBoxInject{boxId, sender, amount}
}

// Route Implements Msg.
func (msg MsgBoxInject) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgBoxInject) Type() string { return types.TypeMsgBoxInject }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgBoxInject) ValidateBasic() sdk.Error {
	if len(msg.Id) == 0 {
		return errors.ErrUnknownBox("")
	}
	if msg.Amount.IsZero() || msg.Amount.IsNegative() {
		return errors.ErrAmountNotValid(msg.Amount.Denom)
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBoxInject) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBoxInject) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgBoxInject) String() string {
	return fmt.Sprintf("MsgBoxInject{%s}", msg.Id)
}
