package errors

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/record/types"
)

const (
	CodeRecordExist                sdk.CodeType = 1
	CodeRecordHashNotValid         sdk.CodeType = 2
	CodeRecordIDNotValid           sdk.CodeType = 3
	CodeRecordNumberNotValid           sdk.CodeType = 4
	CodeRecordNameNotValid         sdk.CodeType = 5
	CodeRecordAuthorNotValid       sdk.CodeType = 6
	CodeRecordTypeNotValid     	   sdk.CodeType = 7
	CodeUnknownRecord              sdk.CodeType = 8
	CodeUnknownAuthor              sdk.CodeType = 9
)

//convert sdk.Error to error
func Errorf(err sdk.Error) error {
	return fmt.Errorf(err.Stacktrace().Error())
}

// Error constructors
func ErrRecordExist(record string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeRecordExist, fmt.Sprintf("Record %s exist", record))
}
func ErrRecordHashNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeRecordHashNotValid, fmt.Sprintf("Record hash length must be %d character", types.HashLength))
}
func ErrRecordNumberNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeRecordNumberNotValid, fmt.Sprintf("RecordNo. max length is %d character", types.RecordNoMaxLength))
}
func ErrRecordAuthorNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeRecordAuthorNotValid, fmt.Sprintf("Author max length is %d character", types.AuthorMaxLength))
}
func ErrRecordNameNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeRecordNameNotValid, fmt.Sprintf("The length of the name is between %d and %d", types.NameMinLength, types.NameMaxLength))
}
func ErrRecordTypeNotValid() sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeRecordTypeNotValid, fmt.Sprintf("The max length of record type is %d", types.RecordTypeMaxLength))
}
func ErrRecordIDNotValid(recordID string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeRecordIDNotValid, fmt.Sprintf("Record-id %s is not a valid recordId", recordID))
}
func ErrUnknownRecord(record string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeUnknownRecord, fmt.Sprintf("Unknown record %s", record))
}
func ErrUnknownAuthor(author string) sdk.Error {
	return sdk.NewError(types.DefaultCodespace, CodeUnknownAuthor, fmt.Sprintf("Unknown record %s", author))
}
