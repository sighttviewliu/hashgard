// +build cli_test

package clitest

import (
	"encoding/hex"
	"fmt"
	"testing"
	"time"

	"github.com/hashgard/hashgard/x/box/types"

	"github.com/hashgard/hashgard/x/box/utils"

	"github.com/hashgard/hashgard/x/box/params"

	"github.com/cosmos/cosmos-sdk/tests"
	sdk "github.com/cosmos/cosmos-sdk/types"
	boxtests "github.com/hashgard/hashgard/x/box/tests"
	issueutils "github.com/hashgard/hashgard/x/issue/utils"
	"github.com/stretchr/testify/require"
)

var (
	decimals        = uint(18)
	IssueCoinAmount = sdk.NewInt(10000)
	BoxAmount       = sdk.NewInt(1000)
	Interest        = sdk.NewInt(200)
	DefaultFlag     = "--gas 200000 -y"

	futureBoxReceivers = []string{keyBar, keyBaz, keyVesting}
)

func AddIssue(t *testing.T, f *Fixtures, sender sdk.AccAddress) string {
	// create issue
	f.TxIssueCreate(keyIssue, "foocoin", "FOO", uint64(IssueCoinAmount.Int64()), fmt.Sprintf("--decimals %d --gas 200000 -y", decimals))
	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	txs := f.QueryTxs(1, 50, "action:issue", fmt.Sprintf("sender:%s", sender))
	require.NotNil(t, txs)
	bytes, _ := hex.DecodeString(txs[len(txs)-1].Data)
	issueId := string(bytes[2:])
	return issueId
}
func CreateLockBox(t *testing.T, f *Fixtures, issueID string, sender sdk.AccAddress) (string, *params.BoxLockParams) {
	params := boxtests.GetLockBoxInfo()

	params.TotalAmount.Token.Denom = issueID
	params.TotalAmount.Token.Amount = BoxAmount
	f.TxLockBoxCreate(sender.String(), params, DefaultFlag)

	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	txs1 := f.QueryTxs(1, 50, "action:"+types.TypeMsgBoxCreate, fmt.Sprintf("sender:%s", sender))
	require.Len(t, txs1, 1)
	bytes, _ := hex.DecodeString(txs1[0].Data)
	boxId := string(bytes[2:])
	return boxId, params
}
func CreateDepositBox(t *testing.T, f *Fixtures, issueAID string, issueBID string, sender sdk.AccAddress) (string, *params.BoxDepositParams) {
	params := boxtests.GetDepositBoxInfo()

	params.TotalAmount.Token.Amount = BoxAmount
	params.TotalAmount.Token.Denom = issueAID
	params.Deposit.Interest.Token.Denom = issueBID
	params.Deposit.Interest.Token.Amount = Interest

	params.Deposit.BottomLine = BoxAmount.QuoRaw(4)
	params.Deposit.Price = BoxAmount.QuoRaw(10)

	f.TxDepositBoxCreate(sender.String(), params, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	txs1 := f.QueryTxs(1, 50, "action:"+types.TypeMsgBoxCreate, fmt.Sprintf("sender:%s", sender))
	require.Len(t, txs1, 1)
	bytes, _ := hex.DecodeString(txs1[0].Data)
	boxId := string(bytes[2:])
	return boxId, params
}
func CreateFutureBox(t *testing.T, f *Fixtures, issueID string, sender sdk.AccAddress) (string, *params.BoxFutureParams) {
	params := boxtests.GetFutureBoxInfo()

	params.TotalAmount.Token.Amount = issueutils.QuoDecimals(params.TotalAmount.Token.Amount, decimals)
	params.TotalAmount.Token.Denom = issueID

	for i, items := range params.Future.Receivers {
		for j, v := range items {
			if j == 0 {
				params.Future.Receivers[i][j] = f.KeyAddress(futureBoxReceivers[i]).String()
				continue
			}
			amount, _ := sdk.NewIntFromString(v)
			params.Future.Receivers[i][j] = issueutils.QuoDecimals(amount, decimals).String()
		}
	}

	f.TxFutureBoxCreate(sender.String(), params, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	// Ensure transaction tags can be queried
	txs1 := f.QueryTxs(1, 50, "action:"+types.TypeMsgBoxCreate, fmt.Sprintf("sender:%s", sender))
	require.Len(t, txs1, 1)
	bytes, _ := hex.DecodeString(txs1[0].Data)
	boxId := string(bytes[2:])
	return boxId, params
}
func getWaitBlocks(endTime int64) int64 {
	duration := endTime - time.Now().Unix()
	if duration < 0 {
		duration = 1
	}
	return duration + 2
}
func TestHashgardCLILockBox(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)
	// start hashgard server
	proc := f.HGStart()
	defer proc.Stop(false)
	// Save key addresses for later use
	issueAddr := f.KeyAddress(keyIssue)
	//barAddr := f.KeyAddress(keyBar)
	issueID := AddIssue(t, f, issueAddr)
	boxID, params := CreateLockBox(t, f, issueID, issueAddr)

	issueAcc := f.QueryAccount(issueAddr)
	//require.Equal(t, issueAcc.GetCoins().AmountOf(issueID), IssueCoinAmount.Sub(BoxAmount))
	require.Equal(t, issueAcc.GetCoins().AmountOf(boxID), issueutils.MulDecimals(params.TotalAmount.Token.Amount, decimals))

	tests.WaitForNextNBlocksTM(getWaitBlocks(params.Lock.EndTime), f.Port)

	issueAcc = f.QueryAccount(issueAddr)
	//require.Equal(t, issueAcc.GetCoins().AmountOf(issueID), IssueCoinAmount)
	require.Equal(t, issueAcc.GetCoins().AmountOf(boxID), sdk.ZeroInt())
}

func TestHashgardCLIDepositBox(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)
	// start hashgard server
	proc := f.HGStart()
	defer proc.Stop(false)
	// Save key addresses for later use
	issueAddr := f.KeyAddress(keyIssue)
	barAddr := f.KeyAddress(keyBar)
	//barAddr := f.KeyAddress(keyBar)
	issueAID := AddIssue(t, f, issueAddr)
	issueBID := AddIssue(t, f, issueAddr)

	boxID, params := CreateDepositBox(t, f, issueAID, issueBID, issueAddr)

	f.TxDepositBoxInterestInject(keyIssue, boxID, params.Deposit.Interest.Token.Amount, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	txsInject := f.QueryTxs(1, 50, "action:"+types.TypeMsgBoxInterestInject, fmt.Sprintf("sender:%s", issueAddr))
	require.Len(t, txsInject, 1)

	f.TxDepositBoxInterestCancel(keyIssue, boxID, params.Deposit.Interest.Token.Amount, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	txsInject = f.QueryTxs(1, 50, "action:"+types.TypeMsgBoxInterestCancel, fmt.Sprintf("sender:%s", issueAddr))
	require.Len(t, txsInject, 1)

	f.TxDepositBoxInterestInject(keyIssue, boxID, params.Deposit.Interest.Token.Amount, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	txsInject = f.QueryTxs(1, 50, "action:"+types.TypeMsgBoxInterestInject, fmt.Sprintf("sender:%s", issueAddr))
	require.Len(t, txsInject, 2)

	depositTo := BoxAmount.QuoRaw(int64(2))
	f.TxSend(keyIssue, barAddr, sdk.NewCoin(issueAID, depositTo), DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)

	tests.WaitForNextNBlocksTM(getWaitBlocks(params.Deposit.StartTime), f.Port)

	f.TxInject(keyBar, boxID, depositTo, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	txsDeposit := f.QueryTxs(1, 50, "action:"+types.TypeMsgBoxInject, fmt.Sprintf("sender:%s", barAddr.String()))
	require.Len(t, txsDeposit, 1)

	f.TxDepositCancel(keyBar, boxID, depositTo, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	txsDeposit = f.QueryTxs(1, 50, "action:"+types.TypeMsgBoxCancel, fmt.Sprintf("sender:%s", barAddr.String()))
	require.Len(t, txsDeposit, 1)

	f.TxInject(keyBar, boxID, depositTo, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	txsDeposit = f.QueryTxs(1, 50, "action:"+types.TypeMsgBoxInject, fmt.Sprintf("sender:%s", barAddr.String()))
	require.Len(t, txsDeposit, 2)

	issueAcc := f.QueryAccount(barAddr)
	require.Equal(t, issueAcc.GetCoins().AmountOf(boxID), depositTo.Quo(params.Deposit.Price))

	tests.WaitForNextNBlocksTM(getWaitBlocks(params.Deposit.MaturityTime), f.Port)

	f.TxWithdraw(keyBar, boxID, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	txsWithdraw := f.QueryTxs(1, 50, "action:"+types.TypeMsgBoxWithdraw, fmt.Sprintf("sender:%s", barAddr))
	require.Len(t, txsWithdraw, 1)

	issueAcc = f.QueryAccount(barAddr)
	require.Equal(t, issueAcc.GetCoins().AmountOf(boxID), sdk.ZeroInt())
}

func TestHashgardCLIFutureBox(t *testing.T) {
	t.Parallel()
	f := InitFixtures(t)
	// start hashgard server
	proc := f.HGStart()
	defer proc.Stop(false)
	// Save key addresses for later use
	issueAddr := f.KeyAddress(keyIssue)
	barAddr := f.KeyAddress(keyBar)
	issueID := AddIssue(t, f, issueAddr)
	boxID, params := CreateFutureBox(t, f, issueID, issueAddr)

	depositTo := params.TotalAmount.Token.Amount

	f.TxSend(keyIssue, barAddr, sdk.NewCoin(issueID, depositTo.QuoRaw(2)), DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.TxInject(keyBar, boxID, depositTo.QuoRaw(2), DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	txsDeposit := f.QueryTxs(1, 50, "action:"+types.TypeMsgBoxInject, fmt.Sprintf("sender:%s", barAddr.String()))
	require.Len(t, txsDeposit, 1)

	f.TxDepositCancel(keyBar, boxID, depositTo.QuoRaw(2), DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	txsDeposit = f.QueryTxs(1, 50, "action:"+types.TypeMsgBoxCancel, fmt.Sprintf("sender:%s", barAddr.String()))
	require.Len(t, txsDeposit, 1)

	f.TxSend(keyBar, issueAddr, sdk.NewCoin(issueID, depositTo.QuoRaw(2)), DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)

	f.TxInject(keyIssue, boxID, depositTo, DefaultFlag)
	tests.WaitForNextNBlocksTM(1, f.Port)
	txsDeposit = f.QueryTxs(1, 50, "action:"+types.TypeMsgBoxInject, fmt.Sprintf("sender:%s", issueAddr.String()))
	require.Len(t, txsDeposit, 1)

	tests.WaitForNextNBlocksTM(getWaitBlocks(params.Future.TimeLine[len(params.Future.TimeLine)-1]), f.Port)

	var address sdk.AccAddress
	for i, v := range params.Future.Receivers {
		address, _ = sdk.AccAddressFromBech32(v[0])
		account := f.QueryAccount(address)
		totalAmount := sdk.ZeroInt()
		for _, coin := range account.GetCoins() {
			if utils.IsId(coin.Denom) {
				f.TxWithdraw(futureBoxReceivers[i], coin.Denom, DefaultFlag)
				tests.WaitForNextNBlocksTM(1, f.Port)
				totalAmount = totalAmount.Add(coin.Amount)
			}
		}
		account = f.QueryAccount(address)
		require.Equal(t, account.GetCoins().AmountOf(params.TotalAmount.Token.Denom), totalAmount)
	}
}
