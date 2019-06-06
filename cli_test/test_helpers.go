package clitest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	cmn "github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/crypto"

	clientkeys "github.com/cosmos/cosmos-sdk/client/keys"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/hashgard/hashgard/app"
	hashgardInit "github.com/hashgard/hashgard/init"
	"github.com/hashgard/hashgard/x/exchange"
	"github.com/hashgard/hashgard/x/gov"
	"github.com/hashgard/hashgard/x/issue"
)

func init() {
	hashgardInit.InitBech32Prefix()
}

const (
	denom        = "agard"
	keyFoo       = "foo"
	keyBar       = "bar"
	fooDenom     = "footoken"
	feeDenom     = "feetoken"
	fee2Denom    = "fee2token"
	keyBaz       = "baz"
	keyVesting   = "vesting"
	keyFooBarBaz = "foobarbaz"
)

var (
	startCoins = sdk.Coins{
		sdk.NewCoin(feeDenom, sdk.TokensFromTendermintPower(1000000)),
		sdk.NewCoin(fee2Denom, sdk.TokensFromTendermintPower(1000000)),
		sdk.NewCoin(fooDenom, sdk.TokensFromTendermintPower(1000)),
		sdk.NewCoin(denom, sdk.TokensFromTendermintPower(150)),
	}

	vestingCoins = sdk.Coins{
		sdk.NewCoin(feeDenom, sdk.TokensFromTendermintPower(500000)),
	}
)

//___________________________________________________________________________________
// Fixtures

// Fixtures is used to setup the testing environment
type Fixtures struct {
	ChainID  string
	RPCAddr  string
	Port     string
	GDHome   string
	GCLIHome string
	P2PAddr  string
	T        *testing.T
}

// NewFixtures creates a new instance of Fixtures with many vars set
func NewFixtures(t *testing.T) *Fixtures {
	tmpDir, err := ioutil.TempDir("", "hashgard_integration_"+t.Name()+"_")
	require.NoError(t, err)

	servAddr, port, err := server.FreeTCPAddr()
	require.NoError(t, err)

	p2pAddr, _, err := server.FreeTCPAddr()
	require.NoError(t, err)

	return &Fixtures{
		T:        t,
		GDHome:   filepath.Join(tmpDir, ".hashgard"),
		GCLIHome: filepath.Join(tmpDir, ".hashgardcli"),
		RPCAddr:  servAddr,
		P2PAddr:  p2pAddr,
		Port:     port,
	}
}

// GenesisFile returns the path of the genesis file
func (f Fixtures) GenesisFile() string {
	return filepath.Join(f.GDHome, "config", "genesis.json")
}

// GenesisFile returns the application's genesis state
func (f Fixtures) GenesisState() app.GenesisState {
	cdc := codec.New()
	genDoc, err := hashgardInit.LoadGenesisDoc(cdc, f.GenesisFile())
	require.NoError(f.T, err)

	var appState app.GenesisState
	require.NoError(f.T, cdc.UnmarshalJSON(genDoc.AppState, &appState))
	return appState
}

// InitFixtures is called at the beginning of a test  and initializes a chain
// with 1 validator.
func InitFixtures(t *testing.T) (f *Fixtures) {
	f = NewFixtures(t)

	// reset test state
	f.UnsafeResetAll()

	// ensure keystore has foo and bar keys
	f.KeysDelete(keyFoo)
	f.KeysDelete(keyBar)
	f.KeysDelete(keyBar)
	f.KeysDelete(keyFooBarBaz)
	f.KeysAdd(keyFoo)
	f.KeysAdd(keyBar)
	f.KeysAdd(keyBaz)
	f.KeysAdd(keyVesting)
	f.KeysAdd(keyFooBarBaz, "--multisig-threshold=2", fmt.Sprintf(
		"--multisig=%s,%s,%s", keyFoo, keyBar, keyBaz))

	// ensure that CLI output is in JSON format
	f.CLIConfig("output", "json")

	// NOTE: HGInit sets the ChainID
	f.HGInit(keyFoo)

	f.CLIConfig("chain-id", f.ChainID)
	f.CLIConfig("broadcast-mode", "block")

	// start an account with tokens
	f.AddGenesisAccount(f.KeyAddress(keyFoo), startCoins)
	f.AddGenesisAccount(
		f.KeyAddress(keyVesting), startCoins,
		fmt.Sprintf("--vesting-amount=%s", vestingCoins),
		fmt.Sprintf("--vesting-start-time=%d", time.Now().UTC().UnixNano()),
		fmt.Sprintf("--vesting-end-time=%d", time.Now().Add(60*time.Second).UTC().UnixNano()),
	)

	f.AddGenesisAccount(sdk.AccAddress(crypto.AddressHash([]byte("foundationaddress"))), sdk.NewCoins(sdk.NewCoin(denom, sdk.TokensFromTendermintPower(300000000))))

	f.GenTx(keyFoo)
	f.CollectGenTxs()

	return
}

