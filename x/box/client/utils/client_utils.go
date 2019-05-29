package utils

import (
	"fmt"
	"reflect"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	boxqueriers "github.com/hashgard/hashgard/x/box/client/queriers"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/types"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
	issueclientutils "github.com/hashgard/hashgard/x/issue/client/utils"
)

func GetCliContext(cdc *codec.Codec) (authtxb.TxBuilder, context.CLIContext, auth.Account, error) {
	txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))
	cliCtx := context.NewCLIContext().
		WithCodec(cdc).
		WithAccountDecoder(cdc)
	from := cliCtx.GetFromAddress()
	account, err := cliCtx.GetAccount(from)

	return txBldr, cliCtx, account, err
}

func GetBoxByID(cdc *codec.Codec, cliCtx context.CLIContext, id string) (types.Box, error) {
	var boxInfo types.Box
	// Query the box
	res, err := boxqueriers.QueryBoxByID(id, cliCtx)
	if err != nil {
		return nil, err
	}
	cdc.MustUnmarshalJSON(res, &boxInfo)
	return boxInfo, nil
}

func BoxOwnerCheck(cdc *codec.Codec, cliCtx context.CLIContext, sender auth.Account, id string) (types.Box, error) {
	boxInfo, err := GetBoxByID(cdc, cliCtx, id)
	if err != nil {
		return nil, err
	}
	if !sender.GetAddress().Equals(boxInfo.GetOwner()) {
		return nil, errors.Errorf(errors.ErrOwnerMismatch(id))
	}
	return boxInfo, nil
}

func GetCoinDecimal(cdc *codec.Codec, cliCtx context.CLIContext, coin sdk.Coin) (uint, error) {
	if coin.Denom == types.Agard {
		return types.AgardDecimals, nil
	}
	issueInfo, err := issueclientutils.GetIssueByID(cdc, cliCtx, coin.Denom)
	if err != nil {
		return 0, err
	}
	return issueInfo.GetDecimals(), nil
}
func GetBoxInfo(box types.BoxInfo) fmt.Stringer {
	switch box.BoxType {
	case types.Lock:
		var clientBox LockBoxInfo
		StructCopy(&clientBox, &box)
		return clientBox
	case types.Deposit:
		return processDepositBoxInfo(box)
	case types.Future:
		var clientBox FutureBoxInfo
		StructCopy(&clientBox, &box)
		return clientBox
	default:
		return box
	}
}
func processDepositBoxInfo(box types.BoxInfo) DepositBoxInfo {
	var clientBox DepositBoxInfo
	StructCopy(&clientBox, &box)
	return clientBox
}
func GetBoxList(boxs types.BoxInfos, boxType string) fmt.Stringer {
	switch boxType {
	case types.Lock:
		var boxInfos = make(LockBoxInfos, 0, len(boxs))
		for _, box := range boxs {
			var clientBox LockBoxInfo
			StructCopy(&clientBox, &box)
			boxInfos = append(boxInfos, clientBox)
		}

		return boxInfos
	case types.Deposit:
		var boxInfos = make(DepositBoxInfos, 0, len(boxs))
		for _, box := range boxs {
			boxInfos = append(boxInfos, processDepositBoxInfo(box))
		}

		return boxInfos
	case types.Future:
		var boxInfos = make(FutureBoxInfos, 0, len(boxs))
		for _, box := range boxs {
			var clientBox FutureBoxInfo
			StructCopy(&clientBox, &box)
			boxInfos = append(boxInfos, clientBox)
		}

		return boxInfos
	}
	return boxs
}
func deepFields(faceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField
	for i := 0; i < faceType.NumField(); i++ {
		v := faceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, deepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}
	return fields
}

