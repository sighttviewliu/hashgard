package keeper

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/tendermint/tendermint/crypto"

	"github.com/hashgard/hashgard/x/exchange/types"
)

var (
	ParamsStoreKeyExchangeParams = []byte("exchangeparams")

	// TODO: Find another way to implement this without using accounts, or find a cleaner way to implement it using accounts.
	FrozenCoinsAccAddr = sdk.AccAddress(crypto.AddressHash([]byte("exchangeFrozenCoins")))
)

func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable(
		ParamsStoreKeyExchangeParams, types.ExchangeParams{},
	)
}

type Keeper struct {
	storeKey     sdk.StoreKey
	cdc          *codec.Codec
	paramsKeeper params.Keeper
	paramSpace   params.Subspace
	bankKeeper   types.BankKeeper
	codespace    sdk.CodespaceType
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramsKeeper params.Keeper,
	paramSpace params.Subspace, bankKeeper types.BankKeeper, codespace sdk.CodespaceType) Keeper {
	return Keeper{
		storeKey:     key,
		cdc:          cdc,
		paramsKeeper: paramsKeeper,
		paramSpace:   paramSpace.WithKeyTable(ParamKeyTable()),
		bankKeeper:   bankKeeper,
		codespace:    codespace,
	}
}
func (keeper Keeper) AddOrder(ctx sdk.Context, order types.Order) {
	keeper.setOrder(ctx, order)
	idArr := keeper.GetAddressOrders(ctx, order.Seller)
	idArr = append(idArr, order.OrderId)
	keeper.setAddressOrders(ctx, order.Seller, idArr)
}
func (keeper Keeper) Make(ctx sdk.Context, seller sdk.AccAddress,
	supply sdk.Coin, target sdk.Coin) (order *types.Order, err sdk.Error) {
	id, err := keeper.getNewOrderId(ctx)
	if err != nil {
		return nil, err
	}

	createTime := ctx.BlockHeader().Time

	order = &types.Order{
		OrderId:    id,
		Seller:     seller,
		Supply:     supply,
		Target:     target,
		Remains:    supply,
		CreateTime: createTime,
	}

	err = keeper.bankKeeper.SendCoins(ctx, seller, FrozenCoinsAccAddr, []sdk.Coin{supply})
	if err != nil {
		return nil, err
	}

	keeper.AddOrder(ctx, *order)
	return order, nil
}

func (keeper Keeper) Cancel(ctx sdk.Context, id uint64, addr sdk.AccAddress) (amt sdk.Coin, err sdk.Error) {
	order, err := keeper.GetOrder(ctx, id)
	if err != nil {
		return amt, err
	}

	if !order.Seller.Equals(addr) {
		return amt, sdk.NewError(keeper.codespace, types.CodeNoPermission, fmt.Sprintf("order %d isn't belong to %s", id, addr))
	}

	amt = order.Remains
	err = keeper.bankKeeper.SendCoins(ctx, FrozenCoinsAccAddr, addr, []sdk.Coin{amt})
	if err != nil {
		return amt, err
	}

	keeper.deleteOrder(ctx, id)
	idArr := keeper.GetAddressOrders(ctx, addr)
	for index := 0; index < len(idArr); {
		if idArr[index] == id {
			idArr = append(idArr[:index], idArr[index+1:]...)
			break
		}
		index++
	}

	if len(idArr) == 0 {
		keeper.deleteAddressOrders(ctx, addr)
	} else {
		keeper.setAddressOrders(ctx, addr, idArr)
	}

	return amt, nil
}

