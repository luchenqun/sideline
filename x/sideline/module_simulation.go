package sideline

import (
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
	"sideline/testutil/sample"
	sidelinesimulation "sideline/x/sideline/simulation"
	"sideline/x/sideline/types"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = sidelinesimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgRegistEmployer = "op_weight_msg_regist_employer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegistEmployer int = 100

	opWeightMsgRegistDeveloper = "op_weight_msg_regist_developer"
	// TODO: Determine the simulation weight value
	defaultWeightMsgRegistDeveloper int = 100

	opWeightMsgCreateTask = "op_weight_msg_create_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCreateTask int = 100

	opWeightMsgDoTask = "op_weight_msg_do_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgDoTask int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	sidelineGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&sidelineGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgRegistEmployer int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegistEmployer, &weightMsgRegistEmployer, nil,
		func(_ *rand.Rand) {
			weightMsgRegistEmployer = defaultWeightMsgRegistEmployer
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegistEmployer,
		sidelinesimulation.SimulateMsgRegistEmployer(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgRegistDeveloper int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgRegistDeveloper, &weightMsgRegistDeveloper, nil,
		func(_ *rand.Rand) {
			weightMsgRegistDeveloper = defaultWeightMsgRegistDeveloper
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgRegistDeveloper,
		sidelinesimulation.SimulateMsgRegistDeveloper(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCreateTask int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCreateTask, &weightMsgCreateTask, nil,
		func(_ *rand.Rand) {
			weightMsgCreateTask = defaultWeightMsgCreateTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCreateTask,
		sidelinesimulation.SimulateMsgCreateTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgDoTask int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgDoTask, &weightMsgDoTask, nil,
		func(_ *rand.Rand) {
			weightMsgDoTask = defaultWeightMsgDoTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgDoTask,
		sidelinesimulation.SimulateMsgDoTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
