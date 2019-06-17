package box

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/queriers"
	"github.com/hashgard/hashgard/x/box/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hashgard/hashgard/x/box/keeper"
)

//New Querier Instance
func NewQuerier(keeper keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryParams:
			return queriers.QueryParams(ctx, keeper)
		case types.QueryBox:
			return queriers.QueryBox(ctx, path[1], keeper)
		case types.QuerySearch:
			return queriers.QueryName(ctx, path[1], path[2], keeper)
		case types.QueryList:
			return queriers.QueryList(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown box query endpoint")
		}
	}
}
