package params

import (
	"github.com/hashgard/hashgard/x/box/types"
)

type BoxFutureParams struct {
	Name             string          `json:"name"`
	TotalAmount      types.BoxToken  `json:"total_amount"`
	Description      string          `json:"description"`
	TransferDisabled bool            `json:"transfer_disabled"`
	Future           types.FutureBox `json:"future"`
}
