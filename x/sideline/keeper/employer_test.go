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

func createNEmployer(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Employer {
	items := make([]types.Employer, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetEmployer(ctx, items[i])
	}
	return items
}

func TestEmployerGet(t *testing.T) {
	keeper, ctx := keepertest.SidelineKeeper(t)
	items := createNEmployer(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetEmployer(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestEmployerRemove(t *testing.T) {
	keeper, ctx := keepertest.SidelineKeeper(t)
	items := createNEmployer(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveEmployer(ctx,
			item.Index,
		)
		_, found := keeper.GetEmployer(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestEmployerGetAll(t *testing.T) {
	keeper, ctx := keepertest.SidelineKeeper(t)
	items := createNEmployer(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllEmployer(ctx)),
	)
}
