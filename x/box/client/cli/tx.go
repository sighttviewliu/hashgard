package cli

import (
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	clientutils "github.com/hashgard/hashgard/x/box/client/utils"
	"github.com/hashgard/hashgard/x/box/errors"
	"github.com/hashgard/hashgard/x/box/msgs"
	"github.com/hashgard/hashgard/x/box/types"
	boxutils "github.com/hashgard/hashgard/x/box/utils"
)

func ProcessBoxDescriptionCmd(cdc *codec.Codec, boxType string, id string, filename string) error {
	if boxutils.GetBoxTypeByValue(id) != boxType {
		return errors.Errorf(errors.ErrUnknownBox(id))
	}
	if err := boxutils.CheckId(id); err != nil {
		return errors.Errorf(err)
	}
	txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
	if err != nil {
		return err
	}
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	buffer := bytes.Buffer{}
	err = json.Compact(&buffer, contents)
	if err != nil {
		return errors.ErrBoxDescriptionNotValid()
	}
	contents = buffer.Bytes()

	_, err = clientutils.BoxOwnerCheck(cdc, cliCtx, account, id)
	if err != nil {
		return err
	}
	if len(contents) <= 0 || !json.Valid(contents) {
		return errors.ErrBoxDescriptionNotValid()
	}
	msg := msgs.NewMsgBoxDescription(id, account.GetAddress(), contents)

	validateErr := msg.ValidateBasic()

	if validateErr != nil {
		return errors.Errorf(validateErr)
	}
	return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
}

func ProcessBoxDisableFeatureCmd(cdc *codec.Codec, boxType string, id string, feature string) error {
	if boxutils.GetBoxTypeByValue(id) != boxType {
		return errors.Errorf(errors.ErrUnknownBox(id))
	}
	_, ok := types.Features[feature]
	if !ok {
		return errors.Errorf(errors.ErrUnknownFeatures())
	}
	if err := boxutils.CheckId(id); err != nil {
		return errors.Errorf(err)
	}
	txBldr, cliCtx, account, err := clientutils.GetCliContext(cdc)
	if err != nil {
		return err
	}
	boxInfo, err := clientutils.BoxOwnerCheck(cdc, cliCtx, account, id)
	if err != nil {
		return err
	}

	if feature == types.Transfer && boxInfo.GetBoxType() == types.Lock {
		return errors.Errorf(errors.ErrNotSupportOperation())
	}

	msg := msgs.NewMsgBoxDisableFeature(id, account.GetAddress(), feature)
	validateErr := msg.ValidateBasic()
	if validateErr != nil {
		return errors.Errorf(validateErr)
	}
	return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
}
