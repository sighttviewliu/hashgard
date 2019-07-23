package record

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hashgard/hashgard/x/record/keeper"
	"github.com/hashgard/hashgard/x/record/queriers"
	"github.com/hashgard/hashgard/x/record/types"
)

//New Querier Instance
func NewQuerier(keeper keeper.Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case types.QueryRecord:
			return queriers.QueryRecord(ctx, path[1], keeper)
		case types.QueryRecords:
			return queriers.QueryRecords(ctx, req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown record query endpoint")
		}
	}
}
