package issue

import (
	"github.com/hashgard/hashgard/x/issue/client"
	"github.com/hashgard/hashgard/x/issue/client/cli"
	"github.com/hashgard/hashgard/x/issue/config"
	"github.com/hashgard/hashgard/x/issue/keeper"
	"github.com/hashgard/hashgard/x/issue/msgs"
	"github.com/hashgard/hashgard/x/issue/types"
)

type (
	Keeper        = keeper.Keeper
	CoinIssueInfo = types.CoinIssueInfo
	Approval      = types.Approval
	IssueFreeze   = types.IssueFreeze
	Params        = config.Params
	Hooks         = keeper.Hooks
)

var (
	MsgCdc          = msgs.MsgCdc
	NewKeeper       = keeper.NewKeeper
	NewModuleClient = client.NewModuleClient
	//GetAccountCmd   = cli.GetAccountCmd
	QueryCmd      = cli.QueryCmd
	RegisterCodec = msgs.RegisterCodec
	DefaultParams = config.DefaultParams
)

const (
	StoreKey          = types.StoreKey
	RouterKey         = types.RouterKey
	QuerierRoute      = types.QuerierRoute
	DefaultParamspace = types.DefaultParamspace
	DefaultCodespace  = types.DefaultCodespace
)
