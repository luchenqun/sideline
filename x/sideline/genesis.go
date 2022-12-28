package sideline

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/keeper"
	"sideline/x/sideline/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the employer
	for _, elem := range genState.EmployerList {
		k.SetEmployer(ctx, elem)
	}
	// Set all the developer
	for _, elem := range genState.DeveloperList {
		k.SetDeveloper(ctx, elem)
	}
	// Set all the task
	for _, elem := range genState.TaskList {
		k.SetTask(ctx, elem)
	}

	// Set task count
	k.SetTaskCount(ctx, genState.TaskCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.EmployerList = k.GetAllEmployer(ctx)
	genesis.DeveloperList = k.GetAllDeveloper(ctx)
	genesis.TaskList = k.GetAllTask(ctx)
	genesis.TaskCount = k.GetTaskCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
