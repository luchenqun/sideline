package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "sideline/testutil/keeper"
	"sideline/testutil/nullify"
	"sideline/x/sideline/keeper"
	"sideline/x/sideline/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNDeveloper(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Developer {
	items := make([]types.Developer, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetDeveloper(ctx, items[i])
	}
	return items
}

func TestDeveloperGet(t *testing.T) {
	keeper, ctx := keepertest.SidelineKeeper(t)
	items := createNDeveloper(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetDeveloper(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestDeveloperRemove(t *testing.T) {
	keeper, ctx := keepertest.SidelineKeeper(t)
	items := createNDeveloper(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveDeveloper(ctx,
			item.Index,
		)
		_, found := keeper.GetDeveloper(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestDeveloperGetAll(t *testing.T) {
	keeper, ctx := keepertest.SidelineKeeper(t)
	items := createNDeveloper(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllDeveloper(ctx)),
	)
}
