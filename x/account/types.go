package account

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strings"
)

// Account interface
type Account interface {
	GetSender() sdk.AccAddress
	SetSender(sdk.AccAddress)

	GetMustMemo() bool
	SetMustMemo(bool)

	String() string
}

// Account Info
type AccountInfo struct {
	Sender   sdk.AccAddress `json:"sender"`
	MustMemo bool           `json:"must_memo"`
}

// Accounts is an array of AccountInfo
type Accounts []AccountInfo

// Implements Account Interface
var _ Account = (*AccountInfo)(nil)

//nolint
func (ci AccountInfo) GetSender() sdk.AccAddress {
	return ci.Sender
}
func (ci *AccountInfo) SetSender(operator sdk.AccAddress) {
	ci.Sender = operator
}
func (ci AccountInfo) GetMustMemo() bool {
	return ci.MustMemo
}
func (ci *AccountInfo) SetMustMemo(mustMemo bool) {
	ci.MustMemo = mustMemo
}

//nolint
func (ci AccountInfo) String() string {
	return fmt.Sprintf(`Account:
  Sender:           			%s
  MustMemo:            			%s`,
		ci.Sender.String(), ci.MustMemo)
}

//nolint
func (ls Accounts) String() string {
	out := fmt.Sprintf("%-44s|%s\n", "Sender", "MustMemo")
	for _, acc := range ls {
		out += fmt.Sprintf("%-44s|%s\n",
			acc.GetSender().String(), acc.MustMemo)
	}
	return strings.TrimSpace(out)
}
