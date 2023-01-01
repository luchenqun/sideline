package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) RegistDeveloper(goCtx context.Context, msg *types.MsgRegistDeveloper) (*types.MsgRegistDeveloperResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetDeveloper(ctx, msg.Creator)
	if found {
		return &types.MsgRegistDeveloperResponse{}, errors.New("you have regist")
	}

	_, found = k.GetEmployer(ctx, msg.Creator)
	if found {
		return &types.MsgRegistDeveloperResponse{}, errors.New("you have regist for a employer")
	}

	developer := types.Developer{
		Index:     msg.Creator, // 以开发者地址为map的索引
		Name:      msg.Name,
		Introduce: msg.Introduce,
		Email:     msg.Email,
		Avatar:    msg.Avatar,
		Address:   msg.Creator,
		Education: msg.Education,
		Major:     msg.Major,
		Skills:    msg.Skills,
		TaskIds:   nil,
	}

	k.SetDeveloper(ctx, developer)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRegistDeveloper,
			sdk.NewAttribute(types.AttributeKeyDeveloper, msg.Creator),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgRegistDeveloperResponse{}, nil
}
