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

	opWeightMsgSubmitTask = "op_weight_msg_submit_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSubmitTask int = 100

	opWeightMsgUndoneTask = "op_weight_msg_undone_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgUndoneTask int = 100

	opWeightMsgSuccessTask = "op_weight_msg_success_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgSuccessTask int = 100

	opWeightMsgFailTask = "op_weight_msg_fail_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFailTask int = 100

	opWeightMsgStartJudgeTask = "op_weight_msg_start_judge_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgStartJudgeTask int = 100

	opWeightMsgCancelTask = "op_weight_msg_cancel_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgCancelTask int = 100

	opWeightMsgJudgeTask = "op_weight_msg_judge_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgJudgeTask int = 100

	opWeightMsgVoteTask = "op_weight_msg_vote_task"
	// TODO: Determine the simulation weight value
	defaultWeightMsgVoteTask int = 100

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

	var weightMsgSubmitTask int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSubmitTask, &weightMsgSubmitTask, nil,
		func(_ *rand.Rand) {
			weightMsgSubmitTask = defaultWeightMsgSubmitTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSubmitTask,
		sidelinesimulation.SimulateMsgSubmitTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgUndoneTask int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgUndoneTask, &weightMsgUndoneTask, nil,
		func(_ *rand.Rand) {
			weightMsgUndoneTask = defaultWeightMsgUndoneTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgUndoneTask,
		sidelinesimulation.SimulateMsgUndoneTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgSuccessTask int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgSuccessTask, &weightMsgSuccessTask, nil,
		func(_ *rand.Rand) {
			weightMsgSuccessTask = defaultWeightMsgSuccessTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgSuccessTask,
		sidelinesimulation.SimulateMsgSuccessTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgFailTask int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgFailTask, &weightMsgFailTask, nil,
		func(_ *rand.Rand) {
			weightMsgFailTask = defaultWeightMsgFailTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFailTask,
		sidelinesimulation.SimulateMsgFailTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgStartJudgeTask int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgStartJudgeTask, &weightMsgStartJudgeTask, nil,
		func(_ *rand.Rand) {
			weightMsgStartJudgeTask = defaultWeightMsgStartJudgeTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgStartJudgeTask,
		sidelinesimulation.SimulateMsgStartJudgeTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgCancelTask int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgCancelTask, &weightMsgCancelTask, nil,
		func(_ *rand.Rand) {
			weightMsgCancelTask = defaultWeightMsgCancelTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgCancelTask,
		sidelinesimulation.SimulateMsgCancelTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgJudgeTask int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgJudgeTask, &weightMsgJudgeTask, nil,
		func(_ *rand.Rand) {
			weightMsgJudgeTask = defaultWeightMsgJudgeTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgJudgeTask,
		sidelinesimulation.SimulateMsgJudgeTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	var weightMsgVoteTask int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgVoteTask, &weightMsgVoteTask, nil,
		func(_ *rand.Rand) {
			weightMsgVoteTask = defaultWeightMsgVoteTask
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgVoteTask,
		sidelinesimulation.SimulateMsgVoteTask(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
