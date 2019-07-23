package queriers

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/record/errors"
	"github.com/hashgard/hashgard/x/record/keeper"
	"github.com/hashgard/hashgard/x/record/params"
	abci "github.com/tendermint/tendermint/abci/types"
)

func QueryRecord(ctx sdk.Context, hash string, keeper keeper.Keeper) ([]byte, sdk.Error) {
	record := keeper.GetRecord(ctx, hash)
	if record == nil {
		return nil, errors.ErrUnknownRecord(hash)
	}

	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), record)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}

func QueryRecords(ctx sdk.Context, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, sdk.Error) {
	var params params.RecordQueryParams
	err := keeper.Getcdc().UnmarshalJSON(req.Data, &params)
	if err != nil {
		return nil, sdk.ErrUnknownRequest(sdk.AppendMsgToErr("incorrectly formatted request data", err.Error()))
	}
	records := keeper.List(ctx, params)
	bz, err := codec.MarshalJSONIndent(keeper.Getcdc(), records)
	if err != nil {
		return nil, sdk.ErrInternal(sdk.AppendMsgToErr("could not marshal result to JSON", err.Error()))
	}
	return bz, nil
}
