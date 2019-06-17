package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type FutureBox struct {
	Injects         []AddressInject `json:"injects"`
	TimeLine        []int64         `json:"time"`
	Receivers       [][]string      `json:"receivers"`
	TotalWithdrawal sdk.Int         `json:"total_withdrawal"`
}

//nolint
func (bi FutureBox) String() string {
	return fmt.Sprintf(`FutureInfo:
  Injects:			%s			
  TimeLine:			%d
  Receivers:			%s
  TotalWithdrawal:			%s`,
		bi.Injects, bi.TimeLine, bi.Receivers, bi.TotalWithdrawal.String())
}
