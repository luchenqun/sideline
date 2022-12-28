package sideline_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "sideline/testutil/keeper"
	"sideline/testutil/nullify"
	"sideline/x/sideline"
	"sideline/x/sideline/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		EmployerList: []types.Employer{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SidelineKeeper(t)
	sideline.InitGenesis(ctx, *k, genesisState)
	got := sideline.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.EmployerList, got.EmployerList)
	// this line is used by starport scaffolding # genesis/test/assert
}
