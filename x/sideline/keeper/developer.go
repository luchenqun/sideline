package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

// SetDeveloper set a specific developer in the store from its index
func (k Keeper) SetDeveloper(ctx sdk.Context, developer types.Developer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeveloperKeyPrefix))
	b := k.cdc.MustMarshal(&developer)
	store.Set(types.DeveloperKey(
		developer.Index,
	), b)
}

// GetDeveloper returns a developer from its index
func (k Keeper) GetDeveloper(
	ctx sdk.Context,
	index string,

) (val types.Developer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeveloperKeyPrefix))

	b := store.Get(types.DeveloperKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveDeveloper removes a developer from the store
func (k Keeper) RemoveDeveloper(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeveloperKeyPrefix))
	store.Delete(types.DeveloperKey(
		index,
	))
}

// GetAllDeveloper returns all developer
func (k Keeper) GetAllDeveloper(ctx sdk.Context) (list []types.Developer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.DeveloperKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Developer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
