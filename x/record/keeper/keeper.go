package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/hashgard/hashgard/x/record/errors"
	"github.com/hashgard/hashgard/x/record/types"
	recordparams "github.com/hashgard/hashgard/x/record/params"
)

var (
	ParamsStoreKeyRecordParams = []byte("recordparams")
)
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamsStoreKeyRecordParams, recordparams.RecordParams{},
	)
}

// Issue Keeper
type Keeper struct {
	// The reference to the Param Keeper to get and set Global Params
	paramsKeeper params.Keeper
	// The reference to the Paramstore to get and set issue specific params
	paramSpace params.Subspace
	// The (unexposed) keys used to access the stores from the Context.
	storeKey sdk.StoreKey
	// The codec codec for binary encoding/decoding.
	cdc *codec.Codec
	// Reserved codespace
	codespace sdk.CodespaceType
}

//Get issue codec
func (keeper Keeper) Getcdc() *codec.Codec {
	return keeper.cdc
}

//New record keeper Instance
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramsKeeper params.Keeper,
	paramSpace params.Subspace, codespace sdk.CodespaceType) Keeper {
	return Keeper{
		storeKey:            key,
		paramsKeeper:        paramsKeeper,
		paramSpace:   		 paramSpace.WithKeyTable(ParamKeyTable()),
		cdc:                 cdc,
		codespace:           codespace,
	}
}

//add a record
func (keeper Keeper) AddRecord(ctx sdk.Context, recordInfo *types.RecordInfo) {
	store := ctx.KVStore(keeper.storeKey)

	// save recordInfo with key: record.id
	store.Set(KeyRecordId(recordInfo.ID), keeper.cdc.MustMarshalBinaryLengthPrefixed(recordInfo))

	// save record.id with key: record.hash
	store.Set(KeyRecord(recordInfo.Hash), keeper.cdc.MustMarshalBinaryLengthPrefixed(recordInfo.ID))
}
//Create a record
func (keeper Keeper) CreateRecord(ctx sdk.Context, recordInfo *types.RecordInfo) sdk.Error {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyRecord(recordInfo.Hash))
	if len(bz) > 0 {
		return errors.ErrRecordExist(recordInfo.Hash)
	}
	id, err := keeper.getNewRecordID(store)
	if err != nil {
		return err
	}
	recordID := KeyRecordIdStr(id)
	recordInfo.RecordTime = ctx.BlockHeader().Time.Unix()
	recordInfo.ID = recordID

	keeper.AddRecord(ctx, recordInfo)

	return nil
}

//Returns record by Hash
func (keeper Keeper) GetRecord(ctx sdk.Context, record string) *types.RecordInfo {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyRecord(record))
	if len(bz) == 0 {
		return nil
	}
	var recordId string
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &recordId)
	return keeper.getRecordByID(ctx, recordId)
}
func (keeper Keeper) getRecordByID(ctx sdk.Context, id string) *types.RecordInfo {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyRecordId(id))
	if len(bz) == 0 {
		return nil
	}
	var recordInfo types.RecordInfo
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &recordInfo)
	return &recordInfo
}

func (keeper Keeper) List(ctx sdk.Context, params recordparams.RecordQueryParams) []*types.RecordInfo {
	iterator := keeper.Iterator(ctx, params.StartRecordId)
	defer iterator.Close()
	if params.Limit < 1 {
		params.Limit = 20
	}
	list := make([]*types.RecordInfo, 0, params.Limit)
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) == 0 {
			continue
		}
		var info types.RecordInfo
		keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &info)
		if params.Sender != nil && !params.Sender.Empty() && params.Author != "" {
			if info.Sender.Equals(params.Sender) && info.Author == params.Author {
				list = append(list, &info)
			}
			continue
		}
		if params.Sender != nil && !params.Sender.Empty() {
			if info.Sender.Equals(params.Sender) {
				list = append(list, &info)
			}
			continue
		}
		if params.Author != "" {
			if info.Author == params.Author {
				list = append(list, &info)
			}
			continue
		}
		list = append(list, &info)
		if len(list) >= params.Limit {
			break
		}
	}
	return list
}
func (keeper Keeper) Iterator(ctx sdk.Context, startRecordId string) sdk.Iterator {
	store := ctx.KVStore(keeper.storeKey)
	endRecordId := startRecordId

	if len(startRecordId) == 0 {
		endRecordId = KeyRecordIdStr(types.RecordMaxId)
		startRecordId = KeyRecordIdStr(types.RecordMinId - 1)
	} else {
		startRecordId = KeyRecordIdStr(types.RecordMinId - 1)
	}
	iterator := store.ReverseIterator(KeyRecordId(startRecordId), KeyRecordId(endRecordId))
	return iterator
}

//Set the initial record id
func (keeper Keeper) SetInitialRecordStartingRecordId(ctx sdk.Context, recordID uint64) sdk.Error {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextRecordID)
	if bz != nil {
		return sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "Initial RecordId already set")
	}
	bz = keeper.cdc.MustMarshalBinaryLengthPrefixed(recordID)
	store.Set(KeyNextRecordID, bz)
	return nil
}

// Get the last used recordID
func (keeper Keeper) GetLastRecordID(ctx sdk.Context) (recordID uint64) {
	recordID, err := keeper.PeekCurrentRecordID(ctx)
	if err != nil {
		return 0
	}
	recordID--
	return
}

// Gets the next available issueID and increase it
func (keeper Keeper) getNewRecordID(store sdk.KVStore) (recordID uint64, err sdk.Error) {
	bz := store.Get(KeyNextRecordID)
	if bz == nil {
		return 0, sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "InitialRecordID never set")
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &recordID)
	bz = keeper.cdc.MustMarshalBinaryLengthPrefixed(recordID + 1)
	store.Set(KeyNextRecordID, bz)
	return recordID, nil
}

// Peeks the next available recordID without increasing it
func (keeper Keeper) PeekCurrentRecordID(ctx sdk.Context) (recordID uint64, err sdk.Error) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextRecordID)
	if bz == nil {
		return 0, sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "InitialRecordID never set")
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &recordID)
	return recordID, nil
}
