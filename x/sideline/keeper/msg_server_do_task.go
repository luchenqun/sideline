package keeper

import (
	"context"
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) DoTask(goCtx context.Context, msg *types.MsgDoTask) (*types.MsgDoTaskResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	task, found := k.GetTask(ctx, msg.Id)
	if !found {
		return nil, errors.Wrapf(types.ErrTaskID, "task id = %s is not exist", msg.Id)
	}
	if task.Status != types.TaskStatusCreated {
		return nil, errors.Wrapf(types.ErrTaskStatus, "task status = %s, forbid do this task", task.Status)
	}

	developer, _ := sdk.AccAddressFromBech32(msg.Creator)

	collateral, err := sdk.ParseCoinNormalized(task.Collateral)
	if err != nil {
		panic(err)
	}

	sdkError := k.bankKeeper.SendCoinsFromAccountToModule(ctx, developer, types.ModuleName, sdk.Coins{collateral})
	if sdkError != nil {
		return nil, sdkError
	}

	task.Developer = msg.Creator
	task.Status = types.TaskStatusDoing

	k.SetTask(ctx, task)

	return &types.MsgDoTaskResponse{}, nil
}
