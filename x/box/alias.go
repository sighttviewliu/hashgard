package box

import (
	"github.com/hashgard/hashgard/x/box/client/cli"
	"github.com/hashgard/hashgard/x/box/config"
	"github.com/hashgard/hashgard/x/box/keeper"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/types"
)

type (
	Keeper  = keeper.Keeper
	BoxInfo = types.BoxInfo
	Params  = config.Params
	Hooks   = keeper.Hooks
)

var (
	MsgCdc        = msgs.MsgCdc
	NewKeeper     = keeper.NewKeeper
	RegisterCodec = msgs.RegisterCodec
	SendTxCmd     = cli.SendTxCmd
	QueryCmd      = cli.QueryCmd
	WithdrawCmd   = cli.WithdrawCmd
	DefaultParams = config.DefaultParams
)

const (
	StoreKey          = types.StoreKey
	RouterKey         = types.RouterKey
	QuerierRoute      = types.QuerierRoute
	DefaultParamspace = types.DefaultParamspace
	DefaultCodespace  = types.DefaultCodespace
)
