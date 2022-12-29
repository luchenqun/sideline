package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgStartJudgeTask = "start_judge_task"

var _ sdk.Msg = &MsgStartJudgeTask{}

func NewMsgStartJudgeTask(creator string, id uint64) *MsgStartJudgeTask {
	return &MsgStartJudgeTask{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgStartJudgeTask) Route() string {
	return RouterKey
}

func (msg *MsgStartJudgeTask) Type() string {
	return TypeMsgStartJudgeTask
}

func (msg *MsgStartJudgeTask) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgStartJudgeTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStartJudgeTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
