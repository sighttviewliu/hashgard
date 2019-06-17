package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/types"
)

// MsgBoxWithdraw
type MsgBoxWithdraw struct {
	Id     string         `json:"id"`
	Sender sdk.AccAddress `json:"sender"`
}

//New MsgBoxWithdraw Instance
func NewMsgBoxWithdraw(boxId string, sender sdk.AccAddress) MsgBoxWithdraw {
	return MsgBoxWithdraw{boxId, sender}
}

// Route Implements Msg.
func (msg MsgBoxWithdraw) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgBoxWithdraw) Type() string { return types.TypeMsgBoxWithdraw }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgBoxWithdraw) ValidateBasic() sdk.Error {
	if len(msg.Id) == 0 {
		return errors.ErrUnknownBox("")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBoxWithdraw) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBoxWithdraw) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgBoxWithdraw) String() string {
	return fmt.Sprintf("MsgBoxWithdraw{%s}", msg.Id)
}
