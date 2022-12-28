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
		DeveloperList: []types.Developer{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		TaskList: []types.Task{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TaskCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SidelineKeeper(t)
	sideline.InitGenesis(ctx, *k, genesisState)
	got := sideline.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.EmployerList, got.EmployerList)
	require.ElementsMatch(t, genesisState.DeveloperList, got.DeveloperList)
	require.ElementsMatch(t, genesisState.TaskList, got.TaskList)
	require.Equal(t, genesisState.TaskCount, got.TaskCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
