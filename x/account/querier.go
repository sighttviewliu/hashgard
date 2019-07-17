package account

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func QueryAddressStatus(ctx sdk.Context, addr string, keeper Keeper) ([]byte, sdk.Error) {
	res, err := keeper.GetMustMemoAddress(ctx, addr)
	if res == nil {
		return nil, sdk.ErrUnknownAddress(addr)
	}

	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), res)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}

func QueryAll(ctx sdk.Context, keeper Keeper) ([]byte, sdk.Error) {
	list := keeper.GetMustMemoAddressList(ctx)
	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), list)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}

//New Querier Instance
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, sdk.Error) {
		switch path[0] {
		case QueryMustMemoAddress:
			return QueryAddressStatus(ctx, path[1], keeper)
		case QueryMustMemoAddresses:
			return QueryAll(ctx, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown record query endpoint")
		}
	}
}
