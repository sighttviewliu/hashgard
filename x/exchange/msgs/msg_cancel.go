package msgs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/exchange/types"
)

var _ sdk.Msg = MsgCancel{}

type MsgCancel struct {
	OrderId uint64         `json:"id"`
	Seller  sdk.AccAddress `json:"seller"`
}

func NewMsgCancel(id uint64, seller sdk.AccAddress) MsgCancel {
	return MsgCancel{
		OrderId: id,
		Seller:  seller,
	}
}

// implement Msg interface
func (msg MsgCancel) Route() string {
	return types.RouterKey
}

func (msg MsgCancel) Type() string {
	return types.TypeMsgCancel
}

func (msg MsgCancel) ValidateBasic() sdk.Error {
	if msg.OrderId == 0 {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "id is invalid")
	}
	if msg.Seller.Empty() {
		return sdk.NewError(types.DefaultCodespace, types.CodeInvalidInput, "seller address is nil")
	}

	return nil
}

func (msg MsgCancel) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCancel) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Seller}
}
