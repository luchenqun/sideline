package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgRegistEmployer{}, "sideline/RegistEmployer", nil)
	cdc.RegisterConcrete(&MsgRegistDeveloper{}, "sideline/RegistDeveloper", nil)
	cdc.RegisterConcrete(&MsgCreateTask{}, "sideline/CreateTask", nil)
	cdc.RegisterConcrete(&MsgDoTask{}, "sideline/DoTask", nil)
	cdc.RegisterConcrete(&MsgSubmitTask{}, "sideline/SubmitTask", nil)
	cdc.RegisterConcrete(&MsgUndoneTask{}, "sideline/UndoneTask", nil)
	cdc.RegisterConcrete(&MsgSuccessTask{}, "sideline/SuccessTask", nil)
	cdc.RegisterConcrete(&MsgFailTask{}, "sideline/FailTask", nil)
	cdc.RegisterConcrete(&MsgStartJudgeTask{}, "sideline/StartJudgeTask", nil)
	cdc.RegisterConcrete(&MsgCancelTask{}, "sideline/CancelTask", nil)
	cdc.RegisterConcrete(&MsgJudgeTask{}, "sideline/JudgeTask", nil)
	cdc.RegisterConcrete(&MsgVoteTask{}, "sideline/VoteTask", nil)
	// this line is used by starport scaffolding # 2
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegistEmployer{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgRegistDeveloper{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateTask{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgDoTask{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSubmitTask{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUndoneTask{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgSuccessTask{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgFailTask{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgStartJudgeTask{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCancelTask{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgJudgeTask{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgVoteTask{},
	)
	// this line is used by starport scaffolding # 3

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	Amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
