package mint

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Inflate every block, update inflation parameters once per hour
func BeginBlocker(ctx sdk.Context, k Keeper) {

	// fetch stored minter & params
	minter := k.GetMinter(ctx)
	params := k.GetParams(ctx)

	// calculate annual provisions
	annualProvisions := minter.NextAnnualProvisions(params)

	// mint coins, add to collected fees, update supply
	mintedCoin := minter.BlockProvision(annualProvisions)

	foundationAddress := k.GetParams(ctx).FoundationAddress
	_, err := k.bankKeeper.SubtractCoins(ctx, foundationAddress, sdk.NewCoins(mintedCoin))
	if err != nil {
		logger := ctx.Logger().With("module", "x/mint")
		logger.Info(fmt.Sprintf("the fund of foundation address(%s) is insufficient", foundationAddress))
		panic(fmt.Errorf("the fund of foundation address(%s) is insufficient", foundationAddress))
	}
	k.fck.AddCollectedFees(ctx, sdk.Coins{mintedCoin})

	// first year do not inflate total supply
	// k.sk.InflateSupply(ctx, mintedCoin.Amount)
}
