package keeper

import (
	"context"
	"cosmossdk.io/errors"
	err "errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"sideline/x/sideline/types"
)

func (k msgServer) RegistEmployer(goCtx context.Context, msg *types.MsgRegistEmployer) (*types.MsgRegistEmployerResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	_, found := k.GetEmployer(ctx, msg.Creator)
	if found {
		return &types.MsgRegistEmployerResponse{}, err.New("you have regist")
	}

	_, found = k.GetDeveloper(ctx, msg.Creator)
	if found {
		return &types.MsgRegistEmployerResponse{}, err.New("you have regist for a developer")
	}

	// 把注册费分给验证者
	registrationFee, _ := sdk.ParseCoinNormalized(k.RegistrationFee(ctx))
	address, _ := sdk.AccAddressFromBech32(msg.Creator)
	balance := k.bankKeeper.GetBalance(ctx, address, types.SidelineDenom)
	if balance.IsLT(registrationFee) {
		return nil, errors.Wrapf(types.ErrCoinAmount, "registration fee = %s, you amount = %s", registrationFee.String(), balance.String())
	}
	validators := k.stakingKeeper.GetLastValidators(ctx)
	average := sdk.NewCoin(types.SidelineDenom, registrationFee.Amount.QuoRaw(int64(len(validators)))) // 每个验证人能分到多少
	for i, validator := range validators {
		validatorAddress := sdk.AccAddress(validator.GetOperator())
		curCoin := average // 默认分平均的
		if i == len(validators)-1 {
			curCoin = registrationFee // 处理除不完的情况，除不完就把剩下的给最后一个
		} else {
			registrationFee = registrationFee.Sub(average)
		}
		sdkError := k.bankKeeper.SendCoins(ctx, address, validatorAddress, sdk.Coins{curCoin})
		if sdkError != nil {
			return nil, sdkError
		}
	}

	employer := types.Employer{
		Index:     msg.Creator, // 以用户地址为map的索引
		Name:      msg.Name,
		Introduce: msg.Introduce,
		Email:     msg.Email,
		Avatar:    msg.Avatar,
		Address:   msg.Creator,
		TaskIds:   nil,
	}

	k.SetEmployer(ctx, employer)

	ctx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			types.EventTypeRegistEmployer,
			sdk.NewAttribute(types.AttributeKeyEmployer, msg.Creator),
			sdk.NewAttribute(types.AttributeKeyRegistrationFee, registrationFee.String()),
			sdk.NewAttribute(types.AttributeKeyRegistrationAverageFee, average.String()),
		),
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
		),
	})

	return &types.MsgRegistEmployerResponse{}, nil
}
