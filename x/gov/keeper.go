package gov

import (
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"

	"github.com/tendermint/tendermint/crypto"
)

const (
	// ModuleKey is the name of the module
	ModuleName = "gov"

	// StoreKey is the store key string for gov
	StoreKey = ModuleName

	// RouterKey is the message route for gov
	RouterKey = ModuleName

	// QuerierRoute is the querier route for gov
	QuerierRoute = ModuleName

	// Parameter store default namestore
	DefaultParamspace = ModuleName
)

// Parameter store key
var (
	ParamStoreKeyDepositParams = []byte("depositparams")
	ParamStoreKeyVotingParams  = []byte("votingparams")
	ParamStoreKeyTallyParams   = []byte("tallyparams")

	// TODO: Find another way to implement this without using accounts, or find a cleaner way to implement it using accounts.
	DepositedCoinsAccAddr     = sdk.AccAddress(crypto.AddressHash([]byte("govDepositedCoins")))
	BurnedDepositCoinsAccAddr = sdk.AccAddress(crypto.AddressHash([]byte("govBurnedDepositCoins")))
)

// Key declaration for parameters
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamStoreKeyDepositParams, DepositParams{},
		ParamStoreKeyVotingParams, VotingParams{},
		ParamStoreKeyTallyParams, TallyParams{},
	)
}

// Governance Keeper
type Keeper struct {
	// The reference to the Param Keeper to get and set Global Params
	paramsKeeper params.Keeper

	// The reference to the Paramstore to get and set gov specific params
	paramSpace params.Subspace

	// The reference to the CoinKeeper to modify balances
	ck BankKeeper

	// The reference to these Keepers to modify parameters
	authKeeper         AuthKeeper
	distributionKeeper DistributionKeeper
	mintKeeper         MintKeeper
	slashingKeeper     SlashingKeeper
	stakingKeeper      StakingKeeper

	// The ValidatorSet to get information about validators
	vs sdk.ValidatorSet

	// The reference to the DelegationSet to get information about delegators
	ds sdk.DelegationSet

	// The (unexposed) keys used to access the stores from the Context.
	storeKey sdk.StoreKey

	// The codec codec for binary encoding/decoding.
	cdc *codec.Codec

	// Reserved codespace
	codespace sdk.CodespaceType
}

// NewKeeper returns a governance keeper. It handles:
// - submitting governance proposals
// - depositing funds into proposals, and activating upon sufficient funds being deposited
// - users voting on proposals, with weight proportional to stake in the system
// - and tallying the result of the vote.
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramsKeeper params.Keeper,
	paramSpace params.Subspace, ck BankKeeper, ds sdk.DelegationSet, codespace sdk.CodespaceType,
	authKeeper AuthKeeper, distributionKeeper DistributionKeeper, mintKeeper MintKeeper,
	slashingKeeper SlashingKeeper, stakingKeeper StakingKeeper) Keeper {

	return Keeper{
		storeKey:           key,
		paramsKeeper:       paramsKeeper,
		paramSpace:         paramSpace.WithKeyTable(ParamKeyTable()),
		ck:                 ck,
		ds:                 ds,
		vs:                 ds.GetValidatorSet(),
		cdc:                cdc,
		codespace:          codespace,
		authKeeper:         authKeeper,
		distributionKeeper: distributionKeeper,
		mintKeeper:         mintKeeper,
		slashingKeeper:     slashingKeeper,
		stakingKeeper:      stakingKeeper,
	}
}

// Proposals
func (keeper Keeper) SubmitProposal(ctx sdk.Context, content ProposalContent) (proposal Proposal, err sdk.Error) {
	proposalID, err := keeper.getNewProposalID(ctx)
	if err != nil {
		return
	}

	submitTime := ctx.BlockHeader().Time
	depositPeriod := keeper.GetDepositParams(ctx).MaxDepositPeriod

	proposal = Proposal{
		ProposalContent: content,
		ProposalID:      proposalID,

		Status:           StatusDepositPeriod,
		FinalTallyResult: EmptyTallyResult(),
		TotalDeposit:     sdk.NewCoins(),
		SubmitTime:       submitTime,
		DepositEndTime:   submitTime.Add(depositPeriod),
	}

	keeper.SetProposal(ctx, proposal)
	keeper.InsertInactiveProposalQueue(ctx, proposal.DepositEndTime, proposalID)
	return
}

