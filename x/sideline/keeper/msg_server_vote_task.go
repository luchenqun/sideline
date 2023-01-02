package keeper

import (
	"context"
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
	"strconv"
	"strings"
)

func (k msgServer) VoteTask(goCtx context.Context, msg *types.MsgVoteTask) (*types.MsgVoteTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %d is not exist", msg.Id)
	}

	// 只有进入仲裁状态的才允许投票
	if !(task.Status == types.TaskStatusJudging) {
		return nil, errors.Wrapf(types.ErrTaskStatus, "task status = %d, forbid undone this task", task.Status)
	}

	// 必须在仲裁周期内
	if task.JudgeHeight+k.MinConfirmJudgeHeight(ctx) < uint64(ctx.BlockHeight()) {
		return nil, errors.Wrapf(types.ErrTime, "current height = %d, right height = %d", ctx.BlockHeight(), task.JudgeHeight+k.MinConfirmJudgeHeight(ctx))
	}

	// 只允许验证人投票
	accAddress, _ := sdk.AccAddressFromBech32(msg.Creator)
	valAddress := sdk.ValAddress(accAddress)
	_, found = k.stakingKeeper.GetValidator(ctx, valAddress)
	if !found {
		return nil, errors.Wrapf(types.ErrPermission, "%s is not a validator", valAddress.String())
	}

	// 不能重复投票
	for _, account := range task.VotedAccounts {
		if strings.Contains(account, msg.Creator) {
			return nil, errors.Wrapf(types.ErrPermission, "you have voted")
		}
	}

	if msg.Option == 0 {
		task.VoteYes += 1
	} else {
		task.VoteNo += 1
	}

	task.VotedAccounts = append(task.VotedAccounts, msg.Creator+"|"+strconv.FormatUint(msg.Option, 10))

	k.SetTask(ctx, task)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeVoteTask,
			sdk.NewAttribute(types.AttributeKeyTaskId, strconv.FormatUint(task.Id, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgVoteTaskResponse{}, nil
}
