package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) SuccessTask(goCtx context.Context, msg *types.MsgSuccessTask) (*types.MsgSuccessTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %d is not exist", msg.Id)
	}

	if task.Employer == msg.Creator {
		// 只要有开发者提交过任务，雇佣者就可以标记为成功
		if !(task.Status == types.TaskStatusSubmitted || task.Status == types.TaskStatusUndone || task.Status == types.TaskStatusJudging) {
			return nil, errors.Wrapf(types.ErrTaskStatus, "nobody's doing your task")
		}
	} else if task.Developer == msg.Creator {
		// 开发者提交了任务结果，但是雇佣者不确认，超过一定高度，我们就可以发起交易确认任务已经做完了
		if !(task.Status == types.TaskStatusSubmitted && uint64(ctx.BlockHeight()) > task.DeliverHeight+k.MinConfirmSubmitHeight(ctx)) {
			return nil, errors.Wrapf(types.ErrTaskStatus, "It's too early to confirm")
		}
	} else {
		return nil, errors.Wrapf(types.ErrPermission, "employer = %s, developer = %s, creator = %s", task.Employer, task.Developer, msg.Creator)
	}

	// 把酬金给到开发者，抵押物返回给开发者
	collateral, _ := sdk.ParseCoinNormalized(task.Collateral)
	remuneration, _ := sdk.ParseCoinNormalized(task.Remuneration)

	developer, _ := sdk.AccAddressFromBech32(task.Developer)

	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, developer, sdk.Coins{collateral.Add(remuneration)})
	if sdkError != nil {
		return nil, sdkError
	}

	// 把酬金的部分作为手续费分发给验证者，奖励验证者维护节点、参与仲裁
	commission := sdk.NewCoin(types.SidelineDenom, remuneration.Amount.QuoRaw(100).MulRaw(int64(k.ValidatorCommission(ctx))))
	validators := k.stakingKeeper.GetLastValidators(ctx)
	average := sdk.NewCoin(types.SidelineDenom, commission.Amount.QuoRaw(int64(len(validators)))) // 每个验证人能分到多少
	for i, validator := range validators {
		validatorAddress := sdk.AccAddress(validator.GetOperator())
		curCoin := average // 默认分平均的
		if i == len(validators)-1 {
			curCoin = commission // 处理除不完的情况，除不完就把剩下的给最后一个
		} else {
			commission = commission.Sub(average)
		}
		sdkError := k.bankKeeper.SendCoins(ctx, developer, validatorAddress, sdk.Coins{curCoin})
		if sdkError != nil {
			return nil, sdkError
		}
	}

	// 把保证金返回给雇佣者。如果是进入了仲裁，那么相当于雇佣者认输，保证金给到开发者
	deposit, _ := sdk.ParseCoinNormalized(task.Deposit)
	address, _ := sdk.AccAddressFromBech32(task.Employer)
	if task.Status == types.TaskStatusJudging {
		address = developer
	}
	sdkError = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, address, sdk.Coins{deposit})
	if sdkError != nil {
		return nil, sdkError
	}

	task.Status = types.TaskStatusSuccess

	k.SetTask(ctx, task)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSuccessTask,
			sdk.NewAttribute(types.AttributeKeyTaskId, strconv.FormatUint(task.Id, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgSuccessTaskResponse{}, nil
}
