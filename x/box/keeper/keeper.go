package keeper

import (
	"fmt"

	"github.com/hashgard/hashgard/x/box/config"

	"github.com/hashgard/hashgard/x/box/utils"

	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/hashgard/hashgard/x/box/errors"
	boxparams "github.com/hashgard/hashgard/x/box/params"
	"github.com/hashgard/hashgard/x/box/types"
	issueerr "github.com/hashgard/hashgard/x/issue/errors"
)

// Box Keeper
type Keeper struct {
	// The reference to the Param Keeper to get and set Global Params
	paramsKeeper params.Keeper
	// The reference to the Paramstore to get and set box specific params
	paramSpace params.Subspace
	// The (unexposed) keys used to access the stores from the Context.
	storeKey sdk.StoreKey
	// The reference to the CoinKeeper to modify balances
	ck BankKeeper
	// The reference to the IssueKeeper to get issue info
	ik IssueKeeper
	// The reference to the FeeCollectionKeeper to add fee
	feeCollectionKeeper FeeCollectionKeeper
	// The codec codec for binary encoding/decoding.
	cdc *codec.Codec
	// Reserved codespace
	codespace sdk.CodespaceType
}

//Get box codec
func (keeper Keeper) Getcdc() *codec.Codec {
	return keeper.cdc
}

//Get box bankKeeper
func (keeper Keeper) GetBankKeeper() BankKeeper {
	return keeper.ck
}

//Get box issueKeeper
func (keeper Keeper) GetIssueKeeper() IssueKeeper {
	return keeper.ik
}

//Get box feeCollectionKeeper
func (keeper Keeper) GetFeeCollectionKeeper() FeeCollectionKeeper {
	return keeper.feeCollectionKeeper
}

//New box keeper Instance
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramsKeeper params.Keeper,
	paramSpace params.Subspace, ck BankKeeper, ik IssueKeeper,
	feeCollectionKeeper FeeCollectionKeeper, codespace sdk.CodespaceType) Keeper {
	return Keeper{
		storeKey:            key,
		paramsKeeper:        paramsKeeper,
		paramSpace:          paramSpace.WithKeyTable(config.ParamKeyTable()),
		ck:                  ck,
		ik:                  ik,
		feeCollectionKeeper: feeCollectionKeeper,
		cdc:                 cdc,
		codespace:           codespace,
	}
}

func (keeper Keeper) getDepositedCoinsAddress(id string) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(fmt.Sprintf("boxDepositedCoins:%s", id))))

}
func (keeper Keeper) SendDepositedCoin(ctx sdk.Context, fromAddr sdk.AccAddress, amt sdk.Coins, id string) sdk.Error {
	toAddr := keeper.getDepositedCoinsAddress(id)
	return keeper.GetBankKeeper().SendCoins(ctx, fromAddr, toAddr, amt)
}
func (keeper Keeper) CancelDepositedCoin(ctx sdk.Context, toAddr sdk.AccAddress, amt sdk.Coins, id string) sdk.Error {
	fromAddr := keeper.getDepositedCoinsAddress(id)
	return keeper.GetBankKeeper().SendCoins(ctx, fromAddr, toAddr, amt)
}
func (keeper Keeper) SubDepositedCoin(ctx sdk.Context, amt sdk.Coins, id string) sdk.Error {
	_, err := keeper.GetBankKeeper().SubtractCoins(ctx, keeper.getDepositedCoinsAddress(id), amt)
	return err
}
func (keeper Keeper) GetDepositedCoins(ctx sdk.Context, id string) sdk.Coins {
	return keeper.GetBankKeeper().GetCoins(ctx, keeper.getDepositedCoinsAddress(id))
}

//Keys set
//Set box
func (keeper Keeper) setBox(ctx sdk.Context, box *types.BoxInfo) {
	store := ctx.KVStore(keeper.storeKey)
	store.Set(KeyBox(box.Id), keeper.cdc.MustMarshalBinaryLengthPrefixed(box))
}

//Set address
func (keeper Keeper) setAddress(ctx sdk.Context, boxType string, accAddress sdk.AccAddress, ids []string) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(ids)
	store.Set(KeyAddress(boxType, accAddress), bz)
}

//Set name
func (keeper Keeper) setName(ctx sdk.Context, boxType string, name string, ids []string) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(ids)
	store.Set(KeyName(boxType, name), bz)
}

