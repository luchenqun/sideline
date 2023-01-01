package keeper

import (
	"context"
	"cosmossdk.io/errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) SuccessTask(goCtx context.Context, msg *types.MsgSuccessTask) (*types.MsgSuccessTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %d is not exist", msg.Id)
	}

	if task.Employer == msg.Creator {
		// 只要有开发者提交过任务，雇佣者就可以标记为成功
		if !(task.Status == types.TaskStatusSubmited || task.Status == types.TaskStatusUndone || task.Status == types.TaskStatusJudging) {
			return nil, errors.Wrapf(types.ErrTaskStatus, "nobody's doing your task")
		}
	} else if task.Developer == msg.Creator {
		// 开发者提交了任务结果，但是雇佣者不确认，超过一定高度，我们就可以发起交易确认任务已经做完了
		if !(task.Status == types.TaskStatusSubmited && uint64(ctx.BlockHeight()) > task.DeliverHeight+types.MinConfirmSubmitHeight) {
			return nil, errors.Wrapf(types.ErrTaskStatus, "It's too early to confirm")
		}
	} else {
		return nil, errors.Wrapf(types.ErrPermission, "employer = %s, developer = %s, creator = %s", task.Employer, task.Developer, msg.Creator)
	}

	// 把酬金给到开发者，抵押物返回给开发者
	collateral, err := sdk.ParseCoinNormalized(task.Collateral)
	if err != nil {
		panic(err)
	}

	remuneration, err := sdk.ParseCoinNormalized(task.Remuneration)
	if err != nil {
		panic(err)
	}
	developer, _ := sdk.AccAddressFromBech32(task.Developer)

	sdkError := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, developer, sdk.Coins{collateral.Add(remuneration)})
	if sdkError != nil {
		return nil, sdkError
	}

	// 把保证金返回给雇佣者
	deposit, err := sdk.ParseCoinNormalized(task.Deposit)
	if err != nil {
		panic(err)
	}

	employer, _ := sdk.AccAddressFromBech32(task.Employer)
	sdkError = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, employer, sdk.Coins{deposit})
	if sdkError != nil {
		return nil, sdkError
	}

	task.Status = types.TaskStatusSuccess

	k.SetTask(ctx, task)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeSuccessTask,
			sdk.NewAttribute(types.AttributeKeyTaskId, strconv.FormatUint(task.Id, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgSuccessTaskResponse{}, nil
}