// Get Proposal from store by ProposalID
func (keeper Keeper) GetProposal(ctx sdk.Context, proposalID uint64) (proposal Proposal, ok bool) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyProposal(proposalID))
	if bz == nil {
		return
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &proposal)
	return proposal, true
}

// Implements sdk.AccountKeeper.
func (keeper Keeper) SetProposal(ctx sdk.Context, proposal Proposal) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(proposal)
	store.Set(KeyProposal(proposal.ProposalID), bz)
}

// Implements sdk.AccountKeeper.
func (keeper Keeper) DeleteProposal(ctx sdk.Context, proposalID uint64) {
	store := ctx.KVStore(keeper.storeKey)
	proposal, ok := keeper.GetProposal(ctx, proposalID)
	if !ok {
		panic("DeleteProposal cannot fail to GetProposal")
	}
	keeper.RemoveFromInactiveProposalQueue(ctx, proposal.DepositEndTime, proposalID)
	keeper.RemoveFromActiveProposalQueue(ctx, proposal.VotingEndTime, proposalID)
	store.Delete(KeyProposal(proposalID))
}

// Get Proposal from store by ProposalID
// voterAddr will filter proposals by whether or not that address has voted on them
// depositorAddr will filter proposals by whether or not that address has deposited to them
// status will filter proposals by status
// numLatest will fetch a specified number of the most recent proposals, or 0 for all proposals
func (keeper Keeper) GetProposalsFiltered(ctx sdk.Context, voterAddr sdk.AccAddress, depositorAddr sdk.AccAddress, status ProposalStatus, numLatest uint64) []Proposal {

	maxProposalID, err := keeper.peekCurrentProposalID(ctx)
	if err != nil {
		return nil
	}

	matchingProposals := []Proposal{}

	if numLatest == 0 {
		numLatest = maxProposalID
	}

	for proposalID := maxProposalID - numLatest; proposalID < maxProposalID; proposalID++ {
		if voterAddr != nil && len(voterAddr) != 0 {
			_, found := keeper.GetVote(ctx, proposalID, voterAddr)
			if !found {
				continue
			}
		}

		if depositorAddr != nil && len(depositorAddr) != 0 {
			_, found := keeper.GetDeposit(ctx, proposalID, depositorAddr)
			if !found {
				continue
			}
		}

		proposal, ok := keeper.GetProposal(ctx, proposalID)
		if !ok {
			continue
		}

		if validProposalStatus(status) {
			if proposal.Status != status {
				continue
			}
		}

		matchingProposals = append(matchingProposals, proposal)
	}
	return matchingProposals
}

// Set the initial proposal ID
func (keeper Keeper) setInitialProposalID(ctx sdk.Context, proposalID uint64) sdk.Error {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextProposalID)
	if bz != nil {
		return ErrInvalidGenesis(keeper.codespace, "Initial ProposalID already set")
	}
	bz = keeper.cdc.MustMarshalBinaryLengthPrefixed(proposalID)
	store.Set(KeyNextProposalID, bz)
	return nil
}

// Get the last used proposal ID
func (keeper Keeper) GetLastProposalID(ctx sdk.Context) (proposalID uint64) {
	proposalID, err := keeper.peekCurrentProposalID(ctx)
	if err != nil {
		return 0
	}
	proposalID--
	return
}

