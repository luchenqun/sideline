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

	developer := types.Developer{
		Index:       msg.Creator, // 以开发者地址为map的索引
		Name:        msg.Name,
		Introduce:   msg.Introduce,
		Email:       msg.Email,
		Avatar:      msg.Avatar,
		Address:     msg.Creator,
		Education:   msg.Education,
		Major:       msg.Major,
		Skills:      msg.Skills,
		TaskSuccess: 0,
		TaskFail:    0,
		Feedbacks:   nil,
	}

	k.SetDeveloper(ctx, developer)

	return &types.MsgRegistDeveloperResponse{}, nil
}