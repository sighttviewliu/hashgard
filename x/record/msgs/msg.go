package msgs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type MsgFlag interface {
	sdk.Msg

	GetRecordId() string
	SetRecordId(string)

	GetSender() sdk.AccAddress
	SetSender(sdk.AccAddress)
}