// Cleanup is meant to be run at the end of a test to clean up an remaining test state
func (f *Fixtures) Cleanup(dirs ...string) {
	clean := append(dirs, f.GDHome, f.GCLIHome)
	for _, d := range clean {
		err := os.RemoveAll(d)
		require.NoError(f.T, err)
	}
}

// Flags returns the flags necessary for making most CLI calls
func (f *Fixtures) Flags() string {
	return fmt.Sprintf("--home=%s --node=%s", f.GCLIHome, f.RPCAddr)
}

//___________________________________________________________________________________
// hashgard

// UnsafeResetAll is hashgard unsafe-reset-all
func (f *Fixtures) UnsafeResetAll(flags ...string) {
	cmd := fmt.Sprintf("../build/hashgard --home=%s unsafe-reset-all", f.GDHome)
	executeWrite(f.T, addFlags(cmd, flags))
	err := os.RemoveAll(filepath.Join(f.GDHome, "config", "gentx"))
	require.NoError(f.T, err)
}

// HGInit is hashgard init
// NOTE: HGInit sets the ChainID for the Fixtures instance
func (f *Fixtures) HGInit(moniker string, flags ...string) {
	cmd := fmt.Sprintf("../build/hashgard init -o --home=%s --moniker=%s", f.GDHome, moniker)
	_, stderr := tests.ExecuteT(f.T, addFlags(cmd, flags), app.DefaultKeyPass)

	var chainID string
	var initRes map[string]json.RawMessage

	err := json.Unmarshal([]byte(stderr), &initRes)
	require.NoError(f.T, err)

	err = json.Unmarshal(initRes["chain_id"], &chainID)
	require.NoError(f.T, err)

	f.ChainID = chainID
	file := path.Join(f.GDHome, "config", "config.toml")
	content, err := ioutil.ReadFile(file)
	require.NoError(f.T, err)
	str := strings.ReplaceAll(string(content), "timeout_commit = \"5s\"", "timeout_commit = \"1s\"")
	err = ioutil.WriteFile(file, []byte(str), os.ModeDir)
	require.NoError(f.T, err)
}

// AddGenesisAccount is hashgard add-genesis-account
func (f *Fixtures) AddGenesisAccount(address sdk.AccAddress, coins sdk.Coins, flags ...string) {
	cmd := fmt.Sprintf("../build/hashgard add-genesis-account %s %s --home=%s", address, coins, f.GDHome)
	executeWriteCheckErr(f.T, addFlags(cmd, flags))
}

