package msgs

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/hashgard/hashgard/x/record/types"
)

var MsgCdc = codec.New()

// Register concrete types on codec codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgRecord{}, "record/MsgRecord", nil)

	cdc.RegisterInterface((*types.Record)(nil), nil)
	cdc.RegisterConcrete(&types.RecordInfo{}, "record/RecordInfo", nil)
}

//nolint
func init() {
	RegisterCodec(MsgCdc)
}
