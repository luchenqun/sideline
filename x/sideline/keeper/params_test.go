package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "sideline/testutil/keeper"
	"sideline/x/sideline/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SidelineKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
	require.EqualValues(t, params.MinConfirmSubmitHeight, k.MinConfirmSubmitHeight(ctx))
	require.EqualValues(t, params.MinConfirmJudgeHeight, k.MinConfirmJudgeHeight(ctx))
}