//Keys Add
//Add box
func (keeper Keeper) AddBox(ctx sdk.Context, box *types.BoxInfo) {
	ids := keeper.GetIdsByAddress(ctx, box.BoxType, box.Owner)
	ids = append(ids, box.Id)
	keeper.setAddress(ctx, box.BoxType, box.Owner, ids)

	ids = keeper.GetIdsByName(ctx, box.BoxType, box.Name)
	ids = append(ids, box.Id)
	keeper.setName(ctx, box.BoxType, box.Name, ids)

	keeper.setBox(ctx, box)
}

//Keys remove
//Remove box
//func (keeper Keeper) RemoveBox(ctx sdk.Context, box *types.BoxInfo) {
//	store := ctx.KVStore(keeper.storeKey)
//	store.Delete(KeyName(box.BoxType, box.Name))
//	store.Delete(KeyAddress(box.BoxType, box.Owner))
//	store.Delete(KeyBox(box.Id))
//}
func (keeper Keeper) GetCoinDecimals(ctx sdk.Context, coin sdk.Coin) (uint, sdk.Error) {
	if coin.Denom == types.Agard {
		return types.AgardDecimals, nil
	}
	coinIssueInfo := keeper.GetIssueKeeper().GetIssue(ctx, coin.Denom)
	if coinIssueInfo == nil {
		return 0, issueerr.ErrUnknownIssue(coin.Denom)
	}
	return coinIssueInfo.Decimals, nil
}

func (keeper Keeper) Fee(ctx sdk.Context, sender sdk.AccAddress, fee sdk.Coin) sdk.Error {
	if fee.IsZero() || fee.IsNegative() {
		return nil
	}
	_, err := keeper.GetBankKeeper().SubtractCoins(ctx, sender, sdk.NewCoins(fee))
	if err != nil {
		return errors.ErrNotEnoughFee()
	}
	_ = keeper.GetFeeCollectionKeeper().AddCollectedFees(ctx, sdk.NewCoins(fee))
	return nil
}

//Keys return
//Return box by id
func (keeper Keeper) GetBox(ctx sdk.Context, id string) *types.BoxInfo {
	id = utils.GetIdFromBoxSeqID(id)
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyBox(id))
	if len(bz) == 0 {
		return nil
	}
	var box types.BoxInfo
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &box)
	return &box
}

//Return box by id and and check owner
func (keeper Keeper) GetBoxByOwner(ctx sdk.Context, sender sdk.AccAddress, id string) (*types.BoxInfo, sdk.Error) {
	box := keeper.GetBox(ctx, id)
	if box == nil {
		return nil, errors.ErrUnknownBox(id)
	}
	if !box.Owner.Equals(sender) {
		return nil, errors.ErrOwnerMismatch(id)
	}
	return box, nil
}

//Returns box list by type and accAddress
func (keeper Keeper) GetBoxByAddress(ctx sdk.Context, boxType string, accAddress sdk.AccAddress) []*types.BoxInfo {
	ids := keeper.GetIdsByAddress(ctx, boxType, accAddress)
	length := len(ids)
	if length == 0 {
		return []*types.BoxInfo{}
	}
	boxs := make([]*types.BoxInfo, 0, length)

	for _, v := range ids {
		boxs = append(boxs, keeper.GetBox(ctx, v))
	}
	return boxs
}
func (keeper Keeper) CanTransfer(ctx sdk.Context, id string) sdk.Error {
	if !utils.IsId(id) {
		return nil
	}
	box := keeper.GetBox(ctx, id)
	if box == nil {
		return nil
	}
	if box.IsTransferDisabled() {
		return errors.ErrCanNotTransfer(id)
	}
	return nil
}

//Queries

//Search box by name
func (keeper Keeper) SearchBox(ctx sdk.Context, boxType string, name string) []*types.BoxInfo {
	store := ctx.KVStore(keeper.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, KeyName(boxType, name))
	defer iterator.Close()
	list := make([]*types.BoxInfo, 0, 1)
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) == 0 {
			continue
		}
		ids := make([]string, 0, 1)
		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &ids)

		for _, v := range ids {
			list = append(list, keeper.GetBox(ctx, v))
		}
	}
	return list
}

