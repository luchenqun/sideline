package keeper

import (
	"context"
	"github.com/cosmos/cosmos-sdk/types/errors"

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
		Developer:    "",
		Deadline:     msg.Deadline,
		Status:       types.TaskStatusCreated,
	}

	// 只有雇主才能发布任务
	_, found := k.GetEmployer(ctx, msg.Creator)
	if !found {
		return &types.MsgCreateTaskResponse{Id: 0}, errors.Wrap(types.ErrNotRegistForEmployer, "forbid create task")
	}
	// @todo 任务结束快高不能低于现在的快高

	employer, _ := sdk.AccAddressFromBech32(msg.Creator)

	remuneration, err := sdk.ParseCoinNormalized(msg.Remuneration)
	if err != nil {
		panic(err)
	}
	deposit, err := sdk.ParseCoinNormalized(msg.Deposit)
	if err != nil {
		panic(err)
	}
	coin := remuneration.Add(deposit)

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, employer, types.ModuleName, sdk.Coins{coin})
	if sdkError != nil {
		return nil, sdkError
	}

	Id := k.AppendTask(ctx, task)

	return &types.MsgCreateTaskResponse{Id: Id}, nil
}
