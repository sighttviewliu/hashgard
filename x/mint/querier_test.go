package mint

import (
	"testing"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"
)

func TestNewQuerier(t *testing.T) {
	input := newTestInput(t)
	querier := NewQuerier(input.mintKeeper)

	query := abci.RequestQuery{
		Path: "",
		Data: []byte{},
	}

	_, err := querier(input.ctx, []string{QueryParameters}, query)
	require.NoError(t, err)

	_, err = querier(input.ctx, []string{QueryMinter}, query)
	require.NoError(t, err)

	_, err = querier(input.ctx, []string{"foo"}, query)
	require.Error(t, err)
}

func TestQueryParams(t *testing.T) {
	input := newTestInput(t)

	var params Params

	res, sdkErr := queryParams(input.ctx, input.mintKeeper)
	require.NoError(t, sdkErr)

	err := input.cdc.UnmarshalJSON(res, &params)
	require.NoError(t, err)

	require.Equal(t, input.mintKeeper.GetParams(input.ctx), params)
}

func TestQueryMinter(t *testing.T) {
	input := newTestInput(t)

	var minter Minter

	res, sdkErr := queryMinter(input.ctx, input.mintKeeper)
	require.NoError(t, sdkErr)

	err := input.cdc.UnmarshalJSON(res, &minter)
	require.NoError(t, err)

	require.Equal(t, input.mintKeeper.GetMinter(input.ctx), minter)
}
