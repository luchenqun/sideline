package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgUndoneTask = "undone_task"

var _ sdk.Msg = &MsgUndoneTask{}

func NewMsgUndoneTask(creator string, id uint64) *MsgUndoneTask {
	return &MsgUndoneTask{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgUndoneTask) Route() string {
	return RouterKey
}

func (msg *MsgUndoneTask) Type() string {
	return TypeMsgUndoneTask
}

func (msg *MsgUndoneTask) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUndoneTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUndoneTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
