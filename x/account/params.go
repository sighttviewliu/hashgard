package account

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Param for query record
type QueryParams struct {
	Sender        sdk.AccAddress `json:"sender"`
}
