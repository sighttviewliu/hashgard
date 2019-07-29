package app

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashgard/hashgard/x/record"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/hashgard/hashgard/x/box"

	"github.com/hashgard/hashgard/x/issue"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	acc "github.com/cosmos/cosmos-sdk/x/account"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/bank"
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/hashgard/hashgard/x/distribution"
	"github.com/hashgard/hashgard/x/exchange"
	"github.com/hashgard/hashgard/x/gov"
	"github.com/hashgard/hashgard/x/mint"
)

var (
	// bonded tokens given to genesis validators/accounts
	FreeFermionsAcc                    = sdk.NewIntWithDecimal(150, 18)
	defaultUnbondingTime time.Duration = 60 * 10 * time.Second

	StakeDenom = "agard"
)

// State to Unmarshal
type GenesisState struct {
	Accounts         []GenesisAccount          `json:"accounts"`
	AccMustMemoData  acc.GenesisState          `json:"must-memo-addresses"`
	AuthData         auth.GenesisState         `json:"auth"`
	BankData         bank.GenesisState         `json:"bank"`
	StakingData      staking.GenesisState      `json:"staking"`
	MintData         mint.GenesisState         `json:"mint"`
	DistributionData distribution.GenesisState `json:"distribution"`
	SlashingData     slashing.GenesisState     `json:"slashing"`
	GovData          gov.GenesisState          `json:"gov"`
	ExchangeData     exchange.GenesisState     `json:"exchange"`
	IssueData        issue.GenesisState        `json:"issue"`
	RecordData       record.GenesisState       `json:"record"`
	BoxData          box.GenesisState          `json:"box"`
	CrisisData       crisis.GenesisState       `json:"crisis"`
	GenTxs           []json.RawMessage         `json:"gentxs"`
}

func NewGenesisState(
	accounts []GenesisAccount,
	accMustMemoData acc.GenesisState,
	authData auth.GenesisState,
	bankData bank.GenesisState,
	stakingData staking.GenesisState,
	mintData mint.GenesisState,
	distributionData distribution.GenesisState,
	govData gov.GenesisState,
	slashingData slashing.GenesisState,
	exchangeData exchange.GenesisState,
	issueData issue.GenesisState,
	recordData record.GenesisState,
	boxData box.GenesisState,
	crisisData crisis.GenesisState,
) GenesisState {

	return GenesisState{
		Accounts:         accounts,
		AccMustMemoData:  accMustMemoData,
		AuthData:         authData,
		BankData:         bankData,
		StakingData:      stakingData,
		MintData:         mintData,
		DistributionData: distributionData,
		GovData:          govData,
		SlashingData:     slashingData,
		IssueData:        issueData,
		RecordData:       recordData,
		BoxData:          boxData,
		ExchangeData:     exchangeData,
		CrisisData:       crisisData,
	}
}

// Sanitize sorts accounts and coin sets.
func (gs GenesisState) Sanitize() {
	sort.Slice(gs.Accounts, func(i, j int) bool {
		return gs.Accounts[i].AccountNumber < gs.Accounts[j].AccountNumber
	})

	for _, acc := range gs.Accounts {
		acc.Coins = acc.Coins.Sort()
	}
}

// NewDefaultGenesisState generates the default state for hashgard.
func NewDefaultGenesisState() GenesisState {
	return GenesisState{
		Accounts:         nil,
		AccMustMemoData:  acc.DefaultGenesisState(),
		AuthData:         auth.DefaultGenesisState(),
		BankData:         bank.DefaultGenesisState(),
		StakingData:      createStakingGenesisState(),
		MintData:         mint.DefaultGenesisState(),
		DistributionData: distribution.DefaultGenesisState(),
		GovData:          createGovGenesisState(),
		SlashingData:     slashing.DefaultGenesisState(),
		ExchangeData:     exchange.DefaultGenesisState(),
		IssueData:        createIssueGenesisState(),
		RecordData:       createRecordGenesisState(),
		BoxData:          createBoxGenesisState(),
		CrisisData:       createCrisisGenesisState(),
		GenTxs:           nil,
	}
}

func createStakingGenesisState() staking.GenesisState {
	return staking.GenesisState{
		Pool: staking.Pool{
			NotBondedTokens: sdk.ZeroInt(),
			BondedTokens:    sdk.ZeroInt(),
		},
		Params: staking.Params{
			UnbondingTime: defaultUnbondingTime,
			MaxValidators: 100,
			MaxEntries:    7,
			BondDenom:     StakeDenom,
		},
	}
}

