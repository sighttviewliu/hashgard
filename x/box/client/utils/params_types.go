package utils

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type LockBoxParams struct {
	LockCreateFee sdk.Coin `json:"lock_create_fee"`
	DescribeFee   sdk.Coin `json:"describe_fee"`
}
type DepositBoxParams struct {
	DepositBoxCreateFee sdk.Coin `json:"deposit_box_create_fee"`
	DisableFeatureFee   sdk.Coin `json:"disable_feature_fee"`
	DescribeFee         sdk.Coin `json:"describe_fee"`
}
type FutureBoxParams struct {
	FutureBoxCreateFee sdk.Coin `json:"future_box_create_fee"`
	DisableFeatureFee  sdk.Coin `json:"disable_feature_fee"`
	DescribeFee        sdk.Coin `json:"describe_fee"`
}

//nolint
func (dp LockBoxParams) String() string {
	return fmt.Sprintf(`Params:
  LockCreateFee:			%s
  DescribeFee:			%s`,
		dp.LockCreateFee.String(),
		dp.DescribeFee.String())
}

//nolint
func (dp DepositBoxParams) String() string {
	return fmt.Sprintf(`Params:
  DepositBoxCreateFee:			%s
  DisableFeatureFee:			%s
  DescribeFee:			%s`,
		dp.DepositBoxCreateFee.String(),
		dp.DisableFeatureFee.String(),
		dp.DescribeFee.String())
}

//nolint
func (dp FutureBoxParams) String() string {
	return fmt.Sprintf(`Params:
  FutureBoxCreateFee:			%s
  DisableFeatureFee:			%s
  DescribeFee:			%s`,
		dp.FutureBoxCreateFee.String(),
		dp.DisableFeatureFee.String(),
		dp.DescribeFee.String())
}
