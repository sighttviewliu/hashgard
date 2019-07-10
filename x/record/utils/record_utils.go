package utils

import (
	"math/rand"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/record/errors"
	"github.com/hashgard/hashgard/x/record/types"
)

var (
	randomBytes = []rune("abcdefghijklmnopqrstuvwxyz")
)

func GetRandomString(l int) string {
	result := make([]rune, l)
	length := len(randomBytes)
	for i := range result {
		result[i] = randomBytes[rand.Intn(length)]
	}
	return string(result)
}
func IsRecordId(id string) bool {
	return strings.HasPrefix(id, types.IDPreStr)
}

func CheckRecordId(issueID string) sdk.Error {
	if !IsRecordId(issueID) {
		return errors.ErrRecordIDNotValid(issueID)
	}
	return nil
}

func CheckRecordHash(hash string) sdk.Error {
	if len(hash) != 64 {
		return errors.ErrRecordHashNotValid(hash)
	}
	return nil
}