func (keeper Keeper) Take(ctx sdk.Context, id uint64, buyer sdk.AccAddress, val sdk.Coin) (supplyTurnover sdk.Coin,
	targetTurnover sdk.Coin, soldOut bool, err sdk.Error) {
	order, err := keeper.GetOrder(ctx, id)
	if err != nil {
		return supplyTurnover, targetTurnover, soldOut, err
	}

	if val.Denom != order.Target.Denom {
		return supplyTurnover, targetTurnover, soldOut, sdk.NewError(keeper.codespace, types.CodeNotMatchTarget, fmt.Sprintf("%s doesn't match order's target(%s)", val.Denom, order.Target.Denom))
	}

	divisor := GetGratestDivisor(order.Supply.Amount, order.Target.Amount)
	supplyPrice := order.Supply.Amount.Quo(divisor)
	remainShares := order.Remains.Amount.Quo(supplyPrice)
	sharePrice := order.Target.Amount.Quo(divisor)

	if val.Amount.LT(sharePrice) {
		return supplyTurnover, targetTurnover, soldOut, sdk.NewError(keeper.codespace, types.CodeTooLess, fmt.Sprintf("minimum purchase threshold is %s%s", sharePrice.String(), order.Target.Denom))
	}

	shares := val.Amount.Quo(sharePrice)

	if shares.GTE(remainShares) {
		shares = remainShares
		soldOut = true
	}

	supplyTurnover = sdk.NewCoin(order.Supply.Denom, order.Supply.Amount.Quo(divisor).Mul(shares))
	targetTurnover = sdk.NewCoin(order.Target.Denom, sharePrice.Mul(shares))

	err = keeper.bankKeeper.SendCoins(ctx, buyer, order.Seller, []sdk.Coin{targetTurnover})
	if err != nil {
		return supplyTurnover, targetTurnover, soldOut, err
	}
	err = keeper.bankKeeper.SendCoins(ctx, FrozenCoinsAccAddr, buyer, []sdk.Coin{supplyTurnover})
	if err != nil {
		return supplyTurnover, targetTurnover, soldOut, err
	}

	if soldOut {
		keeper.deleteOrder(ctx, id)
		idArr := keeper.GetAddressOrders(ctx, order.Seller)
		for index := 0; index < len(idArr); {
			if idArr[index] == id {
				idArr = append(idArr[:index], idArr[index+1:]...)
				break
			}
			index++
		}
		if len(idArr) == 0 {
			keeper.deleteAddressOrders(ctx, order.Seller)
		} else {
			keeper.setAddressOrders(ctx, order.Seller, idArr)
		}
	} else {
		remains := order.Remains.Sub(supplyTurnover)
		newOrder := types.Order{
			OrderId:    id,
			Seller:     order.Seller,
			Supply:     order.Supply,
			Target:     order.Target,
			Remains:    remains,
			CreateTime: order.CreateTime,
		}
		keeper.setOrder(ctx, newOrder)
	}

	return supplyTurnover, targetTurnover, soldOut, nil
}

func (keeper Keeper) GetOrdersByAddr(ctx sdk.Context, addr sdk.AccAddress) (orders types.Orders, err sdk.Error) {
	idArr := keeper.GetAddressOrders(ctx, addr)
	for _, id := range idArr {
		order, err := keeper.GetOrder(ctx, id)
		if err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}

func (keeper Keeper) GetFrozenFundByAddr(ctx sdk.Context, addr sdk.AccAddress) (fund sdk.Coins, err sdk.Error) {
	idArr := keeper.GetAddressOrders(ctx, addr)
	for _, id := range idArr {
		order, err := keeper.GetOrder(ctx, id)
		if err != nil {
			return sdk.NewCoins(), err
		}
		fund = fund.Add([]sdk.Coin{order.Remains})
	}

	return fund, nil
}

// Store level
func (keeper Keeper) GetOrder(ctx sdk.Context, id uint64) (*types.Order, sdk.Error) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyOrder(id))
	if bz == nil {
		return nil, sdk.NewError(keeper.codespace, types.CodeOrderNotExist, fmt.Sprintf("this id is invalid : %d", id))
	}
	var order types.Order
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &order)
	return &order, nil
}

func (keeper Keeper) setOrder(ctx sdk.Context, order types.Order) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(order)
	store.Set(KeyOrder(order.OrderId), bz)
}

