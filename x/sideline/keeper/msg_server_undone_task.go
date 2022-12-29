package keeper

import (
	"context"
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) UndoneTask(goCtx context.Context, msg *types.MsgUndoneTask) (*types.MsgUndoneTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %s is not exist", msg.Id)
	}
	if !(task.Status == types.TaskStatusSubmited) {
		return nil, errors.Wrapf(types.ErrTaskStatus, "task status = %s, forbid undone this task", task.Status)
	}

	if !(task.Employer == msg.Creator) {
		return nil, errors.Wrapf(types.ErrPermission, "employer = %s, creator = %s", task.Employer, msg.Creator)
	}

	task.Status = types.TaskStatusUndone

	k.SetTask(ctx, task)

	return &types.MsgUndoneTaskResponse{}, nil
}
