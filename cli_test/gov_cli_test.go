// +build cli_test

package clitest

import (
	"fmt"
	"strings"
	"testing"

	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/app"
	"github.com/hashgard/hashgard/x/gov"
	"github.com/stretchr/testify/require"
)

func TestHashgardCLISubmitProposalParameterChange(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)

	// start hashgard server
	proc := f.HGStart()
	defer proc.Stop(false)

	fooAddr := f.KeyAddress(keyFoo)

	fooAcc := f.QueryAccount(fooAddr)
	startTokens := sdk.TokensFromTendermintPower(50)
	require.Equal(t, startTokens, fooAcc.GetCoins().AmountOf(app.StakeDenom))

	proposalsQuery := f.QueryGovProposals()
	require.Empty(t, proposalsQuery)

	// Test submit generate only for submit proposal
	proposalTokens := sdk.TokensFromTendermintPower(50000)

	niceVal, params := gov.GetProposalParam(f.KeyAddress(keyIssue).String())
	paramStrs := make([]string, 0)
	for _, param := range params {
		paramStrs = append(paramStrs, fmt.Sprintf("%s=%s", param.Key, param.Value))
	}
	// Create the proposal
	f.TxGovSubmitProposal(keyIssue, "ParameterChange", "Test", "test",
		sdk.NewCoin(denom, proposalTokens), "-y", "--param="+strings.Join(paramStrs, ","))
	tests.WaitForNextNBlocksTM(1, f.Port)

	// Vote on the proposal
	success, _, stderr := f.TxGovVote(1, gov.OptionYes, keyFoo, "-y")
	require.True(f.T, success)
	require.Empty(f.T, stderr)
	tests.WaitForNextNBlocksTM(10, f.Port)

	depositParams := f.QueryGovParamDeposit()
	require.Equal(t, depositParams.MinDeposit.AmountOf(sdk.DefaultBondDenom), niceVal.Amount)

	f.Cleanup()
}