func createGovGenesisState() gov.GenesisState {
	return gov.GenesisState{
		StartingProposalID: 1,
		DepositParams: gov.DepositParams{
			MinDeposit:       sdk.NewCoins(sdk.NewCoin(StakeDenom, sdk.NewIntWithDecimal(10, 18))),
			MaxDepositPeriod: time.Duration(172800) * time.Second,
		},
		VotingParams: gov.VotingParams{
			VotingPeriod: time.Duration(172800) * time.Second,
		},
		TallyParams: gov.TallyParams{
			Quorum:    sdk.NewDecWithPrec(334, 3),
			Threshold: sdk.NewDecWithPrec(5, 1),
			Veto:      sdk.NewDecWithPrec(334, 3),
		},
	}
}

func createCrisisGenesisState() crisis.GenesisState {
	return crisis.GenesisState{
		ConstantFee: sdk.NewCoin(StakeDenom, sdk.NewIntWithDecimal(1000, 18)),
	}
}

func createBoxGenesisState() box.GenesisState {
	genesisState := box.DefaultGenesisState()
	genesisState.Params = box.DefaultParams(StakeDenom)
	return genesisState
}
func createIssueGenesisState() issue.GenesisState {
	genesisState := issue.DefaultGenesisState()
	genesisState.Params = issue.DefaultParams(StakeDenom)
	return genesisState
}
func createRecordGenesisState() record.GenesisState {
	genesisState := record.DefaultGenesisState()
	return genesisState
}

// nolint
type GenesisAccount struct {
	Address       sdk.AccAddress `json:"address"`
	Coins         sdk.Coins      `json:"coins"`
	Sequence      uint64         `json:"sequence_number"`
	AccountNumber uint64         `json:"account_number"`

	// vesting account fields
	OriginalVesting  sdk.Coins `json:"original_vesting"`  // total vesting coins upon initialization
	DelegatedFree    sdk.Coins `json:"delegated_free"`    // delegated vested coins at time of delegation
	DelegatedVesting sdk.Coins `json:"delegated_vesting"` // delegated vesting coins at time of delegation
	StartTime        int64     `json:"start_time"`        // vesting start time (UNIX Epoch time)
	EndTime          int64     `json:"end_time"`          // vesting end time (UNIX Epoch time)
}

func NewGenesisAccount(acc *auth.BaseAccount) GenesisAccount {
	return GenesisAccount{
		Address:       acc.Address,
		Coins:         acc.Coins,
		AccountNumber: acc.AccountNumber,
		Sequence:      acc.Sequence,
	}
}

func NewGenesisAccountI(acc auth.Account) GenesisAccount {
	gacc := GenesisAccount{
		Address:       acc.GetAddress(),
		Coins:         acc.GetCoins(),
		AccountNumber: acc.GetAccountNumber(),
		Sequence:      acc.GetSequence(),
	}

	vacc, ok := acc.(auth.VestingAccount)
	if ok {
		gacc.OriginalVesting = vacc.GetOriginalVesting()
		gacc.DelegatedFree = vacc.GetDelegatedFree()
		gacc.DelegatedVesting = vacc.GetDelegatedVesting()
		gacc.StartTime = vacc.GetStartTime()
		gacc.EndTime = vacc.GetEndTime()
	}

	return gacc
}

// convert GenesisAccount to auth.BaseAccount
func (ga *GenesisAccount) ToAccount() auth.Account {
	bacc := &auth.BaseAccount{
		Address:       ga.Address,
		Coins:         ga.Coins.Sort(),
		AccountNumber: ga.AccountNumber,
		Sequence:      ga.Sequence,
	}

	if !ga.OriginalVesting.IsZero() {
		baseVestingAcc := &auth.BaseVestingAccount{
			BaseAccount:      bacc,
			OriginalVesting:  ga.OriginalVesting,
			DelegatedFree:    ga.DelegatedFree,
			DelegatedVesting: ga.DelegatedVesting,
			EndTime:          ga.EndTime,
		}

		if ga.StartTime != 0 && ga.EndTime != 0 {
			return &auth.ContinuousVestingAccount{
				BaseVestingAccount: baseVestingAcc,
				StartTime:          ga.StartTime,
			}
		} else if ga.EndTime != 0 {
			return &auth.DelayedVestingAccount{
				BaseVestingAccount: baseVestingAcc,
			}
		} else {
			panic(fmt.Sprintf("invalid genesis vesting account: %+v", ga))
		}
	}

	return bacc
}

func NewDefaultGenesisAccount(addr sdk.AccAddress) GenesisAccount {
	accAuth := auth.NewBaseAccountWithAddress(addr)
	coins := sdk.NewCoins(
		sdk.NewCoin(StakeDenom, FreeFermionsAcc),
	)

	coins.Sort()

	accAuth.Coins = coins
	return NewGenesisAccount(&accAuth)
}

