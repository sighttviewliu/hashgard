package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Param for query record
type RecordQueryParams struct {
	StartRecordId string         `json:"start_record_id"`
	Sender        sdk.AccAddress `json:"sender"`
	Author        string 		 `json:"author"`
	Limit         int            `json:"limit"`
}