func StructCopy(destPtr interface{}, srcPtr interface{}) {
	srcv := reflect.ValueOf(srcPtr)
	dstv := reflect.ValueOf(destPtr)
	srct := reflect.TypeOf(srcPtr)
	dstt := reflect.TypeOf(destPtr)
	if srct.Kind() != reflect.Ptr || dstt.Kind() != reflect.Ptr ||
		srct.Elem().Kind() == reflect.Ptr || dstt.Elem().Kind() == reflect.Ptr {
		panic("Fatal error:type of parameters must be Ptr of value")
	}
	if srcv.IsNil() || dstv.IsNil() {
		panic("Fatal error:value of parameters should not be nil")
	}
	srcV := srcv.Elem()
	dstV := dstv.Elem()
	srcfields := deepFields(reflect.ValueOf(srcPtr).Elem().Type())
	for _, v := range srcfields {
		if v.Anonymous {
			continue
		}
		dst := dstV.FieldByName(v.Name)
		src := srcV.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return
}
func GetDepositMsg(cdc *codec.Codec, cliCtx context.CLIContext, account auth.Account,
	id string, amountStr string, operation string, cli bool) (sdk.Msg, error) {

	if err := boxutils.CheckId(id); err != nil {
		return nil, errors.Errorf(err)
	}
	amount, ok := sdk.NewIntFromString(amountStr)
	if !ok {
		return nil, errors.Errorf(errors.ErrAmountNotValid(amountStr))
	}

	boxInfo, err := GetBoxByID(cdc, cliCtx, id)
	if err != nil {
		return nil, err
	}
	if types.BoxDepositing != boxInfo.GetStatus() {
		return nil, errors.Errorf(errors.ErrNotAllowedOperation(boxInfo.GetStatus()))
	}
	if cli {
		decimal, err := GetCoinDecimal(cdc, cliCtx, boxInfo.GetTotalAmount().Token)
		if err != nil {
			return nil, err
		}
		amount = boxutils.MulDecimals(boxutils.ParseCoin(boxInfo.GetTotalAmount().Token.Denom, amount), decimal)
	}
	var msg sdk.Msg
	switch operation {
	case types.DepositTo:
		if err = checkAmountByDepositTo(amount, boxInfo); err != nil {
			return nil, err
		}
		msg = msgs.NewMsgBoxDepositTo(id, account.GetAddress(),
			sdk.NewCoin(boxInfo.GetTotalAmount().Token.Denom, amount))
	case types.Fetch:
		if err = checkAmountByFetch(amount, boxInfo, account); err != nil {
			return nil, err
		}
		msg = msgs.NewMsgBoxDepositFetch(id, account.GetAddress(),
			sdk.NewCoin(boxInfo.GetTotalAmount().Token.Denom, amount))
	default:
		return nil, errors.ErrNotSupportOperation()
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, errors.Errorf(err)
	}
	return msg, nil
}
func GetInterestMsg(cdc *codec.Codec, cliCtx context.CLIContext, account auth.Account,
	id string, amountStr string, operation string, cli bool) (sdk.Msg, error) {

	if err := boxutils.CheckId(id); err != nil {
		return nil, errors.Errorf(err)
	}
	amount, ok := sdk.NewIntFromString(amountStr)
	if !ok {
		return nil, errors.Errorf(errors.ErrAmountNotValid(amountStr))
	}
	box, err := GetBoxByID(cdc, cliCtx, id)
	if err != nil {
		return nil, err
	}
	if box.GetBoxType() != types.Deposit {
		return nil, errors.Errorf(errors.ErrNotSupportOperation())
	}
	if box.GetStatus() != types.BoxCreated {
		return nil, errors.Errorf(errors.ErrNotSupportOperation())
	}
	if cli {
		decimal, err := GetCoinDecimal(cdc, cliCtx, box.GetDeposit().Interest.Token)
		if err != nil {
			return nil, err
		}
		amount = boxutils.MulDecimals(boxutils.ParseCoin(box.GetDeposit().Interest.Token.Denom, amount), decimal)
	}
	var msg sdk.Msg
	switch operation {
	case types.Fetch:
		flag := true
		for i, v := range box.GetDeposit().InterestInjections {
			if v.Address.Equals(account.GetAddress()) {
				if box.GetDeposit().InterestInjections[i].Amount.GTE(amount) {
					flag = false
					break
				}
			}
		}
		if flag {
			return nil, errors.ErrNotEnoughAmount()
		}
		msg = msgs.NewMsgBoxInterestFetch(id, account.GetAddress(), sdk.NewCoin(box.GetDeposit().Interest.Token.Denom, amount))
	case types.Injection:
		if box.GetDeposit().InterestInjections != nil {
			totalInterest := sdk.ZeroInt()
			for _, v := range box.GetDeposit().InterestInjections {
				if v.Address.Equals(account.GetAddress()) {
					totalInterest = totalInterest.Add(v.Amount)
				}
			}
			if totalInterest.Add(amount).GT(box.GetDeposit().Interest.Token.Amount) {
				return nil, errors.Errorf(errors.ErrInterestInjectionNotValid(
					sdk.NewCoin(box.GetDeposit().Interest.Token.Denom, amount)))
			}
		}
		msg = msgs.NewMsgBoxInterestInjection(id, account.GetAddress(), sdk.NewCoin(box.GetDeposit().Interest.Token.Denom, amount))
	default:
		return nil, errors.ErrUnknownOperation()
	}
	if err := msg.ValidateBasic(); err != nil {
		return nil, errors.Errorf(err)
	}
	return msg, nil
}
func GetWithdrawMsg(cdc *codec.Codec, cliCtx context.CLIContext, account auth.Account, id string) (sdk.Msg, error) {
	if account.GetCoins().AmountOf(id).IsZero() {
		return nil, errors.Errorf(errors.ErrNotEnoughAmount())
	}
	boxInfo, err := GetBoxByID(cdc, cliCtx, id)
	if err != nil {
		return nil, err
	}
	switch boxInfo.GetBoxType() {
	case types.Deposit:
		if types.BoxFinished != boxInfo.GetStatus() {
			return nil, errors.Errorf(errors.ErrNotAllowedOperation(boxInfo.GetStatus()))
		}
	case types.Future:
		if types.BoxCreated == boxInfo.GetStatus() {
			return nil, errors.Errorf(errors.ErrNotAllowedOperation(boxInfo.GetStatus()))
		}
		seq := boxutils.GetSeqFromFutureBoxSeq(id)
		if boxInfo.GetFuture().TimeLine[seq-1] > time.Now().Unix() {
			return nil, errors.Errorf(errors.ErrNotAllowedOperation(types.BoxUndue))
		}
	default:
		return nil, errors.Errorf(errors.ErrUnknownBox(id))
	}
	msg := msgs.NewMsgBoxWithdraw(id, account.GetAddress())
	if err := msg.ValidateBasic(); err != nil {
		return nil, errors.Errorf(err)
	}
	return msg, nil
}
func checkAmountByFetch(amount sdk.Int, boxInfo types.Box, account auth.Account) error {
	switch boxInfo.GetBoxType() {
	case types.Deposit:
		if !amount.Mod(boxInfo.GetDeposit().Price).IsZero() {
			return errors.ErrAmountNotValid(amount.String())
		}
		if account.GetCoins().AmountOf(boxInfo.GetId()).LT(amount.Quo(boxInfo.GetDeposit().Price)) {
			return errors.Errorf(errors.ErrNotEnoughAmount())
		}
	case types.Future:
		if boxInfo.GetFuture().Deposits == nil {
			return errors.Errorf(errors.ErrNotEnoughAmount())
		}
		for _, v := range boxInfo.GetFuture().Deposits {
			if v.Address.Equals(account.GetAddress()) {
				if v.Amount.GTE(amount) {
					return nil
				}
			}
		}
		return errors.Errorf(errors.ErrNotEnoughAmount())
	default:
		return errors.Errorf(errors.ErrNotSupportOperation())
	}
	return nil
}
func checkAmountByDepositTo(amount sdk.Int, boxInfo types.Box) error {
	switch boxInfo.GetBoxType() {
	case types.Deposit:
		if !amount.Mod(boxInfo.GetDeposit().Price).IsZero() {
			return errors.ErrAmountNotValid(amount.String())
		}
		if amount.Add(boxInfo.GetDeposit().TotalDeposit).GT(boxInfo.GetTotalAmount().Token.Amount) {
			return errors.Errorf(errors.ErrNotEnoughAmount())
		}
	case types.Future:
		total := sdk.ZeroInt()
		if boxInfo.GetFuture().Deposits != nil {
			for _, v := range boxInfo.GetFuture().Deposits {
				total = total.Add(v.Amount)
			}
		}
		if amount.Add(total).GT(boxInfo.GetTotalAmount().Token.Amount) {
			return errors.Errorf(errors.ErrNotEnoughAmount())
		}
	default:
		return errors.Errorf(errors.ErrNotSupportOperation())
	}
	return nil
}
