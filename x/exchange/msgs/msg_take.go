package msgs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/exchange/types"
)

var _ sdk.Msg = MsgTake{}

type MsgTake struct {
	OrderId uint64         `json:"id"`
	Buyer   sdk.AccAddress `json:"buyer"`
	Value   sdk.Coin       `json:"value"`
}

func NewMsgTake(id uint64, buyer sdk.AccAddress, val sdk.Coin) MsgTake {
	return MsgTake{
		OrderId: id,
		Buyer:   buyer,
		Value:   val,
	}
}

// implement Msg interface
func (msg MsgTake) Route() string {
	return types.RouterKey
}

func (msg MsgTake) Type() string {
	return types.TypeMsgTake
}

func (msg MsgTake) ValidateBasic() sdk.Error {
	if msg.OrderId == 0 {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "id is invalid")
	}
	if msg.Value.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "value is invalid: "+msg.Value.String())
	}

	return nil
}

func (msg MsgTake) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgTake) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Buyer}
}
