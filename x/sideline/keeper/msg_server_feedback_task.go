package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) FeedbackTask(goCtx context.Context, msg *types.MsgFeedbackTask) (*types.MsgFeedbackTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %d is not exist", msg.Id)
	}

	// 任务为终结状态才能互相反馈
	if !(task.Status == types.TaskStatusSuccess || task.Status == types.TaskStatusFail || task.Status == types.TaskStatusEmployerWin || task.Status == types.TaskStatusDeveloperWin) {
		return nil, errors.Wrapf(types.ErrTaskStatus, "status = %d", task.Status)
	}

	// 只允许反馈一次
	if task.Employer == msg.Creator {
		if task.FeedbackByEmployer != "" {
			return nil, errors.Wrapf(types.ErrFeedback, "")
		}
		task.FeedbackByEmployer = msg.Feedback
	} else if task.Developer == msg.Creator {
		if task.FeedbackByDeveloper != "" {
			return nil, errors.Wrapf(types.ErrFeedback, "")
		}
		task.FeedbackByDeveloper = msg.Feedback
	} else {
		return nil, errors.Wrapf(types.ErrPermission, "employer = %s, developer = %s, creator = %s", task.Employer, task.Developer, msg.Creator)
	}

	k.SetTask(ctx, task)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeFeedbackTask,
			sdk.NewAttribute(types.AttributeKeyTaskId, strconv.FormatUint(task.Id, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgFeedbackTaskResponse{}, nil
}
