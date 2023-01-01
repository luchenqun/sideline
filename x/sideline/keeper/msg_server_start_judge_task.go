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
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %d is not exist", msg.Id)
	}

	// 只有开发者或者雇佣者才能发起仲裁
	if !(task.Employer == msg.Creator || task.Developer == msg.Creator) {
		return nil, errors.Wrapf(types.ErrPermission, "you are not a stakeholder, employer = %s, developer = %s, creator = %s", task.Employer, task.Developer, msg.Creator)
	}

	// 雇佣者可以在开发者提交任务之后，可以发起仲裁
	if task.Employer == msg.Creator {
		if !(task.Status == types.TaskStatusUndone || task.Status == types.TaskStatusSubmited) {
			return nil, errors.Wrapf(types.ErrTaskStatus, "employer = %s status = %d forbid start judge task", msg.Creator, task.Status)
		}
	}

	// 开发者在提交任务之后，如果雇佣者不认可开发者结果，可以发起仲裁
	if task.Developer == msg.Creator {
		if task.Status != types.TaskStatusUndone {
			return nil, errors.Wrapf(types.ErrTaskStatus, "developer = %s status = %d forbid start judge task", msg.Creator, task.Status)
		}
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
