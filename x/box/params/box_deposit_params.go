package params

import (
	"github.com/hashgard/hashgard/x/box/types"
)

type BoxDepositParams struct {
	Name             string           `json:"name"`
	TotalAmount      types.BoxToken   `json:"total_amount"`
	Description      string           `json:"description"`
	TransferDisabled bool             `json:"transfer_disabled"`
	Deposit          types.DepositBox `json:"deposit"`
}
