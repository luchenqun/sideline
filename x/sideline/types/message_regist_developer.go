package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgRegistDeveloper = "regist_developer"

var _ sdk.Msg = &MsgRegistDeveloper{}

func NewMsgRegistDeveloper(creator string, name string, introduce string, email string, avatar string, education string, major string, skills []string) *MsgRegistDeveloper {
	return &MsgRegistDeveloper{
		Creator:   creator,
		Name:      name,
		Introduce: introduce,
		Email:     email,
		Avatar:    avatar,
		Education: education,
		Major:     major,
		Skills:    skills,
	}
}

func (msg *MsgRegistDeveloper) Route() string {
	return RouterKey
}

func (msg *MsgRegistDeveloper) Type() string {
	return TypeMsgRegistDeveloper
}

func (msg *MsgRegistDeveloper) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgRegistDeveloper) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRegistDeveloper) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