// Gets the next available ProposalID and increments it
func (keeper Keeper) getNewProposalID(ctx sdk.Context) (proposalID uint64, err sdk.Error) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextProposalID)
	if bz == nil {
		return 0, ErrInvalidGenesis(keeper.codespace, "InitialProposalID never set")
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &proposalID)
	bz = keeper.cdc.MustMarshalBinaryLengthPrefixed(proposalID + 1)
	store.Set(KeyNextProposalID, bz)
	return proposalID, nil
}

// Peeks the next available ProposalID without incrementing it
func (keeper Keeper) peekCurrentProposalID(ctx sdk.Context) (proposalID uint64, err sdk.Error) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextProposalID)
	if bz == nil {
		return 0, ErrInvalidGenesis(keeper.codespace, "InitialProposalID never set")
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &proposalID)
	return proposalID, nil
}

func (keeper Keeper) activateVotingPeriod(ctx sdk.Context, proposal Proposal) {
	proposal.VotingStartTime = ctx.BlockHeader().Time
	votingPeriod := keeper.GetVotingParams(ctx).VotingPeriod
	proposal.VotingEndTime = proposal.VotingStartTime.Add(votingPeriod)
	proposal.Status = StatusVotingPeriod
	keeper.SetProposal(ctx, proposal)

	keeper.RemoveFromInactiveProposalQueue(ctx, proposal.DepositEndTime, proposal.ProposalID)
	keeper.InsertActiveProposalQueue(ctx, proposal.VotingEndTime, proposal.ProposalID)
}

// Params

// Returns the current DepositParams from the global param store
func (keeper Keeper) GetDepositParams(ctx sdk.Context) DepositParams {
	var depositParams DepositParams
	keeper.paramSpace.Get(ctx, ParamStoreKeyDepositParams, &depositParams)
	return depositParams
}

// Returns the current VotingParams from the global param store
func (keeper Keeper) GetVotingParams(ctx sdk.Context) VotingParams {
	var votingParams VotingParams
	keeper.paramSpace.Get(ctx, ParamStoreKeyVotingParams, &votingParams)
	return votingParams
}

// Returns the current TallyParam from the global param store
func (keeper Keeper) GetTallyParams(ctx sdk.Context) TallyParams {
	var tallyParams TallyParams
	keeper.paramSpace.Get(ctx, ParamStoreKeyTallyParams, &tallyParams)
	return tallyParams
}

func (keeper Keeper) setDepositParams(ctx sdk.Context, depositParams DepositParams) {
	keeper.paramSpace.Set(ctx, ParamStoreKeyDepositParams, &depositParams)
}

func (keeper Keeper) setVotingParams(ctx sdk.Context, votingParams VotingParams) {
	keeper.paramSpace.Set(ctx, ParamStoreKeyVotingParams, &votingParams)
}

func (keeper Keeper) setTallyParams(ctx sdk.Context, tallyParams TallyParams) {
	keeper.paramSpace.Set(ctx, ParamStoreKeyTallyParams, &tallyParams)
}

// Votes

// Adds a vote on a specific proposal
func (keeper Keeper) AddVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress, option VoteOption) sdk.Error {
	proposal, ok := keeper.GetProposal(ctx, proposalID)
	if !ok {
		return ErrUnknownProposal(keeper.codespace, proposalID)
	}
	if proposal.Status != StatusVotingPeriod {
		return ErrInactiveProposal(keeper.codespace, proposalID)
	}

	if !validVoteOption(option) {
		return ErrInvalidVote(keeper.codespace, option)
	}

	vote := Vote{
		ProposalID: proposalID,
		Voter:      voterAddr,
		Option:     option,
	}
	keeper.setVote(ctx, proposalID, voterAddr, vote)

	return nil
}

// Gets the vote of a specific voter on a specific proposal
func (keeper Keeper) GetVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) (Vote, bool) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyVote(proposalID, voterAddr))
	if bz == nil {
		return Vote{}, false
	}
	var vote Vote
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &vote)
	return vote, true
}

func (keeper Keeper) setVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress, vote Vote) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(vote)
	store.Set(KeyVote(proposalID, voterAddr), bz)
}

