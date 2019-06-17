package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DepositBox struct {
	StartTime          int64           `json:"start_time"`
	EstablishTime      int64           `json:"establish_time"`
	MaturityTime       int64           `json:"maturity_time"`
	BottomLine         sdk.Int         `json:"bottom_line"`
	Interest           BoxToken        `json:"interest"`
	Price              sdk.Int         `json:"price"`
	PerCoupon          sdk.Dec         `json:"per_coupon"`
	Share              sdk.Int         `json:"share"`
	TotalInject        sdk.Int         `json:"total_inject"`
	WithdrawalShare    sdk.Int         `json:"withdrawal_share"`
	WithdrawalInterest sdk.Int         `json:"withdrawal_interest"`
	InterestInjects    []AddressInject `json:"interest_injects"`
}

type DepositBoxInjectInterest struct {
	Address  sdk.AccAddress `json:"address"`
	Amount   sdk.Int        `json:"amount"`
	Interest sdk.Int        `json:"interest"`
}

type DepositBoxInjectInterestList []DepositBoxInjectInterest

//nolint
func (bi DepositBoxInjectInterest) String() string {
	return fmt.Sprintf(`
  Address:			%s
  Amount:			%s
  Interest:			%s`,
		bi.Address.String(), bi.Amount.String(), bi.Interest.String())
}

//nolint
func (bi DepositBox) String() string {
	return fmt.Sprintf(`DepositInfo:
  StartTime:			%d
  EstablishTime:		%d
  MaturityTime:			%d
  BottomLine:			%s
  Interest:			%s
  Price:			%s
  PerCoupon:			%s
  Share:			%s
  TotalInject:			%s
  WithdrawalShare:			%s,
  WithdrawalInterest:			%s,
  InterestInject:			%s`,
		bi.StartTime,
		bi.EstablishTime,
		bi.MaturityTime,
		bi.BottomLine.String(),
		bi.Interest.String(),
		bi.Price.String(),
		bi.PerCoupon.String(),
		bi.Share.String(),
		bi.TotalInject.String(),
		bi.WithdrawalShare.String(),
		bi.WithdrawalInterest.String(),
		bi.InterestInjects)
}

//nolint
func (bi DepositBoxInjectInterestList) String() string {
	out := fmt.Sprintf("%-44s|%-40s|%s\n",
		"Address", "Amount", "Interest")
	for _, box := range bi {
		out += fmt.Sprintf("%-44s|%-40s|%s\n",
			box.Address.String(), box.Amount.String(), box.Interest.String())
	}
	return strings.TrimSpace(out)
}