// HashgardAppGenState but with JSON
func HashgardAppGenStateJSON(cdc *codec.Codec, genDoc tmtypes.GenesisDoc, appGenTxs []json.RawMessage) (appState json.RawMessage, err error) {

	// create the final app state
	genesisState, err := HashgardAppGenState(cdc, genDoc, appGenTxs)
	if err != nil {
		return nil, err
	}
	return codec.MarshalJSONIndent(cdc, genesisState)
}

// Create the core parameters for genesis initialization for hashgard
// note that the pubkey input is this machines pubkey
func HashgardAppGenState(cdc *codec.Codec, genDoc tmtypes.GenesisDoc, appGenTxs []json.RawMessage) (genesisState GenesisState, err error) {

	if err = cdc.UnmarshalJSON(genDoc.AppState, &genesisState); err != nil {
		return genesisState, err
	}

	// if there are no gen txs to be processed, return the default empty state
	if len(appGenTxs) == 0 {
		return genesisState, errors.New("there must be at least one genesis tx")
	}

	stakingData := genesisState.StakingData
	for i, genTx := range appGenTxs {
		var tx auth.StdTx
		if err := cdc.UnmarshalJSON(genTx, &tx); err != nil {
			return genesisState, err
		}
		msgs := tx.GetMsgs()
		if len(msgs) != 1 {
			return genesisState, errors.New(
				"must provide genesis StdTx with exactly 1 CreateValidator message")
		}
		if _, ok := msgs[0].(staking.MsgCreateValidator); !ok {
			return genesisState, fmt.Errorf(
				"Genesis transaction %v does not contain a MsgCreateValidator", i)
		}
	}

	for _, acc := range genesisState.Accounts {
		// create the genesis account, give'm few steaks and a buncha token with there name
		for _, coin := range acc.Coins {
			if coin.Denom == genesisState.StakingData.Params.BondDenom {
				stakingData.Pool.NotBondedTokens = stakingData.Pool.NotBondedTokens.
					Add(coin.Amount) // increase the supply
			}
		}
	}

	genesisState.StakingData = stakingData
	genesisState.GenTxs = appGenTxs
	return genesisState, nil
}

// HashgardValidateGenesisState ensures that the genesis state obeys the expected invariants
// TODO: No validators are both bonded and jailed (#2088)
// TODO: Error if there is a duplicate validator (#1708)
// TODO: Ensure all state machine parameters are in genesis (#1704)
func HashgardValidateGenesisState(genesisState GenesisState) error {
	if err := validateGenesisStateAccounts(genesisState.Accounts); err != nil {
		return err
	}

	// skip stakingData validation as genesis is created from txs
	if len(genesisState.GenTxs) > 0 {
		return nil
	}

	if err := acc.ValidateGenesis(genesisState.AccMustMemoData); err != nil {
		return err
	}
	if err := auth.ValidateGenesis(genesisState.AuthData); err != nil {
		return err
	}
	if err := bank.ValidateGenesis(genesisState.BankData); err != nil {
		return err
	}
	if err := staking.ValidateGenesis(genesisState.StakingData); err != nil {
		return err
	}
	if err := mint.ValidateGenesis(genesisState.MintData); err != nil {
		return err
	}
	if err := distribution.ValidateGenesis(genesisState.DistributionData); err != nil {
		return err
	}
	if err := gov.ValidateGenesis(genesisState.GovData); err != nil {
		return err
	}
	if err := issue.ValidateGenesis(genesisState.IssueData); err != nil {
		return err
	}
	if err := record.ValidateGenesis(genesisState.RecordData); err != nil {
		return err
	}
	if err := box.ValidateGenesis(genesisState.BoxData); err != nil {
		return err
	}
	if err := exchange.ValidateGenesis(genesisState.ExchangeData); err != nil {
		return err
	}
	if err := crisis.ValidateGenesis(genesisState.CrisisData); err != nil {
		return err
	}

	return slashing.ValidateGenesis(genesisState.SlashingData)
}

