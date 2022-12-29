package keeper

import (
	"context"
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) SubmitTask(goCtx context.Context, msg *types.MsgSubmitTask) (*types.MsgSubmitTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %s is not exist", msg.Id)
	}
	if !(task.Status == types.TaskStatusDoing || task.Status == types.TaskStatusUndone) {
		return nil, errors.Wrapf(types.ErrTaskStatus, "task status = %s, forbid submit this task", task.Status)
	}

	if !(task.Developer == msg.Creator) {
		return nil, errors.Wrapf(types.ErrPermission, "developer = %s, creator = %s", task.Developer, msg.Creator)
	}

	task.Deliver = msg.Deliver
	task.Status = types.TaskStatusSubmited
	task.DeliverHeight = uint64(ctx.BlockHeight())

	k.SetTask(ctx, task)

	return &types.MsgSubmitTaskResponse{}, nil
}
