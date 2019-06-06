package gov

import (
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	communityTax          = "distribution/community_tax"
	foundationAddress     = "mint/foundation_address"
	inflation             = "mint/inflation"
	inflationBase         = "mint/inflation_base"
	minDeposit            = "gov/min_deposit"
	signedBlocksWindow    = "slashing/signed_blocks_window"
	minSignedPerWindow    = "slashing/min_signed_per_window"
	slashFractionDowntime = "slashing/slash_fraction_downtime"
	downtimeJailDuration  = "slashing/downtime_jail_duration"
	maxValidators         = "staking/max_validators"
	unbondingTime         = "staking/unbonding_time"
)

var cdc = MakeCodec()

type ProposalParam struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ProposalParams []ProposalParam

func MakeCodec() *codec.Codec {
	cdc := codec.New()
	sdk.RegisterCodec(cdc)
	codec.RegisterCrypto(cdc)
	return cdc
}

func ValidateProposalParam(proposalParam ProposalParam) sdk.Error {
	// check key
	switch proposalParam.Key {
	case communityTax, inflation, minSignedPerWindow, slashFractionDowntime:
		_, err := sdk.NewDecFromStr(proposalParam.Value)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case minDeposit:
		_, err := sdk.ParseCoins(proposalParam.Value)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case downtimeJailDuration, unbondingTime:
		_, err := time.ParseDuration(proposalParam.Value)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case signedBlocksWindow:
		_, err := strconv.ParseInt(proposalParam.Value, 10, 64)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case maxValidators:
		_, err := strconv.ParseUint(proposalParam.Value, 10, 16)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	case inflationBase:
		_, ok := sdk.NewIntFromString(proposalParam.Value)
		if !ok {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, "invalid string to Int")
		}
	case foundationAddress:
		_, err := sdk.AccAddressFromBech32(proposalParam.Value)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
	default:
		return ErrInvalidParamKey(DefaultCodespace, proposalParam.Key)
	}

	return nil
}
