package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"

	"strings"

	"github.com/spf13/cobra"

	"github.com/hashgard/hashgard/x/gov"
	govClientUtils "github.com/hashgard/hashgard/x/gov/client/utils"
)

type proposal struct {
	Title       string
	Description string
	Type        string
	Deposit     string
	Params      []string
	Usage       string
	Percent     string
	DestAddress string
}

var proposalFlags = []string{
	flagTitle,
	flagDescription,
	flagDescription,
	flagProposalType,
	flagDeposit,
}

// GetCmdSubmitProposal implements submitting a proposal transaction command.
func GetCmdSubmitProposal(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-proposal",
		Short: "Submit a proposal along with an initial deposit",
		Long: strings.TrimSpace(`
Submit a proposal along with an initial deposit. Proposal title, description, type and deposit can be given directly or through a proposal JSON file. For example:

$ hashgardcli gov submit-proposal --proposal="path/to/proposal.json" --from mykey

where proposal.json contains:

{
  "title": "Test Proposal",
  "description": "My awesome proposal",
  "type": "Text",
  "deposit": "10test"
}

is equivalent to

$ hashgardcli gov submit-proposal --title="Test Proposal" --description="My awesome proposal" --type="Text" --deposit="10test" --from mykey
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			proposal, err := parseSubmitProposalFlags()
			if err != nil {
				return err
			}

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)

			// Get from address
			from := cliCtx.GetFromAddress()

			// Pull associated account
			account, err := cliCtx.GetAccount(from)
			if err != nil {
				return err
			}

			// Find deposit amount
			amount, err := sdk.ParseCoins(proposal.Deposit)
			if err != nil {
				return err
			}

			// ensure account has enough coins
			if !account.GetCoins().IsAllGTE(amount) {
				return fmt.Errorf("address %s doesn't have enough coins to pay for this transaction", from)
			}

			proposalType, err := gov.ProposalTypeFromString(proposal.Type)
			if err != nil {
				return err
			}

			var proposalParams gov.ProposalParams
			if proposalType == gov.ProposalTypeParameterChange {
				proposalParams, err = getParamFromString(proposal.Params)
				if err != nil {
					return err
				}
			}

			var taxUsage gov.TaxUsage
			taxUsage.Usage, err = gov.UsageTypeFromString(proposal.Usage)
			if err != nil {
				return err
			}

			taxUsage.DestAddress, err = sdk.AccAddressFromBech32(proposal.DestAddress)
			if err != nil {
				return err
			}

			taxUsage.Percent, err = sdk.NewDecFromStr(proposal.Percent)
			if err != nil {
				return err
			}

			msg := gov.NewMsgSubmitProposal(proposal.Title, proposal.Description, proposalType, from, amount, proposalParams, taxUsage)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}

	cmd.Flags().String(flagTitle, "", "title of proposal")
	cmd.Flags().String(flagDescription, "", "description of proposal")
	cmd.Flags().String(flagProposalType, "", "proposalType of proposal, types: Text/ParameterChange/SoftwareUpgrade/TaxUsage")
	cmd.Flags().String(flagDeposit, "", "deposit of proposal")
	cmd.Flags().String(flagProposal, "", "proposal file path (if this path is given, other proposal flags are ignored)")
	cmd.Flags().StringSlice(flagParam, []string{}, "parameter of proposal,eg. \"mint/Inflation=0.03\",\"staking/max_validators\"")
	cmd.Flags().String(flagUsage, "", "the transaction fee tax usage type, valid values can be Burn and Grant")
	cmd.Flags().String(flagPercent, "0", "percent of transaction fee tax pool to use, integer or decimal >0 and <=1")
	cmd.Flags().String(flagDestAddress, "", "the destination trustee address")

	return cmd
}

// GetCmdDeposit implements depositing tokens for an active proposal.
func GetCmdDeposit(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "deposit [proposal-id] [deposit]",
		Args:  cobra.ExactArgs(2),
		Short: "Deposit tokens for activing proposal",
		Long: strings.TrimSpace(`
Submit a deposit for an acive proposal. You can find the proposal-id by running hashgardcli query gov proposals:

$ hashgardcli gov deposit 1 10gard --from mykey
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)

			// validate that the proposal id is a uint
			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("proposal-id %s not a valid uint, please input a valid proposal-id", args[0])
			}

			// check to see if the proposal is in the store
			_, err = govClientUtils.QueryProposalByID(proposalID, cliCtx, cdc, queryRoute)
			if err != nil {
				return fmt.Errorf("Failed to fetch proposal-id %d: %s", proposalID, err)
			}

			from := cliCtx.GetFromAddress()

			// Fetch associated account
			account, err := cliCtx.GetAccount(from)
			if err != nil {
				return err
			}

			// Get amount of coins
			amount, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}

			// ensure account has enough coins
			if !account.GetCoins().IsAllGTE(amount) {
				return fmt.Errorf("address %s doesn't have enough coins to pay for this transaction", from)
			}

			msg := gov.NewMsgDeposit(from, proposalID, amount)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
}

// GetCmdVote implements creating a new vote command.
func GetCmdVote(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "vote [proposal-id] [option]",
		Args:  cobra.ExactArgs(2),
		Short: "Vote for an active proposal, options: yes/no/no_with_veto/abstain",
		Long: strings.TrimSpace(`
Submit a vote for an acive proposal. You can find the proposal-id by running hashgardcli gov query-proposals:

$ hashgardcli gov vote 1 yes --from mykey
`),
		RunE: func(cmd *cobra.Command, args []string) error {
			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
			cliCtx := context.NewCLIContext().
				WithCodec(cdc).
				WithAccountDecoder(cdc)

			// Get voting address
			from := cliCtx.GetFromAddress()

			// validate that the proposal id is a uint
			proposalID, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return fmt.Errorf("proposal-id %s not a valid int, please input a valid proposal-id", args[0])
			}

			// check to see if the proposal is in the store
			_, err = govClientUtils.QueryProposalByID(proposalID, cliCtx, cdc, queryRoute)
			if err != nil {
				return fmt.Errorf("Failed to fetch proposal-id %d: %s", proposalID, err)
			}

			// Find out which vote option user chose
			byteVoteOption, err := gov.VoteOptionFromString(govClientUtils.NormalizeVoteOption(args[1]))
			if err != nil {
				return err
			}

			// Build vote message and run basic validation
			msg := gov.NewMsgVote(from, proposalID, byteVoteOption)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
}

// DONTCOVER

func getParamFromString(paramStrArr []string) (gov.ProposalParams, error) {
	var proposalParams gov.ProposalParams
	for _, paramStr := range paramStrArr {
		str := strings.Split(paramStr, "=")
		if len(str) != 2 {
			return gov.ProposalParams{}, fmt.Errorf("%s is not valid", paramStr)
		}
		//str = []string{"mint/Inflation","0.0000000000"}
		//params.GetParamSpaceFromKey(str[0]) == "mint"
		//params.GetParamKey(str[0])          == "Inflation"
		proposalParams = append(proposalParams,
			gov.ProposalParam{
				Key:   str[0],
				Value: str[1],
			})
	}
	return proposalParams, nil
}
