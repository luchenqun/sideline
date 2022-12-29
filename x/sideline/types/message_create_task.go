package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateTask = "create_task"

var _ sdk.Msg = &MsgCreateTask{}

func NewMsgCreateTask(creator string, title string, description string, remuneration string, deposit string, collateral string, deadline uint64) *MsgCreateTask {
	return &MsgCreateTask{
		Creator:      creator,
		Title:        title,
		Description:  description,
		Remuneration: remuneration,
		Deposit:      deposit,
		Collateral:   collateral,
		Deadline:     deadline,
	}
}

func (msg *MsgCreateTask) Route() string {
	return RouterKey
}

func (msg *MsgCreateTask) Type() string {
	return TypeMsgCreateTask
}

func (msg *MsgCreateTask) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateTask) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateTask) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
