package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"gitlab.com/rarimo/rarimo-core/x/tokenmanager/types"
)

// SetInfo set a specific info in the store from its index
func (k Keeper) SetInfo(ctx sdk.Context, info types.Info) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InfoKeyPrefix))
	b := k.cdc.MustMarshal(&info)
	store.Set(types.InfoKey(
		info.Index,
	), b)
}

// GetInfo returns a info from its index
func (k Keeper) GetInfo(
	ctx sdk.Context,
	index string,

) (val types.Info, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InfoKeyPrefix))

	b := store.Get(types.InfoKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveInfo removes a info from the store
func (k Keeper) RemoveInfo(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InfoKeyPrefix))
	store.Delete(types.InfoKey(
		index,
	))
}

// GetAllInfo returns all info
func (k Keeper) GetAllInfo(ctx sdk.Context) (list []types.Info) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.InfoKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Info
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}