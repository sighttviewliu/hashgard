package mint

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func TestBlockProvision(t *testing.T) {
	secondsPerYear := uint64(60 * 60 * 8766)
	minter := NewMinter("foo", secondsPerYear)

	tests := []struct {
		inflation sdk.Dec
		inflationBase sdk.Int
		expProvisions int64
	} {
		{sdk.NewDecWithPrec(1, 1), sdk.NewInt(int64(secondsPerYear*100)), 10},
		{sdk.NewDecWithPrec(1, 1), sdk.NewInt(int64(secondsPerYear*200)), 20},
		{sdk.NewDecWithPrec(2, 1), sdk.NewInt(int64(secondsPerYear*100)), 20},
	}

	for i, tc := range tests {
		params := NewParams(tc.inflation, tc.inflationBase, nil)
		annualProvisions := minter.NextAnnualProvisions(params)
		provisions := minter.BlockProvision(annualProvisions)

		expProvisions := sdk.NewCoin(minter.MintDenom,
			sdk.NewInt(tc.expProvisions))

		require.True(t, expProvisions.IsEqual(provisions),
			"test: %v\n\tExp: %v\n\tGot: %v\n",
			i, tc.expProvisions, provisions)
	}
}


