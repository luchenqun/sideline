package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) SubmitTask(goCtx context.Context, msg *types.MsgSubmitTask) (*types.MsgSubmitTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %d is not exist", msg.Id)
	}

	// 截止日禁止提交任务
	if uint64(ctx.BlockHeight()) > task.Deadline {
		return nil, errors.Wrapf(types.ErrTaskStatus, "you submit to later")
	}

	// 第一次提交，或者雇佣者打回任务之后可以再次提交
	if !(task.Status == types.TaskStatusDoing || task.Status == types.TaskStatusUndone) {
		return nil, errors.Wrapf(types.ErrTaskStatus, "task status = %d, forbid submit this task", task.Status)
	}

	// 只能开发者本人提交任务
	if !(task.Developer == msg.Creator) {
		return nil, errors.Wrapf(types.ErrPermission, "developer = %s, creator = %s", task.Developer, msg.Creator)
	}

	task.Deliver = msg.Deliver
	task.Status = types.TaskStatusSubmitted
	task.DeliverHeight = uint64(ctx.BlockHeight())

	k.SetTask(ctx, task)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSubmitTask,
			sdk.NewAttribute(types.AttributeKeyTaskId, strconv.FormatUint(task.Id, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgSubmitTaskResponse{}, nil
}
