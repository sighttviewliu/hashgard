package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

//Box interface
type Box interface {
	GetId() string
	SetId(string)

	GetBoxType() string
	SetBoxType(string)

	GetStatus() string
	SetStatus(string)

	GetOwner() sdk.AccAddress
	SetOwner(sdk.AccAddress)

	GetCreatedTime() int64
	SetCreatedTime(int64)

	GetName() string
	SetName(string)

	GetTotalAmount() BoxToken
	SetTotalAmount(BoxToken)

	GetDescription() string
	SetDescription(string)

	IsTransferDisabled() bool
	SetTransferDisabled(bool)

	GetLock() LockBox
	SetLock(LockBox)

	GetDeposit() DepositBox
	SetDeposit(DepositBox)

	GetFuture() FutureBox
	SetFuture(FutureBox)

	String() string
}

// BoxInfos is an array of BoxInfo
type BoxInfos []BoxInfo

//type BaseBoxInfo struct {
//}
type BoxInfo struct {
	Id               string         `json:"id"`
	Status           string         `json:"status"`
	Owner            sdk.AccAddress `json:"owner"`
	Name             string         `json:"name"`
	BoxType          string         `json:"type"`
	CreatedTime      int64          `json:"created_time"`
	TotalAmount      BoxToken       `json:"total_amount"`
	Description      string         `json:"description"`
	TransferDisabled bool           `json:"transfer_disabled"`
	Lock             LockBox        `json:"lock"`
	Deposit          DepositBox     `json:"deposit"`
	Future           FutureBox      `json:"future"`
}

// Implements Box Interface
var _ Box = (*BoxInfo)(nil)

func (bi BoxInfo) GetId() string {
	return bi.Id
}
func (bi *BoxInfo) SetId(boxId string) {
	bi.Id = boxId
}
func (bi BoxInfo) GetBoxType() string {
	return bi.BoxType
}
func (bi *BoxInfo) SetBoxType(boxType string) {
	bi.BoxType = boxType
}
func (bi BoxInfo) GetStatus() string {
	return bi.Status
}
func (bi *BoxInfo) SetStatus(boxStatus string) {
	bi.Status = boxStatus
}
func (bi BoxInfo) GetOwner() sdk.AccAddress {
	return bi.Owner
}
func (bi *BoxInfo) SetOwner(owner sdk.AccAddress) {
	bi.Owner = owner
}
func (bi BoxInfo) GetCreatedTime() int64 {
	return bi.CreatedTime
}
func (bi *BoxInfo) SetCreatedTime(createdTime int64) {
	bi.CreatedTime = createdTime
}
func (bi BoxInfo) GetName() string {
	return bi.Name
}
func (bi *BoxInfo) SetName(name string) {
	bi.Name = name
}
func (bi BoxInfo) GetTotalAmount() BoxToken {
	return bi.TotalAmount
}
func (bi *BoxInfo) SetTotalAmount(totalAmount BoxToken) {
	bi.TotalAmount = totalAmount
}
func (bi BoxInfo) GetDescription() string {
	return bi.Description
}
func (bi *BoxInfo) SetDescription(description string) {
	bi.Description = description
}

func (bi BoxInfo) IsTransferDisabled() bool {
	return bi.TransferDisabled
}

func (bi *BoxInfo) SetTransferDisabled(transferDisabled bool) {
	bi.TransferDisabled = transferDisabled
}

func (bi BoxInfo) GetLock() LockBox {
	return bi.Lock
}
func (bi *BoxInfo) SetLock(lock LockBox) {
	bi.Lock = lock
}

func (bi BoxInfo) GetDeposit() DepositBox {
	return bi.Deposit
}
func (bi *BoxInfo) SetDeposit(deposit DepositBox) {
	bi.Deposit = deposit
}

func (bi BoxInfo) GetFuture() FutureBox {
	return bi.Future
}
func (bi *BoxInfo) SetFuture(future FutureBox) {
	bi.Future = future
}

type AddressInject struct {
	Address sdk.AccAddress `json:"address"`
	Amount  sdk.Int        `json:"amount"`
}

func NewAddressInject(address sdk.AccAddress, amount sdk.Int) AddressInject {
	return AddressInject{address, amount}
}
func (bi AddressInject) String() string {
	return fmt.Sprintf(`
  Address:			%s
  Amount:			%s`,
		bi.Address.String(), bi.Amount.String())
}

//nolint
func (bi BoxInfo) String() string {
	return fmt.Sprintf(`Box:
  Id: 	         			%s
  Status:					%s
  Owner:           				%s
  Name:             			%s
  TotalAmount:      			%s
  CreatedTime:					%d
  Description:	    			%s
  TransferDisabled:			%t`,
		bi.Id, bi.Status, bi.Owner.String(), bi.Name, bi.TotalAmount.String(),
		bi.CreatedTime, bi.Description, bi.TransferDisabled)
}

//nolint
func (bi BoxInfos) String() string {
	out := fmt.Sprintf("%-17s|%-10s|%-44s|%-16s|%s\n",
		"BoxID", "Status", "Owner", "Name", "BoxTime")
	for _, box := range bi {
		out += fmt.Sprintf("%-17s|%-10s|%-44s|%-16s|%d\n",
			box.GetId(), box.GetStatus(), box.GetOwner().String(), box.GetName(), box.GetCreatedTime())
	}
	return strings.TrimSpace(out)
}
