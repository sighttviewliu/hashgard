package account

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgSetMustMemo{}, "account/MsgSetMustMemo", nil)
}

var MsgCdc = codec.New()

func init() {
	RegisterCodec(MsgCdc)
}

// Msg
type MsgSetMustMemo struct {
	Sender   sdk.AccAddress `json:"sender"`
	MustMemo bool           `json:"must_memo"`
}

var _ sdk.Msg = MsgSetMustMemo{}

// NewMsg - construct msg.
func NewMsgSetMustMemo(sender sdk.AccAddress, mustMemo bool) MsgSetMustMemo {
	return MsgSetMustMemo{Sender: sender, MustMemo: mustMemo}
}

// Route Implements Msg.
func (msg MsgSetMustMemo) Route() string { return RouterKey }

// Type Implements Msg.
func (msg MsgSetMustMemo) Type() string { return TypeMsgSetMustMemo }

// ValidateBasic Implements Msg.
func (msg MsgSetMustMemo) ValidateBasic() sdk.Error {
	if msg.Sender.Empty() {
		return sdk.ErrInvalidAddress("missing sender address")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSetMustMemo) GetSignBytes() []byte {
	return sdk.MustSortJSON(MsgCdc.MustMarshalJSON(msg))
}

// GetSigners Implements Msg.
func (msg MsgSetMustMemo) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}
