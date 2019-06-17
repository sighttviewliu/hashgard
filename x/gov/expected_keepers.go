package gov

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"

	"github.com/hashgard/hashgard/x/mint"
)

// expected bank keeper
type BankKeeper interface {
	GetCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	// TODO remove once governance doesn't require use of accounts
	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) sdk.Error
	SetSendEnabled(ctx sdk.Context, enabled bool)
}

type AuthKeeper interface {
	GetParams(ctx sdk.Context) (params auth.Params)
	SetParams(ctx sdk.Context, params auth.Params)
}

type DistributionKeeper interface {
	SetCommunityTax(ctx sdk.Context, percent sdk.Dec)
	AllocateCommunityPool(ctx sdk.Context, destAddr sdk.AccAddress, percent sdk.Dec, burn bool) sdk.Error
}

type MintKeeper interface {
	GetParams(ctx sdk.Context) mint.Params
	SetParams(ctx sdk.Context, params mint.Params)
}

type SlashingKeeper interface {
	GetParams(ctx sdk.Context) (params slashing.Params)
	SetParams(ctx sdk.Context, params slashing.Params)
}

type StakingKeeper interface {
	GetParams(ctx sdk.Context) staking.Params
	SetParams(ctx sdk.Context, params staking.Params)
}