//List
func (keeper Keeper) List(ctx sdk.Context, params boxparams.BoxQueryParams) []*types.BoxInfo {
	if params.Owner != nil && !params.Owner.Empty() {
		return keeper.GetBoxByAddress(ctx, params.BoxType, params.Owner)
	}
	iterator := keeper.Iterator(ctx, params.BoxType, params.StartId)
	defer iterator.Close()
	list := make([]*types.BoxInfo, 0, params.Limit)
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) == 0 {
			continue
		}
		var info types.BoxInfo
		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &info)
		list = append(list, &info)
		if len(list) >= params.Limit {
			break
		}
	}
	return list
}
func (keeper Keeper) ListAll(ctx sdk.Context, boxType string) []types.BoxInfo {
	iterator := keeper.Iterator(ctx, boxType, "")
	defer iterator.Close()
	list := make([]types.BoxInfo, 0)
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) == 0 {
			continue
		}
		var info types.BoxInfo
		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &info)
		list = append(list, info)
	}
	return list
}
func (keeper Keeper) Iterator(ctx sdk.Context, boxType string, startId string) sdk.Iterator {
	store := ctx.KVStore(keeper.storeKey)
	endId := startId
	if len(startId) == 0 {
		endId = KeyIdStr(boxType, types.BoxMaxId)
		startId = KeyIdStr(boxType, types.BoxMinId-1)
	} else {
		startId = KeyIdStr(boxType, types.BoxMinId-1)
	}
	iterator := store.ReverseIterator(KeyBox(startId), KeyBox(endId))
	return iterator
}

//Create a box
func (keeper Keeper) CreateBox(ctx sdk.Context, box *types.BoxInfo) sdk.Error {
	decimal, err := keeper.GetCoinDecimals(ctx, box.TotalAmount.Token)
	if err != nil {
		return err
	}
	if box.TotalAmount.Decimals != decimal {
		return errors.ErrDecimalsNotValid(box.TotalAmount.Decimals)
	}
	store := ctx.KVStore(keeper.storeKey)
	id, err := keeper.getNewBoxID(store, box.BoxType)
	if err != nil {
		return err
	}
	box.Id = KeyIdStr(box.BoxType, id)
	box.CreatedTime = ctx.BlockHeader().Time.Unix()

	switch box.BoxType {
	case types.Lock:
		err = keeper.ProcessLockBoxCreate(ctx, box)
	case types.Deposit:
		err = keeper.ProcessDepositBoxCreate(ctx, box)
	case types.Future:
		err = keeper.ProcessFutureBoxCreate(ctx, box)
	default:
		return errors.ErrUnknownBoxType()
	}
	if err != nil {
		return err
	}
	keeper.AddBox(ctx, box)
	return nil
}

func (keeper Keeper) ProcessInjectBox(ctx sdk.Context, id string, sender sdk.AccAddress, amount sdk.Coin, operation string) (*types.BoxInfo, sdk.Error) {
	box := keeper.GetBox(ctx, id)
	if box == nil {
		return nil, errors.ErrUnknownBox(id)
	}
	if types.BoxInjecting != box.Status && types.BoxClosed != box.Status {
		return nil, errors.ErrNotAllowedOperation(box.Status)
	}
	switch box.BoxType {
	case types.Deposit:
		return box, keeper.processDepositBoxInject(ctx, box, sender, amount, operation)
	case types.Future:
		return box, keeper.processFutureBoxInject(ctx, box, sender, amount, operation)
	}
	return nil, errors.ErrUnknownBoxType()
}
func (keeper Keeper) ProcessBoxWithdraw(ctx sdk.Context, id string, sender sdk.AccAddress) (sdk.Int, *types.BoxInfo, sdk.Error) {
	if keeper.GetBankKeeper().GetCoins(ctx, sender).AmountOf(id).IsZero() {
		return sdk.ZeroInt(), nil, errors.ErrNotEnoughAmount()
	}
	boxType := utils.GetBoxTypeByValue(id)
	switch boxType {
	case types.Deposit:
		return keeper.processDepositBoxWithdraw(ctx, id, sender)
	case types.Future:
		boxInfo, err := keeper.processFutureBoxWithdraw(ctx, id, sender)
		return sdk.ZeroInt(), boxInfo, err
	}
	return sdk.ZeroInt(), nil, errors.ErrUnknownBoxType()
}