// Gets all the votes on a specific proposal
func (keeper Keeper) GetVotes(ctx sdk.Context, proposalID uint64) sdk.Iterator {
	store := ctx.KVStore(keeper.storeKey)
	return sdk.KVStorePrefixIterator(store, KeyVotesSubspace(proposalID))
}

func (keeper Keeper) deleteVote(ctx sdk.Context, proposalID uint64, voterAddr sdk.AccAddress) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(KeyVote(proposalID, voterAddr))
}

// Deposits

// Gets the deposit of a specific depositor on a specific proposal
func (keeper Keeper) GetDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress) (Deposit, bool) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyDeposit(proposalID, depositorAddr))
	if bz == nil {
		return Deposit{}, false
	}
	var deposit Deposit
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &deposit)
	return deposit, true
}

func (keeper Keeper) setDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress, deposit Deposit) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(deposit)
	store.Set(KeyDeposit(proposalID, depositorAddr), bz)
}

// Adds or updates a deposit of a specific depositor on a specific proposal
// Activates voting period when appropriate
func (keeper Keeper) AddDeposit(ctx sdk.Context, proposalID uint64, depositorAddr sdk.AccAddress, depositAmount sdk.Coins) (sdk.Error, bool) {
	// Checks to see if proposal exists
	proposal, ok := keeper.GetProposal(ctx, proposalID)
	if !ok {
		return ErrUnknownProposal(keeper.codespace, proposalID), false
	}

	// Check if proposal is still depositable
	if (proposal.Status != StatusDepositPeriod) && (proposal.Status != StatusVotingPeriod) {
		return ErrAlreadyFinishedProposal(keeper.codespace, proposalID), false
	}

	// Send coins from depositor's account to DepositedCoinsAccAddr account
	// TODO: Don't use an account for this purpose; it's clumsy and prone to misuse.
	err := keeper.ck.SendCoins(ctx, depositorAddr, DepositedCoinsAccAddr, depositAmount)
	if err != nil {
		return err, false
	}

	// Update proposal
	proposal.TotalDeposit = proposal.TotalDeposit.Add(depositAmount)
	keeper.SetProposal(ctx, proposal)

	// Check if deposit has provided sufficient total funds to transition the proposal into the voting period
	activatedVotingPeriod := false
	if proposal.Status == StatusDepositPeriod && proposal.TotalDeposit.IsAllGTE(keeper.GetDepositParams(ctx).MinDeposit) {
		keeper.activateVotingPeriod(ctx, proposal)
		activatedVotingPeriod = true
	}

	// Add or update deposit object
	currDeposit, found := keeper.GetDeposit(ctx, proposalID, depositorAddr)
	if !found {
		newDeposit := Deposit{depositorAddr, proposalID, depositAmount}
		keeper.setDeposit(ctx, proposalID, depositorAddr, newDeposit)
	} else {
		currDeposit.Amount = currDeposit.Amount.Add(depositAmount)
		keeper.setDeposit(ctx, proposalID, depositorAddr, currDeposit)
	}

	return nil, activatedVotingPeriod
}

// Gets all the deposits on a specific proposal as an sdk.Iterator
func (keeper Keeper) GetDeposits(ctx sdk.Context, proposalID uint64) sdk.Iterator {
	store := ctx.KVStore(keeper.storeKey)
	return sdk.KVStorePrefixIterator(store, KeyDepositsSubspace(proposalID))
}

// Refunds and deletes all the deposits on a specific proposal
func (keeper Keeper) RefundDeposits(ctx sdk.Context, proposalID uint64) {
	store := ctx.KVStore(keeper.storeKey)
	depositsIterator := keeper.GetDeposits(ctx, proposalID)
	defer depositsIterator.Close()
	for ; depositsIterator.Valid(); depositsIterator.Next() {
		deposit := &Deposit{}
		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(depositsIterator.Value(), deposit)

		err := keeper.ck.SendCoins(ctx, DepositedCoinsAccAddr, deposit.Depositor, deposit.Amount)
		if err != nil {
			panic("should not happen")
		}

		store.Delete(depositsIterator.Key())
	}
}

