package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSubmitTask = "submit_task"

var _ sdk.Msg = &MsgSubmitTask{}

func NewMsgSubmitTask(creator string, id uint64, deliver string) *MsgSubmitTask {
	return &MsgSubmitTask{
		Creator: creator,
		Id:      id,
		Deliver: deliver,
	}
}

func (msg *MsgSubmitTask) Route() string {
	return RouterKey
}

func (msg *MsgSubmitTask) Type() string {
	return TypeMsgSubmitTask
}

func (msg *MsgSubmitTask) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSubmitTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSubmitTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
