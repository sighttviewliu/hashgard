package gov

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/hashgard/hashgard/x/gov/tags"
)

// Handle all "gov" type messages.
func NewHandler(keeper Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result {
		switch msg := msg.(type) {
		case MsgDeposit:
			return handleMsgDeposit(ctx, keeper, msg)
		case MsgSubmitProposal:
			return handleMsgSubmitProposal(ctx, keeper, msg)
		case MsgVote:
			return handleMsgVote(ctx, keeper, msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized gov msg type: %T", msg)
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}

func handleMsgSubmitProposal(ctx sdk.Context, keeper Keeper, msg MsgSubmitProposal) sdk.Result {
	var content ProposalContent
	switch msg.ProposalType {
	case ProposalTypeText:
		textProposal := NewTextProposal(msg.Title, msg.Description)
		content = &textProposal
	case ProposalTypeSoftwareUpgrade:
		softwareUpgradeProposal := NewSoftwareUpgradeProposal(msg.Title, msg.Description)
		content = &softwareUpgradeProposal
	case ProposalTypeParameterChange:
		for _, param := range msg.ProposalParams {
			if strings.HasPrefix(param.Key, BoxModule) || strings.HasPrefix(param.Key, IssueModule) {
				subspace := GetParamSpaceFromKey(param.Key)
				if p, ok := keeper.paramsKeeper.GetSubspace(subspace); ok {
					key := CamelString(GetParamKey(param.Key))
					if !p.Has(ctx, []byte(key)) {
						return ErrInvalidParamValue(DefaultCodespace, param.Key, param.Value, "undefined").Result()
					}
					if strings.HasSuffix(key, Fee) {
						coin, err := sdk.ParseCoin(param.Value)
						if err != nil {
							return sdk.ErrInvalidCoins(param.Value).Result()
						}
						if coin.IsNegative() {
							return sdk.ErrInvalidCoins(param.Value).Result()
						}
					}
				} else {
					return ErrInvalidParamValue(DefaultCodespace, subspace, subspace, "undefined").Result()
				}
			}
		}
		parameterChangeProposal := NewParameterChangeProposal(msg.Title, msg.Description, msg.ProposalParams)
		content = &parameterChangeProposal
	case ProposalTypeTaxUsage:
		taxUsageProposal := NewTaxUsageProposal(msg.Title, msg.Description, msg.TaxUsage)
		content = &taxUsageProposal
	default:
		return ErrInvalidProposalType(keeper.codespace, msg.ProposalType).Result()
	}

	proposal, err := keeper.SubmitProposal(ctx, content)
	if err != nil {
		return err.Result()
	}
	proposalID := proposal.ProposalID
	proposalIDStr := fmt.Sprintf("%d", proposalID)

	err, votingStarted := keeper.AddDeposit(ctx, proposalID, msg.Proposer, msg.InitialDeposit)
	if err != nil {
		return err.Result()
	}

	resTags := sdk.NewTags(
		tags.ProposalID, proposalIDStr,
		tags.Category, tags.TxCategory,
		tags.Sender, msg.Proposer.String(),
		tags.ProposalType, msg.ProposalType.String(),
	)

	if votingStarted {
		resTags = resTags.AppendTag(tags.VotingPeriodStart, proposalIDStr)
	}

	return sdk.Result{
		Data: keeper.cdc.MustMarshalBinaryLengthPrefixed(proposalID),
		Tags: resTags,
	}
}

func handleMsgDeposit(ctx sdk.Context, keeper Keeper, msg MsgDeposit) sdk.Result {
	err, votingStarted := keeper.AddDeposit(ctx, msg.ProposalID, msg.Depositor, msg.Amount)
	if err != nil {
		return err.Result()
	}

	proposalIDStr := fmt.Sprintf("%d", msg.ProposalID)

	resTags := sdk.NewTags(
		tags.ProposalID, proposalIDStr,
		tags.Category, tags.TxCategory,
		tags.Sender, msg.Depositor.String(),
	)

	if votingStarted {
		resTags = resTags.AppendTag(tags.VotingPeriodStart, proposalIDStr)
	}

	return sdk.Result{
		Tags: resTags,
	}
}

func handleMsgVote(ctx sdk.Context, keeper Keeper, msg MsgVote) sdk.Result {
	err := keeper.AddVote(ctx, msg.ProposalID, msg.Voter, msg.Option)
	if err != nil {
		return err.Result()
	}

	proposalIDStr := fmt.Sprintf("%d", msg.ProposalID)

	return sdk.Result{
		Tags: sdk.NewTags(
			tags.ProposalID, proposalIDStr,
			tags.Category, tags.TxCategory,
			tags.Sender, msg.Voter.String(),
		),
	}
}
