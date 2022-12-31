package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) CancelTask(goCtx context.Context, msg *types.MsgCancelTask) (*types.MsgCancelTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %s is not exist", msg.Id)
	}
	// 没人接任务才能取消任务
	if task.Status != types.TaskStatusCreated {
		return nil, errors.Wrapf(types.ErrTaskStatus, "task status = %s, forbid cancel task", task.Status)
	}

	// 只允许自己取消任务
	if msg.Creator != task.Employer {
		return nil, errors.Wrapf(types.ErrPermission, "employer = %s, creator = %s", task.Employer, msg.Creator)
	}

	// 退回所有的金额
	employer, _ := sdk.AccAddressFromBech32(msg.Creator)
	remuneration, err := sdk.ParseCoinNormalized(task.Remuneration)
	if err != nil {
		panic(err)
	}
	deposit, err := sdk.ParseCoinNormalized(task.Deposit)
	if err != nil {
		panic(err)
	}
	coin := remuneration.Add(deposit)
	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, employer, sdk.Coins{coin})
	if sdkError != nil {
		return nil, sdkError
	}

	task.Status = types.TaskStatusCanced

	k.SetTask(ctx, task)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCancelTask,
			sdk.NewAttribute(types.AttributeKeyTaskId, strconv.FormatUint(task.Id, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgCancelTaskResponse{}, nil
}
