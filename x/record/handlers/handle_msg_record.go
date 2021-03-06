package handlers

import (
	"github.com/hashgard/hashgard/x/record/utils"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/record/keeper"
	"github.com/hashgard/hashgard/x/record/msgs"
	"github.com/hashgard/hashgard/x/record/types"
)

//Handle MsgRecord
func HandleMsgRecord(ctx sdk.Context, keeper keeper.Keeper, msg msgs.MsgRecord) sdk.Result {
	recordInfo := types.RecordInfo{
		Sender:      msg.Sender,
		Hash:        msg.Hash,
		Name:        msg.Name,
		Author:      msg.Author,
		RecordType:  msg.RecordType,
		RecordNo:    msg.RecordNo,
		Description: msg.Description,
	}

	err := keeper.CreateRecord(ctx, &recordInfo)
	if err != nil {
		return err.Result()
	}

	return sdk.Result{
		Tags: utils.GetRecordTags(&recordInfo),
	}
}
