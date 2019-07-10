package record

import (
	"github.com/hashgard/hashgard/x/record/client"
	"github.com/hashgard/hashgard/x/record/keeper"
	"github.com/hashgard/hashgard/x/record/types"
)

type (
	Keeper        = keeper.Keeper
	RecordInfo = types.RecordInfo
)

var (
	NewKeeper       = keeper.NewKeeper
	NewModuleClient = client.NewModuleClient
)

const (
	StoreKey          = types.StoreKey
	RouterKey         = types.RouterKey
	QuerierRoute      = types.QuerierRoute
	DefaultParamspace = types.DefaultParamspace
	DefaultCodespace  = types.DefaultCodespace
)
