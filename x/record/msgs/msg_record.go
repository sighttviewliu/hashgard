package msgs

import (
	"fmt"

	"github.com/hashgard/hashgard/x/record/params"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/record/errors"
	"github.com/hashgard/hashgard/x/record/types"
)

// MsgRecord to allow a registered recordr
// to record new coins.
type MsgRecord struct {
	Sender              sdk.AccAddress  `json:"sender"`
	*params.RecordParams 				`json:"params"`
}

//New MsgRecord Instance
func NewMsgRecord(sender sdk.AccAddress, params *params.RecordParams) MsgRecord {
	return MsgRecord{sender, params}
}

// Route Implements Msg.
func (msg MsgRecord) Route() string { return types.RouterKey }

// Type Implements Msg.
func (msg MsgRecord) Type() string { return types.TypeMsgRecord }

// Implements Msg. Ensures addresses are valid and Coin is positive
func (msg MsgRecord) ValidateBasic() sdk.Error {
	if len(msg.Sender) == 0 {
		return sdk.ErrInvalidAddress("Sender address cannot be empty")
	}
	// record hash must be 64 characters length
	if len(msg.Hash) != types.HashLength {
		return errors.ErrRecordHashNotValid()
	}
	if len(msg.Name) < types.NameMinLength || len(msg.Name) > types.NameMaxLength {
		return errors.ErrRecordNameNotValid()
	}
	if len(msg.Author) < types.AuthorMinLength || len(msg.Author) > types.AuthorMaxLength {
		return errors.ErrRecordAuthorNotValid()
	}
	if len(msg.Type()) < types.TypeMinLength || len(msg.Author) > types.TypeMaxLength {
		return errors.ErrRecordTypeNotValid()
	}
	if len(msg.RecordNo) < types.RecordNoMinLength || len(msg.Author) > types.RecordNoMaxLength {
		return errors.ErrRecordNumberNotValid()
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgRecord) GetSignBytes() []byte {
	bz := MsgCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners Implements Msg.
func (msg MsgRecord) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

func (msg MsgRecord) String() string {
	return fmt.Sprintf("MsgRecord{%s - %s}", "", msg.Sender.String())
}
