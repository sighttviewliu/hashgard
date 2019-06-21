package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/hashgard/hashgard/x/box/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type LockBoxInfo struct {
	Id               string         `json:"id"`
	BoxType          string         `json:"type"`
	Status           string         `json:"status"`
	Owner            sdk.AccAddress `json:"owner"`
	Name             string         `json:"name"`
	CreatedTime      int64          `json:"created_time"`
	TotalAmount      types.BoxToken `json:"total_amount"`
	Description      string         `json:"description"`
	TransferDisabled bool           `json:"transfer_disabled"`
	Lock             types.LockBox  `json:"lock"`
}
type DepositBoxInfo struct {
	Id               string           `json:"id"`
	BoxType          string           `json:"type"`
	Status           string           `json:"status"`
	Owner            sdk.AccAddress   `json:"owner"`
	Name             string           `json:"name"`
	CreatedTime      int64            `json:"created_time"`
	TotalAmount      types.BoxToken   `json:"total_amount"`
	Description      string           `json:"description"`
	TransferDisabled bool             `json:"transfer_disabled"`
	Deposit          types.DepositBox `json:"deposit"`
}
type FutureBoxInfo struct {
	Id               string          `json:"id"`
	BoxType          string          `json:"type"`
	Status           string          `json:"status"`
	Owner            sdk.AccAddress  `json:"owner"`
	Name             string          `json:"name"`
	CreatedTime      int64           `json:"created_time"`
	TotalAmount      types.BoxToken  `json:"total_amount"`
	Description      string          `json:"description"`
	TransferDisabled bool            `json:"transfer_disabled"`
	Future           types.FutureBox `json:"future"`
}
type LockBoxInfos []LockBoxInfo
type DepositBoxInfos []DepositBoxInfo
type FutureBoxInfos []FutureBoxInfo

//nolint
func getString(Id string, Status string, Owner sdk.AccAddress, Name string, CreatedTime int64,
	TotalAmount types.BoxToken, Description string, TransferDisabled bool) string {
	return fmt.Sprintf(`BoxInfo:
  Id:				%s
  Status:			%s
  Owner:			%s
  Name:				%s
  TotalAmount:			%s
  CreatedTime:			%d
  Description:			%s
  TransferDisabled:		%t`,
		Id, Status, Owner.String(), Name, TotalAmount.String(),
		CreatedTime, Description, TransferDisabled)
}

//nolint
func (bi LockBoxInfo) String() string {
	str := getString(bi.Id, bi.Status, bi.Owner, bi.Name,
		bi.CreatedTime, bi.TotalAmount, bi.Description, bi.TransferDisabled)

	return fmt.Sprintf(`%s
%s`, str, bi.Lock.String())
}

//nolint
func (bi DepositBoxInfo) String() string {
	str := getString(bi.Id, bi.Status, bi.Owner, bi.Name,
		bi.CreatedTime, bi.TotalAmount, bi.Description, bi.TransferDisabled)

	return fmt.Sprintf(`%s
%s`, str, bi.Deposit.String())
}

//nolint
func (bi FutureBoxInfo) String() string {
	str := getString(bi.Id, bi.Status, bi.Owner, bi.Name,
		bi.CreatedTime, bi.TotalAmount, bi.Description, bi.TransferDisabled)

	return fmt.Sprintf(`%s
%s`, str, bi.Future.String())
}

//nolint
func (bi LockBoxInfos) String() string {
	out := fmt.Sprintf("%-17s|%-44s|%-16s|%-40s|%s\n",
		"BoxID", "Owner", "Name", "TotalAmount", "EndTime")
	for _, box := range bi {
		out += fmt.Sprintf("%-17s|%-44s|%-16s|%-40s|%s\n",
			box.Id, box.Owner.String(), box.Name, box.TotalAmount.Token.String(), time.Unix(box.Lock.EndTime, 0).String())
	}
	return strings.TrimSpace(out)
}

//nolint
func (bi DepositBoxInfos) String() string {
	out := fmt.Sprintf("%-17s|%-44s|%-16s|%-40s|%s\n",
		"BoxID", "Owner", "Name", "TotalAmount", "CreatedTime")
	for _, box := range bi {
		out += fmt.Sprintf("%-17s|%-44s|%-16s|%-40s|%d\n",
			box.Id, box.Owner.String(), box.Name, box.TotalAmount.Token.String(), box.CreatedTime)
	}
	return strings.TrimSpace(out)
}

//nolint
func (bi FutureBoxInfos) String() string {
	out := fmt.Sprintf("%-17s|%-44s|%-16s|%-40s|%s\n",
		"BoxID", "Owner", "Name", "TotalAmount", "CreatedTime")
	for _, box := range bi {
		out += fmt.Sprintf("%-17s|%-44s|%-16s|%-40s|%d\n",
			box.Id, box.Owner.String(), box.Name, box.TotalAmount.Token.String(), box.CreatedTime)
	}
	return strings.TrimSpace(out)
}