func (keeper Keeper) SetOrder(ctx sdk.Context, order types.Order) {
	keeper.setOrder(ctx, order)
}

func (keeper Keeper) deleteOrder(ctx sdk.Context, id uint64) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(KeyOrder(id))
}

func (keeper Keeper) HasOrder(ctx sdk.Context, id uint64) bool {
	store := ctx.KVStore(keeper.storeKey)
	return store.Has(KeyOrder(id))
}

// Get the next available OrderId and increments it
func (keeper Keeper) getNewOrderId(ctx sdk.Context) (id uint64, err sdk.Error) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextOrderId)
	if bz == nil {
		return 0, sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "InitialOrderId never set")
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &id)
	bz = keeper.cdc.MustMarshalBinaryLengthPrefixed(id + 1)
	store.Set(KeyNextOrderId, bz)
	return id, nil
}

// Peeks the next available id without incrementing it
func (keeper Keeper) PeekCurrentOrderId(ctx sdk.Context) (id uint64, err sdk.Error) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextOrderId)
	if bz == nil {
		return 0, sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "InitialOrderId never set")
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &id)
	return id, nil
}

// Set the initial order ID
func (keeper Keeper) SetInitialOrderId(ctx sdk.Context, id uint64) sdk.Error {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyNextOrderId)
	if bz != nil {
		return sdk.NewError(keeper.codespace, types.CodeInvalidGenesis, "Initial ProposalID already set")
	}
	bz = keeper.cdc.MustMarshalBinaryLengthPrefixed(id)
	store.Set(KeyNextOrderId, bz)
	return nil
}

func (keeper Keeper) GetAddressOrders(ctx sdk.Context, addr sdk.AccAddress) (idArr []uint64) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(KeyAddressOrders(addr))
	if bz == nil {
		return []uint64{}
	}
	keeper.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &idArr)
	return idArr
}

func (keeper Keeper) setAddressOrders(ctx sdk.Context, addr sdk.AccAddress, idArr []uint64) {
	store := ctx.KVStore(keeper.storeKey)
	bz := keeper.cdc.MustMarshalBinaryLengthPrefixed(idArr)
	store.Set(KeyAddressOrders(addr), bz)
}

func (keeper Keeper) SetAddressOrders(ctx sdk.Context, addr sdk.AccAddress, idArr []uint64) {
	keeper.setAddressOrders(ctx, addr, idArr)
}

func (keeper Keeper) deleteAddressOrders(ctx sdk.Context, addr sdk.AccAddress) {
	store := ctx.KVStore(keeper.storeKey)
	store.Delete(KeyAddressOrders(addr))
}

// Get Order from store by OrderId
// seller will filter orders by creator
// supplyDenom will filter orders by supply token denom
// targetDenom will filter orders by target token denom
// numLatest will fetch a specified number of the most recent orders, or 0 for all orders
func (keeper Keeper) GetOrdersFiltered(ctx sdk.Context,
	seller sdk.AccAddress, supplyDenom string, targetDenom string, numLatest uint64) []*types.Order {

	maxOrderId, err := keeper.PeekCurrentOrderId(ctx)
	if err != nil {
		return nil
	}
	matchOrders := make([]*types.Order, 0)
	if numLatest == 0 {
		numLatest = maxOrderId
	}
	for id := maxOrderId - numLatest; id < maxOrderId; id++ {
		order, err := keeper.GetOrder(ctx, id+types.DefaultStartingOrderId)
		if err != nil {
			continue
		}
		if seller != nil && len(seller) != 0 {
			if !bytes.Equal(order.Seller, seller) {
				continue
			}
		}
		if supplyDenom != "" {
			if order.Supply.Denom != supplyDenom {
				continue
			}
		}
		if targetDenom != "" {
			if order.Target.Denom != targetDenom {
				continue
			}
		}
		matchOrders = append(matchOrders, order)
	}
	return matchOrders
}
