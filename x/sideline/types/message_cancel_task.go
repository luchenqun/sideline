package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCancelTask = "cancel_task"

var _ sdk.Msg = &MsgCancelTask{}

func NewMsgCancelTask(creator string, id uint64) *MsgCancelTask {
	return &MsgCancelTask{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgCancelTask) Route() string {
	return RouterKey
}

func (msg *MsgCancelTask) Type() string {
	return TypeMsgCancelTask
}

func (msg *MsgCancelTask) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCancelTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCancelTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