func (keeper Keeper) SetBoxDescription(ctx sdk.Context, id string, sender sdk.AccAddress, description []byte) (*types.BoxInfo, sdk.Error) {
	box, err := keeper.GetBoxByOwner(ctx, sender, id)
	if err != nil {
		return box, err
	}
	box.Description = string(description)
	keeper.setBox(ctx, box)
	return box, nil
}
func (keeper Keeper) DisableFeature(ctx sdk.Context, sender sdk.AccAddress, id string, feature string) (*types.BoxInfo, sdk.Error) {
	boxInfo, err := keeper.GetBoxByOwner(ctx, sender, id)
	if err != nil {
		return nil, err
	}
	switch feature {
	case types.Transfer:
		return boxInfo, keeper.disableTransfer(ctx, sender, boxInfo)
	default:
		return nil, errors.ErrUnknownFeatures()
	}
}
func (keeper Keeper) disableTransfer(ctx sdk.Context, sender sdk.AccAddress, boxInfo *types.BoxInfo) sdk.Error {
	if boxInfo.GetBoxType() == types.Lock {
		return errors.ErrNotSupportOperation()
	}
	if !boxInfo.IsTransferDisabled() {
		return nil
	}
	boxInfo.TransferDisabled = false
	keeper.setBox(ctx, boxInfo)
	return nil
}

//Send coins
func (keeper Keeper) SendCoins(ctx sdk.Context,
	fromAddr sdk.AccAddress, toAddr sdk.AccAddress,
	amt sdk.Coins) sdk.Error {
	return keeper.ck.SendCoins(ctx, fromAddr, toAddr, amt)
}

//Get name from a box
func (keeper Keeper) GetIdsByName(ctx sdk.Context, boxType string, name string) (ids []string) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyName(boxType, name))
	if bz == nil {
		return []string{}
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &ids)
	return ids
}

//Get address from a box
func (keeper Keeper) GetIdsByAddress(ctx sdk.Context, boxType string, accAddress sdk.AccAddress) (ids []string) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyAddress(boxType, accAddress))
	if bz == nil {
		return []string{}
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &ids)
	return ids
}

// BoxQueues

// Returns an iterator for all the box in the Active Queue that expire by time
func (keeper Keeper) ActiveBoxQueueIterator(ctx sdk.Context, endTime int64) sdk.Iterator {
	store := ctx.KVStore(keeper.storeKey)
	return store.Iterator(PrefixActiveQueue, sdk.PrefixEndBytes(PrefixActiveBoxQueueTime(endTime)))
}

// Inserts a id into the active box queue at time
func (keeper Keeper) InsertActiveBoxQueue(ctx sdk.Context, endTime int64, boxIdStr string) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(boxIdStr)
	store.Set(KeyActiveBoxQueue(endTime, boxIdStr), bz)
}

// removes a id from the Active box Queue
func (keeper Keeper) RemoveFromActiveBoxQueue(ctx sdk.Context, endTime int64, boxIdStr string) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(KeyActiveBoxQueue(endTime, boxIdStr))
}
func (keeper Keeper) RemoveFromActiveBoxQueueByKey(ctx sdk.Context, key []byte) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(key)
}

// -----------------------------------------------------------------------------
// Params

// SetParams sets the auth module's parameters.
func (ak Keeper) SetParams(ctx sdk.Context, params config.Params) {
	ak.paramSpace.SetParamSet(ctx, &params)
}

// GetParams gets the auth module's parameters.
func (ak Keeper) GetParams(ctx sdk.Context) (params config.Params) {
	ak.paramSpace.GetParamSet(ctx, &params)
	return
}

//Set the initial boxCount
func (keeper Keeper) SetInitialBoxStartingId(ctx sdk.Context, boxType string, id uint64) sdk.Error {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextBoxID(boxType))
	if bz != nil {
		return sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "Initial Id already set")
	}
	bz = keeper.cdc.MustMarshalBinaryLengthPrefixed(id)
	store.Set(KeyNextBoxID(boxType), bz)
	return nil
}

// Get the last used id
func (keeper Keeper) GetLastBoxID(ctx sdk.Context, boxType string) (id uint64) {
	id, err := keeper.PeekCurrentBoxID(ctx, boxType)
	if err != nil {
		return 0
	}
	id--
	return
}

// Gets the next available id and increments it
func (keeper Keeper) getNewBoxID(store sdk.KVStore, boxType string) (id uint64, err sdk.Error) {
	bz := store.Get(KeyNextBoxID(boxType))
	if bz == nil {
		return 0, sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "InitialBoxID never set")
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &id)
	bz = keeper.cdc.MustMarshalBinaryLengthPrefixed(id + 1)
	store.Set(KeyNextBoxID(boxType), bz)
	return id, nil
}

// Peeks the next available BoxID without incrementing it
func (keeper Keeper) PeekCurrentBoxID(ctx sdk.Context, boxType string) (id uint64, err sdk.Error) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextBoxID(boxType))
	if bz == nil {
		return 0, sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "InitialBoxID never set")
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &id)
	return id, nil
}
