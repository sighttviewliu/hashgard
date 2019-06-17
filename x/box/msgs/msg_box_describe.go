package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/box/types"
)

// MsgBoxDescription to allow a registered owner
// to box new coins.
type MsgBoxDescription struct {
	Id          string         `json:"id"`
	Sender      sdk.AccAddress `json:"sender"`
	Description []byte         `json:"description"`
}

//New MsgBoxDescription Instance
func NewMsgBoxDescription(boxId string, sender sdk.AccAddress, description []byte) MsgBoxDescription {
	return MsgBoxDescription{boxId, sender, description}
}

// Route Implements Msg.
func (msg MsgBoxDescription) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgBoxDescription) Type() string { return types.TypeMsgBoxDescription }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgBoxDescription) ValidateBasic() sdk.Error {
	if len(msg.Id) == 0 {
		return errors.ErrUnknownBox("")
	}
	if len(msg.Description) > types.BoxDescriptionMaxLength {
		return errors.ErrBoxDescriptionMaxLengthNotValid()
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgBoxDescription) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgBoxDescription) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgBoxDescription) String() string {
	return fmt.Sprintf("MsgBoxDescription{%s}", msg.Id)
}
