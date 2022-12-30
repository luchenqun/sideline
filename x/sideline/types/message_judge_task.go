package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgJudgeTask = "judge_task"

var _ sdk.Msg = &MsgJudgeTask{}

func NewMsgJudgeTask(creator string, id uint64) *MsgJudgeTask {
	return &MsgJudgeTask{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgJudgeTask) Route() string {
	return RouterKey
}

func (msg *MsgJudgeTask) Type() string {
	return TypeMsgJudgeTask
}

func (msg *MsgJudgeTask) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgJudgeTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgJudgeTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
