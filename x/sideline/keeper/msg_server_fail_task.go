package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) FailTask(goCtx context.Context, msg *types.MsgFailTask) (*types.MsgFailTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %s is not exist", msg.Id)
	}

	if task.Employer == msg.Creator {
		// 雇佣者只有在任务未完成的情况下有权告知任务失败
		if !(uint64(ctx.BlockHeight()) > task.Deadline && task.Status == types.TaskStatusDoing) {
			return nil, errors.Wrapf(types.ErrTaskStatus, "It's too early to fail a task")
		}
	} else if task.Developer == msg.Creator {
		// 开发者能置为失败的状态
		if !(task.Status == types.TaskStatusDoing || task.Status == types.TaskStatusSubmited || task.Status == types.TaskStatusUndone || task.Status == types.TaskStatusJudging) {
			return nil, errors.Wrapf(types.ErrTaskStatus, "status = %s forbid fail task", task.Status)
		}
	} else {
		return nil, errors.Wrapf(types.ErrPermission, "employer = %s, developer = %s, creator = %s", task.Employer, task.Developer, msg.Creator)
	}

	// 扣除开发者抵押物，把保证金以及酬金返回给雇佣者
	collateral, err := sdk.ParseCoinNormalized(task.Collateral)
	if err != nil {
		panic(err)
	}

	remuneration, err := sdk.ParseCoinNormalized(task.Remuneration)
	if err != nil {
		panic(err)
	}

	deposit, err := sdk.ParseCoinNormalized(task.Deposit)
	if err != nil {
		panic(err)
	}

	employer, _ := sdk.AccAddressFromBech32(task.Employer)
	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, employer, sdk.Coins{collateral.Add(remuneration).Add(deposit)})
	if sdkError != nil {
		return nil, sdkError
	}

	task.Status = types.TaskStatusFail

	k.SetTask(ctx, task)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeDoTask,
			sdk.NewAttribute(types.AttributeKeyTaskId, strconv.FormatUint(task.Id, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgFailTaskResponse{}, nil
}