// Deletes all the deposits on a specific proposal without refunding them
func (keeper Keeper) DeleteDeposits(ctx sdk.Context, proposalID uint64) {
	store := ctx.KVStore(keeper.storeKey)
	depositsIterator := keeper.GetDeposits(ctx, proposalID)
	defer depositsIterator.Close()
	for ; depositsIterator.Valid(); depositsIterator.Next() {
		deposit := &Deposit{}
		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(depositsIterator.Value(), deposit)

		// TODO: Find a way to do this without using accounts.
		err := keeper.ck.SendCoins(ctx, DepositedCoinsAccAddr, BurnedDepositCoinsAccAddr, deposit.Amount)
		if err != nil {
			panic("should not happen")
		}

		store.Delete(depositsIterator.Key())
	}
}

// ProposalQueues

// Returns an iterator for all the proposals in the Active Queue that expire by endTime
func (keeper Keeper) ActiveProposalQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(keeper.storeKey)
	return store.Iterator(PrefixActiveProposalQueue, sdk.PrefixEndBytes(PrefixActiveProposalQueueTime(endTime)))
}

// Inserts a ProposalID into the active proposal queue at endTime
func (keeper Keeper) InsertActiveProposalQueue(ctx sdk.Context, endTime time.Time, proposalID uint64) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(proposalID)
	store.Set(KeyActiveProposalQueueProposal(endTime, proposalID), bz)
}

// removes a proposalID from the Active Proposal Queue
func (keeper Keeper) RemoveFromActiveProposalQueue(ctx sdk.Context, endTime time.Time, proposalID uint64) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(KeyActiveProposalQueueProposal(endTime, proposalID))
}

// Returns an iterator for all the proposals in the Inactive Queue that expire by endTime
func (keeper Keeper) InactiveProposalQueueIterator(ctx sdk.Context, endTime time.Time) sdk.Iterator {
	store := ctx.KVStore(keeper.storeKey)
	return store.Iterator(PrefixInactiveProposalQueue, sdk.PrefixEndBytes(PrefixInactiveProposalQueueTime(endTime)))
}

// Inserts a ProposalID into the inactive proposal queue at endTime
func (keeper Keeper) InsertInactiveProposalQueue(ctx sdk.Context, endTime time.Time, proposalID uint64) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(proposalID)
	store.Set(KeyInactiveProposalQueueProposal(endTime, proposalID), bz)
}

// removes a proposalID from the Inactive Proposal Queue
func (keeper Keeper) RemoveFromInactiveProposalQueue(ctx sdk.Context, endTime time.Time, proposalID uint64) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(KeyInactiveProposalQueueProposal(endTime, proposalID))
}

func (keeper Keeper) ExecuteProposal(ctx sdk.Context, proposal Proposal) sdk.Error {
	switch proposal.ProposalType() {
	case ProposalTypeParameterChange:
		proposalParams := proposal.ProposalContent.(*ParameterChangeProposal).ProposalParams
		for _, param := range proposalParams {
			err := keeper.SetProposalParam(ctx, param)
			if err != nil {
				return err
			}
		}

	case ProposalTypeTaxUsage:
		burn := false
		taxUsage := proposal.ProposalContent.(*TaxUsageProposal).TaxUsage

		if taxUsage.Usage == UsageTypeBurn {
			burn = true
		}
		err := keeper.distributionKeeper.AllocateCommunityPool(ctx, taxUsage.DestAddress, taxUsage.Percent, burn)
		if err != nil {
			return err
		}
	}

	return nil
}

