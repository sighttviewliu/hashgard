package app

import (
	"fmt"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func (app *HashgardApp) assertRuntimeInvariants() {
	ctx := app.NewContext(false, abci.Header{Height: app.LastBlockHeight() + 1})
	app.assertRuntimeInvariantsOnContext(ctx)
}

func (app *HashgardApp) assertRuntimeInvariantsOnContext(ctx sdk.Context) {
	start := time.Now()
	invarRoutes := app.crisisKeeper.Routes()
	for _, ir := range invarRoutes {
		if err := ir.Invar(ctx); err != nil {
			panic(fmt.Errorf("invariant broken: %s\n"+
				"\tCRITICAL please submit the following transaction:\n"+
				"\t\t hashgardcli crisis invariant-broken %v %v", err, ir.ModuleName, ir.Route))
		}
	}
	end := time.Now()
	diff := end.Sub(start)
	app.BaseApp.Logger().With("module", "invariants").Info(
		"Asserted all invariants", "duration", diff, "height", app.LastBlockHeight())
}
