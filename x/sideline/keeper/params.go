package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

// GetParams get all parameters as types.Params
func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.MinConfirmSubmitHeight(ctx),
		k.MinConfirmJudgeHeight(ctx),
	)
}

// SetParams set the params
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramstore.SetParamSet(ctx, &params)
}

// MinConfirmSubmitHeight returns the MinConfirmSubmitHeight param
func (k Keeper) MinConfirmSubmitHeight(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMinConfirmSubmitHeight, &res)
	return
}

// MinConfirmJudgeHeight returns the MinConfirmJudgeHeight param
func (k Keeper) MinConfirmJudgeHeight(ctx sdk.Context) (res uint64) {
	k.paramstore.Get(ctx, types.KeyMinConfirmJudgeHeight, &res)
	return
}
