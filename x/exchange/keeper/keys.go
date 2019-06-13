package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	KeyDelimiter = []byte(":")

	KeyNextOrderId = []byte("newOrderId")
)

// Key for getting a specific order from the store
func KeyOrder(id uint64) []byte {
	return []byte(fmt.Sprintf("orders:%d", id))
}

// Key for getting all orders of a seller from the store
func KeyAddressOrders(addr sdk.AccAddress) []byte {
	return []byte(fmt.Sprintf("address:%d", addr))
}
