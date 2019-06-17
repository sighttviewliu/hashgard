package box

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/tags"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/hashgard/hashgard/x/box/types"
)

// Called every block, process inflation, update validator set
func EndBlocker(ctx sdk.Context, keeper Keeper) sdk.Tags {
	logger := ctx.Logger().With("module", "x/"+types.ModuleName)
	resTags := sdk.NewTags()
	// fetch active proposals whose voting periods have ended (are passed the block time)
	activeIterator := keeper.ActiveBoxQueueIterator(ctx, ctx.BlockHeader().Time.Unix())
	defer activeIterator.Close()
	count := 0
	for ; activeIterator.Valid(); activeIterator.Next() {
		if count == types.BoxMaxInstalment {
			break
		}
		var id string
		//var seq int
		keeper.Getcdc().MustUnmarshalBinaryLengthPrefixed(activeIterator.Value(), &id)
		//if strings.Contains(id, types.KeyDelimiterString) {
		//	ids := strings.Split(id, types.KeyDelimiterString)
		//	id = ids[0]
		//	seq, _ = strconv.Atoi(ids[1])
		//}
		boxInfo := keeper.GetBox(ctx, id)
		if boxInfo == nil {
			panic(fmt.Sprintf("box %s does not exist", id))
			//continue
		}
		switch boxInfo.BoxType {
		case types.Lock:
			if err := keeper.ProcessLockBoxByEndBlocker(ctx, boxInfo); err != nil {
				panic(err)
			}
			logger.Debug(fmt.Sprintf("lockbox %s (%s) unlocked", id, boxInfo.Name))
			resTags = resTags.AppendTag(tags.BoxID, id).AppendTag(tags.Category, boxInfo.GetBoxType()).AppendTag(tags.Status, types.LockBoxUnlocked)
		case types.Deposit:
			if err := keeper.ProcessDepositBoxByEndBlocker(ctx, boxInfo); err != nil {
				panic(err)
			}
			logger.Debug(fmt.Sprintf("depositbox %s (%s) status:%s", id, boxInfo.Name, boxInfo.Status))
			resTags = resTags.AppendTag(tags.BoxID, id).AppendTag(tags.Category, boxInfo.GetBoxType()).AppendTag(tags.Status, boxInfo.Status)
		case types.Future:
			if err := keeper.ProcessFutureBoxByEndBlocker(ctx, boxInfo); err != nil {
				panic(err)
			}
			logger.Debug(fmt.Sprintf("futurebox %s (%s) status:%s", id, boxInfo.Name, boxInfo.Status))
			resTags = resTags.AppendTag(tags.BoxID, id).AppendTag(tags.Category, boxInfo.GetBoxType()).AppendTag(tags.Status, boxInfo.Status)
		}
		count = count + 1
	}
	return resTags
}
