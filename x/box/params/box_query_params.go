package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Param query box
type BoxQueryParams struct {
	StartId string         `json:"start_id"`
	Owner   sdk.AccAddress `json:"owner"`
	BoxType string         `json:"type"`
	Limit   int            `json:"limit"`
}

// Param query deposit
type BoxQueryDepositListParams struct {
	Id    string         `json:"id"`
	Owner sdk.AccAddress `json:"owner"`
	//Limit int            `json:"limit"`
}