func (keeper Keeper) SetProposalParam(ctx sdk.Context, proposalParam ProposalParam) sdk.Error {

	if strings.HasPrefix(proposalParam.Key, BoxModule) || strings.HasPrefix(proposalParam.Key, IssueModule) {
		subspaceKey := GetParamSpaceFromKey(proposalParam.Key)
		subspaceSet, _ := keeper.paramsKeeper.GetSubspace(subspaceKey)
		key := CamelString(GetParamKey(proposalParam.Key))
		if strings.HasSuffix(key, Fee) {
			coin, _ := sdk.ParseCoin(proposalParam.Value)
			subspaceSet.Set(ctx, []byte(key), &coin)
		} else {
			subspaceSet.Set(ctx, []byte(key), &proposalParam.Value)
		}
		return nil
	}

	switch proposalParam.Key {
	case communityTax:
		val, err := sdk.NewDecFromStr(proposalParam.Value)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
		keeper.distributionKeeper.SetCommunityTax(ctx, val)
		return nil

	case minDeposit:
		coins, err := sdk.ParseCoins(proposalParam.Value)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
		tParams := keeper.GetDepositParams(ctx)
		tParams.MinDeposit = coins
		keeper.setDepositParams(ctx, tParams)
		return nil

	case inflation:
		val, err := sdk.NewDecFromStr(proposalParam.Value)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
		tParams := keeper.mintKeeper.GetParams(ctx)
		tParams.Inflation = val
		keeper.mintKeeper.SetParams(ctx, tParams)
		return nil

	case inflationBase:
		val, ok := sdk.NewIntFromString(proposalParam.Value)
		if !ok {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, "invalid string to Int")
		}
		tParams := keeper.mintKeeper.GetParams(ctx)
		tParams.InflationBase = val
		keeper.mintKeeper.SetParams(ctx, tParams)
		return nil

	case signedBlocksWindow:
		val, err := strconv.ParseInt(proposalParam.Value, 10, 64)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
		tParams := keeper.slashingKeeper.GetParams(ctx)
		tParams.SignedBlocksWindow = val
		keeper.slashingKeeper.SetParams(ctx, tParams)
		return nil

	case minSignedPerWindow:
		val, err := sdk.NewDecFromStr(proposalParam.Value)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
		tParams := keeper.slashingKeeper.GetParams(ctx)
		tParams.MinSignedPerWindow = val
		keeper.slashingKeeper.SetParams(ctx, tParams)
		return nil

	case downtimeJailDuration:
		val, err := time.ParseDuration(proposalParam.Value)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
		tParams := keeper.slashingKeeper.GetParams(ctx)
		tParams.DowntimeJailDuration = val
		keeper.slashingKeeper.SetParams(ctx, tParams)
		return nil

	case slashFractionDowntime:
		val, err := sdk.NewDecFromStr(proposalParam.Value)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
		tParams := keeper.slashingKeeper.GetParams(ctx)
		tParams.SlashFractionDowntime = val
		keeper.slashingKeeper.SetParams(ctx, tParams)
		return nil

	case unbondingTime:
		val, err := time.ParseDuration(proposalParam.Value)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
		tParams := keeper.stakingKeeper.GetParams(ctx)
		tParams.UnbondingTime = val
		keeper.stakingKeeper.SetParams(ctx, tParams)
		return nil

	case maxValidators:
		val, err := strconv.ParseUint(proposalParam.Value, 10, 16)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
		tParams := keeper.stakingKeeper.GetParams(ctx)
		tParams.MaxValidators = uint16(val)
		keeper.stakingKeeper.SetParams(ctx, tParams)
		return nil

	case foundationAddress:
		val, err := sdk.AccAddressFromBech32(proposalParam.Value)
		if err != nil {
			return ErrInvalidParamValue(DefaultCodespace, proposalParam.Key, proposalParam.Value, err.Error())
		}
		tParams := keeper.mintKeeper.GetParams(ctx)
		tParams.FoundationAddress = val
		keeper.mintKeeper.SetParams(ctx, tParams)
		return nil

	default:
		return ErrInvalidParamKey(DefaultCodespace, proposalParam.Key)
	}
}
