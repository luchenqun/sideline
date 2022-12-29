package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgSuccessTask = "success_task"

var _ sdk.Msg = &MsgSuccessTask{}

func NewMsgSuccessTask(creator string, id uint64) *MsgSuccessTask {
	return &MsgSuccessTask{
		Creator: creator,
		Id:      id,
	}
}

func (msg *MsgSuccessTask) Route() string {
	return RouterKey
}

func (msg *MsgSuccessTask) Type() string {
	return TypeMsgSuccessTask
}

func (msg *MsgSuccessTask) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgSuccessTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSuccessTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
