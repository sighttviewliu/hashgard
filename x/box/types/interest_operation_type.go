package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var InterestOperation = map[string]uint{Inject: 1, Cancel: 2}

func CheckInterestOperation(interestOperation string) sdk.Error {
	_, ok := InterestOperation[interestOperation]
	if !ok {
		return sdk.ErrInternal("unknown interest operation:" + interestOperation)
	}
	return nil
}
