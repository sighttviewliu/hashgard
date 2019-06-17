package types

import (
	"fmt"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Order struct {
	OrderId    uint64         `json:"id"`
	Seller     sdk.AccAddress `json:"seller"`
	Supply     sdk.Coin       `json:"supply"`
	Target     sdk.Coin       `json:"target"`
	Remains    sdk.Coin       `json:"remains"`
	CreateTime time.Time      `json:"create_time"`
}

func (order Order) String() string {
	return fmt.Sprintf(`Order %d:
		  Seller:			%s
		  Supply:			%s
		  Target:			%s
		  Remains:			%s
		  Create Time:		%s`, order.OrderId, order.Seller, order.Supply,
		order.Target, order.Remains, order.CreateTime)
}

// Orders is an array of order
type Orders []*Order

func (orders Orders) String() string {
	out := fmt.Sprintf("%10s - (%15s) - (%40s) - [%10s] - Create Time\n", "ID", "Supply", "Target", "Remains")
	for _, order := range orders {
		out += fmt.Sprintf("%10d - (%15s) - (%40s) - [%10s] - %s\n",
			order.OrderId, order.Supply, order.Target, order.Remains, order.CreateTime)
	}

	return strings.TrimSpace(out)
}