// GenTx is hashgard gentx
func (f *Fixtures) GenTx(name string, flags ...string) {
	cmd := fmt.Sprintf("../build/hashgard gentx --name=%s --home=%s --home-client=%s", name, f.GDHome, f.GCLIHome)
	executeWriteCheckErr(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// CollectGenTxs is hashgard collect-gentxs
func (f *Fixtures) CollectGenTxs(flags ...string) {
	cmd := fmt.Sprintf("../build/hashgard collect-gentxs --home=%s", f.GDHome)
	executeWriteCheckErr(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// HGStart runs hashgard start with the appropriate flags and returns a process
func (f *Fixtures) HGStart(flags ...string) *tests.Process {
	cmd := fmt.Sprintf("../build/hashgard start --home=%s --rpc.laddr=%v --p2p.laddr=%v", f.GDHome, f.RPCAddr, f.P2PAddr)
	proc := tests.GoExecuteTWithStdout(f.T, addFlags(cmd, flags))
	tests.WaitForTMStart(f.Port)
	tests.WaitForNextNBlocksTM(1, f.Port)
	return proc
}

// HGTendermint returns the results of hashgard tendermint [query]
func (f *Fixtures) HGTendermint(query string) string {
	cmd := fmt.Sprintf("../build/hashgard tendermint %s --home=%s", query, f.GDHome)
	success, stdout, stderr := executeWriteRetStdStreams(f.T, cmd)
	require.Empty(f.T, stderr)
	require.True(f.T, success)
	return strings.TrimSpace(stdout)
}

// ValidateGenesis runs hashgard validate-genesis
func (f *Fixtures) ValidateGenesis() {
	cmd := fmt.Sprintf("../build/hashgard validate-genesis --home=%s", f.GDHome)
	executeWriteCheckErr(f.T, cmd)
}

//___________________________________________________________________________________
// hashgardcli keys

// KeysDelete is hashgardcli keys delete
func (f *Fixtures) KeysDelete(name string, flags ...string) {
	cmd := fmt.Sprintf("../build/hashgardcli keys delete --home=%s %s", f.GCLIHome, name)
	executeWrite(f.T, addFlags(cmd, append(append(flags, "-y"), "-f")))
}

// KeysAdd is hashgardcli keys add
func (f *Fixtures) KeysAdd(name string, flags ...string) {
	cmd := fmt.Sprintf("../build/hashgardcli keys add --home=%s %s", f.GCLIHome, name)
	executeWriteCheckErr(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// KeysAddRecover prepares hashgardcli keys add --recover
func (f *Fixtures) KeysAddRecover(name, mnemonic string, flags ...string) (exitSuccess bool, stdout, stderr string) {
	cmd := fmt.Sprintf("../build/hashgardcli keys add --home=%s --recover %s", f.GCLIHome, name)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass, mnemonic)
}

// KeysAddRecoverHDPath prepares hashgardcli keys add --recover --account --index
func (f *Fixtures) KeysAddRecoverHDPath(name, mnemonic string, account uint32, index uint32, flags ...string) {
	cmd := fmt.Sprintf("../build/hashgardcli keys add --home=%s --recover %s --account %d --index %d", f.GCLIHome, name, account, index)
	executeWriteCheckErr(f.T, addFlags(cmd, flags), app.DefaultKeyPass, mnemonic)
}

// KeysShow is hashgardcli keys show
func (f *Fixtures) KeysShow(name string, flags ...string) keys.KeyOutput {
	cmd := fmt.Sprintf("../build/hashgardcli keys show --home=%s %s", f.GCLIHome, name)
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var ko keys.KeyOutput
	err := clientkeys.UnmarshalJSON([]byte(out), &ko)
	require.NoError(f.T, err)
	return ko
}

// KeyAddress returns the SDK account address from the key
func (f *Fixtures) KeyAddress(name string) sdk.AccAddress {
	ko := f.KeysShow(name)
	accAddr, err := sdk.AccAddressFromBech32(ko.Address)
	require.NoError(f.T, err)
	return accAddr
}

//___________________________________________________________________________________
// hashgardcli config

// CLIConfig is hashgardcli config
func (f *Fixtures) CLIConfig(key, value string, flags ...string) {
	cmd := fmt.Sprintf("../build/hashgardcli config --home=%s %s %s", f.GCLIHome, key, value)
	executeWriteCheckErr(f.T, addFlags(cmd, flags))
}

//___________________________________________________________________________________
// hashgardcli tx send/sign/broadcast

// TxSend is hashgardcli bank send
func (f *Fixtures) TxSend(from string, to sdk.AccAddress, amount sdk.Coin, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli bank send %s %s %v --from=%s", to, amount, f.Flags(), from)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

func (f *Fixtures) txSendWithConfirm(
	from string, to sdk.AccAddress, amount sdk.Coin, confirm string, flags ...string,
) (bool, string, string) {

	cmd := fmt.Sprintf("../build/hashgardcli bank send %s %s %v --from=%s", to, amount, f.Flags(), from)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), confirm, app.DefaultKeyPass)
}

// TxSign is hashgardcli bank sign
func (f *Fixtures) TxSign(signer, fileName string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli bank sign %v --from=%s %v", f.Flags(), signer, fileName)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxBroadcast is hashgardcli bank broadcast
func (f *Fixtures) TxBroadcast(fileName string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli bank broadcast %v %v", f.Flags(), fileName)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxEncode is hashgardcli bank encode
func (f *Fixtures) TxEncode(fileName string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli bank encode %v %v", f.Flags(), fileName)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxMultisign is hashgardcli bank multisign
func (f *Fixtures) TxMultisign(fileName, name string, signaturesFiles []string,
	flags ...string) (bool, string, string) {

	cmd := fmt.Sprintf("../build/hashgardcli bank multisign %v %s %s %s", f.Flags(),
		fileName, name, strings.Join(signaturesFiles, " "),
	)
	return executeWriteRetStdStreams(f.T, cmd)
}

//___________________________________________________________________________________
// hashgardcli tx stake

// TxStakingCreateValidator is hashgardcli stake create-validator
func (f *Fixtures) TxStakingCreateValidator(from, consPubKey string, amount sdk.Coin, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli stake create-validator %v --from=%s --pubkey=%s", f.Flags(), from, consPubKey)
	cmd += fmt.Sprintf(" --amount=%v --moniker=%v --commission-rate=%v", amount, from, "0.05")
	cmd += fmt.Sprintf(" --commission-max-rate=%v --commission-max-change-rate=%v", "0.20", "0.10")
	cmd += fmt.Sprintf(" --min-self-delegation=%v", "1")
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxStakingUnbond is hashgardcli stake unbond
func (f *Fixtures) TxStakingUnbond(from, shares string, validator sdk.ValAddress, flags ...string) bool {
	cmd := fmt.Sprintf("../build/hashgardcli stake unbond %s %v --from=%s %v", validator, shares, from, f.Flags())
	return executeWrite(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

//___________________________________________________________________________________
// hashgardcli tx gov

// TxGovSubmitProposal is hashgardcli gov submit-proposal
func (f *Fixtures) TxGovSubmitProposal(from, typ, title, description string, deposit sdk.Coin, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli gov submit-proposal %v --from=%s --type=%s", f.Flags(), from, typ)
	cmd += fmt.Sprintf(" --title=%s --description=%s --deposit=%s", title, description, deposit)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxGovDeposit is hashgardcli gov deposit
func (f *Fixtures) TxGovDeposit(proposalID int, from string, amount sdk.Coin, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli gov deposit %d %s --from=%s %v", proposalID, amount, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxGovVote is hashgardcli gov vote
func (f *Fixtures) TxGovVote(proposalID int, option gov.VoteOption, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli gov vote %d %s --from=%s %v", proposalID, option, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

//___________________________________________________________________________________
// hashgardcli query account

// QueryAccount is hashgardcli bank account
func (f *Fixtures) QueryAccount(address sdk.AccAddress, flags ...string) auth.BaseAccount {
	cmd := fmt.Sprintf("../build/hashgardcli bank account %s %v", address, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var initRes map[string]json.RawMessage
	err := json.Unmarshal([]byte(out), &initRes)
	require.NoError(f.T, err, "out %v, err %v", out, err)
	value := initRes["value"]
	var acc auth.BaseAccount
	cdc := codec.New()
	codec.RegisterCrypto(cdc)
	err = cdc.UnmarshalJSON(value, &acc)
	require.NoError(f.T, err, "value %v, err %v", string(value), err)
	return acc
}

//___________________________________________________________________________________
// hashgardcli query txs

// QueryTxs is hashgardcli tendermint txs
func (f *Fixtures) QueryTxs(page, limit int, tags ...string) []sdk.TxResponse {
	cmd := fmt.Sprintf("../build/hashgardcli tendermint txs --page=%d --limit=%d --tags='%s' %v", page, limit, queryTags(tags), f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	var txs []sdk.TxResponse
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &txs)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return txs
}

// QueryTxsInvalid query txs with wrong parameters and compare expected error
func (f *Fixtures) QueryTxsInvalid(expectedErr error, page, limit int, tags ...string) {
	cmd := fmt.Sprintf("../build/hashgardcli tendermint txs --page=%d --limit=%d --tags='%s' %v", page, limit, queryTags(tags), f.Flags())
	_, err := tests.ExecuteT(f.T, cmd, "")
	require.EqualError(f.T, expectedErr, err)
}

//___________________________________________________________________________________
// hashgardcli stake

// QueryStakingValidator is hashgardcli stake validator
func (f *Fixtures) QueryStakingValidator(valAddr sdk.ValAddress, flags ...string) staking.Validator {
	cmd := fmt.Sprintf("../build/hashgardcli stake validator %s %v", valAddr, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var validator staking.Validator
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &validator)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return validator
}

// QueryStakingUnbondingDelegationsFrom is hashgardcli stake unbonding-delegations-from
func (f *Fixtures) QueryStakingUnbondingDelegationsFrom(valAddr sdk.ValAddress, flags ...string) []staking.UnbondingDelegation {
	cmd := fmt.Sprintf("../build/hashgardcli stake unbonding-delegations-from %s %v", valAddr, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var ubds []staking.UnbondingDelegation
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &ubds)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return ubds
}

// QueryStakingDelegationsTo is hashgardcli stake delegations-to
func (f *Fixtures) QueryStakingDelegationsTo(valAddr sdk.ValAddress, flags ...string) []staking.Delegation {
	cmd := fmt.Sprintf("../build/hashgardcli stake delegations-to %s %v", valAddr, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var delegations []staking.Delegation
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &delegations)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return delegations
}

// QueryStakingPool is hashgardcli stake pool
func (f *Fixtures) QueryStakingPool(flags ...string) staking.Pool {
	cmd := fmt.Sprintf("../build/hashgardcli stake pool %v", f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var pool staking.Pool
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &pool)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return pool
}

// QueryStakingParameters is hashgardcli stake parameters
func (f *Fixtures) QueryStakingParameters(flags ...string) staking.Params {
	cmd := fmt.Sprintf("../build/hashgardcli stake params %v", f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var params staking.Params
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &params)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return params
}

//___________________________________________________________________________________
// hashgardcli query gov

// QueryGovParamDeposit is hashgardcli gov param deposit
func (f *Fixtures) QueryGovParamDeposit() gov.DepositParams {
	cmd := fmt.Sprintf("../build/hashgardcli gov param deposit %s", f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	var depositParam gov.DepositParams
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &depositParam)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return depositParam
}

// QueryGovParamVoting is hashgardcli gov param voting
func (f *Fixtures) QueryGovParamVoting() gov.VotingParams {
	cmd := fmt.Sprintf("../build/hashgardcli gov param voting %s", f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	var votingParam gov.VotingParams
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &votingParam)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return votingParam
}

// QueryGovParamTallying is hashgardcli gov param tallying
func (f *Fixtures) QueryGovParamTallying() gov.TallyParams {
	cmd := fmt.Sprintf("../build/hashgardcli gov param tallying %s", f.Flags())
	out, _ := tests.ExecuteT(f.T, cmd, "")
	var tallyingParam gov.TallyParams
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &tallyingParam)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return tallyingParam
}

// QueryGovProposals is hashgardcli gov proposals
func (f *Fixtures) QueryGovProposals(flags ...string) gov.Proposals {
	cmd := fmt.Sprintf("../build/hashgardcli gov proposals %v", f.Flags())
	stdout, stderr := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	if strings.Contains(stderr, "No matching proposals found") {
		return gov.Proposals{}
	}
	require.Empty(f.T, stderr)
	var out gov.Proposals
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(stdout), &out)
	require.NoError(f.T, err)
	return out
}

// QueryGovProposal is hashgardcli gov proposal
func (f *Fixtures) QueryGovProposal(proposalID int, flags ...string) gov.Proposal {
	cmd := fmt.Sprintf("../build/hashgardcli gov proposal %d %v", proposalID, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var proposal gov.Proposal
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &proposal)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return proposal
}

// QueryGovVote is hashgardcli gov query-vote
func (f *Fixtures) QueryGovVote(proposalID int, voter sdk.AccAddress, flags ...string) gov.Vote {
	cmd := fmt.Sprintf("../build/hashgardcli gov query-vote %d %s %v", proposalID, voter, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var vote gov.Vote
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &vote)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return vote
}

// QueryGovVotes is hashgardcli gov query-votes
func (f *Fixtures) QueryGovVotes(proposalID int, flags ...string) []gov.Vote {
	cmd := fmt.Sprintf("../build/hashgardcli gov query-votes %d %v", proposalID, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var votes []gov.Vote
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &votes)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return votes
}

// QueryGovDeposit is hashgardcli gov query-deposit
func (f *Fixtures) QueryGovDeposit(proposalID int, depositor sdk.AccAddress, flags ...string) gov.Deposit {
	cmd := fmt.Sprintf("../build/hashgardcli gov query-deposit %d %s %v", proposalID, depositor, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var deposit gov.Deposit
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &deposit)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return deposit
}

// QueryGovDeposits is hashgardcli gov query-deposits
func (f *Fixtures) QueryGovDeposits(propsalID int, flags ...string) []gov.Deposit {
	cmd := fmt.Sprintf("../build/hashgardcli gov query-deposits %d %v", propsalID, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var deposits []gov.Deposit
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &deposits)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return deposits
}

//___________________________________________________________________________________
// query slashing

// QuerySigningInfo returns the signing info for a validator
func (f *Fixtures) QuerySigningInfo(val string) slashing.ValidatorSigningInfo {
	cmd := fmt.Sprintf("../build/hashgardcli slashing signing-info %s %s", val, f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var sinfo slashing.ValidatorSigningInfo
	err := cdc.UnmarshalJSON([]byte(res), &sinfo)
	require.NoError(f.T, err)
	return sinfo
}

// QuerySlashingParams is hashgardcli slashing params
func (f *Fixtures) QuerySlashingParams() slashing.Params {
	cmd := fmt.Sprintf("../build/hashgardcli slashing params %s", f.Flags())
	res, errStr := tests.ExecuteT(f.T, cmd, "")
	require.Empty(f.T, errStr)
	cdc := app.MakeCodec()
	var params slashing.Params
	err := cdc.UnmarshalJSON([]byte(res), &params)
	require.NoError(f.T, err)
	return params
}

//___________________________________________________________________________________
// executors

func executeWriteCheckErr(t *testing.T, cmdStr string, writes ...string) {
	require.True(t, executeWrite(t, cmdStr, writes...))
}

func executeWrite(t *testing.T, cmdStr string, writes ...string) (exitSuccess bool) {
	exitSuccess, _, _ = executeWriteRetStdStreams(t, cmdStr, writes...)
	return
}

func executeWriteRetStdStreams(t *testing.T, cmdStr string, writes ...string) (bool, string, string) {
	proc := tests.GoExecuteT(t, cmdStr)

	// Enables use of interactive commands
	for _, write := range writes {
		_, err := proc.StdinPipe.Write([]byte(write + "\n"))
		require.NoError(t, err)
	}

	// Read both stdout and stderr from the process
	stdout, stderr, err := proc.ReadAll()
	if err != nil {
		fmt.Println("Err on proc.ReadAll()", err, cmdStr)
	}

	// Log output.
	if len(stdout) > 0 {
		t.Log(time.Now().Format("2006-01-02 15:04:05"), "Stdout:", cmn.Green(string(stdout)))
	}
	if len(stderr) > 0 {
		t.Log(time.Now().Format("2006-01-02 15:04:05"), "Stderr:", cmn.Red(string(stderr)))
	}

	// Wait for process to exit
	proc.Wait()

	// Return succes, stdout, stderr
	return proc.ExitState.Success(), string(stdout), string(stderr)
}

//___________________________________________________________________________________
// utils

func addFlags(cmd string, flags []string) string {
	for _, f := range flags {
		cmd += " " + f
	}
	return strings.TrimSpace(cmd)
}

func queryTags(tags []string) (out string) {
	for _, tag := range tags {
		out += tag + "&"
	}
	return strings.TrimSuffix(out, "&")
}

// Write the given string to a new temporary file
func WriteToNewTempFile(t *testing.T, s string) *os.File {
	fp, err := ioutil.TempFile(os.TempDir(), "hashgard_cli_test_")
	require.Nil(t, err)
	_, err = fp.WriteString(s)
	require.Nil(t, err)
	return fp
}

func marshalStdTx(t *testing.T, stdTx auth.StdTx) []byte {
	cdc := app.MakeCodec()
	bz, err := cdc.MarshalBinaryBare(stdTx)
	require.NoError(t, err)
	return bz
}

func unmarshalStdTx(t *testing.T, s string) (stdTx auth.StdTx) {
	cdc := app.MakeCodec()
	require.Nil(t, cdc.UnmarshalJSON([]byte(s), &stdTx))
	return
}

//___________________________________________________________________________________
// hashgardcli exchange

// TxExchangeCreateOrder is hashgardcli exchange create-order
func (f *Fixtures) TxExchangeCreateOrder(from string, supply, target sdk.Coin, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli exchange create-order %v --from=%s --supply=%s --target=%s", f.Flags(), from, supply, target)
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxExchangeWithdrawalOrder is hashgardcli exchange withdrawal-order
func (f *Fixtures) TxExchangeWithdrawalOrder(orderId int, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli exchange withdrawal-order %d --from=%s %v", orderId, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxExchangeTakeOrder is hashgardcli exchange take-order
func (f *Fixtures) TxExchangeTakeOrder(orderId int, amount sdk.Coin, from string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli exchange take-order %d --amount=%s --from=%s %v", orderId, amount, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// QueryExchangeOrder is hashgardcli exchange query-order
func (f *Fixtures) QueryExchangeOrder(orderId int, flags ...string) exchange.Order {
	cmd := fmt.Sprintf("../build/hashgardcli exchange query-order %d %v", orderId, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var order exchange.Order
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &order)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return order
}

// QueryExchangeOrders is hashgardcli exchange query-orders
func (f *Fixtures) QueryExchangeOrders(addr sdk.AccAddress, flags ...string) []exchange.Order {
	cmd := fmt.Sprintf("../build/hashgardcli exchange query-order %s %v", addr, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var orders []exchange.Order
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &orders)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return orders
}

// QueryExchangeFrozen is hashgardcli exchange query-frozen
func (f *Fixtures) QueryExchangeFrozen(addr sdk.AccAddress, flags ...string) sdk.Coins {
	cmd := fmt.Sprintf("../build/hashgardcli exchange query-frozen %s %v", addr, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var coins sdk.Coins
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &coins)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return coins
}

//___________________________________________________________________________________
// hashgardcli issue

// TxIssueCreate is hashgardcli issue create
func (f *Fixtures) TxIssueCreate(from string, name string, symbol string, totalSupply uint64, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli issue create %s %s %d --from=%s %v", name, symbol, totalSupply, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxIssueTransferOwnership is hashgardcli issue transfer-ownership
func (f *Fixtures) TxIssueTransferOwnership(from string, issueId string, addr sdk.AccAddress, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli issue transfer-ownership %s %s --from=%s %v", issueId, addr, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxIssueDescribe is hashgardcli issue describe
func (f *Fixtures) TxIssueDescribe(from string, issueId string, filePath string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli issue describe %s %s --from=%s %v", issueId, filePath, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxIssueMint is hashgardcli issue Mint
func (f *Fixtures) TxIssueMint(from string, issueId string, amount uint64, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli issue mint %s %d --from=%s %v", issueId, amount, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxIssueDisable is hashgardcli issue disable
func (f *Fixtures) TxIssueDisable(from string, issueId string, feature string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli issue disable %s %s --from=%s %v", issueId, feature, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxIssueFreeze is hashgardcli issue freeze
func (f *Fixtures) TxIssueFreeze(from string, freezeType string, issueId string, addr string, endTime string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli issue freeze %s %s %s %s --from=%s %v", freezeType, issueId, addr, endTime, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxIssueUnFreeze is hashgardcli issue unfreeze
func (f *Fixtures) TxIssueUnFreeze(from string, freezeType string, issueId string, addr string, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli issue unfreeze %s %s %s --from=%s %v", freezeType, issueId, addr, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxIssueBurn is hashgardcli issue burn
func (f *Fixtures) TxIssueBurn(from string, issueId string, amount uint64, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli issue burn %s %d --from=%s %v", issueId, amount, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxIssueBurnFrom is hashgardcli issue burn-from
func (f *Fixtures) TxIssueBurnFrom(from string, issueId string, account string, amount uint64, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli issue burn-from %s %s %d --from=%s %v", issueId, account, amount, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxIssueSendFrom is hashgardcli issue send-from
func (f *Fixtures) TxIssueSendFrom(from string, issueId string, fromAddr string, toAddr string, amount uint64, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli issue send-from %s %s %s %d --from=%s %v", issueId, fromAddr, toAddr, amount, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxIssueApprove is hashgardcli issue approve
func (f *Fixtures) TxIssueApprove(from string, issueId string, addr string, amount uint64, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli issue approve %s %s %d --from=%s %v", issueId, addr, amount, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxIssueIncreaseApprove is hashgardcli issue increase-approval
func (f *Fixtures) TxIssueIncreaseApprove(from string, issueId string, addr string, amount uint64, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli issue increase-approval %s %s %d --from=%s %v", issueId, addr, amount, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// TxIssueDecreaseApprove is hashgardcli issue decrease-approval
func (f *Fixtures) TxIssueDecreaseApprove(from string, issueId string, addr string, amount uint64, flags ...string) (bool, string, string) {
	cmd := fmt.Sprintf("../build/hashgardcli issue decrease-approval %s %s %d --from=%s %v", issueId, addr, amount, from, f.Flags())
	return executeWriteRetStdStreams(f.T, addFlags(cmd, flags), app.DefaultKeyPass)
}

// QueryIssueIssue is hashgardcli issue query-issue
func (f *Fixtures) QueryIssueIssue(issueId string, flags ...string) issue.CoinIssueInfo {
	cmd := fmt.Sprintf("../build/hashgardcli issue query-issue %s %v", issueId, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var coinsIssueInfo issue.CoinIssueInfo
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &coinsIssueInfo)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return coinsIssueInfo
}

// QueryIssueAllowance is hashgardcli issue query-allowance
func (f *Fixtures) QueryIssueAllowance(issueId string, owner string, spender string, flags ...string) issue.Approval {
	cmd := fmt.Sprintf("../build/hashgardcli issue query-allowance %s %s %s %v", issueId, owner, spender, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var approval issue.Approval
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &approval)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return approval
}

// QueryIssueFreeze is hashgardcli issue query-freeze
func (f *Fixtures) QueryIssueFreeze(issueId string, addr string, flags ...string) issue.IssueFreeze {
	cmd := fmt.Sprintf("../build/hashgardcli issue query-freeze %s %s %v", issueId, addr, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var freeze issue.IssueFreeze
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &freeze)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return freeze
}

// QueryIssueIssues is hashgardcli issue list-issues
func (f *Fixtures) QueryIssueIssues(owner string, flags ...string) []issue.CoinIssueInfo {
	cmd := fmt.Sprintf("../build/hashgardcli issue list-issues --address=%s %v", owner, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var issueList []issue.CoinIssueInfo
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &issueList)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return issueList
}

// QueryIssueSearch is hashgardcli issue search
func (f *Fixtures) QueryIssueSearch(symbol string, flags ...string) []issue.CoinIssueInfo {
	cmd := fmt.Sprintf("../build/hashgardcli issue search %s %v", symbol, f.Flags())
	out, _ := tests.ExecuteT(f.T, addFlags(cmd, flags), "")
	var issueList []issue.CoinIssueInfo
	cdc := app.MakeCodec()
	err := cdc.UnmarshalJSON([]byte(out), &issueList)
	require.NoError(f.T, err, "out %v\n, err %v", out, err)
	return issueList
}
