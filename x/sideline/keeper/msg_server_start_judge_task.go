package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) StartJudgeTask(goCtx context.Context, msg *types.MsgStartJudgeTask) (*types.MsgStartJudgeTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %s is not exist", msg.Id)
	}

	if !(task.Employer == msg.Creator || task.Developer == msg.Creator) {
		return nil, errors.Wrapf(types.ErrPermission, "you are not a stakeholder, employer = %s, developer = %s, creator = %s", task.Employer, task.Developer, msg.Creator)
	}

	if !(task.Employer == msg.Creator && (task.Status == types.TaskStatusUndone || task.Status == types.TaskStatusSubmited)) {
		return nil, errors.Wrapf(types.ErrTaskStatus, "employer = %s status = %s forbid start judge task", msg.Creator, task.Status)
	}

	if !(task.Developer == msg.Creator && task.Status == types.TaskStatusUndone) {
		return nil, errors.Wrapf(types.ErrTaskStatus, "developer = %s status = %s forbid start judge task", msg.Creator, task.Status)
	}

	task.Accuser = msg.Creator
	task.Status = types.TaskStatusJudging
	task.JudgeHeight = uint64(ctx.BlockHeight())

	k.SetTask(ctx, task)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeStartJudgeTask,
			sdk.NewAttribute(types.AttributeKeyTaskId, strconv.FormatUint(task.Id, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgStartJudgeTaskResponse{}, nil
}
