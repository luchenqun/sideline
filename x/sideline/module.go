package sideline

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	// this line is used by starport scaffolding # 1

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"sideline/x/sideline/client/cli"
	"sideline/x/sideline/keeper"
	"sideline/x/sideline/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// ----------------------------------------------------------------------------
// AppModuleBasic
// ----------------------------------------------------------------------------

// AppModuleBasic implements the AppModuleBasic interface that defines the independent methods a Cosmos SDK module needs to implement.
type AppModuleBasic struct {
	cdc codec.BinaryCodec
}

func NewAppModuleBasic(cdc codec.BinaryCodec) AppModuleBasic {
	return AppModuleBasic{cdc: cdc}
}

// Name returns the name of the module as a string
func (AppModuleBasic) Name() string {
	return types.ModuleName
}

// RegisterLegacyAminoCodec registers the amino codec for the module, which is used to marshal and unmarshal structs to/from []byte in order to persist them in the module's KVStore
func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

// RegisterInterfaces registers a module's interface types and their concrete implementations as proto.Message
func (a AppModuleBasic) RegisterInterfaces(reg cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(reg)
}

// DefaultGenesis returns a default GenesisState for the module, marshalled to json.RawMessage. The default GenesisState need to be defined by the module developer and is primarily used for testing
func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesis())
}

// ValidateGenesis used to validate the GenesisState, given in its json.RawMessage form
func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var genState types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &genState); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return genState.Validate()
}

// RegisterGRPCGatewayRoutes registers the gRPC Gateway routes for the module
func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx client.Context, mux *runtime.ServeMux) {
	types.RegisterQueryHandlerClient(context.Background(), mux, types.NewQueryClient(clientCtx))
}

// GetTxCmd returns the root Tx command for the module. The subcommands of this root command are used by end-users to generate new transactions containing messages defined in the module
func (a AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}

// GetQueryCmd returns the root query command for the module. The subcommands of this root command are used by end-users to generate new queries to the subset of the state defined by the module
func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd(types.StoreKey)
}

// ----------------------------------------------------------------------------
// AppModule
// ----------------------------------------------------------------------------

// AppModule implements the AppModule interface that defines the inter-dependent methods that modules need to implement
type AppModule struct {
	AppModuleBasic

	keeper        keeper.Keeper
	accountKeeper types.AccountKeeper
	bankKeeper    types.BankKeeper
	stakingKeeper types.StakingKeeper
}

func NewAppModule(
	cdc codec.Codec,
	keeper keeper.Keeper,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
) AppModule {
	return AppModule{
		AppModuleBasic: NewAppModuleBasic(cdc),
		keeper:         keeper,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
		stakingKeeper:  stakingKeeper,
	}
}

// Deprecated: use RegisterServices
func (am AppModule) Route() sdk.Route { return sdk.Route{} }

// Deprecated: use RegisterServices
func (AppModule) QuerierRoute() string { return types.RouterKey }

// Deprecated: use RegisterServices
func (am AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	return nil
}

// RegisterServices registers a gRPC query service to respond to the module-specific gRPC queries
func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	types.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}

// RegisterInvariants registers the invariants of the module. If an invariant deviates from its predicted value, the InvariantRegistry triggers appropriate logic (most often the chain will be halted)
func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

// InitGenesis performs the module's genesis initialization. It returns no validator updates.
func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, gs json.RawMessage) []abci.ValidatorUpdate {
	var genState types.GenesisState
	// Initialize global index to index in genesis state
	cdc.MustUnmarshalJSON(gs, &genState)

	InitGenesis(ctx, am.keeper, genState)

	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns the module's exported genesis state as raw JSON bytes.
func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genState := ExportGenesis(ctx, am.keeper)
	return cdc.MustMarshalJSON(genState)
}

// ConsensusVersion is a sequence number for state-breaking change of the module. It should be incremented on each consensus-breaking change introduced by the module. To avoid wrong/empty versions, the initial version should be set to 1
func (AppModule) ConsensusVersion() uint64 { return 1 }

// BeginBlock contains the logic that is automatically triggered at the beginning of each block
func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

