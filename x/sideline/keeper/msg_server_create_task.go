package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) CreateTask(goCtx context.Context, msg *types.MsgCreateTask) (*types.MsgCreateTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task := types.Task{
		Title:        msg.Title,
		Description:  msg.Description,
		Remuneration: msg.Remuneration,
		Deposit:      msg.Deposit,
		Collateral:   msg.Collateral,
		Employer:     msg.Creator,
		Deadline:     msg.Deadline,
		Status:       types.TaskStatusCreated,
	}

	// 只有雇主才能发布任务
	employer, found := k.GetEmployer(ctx, msg.Creator)
	if !found {
		return nil, errors.Wrap(types.ErrNotRegistForEmployer, "forbid create task")
	}
	// 任务结束快高不能低于现在的快高
	if uint64(ctx.BlockHeight()) >= msg.Deadline {
		return nil, errors.Wrapf(types.ErrTime, "current height = %d, deadline height = %d", ctx.BlockHeight(), msg.Deadline)
	}

	employerAddress, _ := sdk.AccAddressFromBech32(msg.Creator)

	remuneration, err := sdk.ParseCoinNormalized(msg.Remuneration)
	if err != nil {
		panic(err)
	}
	deposit, err := sdk.ParseCoinNormalized(msg.Deposit)
	if err != nil {
		panic(err)
	}
	coin := remuneration.Add(deposit)

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, employerAddress, types.ModuleName, sdk.Coins{coin})
	if sdkError != nil {
		return nil, sdkError
	}

	Id := k.AppendTask(ctx, task)

	employer.TaskIds = append(employer.TaskIds, Id)
	k.SetEmployer(ctx, employer)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeCreateTask,
			sdk.NewAttribute(types.AttributeKeyTaskId, strconv.FormatUint(task.Id, 10)),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgCreateTaskResponse{Id: Id}, nil
}
