package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFailTask = "fail_task"

var _ sdk.Msg = &MsgFailTask{}

func NewMsgFailTask(creator string, id uint64) *MsgFailTask {
	return &MsgFailTask{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgFailTask) Route() string {
	return RouterKey
}

func (msg *MsgFailTask) Type() string {
	return TypeMsgFailTask
}

func (msg *MsgFailTask) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFailTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFailTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
