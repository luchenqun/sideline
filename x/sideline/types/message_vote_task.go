package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgVoteTask = "vote_task"

var _ sdk.Msg = &MsgVoteTask{}

func NewMsgVoteTask(creator string, id uint64, option uint64) *MsgVoteTask {
	return &MsgVoteTask{
		Creator: creator,
		Id:      id,
		Option:  option,
	}
}

func (msg *MsgVoteTask) Route() string {
	return RouterKey
}

func (msg *MsgVoteTask) Type() string {
	return TypeMsgVoteTask
}

func (msg *MsgVoteTask) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgVoteTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVoteTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
