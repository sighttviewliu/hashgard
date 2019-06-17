package msgs

import (
	"github.com/cosmos/cosmos-sdk/codec"
	boxclienttype "github.com/hashgard/hashgard/x/box/client/types"
	"github.com/hashgard/hashgard/x/box/types"
)

var MsgCdc = codec.New()

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgLockBox{}, "box/MsgLockBox", nil)
	cdc.RegisterConcrete(MsgDepositBox{}, "box/MsgDepositBox", nil)
	cdc.RegisterConcrete(MsgFutureBox{}, "box/MsgFutureBox", nil)
	cdc.RegisterConcrete(MsgBoxInterestInject{}, "box/MsgBoxInterestInject", nil)
	cdc.RegisterConcrete(MsgBoxInterestCancel{}, "box/MsgBoxInterestCancel", nil)
	cdc.RegisterConcrete(MsgBoxInject{}, "box/MsgBoxInject", nil)
	cdc.RegisterConcrete(MsgBoxInjectCancel{}, "box/MsgBoxInjectCancel", nil)
	cdc.RegisterConcrete(MsgBoxWithdraw{}, "box/MsgBoxWithdraw", nil)
	cdc.RegisterConcrete(MsgBoxDescription{}, "box/MsgBoxDescription", nil)
	cdc.RegisterConcrete(MsgBoxDisableFeature{}, "box/MsgBoxDisableFeature", nil)

	cdc.RegisterConcrete(boxclienttype.LockBoxInfo{}, "box/LockBoxInfo", nil)
	cdc.RegisterConcrete(boxclienttype.DepositBoxInfo{}, "box/DepositBoxInfo", nil)
	cdc.RegisterConcrete(boxclienttype.FutureBoxInfo{}, "box/FutureBoxInfo", nil)

	cdc.RegisterInterface((*types.Box)(nil), nil)
	cdc.RegisterConcrete(&types.BoxInfo{}, "box/BoxInfo", nil)
}

//nolint
func init() {
	RegisterCodec(MsgCdc)
}
