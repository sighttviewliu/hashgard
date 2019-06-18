package gov

import (
	"strconv"
	"strings"
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

//fee
const (
	Fee         = "Fee"
	BoxModule   = "box/"
	IssueModule = "issue/"
)

//fee/box/lock_create=10000,fee/box/deposit_box_create=10000,fee/box/future_box_create=10000,
//fee/box/disable_feature=10000,fee/box/describe_fee=10000,fee/issue/create=10000,fee/issue/mint=10000
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
	if strings.HasPrefix(proposalParam.Key, BoxModule) || strings.HasPrefix(proposalParam.Key, IssueModule) {
		return nil
	}
	// check key
	switch proposalParam.Key {
	case communityTax, inflation, minSignedPerWindow, slashFractionDowntime:
		dec, err := sdk.NewDecFromStr(proposalParam.Value)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
		if dec.LT(sdk.NewDec(0)) || dec.GTE(sdk.NewDec(1)) {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, "the value is 0~1")
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
		v, err := strconv.ParseInt(proposalParam.Value, 10, 64)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
		if v < 1 {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, "the value must >1")
		}
	case maxValidators:
		v, err := strconv.ParseUint(proposalParam.Value, 10, 16)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
		if v < 4 {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, "the value must >4")
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
