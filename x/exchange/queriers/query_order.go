package queriers

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/hashgard/hashgard/x/exchange/keeper"
)

type QueryOrderParams struct {
	OrderId uint64
}

func NewQueryOrderParams(id uint64) QueryOrderParams {
	return QueryOrderParams{
		OrderId: id,
	}
}

func QueryOrder(ctx sdk.Context, cdc *codec.Codec, req abci.RequestQuery, keeper keeper.Keeper) ([]byte, sdk.Error) {
	var params QueryOrderParams
	err := cdc.UnmarshalJSON(req.Data, &params)
	if err != nil {
		return []byte{}, sdk.ErrInternal(fmt.Sprintf("failed to parse params: %s", err))
	}

	order, orderErr := keeper.GetOrder(ctx, params.OrderId)
	if orderErr != nil {
		return nil, orderErr
	}

	bz, err := codec.MarshalJSONIndent(cdc, order)
	if err != nil {
		return nil, sdk.ErrInternal(fmt.Sprintf("could not marshal result to JSON: %s", err))
	}

	return bz, nil
}
