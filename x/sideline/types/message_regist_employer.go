package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegistEmployer = "regist_employer"

var _ sdk.Msg = &MsgRegistEmployer{}

func NewMsgRegistEmployer(creator string, name string, introduce string, email string, avatar string) *MsgRegistEmployer {
	return &MsgRegistEmployer{
		Creator:   creator,
		Name:      name,
		Introduce: introduce,
		Email:     email,
		Avatar:    avatar,
	}
}

func (msg *MsgRegistEmployer) Route() string {
	return RouterKey
}

func (msg *MsgRegistEmployer) Type() string {
	return TypeMsgRegistEmployer
}

func (msg *MsgRegistEmployer) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegistEmployer) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegistEmployer) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
