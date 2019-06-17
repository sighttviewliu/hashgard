package utils

import (
	"fmt"
	"strconv"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/types"

	issueutils "github.com/hashgard/hashgard/x/issue/utils"
)

func MulDecimals(coin sdk.Coin, decimals uint) sdk.Int {
	if coin.Denom == types.Agard {
		return coin.Amount
	}
	return issueutils.MulDecimals(coin.Amount, decimals)
}
func ParseCoin(denom string, amount sdk.Int) sdk.Coin {
	if denom == types.Agard {
		denom = types.Gard
	}
	coin, _ := sdk.ParseCoin(fmt.Sprintf("%s%s", amount, denom))
	return coin
}
func CalcInterest(perCoupon sdk.Dec, share sdk.Int, interest types.BoxToken) sdk.Int {
	dec := perCoupon.MulInt(share)
	decimals := interest.Decimals
	if interest.Token.Denom == types.Agard {
		decimals = types.GardDecimals
	}
	dec = GetMaxPrecision(dec, decimals)
	return dec.MulInt(issueutils.GetDecimalsInt(decimals)).TruncateInt()
}

func IsId(id string) bool {
	return strings.HasPrefix(id, types.IDPreStr)
}

func CheckId(id string) sdk.Error {
	if !IsId(id) {
		return errors.ErrBoxID(id)
	}
	return nil
}

func CalcInterestRate(totalAmount sdk.Int, price sdk.Int, interest sdk.Coin, decimals uint) sdk.Dec {
	totalCoupon := totalAmount.Quo(price)
	perCoupon := sdk.NewDecFromBigInt(interest.Amount.BigInt()).QuoInt(totalCoupon)
	if interest.Denom == types.Agard {
		decimals = types.GardDecimals
	}
	return quoMaxPrecisionByDecimal(perCoupon, decimals)
}

func quoMaxPrecisionByDecimal(dec sdk.Dec, decimals uint) sdk.Dec {
	dec = dec.QuoInt(issueutils.GetDecimalsInt(decimals))
	dec = GetMaxPrecision(dec, decimals)
	return dec
}

func GetBoxTypeByValue(value string) string {
	value = strings.ReplaceAll(value, types.IDPreStr, "")
	for k, v := range types.BoxType {
		if strings.HasPrefix(value, v) {
			return k
		}
	}
	return ""
}
func GetCoinDenomByFutureBoxSeq(id string, seq int) string {
	return fmt.Sprintf("%s%02d", id, seq)
}
func GetIdFromBoxSeqID(idSeq string) string {
	if len(idSeq) > types.IdLength {
		return idSeq[:types.IdLength]
	}
	return idSeq
}
func GetSeqFromFutureBoxSeq(boxSeqStr string) int {
	seqStr := boxSeqStr[len(boxSeqStr)-2:]
	seq, _ := strconv.Atoi(seqStr)
	return seq
}
func GetMaxPrecision(dec sdk.Dec, decimals uint) sdk.Dec {
	precision := types.MaxPrecision
	if decimals < types.MaxPrecision {
		precision = decimals
	}
	decStr := dec.String()
	len := strings.Index(decStr, ".") + int(precision)
	str := decStr[0 : len+1]
	dec, _ = sdk.NewDecFromStr(str)
	return dec
}
