package keeper

import (
	"context"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) RegistEmployer(goCtx context.Context, msg *types.MsgRegistEmployer) (*types.MsgRegistEmployerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetEmployer(ctx, msg.Creator)
	if found {
		return &types.MsgRegistEmployerResponse{}, errors.New("you have regist")
	}

	employer := types.Employer{
		Index:     msg.Creator, // 以用户地址为map的索引
		Name:      msg.Name,
		Introduce: msg.Introduce,
		Email:     msg.Email,
		Avatar:    msg.Avatar,
		Address:   msg.Creator,
		TaskCount: 0,
		TaskFail:  0,
		Feedbacks: nil,
	}

	k.SetEmployer(ctx, employer)

	return &types.MsgRegistEmployerResponse{}, nil
}
