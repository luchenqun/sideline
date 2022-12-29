package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgDoTask = "do_task"

var _ sdk.Msg = &MsgDoTask{}

func NewMsgDoTask(creator string, id uint64) *MsgDoTask {
	return &MsgDoTask{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgDoTask) Route() string {
	return RouterKey
}

func (msg *MsgDoTask) Type() string {
	return TypeMsgDoTask
}

func (msg *MsgDoTask) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDoTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDoTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
