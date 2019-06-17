package cli

import (
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientutils "github.com/hashgard/hashgard/x/box/client/utils"
)

func ProcessBoxInject(cdc *codec.Codec, id string, amountStr string, operation string) error {
	txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
	if err != nil {
		return err
	}
	msg, err := clientutils.GetInjectMsg(cdc, cliCtx, account, id, amountStr, operation, true)
	if err != nil {
		return err
	}
	return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
}