// validateGenesisStateAccounts performs validation of genesis accounts. It
// ensures that there are no duplicate accounts in the genesis state and any
// provided vesting accounts are valid.
func validateGenesisStateAccounts(accs []GenesisAccount) (err error) {
	addrMap := make(map[string]bool, len(accs))
	for _, acc := range accs {
		addrStr := acc.Address.String()

		// disallow any duplicate accounts
		if _, ok := addrMap[addrStr]; ok {
			return fmt.Errorf("duplicate account found in genesis state; address: %s", addrStr)
		}

		// validate any vesting fields
		if !acc.OriginalVesting.IsZero() {
			if acc.EndTime == 0 {
				return fmt.Errorf("missing end time for vesting account; address: %s", addrStr)
			}

			if acc.StartTime >= acc.EndTime {
				return fmt.Errorf(
					"vesting start time must before end time; address: %s, start: %s, end: %s",
					addrStr,
					time.Unix(acc.StartTime, 0).UTC().Format(time.RFC3339),
					time.Unix(acc.EndTime, 0).UTC().Format(time.RFC3339),
				)
			}
		}

		addrMap[addrStr] = true
	}

	return nil
}

// CollectStdTxs processes and validates application's genesis StdTxs and returns
// the list of appGenTxs, and persistent peers required to generate genesis.json.
func CollectStdTxs(cdc *codec.Codec, moniker string, genTxsDir string, genDoc tmtypes.GenesisDoc) (
	appGenTxs []auth.StdTx, persistentPeers string, err error) {

	var fos []os.FileInfo
	fos, err = ioutil.ReadDir(genTxsDir)
	if err != nil {
		return appGenTxs, persistentPeers, err
	}

	// prepare a map of all accounts in genesis state to then validate
	// against the validators addresses
	var appState GenesisState
	if err := cdc.UnmarshalJSON(genDoc.AppState, &appState); err != nil {
		return appGenTxs, persistentPeers, err
	}

	addrMap := make(map[string]GenesisAccount, len(appState.Accounts))
	for i := 0; i < len(appState.Accounts); i++ {
		acc := appState.Accounts[i]
		addrMap[acc.Address.String()] = acc
	}

	// addresses and IPs (and port) validator server info
	var addressesIPs []string

	for _, fo := range fos {
		filename := filepath.Join(genTxsDir, fo.Name())
		if !fo.IsDir() && (filepath.Ext(filename) != ".json") {
			continue
		}

		// get the genStdTx
		var jsonRawTx []byte
		if jsonRawTx, err = ioutil.ReadFile(filename); err != nil {
			return appGenTxs, persistentPeers, err
		}
		var genStdTx auth.StdTx
		if err = cdc.UnmarshalJSON(jsonRawTx, &genStdTx); err != nil {
			return appGenTxs, persistentPeers, err
		}
		appGenTxs = append(appGenTxs, genStdTx)

		// the memo flag is used to store
		// the ip and node-id, for example this may be:
		// "528fd3df22b31f4969b05652bfe8f0fe921321d5@192.168.2.37:26656"
		nodeAddrIP := genStdTx.GetMemo()
		if len(nodeAddrIP) == 0 {
			return appGenTxs, persistentPeers, fmt.Errorf(
				"couldn't find node's address and IP in %s", fo.Name())
		}

		// genesis transactions must be single-message
		msgs := genStdTx.GetMsgs()
		if len(msgs) != 1 {

			return appGenTxs, persistentPeers, errors.New(
				"each genesis transaction must provide a single genesis message")
		}

		msg := msgs[0].(staking.MsgCreateValidator)
		// validate delegator and validator addresses and funds against the accounts in the state
		delAddr := msg.DelegatorAddress.String()
		valAddr := sdk.AccAddress(msg.ValidatorAddress).String()

		delAcc, delOk := addrMap[delAddr]
		_, valOk := addrMap[valAddr]

		accsNotInGenesis := []string{}
		if !delOk {
			accsNotInGenesis = append(accsNotInGenesis, delAddr)
		}
		if !valOk {
			accsNotInGenesis = append(accsNotInGenesis, valAddr)
		}
		if len(accsNotInGenesis) != 0 {
			return appGenTxs, persistentPeers, fmt.Errorf(
				"account(s) %v not in genesis.json: %+v", strings.Join(accsNotInGenesis, " "), addrMap)
		}

		if delAcc.Coins.AmountOf(msg.Value.Denom).LT(msg.Value.Amount) {
			return appGenTxs, persistentPeers, fmt.Errorf(
				"insufficient fund for delegation %v: %v < %v",
				delAcc.Address, delAcc.Coins.AmountOf(msg.Value.Denom), msg.Value.Amount,
			)
		}

		// exclude itself from persistent peers
		if msg.Description.Moniker != moniker {
			addressesIPs = append(addressesIPs, nodeAddrIP)
		}
	}

	sort.Strings(addressesIPs)
	persistentPeers = strings.Join(addressesIPs, ",")

	return appGenTxs, persistentPeers, nil
}
