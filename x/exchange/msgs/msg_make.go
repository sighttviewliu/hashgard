package msgs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/exchange/types"
)

var _ sdk.Msg = MsgMake{}

type MsgMake struct {
	Seller sdk.AccAddress `json:"seller"`
	Supply sdk.Coin       `json:"supply"`
	Target sdk.Coin       `json:"target"`
}

func NewMsgMake(seller sdk.AccAddress, supply sdk.Coin, target sdk.Coin) MsgMake {
	return MsgMake{
		Seller: seller,
		Supply: supply,
		Target: target,
	}
}

// implement Msg interface
func (msg MsgMake) Route() string {
	return types.RouterKey
}

func (msg MsgMake) Type() string {
	return types.TypeMsgMake
}

func (msg MsgMake) ValidateBasic() sdk.Error {
	if msg.Seller.Empty() {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "seller address is nil")
	}
	if msg.Supply.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "supply amount is invalid: "+msg.Supply.String())
	}
	if msg.Target.Amount.LTE(sdk.ZeroInt()) {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "target amount is invalid: "+msg.Target.String())
	}
	if msg.Target.Denom == msg.Supply.Denom {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "target coin can't be same with supply coin")
	}

	return nil
}

func (msg MsgMake) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgMake) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Seller}
}
