package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFeedbackTask = "feedback_task"

var _ sdk.Msg = &MsgFeedbackTask{}

func NewMsgFeedbackTask(creator string, id uint64) *MsgFeedbackTask {
	return &MsgFeedbackTask{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgFeedbackTask) Route() string {
	return RouterKey
}

func (msg *MsgFeedbackTask) Type() string {
	return TypeMsgFeedbackTask
}

func (msg *MsgFeedbackTask) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFeedbackTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFeedbackTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
