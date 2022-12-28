package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

// SetEmployer set a specific employer in the store from its index
func (k Keeper) SetEmployer(ctx sdk.Context, employer types.Employer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EmployerKeyPrefix))
	b := k.cdc.MustMarshal(&employer)
	store.Set(types.EmployerKey(
		employer.Index,
	), b)
}

// GetEmployer returns a employer from its index
func (k Keeper) GetEmployer(
	ctx sdk.Context,
	index string,

) (val types.Employer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EmployerKeyPrefix))

	b := store.Get(types.EmployerKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveEmployer removes a employer from the store
func (k Keeper) RemoveEmployer(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EmployerKeyPrefix))
	store.Delete(types.EmployerKey(
		index,
	))
}

// GetAllEmployer returns all employer
func (k Keeper) GetAllEmployer(ctx sdk.Context) (list []types.Employer) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.EmployerKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Employer
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return list
}
