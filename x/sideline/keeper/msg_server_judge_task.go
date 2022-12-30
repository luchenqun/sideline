package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) JudgeTask(goCtx context.Context, msg *types.MsgJudgeTask) (*types.MsgJudgeTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %s is not exist", msg.Id)
	}
	// 任何人都可以结束仲裁
	// 只有进入仲裁的任务，才有可能给仲裁结果
	if task.Status != types.TaskStatusJudging {
		return nil, errors.Wrapf(types.ErrTaskStatus, "status = %s forbid success task", task.Status)
	}

	// 时间未到
	if uint64(ctx.BlockHeight()) < task.JudgeHeight+types.MinConfirmJudgeHeight {
		return nil, errors.Wrapf(types.ErrTime, "current height = %d, right height = %d", ctx.BlockHeight(), task.JudgeHeight+types.MinConfirmJudgeHeight)
	}

	// 默认开发者赢
	winner, _ := sdk.AccAddressFromBech32(task.Developer)
	task.Status = types.TaskStatusEmployerWin

	// 因为默认是开发者赢，所以我们只要判断什么情况下雇佣者赢就完事了
	// 情况1: 原告是雇佣者，且赞同投票大（默认投票数一样时，原告赢）。情况2: 原告是开发者，给开发者投反对票多
	if (task.Accuser == task.Employer && task.VoteYes >= task.VoteNo) || (task.Accuser == task.Developer && task.VoteNo > task.VoteYes) {
		winner, _ = sdk.AccAddressFromBech32(task.Employer)
		task.Status = types.TaskStatusEmployerWin
	}

	collateral, err := sdk.ParseCoinNormalized(task.Collateral)
	if err != nil {
		panic(err)
	}

	remuneration, err := sdk.ParseCoinNormalized(task.Remuneration)
	if err != nil {
		panic(err)
	}

	deposit, err := sdk.ParseCoinNormalized(task.Deposit)
	if err != nil {
		panic(err)
	}

	// 赢家通吃
	coins := sdk.Coins{collateral.Add(remuneration).Add(deposit)}
	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, winner, coins)
	if sdkError != nil {
		return nil, sdkError
	}

	k.SetTask(ctx, task)

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

	return &types.MsgJudgeTaskResponse{}, nil
}
