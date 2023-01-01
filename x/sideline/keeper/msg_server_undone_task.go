package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) UndoneTask(goCtx context.Context, msg *types.MsgUndoneTask) (*types.MsgUndoneTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %d is not exist", msg.Id)
	}
	if !(task.Status == types.TaskStatusSubmited) {
		return nil, errors.Wrapf(types.ErrTaskStatus, "task status = %d, forbid undone this task", task.Status)
	}

	if !(task.Employer == msg.Creator) {
		return nil, errors.Wrapf(types.ErrPermission, "employer = %s, creator = %s", task.Employer, msg.Creator)
	}

	task.Status = types.TaskStatusUndone

	k.SetTask(ctx, task)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeUndoneTask,
			sdk.NewAttribute(types.AttributeKeyTaskId, strconv.FormatUint(task.Id, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgUndoneTaskResponse{}, nil
}