// EndBlock contains the logic that is automatically triggered at the end of each block
func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	k := am.keeper
	tasks := k.GetAllTask(ctx)
	blockHeight := uint64(ctx.BlockHeight())
	for _, task := range tasks {
		// 开发者失败了
		if task.Status == types.TaskStatusDoing && blockHeight > task.Deadline {
			// 扣除开发者抵押物，把保证金以及酬金返回给雇佣者
			collateral, _ := sdk.ParseCoinNormalized(task.Collateral)
			remuneration, _ := sdk.ParseCoinNormalized(task.Remuneration)
			deposit, _ := sdk.ParseCoinNormalized(task.Deposit)

			employer, _ := sdk.AccAddressFromBech32(task.Employer)
			am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, employer, sdk.Coins{collateral.Add(remuneration).Add(deposit)})

			task.Status = types.TaskStatusFail

			k.SetTask(ctx, task)

			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.EventTypeFailTask,
					sdk.NewAttribute(types.AttributeKeyTaskId, strconv.FormatUint(task.Id, 10)),
				),
			})
		}

		// 超过规定高度雇佣者未确认，直接成功
		if task.Status == types.TaskStatusSubmitted && blockHeight > task.DeliverHeight+k.MinConfirmSubmitHeight(ctx) {
			// 把酬金给到开发者，抵押物返回给开发者
			collateral, _ := sdk.ParseCoinNormalized(task.Collateral)
			remuneration, _ := sdk.ParseCoinNormalized(task.Remuneration)

			developer, _ := sdk.AccAddressFromBech32(task.Developer)

			am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, developer, sdk.Coins{collateral.Add(remuneration)})

			// 把酬金的部分作为手续费分发给验证者，奖励验证者维护节点、参与仲裁
			commission := sdk.NewCoin(types.SidelineDenom, remuneration.Amount.QuoRaw(100).MulRaw(int64(k.ValidatorCommission(ctx))))
			validators := am.stakingKeeper.GetLastValidators(ctx)
			average := sdk.NewCoin(types.SidelineDenom, commission.Amount.QuoRaw(int64(len(validators)))) // 每个验证人能分到多少
			for i, validator := range validators {
				validatorAddress := sdk.AccAddress(validator.GetOperator())
				curCoin := average // 默认分平均的
				if i == len(validators)-1 {
					curCoin = commission // 处理除不完的情况，除不完就把剩下的给最后一个
				} else {
					commission = commission.Sub(average)
				}
				am.bankKeeper.SendCoins(ctx, developer, validatorAddress, sdk.Coins{curCoin})
			}

			// 把保证金返回给雇佣者
			deposit, _ := sdk.ParseCoinNormalized(task.Deposit)
			employer, _ := sdk.AccAddressFromBech32(task.Employer)
			am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, employer, sdk.Coins{deposit})

			task.Status = types.TaskStatusSuccess

			k.SetTask(ctx, task)

			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.EventTypeSuccessTask,
					sdk.NewAttribute(types.AttributeKeyTaskId, strconv.FormatUint(task.Id, 10)),
				),
			})
		}

		// 投票结束啦
		if task.Status == types.TaskStatusJudging && blockHeight > task.JudgeHeight+k.MinConfirmJudgeHeight(ctx) {
			// 默认开发者赢
			winner, _ := sdk.AccAddressFromBech32(task.Developer)
			task.Status = types.TaskStatusDeveloperWin

			// 因为默认是开发者赢，所以我们只要判断什么情况下雇佣者赢就完事了
			// 情况1: 原告是雇佣者，且赞同投票大（默认投票数一样时，原告赢）。情况2: 原告是开发者，给开发者投反对票多
			if (task.Accuser == task.Employer && task.VoteYes >= task.VoteNo) || (task.Accuser == task.Developer && task.VoteNo > task.VoteYes) {
				winner, _ = sdk.AccAddressFromBech32(task.Employer)
				task.Status = types.TaskStatusEmployerWin
			}

			collateral, _ := sdk.ParseCoinNormalized(task.Collateral)
			remuneration, _ := sdk.ParseCoinNormalized(task.Remuneration)
			deposit, _ := sdk.ParseCoinNormalized(task.Deposit)

			// 赢家通吃
			coins := sdk.Coins{collateral.Add(remuneration).Add(deposit)}
			am.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, winner, coins)

			k.SetTask(ctx, task)

			ctx.EventManager().EmitEvents(sdk.Events{
				sdk.NewEvent(
					types.EventTypeJudgeTask,
					sdk.NewAttribute(types.AttributeKeyTaskId, strconv.FormatUint(task.Id, 10)),
				),
			})
		}
	}
	return []abci.ValidatorUpdate{}
}
