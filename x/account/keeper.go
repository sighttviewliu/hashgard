package account

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"strings"
)

var (
	ParamsStoreKeyRecordParams = []byte("accountparams")
)

func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamsStoreKeyRecordParams, false,
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
func NewKeeper(cdc *codec.Codec, key sdk.StoreKey,
	paramSpace params.Subspace, codespace sdk.CodespaceType) Keeper {
	return Keeper{
		storeKey:   key,
		paramSpace: paramSpace.WithKeyTable(ParamKeyTable()),
		cdc:        cdc,
		codespace:  codespace,
	}
}

func KeyMustMemoAddress(addr string) []byte {
	return []byte(fmt.Sprintf("addr:%s", addr))
}
func PrefixMustMemoAddress() []byte {
	return []byte("addr:")
}

// set must memo
func (keeper Keeper) SetMustMemoAddress(ctx sdk.Context, acc *AccountInfo) {
	store := ctx.KVStore(keeper.storeKey)

	if acc.MustMemo {
		store.Set(KeyMustMemoAddress(acc.Sender.String()), keeper.cdc.MustMarshalBinaryLengthPrefixed(1))
	} else {
		store.Delete(KeyMustMemoAddress(acc.Sender.String()))
	}
}

// get must memo status of address
func (keeper Keeper) GetMustMemoAddress(ctx sdk.Context, addr string) (*AccountInfo, error) {
	store := ctx.KVStore(keeper.storeKey)

	address, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, err
	}

	var res = AccountInfo{
		Sender:   address,
		MustMemo: false,
	}

	bz := store.Get(KeyMustMemoAddress(addr))
	if len(bz) > 0 {
		res.MustMemo = true
	}
	return &res, nil
}
func (keeper Keeper) GetMustMemoAddressList(ctx sdk.Context) []AccountInfo {
	store := ctx.KVStore(keeper.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, PrefixMustMemoAddress())
	defer iterator.Close()
	list := make([]AccountInfo, 0)
	for ; iterator.Valid(); iterator.Next() {
		bz := iterator.Value()
		if len(bz) > 0 {
			addr, err := sdk.AccAddressFromBech32(strings.Split(string(iterator.Key()[:]), ":")[1])
			if err != nil {
				continue
			}
			list = append(list, AccountInfo{
				Sender:   addr,
				MustMemo: true,
			})
		}
	}
	return list
}
