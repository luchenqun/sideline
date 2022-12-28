package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) RegistEmployer(goCtx context.Context, msg *types.MsgRegistEmployer) (*types.MsgRegistEmployerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgRegistEmployerResponse{}, nil
}
