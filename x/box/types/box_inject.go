package types

//
//import (
//	"fmt"
//
//	sdk "github.com/cosmos/cosmos-sdk/types"
//)
//
//type BoxInject struct {
//	Amount   sdk.Int `json:"amount"`
//	Interest sdk.Int `json:"interest"`
//}
//
//func NewZeroBoxInject() *BoxInject {
//	return &BoxInject{Amount: sdk.ZeroInt(), Interest: sdk.ZeroInt()}
//}
//func NewBoxInject(amount sdk.Int) *BoxInject {
//	return &BoxInject{Amount: amount, Interest: sdk.ZeroInt()}
//}
//
////nolint
//func (bi BoxInject) String() string {
//	return fmt.Sprintf(`
//  Amount:			%s
//  Interest:			%s`,
//		bi.Amount.String(), bi.Interest.String())
//}
