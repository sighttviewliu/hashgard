package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	Lock    = "lock"
	Deposit = "deposit"
	Future  = "future"
)

var BoxType = map[string]string{Lock: "aa", Deposit: "ab", Future: "ac"}

func GetMustBoxTypeValue(boxType string) string {
	value, ok := BoxType[boxType]
	if !ok {
		panic("unknown type")
	}
	return value
}

func CheckBoxType(boxType string) sdk.Error {
	_, ok := BoxType[boxType]
	if !ok {
		return sdk.ErrInternal("unknown type:" + boxType)
	}
	return nil
}

func GetBoxTypeValue(boxType string) (string, error) {
	value, ok := BoxType[boxType]
	if !ok {
		return "", fmt.Errorf("unknown type:%s", boxType)
	}
	return value, nil
}
