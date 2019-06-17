package mint

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Minter represents the minting state.
type Minter struct {
	MintDenom           string  `json:"mint_denom"`            // type of coin to mint
	BlocksPerYear       uint64  `json:"blocks_per_year"`       // expected blocks per year
}

// NewMinter returns a new Minter object with the given inflation and annual
// provisions values.
func NewMinter(mintDenom string, blocksPerYear uint64) Minter {
	return Minter{
		MintDenom:      mintDenom,
		BlocksPerYear:	blocksPerYear,
	}
}

// DefaultMinter returns a default initial Minter object for a new chain
func DefaultMinter() Minter {
	return NewMinter(
		"agard",
		uint64(60 * 60 * 8766 / 5),
	)
}

func validateMinter(minter Minter) error {
	if len(minter.MintDenom) == 0 {
		return fmt.Errorf("minter token denom should not be empty")
	}

	if minter.BlocksPerYear == 0 {
		return fmt.Errorf("blocks per year should not be zero")
	}

	return nil
}

// NextAnnualProvisions returns the annual provisions based on inflation base and inflation rate.
func (m Minter) NextAnnualProvisions(params Params) sdk.Dec {
	return params.Inflation.MulInt(params.InflationBase)
}

// BlockProvision returns the provisions for a block based on the annual
// provisions rate.
func (m Minter) BlockProvision(annualProvisions sdk.Dec) sdk.Coin {
	provisionAmt := annualProvisions.QuoInt(sdk.NewInt(int64(m.BlocksPerYear)))
	return sdk.NewCoin(m.MintDenom, provisionAmt.TruncateInt())
}

func (p Minter) String() string {
	return fmt.Sprintf(`Minter Info:
  Mint Denom :			%s
  Blocks Per Year :     %d
`,
		p.MintDenom, p.BlocksPerYear,
	)
}