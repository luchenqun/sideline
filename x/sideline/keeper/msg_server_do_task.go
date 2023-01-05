package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) DoTask(goCtx context.Context, msg *types.MsgDoTask) (*types.MsgDoTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %d is not exist", msg.Id)
	}
	if task.Status != types.TaskStatusCreated {
		return nil, errors.Wrapf(types.ErrTaskStatus, "task status = %d, forbid do this task", task.Status)
	}

	// 只有开发者才能发布任务
	developer, found := k.GetDeveloper(ctx, msg.Creator)
	if !found {
		return nil, errors.Wrap(types.ErrNotRegistForDeveloper, "forbid do task")
	}

	// 抵押物
	developerAddress, _ := sdk.AccAddressFromBech32(msg.Creator)

	collateral, _ := sdk.ParseCoinNormalized(task.Collateral)

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, developerAddress, types.ModuleName, sdk.Coins{collateral})
	if sdkError != nil {
		return nil, sdkError
	}

	developer.TaskIds = append(developer.TaskIds, task.Id)

	task.Developer = msg.Creator
	task.Status = types.TaskStatusDoing

	k.SetTask(ctx, task)
	k.SetDeveloper(ctx, developer)

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

	return &types.MsgDoTaskResponse{}, nil
}
